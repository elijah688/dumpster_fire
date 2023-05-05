package server

import (
	context "context"
	"items/domain/items"
	"items/service/db/db"
)

type Server struct {
	*items.UnimplementedItemsServiceServer
	pool *db.Pool
}

func New(pool *db.Pool) items.ItemsServiceServer {
	return &Server{
		new(items.UnimplementedItemsServiceServer),
		pool,
	}
}

func (s *Server) Create(ctx context.Context, in *items.Item) (*items.Item, error) {
	return s.pool.CreateItem(ctx, in)
}

func (s *Server) Update(ctx context.Context, in *items.Item) (*items.Item, error) {
	return s.pool.UpdateItem(ctx, in)
}

func (s *Server) Get(ctx context.Context, in *items.Item) (*items.Item, error) {
	return nil, nil
}

func (s *Server) Delete(ctx context.Context, in *items.Item) (*items.Item, error) {
	return nil, nil
}

func (s *Server) GetAll(ctx context.Context, in *items.Empty) (*items.PostList, error) {
	return nil, nil
}
