package main

import (
	"github.com/gclitheroe/grpc-exp/field"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestFieldAuth(t *testing.T) {
	c := field.NewFieldClient(conn)

	if _, err := c.DeviceSave(context.Background(), &field.Device{}); err != nil {
		t.Error(err)
	}

	c = field.NewFieldClient(connRead)

	if _, err := c.DeviceSave(context.Background(), &field.Device{DeviceID: "device-id"}); grpc.Code(err) != codes.Unauthenticated {
		t.Errorf("should get unuathenicated error %+v.", err)
	}

	c = field.NewFieldClient(connNoCreds)

	if _, err := c.DeviceSave(context.Background(), &field.Device{DeviceID: "device-id"}); grpc.Code(err) != codes.Unauthenticated {
		t.Errorf("should get unuathenicated error %+v.", err)
	}
}
