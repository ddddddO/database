run:
	go run cmd/database/main.go

send:
	echo '{"raw_query":"select 1"}' | nc localhost 9999

test:
	go test ./... -v
