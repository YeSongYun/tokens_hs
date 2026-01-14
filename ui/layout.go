package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// BuildMainLayout 构建双栏主布局
func BuildMainLayout(state *AppState) fyne.CanvasObject {
	// 标题区
	header := BuildHeader()

	// 左栏：单价设置 + 消费计算
	priceCard := NewPriceSettingCard(state)
	costCard := NewCostCalcCard(state)
	leftColumn := container.NewVBox(
		container.NewPadded(priceCard),
		container.NewPadded(costCard),
	)

	// 右栏：反向计算 + 单价计算
	reverseCard := NewReverseCalcCard(state)
	priceCalcCard := NewPriceCalcCard(state)
	rightColumn := container.NewVBox(
		container.NewPadded(reverseCard),
		container.NewPadded(priceCalcCard),
	)

	// 双栏布局（等宽两列）
	twoColumns := container.NewGridWithColumns(2,
		leftColumn,
		rightColumn,
	)

	// 主布局：标题 + 分隔线 + 双栏内容
	return container.NewBorder(
		container.NewVBox(
			container.NewPadded(header),
			widget.NewSeparator(),
		),
		nil, nil, nil,
		container.NewPadded(twoColumns),
	)
}
