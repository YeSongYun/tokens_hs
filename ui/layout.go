package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// BuildMainLayout 构建分层主布局
func BuildMainLayout(state *AppState) fyne.CanvasObject {
	// 标题区
	header := BuildHeader()

	// 顶部：单价设置（全宽）
	priceSettingCard := NewPriceSettingCard(state)

	// 中间：消费计算 + 反向计算（并排，体现关联性）
	middleRow := container.NewGridWithColumns(2,
		NewCostCalcCard(state),
		NewReverseCalcCard(state),
	)

	// 底部：单价计算（独立功能）
	priceCalcCard := NewPriceCalcCard(state)

	// 主布局：标题 + 分隔线 + 分层内容
	return container.NewBorder(
		container.NewVBox(
			container.NewPadded(header),
			widget.NewSeparator(),
		),
		nil, nil, nil,
		container.NewPadded(
			container.NewVBox(
				priceSettingCard,
				middleRow,
				priceCalcCard,
			),
		),
	)
}
