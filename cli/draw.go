package cli

import (
	"fmt"
	"github.com/gomachan46/satistuffed/model"
	"strings"
)

func Draw(facility *model.Facility, depth int) {
	fmt.Printf(" %s (産: %g)\n", facility.Recipe.Name, facility.Amount)
	drawChildren("", facility, depth)
	fmt.Println("")
}

func drawChildren(indent string, facility *model.Facility, depth int) {
	children := facility.Children
	if strings.Contains(facility.Recipe.Name, " x ") {
		children = children[0:1]
	}

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
			fmt.Printf("%s %s (産: %g, 消: %g, 余: %g)\n", s, v.Facility.Recipe.Name, v.Facility.Amount, v.Facility.Amount-v.Remain, v.Remain)
		} else {
			fmt.Printf("%s %s (消: %g)\n", s, v.Facility.Recipe.Name, v.Facility.Amount)
		}

		drawChildren(indent+a, v.Facility, depth-1)
	}
}
