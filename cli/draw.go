package cli

import (
	"fmt"
	"github.com/gomachan46/satistuffed/model"
)

func Draw(facility *model.Facility) {
	fmt.Printf(" %s\n", facility.Recipe.Name)
	drawChildren("", facility)
	fmt.Println("")
}

func drawChildren(indent string, facility *model.Facility) {
	children := *facility.Children

	for i, v := range children {
		s := indent + " ├─"
		a := " │ "
		if i == len(children)-1 {
			s = indent + " └─"
			a = "   "
		}

		fmt.Printf("%s %s\n", s, v.Recipe.Name)

		drawChildren(indent+a, &v)
	}
}
