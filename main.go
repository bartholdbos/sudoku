package main 

import (
	"github.com/bartholdbos/sudoku/gridpkg"
	"fmt"
)

func main() {
	fmt.Println("Reading grid string")
	
	grid, err := gridpkg.NewGrid(".4..7...39.215...86...2.41.57...6..9..834.6..3...1..27.19.6...48...319.24...9..8.")
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
}