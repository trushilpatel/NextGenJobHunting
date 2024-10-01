# NextGenJobHunting

Automating resume and cover letter generation, allowing users to easily tailor their applications to specific job opportunities while highlighting their unique skills and experiences.

## Next-Gen Job Hunting Backend Setup

This guide will walk you through setting up and running the backend for the Next-Gen Job Hunting project, including setting up Docker, installing Go, setting up live reloading with Air, and using `make` to compile and run the project.

### Prerequisites

Ensure the following tools are installed on your machine:

- [Docker](https://docs.docker.com/get-docker/)
- [Go](https://go.dev/doc/install)
- [Air](https://github.com/cosmtrek/air) (for live reloading)
- [Make](https://www.gnu.org/software/make/)

### Setup Instructions

### Step 1: Clone the Repository

Start by cloning the project repository.

```bash
git clone https://github.com/your-username/next-gen-job-hunting-backend.git
cd next-gen-job-hunting-backend
```

### Step 2: Set Up Docker and DB

Make sure Docker is running on your machine. Use Docker for containerized database management and services. You can spin up a database container with the necessary configurations.

**Build Docker Compose**

```sh
make build_docker
```

**Start the Development Database**

```sh
make dev_db
```

### Step 3: Install Go and Air

##### 1. Install Go

If you haven't already installed Go, follow the instructions [here](https://go.dev/doc/install)

##### 2. Install Air

Install Air for live reloading during development.

```sh
go install github.com/cosmtrek/air@latest
```

##### 3. Install wire (Golang dependency Injector)

```sh
go install github.com/google/wire/cmd/wire@latest
```

make sure you can run "air" and "wire" command in terminal.

### Step 4: Configure Environment Variables

Ensure that the environment variables required by the project are defined in the .env file located in the root directory.

### Step 5: Compile and Run the Backend

The Makefile provided automates the process of compiling, running, and managing the backend during development.

##### 1. Compile the Project

To compile the Go project, run the following:

```sh
make build_backend
```

This command generates the next-gen-job-hunting binary by wiring dependencies using Go’s wire and building the project.

##### 2. Run the Project

To run the project manually:

```sh
make run_backend
```

This runs the next-gen-job-hunting.go file, starting the backend server.

### Step 6: Live Reloading with Air

During development, you can use Air to enable live reloading of the Go backend when code changes are detected.

To start live reloading:

```sh
make dev_backend
```

This command will:

1. Kill any running processes on the port specified in your .env file.
2. Rebuild the project using wire and Go.
3. Start the backend with Air for live reloading.

### Step 7: Kill Running Processes on Port

To manually kill any process running on a specific port (default is 8080), you can use:

```sh
make kill_port
```

### Step 8: Print Database Configurations

To print the database configuration from your .env file:

```sh
make print_db_config
```

````

```sh
make build_backend
````

This command generates the next-gen-job-hunting binary by wiring dependencies using Go’s wire and building the project.

**Run the Project**
To run the project manually:

```sh
make run_backend
```

This runs the next-gen-job-hunting.go file, starting the backend server.

### Step 6: Live Reloading with Air

During development, you can use Air to enable live reloading of the Go backend when code changes are detected.

To start live reloading:

```sh
make dev_backend
```

This command will:
Kill any running processes on the port specified in your .env file.
Rebuild the project using wire and Go.
Start the backend with Air for live reloading.

### Step 7: Kill Running Processes on Port

To manually kill any process running on a specific port (default is 8080), you can use:

```sh
make kill_port
```

### Step 8: Print Database Configurations

To print the database configuration from your .env file:

```sh
make print_db_config
```

### Git Branch Naming Conventions

To maintain a consistent workflow, use the following branch naming conventions based on the type of work being performed:

- **Feature Branches**: For new features or enhancements.
  ```
  git checkout -b feature/<feature-name>
  ```
- **Bug Fix Branches**: For fixing bugs or issues.
  ```
  git checkout -b bugfix/<issue-id>-<short-description>
  ```
- **Test Branches**: For adding or modifying tests.
  ```
  git checkout -b test/<test-name>
  ```
- **Documentation Branches**: For updating or creating documentation.
  ```
  git checkout -b docs/<document-name>
  ```
