package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	maze := readMaze()

	for i := range maze {
		for _, character := range maze[i] {
			fmt.Printf(" %s ", string(character))
		}

		fmt.Printf("\n")
	}

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
