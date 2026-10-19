func searchMatrix(matrix [][]int, target int) bool {
	return binMatrix(matrix, target, 0, len(matrix)-1)
}

func binMatrix(mat [][]int, t, l, r int) bool {
	if l > r {
		return false
	}

	m := (l+r)/2
	inMat := t >= mat[m][0] && t <= mat[m][len(mat[m])-1]
	if inMat {
		return bin(mat[m], t, 0, len(mat[m])-1)
	} else if t < mat[m][0] {
		return binMatrix(mat, t, l, m-1)
	} else {
		return binMatrix(mat, t, m+1, r)
	}
}

func bin(mm []int, t, l, r int) bool {
	if l > r {
		return false
	}

	m := (l+r)/2

	if mm[m] == t {
		return true
	} else if mm[m] < t {
		return bin(mm, t, m+1, r)
	} else {
		return bin(mm, t, l, m-1)
	}
}