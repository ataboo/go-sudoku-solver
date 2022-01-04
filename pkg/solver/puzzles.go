package solver

// Easy
var TestGrid1 = []int{
	1, 0, 0, 5, 0, 0, 0, 4, 3,
	4, 7, 0, 1, 0, 2, 0, 5, 0,
	0, 0, 0, 0, 0, 0, 2, 9, 1,
	9, 0, 3, 4, 0, 0, 5, 8, 6,
	6, 0, 0, 0, 0, 5, 0, 2, 0,
	0, 5, 0, 0, 0, 0, 1, 7, 0,
	3, 0, 5, 8, 0, 0, 9, 6, 2,
	0, 2, 6, 3, 0, 9, 8, 0, 5,
	8, 0, 0, 0, 5, 0, 0, 0, 0,
}

// Medium
var TestGrid2 = []int{
	6, 7, 2, 3, 5, 1, 9, 0, 8,
	0, 3, 0, 0, 6, 0, 2, 7, 5,
	0, 0, 9, 0, 0, 7, 6, 3, 1,
	0, 0, 7, 9, 4, 6, 8, 2, 3,
	3, 9, 6, 1, 2, 8, 7, 0, 4,
	2, 0, 0, 0, 7, 3, 1, 6, 0,
	0, 1, 0, 6, 9, 2, 3, 0, 7,
	7, 6, 3, 0, 1, 5, 0, 0, 2,
	9, 2, 0, 7, 3, 4, 0, 1, 6,
}

// Hard
var TestGrid3 = []int{
	0, 0, 0, 9, 2, 0, 1, 0, 3,
	3, 8, 0, 0, 0, 0, 0, 0, 2,
	0, 2, 9, 3, 0, 4, 0, 0, 0,
	0, 0, 6, 0, 0, 0, 0, 0, 0,
	9, 1, 2, 6, 8, 7, 0, 0, 0,
	0, 0, 8, 0, 0, 0, 6, 7, 0,
	2, 0, 3, 0, 5, 6, 0, 0, 7,
	0, 6, 0, 2, 0, 1, 8, 0, 0,
	0, 0, 0, 4, 0, 0, 0, 0, 6,
}

// Evil
var TestGrid4 = []int{
	0, 0, 2, 0, 0, 0, 0, 6, 0,
	5, 6, 0, 3, 0, 0, 0, 0, 7,
	0, 0, 8, 0, 0, 5, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 0, 0, 8,
	6, 3, 0, 0, 0, 9, 0, 1, 0,
	0, 2, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 7, 0, 0, 4, 0, 0,
	9, 1, 0, 0, 0, 3, 0, 8, 0,
	0, 0, 5, 0, 0, 0, 0, 0, 0,
}

// Evil
var TestGrid5 = []int{
	0, 0, 0, 5, 0, 0, 0, 0, 0,
	1, 0, 0, 8, 0, 2, 0, 9, 0,
	0, 0, 9, 0, 0, 0, 4, 0, 0,
	8, 0, 0, 9, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 7, 0, 0, 0, 0,
	0, 6, 0, 0, 0, 0, 0, 0, 3,
	7, 0, 0, 0, 4, 0, 0, 0, 0,
	0, 8, 0, 2, 0, 7, 6, 0, 0,
	0, 0, 0, 0, 5, 0, 0, 2, 0,
}
