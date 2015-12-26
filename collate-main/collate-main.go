// Paul Morrison's "Collate" example on page 91 of "Flow-Based Programming, 2nd Edition"
//
//  +---+
//  | A |---+       +----------+
//  +---+   |       |          |       +--------+
//          +------>|0         |       |        |
//                  |    C  out|------>|    D   |
//          +------>|1         |       |        |
//  +---+   |       |          |       +--------+
//  | B |---+       +----------+
//  +---+
//
// Component A produces "headers".  Component B produces data keyed by headers (sorted).  Component C
// Collates the output from A and B and sends the merged output to D.  C reads a header from A, then keeps
// reading data records from B until the key changes.  Then, C reads the next header from A and repeats.
//
// The implementation below shows only the "happy" case.  Edge cases, e.g. what happens when B skips over
// a header, are not implemented below, but it should be clear how such logic could be added.
//
// This implementation explicitly creates Go channels for A to C, B to C and C to D, then passes theses
// channels as parameters to the various components, thereby "wiring up" the diagram.
//

package main

import (
	"fmt"
	"github.com/guitarvydas/fbp-example/collate"
	"github.com/guitarvydas/fbp-example/header"
	"github.com/guitarvydas/fbp-example/printer"
	"github.com/guitarvydas/fbp-example/record"
)

func main() {
	fmt.Println("Ports as explicit func arguments")
	c1 := make(chan string, 5)
	c2 := make(chan string, 5)
	c3 := make(chan string, 5)
	go collate.Collate(c1, c2, c3)
	go header.Read(c1)
	go record.Read(c2)
	printer.Print(c3)
}
