func countComponents(n int, edges [][]int) int {
	graph := make([][]int,n)
	for _,edge := range edges {
		u := edge[0]
		v:= edge[1]
		graph[u] = append(graph[u],v)
		graph[v] = append(graph[v],u)
	}
	count := 0
	visited := make([]bool,n)
	var dfs func(int)
	dfs = func(node int){
		visited[node] = true
		for _,neighbor := range graph[node]{
			if !visited[neighbor]{
				dfs(neighbor)
			}
		}
	}
	for i := 0 ; i < n ; i++ {
		if !visited[i]{
			count +=1
		}
		dfs(i)
	}
	return count
}