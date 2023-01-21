package cli_test

import (
	"testing"

	"github.com/rommms07/idream-erp/internal/cli"
)

func Test_shouldStartTheCliApp(t *testing.T) {
	cli.Start([]string{})
}
