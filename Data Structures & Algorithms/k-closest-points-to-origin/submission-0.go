func kClosest(points [][]int, k int) [][]int {
	quick(points, 0, len(points)-1)
	return points[:k]
}

func quick(p [][]int, l, r int) {
	if r-l+1 <= 1 {
		return
	}

	pivot := p[r]
	pivotidx := l

	for i := l; i < r; i++ {
		if dist(p[i]) < dist(pivot) {
			p[pivotidx], p[i] = p[i], p[pivotidx]
			pivotidx++
		}
	}

	p[r], p[pivotidx] = p[pivotidx], p[r]

	quick(p, l, pivotidx-1)
	quick(p, pivotidx+1, r)
}

func dist(p []int) float64 {
	if len(p) != 2 {
		return -1
	}

	x, y := p[0], p[1]

	return math.Sqrt(float64(x*x) + float64(y*y))
}