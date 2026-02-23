package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		raw := strings.Split(line, ":")
		strings.TrimSpace(raw)
		

	}
}

func median(n []float64) float64 {
	for i := 0; i < len(n)-1; i++ {
		for j := 0; j < len(n)-1-i; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
	if len(n)%2 == 0 {
		return (n[len(n)/2] + n[len(n)/2-1]) / 2
	}
	return n[len(n)/2]
}

func mean(n []float64) float64 {
	sum := 0.00
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum / float64(len(n))
}
