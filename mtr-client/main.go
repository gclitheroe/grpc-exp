package main

import (
	"crypto/tls"
	"github.com/gclitheroe/grpc-exp/credentials/token"
	"github.com/gclitheroe/grpc-exp/data"
	"github.com/gclitheroe/grpc-exp/field"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"os"
	"time"
)

// use InsecureSkipVerify to ignore server identity with a self signed certificate.
var opts = []grpc.DialOption{
	grpc.WithPerRPCCredentials(token.New(os.Getenv("MTR_TOKEN_WRITE"))),
	grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{ServerName: "", InsecureSkipVerify: true})),
}

func main() {
	conn, err := grpc.Dial(os.Getenv("MTR_SERVER")+":"+os.Getenv("MTR_PORT"), opts...)
	if err != nil {
		log.Printf("did not connect: %v", err)
	}
	defer conn.Close()

	c := field.NewFieldClient(conn)
	d := data.NewDataClient(conn)

	for {
		if _, err := c.DeviceSave(context.Background(), &field.Device{DeviceID: "test-device"}); err != nil {
			log.Printf("could not save device: %v", err)
		}

		if _, err := d.SiteSave(context.Background(), &data.Site{SiteID: "TAUP"}); err != nil {
			log.Printf("could not save site: %v", err)
		}
		time.Sleep(time.Second)
	}
}
