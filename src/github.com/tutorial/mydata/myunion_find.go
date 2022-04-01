package mydata

type UnionFind struct {
	Count int
	Parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := range(parent) {
		parent[i] = i
	}

	return &UnionFind{
		Count: n,
		Parent: parent,
	}
}

func (u *UnionFind) Find(x int) int {
	for u.Parent[x] != x {
		x = u.Parent[x]
	}
	return x
}

func (u *UnionFind) Union(p, q int) {
	rootP := u.Find(p)
	rootQ := u.Find(q)
	if rootP == rootQ {
		return 
	}

	u.Parent[p] = rootQ
	u.Count --
}

func (u *UnionFind) Connected(p, q int) bool {
    rootP := u.Find(p) 
    rootQ := u.Find(q)
    return rootP == rootQ;
}
