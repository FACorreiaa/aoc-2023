package tui

import (
	"github.com/FACorreiaa/aoc-2023/lib"
	"github.com/FACorreiaa/aoc-2023/lib/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg struct{ error }
	// UpdatedEntries holds the new entries from DB
)

const (
	MenuView View = iota
	SubmenuView
	// Add other views as needed
)

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
	itemGenerator       *lib.RandomItemGenerator
	keys                *listKeyMap
	delegateKeys        *lib.DelegateKeyMap
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
	}
}

var (
	itemGenerator lib.RandomItemGenerator
	delegateKeys  = lib.NewDelegateKeyMap()
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
	delegate := NewItemDelegate(delegateKeys)
	menuList := list.New(items, delegate, 0, 0)
	menuList.Title = "Advent of Code"
	menuList.Styles.Title = constants.TitleStyle
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
