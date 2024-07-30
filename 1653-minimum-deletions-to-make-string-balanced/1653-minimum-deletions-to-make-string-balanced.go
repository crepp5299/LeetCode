// func minimumDeletions(s string) int {
//     runes := []rune(s)
//     n := len(runes)
//     countA := make([]int, n)
//     count := 0
//     for i := n - 1; i >= 0; i-- {
//         countA[i] = count
//         if runes[i] == 'a' {
//             count++
//         }
//     }
    
//     count = 0
    
//     for i, _ := range runes {
//         if countA[i] + count < n {
//             n = countA[i] + count
//         }
//         if runes[i] == 'b' {
//             count++
//         }
//     }
//     return n
// }

func minimumDeletions(s string) int {
    // from current index, calculate how many 'b' to delete on left and 'a' to delete on right.
    //in one pass get the acounts and bcounts
    
    acount, bcount := 0,0
    acount = strings.Count(s,"a")
    min:=acount
    for _, v := range s {
        if v == 'b' {
            bcount++
        } else {
            acount--
        }
        min = int(math.Min(float64(min), float64(acount+bcount)))
    }
    return min
}