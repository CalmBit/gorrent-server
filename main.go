package main

import (
	"bytes"
	"os"

	"flag"

	"sync"

	"github.com/CalmBit/gorrent-server/file"
	logging "github.com/op/go-logging"
	yaml "gopkg.in/yaml.v2"
)

const gorrentServerVersion = "0.0.1"

var verboseStartup bool

var log = logging.MustGetLogger("gorrent-server")

var format = logging.MustStringFormatter(
	`%{color}[%{time:2006-01-02 15:04:05.000}] [%{level}] [%{shortfunc}]%{color:reset} %{message}`,
)

func main() {

	/* INITIAL FLAG SETUP/PARSING */

	flag.BoolVar(&verboseStartup, "V", false, "Enables verbose startup (let's see those config options)")

	flag.Parse()

	/* STARTUP */

	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Format := logging.NewBackendFormatter(backend1, format)

	backend1LF := logging.AddModuleLevel(backend1Format)

	var loggingLevel logging.Level

	if verboseStartup {
		loggingLevel = logging.DEBUG
	} else {
		loggingLevel = logging.INFO
	}

	backend1LF.SetLevel(loggingLevel, "")

	logging.SetBackend(backend1LF)

	log.Info("Gorrent Server v.", gorrentServerVersion, " starting up...")

	config := file.ConfigDict{}

	configErr := loadConfiguration(&config, "config/config.yml")

	if configErr != nil {
		log.Critical(configErr)
		return
	}

	if verboseStartup {
		log.Debug("Configuration loaded:")
		log.Debug("-----------------------")
		log.Debug("Min/Max Worker Count: ", config.Workers.MinCount, "/", config.Workers.MaxCount)
	}

	var wg sync.WaitGroup

	for i := 0; i < config.Workers.MinCount; i++ {
		go commandHandler(&wg)
		wg.Add(1)
	}

	wg.Wait()

	log.Info("Bye!")

}

func commandHandler(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Debug("Command Handler Created")
}

func loadConfiguration(conf *file.ConfigDict, configPath string) error {
	f, err := os.Open(configPath)

	if err != nil {
		return err
	}

	buff := bytes.Buffer{}
	buff.ReadFrom(f)

	merr := yaml.Unmarshal(buff.Bytes(), conf)

	return merr
}
