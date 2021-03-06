package main

import "fmt"

/*
链表与文字游戏
 */

type storyPage struct {
	text     string
	nextPage *storyPage
}

func (page *storyPage)playStory()  {
	// 将递归改为循环
	for page != nil{
		fmt.Println(page.text)
		page = page.nextPage
	}
}

func (page *storyPage)addToEnd(text string)  {
	for page.nextPage!=nil  {
		page = page.nextPage
	}
	page.nextPage = &storyPage{text, nil}
}

// O(1)
func (page *storyPage)AddAfter(text string)  {
	newPage := &storyPage{text, page.nextPage}
	page.nextPage = newPage
}

// Delete

func main() {
	page1 := storyPage{"It was a dark and stormy night.", nil}
	page1.addToEnd("You climb into the attic, it is pitch black, you can't see a thing!")
	page1.addToEnd("You are eaten by a Grue")

	page1.AddAfter("Testing AddAfter")
	page1.playStory()
	// Functions - has return value - may also execute commands
	// Procedures - has no return value, just executes commands
	// Methods - functions attached to a struct/object/etc
}
