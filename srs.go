package main

// PIECES[colour][rotation][block][x then y]
var PIECES = [7][4][4][2]int{
	// Red, index 0
	{
		{
			{0, 0},
			{1, 0},
			{1, 1},
			{2, 1},
		},
		{
			{2, 0},
			{2, 1},
			{1, 1},
			{1, 2},
		},
		{
			{0, 1},
			{1, 1},
			{1, 2},
			{2, 2},
		},
		{
			{1, 0},
			{1, 1},
			{0, 1},
			{0, 2},
		},
	},

	// Orange, index 2
	{
		{
			{0, 1},
			{1, 1},
			{2, 1},
			{2, 0},
		},
		{
			{1, 0},
			{1, 1},
			{1, 2},
			{2, 2},
		},
		{
			{0, 2},
			{0, 1},
			{1, 1},
			{2, 1},
		},
		{
			{0, 0},
			{1, 0},
			{1, 1},
			{1, 2},
		},
	},

	// Yellow, index 2
	{
		{
			{1, 0},
			{1, 1},
			{2, 0},
			{2, 1},
		},
		{
			{1, 0},
			{1, 1},
			{2, 0},
			{2, 1},
		},
		{
			{1, 0},
			{1, 1},
			{2, 0},
			{2, 1},
		},
		{
			{1, 0},
			{1, 1},
			{2, 0},
			{2, 1},
		},
	},

	// Green, index 3
	{
		{
			{0, 1},
			{1, 1},
			{1, 0},
			{2, 0},
		},
		{
			{1, 0},
			{1, 1},
			{2, 1},
			{2, 2},
		},
		{
			{0, 2},
			{1, 2},
			{1, 1},
			{2, 1},
		},
		{
			{0, 0},
			{0, 1},
			{1, 1},
			{1, 2},
		},
	},

	// Cyan, index 4
	{
		{
			{0, 1},
			{1, 1},
			{2, 1},
			{3, 1},
		},
		{
			{2, 0},
			{2, 1},
			{2, 2},
			{2, 3},
		},
		{
			{0, 2},
			{1, 2},
			{2, 2},
			{3, 2},
		},
		{
			{1, 0},
			{1, 1},
			{1, 2},
			{1, 3},
		},
	},

	// Blue, index 5
	{
		{
			{0, 0},
			{0, 1},
			{1, 1},
			{2, 1},
		},
		{
			{2, 0},
			{1, 0},
			{1, 1},
			{1, 2},
		},
		{
			{2, 2},
			{0, 1},
			{1, 1},
			{2, 1},
		},
		{
			{1, 0},
			{1, 1},
			{1, 2},
			{0, 2},
		},
	},

	// Pink, index 6
	{
		{
			{1, 1},
			{0, 1},
			{1, 0},
			{2, 1},
		},
		{
			{1, 1},
			{1, 2},
			{1, 0},
			{2, 1},
		},
		{
			{1, 1},
			{1, 2},
			{0, 1},
			{2, 1},
		},
		{
			{1, 1},
			{1, 2},
			{0, 1},
			{1, 0},
		},
	},
}

var KICKS = [8][4][2]int{
	{
		{-1, 0},
		{-1, 1},
		{0, -2},
		{-1, -2},
	},
	{
		{1, 0},
		{1, -1},
		{0, 2},
		{1, 2},
	},
	{
		{1, 0},
		{1, -1},
		{0, 2},
		{1, 2},
	},
	{
		{-1, 0},
		{-1, 1},
		{0, -2},
		{-1, -2},
	},
	{
		{1, 0},
		{1, 1},
		{0, -2},
		{1, -2},
	},
	{
		{-1, 0},
		{-1, -1},
		{0, 2},
		{-1, 2},
	},
	{
		{-1, 0},
		{-1, -1},
		{0, 2},
		{-1, 2},
	},
	{
		{1, 0},
		{1, 1},
		{0, -2},
		{1, -2},
	},
}

var I_KICKS = [8][4][2]int{
	{
		{-2, 0},
		{1, 0},
		{-2, -1},
		{1, 2},
	},
	{
		{2, 0},
		{-1, 0},
		{2, 1},
		{-1, -2},
	},
	{
		{-1, 0},
		{2, 0},
		{-1, -2},
		{2, -1},
	},
	{
		{1, 0},
		{-2, 0},
		{1, -2},
		{-2, 1},
	},
	{
		{2, 0},
		{-1, 0},
		{2, 1},
		{-1, -2},
	},
	{
		{-2, 0},
		{1, 0},
		{-2, -1},
		{1, 2},
	},
	{
		{1, 0},
		{-2, 0},
		{1, -2},
		{-2, 1},
	},
	{
		{-1, 0},
		{2, 0},
		{-1, 2},
		{2, -1},
	},
}
