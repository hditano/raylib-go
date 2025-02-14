package main

// TODO
// NPC - talk to them, fight
// NPC move around the graph
// items that can picked up or placed down

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/hditano/basic/textadventure2/items"
)

type choice struct {
	cmd         string
	description string
	nextNode    *storyNode
}

type storyNode struct {
	text    string
	choices []*choice
	npc     *npc
	item    *item
}

type npc struct {
	name            []string
	dialogue        []string
	nextNode        *storyNode
	currentName     string
	currentDialogue string
}

type player struct {
	name         string
	inventory    []*item
	startingItem item
}

// TODO
type item struct {
	generation  bool
	name        string
	description string
}

func (myPlayer *player) Player(name string) {
	myPlayer.name = name

	myPlayer.startingItem = myPlayer.startingItem.itemGeneration()

	fmt.Printf("My item is: %s %s %s\n", myPlayer.name, myPlayer.startingItem.name, myPlayer.startingItem.description)
}

func (npcNode *npc) addNpc() {
	if npcNode.nextNode != nil {
		name := npcNode.name[rand.Intn(len(npcNode.name))]
		dialogue := npcNode.dialogue[rand.Intn(len(npcNode.dialogue))]

		npcNode.currentName = name
		npcNode.currentDialogue = dialogue
	}
}

func (inventory *item) itemGeneration() item {
	name := items.ItemsNames[rand.Intn(len(items.ItemsNames))]
	description := items.DescriptionNames[rand.Intn(len(items.DescriptionNames))]

	newItem := item{
		generation:  false,
		name:        name,
		description: description,
	}

	return newItem
}

func (node *storyNode) addChoice(cmd string, description string, nextNode *storyNode) {
	choice := &choice{cmd, description, nextNode}
	node.choices = append(node.choices, choice)
}

func (node *storyNode) render() {
	fmt.Println(node.text)
	if node.choices != nil {
		for _, choice := range node.choices {
			fmt.Println(choice.cmd, choice.description)
		}
	}
	if node.npc != nil {
		fmt.Printf("Npc %s says %s\n", node.npc.currentName, node.npc.currentDialogue)
	}
}

func (node *storyNode) executeCmd(cmd string) *storyNode {
	for _, choice := range node.choices {
		if strings.ToLower(choice.cmd) == strings.ToLower(cmd) {
			return choice.nextNode
		}
	}
	fmt.Println("Sorry I didn't understand that..")
	return node
}

var scanner *bufio.Scanner

func (node *storyNode) play() {
	node.render()
	if node.choices != nil {
		scanner.Scan()
		node.executeCmd(scanner.Text()).play()
	}
}

func printArray(a [3]string) {
	for _, e := range a {
		fmt.Println(e)
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)

	start := storyNode{text: `You are in large chamber, deep underground. You see three passages leading out. A north apssage leads into darkness.
							To the south, a passage appears to head upward. The eastern passages appears flat and well traveled`}

	darkRoom := storyNode{text: "it is pitch black. You canont see a thing."}

	darkRoomLit := storyNode{text: `The dark passage is now lit by your latern. You can continue north or head back south`}

	grue := storyNode{text: `While stumbling around in the darkness, you are eaten by a grue.`}

	trap := storyNode{text: `You head down the well traveled path when suddenly a trap door opens and you fall into a pit`}

	treasure := storyNode{text: `You arrive at a small chamber, filled with treasure!`}

	start.addChoice("N", "Go North", &darkRoom)
	start.addChoice("S", "Go South", &darkRoom)
	start.addChoice("E", "Go East", &trap)

	darkRoom.addChoice("S", "Try to go back south", &grue)
	darkRoom.addChoice("O", "Turn on Latern", &darkRoomLit)

	darkRoomLit.addChoice("N", "Go North", &treasure)
	darkRoomLit.addChoice("S", "Go South", &start)

	// NPCs

	npcS := npc{
		[]string{"Zoe", "Gime", "Seba", "Tomas"},
		[]string{"Hello, How are you?", "What are you doing Here!!"},
		&start,
		"",
		"",
	}

	start.npc = &npcS

	start.npc.addNpc()

	// Player

	myplayer := player{}

	myplayer.Player("Hernan")

	// Item Generation

	// Start Game

	start.play()

	fmt.Println()
	fmt.Println("You won")
}
