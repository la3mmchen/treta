package main

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/la3mmchen/treta/internal/commands"
	"github.com/la3mmchen/treta/internal/types"
)

var (
	configFile string = "config.json"
	// AppVersion Version of the app. Must be injected during the build.
	AppVersion string
	// Cfg types.Configuration
	Cfg types.Configuration
)

func main() {
	var Cfg = types.Configuration{
		AppUsage:   "A simple cli to print cards of a defined trello list",
		AppName:    "treta",
		AppVersion: AppVersion,
		Debug:      "false",
	}

	// load config if it is present
	if _, err := os.Stat(filepath.Join(".", configFile)); !os.IsNotExist(err) {
		file, err := os.Open(filepath.Join(".", configFile))
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&Cfg)
		if err != nil {
			os.Exit(1)
		}
	}

	app := commands.GetApp(Cfg)

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}