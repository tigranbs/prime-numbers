package main

import (
	"archive/zip"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	file_index = 1
)

func main() {
	for file_index <= 50 {
		filename := fmt.Sprintf("primes%d.zip", file_index)
		txt_file := fmt.Sprintf("primes%d.txt", file_index)

		fmt.Println("Downloading file -> ", filename)
		resp, err := http.Get(fmt.Sprintf("http://primes.utm.edu/lists/small/millions/primes%d.zip", file_index))
		if err != nil {
			fmt.Println("File Index -> ", file_index)
			panic(err)
		}

		f, err := os.Create(filename)
		if err != nil {
			fmt.Println("File Index -> ", file_index)
			panic(err)
		}

		io.Copy(f, resp.Body)
		f.Close()
		resp.Body.Close()

		fmt.Println("Unziping file -> ", filename, " to ", txt_file)

		unzip(filename, "./")

		fmt.Println("Reading Primes Text -> ", txt_file)

		text_file_data, err := ioutil.ReadFile(txt_file)
		if err != nil {
			panic(err)
		}

		text := string(text_file_data)
		number_data := make([]byte, 0)

		lines := strings.Split(text, "\n")
		for _, line := range lines {
			words := strings.Split(line, " ")
			for _, w := range words {
				num, err := strconv.Atoi(w)
				if err == nil {
					num_bytes := make([]byte, 4)
					binary.BigEndian.PutUint32(num_bytes, uint32(num))
					number_data = append(number_data, num_bytes...)
				}
			}
		}

		fmt.Println("Removing downloaded file -> ", filename)
		os.Remove(filename)
		os.Remove(txt_file)

		fmt.Println("Appending data -> ", len(number_data))
		prime_file, err := os.OpenFile("primes.data", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("File Index -> ", file_index)
			panic(err)
		}

		prime_file.Write(number_data)
		prime_file.Close()

		file_index++
	}

	fmt.Println("Done !!!")
}

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
