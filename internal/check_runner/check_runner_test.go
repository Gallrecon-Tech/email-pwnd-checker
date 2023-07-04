package check_runner_test

import (
	"github.com/Gallrecon-Tech/email-pwnd-checker/internal/check_runner"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckEmail(t *testing.T) {
	runner := check_runner.Runner{}
	t.Run("Should return error when wrong format email is given", func(t *testing.T) {
		_, err := runner.CheckEmail("wrong_email_format")
		assert.Error(t, err)
	})
	t.Run("Should return CheckResult with email address inside for proper email", func(t *testing.T) {
		emailAddress := "email@example.com"
		result, err := runner.CheckEmail(emailAddress)
		assert.Nil(t, err)
		assert.Equal(t, emailAddress, result.Address)
	})
}
