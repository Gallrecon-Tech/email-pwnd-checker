package check_runner

import (
	"log"
	"net/mail"
)

type CheckResult struct {
	Address string
	Pwnd    bool
}

type Runner struct {
}

func (runner *Runner) CheckEmail(email string) (CheckResult, error) {
	log.Printf("Checking email %s", email)
	emailAddress, err := mail.ParseAddress(email)
	if err != nil {
		log.Printf("[ERROR] Wrong email format: %s", err)
		return CheckResult{}, err
	}
	// TODO: Perform check
	return CheckResult{emailAddress.Address, false}, nil
}
