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
	connConfig, err := pgxpool.ParseConfig("postgres://postgres:password@localhost:5432/postgres")
	if err != nil {
		return nil, fmt.Errorf("Unable to parse connection config: %v", err)
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return nil, fmt.Errorf("Unable to create connection pool: %v", err)
	}

	return &Pool{pool}, nil
}

func (p *Pool) CreateItem(ctx context.Context) {
	q := goqu.From("users").Select("id", "name")

	// Execute the query
	sql, params, err := q.ToSQL()
	if err != nil {
		fmt.Println(err)
		// return nil, err
	}

	rows, err := p.Query(ctx, sql, params)
	if err != nil {
		log.Fatalf("failed to execute query: %v", err)
	}

	fmt.Println(rows)
	// Query(context.Background(), "SELECT id, name FROM users")
	// if err != nil {
	// 	log.Fatalf("failed to execute query: %v", err)
	// }
	// defer rows.Close()

	for rows.Next() {

		item := new(items.Item)

		err := rows.Scan(&item.Id, &item.Title)
		if err != nil {
			log.Fatalf("failed to scan row: %v", err)
		}
		fmt.Printf("ITEM: %v", item)
	}
	if err := rows.Err(); err != nil {
		log.Fatalf("error iterating rows: %v", err)
	}
}
