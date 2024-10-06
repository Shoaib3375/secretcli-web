# Makefile

.PHONY: up down logs

# Define the service name
SERVICE_NAME=db

# Command to start the database service
db-up:
	docker-compose up -d $(SERVICE_NAME)

# Command to stop the services
db-down:
	docker-compose down

# Command to view logs for the database service
db-logs:
	docker-compose logs -f $(SERVICE_NAME)

# Command to create a PostgreSQL database
create-db:
	docker-compose exec $(SERVICE_NAME) psql -U admin -c "CREATE DATABASE secretcli;"

# Command to remove the PostgreSQL database
drop-db:
	docker-compose exec $(SERVICE_NAME) psql -U admin -c "DROP DATABASE secretcli;"
