package main

import (
	"github.com/gclitheroe/grpc-exp/data"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"testing"
)

func TestDataAuth(t *testing.T) {
	d := data.NewDataClient(conn)

	if _, err := d.SiteSave(context.Background(), &data.Site{}); err != nil {
		t.Error(err)
	}

	if _, err := d.SiteSearch(context.Background(), &data.SiteSearchRequest{}); err != nil {
		t.Error(err)
	}

	d = data.NewDataClient(connRead)

	if _, err := d.SiteSave(context.Background(), &data.Site{}); grpc.Code(err) != codes.Unauthenticated {
		t.Errorf("should get unuathenicated error %+v.", err)
	}

	if _, err := d.SiteSearch(context.Background(), &data.SiteSearchRequest{}); err != nil {
		t.Error(err)
	}

	d = data.NewDataClient(connNoCreds)

	if _, err := d.SiteSave(context.Background(), &data.Site{}); grpc.Code(err) != codes.Unauthenticated {
		t.Errorf("should get unuathenicated error %+v.", err)
	}

	if _, err := d.SiteSearch(context.Background(), &data.SiteSearchRequest{}); grpc.Code(err) != codes.Unauthenticated {
		t.Errorf("should get unuathenicated error %+v.", err)
	}
}
