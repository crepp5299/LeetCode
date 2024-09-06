func rotate(matrix [][]int)  {
    n := len(matrix)
    curI := 0
    count := 0
    for curI<=n/2 {
        for curJ := curI; curJ < n-1-curI; curJ++{
            tempI, tempJ := curI, curJ
            tempVal := matrix[tempI][tempJ]
            for count < 4 {
                tempI, tempJ = tempJ, n-1-tempI
                tempVal, matrix[tempI][tempJ] = matrix[tempI][tempJ], tempVal
                count++
            }
            count = 0
        }
        curI++
    }
}