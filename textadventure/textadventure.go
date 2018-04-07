package main

type choices struct {
	cmd        string
	desription string
	nextNode       *storyNode
	nextChoice       *choices
}

type storyNode struct {
	text    string
	choices *choices
}

func (node *storyNode)addChoice(cmd string, description string, nextNode *storyNode)  {
	choice := &choices{cmd, description, nextNode, nil}
	if node.choices == nil{
		node.choices = choice
	} else {
		currentChoice := node.choices
		for currentChoice.nextNode!= nil{
			currentChoice = currentChoice.nextChoice
		}
		currentChoice.nextChoice = choice
	}
}


func main() {

}
