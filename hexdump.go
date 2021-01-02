package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Pls specify which file to dump")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	size := len(data)

	for a := 0; a < size; a++ {
		if a%16 == 0 {
			if a > 0 {
				fmt.Printf(" %02x", data[a])
				fmt.Print("  ")
				for b := 15; b > -1; b-- {
					if data[a-b] > 31 && data[a-b] < 128 {
						fmt.Printf("%c", data[a-b])
					} else {
						fmt.Print(".")
					}
				}
			}

			fmt.Printf("%s%08x ", func(a int) string {
				if a == 0 {
					return ""
				}
				return "\n"
			}(a), a)
		} else {
			fmt.Printf(" %02x", data[a])

			if a%8 == 0 && a > 0 {
				fmt.Printf(" ")
			}
		}
	}
}
