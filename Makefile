all:cache-server cache-cli

cache-server:
	go build -o bin/cache-server server/main.go

cache-cli:
	go build -o bin/cache-cli client/main.go

test:
	go test ./...

clean:
	rm -r ./bin