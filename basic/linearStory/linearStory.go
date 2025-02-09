package main

import "fmt"

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage) playStory() {
	for page != nil {
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func (page *storyPage) addToEnd(text string) {
	pageToAdd := &storyPage{text, nil}
	for page.nextPage != nil {
		page = page.nextPage
	}
	page.nextPage = pageToAdd
}

func (page *storyPage) addMiddlePage(text string) {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

// Delete

func checkStruct(page *storyPage) {
	fmt.Println(page.text)
	fmt.Println(page.nextPage)
}

func main() {
	page1 := storyPage{"This is my first page", nil}
	page1.addToEnd("This is my second page")
	page1.addToEnd("This is my third page")

	checkStruct(&page1)

	for page := &page1; page != nil; page = page.nextPage {
		checkStruct(page)
	}

	page1.addMiddlePage("This is a new story")

	page1.playStory()
}
