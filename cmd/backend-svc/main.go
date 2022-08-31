package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/rollbar/rollbar-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend-svc-template/config"
	"backend-svc-template/core"
	"backend-svc-template/middleware"
	"backend-svc-template/server"
	"backend-svc-template/storage"
)

func main() {
	cfg := config.LoadConfig()

	binding.Validator = new(middleware.DefaultValidator)

	rollbar.SetToken(*cfg.RollbarConfig.RollbarToken)
	rollbar.SetEnvironment(*cfg.RollbarConfig.RollbarEnv)
	rollbar.SetCodeVersion(*cfg.CommitSHA)
	rollbar.SetServerHost(*cfg.ServiceConfig.ServiceRevision)
	rollbar.SetServerRoot(*cfg.RollbarConfig.RollbarServerRoot)

	firebaseAuth := middleware.NewFirebaseAuth(&middleware.FirebaseConfig{
		ProjectID:        *cfg.ProjectID,
		ServiceAccountID: *cfg.FirebaseConfig.ServiceAccountID,
	})

	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(*cfg.MongoDBConfig.MongoDBURI))
	if err != nil {
		log.Fatalf("failed to connect to mongo cluster, err: %+v", err)
	}

	bs := storage.NewBeer(mongoClient)
	rs := storage.NewReview(mongoClient)

	srvr := server.New(core.NewBeer(bs), core.NewReview(rs), firebaseAuth)

	srv := http.Server{
		Addr:              fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:           srvr.Mux(),
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM)

	go func() {
		log.Print(srv.ListenAndServe())
	}()
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer func() {
		if err := bs.Close(ctx); err != nil {
			panic(err)
		}

		if err := rs.Close(ctx); err != nil {
			panic(err)
		}

		cancel()
	}()

	_ = srv.Shutdown(ctx)
}
