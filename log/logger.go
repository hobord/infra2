package log

import (
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
)

var Logger *logrus.Entry

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	Logger = logrus.WithFields(logrus.Fields{
		// "service": "go-reverse-proxy",
	})
}

func HttpLoggerHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Logger.Println("HTTP Request Log")
		next.ServeHTTP(w, r) // call original
	})
}
