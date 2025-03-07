package utilities

import (
	"fmt"
	"os"
	"path/filepath"
)

func Delete(path string, verbose bool) {

	path_abs, _ := filepath.Abs(path)

	if path_abs == "/" || path_abs == "C:\\" {
		fmt.Println("Cannot delete root directory")
	}

	if path_abs == filepath.Dir(os.Args[0]) {
		fmt.Println("Cannot delete the directory containing the executable")
	}

	if path_abs == os.Args[0] {
		fmt.Println("Cannot delete the executable")
	}

	var delete_err error
	if stat, err := os.Stat(path_abs); err == nil && stat.IsDir() == false {
		if verbose {
			fmt.Printf("Deleting file: %s\n", path_abs)
		}
		delete_err = os.Remove(path_abs) // remove a single file
		if delete_err != nil {
			fmt.Println(err)
		}
	}

	if stat, err := os.Stat(path_abs); err == nil && stat.IsDir() {
		if verbose {
			fmt.Printf("Recursively deleting directory: %s\n", path_abs)
		}
		delete_err = os.RemoveAll(path_abs) // delete an entire directory
		if delete_err != nil {
			fmt.Println(err)
		}
	}
}
