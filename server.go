package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		DSN string
	}
	Server struct {
		Addr string
	}
}

func main() {
	filename := flag.String("config", "config.yml", "Configuration file")
	flag.Parse()
	config, err := load(*filename)
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}
	createServer(&config)
}

func load(file string) (cfg Config, err error) {
	cfg.Server.Addr = ":8080"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return
	}
	return
}

func createServer(config *Config) {
	hs := &http.Server{
		Addr: config.Server.Addr,
	}
	http.HandleFunc("/", rootHandler)
	log.Fatal(hs.ListenAndServe())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "No. I see you, Game Night Crew.")
}
