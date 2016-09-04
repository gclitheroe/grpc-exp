package main

import (
	"crypto/tls"
	tk "github.com/gclitheroe/grpc-exp/credentials/token"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"testing"
)

var testServer *grpc.Server
var conn *grpc.ClientConn
var connRead *grpc.ClientConn
var connNoCreds *grpc.ClientConn

// TestMain starts the testServer and connections to it.
// A self signed TLS cert is auto generated and verification
// is skipped on the client connections.
func TestMain(m *testing.M) {
	testServer = grpc.NewServer()

	register(testServer)

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
		log.Print("succesfully loaded TLS cert.")
	}

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
	}

	lis, err := tls.Listen("tcp", ":"+os.Getenv("MTR_PORT"), &config)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Print("starting test server")
	go testServer.Serve(lis)

	conn, err = grpc.Dial(os.Getenv("MTR_SERVER")+":"+os.Getenv("MTR_PORT"),
		grpc.WithPerRPCCredentials(tk.New(os.Getenv("MTR_TOKEN_WRITE"))),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{ServerName: "", InsecureSkipVerify: true})))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	connRead, err = grpc.Dial(os.Getenv("MTR_SERVER")+":"+os.Getenv("MTR_PORT"),
		grpc.WithPerRPCCredentials(tk.New(os.Getenv("MTR_TOKEN_READ"))),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{ServerName: "", InsecureSkipVerify: true})))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	connNoCreds, err = grpc.Dial(os.Getenv("MTR_SERVER")+":"+os.Getenv("MTR_PORT"),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{ServerName: "", InsecureSkipVerify: true})))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	code := m.Run()

	conn.Close()
	connRead.Close()
	connNoCreds.Close()
	testServer.Stop()

	os.Exit(code)
}
