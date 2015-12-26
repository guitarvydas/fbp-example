package printer

import "fmt"

func Print(in <-chan string) {
	for {
		merged := <-in
		if merged == "EOF" {
			break
		}
		fmt.Println(merged)
	}
}
