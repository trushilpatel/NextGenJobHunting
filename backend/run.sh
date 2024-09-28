#!/bin/bash
wire ./di

# Build the Go application
go build -o next-gen-job-hunting .

# Run the application
./next-gen-job-hunting
