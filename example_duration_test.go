package relaxduration

import (
	"flag"
	"fmt"
	"time"
)

var dst time.Duration

func ExampleRelaxedDuration_String() {
	fs := flag.NewFlagSet("RelaxedDurationExample", flag.ExitOnError)
	fs.Var(&RelaxedDuration{&dst}, "duration", "A relaxed way to set a duration. e.g: '3 days and 6 hours'")

	fs.Parse([]string{"-duration", "3 days and and 7 hours"})
	fmt.Printf(`Duration is %.0f hours`, dst.Hours())
	// Output: Duration is 79 hours
}
