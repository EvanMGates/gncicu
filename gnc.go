package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/nlopes/slack"

	"gopkg.in/yaml.v2"
)

type Database struct {
	DNS string
}

type Server struct {
	Addr   string
	Assets string
}

type Config struct {
	Database Database
	Server   Server
}

type gnc struct {
	hs     *http.Server
	config *Config
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
	g := gnc{config: config,
		hs: hs,
	}
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/slackbot", slackbotHandler)
	http.HandleFunc("/fortnite/drop", g.fortniteDrop)
	log.Fatal(hs.ListenAndServe())
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "No. I see you, Game Night Crew!")
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
		params.ResponseType = "in_channel"
		params.Channel = "fortnite"
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

func (gnc *gnc) fortniteDrop(w http.ResponseWriter, r *http.Request) {
	landingZone := randomLandingZone()
	buffer := gnc.getLandingZoneImage(landingZone)
	buf := buffer.Bytes()
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buf)))
	if _, err := w.Write(buf); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	// Save landing zone
}
