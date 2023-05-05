package db

import (
	"context"
	"fmt"
	"items/domain/items"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
}

func New() (*Pool, error) {
	connConfig, err := pgxpool.ParseConfig("postgres://postgres:postgres@localhost:5432/items_db")
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection config: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to create connection pool: %v", err)
	}

	_, err = pool.Exec(context.Background(), `
		CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		CREATE TABLE IF NOT EXISTS items (
			id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			src TEXT NOT NULL
		)
	`)

	if err != nil {
		log.Fatalf("failed to create table: %v", err)
	}

	return &Pool{pool}, nil
}

func (p *Pool) CreateItem(ctx context.Context, in *items.Item) (*items.Item, error) {
	insert := goqu.Insert("items").Rows(
		goqu.Record{
			"title":   in.Title,
			"content": in.Content,
			"src":     in.Src,
		},
	).Returning("*")

	sql, _, err := insert.ToSQL()
	if err != nil {
		return nil, err
	}

	row := p.QueryRow(ctx, sql)
	item := new(items.Item)
	err = row.Scan(&item.Id, &item.Title, &item.Content, &item.Src)
	if err != nil {
		log.Printf("failed to scan created item: %v", err)
		return nil, err
	}

	return item, nil
}
func (p *Pool) UpdateItem(ctx context.Context, in *items.Item) (*items.Item, error) {
	update := goqu.Update("items").
		Set(goqu.Record{
			"title":   in.Title,
			"content": in.Content,
			"src":     in.Src,
		}).
		Where(goqu.C("id").Eq(in.Id))

	sql, _, err := update.ToSQL()
	if err != nil {
		return nil, err
	}

	result, err := p.Exec(ctx, sql)
	if err != nil {
		log.Printf("failed to execute update statement: %v", err)
		return nil, err
	}

	if result.RowsAffected() == 0 {
		return nil, fmt.Errorf("no rows were updated")
	}

	return in, nil
}

func (p *Pool) GetAll(ctx context.Context, in *items.Empty) (*items.ItemsList, error) {

	sql, _, err := goqu.From("items").ToSQL()
	if err != nil {
		return nil, err
	}

	rows, err := p.Query(ctx, sql)
	if err != nil {
		log.Printf("failed to execute get all statement: %v", err)
		return nil, err
	}
	defer rows.Close()

	list := new(items.ItemsList)
	for rows.Next() {
		item := new(items.Item)
		err = rows.Scan(&item.Id, &item.Title, &item.Content, &item.Src)
		if err != nil {
			log.Printf("failed to scan created item: %v", err)
			return nil, err
		}

		list.Items = append(list.Items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during iteration of rows: %v", err)
	}

	return list, nil
}

func (p *Pool) Get(ctx context.Context, in *items.Item) (*items.Item, error) {

	sql, args, err := goqu.From("items").
		Select("*").
		Where(goqu.C("id").
			Eq(in.Id)).ToSQL()
	if err != nil {
		return nil, err
	}

	item := new(items.Item)
	if err := p.QueryRow(ctx, sql, args...).Scan(&item.Id, &item.Title, &item.Content, &item.Src); err != nil {
		log.Printf("failed to scan created item: %v", err)
		return nil, err
	}

	return item, nil

}

func (p *Pool) Delete(ctx context.Context, in *items.Item) (*items.Item, error) {

	sql, args, err := goqu.Delete("items").
		Where(goqu.C("id").
			Eq(in.Id)).Returning("*").ToSQL()

	if err != nil {
		return nil, err
	}

	item := new(items.Item)
	if err := p.QueryRow(ctx, sql, args...).Scan(&item.Id, &item.Title, &item.Content, &item.Src); err != nil {
		log.Printf("failed to scan created item: %v", err)
		return nil, err
	}

	return item, nil

}
