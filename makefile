#!make

include ./.env
export $(shell sed 's/=.*//' ./.env)

print_db_config:
	@echo "Database Host: $(DB_HOST)"
	@echo "Database Port: $(DB_PORT)"
	@echo "Database Username: $(DB_USERNAME)"
	@echo "Database Name: $(DB_NAME)"

kill_port:
	@pid=$(@lsof -t -i:${PORT:=8080}) && if [ -n "${pid}" ]; then kill -9 $pid; fi

build_docker:
	cd ./infra && docker-compose -f docker-compose.yml build

dev_db:
	cd ./infra && docker-compose -f docker-compose.yml up -d
	
build_backend:
	cd ./backend && wire ./di
	cd ./backend && go build -o ./next-gen-job-hunting .

run_backend:
	cd ./backend && go run next-gen-job-hunting

dev_backend:
	cd ./backend && air --build.cmd "make kill_port && make build_backend" --build.bin "./next-gen-job-hunting" --build.exclude_dir "tmp" --build.exclude_file "./di/wire_gen.go,*di/*"