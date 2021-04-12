run:
	go run main.go

run-watch:
	nodemon --exec go run main.go --signal SIGTERM