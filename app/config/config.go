package config

import "github.com/namsral/flag"

var (
        Bind string
        Debug bool
        Backend string
)

func init() {
	flag.StringVar(&Bind, "bind", ":5002", "Bind address")
	flag.BoolVar(&Debug, "debug", false, "Debug mode")
	flag.StringVar(&Backend, "backend", ":5001", "Metronero backend host:port")
	flag.Parse()
}
