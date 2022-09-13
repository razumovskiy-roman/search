package main

import (
	"fmt"
	"math"
	"os"
	"encoding/json"
)

const L = 10

var n [L]int

func inc(a [L]int) [L]int {

	for i := L - 1; i >= 0; i-- {
		a[i]++
		if a[i] < 10 {
		break
		}
		if a[i] == 10 {
		 a[i] = 0
		}
	}

return a
}

func count(j int, k int, a[L]int) bool {
	var n int
	n=0
	if j == 10 {j=0}
	for i :=0; i<10; i++ {
		if a[i] == j {
			n++
		}
	}
	if n == k {
		return false
	}
	return true
}

func chek(a[L]int) bool {
	for i:=1; i<11; i++ {
		if count(i, a[i-1], a) {
			return false
			break
		}
	}
	return true
}

func main() {

	file, err := os.Create("result.txt")
	if err != nil{
		fmt.Println("Unable to create file:", err) 
		os.Exit(1) 
	}
	defer file.Close() 
	
	file1, err := os.Create("progress.txt")
	if err != nil{
		fmt.Println("Unable to create file:", err) 
		os.Exit(1) 
	}
	defer file1.Close() 

	for j := 0; j <= int(math.Pow(10, L)); j++ {
		if j%1000000000 == 0 {
			s, _ := json.Marshal(n)
			text := string(s)
			file1.WriteString(text + "\n")
		}
		//fmt.Println(n)
		if chek(n) {
			fmt.Println(n)
			s, _ := json.Marshal(n)
			text := string(s)
			file.WriteString(text + "\n")
		}
		n = inc(n)
	}
}
