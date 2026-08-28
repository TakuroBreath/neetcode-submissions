func getConcatenation(nums []int) []int {
    n := len(nums)
    res := make([]int, 2*n)

    copy(res, nums)
    copy(res[n:], nums)

    return res
}