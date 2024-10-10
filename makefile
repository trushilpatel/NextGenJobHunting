##### Common Variables #####
##### Environment #####
include ./.env
export $(shell sed 's/=.*//' ./.env)

.PHONY: print-db-config kill-port
print-db-config:
	@echo "Database Host: $(DB_HOST)"
	@echo "Database Port: $(DB_PORT)"
	@echo "Database Username: $(DB_USERNAME)"
	@echo "Database Name: $(DB_NAME)"

.PHONY: kill-port
kill-port:
	@pid=$$(lsof -t -i:${PORT:=8080}) && if [ -n "$$pid" ]; then kill -9 $$pid; fi

##### Docker Commands #####
.PHONY: build-docker dev-db

build-docker:
	cd ./infra && docker-compose -f docker-compose.yml build

dev-db:
	cd ./infra && docker-compose -f docker-compose.yml up -d


##### Backend Commands #####
.PHONY: build-backend run-backend dev-backend

build-backend:
	cd ./backend && wire ./di
	cd ./backend && go build -o ./next-gen-job-hunting .

run-backend:
	cd ./backend && go run next-gen-job-hunting

dev-backend:
	cd ./backend && air --build.cmd "make kill-port && make build-backend" --build.bin "./next-gen-job-hunting" --build.exclude_dir "tmp" --build.exclude_file "./di/wire_gen.go,*di/*"


##### Crawler Commands #####
.PHONY: crawler-run-chrome crawler-venv crawler-clean install-crawler crawler-run

# Run Chrome with debugging options based on the OS
crawler-run-chrome:
	@echo "Chrome to maintain a debugging session..."
	make -C crawler run-chrome

crawler-setup:
	@echo "Setting Up Crawler Virtual Environment and Installing libraries..."
	make -C crawler install

crawler-clean:
	@echo "Removing virtual environment..."
	make -C crawler clean


install-crawler:
	@echo "Installing crawler dependencies..."
	make -C crawler install

crawler-run:
	@echo "Running LinkedIn crawler from the main folder..."
	make -C crawler run
