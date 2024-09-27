#!make

include ../.env
export $(shell sed 's/=.*//' ../.env)

print_db_config:
	@echo "Database Host: $(DB_HOST)"
	@echo "Database Port: $(DB_PORT)"
	@echo "Database Username: $(DB_USERNAME)"
	@echo "Database Name: $(DB_NAME)"

backend:
	./run.sh