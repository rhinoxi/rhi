package cs

var cheatsheet = map[string]string{
	"grpc-go": `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <path/to/proto>`,
}
