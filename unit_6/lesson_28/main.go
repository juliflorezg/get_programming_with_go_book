// errors

// lesson goals:
// Write files and handle failure
// Handle errors with a flair of creativity
// Make and identify specific errors
// Keep calm and carry on

package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error
}

func (sf *safeWriter) writeLn(str string) {
	if sf.err != nil {
		return
	}

	_, sf.err = fmt.Fprintln(sf.w, str)
	return
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	sw := safeWriter{w: f}

	sw.writeLn("Errors are values")
	sw.writeLn("Don’t just check errors, handle them gracefully.")
	sw.writeLn("Don't panic.")
	sw.writeLn("Make the zero value useful.")
	sw.writeLn("The bigger the interface, the weaker the abstraction.")
	sw.writeLn("interface{} says nothing.")
	sw.writeLn("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.writeLn("Documentation is for users.")
	sw.writeLn("A little copying is better than a little dependency.")
	sw.writeLn("Clear is better than clever.")
	sw.writeLn("Concurrency is not parallelism.")
	sw.writeLn("Don’t communicate by sharing memory, share memory by communicating.")
	sw.writeLn("Channels orchestrate; mutexes serialize.")

	return sw.err

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

	////////////////////////////////////? sudoku rules
	{

		var g grid

		// err := g.set(10, 0, 5)

		// if err != nil {
		// 	fmt.Printf("An error occurred: %v.\n", err)
		// 	os.Exit(1)
		// }

		err = g.set(0, 0, 9)

		if err != nil {

			// switch err {
			// case ErrBounds, ErrDigit:
			// 	fmt.Println("Les erreurs de paramètres hors limites.") // error de parámetro fuera de rango
			// default:
			// 	fmt.Println(err)
			// }
			// os.Exit(1)

			//? use of type assertion to get the actual type (underlying type) of value err and getting its inner errors stored in []error
			if errs, ok := err.(SudokuError); ok {
				fmt.Println(errs)
				fmt.Printf("%d error(s) occurred: \n", len(errs))
				for _, e := range errs {
					fmt.Printf("- %v\n", e)
				}
			}
			os.Exit(1)
		}
	}

	parseUrl()
}
