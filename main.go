package main

import (
	"strconv"
	"fmt"

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
		total    float64
		count    float64
		current  float64
	)

	tv := widget.NewLabel("0")
	tc := widget.NewLabel("0")
	cv := widget.NewLabel("0")

	tv.Alignment = fyne.TextAlignTrailing
	tc.Alignment = fyne.TextAlignTrailing
	cv.Alignment = fyne.TextAlignTrailing

	f := func(c string) {
		if cv.Text == "0" {
			cv.SetText(c)
		} else {
			cv.SetText(cv.Text + c)
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
			Text:       "現在の値を総合計に加算",
			Importance: widget.HighImportance,
			OnTapped:   func() {
				current, err := strconv.ParseFloat(cv.Text, 64)
				if err != nil {
					fmt.Println("数値の変換に失敗しました")
				}
				total = total + current
				count ++
				tv.SetText(fmt.Sprint(total))
				tc.SetText(fmt.Sprint(count))
				cv.SetText("0")
			},},
		&widget.Button{
			Text:       "直前の値を削除",
			Importance: widget.DangerImportance,
			OnTapped:   func() {
				total = total - current
				count = count - 1
				tv.SetText(fmt.Sprint(total))
				tc.SetText(fmt.Sprint(count))
				cv.SetText("0")
			},},
		),
	)

	w.ShowAndRun()
}
