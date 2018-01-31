NAME="vpnhub"
OWNER="Axsh.co"

build: gen deps client server

test:
	go test -v

gen:
	go generate -v ./api/ ./driver/

client: deps
	go build  -v -o ./hubclient ./bin/hubclient/main.go

server: deps
	go build  -v -o ./hubserver ./bin/hubserver/main.go

deps:
	govendor sync
