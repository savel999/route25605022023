package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	Min, Max int
}

type IntervalsList []Interval

func (e IntervalsList) Len() int {
	return len(e)
}

func (e IntervalsList) Less(i, j int) bool {
	return e[i].Min < e[j].Min
}

func (e IntervalsList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
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

	result := make([]bool, cnt)
	for i := 0; i < cnt; i++ {
		result[i] = true
	}

	for i := 0; i < cnt; i++ {
		var intervalsCnt int
		fmt.Fscan(in, &intervalsCnt)

		isGoodSet := true
		intervalsListStructs := make(IntervalsList, intervalsCnt)

		for j := 0; j < intervalsCnt; j++ {
			var interval string
			fmt.Fscan(in, &interval)

			if !isGoodSet {
				continue
			}

			boundaries := strings.Split(interval, "-")
			aParts := strings.Split(boundaries[0], ":")
			bParts := strings.Split(boundaries[1], ":")

			hA, _ := strconv.Atoi(aParts[0])
			mA, _ := strconv.Atoi(aParts[1])
			sA, _ := strconv.Atoi(aParts[2])

			hB, _ := strconv.Atoi(bParts[0])
			mB, _ := strconv.Atoi(bParts[1])
			sB, _ := strconv.Atoi(bParts[2])

			// validation
			if (hA < 0 || hA > 23) || (mA < 0 || mA > 59) || (sA < 0 || sA > 59) ||
				(hB < 0 || hB > 23) || (mB < 0 || mB > 59) || (sB < 0 || sB > 59) {
				isGoodSet = false

				continue
			}

			// min <= max
			totalSecondsA := hA*3600 + mA*60 + sA
			totalSecondsB := hB*3600 + mB*60 + sB

			if totalSecondsA > totalSecondsB {
				isGoodSet = false

				continue
			}

			intervalsListStructs[j] = Interval{Min: totalSecondsA, Max: totalSecondsB}
		}

		// intersection
		if isGoodSet && intervalsCnt > 1 {
			sort.Sort(IntervalsList(intervalsListStructs))

			for j := 0; j < intervalsCnt-1; j++ {
				first := intervalsListStructs[j]
				second := intervalsListStructs[j+1]

				if first.Min <= second.Max && first.Max >= second.Min {
					isGoodSet = false

					break
				}
			}
		}

		result[i] = isGoodSet
	}

	for _, isOk := range result {
		if isOk {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
