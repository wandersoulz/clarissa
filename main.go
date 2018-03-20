package main

import (
	"clarissa/kernel"
	"os"
	"strconv"

	"github.com/wandersoulz/godes"
)

func main() {
	seed, _ := strconv.ParseInt(os.Args[1], 10, 64)
	godes.SetSeed(seed)

	// CLARISSA SPECIFIC CODE

	godes.Run()

	// CLARISSA SPECIFIC CODE
	c := kernel.Init()
	c.ReportCurrentMajor()

	godes.AddRunner(c)
	godes.WaitUntilDone() // waits for all the runners to finish the Run()

	c.ReportCurrentMajor()
}
