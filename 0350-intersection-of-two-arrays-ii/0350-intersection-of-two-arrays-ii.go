func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    res:= []int{}
    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] == nums2[j] {
            res = append(res, nums1[i])
            i++
            j++
            continue
        }
        if nums1[i] > nums2[j] {
            j++
            continue
        }
        i++
    }
    return res
}