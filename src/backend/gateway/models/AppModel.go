package models

import (
	"context"

	conn "gateway/db"

	"github.com/go-chi/chi/v5"
	"github.com/uptrace/bun"
)

type (
	MuxServer struct {
		Mux      *chi.Mux
		Endpoint string
		Proxy    string
	}

	ErrorObject struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
)

func Create(model interface{}) (q *bun.InsertQuery, ctx context.Context) {
	ctx = context.Background()

	db := conn.InitDBConnect()
	q = db.NewInsert().Model(model)
	return
}

func Read(model interface{}) (q *bun.SelectQuery, ctx context.Context) {
	ctx = context.Background()

	db := conn.InitDBConnect()
	q = db.NewSelect().Model(model)
	return
}

func Update(model interface{}) (q *bun.UpdateQuery, ctx context.Context) {
	ctx = context.Background()

	db := conn.InitDBConnect()
	q = db.NewUpdate().Model(model)
	return
}

func Delete(model interface{}) (q *bun.DeleteQuery, ctx context.Context) {
	ctx = context.Background()

	db := conn.InitDBConnect()
	q = db.NewDelete().Model(model)
	return
}
