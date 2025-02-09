package main

import (
	"bufio"
	"fmt"
	"os"
)

type storyNode struct {
	text    string
	yesPath *storyNode
	noPath  *storyNode
}

func (node *storyNode) play() {
	fmt.Println(node.text)
	scanner := bufio.NewScanner((os.Stdin))

	if node.yesPath != nil && node.noPath != nil {

		for {
			scanner.Scan()
			answer := scanner.Text()

			if answer == "yes" {
				node.yesPath.play()
				break
			} else if answer == "no" {
				node.noPath.play()
				break
			} else {
				fmt.Println("Thats not an answer")
			}
		}
	}
}

func main() {

	root := storyNode{"You are at the root node", nil, nil}
	winning := storyNode{"You have won", nil, nil}
	losing := storyNode{"You have lost", nil, nil}

	root.yesPath = &losing
	root.noPath = &winning

	root.play()

}
