func mergeSort(pairs []Pair) []Pair {
	return ms(pairs, 0, len(pairs)-1)
}

func ms(p []Pair, l, r int) []Pair {
	if r-l+1 <= 1 {
		return p
	}

	m := (l+r)/2

	ms(p, l, m)
	ms(p, m+1, r)
	merge(p, l, m, r)

	return p
}

func merge(p []Pair, l, m, r int) {
	L := make([]Pair, m-l+1)
	copy(L, p[l:m+1])
	R := make([]Pair, r-m)
	copy(R, p[m+1:r+1])

	i, j, k := 0, 0, l

	for i < len(L) && j < len(R) {
		if L[i].Key <= R[j].Key {
			p[k] = L[i]
			i++
		} else {
			p[k] = R[j]
			j++
		}
		k++
	}

	for i < len(L) {
		p[k] = L[i]
		i++
		k++
	}

	for j < len(R) {
		p[k] = R[j]
		j++
		k++
	}
}