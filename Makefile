telemetry:
	cd ./compose/telemetry && docker compose up -d

telemetry-down:
	cd ./compose/telemetry && docker compose down

telemetry-restart: telemetry-down telemetry