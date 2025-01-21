package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/spiffe/go-spiffe/proto/spiffe/workload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	status "google.golang.org/grpc/status"
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

	if err := os.Chmod(l.Addr().String(), os.ModePerm); err != nil {
		panic(fmt.Errorf("unable to change UDS permissions: %w", err))
	}

	fmt.Printf("Listening on %s\n", l.Addr().String())
	err = s.Serve(l)
	if err != nil {
		panic(fmt.Errorf("failed to serve: %v", err))
	}
}

type agent struct {
	signer   jose.Signer
	idToken  string
	spiffeId string

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
	spiffeId, err := constructSpiffeID(req.IdToken)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to construct spiffe id: %v", err)
	}

	a.idToken = req.IdToken
	a.spiffeId = spiffeId
	return &InitResponse{
		SpiffeId: spiffeId,
	}, nil
}

func (a *agent) FetchJWTSVID(ctx context.Context, req *workload.JWTSVIDRequest) (*workload.JWTSVIDResponse, error) {
	if a.idToken == "" || a.spiffeId == "" {
		return nil, status.Errorf(codes.FailedPrecondition, "id token is not set, please run Init first")
	}

	claims := jwt.Claims{
		Expiry:  jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		Issuer:  "your-issuer",
		Subject: a.spiffeId,
	}

	raw, err := jwt.Signed(a.signer).Claims(claims).CompactSerialize()
	if err != nil {
		return nil, err
	}

	return &workload.JWTSVIDResponse{
		Svids: []*workload.JWTSVID{
			{
				Svid:     raw,
				SpiffeId: a.spiffeId,
			},
		},
	}, nil
}

func constructSpiffeID(idToken string) (string, error) {
	jwt, err := jwt.ParseSigned(idToken)
	if err != nil {
		return "", fmt.Errorf("failed to parse id token: %v", err)
	}

	var claims struct {
		Subject string `json:"sub"`
	}
	if err := jwt.UnsafeClaimsWithoutVerification(&claims); err != nil {
		return "", fmt.Errorf("failed to extract claims from id token: %v", err)
	}

	return fmt.Sprintf("spiffe://example.org/%s", claims.Subject), nil
}
