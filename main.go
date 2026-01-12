package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// 全局变量存储当前单价
var currentPricePerMillion float64 = 0

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())

	myWindow := myApp.NewWindow("Tokens 价格计算器")

	// ========== 标题 ==========
	title := canvas.NewText("Tokens 价格计算器", color.RGBA{R: 100, G: 200, B: 255, A: 255})
	title.TextSize = 24
	title.TextStyle = fyne.TextStyle{Bold: true}
	title.Alignment = fyne.TextAlignCenter

	subtitle := canvas.NewText("快速换算 Token 价格与消费", color.RGBA{R: 150, G: 150, B: 150, A: 255})
	subtitle.TextSize = 12
	subtitle.Alignment = fyne.TextAlignCenter

	// ========== 单价设置区域 ==========
	priceMillionEntry := widget.NewEntry()
	priceMillionEntry.SetPlaceHolder("例如: 10")

	priceThousandEntry := widget.NewEntry()
	priceThousandEntry.SetPlaceHolder("自动计算")

	var updating bool

	priceMillionEntry.OnChanged = func(text string) {
		if updating {
			return
		}
		val, err := strconv.ParseFloat(text, 64)
		if err != nil || val < 0 {
			return
		}
		currentPricePerMillion = val
		updating = true
		priceThousandEntry.SetText(formatNumber(MillionToThousand(val)))
		updating = false
	}

	priceThousandEntry.OnChanged = func(text string) {
		if updating {
			return
		}
		val, err := strconv.ParseFloat(text, 64)
		if err != nil || val < 0 {
			return
		}
		currentPricePerMillion = ThousandToMillion(val)
		updating = true
		priceMillionEntry.SetText(formatNumber(currentPricePerMillion))
		updating = false
	}

	priceCard := widget.NewCard("💰 单价设置", "",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widget.NewLabelWithStyle("元/百万tokens", fyne.TextAlignLeading, fyne.TextStyle{}),
				priceMillionEntry,
			),
			container.NewGridWithColumns(2,
				widget.NewLabelWithStyle("元/千tokens", fyne.TextAlignLeading, fyne.TextStyle{}),
				priceThousandEntry,
			),
		),
	)

	// ========== 消费计算区域 ==========
	tokensEntry := widget.NewEntry()
	tokensEntry.SetPlaceHolder("例如: 1000000")

	costResultText := canvas.NewText("¥ 0.000000", color.RGBA{R: 100, G: 255, B: 150, A: 255})
	costResultText.TextSize = 20
	costResultText.TextStyle = fyne.TextStyle{Bold: true}

	tokensEntry.OnChanged = func(text string) {
		tokens, err := strconv.ParseFloat(text, 64)
		if err != nil || tokens < 0 {
			costResultText.Text = "¥ 0.000000"
			costResultText.Refresh()
			return
		}
		cost := CalculateCost(tokens, currentPricePerMillion)
		costResultText.Text = fmt.Sprintf("¥ %.6f", cost)
		costResultText.Refresh()
	}

	costCard := widget.NewCard("📊 消费计算", "",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widget.NewLabelWithStyle("Tokens 数量", fyne.TextAlignLeading, fyne.TextStyle{}),
				tokensEntry,
			),
			widget.NewSeparator(),
			container.NewHBox(
				widget.NewLabelWithStyle("消费金额:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				costResultText,
			),
		),
	)

	// ========== 反向计算区域 ==========
	costEntry := widget.NewEntry()
	costEntry.SetPlaceHolder("例如: 10")

	tokensResultText := canvas.NewText("0", color.RGBA{R: 255, G: 200, B: 100, A: 255})
	tokensResultText.TextSize = 20
	tokensResultText.TextStyle = fyne.TextStyle{Bold: true}

	costEntry.OnChanged = func(text string) {
		cost, err := strconv.ParseFloat(text, 64)
		if err != nil || cost < 0 {
			tokensResultText.Text = "0"
			tokensResultText.Refresh()
			return
		}
		tokens := CalculateTokens(cost, currentPricePerMillion)
		tokensResultText.Text = formatTokens(tokens)
		tokensResultText.Refresh()
	}

	reverseCard := widget.NewCard("🔄 反向计算", "",
		container.NewVBox(
			container.NewGridWithColumns(2,
				widget.NewLabelWithStyle("消费金额 (元)", fyne.TextAlignLeading, fyne.TextStyle{}),
				costEntry,
			),
			widget.NewSeparator(),
			container.NewHBox(
				widget.NewLabelWithStyle("Tokens 数量:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
				tokensResultText,
			),
		),
	)

	// ========== 主布局 ==========
	content := container.NewVBox(
		container.NewPadded(container.NewVBox(
			title,
			subtitle,
		)),
		widget.NewSeparator(),
		container.NewPadded(priceCard),
		container.NewPadded(costCard),
		container.NewPadded(reverseCard),
	)

	scroll := container.NewScroll(content)

	myWindow.SetContent(scroll)
	myWindow.Resize(fyne.NewSize(420, 520))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}

// formatNumber 格式化数字，去除不必要的尾零
func formatNumber(val float64) string {
	s := fmt.Sprintf("%.10f", val)
	for len(s) > 1 && s[len(s)-1] == '0' && s[len(s)-2] != '.' {
		s = s[:len(s)-1]
	}
	if s[len(s)-1] == '.' {
		s = s + "0"
	}
	return s
}

// formatTokens 格式化tokens数量，大数字使用千分位
func formatTokens(val float64) string {
	if val >= 1000000 {
		return fmt.Sprintf("%.2fM", val/1000000)
	} else if val >= 1000 {
		return fmt.Sprintf("%.2fK", val/1000)
	}
	return fmt.Sprintf("%.2f", val)
}
