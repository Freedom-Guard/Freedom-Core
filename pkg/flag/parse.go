package flags

import (
	"flag"
	"fmt"
	"os"

	"github.com/Freedom-Guard/freedom-core/pkg/logger"
	"github.com/Freedom-Guard/freedom-core/pkg/updater"
)

var AppVersion = "0.13.0"

func Parse() {
	host := flag.String("host", "127.0.0.1", "Server host")
	port := flag.Int("port", 8087, "Server port")
	showVersion := flag.Bool("version", false, "Show Freedom Core version")
	update := flag.Bool("update", false, "Update Freedom Core to the latest version")
	updateCores := flag.Bool("update-cores", false, "Update Cores to the latest version")

	flag.Parse()

	if *showVersion {
		fmt.Println("Freedom Core by Freedom Guard")
		fmt.Printf("Version: %s\n", AppVersion)
		fmt.Println("GitHub: https://github.com/Freedom-Guard/Freedom-Core")
		os.Exit(0)
	}

	if *update {
		logger.Log(logger.INFO, "Checking for updates...")
		if err := updater.Update(); err != nil {
			logger.Log(logger.INFO, "Update failed:"+err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	if *updateCores {
		logger.Log(logger.INFO, "Checking for update cores...")
		if err := updater.DeleteCores(); err != nil {
			logger.Log(logger.INFO, "Update failed:"+err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	AppConfig = &Config{
		Host:    *host,
		Port:    *port,
		Version: AppVersion,
	}
}
