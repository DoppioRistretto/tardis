/*
Copyright © 2025 Brady Semm <btsemm@protonmail.com>
*/
package cmd

import (
	"tardis/utilities"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a tar archive file",
	Long: `Create a tar archive file from the specified paths.
Supports gzip compression.`,

	Run: func(cmd *cobra.Command, args []string) {

		archiveName, _ := cmd.Flags().GetString("archive")
		paths, _ := cmd.Flags().GetStringArray("path")
		delete, _ := cmd.Flags().GetBool("delete")
		verbose, _ := cmd.Flags().GetBool("verbose")

		err := utilities.Tar(archiveName, paths)
		if err == nil && delete {
			// delete the source files
			for _, path := range paths {
				utilities.Delete(path, verbose)
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	var archiveName string
	var paths []string
	var delete bool

	createCmd.Flags().StringVarP(&archiveName, "archive", "a", "", "The tar file to create. Use .gz or .tgz extension for gzip compression")
	createCmd.MarkFlagRequired("archive")

	createCmd.Flags().StringArrayVarP(&paths, "path", "p", []string{}, "The paths to include in the tar file")
	createCmd.MarkFlagRequired("path")

	createCmd.Flags().BoolVar(&delete, "delete", false, "Delete the source files after creating the tar file")
}
