package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func mergeCsvFiles(f1, f2, outputFile string) {
	file1, err := os.Open(f1)
	if err != nil {
		panic(err)
	}
	defer file1.Close()

	file2, err := os.Open(f2)
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	file3, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer file3.Close()

	reader1 := csv.NewReader(file1)
	reader2 := csv.NewReader(file2)
	writer := csv.NewWriter(file3)
	defer writer.Flush()

	// reading headers from file1, file2
	headers1, err := reader1.Read()
	if err != nil {
		panic(err)
	}

	if err := writer.Write(headers1); err != nil {
		panic(err)
	}

	_, err = reader2.Read()
	if err != nil {
		panic(err)
	}

	rec1, err := readRecords(reader1)
	if err != nil {
		panic(err)
	}

	rec2, err := readRecords(reader2)
	if err != nil {
		panic(err)
	}

	writeRecords(writer, rec1)
	fmt.Println("file1 rows")

	writeRecords(writer, rec2)
	fmt.Println("file2 rows")

}

func readRecords(r *csv.Reader) ([][]string, error) {
	rec, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rec, nil
}

func writeRecords(w *csv.Writer, rec [][]string) {
	if err := w.WriteAll(rec); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("CSV Reader")

	mergeCsvFiles("winequality-red.csv", "winequality-white.csv", "winequality_combined.csv")

	fmt.Println("merging completed")
}
