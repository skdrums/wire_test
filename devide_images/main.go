package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("input? ")
	// Scannerを使って一行読み
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println("input is", scanner.Text())
}
