# Crawler Makefile

# Define the virtual environment directory
VENV_DIR := ./venv

# Target to clean up the virtual environment
.PHONY: clean
clean:
	@echo "Removing virtual environment..."
	rm -rf $(VENV_DIR)

# Target to create a virtual environment
.PHONY: venv
venv: clean
	@echo "Creating virtual environment..."
	python3 -m venv $(VENV_DIR)

.PHONY: connect-venv
connect-venv:
	@echo "Connecting to the virtual environment..."
	source $(VENV_DIR)/bin/activate


# Target to install dependencies in the virtual environment
.PHONY: install
install:
	@echo "Installing dependencies..."
	@source $(VENV_DIR)/bin/activate && pip install -r requirements.txt

# Target to run the LinkedIn crawler within the virtual environment
.PHONY: run
run: install
	@echo "Running LinkedIn crawler..."
	@source $(VENV_DIR)/bin/activate && python crawlers.py

# Target to run Chrome with debugging options based on the OS
.PHONY: run-chrome
run-chrome:
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
