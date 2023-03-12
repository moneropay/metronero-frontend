package config

import "github.com/namsral/flag"

var (
	Bind      string
	Debug     bool
	Backend   string
	Uname     string
	JwtSecret string
)

func Load() {
	flag.StringVar(&Bind, "bind", ":5002", "Bind address")
	flag.BoolVar(&Debug, "debug", false, "Debug mode")
	flag.StringVar(&Backend, "backend", "http://localhost:5001", "Metronero backend host:port")
	flag.StringVar(&JwtSecret, "token-secret", "", "Secret for authentication tokens, same as backend")
	flag.Parse()
	Uname = "Siren"
}
