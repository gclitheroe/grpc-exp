package main

import (
	"github.com/gclitheroe/grpc-exp/field"
	"golang.org/x/net/context"
	"log"
)

// DeviceSave implements FieldServer.DeviceSave
func (s *server) DeviceSave(ctx context.Context, in *field.Device) (*field.Result, error) {
	if err := write(ctx); err != nil {
		return nil, err
	}

	log.Printf("saved device %s", in.DeviceID)

	return &field.Result{}, nil
}
