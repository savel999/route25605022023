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

	workersDays := make([][]int, cnt)

	for i := 0; i < cnt; i++ {
		var workerDaysCnt int
		fmt.Fscan(in, &workerDaysCnt)

		workersDays[i] = make([]int, workerDaysCnt)

		for j := 0; j < workerDaysCnt; j++ {
			var day int
			fmt.Fscan(in, &day)

			workersDays[i][j] = day
		}
	}

	for _, workerDaysTasks := range workersDays {
		tasksMap := make(map[int]int)
		curTask := workerDaysTasks[0]

		tasksMap[curTask]++

		for i := 1; i < len(workerDaysTasks); i++ {
			if workerDaysTasks[i] != curTask {
				tasksMap[workerDaysTasks[i]]++
				curTask = workerDaysTasks[i]
			}
		}

		isOk := true
		for _, val := range tasksMap {
			if val > 1 {
				isOk = false

				break
			}
		}

		if isOk {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
