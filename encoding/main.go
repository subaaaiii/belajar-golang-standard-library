package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	// base 64
	value := "Subairi belajar golang"
	encoded := base64.StdEncoding.EncodeToString([]byte(value))

	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}

	//csv
	// csv reader
	csvString := "eko, kurniawan, khannedy\n" +
		"budy, pratama, nugraha\n" +
		"joko, morro, diah"
	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	// csv writer

	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"eko", "kurniawan", "khannedy"})
	_ = writer.Write([]string{"budi", "nugraha", "pratama"})
	_ = writer.Write([]string{"joko", "muro", "jaah"})

	writer.Flush()

}
