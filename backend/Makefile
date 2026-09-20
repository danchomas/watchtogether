include .env
export

env-up:
	docker compose up postgres -d

env-down:
	docker compose down postgres

env-cleanup:
	@read -p "are you sure than you need cleanup all database files? [y/n]: "; \
	if [ "$$REPLY" = "y" ]; then \
		docker compose down postgres && \
		sudo rm -rf out/pgdata && \
		echo "environment files are cleanup"; \
	else \
		echo "cleanup was canceled"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "u need write param seq. target: make migrate-create seq=init"; \
		exit 1; \
	fi
	@mkdir -p migrations
	docker compose run --rm --user "$$(id -u):$$(id -g)" migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@$(MAKE) migrate-action action=up

migrate-down:
	@$(MAKE) migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "u need write param seq. target: make migrate-create seq=init"; \
		exit 1; \
	fi
	@docker compose run --rm migrate \
		-path /migrations \
		-database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@postgres:5432/$(POSTGRES_DB)?sslmode=disable" \
		"$(action)"
