package main 

import (
	"github.com/bartholdbos/sudoku/gridpkg"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading grid string")
	
	grid, err := gridpkg.NewGrid(os.Args[1])
	if err != nil{
		panic(err);
	}
	
	grid.PrintGrid()
	
	fmt.Println("Mapping grid")
	grid.MapGrid()
	
	for true {
		empty := grid.GetEmpty()
		if len(empty) == 0 {
			break;
		}
		
		for _, v := range empty{
			possibilities := gridpkg.GetPossible(v.GetBlock(), v.GetRow(), v.GetColum())
				
			if(len(possibilities) == 1){
				v.SetNumber(possibilities[0])
			}
		}
		grid.PrintGrid()
		fmt.Println("===============================")
	}

	fmt.Println("Solved!")
}