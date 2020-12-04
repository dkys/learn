package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

var mw *walk.MainWindow

type menu struct {
	Text        string
	IsMenu      bool
	items       []menu
	OnTriggered walk.EventHandler
}

//m,_ := NewMainWindow()
func main() {
	mw, _ = walk.NewMainWindow()
	mw.SetTitle(`menus`)
	mw.SetSize(walk.Size{400, 300})
	mw.SetX((int(win.GetSystemMetrics(win.SM_CXSCREEN)) - mw.Width()) / 2)
	mw.SetY((int(win.GetSystemMetrics(win.SM_CYSCREEN)) - mw.Width()) / 2)
	img, _ := walk.ImageFrom(`check.ico`)
	mw.SetIcon(img)
	mw.SetLayout(walk.NewVBoxLayout())

	// 设置菜单栏
	menuItems := []menu{
		{
			Text:   `&File`,
			IsMenu: true,
			items: []menu{
				menu{
					Text:   `Recent`,
					IsMenu: true,
					items: []menu{
						menu{
							Text: `A`,
						},
						menu{
							Text: `B`,
						},
						menu{
							Text: `C`,
						},
					},
				},
				menu{
					Text:        `E&xit`,
					OnTriggered: func() { mw.Close() },
				},
			},
		},
		{
			Text:   `View`,
			IsMenu: true,
			items: []menu{
				menu{
					Text:   `Recent`,
					IsMenu: true,
					items: []menu{
						menu{
							Text: `A`,
						},
						menu{
							Text: `B`,
						},
						menu{
							Text: `C`,
						},
					},
				},
				menu{
					Text:        `E&xit`,
					OnTriggered: func() { mw.Close() },
				},
			},
		},
	}
	for _, item := range menuItems {
		setMenuBar(mw.Menu(), item)
	}

	// 设置工具栏
	toolBar, _ := walk.NewToolBar(mw)
	toolBar.SetToolTipText(`工具栏`)
	m, _ := walk.NewMenu()
	ac, _ := toolBar.Actions().AddMenu(m)
	ac.SetText(`打开`)

	mw.Show()
	mw.Run()
}

// 设置菜单栏
func setMenuBar2(menu *walk.Menu, item Menu) {
	//for _, text := range texts {
	m, _ := walk.NewMenu()
	ac, _ := menu.Actions().AddMenu(m)
	ac.SetText(item.Text)
	ac.SetEnabled(true)
	ac.SetVisible(true)

	ac2 := walk.NewAction()
	ac2.SetText(`Test`)
	ac.Menu().Actions().Add(ac2)

	//m2, _ := NewMenu()
	//ac2, _ := ac.Menu().Actions().AddMenu(m2)
	//ac2.SetText(`test`)
	//}
}

// 设置菜单栏
func setMenuBar(menu *walk.Menu, item menu) {

	if item.IsMenu {
		m, _ := walk.NewMenu()
		ac, _ := menu.Actions().AddMenu(m)
		ac.SetText(item.Text)
		//ac.SetEnabled(true)
		//ac.SetVisible(true)
		if item.items != nil {
			for _, item := range item.items {
				setMenuBar(m, item)
			}
		}

		OnTriggered(item, ac)

	} else {
		ac2 := walk.NewAction()
		ac2.SetText(item.Text)
		menu.Actions().Add(ac2)
		OnTriggered(item, ac2)
	}
}

// 点击事件
func OnTriggered(item menu, action *walk.Action) {
	if item.OnTriggered != nil {
		action.Triggered().Attach(item.OnTriggered)
	}
}
