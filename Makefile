NAME="vpnhub"
OWNER="Axsh.co"
PROTOC_VERSION=3.4.0

build: protobuf deps client server

test:
	@go test $(shell go list ./... | grep -v /vendor/)

protobuf:
ifeq ("$(shell which protoc)", "")
	@echo "Missing protoc"
else ifneq ("$(shell protoc --version | cut -d " " -f2)","$(PROTOC_VERSION)")
	@echo "protoc version mismatch, requires $(PROTOC_VERSION)"
else
	go generate -v ./api/ ./model/
endif


client: deps
	go build  -v -o ./portal-client ./cmd/portal-client

server: deps
	go build  -v -o ./portal-server ./cmd/portal-server

deps:
	govendor sync
