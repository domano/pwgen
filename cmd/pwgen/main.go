package main

import (
	"context"
	"github.com/caarlos0/env/v6"
	handler "github.com/domano/pwgen/internal/http"
	"github.com/domano/pwgen/internal/password"
	"github.com/gorilla/handlers"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type config struct {
	CertFile    string        `env:"CERT_FILE" envDefault:"cert.pem"`
	KeyFile     string        `env:"KEY_FILE" envDefault:"key.unencrypted.pem"`
	Port        int           `env:"PORT" envDefault:"8443"`
	GracePeriod time.Duration `env:"GRACE_PERIOD" envDefault:"5s"`
}

var cfg config

func main() {
	err := parseConfig()
	if err != nil {
		log.WithError(err).Fatalln("Could not parse config, shutting down.")
	}
	// Listen for SIGINT to gracefully close the app
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	err = run(stop)
	if err != nil {
		log.WithError(err).Fatalln("Could not start app. Shutting down.")
	}
}

func parseConfig() error {
	cfg = config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.Wrap(err, "Failed while reading environment variables: ")
	}
	return nil
}

func run(stop chan os.Signal) error {
	log.Infoln("Starting pwgen...")

	// Create a new password handler using our single use PasswordAdapter
	ph := handler.NewPasswordHandler(handler.PassworderFunc(PasswordAdapter))

	server := createServer(ph, "/passwords")
	errChan := startServer(&server)

	// Wait for SIGINT or server error
	select {
	case err := <-errChan:
		return errors.Wrap(err, "Could not start server: ")
	case <-stop:
		log.Infoln("pwgen shuts down now.")

		// Trigger Graceful shutdown with 5 second time limit
		ctx, ctxCancel := context.WithTimeout(context.Background(), cfg.GracePeriod)
		err := server.Shutdown(ctx)
		if err != nil {
			ctxCancel()
			return errors.Wrap(err, "pwgen failed during graceful shutdown")
		}
		log.Infoln("pwgen gracefully shut down.")
		return nil
	}
}

// Starts the server in its own goroutine
func startServer(server *http.Server) <-chan error {
	errChan := make(chan error, 0)
	go func() {
		err := server.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		if err != nil {
			errChan <- errors.Wrap(err, "HTTP Server threw an error, shutting down: ")
		}
	}()
	return errChan
}

func createServer(h http.Handler, route string) http.Server {
	// Wrap the handler with all necessary middlewares
	lh := handler.LoggingHandlerFunc(h)

	// Route /passwords to our handler chain
	mux := http.NewServeMux()
	mux.Handle(route, lh)

	// Add a recovery handler in case anything unexpected happens
	rh := handlers.RecoveryHandler(handlers.RecoveryLogger(log.StandardLogger()), handlers.PrintRecoveryStack(true))(mux)

	return http.Server{Addr: ":8443", Handler: rh}
}

// PasswordAdapter allows us to use a password
// generator to fulfill the Passworder-interface for our handler
func PasswordAdapter(amount, minLength, specialChars, numbers int, swap bool) (passwords []string) {
	generator := password.NewGenerator(
		password.MinLength(minLength),
		password.SpecialChars(specialChars),
		password.Numbers(numbers),
		password.Swap(swap))

	for i := 0; i < amount; i++ {
		passwords = append(passwords, generator.Password())
	}
	return passwords
}
