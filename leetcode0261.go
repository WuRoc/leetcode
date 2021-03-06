/*
LeetCode 261: https://leetcode.com/problems/graph-valid-tree/
*/

package leetcode

func validTree(n int, edges [][]int) bool {
	if n != len(edges)+1 {
		return false
	}

	parents := make([]int, n)
	for i := 0; i < len(parents); i++ {
		parents[i] = i
	}

	for _, edge := range edges {
		parent1 := getParent261(parents, edge[0])
		parent2 := getParent261(parents, edge[1])
		if parent1 == parent2 {
			return false
		}

		union261(parents, parent1, parent2)
	}

	return true
}

func getParent261(parents []int, node int) int {
	parent := parents[node]
	if parent != node {
		parent = getParent261(parents, parent)
		parents[node] = parent
	}

	return parent
}

func union261(parents []int, node1, node2 int) {
	parents[node1] = node2
}
