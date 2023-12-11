package solution

//func NewItemDelegate(keys *DelegateKeyMap) list.DefaultDelegate {
//	d := list.NewDefaultDelegate()
//
//	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
//		var title string
//
//		if i, ok := m.SelectedItem().(Item); ok {
//			title = i.Title()
//		} else {
//			return nil
//		}
//
//		//
//		mapFunction := map[string]func(){
//			"Day 1": dayone.Start,
//			"Day 2": daytwo.Start,
//			"Day 3": daythree.Start,
//			"Day 4": dayfour.Start,
//			"Day 5": dayfive.Start,
//			"Day 6": daysix.Start,
//			"Day 7": dayseven.Start,
//			"Day 8": dayeight.Start,
//			"Day 9": daynine.Start,
//		}
//		switch msg := msg.(type) {
//		case tea.KeyMsg:
//			switch {
//			case key.Matches(msg, keys.Choose):
//				m.NewStatusMessage(constants.StatusMessageStyle("You chose " + title))
//				if action, ok := mapFunction[title]; ok {
//					// Call the Start function for the selected day
//					action()
//					return nil
//				}
//				//index := m.Index()
//				//m.RemoveItem(index)
//				//if len(m.Items()) == 0 {
//				//    keys.Back.SetEnabled(false)
//				//}
//				//return m.NewStatusMessage(constants.StatusMessageStyle("Deleted " + title))
//			case key.Matches(msg, keys.Back):
//
//			}
//		}
//
//		return nil
//	}
//
//	help := []key.Binding{keys.Choose, keys.Back}
//
//	d.ShortHelpFunc = func() []key.Binding {
//		return help
//	}
//
//	d.FullHelpFunc = func() [][]key.Binding {
//		return [][]key.Binding{help}
//	}
//
//	return d
//}
