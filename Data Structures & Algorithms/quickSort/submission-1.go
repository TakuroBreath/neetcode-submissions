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
	pivotidx := l

	for i := l; i < r; i++ {
		if p[i].Key < pivot.Key {
			p[i], p[pivotidx] = p[pivotidx], p[i]
			pivotidx++
		}
	}

	p[pivotidx], p[r] = p[r], p[pivotidx]

	quick(p, l, pivotidx-1)
	quick(p, pivotidx+1, r)
}
