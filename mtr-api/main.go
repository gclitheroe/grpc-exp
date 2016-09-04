package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"github.com/GeoNet/mtr/mtrapp"
	"github.com/gclitheroe/grpc-exp/data"
	"github.com/gclitheroe/grpc-exp/field"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
	"math/big"
	"os"
	"time"
)

var tokenWrite = os.Getenv("MTR_TOKEN_WRITE")
var tokenRead = os.Getenv("MTR_TOKEN_READ")

// server is used to implement field.FieldServer and data.DataServer
type server struct{}

func init() {
	switch "" {
	case tokenWrite:
		log.Panic("empty write token")
	case tokenRead:
		log.Panic("empty read token")
	}
}

func main() {
	var cert tls.Certificate
	var err error

	if cert, err = tls.LoadX509KeyPair("certs/server.crt", "certs/server.key"); err != nil {
		log.Printf("failed to read TLS certs, will generate self signed cert: %s", err.Error())
		if cert, err = selfie(); err != nil {
			log.Fatalf("failed to generate self signed TLS cert: %v", err)
		} else {
			log.Print("succesfully generated self signed TLS cert.")
		}
	} else {
		log.Print("succesfully loaded TLS cert from /certs")
	}

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	lis, err := tls.Listen("tcp", ":"+os.Getenv("MTR_PORT"), &config)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(telemetry))

	register(s)

	log.Print("starting server")
	log.Fatal(s.Serve(lis))
}

func register(s *grpc.Server) {
	field.RegisterFieldServer(s, &server{})
	data.RegisterDataServer(s, &server{})
}

func token(ctx context.Context) string {
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return ""
	}

	t := md["token"]

	if t == nil || len(t) != 1 {
		return ""
	}

	return t[0]
}

func write(ctx context.Context) error {
	switch token(ctx) {
	case tokenWrite:
		return nil
	default:
		return grpc.Errorf(codes.Unauthenticated, "valid write token required.")
	}
}

func read(ctx context.Context) error {
	switch token(ctx) {
	case tokenWrite, tokenRead:
		return nil
	default:
		return grpc.Errorf(codes.Unauthenticated, "valid read or write token required.")
	}
}

// selfie generates a self signed TLS certificate.
func selfie() (tls.Certificate, error) {
	ca := &x509.Certificate{
		SerialNumber: big.NewInt(1337),
		Subject: pkix.Name{
			Organization: []string{"seflie"},
		},
		SignatureAlgorithm:    x509.SHA512WithRSA,
		PublicKeyAlgorithm:    x509.ECDSA,
		NotBefore:             time.Now().AddDate(-1, 0, 0),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		BasicConstraintsValid: true,
		IsCA:        true,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	p, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return tls.Certificate{}, err
	}

	b, err := x509.CreateCertificate(rand.Reader, ca, ca, &p.PublicKey, p)
	if err != nil {
		return tls.Certificate{}, err
	}

	return tls.Certificate{
		Certificate: [][]byte{b},
		PrivateKey:  p,
	}, nil
}

// telemetry is a UnaryServerInterceptor.
// adds timing and metrics.
func telemetry(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	t := mtrapp.Start()

	i, err := handler(ctx, req)

	mtrapp.Requests.Inc()

	if err == nil {
		t.Track(info.FullMethod)

		mtrapp.StatusOK.Inc()

		if t.Taken() > 250 {
			log.Printf("%s took %d (ms)", info.FullMethod, t.Taken())
		}
	} else {
		log.Printf("%s error %s", info.FullMethod, err)

		// Remap the grpc codes to the existing (http based) mtr counters.
		// Could add mtr counters for grpc.
		switch grpc.Code(err) {
		case codes.InvalidArgument:
			mtrapp.StatusBadRequest.Inc()
		case codes.Unauthenticated:
			mtrapp.StatusUnauthorized.Inc()
		case codes.NotFound:
			mtrapp.StatusNotFound.Inc()
		case codes.FailedPrecondition:
			mtrapp.StatusInternalServerError.Inc()
		case codes.Unavailable:
			mtrapp.StatusServiceUnavailable.Inc()
		}
	}

	return i, err
}
