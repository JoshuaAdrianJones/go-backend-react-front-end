backend:
	go fmt
	go clean
	go build -o go_server app.go
	./go_server

frontend:
	cd front_end_example && yarn dev
