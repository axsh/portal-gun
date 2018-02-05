NAME="vpnhub"
OWNER="Axsh.co"

build: protobuf deps client server

test:
	go test -v

protobuf:
	go generate -v ./api/ ./model/

client: deps
	go build  -v -o ./portal-client ./cmd/portal-client

server: deps
	go build  -v -o ./portal-server ./cmd/portal-server

deps:
	govendor sync
