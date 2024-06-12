package main

import (
	"fmt"

	"github.com/shopspring/decimal"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&japaneseTheme{})

	w := a.NewWindow("simplfunce counter")

	var (
		total   decimal.Decimal
		count   decimal.Decimal
		prev    decimal.Decimal
		point   bool = false
	)

	tv := widget.NewLabel("0")
	tc := widget.NewLabel("0")
	cv := widget.NewLabel("0")

	tv.Alignment = fyne.TextAlignTrailing
	tc.Alignment = fyne.TextAlignTrailing
	cv.Alignment = fyne.TextAlignTrailing

	f := func(c string) {
		switch c {
		case "0":
			if cv.Text != "0" {
				cv.SetText(cv.Text + c)
			} else {
				cv.SetText(c)
			}
		case ".":
			if !point {
				cv.SetText(cv.Text + c)
				point = true
			}
		default:
			if cv.Text == "0" {
				cv.SetText(c)
			} else {
				cv.SetText(cv.Text + c)
			}
		}
	}

	b1 := widget.NewButton("1", func() {f("1")})
	b2 := widget.NewButton("2", func() {f("2")})
	b3 := widget.NewButton("3", func() {f("3")})
	b4 := widget.NewButton("4", func() {f("4")})
	b5 := widget.NewButton("5", func() {f("5")})
	b6 := widget.NewButton("6", func() {f("6")})
	b7 := widget.NewButton("7", func() {f("7")})
	b8 := widget.NewButton("8", func() {f("8")})
	b9 := widget.NewButton("9", func() {f("9")})
	b0 := widget.NewButton("0", func() {f("0")})
	pb := widget.NewButton(".", func() {f(".")})

	ltv := widget.NewLabel("総合計")
	ltc := widget.NewLabel("カウント")
	lcv := widget.NewLabel("現在の値")

	w.SetContent(container.NewVBox(
		container.NewGridWithColumns(2,
			ltv, tv,
			ltc, tc,
			lcv, cv,
		),
	 	container.NewGridWithColumns(3,
			b7, b8, b9,
			b4, b5, b6,
			b1, b2, b3,
			b0, pb,
		),
		&widget.Button{
			Text:       "現在の値をクリア",
			Importance: widget.LowImportance,
			OnTapped:   func() {
				cv.SetText("0")
			},},
		&widget.Button{
			Text:       "現在の値を総合計に加算",
			Importance: widget.HighImportance,
			OnTapped:   func() {
				current, err := decimal.NewFromString(cv.Text)
				if err != nil {
					fmt.Println("数値の取得に失敗しました")
				}
				if !current.Equal(decimal.NewFromInt32(0)) {
					total = total.Add(current)
					count = count.Add(decimal.NewFromInt32(1))
					prev = current
					point = false
				}
				tv.SetText(total.String())
				tc.SetText(count.String())
				cv.SetText("0")
			},},
		&widget.Button{
			Text:       "直前の値を削除",
			Importance: widget.DangerImportance,
			OnTapped:   func() {
				if !prev.Equal(decimal.NewFromInt32(0)) {
					total = total.Sub(prev)
					count = count.Sub(decimal.NewFromInt32(1))
					prev = decimal.NewFromInt32(0)
				}
				tv.SetText(total.String())
				tc.SetText(count.String())
				cv.SetText("0")
			},},
		),
	)
	w.ShowAndRun()
}
