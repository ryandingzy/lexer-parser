package main

/*分析表*/
//ACTION符号表
var sym = [12][6]rune {
	{'s',' ',' ','s',' ',' '},
	{' ','s',' ',' ',' ','a'},
	{'r','r','s','r','r','r'},
	{'r','r','r','r','r','r'},
	{'s',' ',' ','s',' ',' '},
	{'r','r','r','r','r','r'},
	{'s',' ',' ','s',' ',' '},
	{'s',' ',' ','s',' ',' '},
	{' ','s',' ','s',' ',' '},
	{'r','r','s','r','r','r'},
	{'r','r','r','r','r','r'},
	{'r','r','r','r','r','r'}}

//ACTION数字表
var snum = [12][6]int {
	{5, 0, 0, 4, 0, 0},
	{0, 6, 0, 0, 0, 0},
	{2, 2, 7, 2, 2, 2},
	{4, 4, 4, 4, 4, 4},
	{5, 0, 0, 4, 0, 0},
	{6, 6, 6, 6, 6, 6},
	{5, 0, 0, 4, 0, 0},
	{5, 0, 0, 4, 0, 0},
	{0, 6, 0, 11, 0, 0},
	{1, 1, 7, 1, 1, 1},
	{3, 3, 3, 3, 3, 3},
	{5, 5, 5, 5, 5, 5}}

//GOTO表
var go2 = [12][3]int {
	{1, 2, 3},
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
	{8, 2, 3},
	{0, 0, 0},
	{0, 9, 3},
	{0, 0, 10},
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0},
	{0, 0, 0}}