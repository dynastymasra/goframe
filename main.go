package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dynastymasra/cookbook/provider"

	"github.com/gorilla/handlers"

	"github.com/dynastymasra/goframe/infrastructure/web"

	"github.com/dynastymasra/cookbook/provider/postgres"

	"github.com/dynastymasra/goframe/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func init() {
	config.Load()
	config.Logger().Setup()
}

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	log := logrus.WithFields(logrus.Fields{
		config.JServiceName: config.ServiceName,
		config.JVersion:     config.Version,
		"web":               config.ServerPort(),
	})

	log.Infoln("Prepare start service")

	// Remove this database if not used
	db, err := config.Postgres().Client()
	if err != nil {
		log.WithError(err).Fatalln("Failed connect to Postgres")
	}

	// Remove this database if not used
	//driver, err := config.Neo4J().Driver()
	//if err != nil {
	//	log.WithError(err).Fatalln("Failed connect to Neo4J")
	//}

	// Remove this database if not used
	//client, err := config.MongoDB().Client()
	//if err != nil {
	//	log.WithError(err).Fatalln("Failed connect to MongoDB")
	//}

	// Remove this database if not used
	//esClient, err := config.Elasticsearch().Client()
	//if err != nil {
	//	log.WithError(err).Fatalln("Failed connect to Elasticsearch")
	//}

	// Remove this database if not used
	//redisClient, err := config.Redis().Client()
	//if err != nil {
	//	log.WithError(err).Fatalln("Failed connect to Redis")
	//}

	clientApp := cli.NewApp()
	clientApp.Name = config.ServiceName
	clientApp.Version = config.Version

	clientApp.Action = func(c *cli.Context) error {
		router := &web.RouterInstance{}

		srv := &http.Server{
			Addr: fmt.Sprintf(":%s", config.ServerPort()),
			Handler: handlers.RecoveryHandler(
				handlers.PrintRecoveryStack(true),
				handlers.RecoveryLogger(log),
			)(router.Router()),
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil {
				log.WithError(err).Fatalln("Failed start web application")
			}
		}()

		log.Infoln("Web application is running")

		select {
		case sig := <-stop:
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				log.WithError(err).Fatalln("Failed shutdown web application")
			}

			log.Warnln(fmt.Sprintf("Service shutdown because %+v", sig))
			os.Exit(0)
		}

		return nil
	}

	clientApp.Commands = []*cli.Command{
		{
			Name:        "migrate:run",
			Description: "Running database migration",
			Action: func(c *cli.Context) error {
				logrus.Infoln("Start database migration")

				migration, err := postgres.Migration(db)
				if err != nil {
					log.WithError(err).Fatalln("Failed run migration")
				}

				if err := provider.RunMigration(migration); err != nil {
					logrus.WithError(err).Fatalln("Failed run database migration")
				}

				logrus.Infoln("Success run database migration ")

				return nil
			},
		}, {
			Name:        "migrate:rollback",
			Description: "Rollback database migration",
			Action: func(c *cli.Context) error {
				logrus.Infoln("Rollback database migration to previous version")

				migration, err := postgres.Migration(db)
				if err != nil {
					log.WithError(err).Fatalln("Failed run migration")
				}

				if err := provider.RollbackMigration(migration); err != nil {
					logrus.WithError(err).Fatalln("Failed rollback database migration")
				}

				logrus.Infoln("Success rollback database migration to latest")

				return nil
			},
		}, {
			Name:        "migrate:create",
			Description: "Create up and down migration files with timestamp",
			Action: func(c *cli.Context) error {
				return postgres.CreateMigrationFiles(c.Args().Get(0))
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		log.WithError(err).Fatalln("Failed start application")
	}
}
