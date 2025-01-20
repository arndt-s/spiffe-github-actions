module github.com/arndt-s/spiffe-github-actions/agent

go 1.23

require google.golang.org/grpc v1.69.4

require (
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/crypto v0.28.0 // indirect
)

require (
	github.com/spiffe/go-spiffe v1.1.0
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gopkg.in/square/go-jose.v2 v2.6.0
)
