#!/usr/bin/env python3
"""Preview or explicitly reset Kencleng resources in the current Podman store."""

import argparse
import json
import subprocess
import sys

PROJECTS = {"kencleng", "kencleng-apps", "kencleng-storage"}
PROJECT_LABELS = ("io.podman.compose.project", "com.docker.compose.project")
LEGACY_VOLUMES = {
    "kencleng_kencleng_pgdata",
    "kencleng_kencleng_miniodata",
    "kencleng-storage_kencleng_miniodata",
}
LEGACY_NETWORKS = {f"{project}_default" for project in PROJECTS}
LEGACY_PODS = {f"pod_{project}" for project in PROJECTS}


def field(resource, name, default=""):
    return next(
        (value for key, value in resource.items() if key.lower() == name.lower()),
        default,
    )


def resource_id(resource):
    return field(resource, "id")


def resource_name(resource):
    names = field(resource, "name") or field(resource, "names")
    return names[0] if isinstance(names, list) else names


def project_owners(resource):
    labels = field(resource, "labels", {}) or {}
    return {labels[key] for key in PROJECT_LABELS if labels.get(key)}


def owned(resource, legacy_names=()):
    owners = project_owners(resource)
    if owners:
        # A foreign project label never qualifies through a familiar name.
        return owners <= PROJECTS
    return resource_name(resource) in legacy_names


def run(*args):
    result = subprocess.run(
        ["podman", *args], capture_output=True, text=True, check=False
    )
    if result.returncode:
        message = f"podman {' '.join(args)} gagal:\n{result.stderr.strip()}"
        if "rootless netns: kill network process: permission denied" in result.stderr:
            message += (
                "\nCleanup jaringan rootless ditolak oleh host. Periksa log AppArmor/pasta; "
                "lihat bagian troubleshooting reset di docs/project/kencleng-repo-setup.md. "
                "Reset belum selesai dan sebagian container mungkin sudah dihapus. "
                "Script tidak memakai sudo, mengubah policy, atau mengabaikan error ini. "
                "Setelah host diperbaiki, jalankan preview lagi sebelum --apply."
            )
        raise RuntimeError(message)
    return result.stdout


def read_json(*args):
    return json.loads(run(*args))


def inspect_pod(pod):
    return read_json("pod", "inspect", resource_id(pod))[0]


def members(pod):
    infra = field(pod, "InfraContainerID")
    entries = field(pod, "containers", None)
    if entries is None and field(pod, "NumContainers", None) == 0:
        return set()
    if not isinstance(entries, list):
        raise RuntimeError("Keanggotaan pod tidak dapat diverifikasi; tidak akan dihapus.")
    return {
        resource_id(container)
        for container in entries
        if resource_id(container) != infra
    }


def inventory():
    all_containers = read_json("ps", "--all", "--no-trunc", "--format=json")
    containers = [item for item in all_containers if owned(item)]
    container_ids = {resource_id(item) for item in containers}
    associated_pods = {field(item, "pod") for item in containers} - {""}
    pods = [
        item
        for item in read_json("pod", "ps", "--no-trunc", "--format=json")
        if owned(item, LEGACY_PODS) or resource_id(item) in associated_pods
    ]
    for pod in pods:
        if project_owners(pod) - PROJECTS:
            raise RuntimeError(
                f"Pod {resource_name(pod)} berlabel project asing. Reset dibatalkan."
            )
        unknown = members(inspect_pod(pod)) - container_ids
        if unknown:
            raise RuntimeError(
                f"Pod {resource_name(pod)} memiliki container di luar target: "
                f"{', '.join(sorted(unknown))}. Reset dibatalkan sebelum perubahan."
            )
    networks = [
        item
        for item in read_json("network", "ls", "--format=json")
        if owned(item, LEGACY_NETWORKS)
    ]
    volumes = [
        item
        for item in read_json("volume", "ls", "--format=json")
        if owned(item, LEGACY_VOLUMES)
    ]
    return containers, pods, networks, volumes


def reset(containers, pods, networks, volumes, delete_data):
    # Explicit IDs only: no prune, --all, recursive dependencies, or volume force.
    for container in containers:
        run("rm", "--force", "--time", "10", resource_id(container))
    for pod in pods:
        # Recheck after container removal: never force-remove an unexpected member.
        if members(inspect_pod(pod)):
            raise RuntimeError(
                f"Pod {resource_name(pod)} masih berisi container; hentikan pemeriksaan di sini."
            )
        run("pod", "rm", "--force", "--time", "10", resource_id(pod))
    for network in networks:
        # No force: Podman must refuse a network still used by another container.
        run("network", "rm", resource_name(network))
    if delete_data:
        for volume in volumes:
            # No force: never remove another container to free a shared volume.
            run("volume", "rm", resource_name(volume))


def main(argv=None):
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--apply", action="store_true", help="Apply after typed confirmation")
    parser.add_argument(
        "--delete-data", action="store_true", help="Also permanently delete selected data volumes"
    )
    args = parser.parse_args(argv)
    containers, pods, networks, volumes = inventory()
    print("Target project labels: " + ", ".join(sorted(PROJECTS)))
    print("Pod/volume/network lama tanpa label hanya dipilih dari daftar nama eksplisit.")
    for title, items in [
        ("Containers", containers), ("Pods", pods), ("Networks", networks),
        ("Volumes — HAPUS DATA" if args.delete_data else "Volumes — DIPERTAHANKAN", volumes),
    ]:
        print(f"\n{title}:")
        for item in items:
            print(f"  {resource_name(item)} {resource_id(item)}".rstrip())
        if not items:
            print("  (tidak ada)")
    if not args.apply:
        print("\nPreview saja. Tambahkan --apply untuk menjalankan daftar ini.")
        return 0
    if not (containers or pods or networks or (args.delete_data and volumes)):
        print("\nTidak ada target yang perlu dihapus.")
        return 0
    phrase = "DESTROY KENCLENG DATA" if args.delete_data else "RESET KENCLENG CONTAINERS"
    if args.delete_data:
        print("\nPERMANEN: isi volume PostgreSQL/object storage/cache terpilih akan hilang. Tidak ada backup otomatis.")
    print("Images, file project, env, serta resource di luar target tidak dihapus.")
    answer = input(f"Ketik {phrase!r} untuk melanjutkan: ")
    if answer != phrase:
        print("Dibatalkan. Tidak ada perubahan.")
        return 1
    reset(containers, pods, networks, volumes, args.delete_data)
    remaining = inventory()
    if any(remaining[:3]) or (args.delete_data and remaining[3]):
        raise RuntimeError("Masih ada resource target. Periksa sebelum memulai ulang.")
    print("\nReset target selesai. Image yang sudah tersedia tetap dapat dipakai ulang.")
    return 0


if __name__ == "__main__":
    try:
        sys.exit(main())
    except (RuntimeError, OSError, ValueError, EOFError, KeyboardInterrupt) as error:
        print(f"\nReset berhenti: {error}", file=sys.stderr)
        sys.exit(1)
