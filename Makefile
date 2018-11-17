build: build-proto
	dep check
	go build

build-proto:
	protoc -I ./proto ./proto/proxysql/proxysql.proto --go_out=plugins=grpc:proto

test:
	go test ./...
