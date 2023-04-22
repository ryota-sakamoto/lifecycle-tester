dev:
	go run cmd/lifecycle-tester/main.go server

test:
	go test -v ./...
