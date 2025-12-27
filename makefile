.PHONY: backend
backend:
	docker compose up --build -d backend
db: 
	docker compose up --build -d postgres 
up: 
	docker compose up --build -d 
down: 
	docker compose down -v 
