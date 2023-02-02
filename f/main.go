package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cnt int
	fmt.Fscan(in, &cnt)

	if cnt > 10000 {
		return
	}

	result := make([]bool, cnt)
	for i := 0; i < cnt; i++ {
		result[i] = true
	}

	for i := 0; i < cnt; i++ {
		var intervalsCnt int
		fmt.Fscan(in, &intervalsCnt)

		for j := 0; j < intervalsCnt; j++ {
			var interval int
			fmt.Fscan(in, &interval)

			if isOk, _ := result[i]; isOk {

			}
		}
	}

	for _, isOk := range result {
		if isOk {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
