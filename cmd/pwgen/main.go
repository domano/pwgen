package main

import (
	"context"
	handler "github.com/domano/pwgen/internal/http"
	"github.com/domano/pwgen/internal/password"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Config struct {
	certFile string
	keyFile string
	port int
	shutdownTimeout time.Duration
}

var config Config

func main() {
	config = Config{
		"cert.pem",
		"key.unencrypted.pem",
		8443,
		5*time.Second,
	}
	run()
}

func run() {
	log.Infoln("Starting pwgen...")

	// Listen for SIGINT to gracefully close the app
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Create a new password handler using our single use PasswordAdapter
	ph := handler.NewPasswordHandler(handler.PassworderFunc(PasswordAdapter))

	server := createServer(ph, "/passwords")
	startServer(server)

	// Wait for SIGINT
	<-stop
	log.Infoln("pwgen shuts down now.")

	// Trigger Graceful shutdown with 5 second time limit
	ctx, _ := context.WithTimeout(context.Background(), config.shutdownTimeout)
	err := server.Shutdown(ctx)
	if err != nil {
		log.WithError(err).Fatalln("pwgen failed during graceful shutdown")
	}
	log.Infoln("pwgen gracefully shut down.")
}

// Starts the server in its own goroutine
func startServer(server http.Server) {
	go func() {
		err := server.ListenAndServeTLS(config.certFile, config.keyFile)
		if err != nil {
			log.WithError(err).Fatal("HTTP Server threw an error, shutting down.")
		}
	}()
}

func createServer(h http.Handler, route string) http.Server {
	// Wrap the handler with all necessary middlewares
	lh := handler.LoggingHandlerFunc(h)

	// Route /passwords to our handler chain
	mux := http.NewServeMux()
	mux.Handle(route, lh)

	return http.Server{Addr: ":8443", Handler: mux}
}

// PasswordAdapter allows us to use a password
// generator to fullfill the Passworder-interface for our handler
func PasswordAdapter(minLength, specialChars, numbers int) string {
	generator := password.NewGenerator(
		password.MinLength(minLength),
		password.SpecialChars(specialChars),
		password.Numbers(numbers))

	return generator.Password()
}
