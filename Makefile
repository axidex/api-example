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

db-down:
	cd ./compose/db && docker compose down

db-restart: db-down db

# App Docker compose
app:
	cd ./compose/app && docker compose up -d

app-down:
	cd ./compose/app && docker compose down

app-restart: app-down app
