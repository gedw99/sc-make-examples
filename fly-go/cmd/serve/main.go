/*
Serve is a very simple static file server in go
Usage:
	-p="8100": port to serve on
	-d=".":    the directory of static files to host
Navigating to http://localhost:8100 will display the index.html or directory
listing file.
*/
package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/sethvargo/go-envconfig"
)

func main() {

	/*
		TODO
		all below needs have excellent examples at: https://github.com/sethvargo

		config from env variables.

		embed web assets and handler to laod them: https://github.com/google/exposure-notifications-verification-server/blob/main/assets/assets.go
		deploy to fly and googel clodu run.
		github actions to deplyo to fly and google cloud run
		https://github.com/mutagen-io/mutagen to debug live server
	*/

	ctx := context.Background()

	var c Config
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	port := flag.String("p", "8100", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

type Config struct {
	Port      int    `env:"PORT, default=8100"`
	Directory string `env:"DIRECTORY"`
}
