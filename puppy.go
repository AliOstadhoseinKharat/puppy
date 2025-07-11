package puppy

import (
	"fmt"

	"github.com/AliOstadhoseinKharat/dog"
)

func Bark() string {
	return "Woof !"
}

func Barks() string {
	return "Woof Woof Woof !"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}

func version1() {
	fmt.Println("I'm from version 1.0.0")
}

func version2() {
	fmt.Println("I'm from version 1.1.0")
}
