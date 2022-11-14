package day5

//https://www.bilibili.com/video/BV1KP4y187qj/?spm_id_from=333.999.0.0&vd_source=1761deea5b8c69ccd1cd089a400b2c54

func findNumberIn2DArray(matrix [][]int, target int) bool {

	//z字形查找
	i := len(matrix) - 1
	j := 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
