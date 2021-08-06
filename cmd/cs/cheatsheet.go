package cs

type csType map[string]interface{}

var csm = csType{
	"grpc-go": `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative <path/to/proto>`,
	"openssl": csType{
		"generate private key & public certificate": `openssl req -newkey rsa:2048 -nodes -keyout <key.pem> -x509 -days 365 -out <certificate.pem>`,
	},
}
