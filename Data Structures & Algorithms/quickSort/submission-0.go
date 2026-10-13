// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func QuickSort(pairs []Pair) []Pair {
	quick(pairs, 0, len(pairs)-1)
	return pairs
}

func quick(p []Pair, l, r int) {
	if r-l+1 <= 1 {
		return
	}

	pivot := p[r]
	pivotIdx := l
	
	for i := l; i < r; i++ {
		if p[i].Key < pivot.Key {
			p[pivotIdx], p[i] = p[i], p[pivotIdx]
			pivotIdx++
		}
	}
	p[pivotIdx], p[r] = p[r], p[pivotIdx]

	quick(p, l, pivotIdx-1)
	quick(p, pivotIdx+1, r)
}
