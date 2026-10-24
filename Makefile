.DEFAULT_GOAL := help

# Docker Engine/Desktop: docker compose. Podman: pass COMPOSE=podman-compose.
COMPOSE ?= docker compose
MINIO_PROJECT ?= kencleng-storage
COMPOSE_BASE = $(COMPOSE) -f docker-compose.yml
# Storage has an independent lifecycle/config hash, especially on podman-compose.
COMPOSE_MINIO = $(COMPOSE) -p $(MINIO_PROJECT) -f docker-compose.minio.yml
COMPOSE_STORAGE = $(COMPOSE) -p kencleng-storage -f docker-compose.storage.yml
APP_SERVICE ?= backend frontend
DEV_UID ?= $(shell id -u)
DEV_GID ?= $(shell id -g)
export DEV_UID DEV_GID
APPS_PROVIDER_FILES =
ifneq ($(filter podman podman-compose,$(notdir $(firstword $(COMPOSE)))),)
APPS_PROVIDER_FILES = -f docker-compose.apps.podman.yml
endif
COMPOSE_APPS = $(COMPOSE) -p kencleng-apps -f docker-compose.apps.yml $(APPS_PROVIDER_FILES)

.PHONY: help setup-backend setup-frontend dev-backend dev-frontend verify \
	up db-wait db-check down status logs config up-podman down-podman minio-up minio-down \
	docker-run docker-stop docker-down docker-status docker-logs docker-config

.PHONY: storage-up storage-down storage-status storage-logs storage-config storage-check storage-persistence-check

help:
	@echo 'App setup/run: cd backend or frontend, then make setup / make run'
	@echo 'Infra: make up / make status / make logs / make down'
	@echo 'Podman: add COMPOSE=podman-compose (or use make up-podman)'
	@echo 'Optional SeaweedFS: make storage-up / storage-check / storage-persistence-check / storage-down'
	@echo 'Container apps: make docker-run / docker-down / docker-status (same COMPOSE choice)'

setup-backend:
	$(MAKE) -C backend setup

setup-frontend:
	$(MAKE) -C frontend setup

dev-backend:
	$(MAKE) -C backend run

dev-frontend:
	$(MAKE) -C frontend run

verify:
	$(MAKE) -C backend verify
	cd frontend && npm run verify

up:
	$(COMPOSE_BASE) up -d postgres
	$(MAKE) db-wait
	$(MAKE) db-check

db-wait:
	$(COMPOSE_BASE) exec -T postgres sh -ec 'attempt=0; \
		until pg_isready -h 127.0.0.1 -U "$$POSTGRES_USER" -d "$$POSTGRES_DB"; do \
			attempt=$$((attempt + 1)); \
			if [ "$$attempt" -ge 30 ]; then echo "PostgreSQL did not become ready" >&2; exit 1; fi; \
			sleep 2; \
		done'

db-check:
	$(COMPOSE_BASE) exec -T postgres sh -ec 'mismatches=$$(psql -X -w -U "$$POSTGRES_USER" -d postgres -v ON_ERROR_STOP=1 -Atc \
		"SELECT count(*) FROM pg_database WHERE datallowconn AND datcollversion IS DISTINCT FROM pg_database_collation_actual_version(oid)"); \
		if [ "$$mismatches" != 0 ]; then \
			psql -X -w -U "$$POSTGRES_USER" -d postgres -v ON_ERROR_STOP=1 -c \
			"SELECT datname, datcollversion, pg_database_collation_actual_version(oid) AS actual FROM pg_database WHERE datallowconn"; \
			echo "Database collation mismatch: inspect image/volume compatibility before development; no metadata was changed." >&2; \
			exit 1; \
		fi'

down:
	$(COMPOSE_APPS) down
	$(COMPOSE_STORAGE) down
	$(COMPOSE_BASE) down

status:
	$(COMPOSE_BASE) ps

logs:
	$(COMPOSE_BASE) logs --tail=100 -f postgres

config:
	$(COMPOSE_BASE) config

up-podman:
	$(MAKE) COMPOSE=podman-compose up

down-podman:
	$(MAKE) COMPOSE=podman-compose down

minio-up:
	$(COMPOSE_MINIO) up -d minio
	$(COMPOSE_MINIO) run --rm minio-init

minio-down:
	$(COMPOSE_MINIO) stop minio

storage-up:
	$(MAKE) up
	$(COMPOSE_STORAGE) pull seaweedfs storage-cli
	$(COMPOSE_STORAGE) up -d seaweedfs
	python3 scripts/local-storage.py --compose '$(COMPOSE)' init

storage-down:
	$(COMPOSE_STORAGE) down

storage-status:
	$(COMPOSE_STORAGE) ps

storage-logs:
	$(COMPOSE_STORAGE) logs --tail=100 -f seaweedfs

storage-config:
	$(COMPOSE_STORAGE) config

storage-check:
	python3 scripts/local-storage.py --compose '$(COMPOSE)' check

storage-persistence-check:
	python3 scripts/local-storage.py --compose '$(COMPOSE)' persistence-check

docker-run:
	$(MAKE) up
	mkdir -p backend/tmp frontend/node_modules frontend/.next
	$(COMPOSE_APPS) up --build $(APP_SERVICE)

docker-stop:
	$(COMPOSE_APPS) stop $(APP_SERVICE)

docker-down:
	$(COMPOSE_APPS) down

docker-status:
	$(COMPOSE_APPS) ps

docker-logs:
	$(COMPOSE_APPS) logs --tail=100 -f $(APP_SERVICE)

docker-config:
	$(COMPOSE_APPS) config
