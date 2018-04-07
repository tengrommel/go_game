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

func main() {
	page1 := storyPage{"It was a dark and stormy night.", nil}
	page2 := storyPage{"You are alone, and you need to find the sacred helmet before the bad guys do", nil}
	page3 := storyPage{"You see a troll ahead", nil}
	page1.nextPage = &page2
	page2.nextPage = &page3

	page1.playStory()
	// Functions - has return value - may also execute commands
	// Procedures - has no return value, just executes commands
	// Methods - functions attached to a struct/object/etc
}
