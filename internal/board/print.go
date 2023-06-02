package board

import "fmt"

const (
	letters   = `    A  B  C  D  E  F  G  H  I  J  K  L  M  N  O`
	horizLine = `   ---------------------------------------------`
)

func Print(b Board) {
	fmt.Print(letters + "\n" + horizLine + "\n")
	for row := 0; row < size; row++ {
		fmt.Printf("%2.f|", float32(row+1))
		for col := 0; col < size; col++ {
			if b[row][col] != nil {
				fmt.Print(" " + b[row][col].GetLetter() + " ")
			} else {
				fmt.Print(" * ")
			}
		}
		fmt.Printf("|%d\n", row+1)
	}
	fmt.Print(horizLine + "\n" + letters + "\n")
}
