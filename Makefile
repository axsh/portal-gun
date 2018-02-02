NAME="vpnhub"
OWNER="Axsh.co"

build: gen deps client server

test:
	go test -v

gen:
	go generate -v ./api/ ./model/

client: deps
	go build  -v -o ./hubclient ./cmd/hubclient

server: deps
	go build  -v -o ./hubserver ./cmd/hubserver

deps:
	govendor sync
