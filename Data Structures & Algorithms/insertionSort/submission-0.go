// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
	res := make([][]Pair, 0, len(pairs))

	for i := range pairs {
		j := i
		for j > 0 && pairs[j].Key < pairs[j-1].Key {
			pairs[j], pairs[j-1] = pairs[j-1], pairs[j]
			j--
		}
		c := make([]Pair, len(pairs))
		copy(c, pairs)
		res = append(res, c)

	}

	return res
}
