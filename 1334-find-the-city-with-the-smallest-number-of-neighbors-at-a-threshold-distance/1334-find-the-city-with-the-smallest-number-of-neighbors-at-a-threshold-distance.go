func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	var graph = make(map[int][][]int)
    
	for _, n := range edges{
		graph[n[0]] = append(graph[n[0]], []int{n[1], n[2]})
		graph[n[1]] = append(graph[n[1]], []int{n[0], n[2]})
	}
    var resMap = make(map[int]int)
    for i:=0; i<n; i++{
        resMap[i] = 0
    }
    
	
	for k, _ := range graph{
		ans := bfs(k, graph, distanceThreshold, n)
		resMap[k] = ans

	}
    min := math.MaxInt32
    max := math.MinInt32

    for k, v := range resMap{
        //one city cannot visit more than max -->(n-1)
        if v > n-1{
            v = n-1
        }
        
        if v < min{
            min = v
            max = k
        }else if v == min{
             max = Max(k, max)
        }
    }


	return max
}



type Node struct{
	val int
	dist int
}

func bfs(node int, graph map[int][][]int, distanceThreshold int,n int) int{
	var visited = make([]bool, n)
	var q []Node
	
	var ans int
	q = append(q, Node{val: node, dist: 0})

	for len(q)>0{

		currNode := q[0]
		q = q[1:]

		visited[currNode.val] = true
		for _, n := range graph[currNode.val]{

			if currNode.dist + n[1] > distanceThreshold{
				continue
			}            

			if !visited[n[0]]{
				ans++
				q = append(q, Node{val: n[0], dist: currNode.dist + n[1]})
			}
		}
	}
	return ans
}

func Max(i, j int) int{
    if i>j{
        return i
    }
    return j
}