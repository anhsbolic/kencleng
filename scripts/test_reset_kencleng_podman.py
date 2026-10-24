"""Safety regression checks. Podman is mocked; no real resources are touched."""

import contextlib
import copy
import importlib.util
import io
import json
from pathlib import Path
import subprocess
import unittest
from unittest.mock import patch

SPEC = importlib.util.spec_from_file_location(
    "reset_podman", Path(__file__).with_name("reset-kencleng-podman.py")
)
reset = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(reset)


class FakePodman:
    def __init__(self, shared=False):
        label = {"io.podman.compose.project": "kencleng"}
        self.containers = [
            {"Id": "db-id", "Names": ["kencleng_postgres_1"], "Labels": label, "Pod": "pod-id"},
            {"Id": "old-id", "Names": ["kencleng_caddy_1"], "Labels": label, "Pod": "pod-id"},
            {"Id": "other-id", "Names": ["other_database"], "Labels": {"com.docker.compose.project": "other"}, "Pod": "pod-id" if shared else ""},
        ]
        self.pods = [{"Id": "pod-id", "Name": "pod_kencleng", "Labels": {}}]
        self.networks = [{"Name": "kencleng_default", "Labels": label}, {"Name": "other_default", "Labels": {}}]
        self.volumes = [
            {"Name": "kencleng_kencleng_pgdata", "Labels": {}},
            {"Name": "storage-data", "Labels": {"com.docker.compose.project": "kencleng-storage"}},
            {"Name": "other-data", "Labels": {}},
        ]
        self.mutations = []

    def __call__(self, *args):
        if args[0] == "ps":
            return json.dumps(self.containers)
        if args[:2] == ("pod", "ps"):
            return json.dumps(self.pods)
        if args[:2] == ("network", "ls"):
            return json.dumps(self.networks)
        if args[:2] == ("volume", "ls"):
            return json.dumps(self.volumes)
        if args[:2] == ("pod", "inspect"):
            return json.dumps([{
                "Id": "pod-id", "Name": "pod_kencleng", "InfraContainerID": "infra-id",
                "Containers": [{"Id": "infra-id"}] + [
                    {"Id": item["Id"]} for item in self.containers if item["Pod"] == "pod-id"
                ],
            }])
        self.mutations.append(args)
        if args[0] == "rm":
            self.containers = [item for item in self.containers if item["Id"] != args[-1]]
        elif args[:2] == ("pod", "rm"):
            self.pods = []
        elif args[:2] == ("network", "rm"):
            self.networks = [item for item in self.networks if item["Name"] != args[-1]]
        elif args[:2] == ("volume", "rm"):
            self.volumes = [item for item in self.volumes if item["Name"] != args[-1]]
        else:
            raise AssertionError(args)
        return ""


class ResetSafetyTests(unittest.TestCase):
    def execute(self, fake, flags, phrase=""):
        with patch.object(reset, "run", fake), patch("builtins.input", return_value=phrase), contextlib.redirect_stdout(io.StringIO()):
            return reset.main(flags)

    def test_preview_never_mutates_even_with_delete_data(self):
        fake = FakePodman()
        self.assertEqual(self.execute(fake, ["--delete-data"]), 0)
        self.assertEqual(fake.mutations, [])

    def test_cancel_never_mutates(self):
        fake = FakePodman()
        self.assertEqual(self.execute(fake, ["--apply", "--delete-data"], "no"), 1)
        self.assertEqual(fake.mutations, [])

    def test_container_reset_preserves_all_volumes_and_foreign_resources(self):
        fake = FakePodman()
        volumes = copy.deepcopy(fake.volumes)
        self.execute(fake, ["--apply"], "RESET KENCLENG CONTAINERS")
        self.assertEqual([item["Id"] for item in fake.containers], ["other-id"])
        self.assertEqual(fake.volumes, volumes)
        self.assertEqual([item["Name"] for item in fake.networks], ["other_default"])

    def test_fresh_reset_removes_only_selected_data_and_is_repeatable(self):
        fake = FakePodman()
        self.execute(fake, ["--apply", "--delete-data"], "DESTROY KENCLENG DATA")
        self.assertEqual([item["Name"] for item in fake.volumes], ["other-data"])
        for command in fake.mutations:
            if command[:2] in [("volume", "rm"), ("network", "rm")]:
                self.assertNotIn("--force", command)
        first_count = len(fake.mutations)
        self.execute(fake, ["--apply", "--delete-data"])
        self.assertEqual(len(fake.mutations), first_count)

    def test_shared_pod_aborts_before_any_mutation(self):
        fake = FakePodman(shared=True)
        with self.assertRaisesRegex(RuntimeError, "di luar target"):
            self.execute(fake, ["--apply", "--delete-data"], "DESTROY KENCLENG DATA")
        self.assertEqual(fake.mutations, [])

    def test_container_mode_resources_are_reset_without_touching_other_projects(self):
        fake = FakePodman()
        label = {"com.docker.compose.project": "kencleng-apps"}
        fake.containers.append({"Id": "app-id", "Names": ["kencleng-apps-frontend-1"], "Labels": label, "Pod": ""})
        fake.volumes.append({"Name": "kencleng-apps_frontend_modules", "Labels": label})
        self.execute(fake, ["--apply", "--delete-data"], "DESTROY KENCLENG DATA")
        self.assertEqual([item["Id"] for item in fake.containers], ["other-id"])
        self.assertEqual([item["Name"] for item in fake.volumes], ["other-data"])

    def test_familiar_names_do_not_override_foreign_project_labels(self):
        for name in ["kencleng_kencleng_pgdata", "kencleng_default", "pod_kencleng"]:
            self.assertFalse(reset.owned({"Name": name, "Labels": {"io.podman.compose.project": "other"}}, {name}))

    def test_unknown_pod_membership_is_not_treated_as_empty(self):
        with self.assertRaises(RuntimeError):
            reset.members({"Id": "pod-id"})

    def test_foreign_pod_owner_aborts_even_with_only_target_members(self):
        fake = FakePodman()
        fake.pods[0]["Labels"] = {"com.docker.compose.project": "other"}
        with self.assertRaisesRegex(RuntimeError, "project asing"):
            self.execute(fake, ["--apply"], "RESET KENCLENG CONTAINERS")
        self.assertEqual(fake.mutations, [])

    def test_new_pod_member_blocks_pod_force_removal(self):
        class GrowingPod(FakePodman):
            def __call__(self, *args):
                result = super().__call__(*args)
                if args[0] == "rm" and args[-1] == "old-id":
                    self.containers[0]["Pod"] = "pod-id"
                return result

        fake = GrowingPod()
        with self.assertRaisesRegex(RuntimeError, "masih berisi"):
            self.execute(fake, ["--apply", "--delete-data"], "DESTROY KENCLENG DATA")
        self.assertEqual([item["Id"] for item in fake.containers], ["other-id"])
        self.assertFalse(any(command[:2] == ("pod", "rm") for command in fake.mutations))
        self.assertEqual(len(fake.volumes), 3)

    def test_rootless_cleanup_error_has_hint_without_automatic_repair(self):
        failure = subprocess.CompletedProcess(
            [], 125, "", "rootless netns: kill network process: permission denied"
        )
        with patch.object(reset.subprocess, "run", return_value=failure) as command:
            with self.assertRaisesRegex(RuntimeError, "Periksa log AppArmor/pasta"):
                reset.run("rm", "--force", "--time", "10", "db-id")
        command.assert_called_once_with(
            ["podman", "rm", "--force", "--time", "10", "db-id"],
            capture_output=True, text=True, check=False,
        )

    def test_container_cleanup_failure_does_not_delete_volumes_or_networks(self):
        class FailingCleanup(FakePodman):
            def __call__(self, *args):
                if args[0] == "rm":
                    raise RuntimeError("rootless netns: kill network process: permission denied")
                return super().__call__(*args)

        fake = FailingCleanup()
        with self.assertRaisesRegex(RuntimeError, "kill network process"):
            self.execute(fake, ["--apply", "--delete-data"], "DESTROY KENCLENG DATA")
        self.assertEqual(fake.mutations, [])
        self.assertEqual(len(fake.volumes), 3)


if __name__ == "__main__":
    unittest.main()
