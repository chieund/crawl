package main

import (
	"crawl/cmd/cronjob"
	"crawl/cmd/migrate"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var version = "1.0.0"
var rootCmd = &cobra.Command{
	Use:     "command",
	Short:   "command",
	Long:    "command",
	Version: version,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "There was an error your CLI '%s'", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(cronjob.CrawlArticleCmd)
	rootCmd.AddCommand(cronjob.CrawlArticleDetailCmd)
	rootCmd.AddCommand(migrate.MigrateCmd)
}

func main() {
	Execute()
}
