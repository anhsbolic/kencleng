"""Exercise dependency-volume refresh without Docker, network or real npm installs."""

import os
from pathlib import Path
import subprocess
import tempfile
import unittest

ENTRYPOINT = Path(__file__).resolve().parents[1] / "frontend/docker-dev.sh"


class FrontendDependencyVolumeTests(unittest.TestCase):
    def test_persistent_volume_refreshes_for_manifest_lock_and_runtime_changes(self):
        with tempfile.TemporaryDirectory(prefix="kencleng-dependency-volume-") as directory:
            root = Path(directory)
            binaries = root / "bin"
            binaries.mkdir()
            (binaries / "node").write_text('#!/bin/sh\nprintf "%s\\n" "$PROBE_NODE_VERSION"\n')
            (binaries / "npm").write_text(
                '#!/bin/sh\nset -eu\n[ "$*" = ci ]\n'
                'printf "install\\n" >> installs.log\nmkdir -p node_modules/.bin\n'
                'printf \'#!/bin/sh\\nprintf "%%s\\n" "$*" > next-args.log\\n\' > node_modules/.bin/next\n'
                'chmod +x node_modules/.bin/next\n'
            )
            for file in binaries.iterdir():
                file.chmod(0o755)
            (root / "package.json").write_text('{"name":"probe"}\n')
            (root / "package-lock.json").write_text('{"version":1}\n')
            env = {**os.environ, "PATH": f"{binaries}:{os.environ['PATH']}", "PROBE_NODE_VERSION": "v24.22.1"}

            def start(*arguments, version="v24.22.1"):
                subprocess.run(["sh", str(ENTRYPOINT), *arguments], cwd=root,
                               env={**env, "PROBE_NODE_VERSION": version}, check=True)

            def installs():
                return len((root / "installs.log").read_text().splitlines())

            start("--install")
            self.assertEqual(installs(), 1)
            self.assertFalse((root / "next-args.log").exists())
            start()
            self.assertEqual(installs(), 1)
            self.assertEqual((root / "next-args.log").read_text().strip(), "dev --hostname 0.0.0.0 --webpack")
            (root / "package-lock.json").write_text('{"version":2}\n')
            start()
            self.assertEqual(installs(), 2)
            (root / "package.json").write_text('{"name":"probe-updated"}\n')
            start()
            self.assertEqual(installs(), 3)
            start(version="v24.22.2")
            self.assertEqual(installs(), 4)

    def test_failed_install_does_not_stamp_or_launch_next(self):
        with tempfile.TemporaryDirectory(prefix="kencleng-failed-install-") as directory:
            root = Path(directory)
            binaries = root / "bin"
            binaries.mkdir()
            (binaries / "node").write_text('#!/bin/sh\necho v24.22.1\n')
            (binaries / "npm").write_text('#!/bin/sh\nexit 7\n')
            for file in binaries.iterdir():
                file.chmod(0o755)
            for name in ["package.json", "package-lock.json"]:
                (root / name).write_text('{}\n')
            result = subprocess.run(["sh", str(ENTRYPOINT)], cwd=root,
                                    env={**os.environ, "PATH": f"{binaries}:{os.environ['PATH']}"})
            self.assertEqual(result.returncode, 7)
            self.assertFalse((root / "node_modules/.kencleng-dev-deps").exists())


if __name__ == "__main__":
    unittest.main()
