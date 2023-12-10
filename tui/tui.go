package tui

import (
	dayeight "github.com/FACorreiaa/aoc-2023/cmd/day-eight"
	dayfive "github.com/FACorreiaa/aoc-2023/cmd/day-five"
	dayfour "github.com/FACorreiaa/aoc-2023/cmd/day-four"
	daynine "github.com/FACorreiaa/aoc-2023/cmd/day-nine"
	dayone "github.com/FACorreiaa/aoc-2023/cmd/day-one"
	dayseven "github.com/FACorreiaa/aoc-2023/cmd/day-seven"
	daysix "github.com/FACorreiaa/aoc-2023/cmd/day-six"
	daythree "github.com/FACorreiaa/aoc-2023/cmd/day-three"
	daytwo "github.com/FACorreiaa/aoc-2023/cmd/day-two"
	"github.com/FACorreiaa/aoc-2023/lib"
	"github.com/FACorreiaa/aoc-2023/tui/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg struct{ error }
)

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(lib.Item); ok {
			title = i.Title()
		} else {
			return nil
		}

		//
		mapFunction := map[string]func(){
			"Day 1": dayone.Start,
			"Day 2": daytwo.Start,
			"Day 3": daythree.Start,
			"Day 4": dayfour.Start,
			"Day 5": dayfive.Start,
			"Day 6": daysix.Start,
			"Day 7": dayseven.Start,
			"Day 8": dayeight.Start,
			"Day 9": daynine.Start,
		}
		switch msg := msg.(type) {

		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.choose):
				m.NewStatusMessage(statusMessageStyle("You chose " + title))
				return func() tea.Msg {
					mapFunction[title]() // Call the selected function
					return nil
				}
			case key.Matches(msg, keys.remove):
				index := m.Index()
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.remove.SetEnabled(false)
				}
				return m.NewStatusMessage(statusMessageStyle("Deleted " + title))
			}
		}

		return nil
	}

	help := []key.Binding{keys.choose, keys.remove}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type delegateKeyMap struct {
	choose key.Binding
	remove key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.

func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.choose,
		d.remove,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.

func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.choose,
			d.remove,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
		remove: key.NewBinding(
			key.WithKeys("x", "backspace"),
			key.WithHelp("x", "delete"),
		),
	}
}

// layout

var (
	appStyle = lipgloss.NewStyle().Padding(1, 2)

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFDF5")).
			Background(lipgloss.Color("#25A065")).
			Padding(0, 1)

	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
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

type mode int

const (
	nav mode = iota
	edit
	create
)

type model struct {
	list          list.Model
	itemGenerator *lib.RandomItemGenerator
	keys          *listKeyMap
	delegateKeys  *delegateKeyMap
	subMenu       bool
	inSubMenu     bool
	mode          mode
	quitting      bool
	input         textinput.Model
}

func NewModel() model {
	var (
		itemGenerator lib.RandomItemGenerator
		delegateKeys  = newDelegateKeyMap()
		listKeys      = newListKeyMap()
	)

	// Make initial list of items
	const numItems = 25
	items := make([]list.Item, numItems)
	for i := 0; i < numItems; i++ {
		items[i] = itemGenerator.Next()
	}

	// Setup list
	delegate := newItemDelegate(delegateKeys)
	menuList := list.New(items, delegate, 0, 0)
	menuList.Title = "Advent of Code"
	menuList.Styles.Title = titleStyle
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
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := appStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		if m.input.Focused() {
			if key.Matches(msg, constants.Keymap.Enter) {
				m.input.SetValue("")
				m.mode = nav
				m.input.Blur()
			}
			if key.Matches(msg, constants.Keymap.Back) {
				m.input.SetValue("")
				m.mode = nav
				m.input.Blur()
			}
			// only log keypresses for the input field when it's focused
			m.input, cmd = m.input.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			switch {
			case key.Matches(msg, constants.Keymap.Create):
				m.mode = create
				m.input.Focus()
				cmd = textinput.Blink
			case key.Matches(msg, constants.Keymap.Quit):
				m.quitting = true
				return m, tea.Quit
			case key.Matches(msg, constants.Keymap.Enter):
				activeProject := m.list.SelectedItem().(lib.Item)
				entry := InitEntry(activeProject.Title(), constants.P)
				return entry.Update(constants.WindowSize)
			case key.Matches(msg, constants.Keymap.Rename):
				m.mode = edit
				m.input.Focus()
				cmd = textinput.Blink
			//case key.Matches(msg, constants.Keymap.Delete):
			//	items := m.list.Items()
			//	if len(items) > 0 {
			//		cmd = deleteProjectCmd(m.getActiveProjectID(), constants.Pr)
			//	}
			default:
				m.list, cmd = m.list.Update(msg)
			}
			cmds = append(cmds, cmd)
		}

		switch {
		case key.Matches(msg, m.keys.toggleSpinner):
			cmd := m.list.ToggleSpinner()
			return m, cmd
		//case key.Matches(msg, m.keys.chooseItem):
		//	cmd := m.list.ToggleSpinner()
		//	fmt.Printf("%#v", msg)
		//	fmt.Printf("%#v", m.keys.chooseItem)
		//	return m, cmd

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
			m.delegateKeys.remove.SetEnabled(true)
			newItem := m.itemGenerator.Next()
			insCmd := m.list.InsertItem(0, newItem)
			statusCmd := m.list.NewStatusMessage(statusMessageStyle("Added " + newItem.Title()))
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
		return m.list.View()
	}

	return appStyle.Render(m.list.View())
}
