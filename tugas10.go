package main

import (
	"fmt"
	"time"
)

func main() {
	var nama = "Hugo Ghally Imanaka"
	var lahir = "20-01-2001"

	fmt.Println(nama)
	time.Sleep(time.Second * 5)
	fmt.Println("sudah 5 detik")
	fmt.Println(lahir)
}
