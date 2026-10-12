# backend/Makefile
COMPOSE ?= podman-compose
GITLEAKS_RUN = $(COMPOSE) -f ../docker-compose.yml --profile tools run --rm --no-deps -T gitleaks
GITLEAKS_SHELL = $(COMPOSE) -f ../docker-compose.yml --profile tools run --rm --no-deps -T --entrypoint /bin/bash gitleaks

.PHONY: verify lint test-unit test-race test-contract security gitleaks gitleaks-history gitleaks-worktree migrate-up migrate-down

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
	$(MAKE) gitleaks
	govulncheck ./...
gitleaks:
	$(MAKE) -k gitleaks-history gitleaks-worktree
gitleaks-history:
	$(GITLEAKS_RUN) git /repo --log-opts="--all --full-history -- backend/" --redact --no-banner --no-color --timeout 180
gitleaks-worktree:
	$(GITLEAKS_SHELL) -lc 'set -euo pipefail; \
		common=(--redact --no-banner --no-color --timeout 180); \
		gitleaks git --pre-commit "$${common[@]}" /repo/backend; \
		gitleaks git --staged "$${common[@]}" /repo/backend; \
		while IFS= read -r -d "" file; do \
			gitleaks dir "$${common[@]}" "/repo/$$file"; \
		done < <(git -C /repo ls-files -z --others --exclude-standard -- backend/)'
migrate-up:
	migrate -path migrations -database "$$DATABASE_URL" up
migrate-down:
	migrate -path migrations -database "$$DATABASE_URL" down -all
