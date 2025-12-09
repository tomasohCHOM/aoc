package main

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := range n {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) Find(i int) int {
	if uf.parent[i] == i {
		return i
	}
	uf.parent[i] = uf.Find(uf.parent[i])
	return uf.parent[i]
}

func (uf *UnionFind) Union(i, j int) bool {
	ri, rj := uf.Find(i), uf.Find(j)
	if ri == rj {
		return false
	}
	if uf.size[ri] < uf.size[rj] {
		uf.parent[ri] = rj
		uf.size[rj] += uf.size[ri]
	} else {
		uf.parent[rj] = ri
		uf.size[ri] += uf.size[rj]
	}
	return true
}

func (uf *UnionFind) GetSetSize(i int) int {
	root := uf.Find(i)
	return uf.size[root]
}
