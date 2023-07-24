package main

import (
	"fmt"
	"github.com/Gallrecon-Tech/email-pwnd-checker/internal/check_runner"
	cfg "github.com/Gallrecon-Tech/email-pwnd-checker/internal/config"
	"log"
	"os"
)

func parseArgs() string {
	if len(os.Args) < 2 {
		log.Println("Missing email address to scan")
		fmt.Printf("Usage: %s <email>\n", os.Args[0])
		os.Exit(1)
	}
	return os.Args[1]
}

func main() {
	log.Println("Welcome in Email Pwnd Checker by GALLRECON.TECH")
	emailToScan := parseArgs()
	config, err := cfg.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}

	runner := check_runner.Runner{}
	runner.CheckEmail(emailToScan, config)
}
