#!make

build_backend:
	wire ./di
	go build -o ./next-gen-job-hunting .

kill_port:
	@pid=$(@lsof -t -i:${PORT:=8080}) && if [ -n "${pid}" ]; then kill -9 $pid; fi
