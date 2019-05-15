package main

import (
	handler "github.com/domano/pwgen/internal/http"
	"github.com/domano/pwgen/internal/password"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	log.Infoln("Starting pwgen...")
	defer log.Infoln("pwgen shuts down now.")

	// Create a new password handler using our single use PasswordAdapter
	ph := handler.NewPasswordHandler(handler.PassworderFunc(PasswordAdapter))

	// Wrap the password handler with all necessary middlewares
	h := handlers.LoggingHandler(os.Stdout, ph)
	http.Handle("/passwords", h)
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.unencrypted.pem", nil)
	if err != nil {
		log.WithError(err).Fatal("HTTP Server threw an error, shutting down.")
	}
}


func PasswordAdapter(minLength, specialChars, numbers int) string {
	generator := password.NewGenerator(
			password.MinLength(minLength),
			password.SpecialChars(specialChars),
			password.Numbers(numbers))

	return generator.Password()


}