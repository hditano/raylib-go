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

func main() {
	page1 := storyPage{"This is my first, page", nil}
	page2 := storyPage{"This is my second page", nil}
	page3 := storyPage{"This is my third page", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	page1.playStory()
}
