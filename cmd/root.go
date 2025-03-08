/*
Copyright © 2025 Brady Semm <btsemm@protonmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tardis",
	Short: "Tar Dis and Untar Dat",
	Long: `A quick and dirty command line utility to archive and un-archive files.
Supports gzip and bzip2 compression.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	var verbose bool
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Print verbose output")

}
