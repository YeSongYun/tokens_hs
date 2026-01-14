package ui

import (
	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

// Run 启动应用
func Run() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())

	myWindow := myApp.NewWindow("Tokens 价格计算器")

	// 加载图标
	iconData, err := ioutil.ReadFile("icon.png")
	if err == nil {
		myWindow.SetIcon(fyne.NewStaticResource("icon", iconData))
	}

	// 创建全局状态
	state := NewAppState()

	// 构建主布局
	content := BuildMainLayout(state)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(750, 580))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
