package collate

import "fmt"

func Collate(in0 <-chan string, in1 <-chan string, out chan<- string) {
	var hdr, rec string
	for {
		switch {
		case rec == "EOF":
			out <- rec
			break
		case hdr == "":
			hdr = <-in0
			out <- fmt.Sprintf("header %s", hdr)
		case rec != "" && rec[0] == hdr[0]:
			out <- fmt.Sprintf("record /%s/", rec[1:])
			rec = <-in1
		case rec != "" && rec[0] != hdr[0]:
			hdr = ""
		case rec == "":
			rec = <-in1
		}
	}
}
