package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	maze := readMaze()
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
			}
		}

		fmt.Println()
	}

}
