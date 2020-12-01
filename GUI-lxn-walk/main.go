package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"strings"
)

func main() {
	window, _ := walk.NewMainWindow()

	// 设置窗体生成在屏幕的正中间
	window.SetX((int(win.GetSystemMetrics(0)) - window.Width()) / 2)
	window.SetY((int(win.GetSystemMetrics(1)) - window.Height()) / 2)

	// 设置窗口图标
	img, _ := walk.ImageFrom(`check.ico`)
	window.SetIcon(img)
	// 设置窗体标题
	window.SetTitle(`hello world`)

	// 布局设置
	l, _ := VBox{}.Create() //样式，纵向
	window.SetLayout(l)

	// 设置窗体的宽高
	window.SetSize(walk.Size{400, 300})

	// 创建容器,容器内元素纵向分割排列
	box, _ := walk.NewHSplitter(window)

	// 创建文本框1
	edit1, _ := walk.NewTextEdit(box)
	edit1.SetX(10)
	edit1.SetY(10)

	// 创建文本框2
	edit2, _ := walk.NewTextEdit(box)
	edit2.SetReadOnly(true)
	edit2.SetX(10)
	edit2.SetY(10)

	// 添加按钮
	btn1, _ := walk.NewPushButton(window)
	btn1.SetText(`提交`)
	btn1.SetWidth(100)
	btn1.SetHeight(30)
	btn1.SetX(10)
	btn1.SetY(60)
	btn1.Clicked().Attach(func() {
		edit2.SetText(strings.ToUpper(edit1.Text()))
	})

	// 设置窗体为显示状态（默认：隐藏状态）
	window.Show()

	// 运行窗体
	window.Run()
}
