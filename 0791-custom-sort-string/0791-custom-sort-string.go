func customSortString(order string, s string) string {
    m := make(map[byte]int)
    for i := range order {
        m[order[i]] = 1
    }
    buff := bytes.NewBufferString("")
    for i := range s {
        if m[s[i]] > 0{
            m[s[i]]++
        } else {
            buff.WriteByte(s[i])
        }
    }  

    for i := range order {
        for m[order[i]] > 1{
            buff.WriteByte(order[i])
            m[order[i]]--
        }
    }
  
    return buff.String()
}