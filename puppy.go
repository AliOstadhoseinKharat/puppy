package puppy

import (
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
