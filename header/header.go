package header

// import "fmt"

func Read(out chan<- string) {
	// fake reading of headers from a file, just generate some headers
	out <- "A"
	out <- "B"
	out <- "C"
}
