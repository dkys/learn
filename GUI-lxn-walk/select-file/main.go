package main

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
	"io"
	"os"
)

func main() {
	mw, _ := walk.NewMainWindow()
	mw.SetTitle(`上传文件`)   // 窗口标题
	mw.SetName(`uploads`) // 程序名称

	// 设置图标
	img, _ := walk.ImageFrom(`check.ico`)
	mw.SetIcon(img)

	// 窗口尺寸设置
	mw.SetSize(walk.Size{300, 40})
	mw.SetMinMaxSize(walk.Size{300, 40}, walk.Size{500, 100})
	// 设置窗口位置：在屏幕的中间
	mw.SetX((int(win.GetSystemMetrics(win.SM_CXSCREEN)) - mw.Width()) / 2)
	mw.SetY((int(win.GetSystemMetrics(win.SM_CYSCREEN)) - mw.Height()) / 2)

	// 窗口布局 ： 纵向排列
	l, _ := VBox{}.Create()
	mw.SetLayout(l)

	// 添加一个分组容器
	gbox, _ := walk.NewGroupBox(mw)
	gbox.SetSize(walk.Size{300, 20})                            // 设置容器尺寸（默认）
	gbox.SetMinMaxSize(walk.Size{300, 50}, walk.Size{500, 500}) // 设置容器最大和最小尺寸
	// 设置分组容器布局： 横向排列
	l2, _ := HBox{}.Create()
	gbox.SetLayout(l2)

	// 在分组容器中添加输入框
	edit, _ := walk.NewTextEdit(gbox)
	//edit.SetMinMaxSizePixels(walk.Size{200,20},walk.Size{200,20})
	edit.SetMinMaxSize(walk.Size{100, 30}, walk.Size{100, 30})
	edit.SetReadOnly(true)
	//edit.SetText("请选择文件")

	// 在分组容器中添加按钮
	btn, _ := walk.NewPushButton(gbox)
	//btn.SetMinMaxSize(walk.Size{100, 30},walk.Size{100, 30})
	btn.SetText(`打开`)
	btn.Clicked().Attach(func() {
		dlg := SelectFile(mw)
		edit.SetText(dlg.FilePath)
		s := fmt.Sprintf("Select : %s\r\n", dlg.FilePath)
		//edit.AppendText(s)
		fmt.Println(s)
		fmt.Println(dlg)
	})

	// 在分组容器中添加按钮
	btnSave, _ := walk.NewPushButton(gbox)
	//btnSave.SetMinMaxSize(walk.Size{100, 30},walk.Size{100, 30})
	btnSave.SetSize(walk.Size{50, 20})
	btnSave.SetText(`另存为`)
	btnSave.Clicked().Attach(func() {
		SaveFile(mw, edit)
	})

	mw.Show()
	mw.Run()

}

// btn1 单击事件
func SelectFile(mw *walk.MainWindow) walk.FileDialog {
	dlg := walk.FileDialog{
		Title:  `选择文件`,
		Filter: `*`,
	}
	dlg.ShowOpen(mw)
	return dlg
}

// btnSave 单击事件
func SaveFile(mw *walk.MainWindow, edit *walk.TextEdit) {
	dlg := walk.FileDialog{
		Title: `另存为`,
	}
	dlg.ShowSave(mw)

	// 打开源文件地址
	fi, _ := os.Open(edit.Text())
	// 创建并打开新文件
	f, _ := os.OpenFile(dlg.FilePath, os.O_CREATE, os.ModePerm)
	// 拷贝内容到新文件
	io.Copy(f, fi)
	// 关闭文件
	fi.Close()
	f.Close()
}
