.DEFAULT_GOAL := help

.PHONY: help setup run nvm-run docker-run docker-stop check-hot-reload dev verify

help:
	@echo 'make setup  - preserve/create .env.local and install locked dependencies'
	@echo 'make run    - start the Next.js development server with Fast Refresh'
	@echo 'make nvm-run - activate the project Node.js version through nvm, then run'
	@echo 'make docker-run - run this app in a development container (COMPOSE=podman-compose supported)'
	@echo 'make check-hot-reload - verify both running apps, temporarily editing/restoring source'
	@echo 'make verify - run lint and unit-test tooling'

setup:
	test -f .env.local || cp .env.example .env.local
	npm ci

run:
	npm run dev

nvm-run:
	@bash -c 'kencleng_nvm_dir="$${NVM_DIR:-$$HOME/.nvm}"; \
		[ -s "$$kencleng_nvm_dir/nvm.sh" ] || { echo "nvm.sh not found; install nvm or set NVM_DIR." >&2; exit 1; }; \
		. "$$kencleng_nvm_dir/nvm.sh" --no-use && nvm use && exec $(MAKE) run'

docker-run:
	$(MAKE) -C .. docker-run APP_SERVICE=frontend

docker-stop:
	$(MAKE) -C .. docker-stop APP_SERVICE=frontend

check-hot-reload:
	node ../scripts/check-hot-reload.mjs

dev: run

verify:
	npm run verify
