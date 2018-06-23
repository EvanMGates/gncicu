package main

import (
	"encoding/json"
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

type gnc struct {
	hs http.Server
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
	http.HandleFunc("/slackbot", slackbotHandler)
	http.HandleFunc("/fortnite/drop", fortniteDrop)
	log.Fatal(hs.ListenAndServe())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "No. I see you, Game Night Crew. Maybe.")
}

func slackbotHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var data slackData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(data.Challenge))
}

func fortniteDrop(w http.ResponseWriter, r *http.Request) {
	landingZone := randomLandingZone()
	fmt.Fprintf(w, "<html>We're dropping at <b>%s</b></html>", landingZone)
	// Save landingZone to RedisDB
}
