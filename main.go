package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dynastymasra/goframe/console"

	"github.com/dynastymasra/goframe/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/tylerb/graceful.v1"
)

func init() {
	config.Load()
	config.Logger().Setup()
}

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	log := logrus.WithFields(logrus.Fields{
		"serviceName": config.ServiceName,
		"version":     config.Version,
	})

	log.Infoln("Prepare start service")

	db, err := config.Postgres().Client()
	if err != nil {
		log.WithError(err).Fatalln("Failed connect to postgres")
	}

	migration, err := console.Migration(db)
	if err != nil {
		log.WithError(err).Fatalln("Failed run migration")
	}

	clientApp := cli.NewApp()
	clientApp.Name = config.ServiceName
	clientApp.Version = config.Version

	clientApp.Action = func(c *cli.Context) error {
		webServer := &graceful.Server{
			Timeout: 0,
		}

		select {
		case sig := <-stop:
			<-webServer.StopChan()

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

				if err := console.RunMigration(migration); err != nil {
					logrus.WithError(err).Errorln("Failed run database migration")
					os.Exit(1)
				}

				logrus.Infoln("Success run database migration ")

				return nil
			},
		}, {
			Name:        "migrate:rollback",
			Description: "Rollback database migration",
			Action: func(c *cli.Context) error {
				logrus.Infoln("Rollback database migration to previous version")

				if err := console.RollbackMigration(migration); err != nil {
					logrus.WithError(err).Errorln("Failed rollback database migration")
					os.Exit(1)
				}

				logrus.Infoln("Success rollback database migration to latest")

				return nil
			},
		}, {
			Name:        "migrate:create",
			Description: "Create up and down migration files with timestamp",
			Action: func(c *cli.Context) error {
				return console.CreateMigrationFiles(c.Args().Get(0))
			},
		},
	}

	if err := clientApp.Run(os.Args); err != nil {
		panic(err)
	}
}
