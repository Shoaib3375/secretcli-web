.PHONY: up down logs db-up db-down prometheus-logs create-db drop-db

# Define service names
DB_SERVICE_NAME=db
PROMETHEUS_SERVICE_NAME=prometheus

# Command to start all services
up:
	docker-compose up -d

# Command to stop all services
down:
	docker-compose down

# Command to start the database service
db-up:
	docker-compose up -d $(DB_SERVICE_NAME)

# Command to stop the database service
db-down:
	docker-compose stop $(DB_SERVICE_NAME)

# Command to view logs for the database service
db-logs:
	docker-compose logs -f $(DB_SERVICE_NAME)

# Command to view logs for the Prometheus service
prometheus-logs:
	docker-compose logs -f $(PROMETHEUS_SERVICE_NAME)

# Command to create the PostgreSQL database
create-db:
	docker-compose exec $(DB_SERVICE_NAME) psql -U admin -c "CREATE DATABASE secretcli;"

# Command to drop the PostgreSQL database
drop-db:
	docker-compose exec $(DB_SERVICE_NAME) psql -U admin -c "DROP DATABASE secretcli;"
