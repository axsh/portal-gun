package model

//go:generate protoc -I../proto --go_out=plugins=grpc:${GOPATH}/src ../proto/model.proto
