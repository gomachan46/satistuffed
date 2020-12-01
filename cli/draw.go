package cli

import (
	"fmt"
	"github.com/gomachan46/satistuffed/model"
)

func Draw(facility *model.Facility, depth int) {
	fmt.Printf(" %s\n", facility.Recipe.Name)
	drawChildren("", facility, depth)
	fmt.Println("")
}

func drawChildren(indent string, facility *model.Facility, depth int) {
	children := facility.Children

	for i, v := range children {
		if depth == 0 {
			continue
		}

		s := indent + " ├─"
		a := " │ "
		if i == len(children)-1 {
			s = indent + " └─"
			a = "   "
		}

		if v.Remain > 0 {
			fmt.Printf("%s %s (余り: %g)\n", s, v.Facility.Recipe.Name, v.Remain)
		} else {
			fmt.Printf("%s %s\n", s, v.Facility.Recipe.Name)
		}

		drawChildren(indent+a, v.Facility, depth-1)
	}
}
