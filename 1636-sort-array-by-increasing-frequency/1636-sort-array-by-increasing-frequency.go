func frequencySort(nums []int) []int {
    m := map[int]int{}
    for i := range nums{
        m[nums[i]]++
    }
    keys := make([]int, 0, len(m))
    for key := range m {
        keys = append(keys, key)
    }
    
    sort.SliceStable(keys, func(i, j int) bool{
        if m[keys[i]] == m[keys[j]] {
            return keys[i] > keys[j]
        }
        return m[keys[i]] < m[keys[j]]
    })
    rs := make([]int, 0, 101)
    for _, k := range keys{
        for m[k] > 0 {
            rs = append(rs, k)
            m[k]--
        }
    }
    return rs
}

// func frequencySort(nums []int) []int  {
//     basket := map[string]int{"orange": 5, "strawberryi": 7, "apple": 7, "strawberryy": 5,
//                              "mango": 3, "strawberry": 9}

//     keys := make([]string, 0, len(basket))

//     for key := range basket {
//         keys = append(keys, key)
//     }
//     sort.SliceStable(keys, func(i, j int) bool{
//         if basket[keys[i]] == basket[keys[j]] {
//             return len(keys[i]) < len(keys[j])
//         }
//         return basket[keys[i]] > basket[keys[j]]
//     })

//     for _, k := range keys{
//         fmt.Println(k, basket[k])
//     }
//     return []int{}
// }