run:
	go run cmd/database/main.go

send:
	echo '{"raw_query":"select * from ttttest;"}' | nc localhost 9999

test:
	go test ./... -v

fmt:
	go fmt ./...