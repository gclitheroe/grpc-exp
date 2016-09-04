package main

import (
	"github.com/gclitheroe/grpc-exp/data"
	"golang.org/x/net/context"
	"log"
)

// SiteSave implements DataServer.SiteSave
func (s *server) SiteSave(ctx context.Context, in *data.Site) (*data.Result, error) {
	if err := write(ctx); err != nil {
		return nil, err
	}

	log.Printf("saved site %s", in.SiteID)

	return &data.Result{}, nil
}

// SiteSaveRequest implements DataServer.SiteSearch
func (s *server) SiteSearch(ctx context.Context, in *data.SiteSearchRequest) (*data.SiteSearchResult, error) {
	if err := read(ctx); err != nil {
		return nil, err
	}

	log.Print("site search request")

	return &data.SiteSearchResult{}, nil
}
