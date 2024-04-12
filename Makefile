build:
	go build .

protoc:
	protoc -I proto/ proto/claims-provider.proto --go_out=plugins=grpc:proto

run:
	./claims-provider s