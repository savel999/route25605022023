package main

import (
	"bufio"
	"fmt"
	"math"
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

	developersSets := make([][]int, cnt)

	for i := 0; i < cnt; i++ {
		var developersCnt int
		fmt.Fscan(in, &developersCnt)

		developersSets[i] = make([]int, developersCnt)

		for j := 0; j < developersCnt; j++ {
			var developerSkill int
			fmt.Fscan(in, &developerSkill)

			developersSets[i][j] = developerSkill
		}

	}

	for _, set := range developersSets {
		setLen := len(set)
		pairs := make([][2]int, setLen/2)

		for pairIdx, _ := range pairs {
			for i := 0; i < setLen; i++ {
				if set[i] == 0 {
					continue
				}

				aIdx := i
				aExp := set[i]
				bIdx := 0

				isPicked := false
				diff := 0

				for j := i + 1; j < setLen; j++ {
					if set[j] == 0 {
						continue
					}

					if !isPicked || diff > int(math.Abs(float64(aExp-set[j]))) {
						bIdx = j
						diff = int(math.Abs(float64(aExp - set[j])))

						isPicked = true
					}
				}

				pairs[pairIdx][0] = aIdx + 1
				pairs[pairIdx][1] = bIdx + 1

				set[aIdx] = 0
				set[bIdx] = 0

				break
			}
		}

		for _, pairItem := range pairs {
			fmt.Fprintln(out, pairItem[0], " ", pairItem[1])
		}

		fmt.Fprintln(out)
	}

	fmt.Fprintln(out)
}
