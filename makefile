.PHONY: backend monitor
backend:
	docker compose up --build -d --force-recreate backend
data:
	docker compose up --build -d postgres redis
monitor:
	docker compose down -v grafana alloy prometheus loki && docker compose up --build -d --force-recreate grafana alloy prometheus loki
up:
	docker compose up -d
build:
	docker compose up --build -d 
down: 
	docker compose down -v 
