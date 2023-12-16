package math

func (left *Matrix) Add(right *Matrix) (out *Matrix) {
	out = new(Matrix)

	out.rows = left.rows
	out.cols = left.cols

	for x := uint64(0); x < out.rows; x++ {
		for y := uint64(0); y < out.cols; y++ {
			out.mat[x][y] = left.mat[x][y] + right.mat[x][y]
		}
	}

	return
}
