package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/djpiper28/setmoji/generator/logger"
	scryfallclient "github.com/djpiper28/setmoji/generator/scryfall_client"
)

const OutputDirectory = "../sets/"

func main() {
	logger.Logger.Info("Starting set generator...")

	logger.Logger.Info("Getting sets...")
	sets, err := scryfallclient.GetSets()
	if err != nil {
		logger.Logger.Fatal("Cannot get sets", "error", err)
	}

	for i, set := range sets {
		filename := filepath.Join(OutputDirectory, set.Code+".svg")
		_, err := os.Lstat(filename)
		if err == nil {
			continue
		}

		logger.Logger.Info("File does not exist - downloading",
			"filename", filename,
			"code", set.Code,
			"setNumber", i,
			"sets", len(sets))

		svgBytes, err := set.GetSvg()
		if err != nil {
			logger.Logger.Fatal("Cannot get svg for set",
				"error", err,
				"code", set.Code,
				"setNumber", i,
				"sets", len(sets))
		}

		err = os.WriteFile(filename, svgBytes, 0777)
		if err != nil {
			logger.Logger.Fatal("Cannot save svg for set",
				"error", err,
				"code", set.Code,
				"setNumber", i,
				"sets", len(sets))
		}

		logger.Logger.Info("Saved set svg",
			"code", set.Code,
			"setNumber", i,
			"sets", len(sets))

		logger.Logger.Debug("Sleeping to avoid rate limit banning...")
		time.Sleep(time.Millisecond * 100)
	}
}
