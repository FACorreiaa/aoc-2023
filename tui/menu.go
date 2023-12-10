package tui

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/lib"
	"github.com/FACorreiaa/aoc-2023/lib/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
	"os"
	"time"
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

func newListKeyMap() *listKeyMap {
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

type model struct {
	list          list.Model
	itemGenerator *lib.RandomItemGenerator
	keys          *listKeyMap
	delegateKeys  *lib.DelegateKeyMap
	subMenu       bool
}

func newModel() model {
	var (
		itemGenerator lib.RandomItemGenerator
		delegateKeys  = lib.NewDelegateKeyMap()
		listKeys      = newListKeyMap()
	)

	// Make initial list of items
	const numItems = 25
	items := make([]list.Item, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = itemGenerator.Next()
	}

	// Setup list
	delegate := lib.NewItemDelegate(delegateKeys)
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

	return model{
		list:          menuList,
		keys:          listKeys,
		delegateKeys:  delegateKeys,
		itemGenerator: &itemGenerator,
	}
}

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var commands []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := constants.AppStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, m.keys.toggleSpinner):
			cmd := m.list.ToggleSpinner()
			return m, cmd

		case key.Matches(msg, m.keys.toggleTitleBar):
			v := !m.list.ShowTitle()
			m.list.SetShowTitle(v)
			m.list.SetShowFilter(v)
			m.list.SetFilteringEnabled(v)
			return m, nil

		case key.Matches(msg, m.keys.toggleStatusBar):
			m.list.SetShowStatusBar(!m.list.ShowStatusBar())
			return m, nil

		case key.Matches(msg, m.keys.togglePagination):
			m.list.SetShowPagination(!m.list.ShowPagination())
			return m, nil

		case key.Matches(msg, m.keys.toggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil

		case key.Matches(msg, m.keys.insertItem):
			m.delegateKeys.Remove.SetEnabled(true)
			newItem := m.itemGenerator.Next()
			insCmd := m.list.InsertItem(0, newItem)
			statusCmd := m.list.NewStatusMessage(constants.StatusMessageStyle("Added " + newItem.Title()))
			return m, tea.Batch(insCmd, statusCmd)
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	commands = append(commands, cmd)

	return m, tea.Batch(commands...)
}

func (m model) View() string {
	if m.subMenu {
		// Render sub-menu UI
		return m.list.View()
	}

	return constants.AppStyle.Render(m.list.View())
}

func Start() error {
	rand.NewSource(time.Now().UTC().UnixNano())

	if _, err := tea.NewProgram(newModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return nil
}
