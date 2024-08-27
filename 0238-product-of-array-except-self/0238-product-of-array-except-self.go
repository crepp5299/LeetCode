func productExceptSelf(nums []int) []int {
    i, j := 1, 1
    n := len(nums)
    rs := make([]int, n)
    for k:=0;k<n;k++{
        if rs[k] == 0 && k < n/2{
            rs[k] = i
            rs[n-1-k] = j
        } else {
            rs[k]*=i
            rs[n-1-k]*=j
        }
        if k == n-1-k {
            rs[k] = i*j
        }
        i = i*nums[k]
        j = j*nums[n-1-k]
    }
    return rs
}