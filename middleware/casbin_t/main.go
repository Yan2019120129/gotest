package main

import (
	"log"
	"os"

	"gotest/middleware/casbin_t/bootstrap"
)

func main() {
	app, err := bootstrap.NewApp()
	if err != nil {
		log.Fatalf("initialize casbin demo: %v", err)
	}

	addr := os.Getenv("CASBIN_T_ADDR")
	if addr == "" {
		addr = ":3001"
	}
	if err := app.Listen(addr); err != nil {
		log.Fatalf("listen on %s: %v", addr, err)
	}

}
