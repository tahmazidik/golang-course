package collector

import (
	"context"
	"log/slog"

	"repo-stat/processor/internal/domain"
	collectorpb "repo-stat/proto/collector"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn *grpc.ClientConn
	pb   collectorpb.CollectorClient
	log  *slog.Logger
}

func NewClient(address string, log *slog.Logger) (*Client, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Error("failed to create gRPC client", "error", err)
		return nil, err
	}

	return &Client{
		conn: conn,
		pb:   collectorpb.NewCollectorClient(conn),
		log:  log,
	}, nil
}

func (c *Client) GetRepositoryInfo(ctx context.Context, url string) (*domain.RepositoryInfo, error) {
	resp, err := c.pb.GetRepositoryInfo(ctx, &collectorpb.GetRepositoryInfoRequest{
		Url: url,
	})
	if err != nil {
		c.log.Error("failed to get repository info", "error", err, "url", url)
		return nil, err
	}
	return &domain.RepositoryInfo{
		FullName:    resp.FullName,
		Description: resp.Description,
		Stars:       resp.Stars,
		Forks:       resp.Forks,
		CreatedAt:   resp.CreatedAt.AsTime(),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
