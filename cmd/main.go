package main

import (
	"os"

	_ "github.com/rommms07/idream-erp/core/models"
	"github.com/rommms07/idream-erp/internal/cli"
)

func main() {
	cli.Start(os.Args)
}
