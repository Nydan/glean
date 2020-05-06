package app

import (
	"context"
	"log"

	"github.com/nydan/glean/internal/config"
	"github.com/nydan/glean/internal/controller/http"
	ordctrl "github.com/nydan/glean/internal/controller/http/order"
	ordrp "github.com/nydan/glean/internal/repository/order"
	"github.com/nydan/glean/internal/server"
	orduc "github.com/nydan/glean/internal/usecase/order"
	"github.com/nydan/glean/pkg/database"
	"github.com/nydan/glean/pkg/database/sqldb"
)

// HTTPServer initializes all dependencies for HTTP server.
// The initialization that happen here are related to HTTP API.
func HTTPServer(cfg config.Config) error {
	db := connectDB(cfg.Database)

	orderRp := ordrp.Order(db)

	orderUc := orduc.Order(orderRp)

	orderCtrl := ordctrl.Order(orderUc)

	ctrl := http.NewController(orderCtrl)

	srv := server.NewHTTPServer(cfg.HTTPServer, ctrl)

	return srv.RunHTTP()
}

func connectDB(cfg config.Database) database.Database {
	ctx := context.Background()
	leaderDB, err := sqldb.Connect(ctx, "postgres", cfg.Master, &sqldb.ConnectionOptions{
		Retry: 1,
	})
	if err != nil {
		log.Fatal("failed to connect to leader DB", err)
	}

	followerDB, err := sqldb.Connect(ctx, "postgres", cfg.Replica, &sqldb.ConnectionOptions{
		Retry: 1,
	})
	if err != nil {
		log.Fatal("failed to connect to follower DB", err)
	}

	wrappedDB, err := sqldb.Wrap(ctx, leaderDB, followerDB)
	if err != nil {
		log.Fatal("failed to wrap DB connection", err)
	}

	return wrappedDB
}
