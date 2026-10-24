# backend/Makefile
.DEFAULT_GOAL := verify
AIR_VERSION := v1.67.4

.PHONY: setup run gvm-run docker-run docker-stop dev-tools dev verify lint test-unit test-race test-contract security migrate-up migrate-down

setup:
	test -f .env || cp ../.env.example .env
	$(MAKE) dev-tools

run: dev

gvm-run:
	@bash -c 'kencleng_gvm_root="$${GVM_ROOT:-$$HOME/.gvm}"; \
		[ -s "$$kencleng_gvm_root/scripts/gvm" ] || { echo "gvm script not found; install gvm or set GVM_ROOT." >&2; exit 1; }; \
		. "$$kencleng_gvm_root/scripts/gvm" && gvm use "$$(cat ../.go-version)" && exec $(MAKE) run'

docker-run:
	$(MAKE) -C .. docker-run APP_SERVICE=backend

docker-stop:
	$(MAKE) -C .. docker-stop APP_SERVICE=backend

dev-tools:
	mkdir -p bin
	GOBIN="$(CURDIR)/bin" go install github.com/air-verse/air@$(AIR_VERSION)

dev:
	@test -x ./bin/air || { echo 'Air is missing. Run make dev-tools first.' >&2; exit 1; }
	./bin/air -c .air.toml

verify: lint test-unit test-race test-contract security
lint:
	staticcheck ./...
	gosec ./...
test-unit:
	go test ./...
test-race:
	go test -race ./...
test-contract:
	go test -tags=contract ./...
security:
	gitleaks detect
	govulncheck ./...
migrate-up:
	migrate -path migrations -database "$$DATABASE_URL" up
migrate-down:
	migrate -path migrations -database "$$DATABASE_URL" down -all
