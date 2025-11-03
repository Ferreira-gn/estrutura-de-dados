package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Dom-Garotom/estruturaDeDados/maze/stack"
)

func main() {
	maze := readMaze()
	printMaze(maze)
	solvedMaze(maze)
	printMaze(maze)
}

func readMaze() [][]rune {
	file, err := os.Open("maze.txt")

	if err != nil {
		fmt.Printf("Erro ao abrir arquivo de mapeamento do labirinto : %v", err)
		return nil
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var maze [][]rune

	for scanner.Scan() {
		// Escaneia a linha atual do arquivo e converte em um array de strings contendo cada caractere da linha.
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			break
		}

		maze = append(maze, []rune(line))
	}

	return maze
}

func printMaze(maze [][]rune) {

	for i := range maze {

		for _, character := range maze[i] {
			switch character {
			case '1':
				fmt.Printf("\033[1;35m %s \033[m", string(character))
			case '0':
				fmt.Printf("\033[1;32m %s \033[m", string(character))
			case 'm':
				fmt.Printf("\033[1;30m %s \033[m", string(character))
			case 'e':
				fmt.Printf("\033[1;33m %s \033[m", string(character))
			case '.':
				fmt.Printf("\033[1;33m %s \033[m", string(character))
			}
		}

		fmt.Println()
	}

}

func solvedMaze(maze [][]rune) {
	start := findCharacter(maze, 'm')
	exit := findCharacter(maze, 'e')

	if start == nil || exit == nil {
		fmt.Print("Erro ao buscar pelos valos do ponto inicial ou ponto de saída do labirinto")
		return
	}

	pathTraveled := stack.NewStack(*start)
	moves := []stack.Position{
		{X: -1, Y: 0}, // cima
		{X: 1, Y: 0},  // baixo
		{X: 0, Y: -1}, // esquerda
		{X: 0, Y: 1},  // direita
	}

	for pathTraveled.Len() > 0 {
		currentPosition := pathTraveled.ViewTheTop()

		if currentPosition.X == exit.X && currentPosition.Y == exit.Y {
			fmt.Printf("Saída do labirinto encontrada\n")
			maze[currentPosition.X][currentPosition.Y] = '.'
			return
		}

		found := false

		for _, currentMove := range moves {
			// Soma o movimento com a posição atual para chegar no novo ponto de interesse
			newDimenssionX := currentMove.X + currentPosition.X
			newDimenssionY := currentMove.Y + currentPosition.Y

			if isValidMove(maze, newDimenssionX, newDimenssionY) {
				pathTraveled.Append(newDimenssionX, newDimenssionY)
				maze[currentPosition.X][currentPosition.Y] = '.'
				found = true
				break
			}
		}

		if !found {
			pathTraveled.Pop()
		}
	}

	fmt.Printf("Saída do labirinto não encontrada\n")
}

func findCharacter(maze [][]rune, findValue rune) *stack.Position {

	for i := range maze {
		for j := range maze[i] {

			if maze[i][j] == findValue {
				return &stack.Position{
					X: i,
					Y: j,
				}
			}

		}
	}

	return nil
}

func isValidMove(maze [][]rune, x, y int) bool {
	linesSize := len(maze)
	columnsSize := len(maze[0])

	if x < 0 || y < 0 || x >= linesSize || y >= columnsSize {
		return false
	}

	return maze[x][y] == 'e' || maze[x][y] == '0'
}
