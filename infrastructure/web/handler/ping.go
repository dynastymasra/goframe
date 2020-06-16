package handler

import (
	"fmt"
	"net/http"

	"github.com/dynastymasra/goframe/config"

	"github.com/jinzhu/gorm"

	"github.com/dynastymasra/cookbook"
	"github.com/sirupsen/logrus"
)

func Ping(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		logrus.Info(r.Context().Value(config.RequestID))

		if err := db.DB().Ping(); err != nil {
			logrus.WithError(err).Errorln("Failed ping database")

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, cookbook.ErrorResponse(err.Error(), r.Context().Value(config.RequestID)))
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, cookbook.SuccessResponse().Stringify())
	}
}
