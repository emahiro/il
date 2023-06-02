gen-pb:
	protoc -I=./proto --go_out=./pb ./proto/*.proto

gen-grpc:
	protoc -I ./proto --go_out ./pb --go_opt paths=source_relative --go-grpc_out ./pb --go-grpc_opt paths=source_relative ./proto/*.proto
