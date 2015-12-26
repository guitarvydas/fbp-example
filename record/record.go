package record

import "fmt"

func Read(out chan<- string) {
	// fake reading records - first character is the key (header)
	for i := 0; i < 5; i++ {
		out <- fmt.Sprintf("A ++%v++", i)
	}
	for i := 15; i < 21; i++ {
		out <- fmt.Sprintf("B ++%v++", i)
	}
	for i := 25; i < 32; i++ {
		out <- fmt.Sprintf("C ++%v++", i)
	}
	out <- "EOF"
}
