package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var salaries []float64
	for scanner.Scan() {
		line := scanner.Text()
		raw := strings.Split(line, ":")
		if len(raw) < 2 {
			continue
		}
		salaried := strings.TrimSpace(raw[1])
		sal, err := strconv.ParseFloat(salaried, 64)
		if err != nil {
			continue
		}
		salaries = append(salaries, sal)
		if len(salaries) == 0 {
			fmt.Println("Invalid salaries entry")
			return
		}
	}
	fmt.Println(mean(salaries))
	fmt.Println(median(salaries))
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
