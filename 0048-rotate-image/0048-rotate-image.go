func rotate(matrix [][]int)  {
    n := len(matrix)
    for curI := 0; curI<=n/2; curI++ {
        for curJ := curI; curJ < n-1-curI; curJ++{
            tempI, tempJ := curI, curJ
            tempVal := matrix[tempI][tempJ]
            for count:=0; count < 4;count++ {
                tempI, tempJ = tempJ, n-1-tempI
                tempVal, matrix[tempI][tempJ] = matrix[tempI][tempJ], tempVal
            }
        }
    }
}