package main

import (
	"github.com/Gallrecon-Tech/email-pwnd-checker/internal/check_runner"
	"log"
	"os"
)

func parseArgs() string {
	if len(os.Args) < 2 {
		log.Println("Missing email address to scan")
		log.Printf("Usage: %s <email>", os.Args[0])
		os.Exit(1)
	}
	return os.Args[1]
}

func main() {
	log.Println("Welcome in Email Pwnd Checker by GALLRECON.TECH")
	emailToScan := parseArgs()
	runner := check_runner.Runner{}
	runner.CheckEmail(emailToScan)
}
