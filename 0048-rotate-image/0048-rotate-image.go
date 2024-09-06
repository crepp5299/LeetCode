func rotate(matrix [][]int)  {
    n := len(matrix)
    for curI := 0; curI<=n/2; curI++ {
        for curJ := curI; curJ < n-1-curI; curJ++{
            tempVal := matrix[curI][curJ]
            for count:=0; count < 4;count++ {
                curI, curJ = curJ, n-1-curI
                tempVal, matrix[curI][curJ] = matrix[curI][curJ], tempVal
            }
        }
    }
}