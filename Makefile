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