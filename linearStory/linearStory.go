package main

import "fmt"

/*
链表与文字游戏
 */

type storyPage struct {
	text     string
	nextPage *storyPage
}

func playStory(page *storyPage)  {
	//如果不进行处理会产生 invalid memory address or nil pointer dereference
	if page == nil{
		return
	}
	fmt.Println(page.text)
	playStory(page.nextPage)
}

func main() {
	page1 := storyPage{"It was a dark and stormy night.", nil}
	page2 := storyPage{"You are alone, and you need to find the sacred helmet before the bad guys do", nil}
	page3 := storyPage{"You see a troll ahead", nil}
	page1.nextPage = &page2
	page2.nextPage = &page3

	playStory(&page2)
}
