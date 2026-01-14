package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// BuildHeader 构建标题区域
func BuildHeader() fyne.CanvasObject {
	title := canvas.NewText("Tokens 价格计算器", color.RGBA{R: 100, G: 200, B: 255, A: 255})
	title.TextSize = 24
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	subtitle := canvas.NewText("快速换算 Token 价格与消费", color.RGBA{R: 150, G: 150, B: 150, A: 255})
	subtitle.TextSize = 12
	subtitle.Alignment = fyne.TextAlignCenter

	return container.NewVBox(title, subtitle)
}
