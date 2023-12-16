package math

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (input *Matrix) ParseFromFile(filename string) {
	file, _ := os.Open(filename)
	// if err != nil {
	// 	panic(fmt.Sprintf("There was an error \"%s\"", err))
	// }
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	first_line := scanner.Text()
	tmp := strings.Split(first_line, " ")

	//TODO: ensure that the first line contains only 2 integers > 0
	rows_, columns_ := tmp[0], tmp[1]

	// omg this does not look good
	rows, _ := strconv.Atoi(rows_)
	// if err != nil {
	// 	panic(fmt.Sprintf("There was an error \"%s\"", err))
	// }

	columns, _ := strconv.Atoi(columns_)
	// if err != nil {
	// 	panic(fmt.Sprintf("There was an error \"%s\"", err))
	// }

	input.rows = uint64(rows)
	input.cols = uint64(columns)

	fmt.Println(rows, columns)

	for i := 0; i < rows; i++ {
		scanner.Scan()
		row := strings.Split(scanner.Text(), " ")
		fmt.Println(row)
		for j := 0; j < columns; j++ {
			el, _ := strconv.Atoi(row[j])
			input.mat = append(input.mat, make([]uint64, columns))
			input.mat[i][j] = uint64(el)
		}
	}
}

func (matrix *Matrix) Print(label string) {
	fmt.Printf("-=-=-=-=-= %s =-=-=-=-=-\n", label)
	fmt.Printf(
		"rows: %d\ncolumns: %d\nfield: %d\n",
		matrix.rows,
		matrix.cols,
		matrix.field,
	)
	for i := uint64(0); i < matrix.rows; i++ {
		for j := uint64(0); j < matrix.cols; j++ {
			fmt.Printf("%d ", matrix.mat[i][j])
		}
		fmt.Println()
	}
}
