package main

import (
	"bufio"
	"fmt"
	"os"
)

type Result struct {
	A, B int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	fmt.Fscan(in, &cnt)

	if cnt > 10000 {
		return
	}

	resultSlice := make([]Result, cnt)

	for i := 0; i < cnt; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)

		if a < -1000 || a > 1000 || b < -1000 || b > 1000 {
			return
		}

		resultSlice[i] = Result{A: a, B: b}
	}

	for _, resItem := range resultSlice {
		fmt.Fprintln(out, resItem.A+resItem.B)
	}
}
