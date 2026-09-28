package graphs

/*
Disjoint Set
- This takes input a graph, after constructing it can tell whether some nodes belong to same connected component or not, basically they belong to same set or not

Explanation:

Method-1: Union By Rank
- This has two things which we maintain, ParentArr which stores parent and RankArr which stores rank this can also be said that a tree of height 2 root has a rank 2
- If ranks are same then we can consider any of the edge as root and increment its rank, if not we can connect to other parent
- To check whether 2 belong to same parent or not we can iteratively check until index and value in parent becomes the same, that is the parent

- There is something called path compression, at every moment you try to store the top most parent at every instance of finding the query
- When finding the root when a new element comes or just querying we try to compress the path

Method-2: Union By Size
- This is almost same as union by rank, but instead we add size to the root if we add children
-
*/
type DisjointSet struct {
	ParentArr []int
	RankArr   []int
	SizeArr   []int
}

func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		ParentArr: make([]int, n+1),
		SizeArr:   make([]int, n+1),
		RankArr:   make([]int, n+1),
	}

	for i := 0; i <= n; i++ {
		ds.ParentArr[i] = i
		ds.SizeArr[i] = 1
	}

	return ds
}

func (ds *DisjointSet) FindRoot(x int) int {

	if x == ds.ParentArr[x] {
		return x
	}

	return ds.FindRoot(ds.ParentArr[x])

}

func (ds *DisjointSet) find(u, v int) bool {
	return ds.FindRoot(u) == ds.FindRoot(v)
}

func (ds *DisjointSet) unionByRank(u, v int) {
	rootU := ds.FindRoot(u)
	rootV := ds.FindRoot(v)

	if rootU == rootV {
		return
	}

	if ds.RankArr[rootU] > ds.RankArr[rootV] {
		ds.ParentArr[rootV] = rootU
	} else if ds.RankArr[rootU] < ds.RankArr[rootV] {
		ds.ParentArr[rootU] = rootV
	} else {
		ds.ParentArr[rootU] = rootV
		ds.RankArr[rootU]++
	}
}

func (ds *DisjointSet) unionBySize(u, v int) {
	rootU := ds.FindRoot(u)
	rootV := ds.FindRoot(v)

	if rootU == rootV {
		return
	}

	if ds.SizeArr[rootU] > ds.SizeArr[rootV] {
		ds.ParentArr[rootV] = rootU
		ds.SizeArr[rootU] += ds.SizeArr[rootV]
	} else {
		ds.ParentArr[rootU] = rootV
		ds.SizeArr[rootV] += ds.SizeArr[rootU]
	}
}
