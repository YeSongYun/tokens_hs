package ui

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// NewPriceSettingCard 创建单价设置卡片
func NewPriceSettingCard(state *AppState) fyne.CanvasObject {
	millionEntry := widget.NewEntry()
	millionEntry.SetPlaceHolder("例如: 10")

	thousandEntry := widget.NewEntry()
	thousandEntry.SetPlaceHolder("自动计算")

	var updating bool

	millionEntry.OnChanged = func(text string) {
		if updating {
			return
		}
		val, err := strconv.ParseFloat(text, 64)
		if err != nil || val < 0 {
			return
		}
		state.SetPrice(val)
		updating = true
		thousandEntry.SetText(FormatNumber(val / 1000))
		updating = false
	}

	thousandEntry.OnChanged = func(text string) {
		if updating {
			return
		}
		val, err := strconv.ParseFloat(text, 64)
		if err != nil || val < 0 {
			return
		}
		state.SetPrice(val * 1000)
		updating = true
		millionEntry.SetText(FormatNumber(val * 1000))
		updating = false
	}

	return widget.NewCard("💰 单价设置", "",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widget.NewLabel("元/百万tokens"),
				millionEntry,
			),
			container.NewGridWithColumns(2,
				widget.NewLabel("元/千tokens"),
				thousandEntry,
			),
		),
	)
}

// NewCostCalcCard 创建消费计算卡片
func NewCostCalcCard(state *AppState) fyne.CanvasObject {
	tokensEntry := widget.NewEntry()
	tokensEntry.SetPlaceHolder("例如: 1000000")

	resultText := canvas.NewText("¥ 0.000000", color.RGBA{R: 100, G: 255, B: 150, A: 255})
	resultText.TextSize = 20
	resultText.TextStyle = fyne.TextStyle{Bold: true}

	updateResult := func() {
		tokens, err := strconv.ParseFloat(tokensEntry.Text, 64)
		if err != nil || tokens < 0 {
			resultText.Text = "¥ 0.000000"
			resultText.Refresh()
			return
		}
		price := state.GetPrice()
		cost := tokens * (price / 1000000)
		resultText.Text = fmt.Sprintf("¥ %.6f", cost)
		resultText.Refresh()
	}

	tokensEntry.OnChanged = func(text string) {
		updateResult()
	}

	state.Subscribe(func(price float64) {
		updateResult()
	})

	return widget.NewCard("📊 消费计算", "",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widget.NewLabel("Tokens 数量"),
				tokensEntry,
			),
			widget.NewSeparator(),
			container.NewHBox(
				widget.NewLabelWithStyle("消费金额:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				resultText,
			),
		),
	)
}

// NewReverseCalcCard 创建反向计算卡片
func NewReverseCalcCard(state *AppState) fyne.CanvasObject {
	costEntry := widget.NewEntry()
	costEntry.SetPlaceHolder("例如: 10")

	resultText := canvas.NewText("0", color.RGBA{R: 255, G: 200, B: 100, A: 255})
	resultText.TextSize = 20
	resultText.TextStyle = fyne.TextStyle{Bold: true}

	updateResult := func() {
		cost, err := strconv.ParseFloat(costEntry.Text, 64)
		if err != nil || cost < 0 {
			resultText.Text = "0"
			resultText.Refresh()
			return
		}
		price := state.GetPrice()
		if price == 0 {
			resultText.Text = "0"
			resultText.Refresh()
			return
		}
		tokens := cost / (price / 1000000)
		resultText.Text = FormatTokens(tokens)
		resultText.Refresh()
	}

	costEntry.OnChanged = func(text string) {
		updateResult()
	}

	state.Subscribe(func(price float64) {
		updateResult()
	})

	return widget.NewCard("🔄 反向计算", "",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widget.NewLabel("消费金额 (元)"),
				costEntry,
			),
			widget.NewSeparator(),
			container.NewHBox(
				widget.NewLabelWithStyle("Tokens 数量:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				resultText,
			),
		),
	)
}

// NewPriceCalcCard 创建单价计算卡片
func NewPriceCalcCard(state *AppState) fyne.CanvasObject {
	costEntry := widget.NewEntry()
	costEntry.SetPlaceHolder("例如: 10")

	tokensEntry := widget.NewEntry()
	tokensEntry.SetPlaceHolder("例如: 1000000")

	resultText := canvas.NewText("0 元/M", color.RGBA{R: 200, G: 150, B: 255, A: 255})
	resultText.TextSize = 20
	resultText.TextStyle = fyne.TextStyle{Bold: true}

	updateResult := func() {
		cost, err1 := strconv.ParseFloat(costEntry.Text, 64)
		tokens, err2 := strconv.ParseFloat(tokensEntry.Text, 64)

		if err1 != nil || err2 != nil || cost < 0 || tokens <= 0 {
			resultText.Text = "0 元/M"
			resultText.Refresh()
			return
		}

		price := cost / tokens * 1000000
		resultText.Text = FormatNumber(price) + " 元/M"
		resultText.Refresh()
	}

	costEntry.OnChanged = func(text string) {
		updateResult()
	}

	tokensEntry.OnChanged = func(text string) {
		updateResult()
	}

	return widget.NewCard("🧮 单价计算", "",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widget.NewLabel("消费金额 (元)"),
				costEntry,
			),
			container.NewGridWithColumns(2,
				widget.NewLabel("Tokens 数量"),
				tokensEntry,
			),
			widget.NewSeparator(),
			container.NewHBox(
				widget.NewLabelWithStyle("单价:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				resultText,
			),
		),
	)
}
