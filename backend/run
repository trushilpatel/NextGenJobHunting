# Kill Port
pid=$(lsof -t -i:${PORT:=8080}) && if [ -n "$pid" ]; then kill -9 $pid; fi
echo "Killed process on port ${PORT:=8080}"

#!/bin/bash
wire ./di

# Build the Go application
go build -o next-gen-job-hunting .


# Run the application
./next-gen-job-hunting
