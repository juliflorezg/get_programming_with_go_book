// errors

// lesson goals:
// Write files and handle failure
// Handle errors with a flair of creativity
// Make and identify specific errors
// Keep calm and carry on

package main

import (
	"fmt"
	"os"
)

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
  defer f.Close()

	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(f, "Don’t just check errors, handle them gracefully.")
	return err

}

func main() {
	files, err := os.ReadDir("../lesson_27")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		fmt.Println(file.Name(), "is Dir??", file.IsDir())
	}

	err2 := proverbs("proverbs.txt")

  if err != nil {
    fmt.Println(err2)
    os.Exit(1)
  }
}
