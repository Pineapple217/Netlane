package database

import (
	"context"
	"log/slog"

	"entgo.io/ent/dialect"
	"github.com/Pineapple217/Netlane/pkg/config"
	"github.com/Pineapple217/Netlane/pkg/ent"
	"github.com/Pineapple217/Netlane/pkg/util"
	_ "github.com/mattn/go-sqlite3"
)

func NewDatabaseClient(c config.Database) *ent.Client {
	slog.Info("Starting database", "connection string", c.Sting)
	client, err := ent.Open(dialect.SQLite, c.Sting)
	util.MaybeDie(err, "Failed to start database")

	err = client.Schema.Create(context.Background())
	util.MaybeDie(err, "Failed to create database schemas")
	return client
}
