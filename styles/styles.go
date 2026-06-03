package styles

import "github.com/charmbracelet/lipgloss"

// var Selected = lipgloss.NewStyle().
//     Background(lipgloss.Color("15")).
//     Foreground(lipgloss.Color("0"))

// var Selected = lipgloss.NewStyle().
//     Bold(true).
//     Foreground(lipgloss.Color("#FFFFFF")).
//     Background(lipgloss.Color("#7B2FBE")).
//     PaddingLeft(2).
//     PaddingRight(2).
//     Border(lipgloss.RoundedBorder()).
//     BorderForeground(lipgloss.Color("#F5A623"))

// var Title = lipgloss.NewStyle().
//     Bold(true).
//     Foreground(lipgloss.Color("#F5A623")).
//     Border(lipgloss.DoubleBorder()).
//     BorderForeground(lipgloss.Color("#7B2FBE")).
//     PaddingLeft(4).
//     PaddingRight(4).
//     Align(lipgloss.Center)


var Selected = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#000000")).
    Background(lipgloss.Color("#FFFFFF")).
    PaddingLeft(1).
    PaddingRight(1)

var Title = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FFFFFF"))


var Welcome = lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#F5A623")).
    Border(lipgloss.RoundedBorder()).
    BorderForeground(lipgloss.Color("#F5A623")).
    PaddingLeft(4).
    PaddingRight(4).
    Align(lipgloss.Center)