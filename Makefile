# Compose

compose: telemetry db

network:
	docker network create api-example

# Telemetry Docker Compose
telemetry:
	cd ./compose/telemetry && docker compose up -d

telemetry-down:
	cd ./compose/telemetry && docker compose down

telemetry-restart: telemetry-down telemetry

# DB Docker Compose
db:
	cd ./compose/db && docker compose up -d

db-logs:
	cd ./compose/db && docker compose logs -f

db-remove:
	cd ./compose/db && docker compose down -v && rm -rf ./postgres_data

db-restart: db-down db

# App Docker compose
app:
	cd ./compose/app && docker compose up -d --pull always

app-down:
	cd ./compose/app && docker compose down

app-restart: app-down app

# Transactions Docker compose
.PHONY: transactions
transactions:
	cd ./compose/transactions && docker compose up -d --pull always

transactions-logs:
	cd ./compose/transactions && docker compose logs -f

transactions-remove:
	cd ./compose/transactions && docker compose down -v

transactions-restart: transactions-down transactions


# GIT
remove-local-branches:
	@for branch in $$(git branch -r | grep 'feature/' | sed 's/^[* ]*//'); do \
		git branch -D $$branch; \
	done

remove-remote-branches:
	@for branch in $$(git branch | grep 'origin/feature/' | sed 's/origin\///'); do \
		git push origin --delete $$branch; \
	done