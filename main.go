package main

import (
	"clarissa/kernel"
	"os"
	"strconv"

	"github.com/wandersoulz/godes"
)

func main() {
	numSeeds, _ := strconv.ParseInt(os.Args[1], 10, 64)
	momFriendship, _ := strconv.ParseFloat(os.Args[2], 64)

	seed := numSeeds
	godes.SetSeed(seed)
	// CLARISSA SPECIFIC CODE

	godes.Run()

	// CLARISSA SPECIFIC CODE
	k := kernel.Init(seed, momFriendship)
	k.ReportCurrentMajor()

	godes.AddRunner(kernel.InitDay(k))
	godes.WaitUntilDone() // waits for all the runners to finish the Run()

	k.ReportCurrentMajor()
	godes.Clear()
}
