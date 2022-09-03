package tourofgo

import (
	"fmt"
)

type fruit struct {
	name string
	qty  int
}

func changeName(f *fruit, newName string) {
	f.name = newName
	fmt.Println("during change: ", f)
}

func main() {
	f := fruit{"apple", 4}

	fmt.Println("before change: ", f)
	changeName(&f, "watermelon")
	fmt.Println("after change: ", f)
}
