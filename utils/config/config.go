package config

import "github.com/namsral/flag"

type Config struct {
        Bind string
        Debug bool
        Backend string
	Uname string
}

func Load() Config {
	var cfg Config
	flag.StringVar(&cfg.Bind, "bind", ":5002", "Bind address")
	flag.BoolVar(&cfg.Debug, "debug", false, "Debug mode")
	flag.StringVar(&cfg.Backend, "backend", ":5001", "Metronero backend host:port")
	flag.Parse()
	cfg.Uname = "Siren"
	return cfg
}
