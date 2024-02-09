package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	filePath := "urls.txt"

	

	file, err := os.Open(*filePath)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	fmt.Printf("%-100s%-40s\n", "URL", "Result")
	fmt.Println("----------------------------------------")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		checkOpenRedirect(url)
	}
}

func checkOpenRedirect(url string) {

	para_path := "parameters.txt"

	file, err := os.Open(para_path)
	if err != nil {
		fmt.Println("Error opening the parameters file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		para_value := scanner.Text()

		full := url + para_value

		resp, err := http.Get(full)
		if err != nil {
			fmt.Printf("%-100s%-40s\n", full, fmt.Sprintf("Error: %v", err))
			return
		}
		defer resp.Body.Close()

		status := "No open redirect found"
		if resp.StatusCode >= 300 && resp.StatusCode < 400 {
			status = " open redirect maybe "


  fmt.Printf("%-100s%-40s%-10d\n", full, status, resp.StatusCode)
		}

		

	}
}
