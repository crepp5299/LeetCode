func calPoints(operations []string) int {
    rs := make([]int, 0)
    for _, c := range operations {
        switch x := string(c); x {
            case "+":
            l := len(rs)
            rs = append(rs, rs[l - 1] + rs[l-2])
            case "D":
                        rs = append(rs, 2*int(rs[len(rs) - 1]))
            case "C":
                rs = rs[:len(rs)-1]
            default:
            xInt, _ :=strconv.Atoi(x)
            rs = append(rs, xInt)
        }
    }
    sum := 0
    for i := range rs {
        sum+= rs[i]
    }
    return sum
}