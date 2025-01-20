package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spiffe/go-spiffe/proto/spiffe/workload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

func main() {
	s := grpc.NewServer()
	a, err := newAgent()
	if err != nil {
		panic(fmt.Errorf("failed to create agent: %v", err))
	}

	reflection.Register(s)
	workload.RegisterSpiffeWorkloadAPIServer(s, a)
	RegisterInitAPIServer(s, a)

	path := os.Getenv("SPIFFE_ENDPOINT_SOCKET")
	if path == "" {
		panic("SPIFFE_ENDPOINT_SOCKET is not set")
	}

	l, err := net.Listen("unix", path)
	if err != nil {
		panic(fmt.Errorf("failed to listen: %v", err))
	}

	fmt.Printf("Listening on %s\n", l.Addr().String())
	err = s.Serve(l)
	if err != nil {
		panic(fmt.Errorf("failed to serve: %v", err))
	}
}

type agent struct {
	signer jose.Signer
	UnimplementedInitAPIServer
	workload.UnimplementedSpiffeWorkloadAPIServer
}

func newAgent() (*agent, error) {
	signingKey := []byte("secret")
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: signingKey}, nil)
	if err != nil {
		return nil, err
	}

	return &agent{
		signer: signer,
	}, nil
}

func (a *agent) Init(ctx context.Context, req *InitRequest) (*InitResponse, error) {
	return &InitResponse{
		SpiffeId: "spiffe://example.org/agent",
	}, nil
}

func (a *agent) FetchJWTSVID(ctx context.Context, req *workload.JWTSVIDRequest) (*workload.JWTSVIDResponse, error) {
	claims := jwt.Claims{
		Expiry:  jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		Issuer:  "your-issuer",
		Subject: "your-subject",
	}

	raw, err := jwt.Signed(a.signer).Claims(claims).CompactSerialize()
	if err != nil {
		return nil, err
	}

	return &workload.JWTSVIDResponse{
		Svids: []*workload.JWTSVID{
			{
				Svid:     raw,
				SpiffeId: "",
			},
		},
	}, nil
}
