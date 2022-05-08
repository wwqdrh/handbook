package io

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Movie will hold our parsed CSV
type Movie struct {
	Title    string
	Director string
	Year     int
}

// ReadCSV gives shows some examples of processing CSV
// that is passed in as an io.Reader
func ReadCSV(b io.Reader) ([]Movie, error) {

	r := csv.NewReader(b)

	// These are some optional configuration options
	r.Comma = ';'
	r.Comment = '-'

	var movies []Movie

	// grab and ignore the header for now
	// we may also wanna use this for a dictionary key or some
	// other form of lookup
	_, err := r.Read()
	if err != nil && err != io.EOF {
		return nil, err
	}

	// loop until it's all processed
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		year, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return nil, err
		}

		m := Movie{record[0], record[1], int(year)}
		movies = append(movies, m)
	}
	return movies, nil
}

// AddMoviesFromText uses the CSV parser with a string
func AddMoviesFromText() error {
	// this is an example of us taking a string, converting
	// it into a buffer, and reading it with the csv package
	in := `
- first our headers
movie title;director;year released

- then some data
Guardians of the Galaxy Vol. 2;James Gunn;2017
Star Wars: Episode VIII;Rian Johnson;2017
`

	b := bytes.NewBufferString(in)
	m, err := ReadCSV(b)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", m)
	return nil
}

// A Book has an Author and Title
type Book struct {
	Author string
	Title  string
}

// Books is a named type for an array of books
type Books []Book

// ToCSV takes a set of Books and writes to an io.Writer
// it returns any errors
func (books *Books) ToCSV(w io.Writer) error {
	n := csv.NewWriter(w)
	err := n.Write([]string{"Author", "Title"})
	if err != nil {
		return err
	}
	for _, book := range *books {
		err := n.Write([]string{book.Author, book.Title})
		if err != nil {
			return err
		}
	}

	n.Flush()
	return n.Error()
}

// WriteCSVOutput initializes a set of books
// and writes the to os.Stdout
func WriteCSVOutput() error {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}

	return b.ToCSV(os.Stdout)
}

// WriteCSVBuffer returns a buffer csv for
// a set of books
func WriteCSVBuffer() (*bytes.Buffer, error) {
	b := Books{
		Book{
			Author: "F Scott Fitzgerald",
			Title:  "The Great Gatsby",
		},
		Book{
			Author: "J D Salinger",
			Title:  "The Catcher in the Rye",
		},
	}

	w := &bytes.Buffer{}
	err := b.ToCSV(w)
	return w, err
}
