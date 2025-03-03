package main

import (
	"flag"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"

	"ordo-map/arango"
	"ordo-map/settings"
	"ordo-map/util/logging"

	"github.com/rs/zerolog/log"
)

type Flags = struct {
	settingsFile string
}

var flags Flags

func main() {

	// The server recovery function
	defer func() {
		r := recover()
		if r != nil {
			log.Error().Err(r.(error)).Bytes("stack", debug.Stack()).Msgf("Recovering panic")
		}
	}()

	// Parse flags
	parseFlags()

	// Load settings
	settings.Load(flags.settingsFile)

	// Configure logging
	logging.Configure()

	// Initialize ArangoDB
	arango.Init()

	// Start the server
	log.Info().Msgf("########## Magic Helper Server Startup ##########")
	log.Info().Msgf("########## running %v on %v CPUs ##########", runtime.Version(), runtime.NumCPU())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "¡Hola desde OrdoMap Backend!")
	})

	fmt.Println("Servidor backend en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}

func parseFlags() {
	flag.StringVar(&flags.settingsFile, "settings", "", "Determines the file from which the settings are loaded.")
	flag.Parse()
}
