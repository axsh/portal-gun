NAME="vpnhub"
OWNER="Axsh.co"

test:
	go test -v

gen:
	go generate -v ./api/ ./model/

build:
	govendor sync
	go build  -v -o ./hubclient ./bin/hubclient/main.go
	go build  -v -o ./hubserver ./bin/hubserver/main.go
