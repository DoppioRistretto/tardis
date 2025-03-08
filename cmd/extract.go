/*
Copyright © 2025 Brady Semm <btsemm@protonmail.com>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"tardis/utilities"

	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Recursively extract a tar archive file",
	Long: `Recursively extract a tar archive file to a specified directory.
Supports gzip, bzip2, and uncompressed tar files.`,

	Run: func(cmd *cobra.Command, args []string) {

		file, _ := cmd.Flags().GetString("file")
		file_abs, _ := filepath.Abs(file)
		destination, _ := cmd.Flags().GetString("destination")
		dest_abs, _ := filepath.Abs(destination)
		verbose, _ := cmd.Flags().GetBool("verbose")

		if _, err := os.Stat(file_abs); os.IsNotExist(err) {
			log.Fatalf("File %s does not exist\n", file_abs)
		}

		if _, err := os.Stat(dest_abs); os.IsNotExist(err) {
			os.MkdirAll(dest_abs, os.FileMode(0755))
		}

		utilities.Untar(file_abs, dest_abs)

		e := filepath.WalkDir(dest_abs, func(path string, d os.DirEntry, err error) error {
			if verbose {
				log.Printf("Visited: %s\n", path)
			}
			if strings.HasSuffix(path, ".tar") || strings.HasSuffix(path, ".gz") || strings.HasSuffix(path, ".tgz") || strings.HasSuffix(path, ".tar.bz2") {
				if verbose {
					log.Printf("Found archive: %s\n", path)
					log.Printf("Untarring archive: %s to %s\n", path, filepath.Dir(path))
				}
				utilities.Untar(path, filepath.Dir(path))
			}
			return nil
		})
		if e != nil {
			log.Fatal(e)
		}
	},
}

func init() {

	rootCmd.AddCommand(extractCmd)

	var file string
	var destination string

	extractCmd.Flags().StringVarP(&file, "file", "f", "", "The tar file to extract")
	extractCmd.MarkFlagRequired("file")

	extractCmd.Flags().StringVarP(&destination, "destination", "d", "", "The destination directory to extract the tar file to")
	extractCmd.MarkFlagRequired("destination")
}
