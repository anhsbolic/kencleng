COMPOSE ?= podman-compose
.DEFAULT_GOAL := verify

.PHONY: verify up down up-podman down-podman gitleaks

gitleaks:
	$(COMPOSE) --profile tools run --rm --no-deps -T gitleaks

verify:
	cd backend && make verify
	cd frontend && npm run verify

up:
	docker-compose up -d

down:
	docker-compose down

up-podman:
	podman-compose up -d

down-podman:
	podman-compose down
