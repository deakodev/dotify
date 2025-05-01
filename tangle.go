package main

import (
	"github.com/deakodev/go-remedy/internal/dot"
	"github.com/deakodev/go-remedy/internal/repl"
	// "github.com/deakodev/go-remedy/internal/session"
)

func main() {
	// var session session.Session
	registry := dot.MakeRegistry()
	repl.Run(registry)
}
