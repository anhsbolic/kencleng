# backend/Makefile
COMPOSE ?= podman-compose
GITLEAKS_RUN = $(COMPOSE) -f ../docker-compose.yml --profile tools run --rm --no-deps -T gitleaks

.PHONY: verify lint gosec test-scanner-gate test-unit test-race test-contract security gitleaks gitleaks-history gitleaks-worktree migrate-up migrate-down

verify: lint test-unit test-race test-contract security
lint:
	staticcheck ./...
	$(MAKE) gosec
gosec:
	python3 scripts/scanner_gate.py gosec -- $(shell command -v gosec)
test-scanner-gate:
	PYTHONDONTWRITEBYTECODE=1 python3 -m unittest discover -s scripts -p 'test_scanner_gate.py'
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
	python3 scripts/scanner_gate.py gitleaks-history -- $(GITLEAKS_RUN)
gitleaks-worktree:
	python3 scripts/scanner_gate.py gitleaks-worktree -- $(GITLEAKS_RUN)
migrate-up:
	migrate -path migrations -database "$$DATABASE_URL" up
migrate-down:
	migrate -path migrations -database "$$DATABASE_URL" down -all
