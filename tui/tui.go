package tui

import (
	dayone "github.com/FACorreiaa/aoc-2023/cmd/day-one"
	"github.com/FACorreiaa/aoc-2023/common"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg struct{ error }
)

const (
	MenuView View = iota
	SubmenuView
	// Add other views as needed
)

type Solution struct {
	Title       string
	Description string
	ReadmePath  string
}

// r.titles = []string{
// "Day 1",
// "Day 2",
// "Day 3",
// "Day 4",
// "Day 5",
// "Day 6",
// "Day 7",
// "Day 8",
// "Day 9",
// }
//
// r.description = []string{
// "Trebuchet?!",
// "Cube Conundrum",
// "Gear Ratios",
// "Scratchcards",
// "If You Give A Seed A Fertilizer",
// "Wait for it",
// "Camel Cards",
// "Haunted Wasteland",
// "Mirage Maintenance",
// }
type listKeyMap struct {
	toggleSpinner    key.Binding
	toggleTitleBar   key.Binding
	toggleStatusBar  key.Binding
	togglePagination key.Binding
	toggleHelpMenu   key.Binding
	insertItem       key.Binding
	chooseItem       key.Binding
}

type View int
type model struct {
	selectedProjectName string
	list                list.Model
	itemGenerator       *common.RandomItemGenerator
	keys                *listKeyMap
	delegateKeys        *common.DelegateKeyMap
	subMenu             bool
	viewport            viewport.Model
	paginator           paginator.Model
	quitting            bool
	error               string
	activeView          View
}

func NewListKeyMap() *listKeyMap {
	return &listKeyMap{
		insertItem: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "add item"),
		),
		toggleSpinner: key.NewBinding(
			key.WithKeys("s"),
			key.WithHelp("s", "toggle spinner"),
		),
		toggleTitleBar: key.NewBinding(
			key.WithKeys("T"),
			key.WithHelp("T", "toggle title"),
		),
		toggleStatusBar: key.NewBinding(
			key.WithKeys("S"),
			key.WithHelp("S", "toggle status"),
		),
		togglePagination: key.NewBinding(
			key.WithKeys("P"),
			key.WithHelp("P", "toggle pagination"),
		),
		toggleHelpMenu: key.NewBinding(
			key.WithKeys("H"),
			key.WithHelp("H", "toggle help"),
		),
		chooseItem: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
	}
}

var (
	itemGenerator common.RandomItemGenerator
	delegateKeys  = common.NewDelegateKeyMap()
	listKeys      = NewListKeyMap()
)

func InitProject() (tea.Model, tea.Cmd) {
	// Make initial list of items
	const numItems = 25
	items := make([]list.Item, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = itemGenerator.Next()
	}

	// Setup list
	delegate := common.NewItemDelegate(delegateKeys, DaySolutions())
	menuList := list.New(items, delegate, 0, 0)
	menuList.Title = "Advent of Code"
	menuList.Styles.Title = common.TitleStyle
	menuList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			listKeys.toggleSpinner,
			listKeys.insertItem,
			listKeys.toggleTitleBar,
			listKeys.toggleStatusBar,
			listKeys.togglePagination,
			listKeys.toggleHelpMenu,
		}
	}
	m := model{list: menuList,
		keys:          listKeys,
		delegateKeys:  delegateKeys,
		itemGenerator: &itemGenerator}

	return m, func() tea.Msg { return errMsg{nil} }
}

// DaySolutions complete with next days later
func DaySolutions() map[string]func() tea.Msg {
	return map[string]func() tea.Msg{
		"Day 1": dayone.Start,
	}
}
