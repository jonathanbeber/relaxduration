# relaxduration

If you accept that 1 day is 24 hours in some situations, you might want
to parse it in Go too.

This package tries to handle situations where durations were used
without much care (e.g. 1 day is exactly 24 hours from now, 1 week is
168 hours from now, etc.) and you have to parse it in `Go` for some
reason.

I decided to create it after reading this [Golang Google Groups
discussion](https://groups.google.com/g/golang-nuts/c/iSuil_hY7w8) and
[this issue](https://github.com/golang/go/issues/11473). Since I
searched, couldn't find one and, it took just some fun minutes, here we
are.

# Usage

Check `example_duration_test.go`, if it's not clear or you do not care
about `flags` try something like:

```golang
package main

import (
	"github.com/jonathanbeber/relaxduration"
)

func main() {
	d, err := relaxduration.Parse("3 days 7 hours and 15 secs")
	if err != nil {
		println("unexpected: %v", err)
	}
	println(d.String())
	// Output: 79h0m15s
}
```
