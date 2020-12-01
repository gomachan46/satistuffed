package cli

import (
	"fmt"
	"github.com/gomachan46/satistuffed/model"
)

func Draw(facility *model.Facility, depth int) {
	fmt.Printf(" %s (総: %g)\n", facility.Recipe.Name, facility.Amount)
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
			fmt.Printf("%s %s (総: %g, 余: %g)\n", s, v.Facility.Recipe.Name, v.Facility.Amount, v.Remain)
		} else {
			fmt.Printf("%s %s (総: %g)\n", s, v.Facility.Recipe.Name, v.Facility.Amount)
		}

		drawChildren(indent+a, v.Facility, depth-1)
	}
}
