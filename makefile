##### Common Variables #####
##### Environment #####
include ./.env
export $(shell sed 's/=.*//' ./.env)

print-db-config:
	@echo "Database Host: $(DB_HOST)"
	@echo "Database Port: $(DB_PORT)"
	@echo "Database Username: $(DB_USERNAME)"
	@echo "Database Name: $(DB_NAME)"

kill-port:
	@pid=$$(lsof -t -i:${PORT:=8080}) && if [ -n "$$pid" ]; then kill -9 $$pid; fi

##### Docker #####
build-docker:
	cd ./infra && docker-compose -f docker-compose.yml build

dev-db:
	cd ./infra && docker-compose -f docker-compose.yml up -d

##### Backend #####
build-backend:
	cd ./backend && wire ./di
	cd ./backend && go build -o ./next-gen-job-hunting .

run-backend:
	cd ./backend && go run next-gen-job-hunting

dev-backend:
	cd ./backend && air --build.cmd "make kill-port && make build-backend" --build.bin "./next-gen-job-hunting" --build.exclude_dir "tmp" --build.exclude_file "./di/wire_gen.go,*di/*"

##### Crawler #####

.PHONY: run-chrome

crawler-run-chrome:
	@echo "Detected OS: $$(uname -s | xargs)"
	@if [ "$$(uname -s | xargs)" = "Linux" ]; then \
		echo "Running Chrome on Linux..."; \
		google-chrome --remote-debugging-port=9222 --user-data-dir="$$(echo ~)/chrome-debug-session"; \
	elif [ "$$(uname -s | xargs)" = "Darwin" ]; then \
		echo "Running Chrome on macOS..."; \
		/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome --remote-debugging-port=9222 --user-data-dir="$$(echo ~)/chrome-debug-session"; \
	elif [ "$$(uname -s | xargs)" = "Windows" ]; then \
		echo "Running Chrome on Windows..."; \
		"C:\Program Files\Google\Chrome\Application\chrome.exe" --remote-debugging-port=9222 --user-data-dir="$$(echo %USERPROFILE%)/chrome-debug-session"; \
	else \
		echo "Unsupported OS: $$(uname -s | xargs)"; \
	fi
