package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type TableWithClicks struct {
	Rows   int
	Cols   int
	Clicks []int
	Table  [][]int
}

type Sortable struct {
	Nums, Idxs []int
}

func (s Sortable) Len() int           { return len(s.Nums) }
func (s Sortable) Less(i, j int) bool { return s.Nums[i] < s.Nums[j] }
func (s Sortable) Swap(i, j int) {
	s.Nums[i], s.Nums[j] = s.Nums[j], s.Nums[i]
	s.Idxs[i], s.Idxs[j] = s.Idxs[j], s.Idxs[i]
}

func SortAndReturnIdxs(nums []int) []int {
	idxs := make([]int, len(nums))
	for i := range idxs {
		idxs[i] = i
	}

	sort.Sort(Sortable{nums, idxs})

	return idxs
}

func CopyTable(table [][]int) [][]int {
	duplicate := make([][]int, len(table))
	for i := range table {
		duplicate[i] = make([]int, len(table[i]))
		copy(duplicate[i], table[i])
	}

	return duplicate
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var setsCnt int
	fmt.Fscan(in, &setsCnt)

	tablesWithClicks := make([]TableWithClicks, setsCnt)

	for i := 0; i < setsCnt; i++ {
		var rowsCnt, colsCnt int
		fmt.Fscan(in, &rowsCnt, &colsCnt)

		// table
		tablesWithClicks[i].Table = make([][]int, rowsCnt)
		tablesWithClicks[i].Rows = rowsCnt
		tablesWithClicks[i].Cols = colsCnt

		for j := 0; j < rowsCnt; j++ {
			tablesWithClicks[i].Table[j] = make([]int, colsCnt)

			for k := 0; k < colsCnt; k++ {
				var cell int
				fmt.Fscan(in, &cell)

				tablesWithClicks[i].Table[j][k] = cell
			}
		}

		// clicks
		var clicksCnt int
		fmt.Fscan(in, &clicksCnt)

		tablesWithClicks[i].Clicks = make([]int, clicksCnt)

		for cl := 0; cl < clicksCnt; cl++ {
			var clickCol int
			fmt.Fscan(in, &clickCol)

			tablesWithClicks[i].Clicks[cl] = clickCol
		}

	}

	// sorting & print
	for _, table := range tablesWithClicks {
		// sorting
		for _, clickCol := range table.Clicks {
			clickColValues := make([]int, table.Rows)

			for i := 0; i < table.Rows; i++ {
				clickColValues[i] = table.Table[i][clickCol-1]
			}

			sortedIndexes := SortAndReturnIdxs(clickColValues)
			resultTable := CopyTable(table.Table)

			for idx, rowIdx := range sortedIndexes {
				copy(resultTable[idx], table.Table[rowIdx])
			}

			for i := 0; i < table.Rows; i++ {
				for j := 0; j < table.Cols; j++ {
					table.Table[i][j] = resultTable[i][j]
				}
			}
		}

		// print
		for i := 0; i < table.Rows; i++ {
			for j := 0; j < table.Cols; j++ {
				if j == table.Cols-1 {
					fmt.Fprintf(out, "%d\n", table.Table[i][j])
				} else {
					fmt.Fprintf(out, "%d ", table.Table[i][j])
				}

			}
		}

		fmt.Fprintln(out)
	}

}
