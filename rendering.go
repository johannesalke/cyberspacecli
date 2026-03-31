package main

import (
	//glamour "charm.land/glamour/v2"
	lipgloss "charm.land/lipgloss/v2"
	"fmt"
	"strings"
)

var (
	basicBox = lipgloss.NewStyle().
			Width(90).
			MarginLeft(2).
			Foreground(lipgloss.Color("#ff9a10")).
			BorderForeground(lipgloss.Color("#744b0f"))

	boxTop = lipgloss.NewStyle().Inherit(basicBox).
		Border(lipgloss.RoundedBorder(), true, true, false, true).
		Padding(0, 2, 0, 2).
		MarginLeft(2).
		MarginTop(2)
	boxSides = lipgloss.NewStyle().Inherit(basicBox).
			Border(lipgloss.RoundedBorder(), false, true, false, true).
			Padding(0, 2, 0, 2).
			MarginLeft(2)
	boxBottom = lipgloss.NewStyle().Inherit(basicBox).
			Border(lipgloss.RoundedBorder(), false, true, true, true).
			Padding(0, 2, 0, 2).
			MarginLeft(2)
)

func RenderPost(elements ...string) error {
	N := len(elements)

	result := boxTop.Render(strings.TrimRight(elements[0], "\n")) + "\n"
	for _, element := range elements[1 : N-1] {
		result += boxSides.Render(strings.TrimRight(element, "\n")) + "\n"

	}
	result += boxBottom.Render(strings.TrimRight(elements[N-1], "\n"))

	fmt.Print(result)
	return nil
}
