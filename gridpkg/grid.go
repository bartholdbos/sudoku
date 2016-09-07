package gridpkg

import (
	"strconv"
	"fmt"
)

type grid struct{
	squares[81]square
	blocks[9][9]*square
	rows[9][9]*square
	colums[9][9]*square
}

type square struct{
	grid   *grid
	number int
	block  int
	row    int
	colum  int
}

func NewGrid(numbers string) (*grid, error){
	g := &grid{}
	for k, v := range []byte(numbers){
		var number int
		
		if(v == '.'){
			number = 0
		}else{
			var err error
			number, err = strconv.Atoi(string(v))
			if err != nil{
				return g, err
			}
		}
		
		g.squares[k].number = number
		g.squares[k].grid = g
	}
	return g, nil
}

func (grid *grid) PrintGrid(){
	fmt.Print("[")
	for i := 0; i < len(grid.squares); i++{
		if(i % 3 == 0 && i != 0){
			fmt.Print("]  ");
		}
		if(i % 9 == 0 && i != 0){
			fmt.Println();
		}
		if(i % 27 == 0 && i != 0){
			fmt.Println();
		}
		if(i % 3 == 0 && i != 0){
			fmt.Print("[");
		}
		if(i % 3 != 0 && i != 0){
			fmt.Print(", ");
		}
		fmt.Print(grid.squares[i].number)
	}
	fmt.Println("]")
}

func (grid *grid) MapGrid(){
	var blockindex int
	
	for i := 0; i < len(grid.blocks); i++ {
		blockindex = (i / 3 * 27) + (i % 3 * 3)
		
		for j := 0; j < len(grid.blocks[i]); j++ {
			grid.blocks[i][j] = &grid.squares[blockindex + (j / 3 * 9) + (j % 3)]
			grid.blocks[i][j].block = i
		}
	}
	
	for i := 0; i < len(grid.rows); i++ {
		for j := 0; j < len(grid.rows[i]); j++ {
			grid.rows[i][j] = &grid.squares[(i * 9) + j]
			grid.rows[i][j].row = i
		}
	}
	
	for i := 0; i < len(grid.colums); i++ {
		for j := 0; j < len(grid.colums[i]); j++ {
			grid.colums[i][j] = &grid.squares[(j * 9) + i]
			grid.colums[i][j].colum = i
		}
	}
}

func (grid *grid) GetBlock(block int) ([9]*square){
	return grid.blocks[block]
}

func (grid *grid) GetRow(row int) ([9]*square){
	return grid.rows[row]
}

func (grid *grid) GetColum(colum int) ([9]*square){
	return grid.colums[colum]
}

func (square *square) GetNumber() (int){
	return square.number
}

func (square *square) SetNumber(number int){
	square.number = number
}

func (square *square) GetBlock() ([9]*square){
	return square.grid.GetBlock(square.block)
}

func (square *square) GetRow() ([9]*square){
	return square.grid.GetRow(square.row)
}

func (square *square) GetColum() ([9]*square){
	return square.grid.GetColum(square.colum)
}

func (grid *grid) GetEmpty() ([]*square){
	var empty []*square
	
	for i := 0; i < len(grid.squares); i++{
		if(grid.squares[i].GetNumber() == 0){
			empty = append(empty, &grid.squares[i])
		}
	}
	
	return empty
}

func GetPossible(numbers ...[9]*square) (possiblen []int){
	var possible [9]bool
	var squares []*square
	for _, v := range numbers{
		squares = append(squares, v[:]...)
	}
	
	for _, v := range squares{
		if v.GetNumber() == 0{
			continue;
		}
		possible[v.GetNumber() - 1] = true
	}
	
	for k, v := range possible{
		if !v{
			possiblen = append(possiblen, k + 1)
		}
	}
	
	return
}