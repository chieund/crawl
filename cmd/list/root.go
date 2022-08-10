package list

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "1.0.0"
var RootCmd = &cobra.Command{
	Use:     "command",
	Short:   "command",
	Long:    "command",
	Version: version,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error your CLI '%s'", err)
		os.Exit(1)
	}
}
