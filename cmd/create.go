/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"tardis/utilities"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
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

	// Here you will define your flags and configuration settings.
	var archiveName string
	var paths []string
	var delete bool

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createCmd.Flags().StringVarP(&archiveName, "archive", "a", "", "The tar file to create. Use .gz extension for gzip compression")
	createCmd.MarkFlagRequired("archive")

	createCmd.Flags().StringArrayVarP(&paths, "path", "p", []string{}, "The paths to include in the tar file")
	createCmd.MarkFlagRequired("path")

	createCmd.Flags().BoolVar(&delete, "delete", false, "Delete the source files after creating the tar file")
}
