package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	col int
	row int
	val int
}

type Cord struct {
	x int
	y int
}

func search(arr [][]int, start, end Cord, sizeX, sizeY int) ([]Cord, error) {
	var err error
	queue := make([]Cord, 0)
	queue = append(queue, start)
	nodes := make(map[Cord]Node)
	nodes[start] = Node{col: start.x, row: start.y, val: arr[start.x][start.y]}
	ans := make([]Cord, 0)
	for i := 0; i < len(queue); i++ {
		for add := 0; add < 4; add++ {
			x := 0
			y := 0
			switch add {
			case 0:
				x = 1
			case 1:
				x = -1
			case 2:
				y = 1
			case 3:
				y = -1
			}
			isValid := check(arr, queue[i].x+x, queue[i].y+y, sizeX, sizeY)
			if isValid {
				if nd, ok := nodes[Cord{x: queue[i].x + x, y: queue[i].y + y}]; ok && nd.val <= nodes[queue[i]].val+arr[queue[i].x][queue[i].y] {
					continue
				}
				queue = append(queue, Cord{queue[i].x + x, queue[i].y + y})
				nodes[Cord{x: queue[i].x + x, y: queue[i].y + y}] = Node{col: queue[i].x, row: queue[i].y, val: nodes[queue[i]].val + arr[queue[i].x][queue[i].y]}
			}
		}
	}

	if _, ok := nodes[end]; !ok {
		err = errors.New("Пути не существует")
	}

	for pair := end; ; {
		ans = append(ans, pair)
		if pair == start {
			break
		}
		pair = Cord{x: nodes[pair].col, y: nodes[pair].row}
	}

	return ans, err

}

func check(arr [][]int, x, y, sizeX, sizeY int) bool {
	if y < 0 || y >= sizeY || x < 0 || x >= sizeX || arr[x][y] == 0 {
		return false
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	size := strings.Split(input, " ")
	wrongInput := "Неверный ввод"
	cords := make([]int, 4)
	if len(size) > 2 {
		log.Fatal(wrongInput)
	}
	rows, err := strconv.Atoi(size[0])
	if err != nil || rows < 0 {
		log.Fatal(wrongInput)
	}
	cols, err := strconv.Atoi(size[1])
	if err != nil || cols < 0 {
		log.Fatal(wrongInput)
	}

	arr := make([][]int, cols)

	for i := 0; i < rows; i++ {
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		rowArr := strings.Split(input, " ")
		for j := 0; j < cols; j++ {
			part, err := strconv.Atoi(rowArr[j])
			if err != nil || part < 0 {
				log.Fatal(wrongInput)
			}
			arr[i] = append(arr[i], part)
		}
	}
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	size = strings.Split(input, " ")
	if len(size) != 4 {
		log.Fatal(wrongInput)
	}
	for idx, val := range size {
		cords[idx], err = strconv.Atoi(val)
		if err != nil || cords[idx] < 0 {
			log.Fatal(wrongInput)
		}
	}

	start := Cord{x: cords[0], y: cords[1]}
	end := Cord{x: cords[2], y: cords[3]}

	ans, err := search(arr, start, end, rows, cols)
	if err != nil {
		log.Fatal(err)
	}
	for i := len(ans) - 1; i >= 0; i-- {
		fmt.Println(ans[i].x, ans[i].y)
	}
	fmt.Println(".")
}
