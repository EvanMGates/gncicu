package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nlopes/slack"

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
	fmt.Fprintln(w, "No. I see you, Game Night Crew. Maybe... Naw.")
}

func slackbotHandler(w http.ResponseWriter, r *http.Request) {
	const verificationToken = "2VMknV1iX74QsAUQlb5iS0ms"
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !s.ValidateToken(verificationToken) {
		w.WriteHeader(http.StatusUnauthorized)
	}

	switch s.Command {
	case "/drop":
		params := &slack.Msg{Text: fmt.Sprintf("*%s*!!", randomLandingZone())}
		b, err := json.Marshal(params)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func fortniteDrop(w http.ResponseWriter, r *http.Request) {
	landingZone := randomLandingZone()
	fmt.Fprintf(w, "*%s*!!", landingZone)
	// Save landingZone to RedisDB
}
