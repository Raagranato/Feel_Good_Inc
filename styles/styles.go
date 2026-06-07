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


// var Selected = lipgloss.NewStyle().
//     Bold(true).
//     Foreground(lipgloss.Color("#000000")).
//     Background(lipgloss.Color("#FFFFFF")).
//     PaddingLeft(1).
//     PaddingRight(1)

// var Title = lipgloss.NewStyle().
//     Bold(true).
//     Foreground(lipgloss.Color("#FFFFFF"))


// var Welcome = lipgloss.NewStyle().
//     Bold(true).
//     Foreground(lipgloss.Color("#F5A623")).
//     Border(lipgloss.RoundedBorder()).
//     BorderForeground(lipgloss.Color("#F5A623")).
//     PaddingLeft(4).
//     PaddingRight(4).
//     Align(lipgloss.Center)

// var Lose = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
// var Win = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
// var Draw = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))


func Selected(s string) string {
    return lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#000000")).
        Background(lipgloss.Color("#FFFFFF")).
        PaddingLeft(1).
        PaddingRight(1).
        Render(s)
}

func Title(s string) string {
    return lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#FFFFFF")).
        Render(s)
}

func Welcome(s string) string {
    return lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color("#F5A623")).
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color("#F5A623")).
        PaddingLeft(4).
        PaddingRight(4).
        Align(lipgloss.Center).
        Render(s)
}

func Lose(s string) string {
    return lipgloss.NewStyle().
        Foreground(lipgloss.Color("9")).
        Render(s)
}

func Win(s string) string {
    return lipgloss.NewStyle().
        Foreground(lipgloss.Color("10")).
        Render(s)
}

func Draw(s string) string {
    return lipgloss.NewStyle().
        Foreground(lipgloss.Color("11")).
        Render(s)
}

func One(s string) string {
    // Azul Claro (Tranquilo)
    return lipgloss.NewStyle().Foreground(lipgloss.Color("#00BFFF")).Render(s) 
}

func Two(s string) string {
    // Verde (Atenção leve)
    return lipgloss.NewStyle().Foreground(lipgloss.Color("#32CD32")).Render(s) 
}

func Three(s string) string {
    // Amarelo/Dourado (Aviso)
    return lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700")).Render(s) 
}

func Four(s string) string {
    // Laranja (Perigo)
    return lipgloss.NewStyle().Foreground(lipgloss.Color("#FF8C00")).Render(s) 
}

func Five(s string) string {
    // Laranja Avermelhado (Muito perigo)
    return lipgloss.NewStyle().Foreground(lipgloss.Color("#FF4500")).Render(s) 
}

func Six(s string) string {
    // Vermelho Vivo (Crítico)
    return lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(s) 
}

func Seven(s string) string {
    // Vermelho Escuro (Quase morte)
    return lipgloss.NewStyle().Foreground(lipgloss.Color("#8B0000")).Render(s) 
}

func Eight(s string) string {
    // Roxo Escuro (Extremo / Cercado de bombas)
    return lipgloss.NewStyle().Foreground(lipgloss.Color("#800080")).Render(s) 
}

func Closed(s string) string {
    return lipgloss.NewStyle().Foreground(lipgloss.Color("#8BAC0F")).Render(s)
}

func Flag(s string) string {
    // Vermelho vibrante para dar destaque à marcação de bomba
    return lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Render(s)
}