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

	ordersSlice := make([]map[int]int, cnt)

	for i := 0; i < cnt; i++ {
		var productsCnt int
		fmt.Fscan(in, &productsCnt)

		ordersSlice[i] = make(map[int]int, productsCnt)

		for j := 0; j < productsCnt; j++ {
			var productPrice int
			fmt.Fscan(in, &productPrice)

			ordersSlice[i][productPrice]++
		}

	}

	for _, order := range ordersSlice {
		sum := 0

		for productPrice, cntProducts := range order {
			sum += 2*productPrice*(cntProducts/3) + productPrice*(cntProducts%3)
		}

		fmt.Fprintln(out, sum)
	}
}
