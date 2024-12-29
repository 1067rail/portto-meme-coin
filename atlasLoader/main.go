// This file is for atlas Go Program Mode

package main

import (
	"fmt"
	"io"
	"os"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/1067rail/portto-meme-coin/models"
)

func main() {
	stmts, err := gormschema.New("postgres").Load(
		&models.MemeCoin{},
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load gorm schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}
