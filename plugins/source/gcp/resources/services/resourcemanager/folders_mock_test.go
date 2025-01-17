// Code generated by codegen; DO NOT EDIT.

package resourcemanager

import (
	"context"
	"fmt"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"testing"

	"cloud.google.com/go/resourcemanager/apiv3"

	pb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v3"

	"google.golang.org/api/option"
)

func createFolders() (*client.Services, error) {
	fakeServer := &fakeFoldersServer{}
	l, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}
	gsrv := grpc.NewServer()
	pb.RegisterFoldersServer(gsrv, fakeServer)
	fakeServerAddr := l.Addr().String()
	go func() {
		if err := gsrv.Serve(l); err != nil {
			panic(err)
		}
	}()

	// Create a client.
	svc, err := resourcemanager.NewFoldersClient(context.Background(),
		option.WithEndpoint(fakeServerAddr),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(grpc.WithTransportCredentials(insecure.NewCredentials())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create grpc client: %w", err)
	}

	return &client.Services{
		ResourcemanagerFoldersClient: svc,
	}, nil
}

type fakeFoldersServer struct {
	pb.UnimplementedFoldersServer
}

func (f *fakeFoldersServer) ListFolders(context.Context, *pb.ListFoldersRequest) (*pb.ListFoldersResponse, error) {
	resp := pb.ListFoldersResponse{}
	if err := faker.FakeObject(&resp); err != nil {
		return nil, fmt.Errorf("failed to fake data: %w", err)
	}
	resp.NextPageToken = ""
	return &resp, nil
}

func TestFolders(t *testing.T) {
	client.MockTestHelper(t, Folders(), createFolders, client.TestOptions{})
}
