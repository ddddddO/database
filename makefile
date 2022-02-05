run:
	go run cmd/database/main.go

send:
	echo '{"raw_query":"AAAA"}' | nc localhost 9999

test:
	go test ./... -v