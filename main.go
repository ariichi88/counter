package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	a.Settings().SetTheme(&japaneseTheme{})

	w := a.NewWindow("simple counter")

	b1 := widget.NewButton("1", func() {})
	b2 := widget.NewButton("2", func() {})
	b3 := widget.NewButton("3", func() {})
	b4 := widget.NewButton("4", func() {})
	b5 := widget.NewButton("5", func() {})
	b6 := widget.NewButton("6", func() {})
	b7 := widget.NewButton("7", func() {})
	b8 := widget.NewButton("8", func() {})
	b9 := widget.NewButton("9", func() {})
	b0 := widget.NewButton("0", func() {})
	pb := widget.NewButton(".", func() {})

	tv := widget.NewLabel("0")
	tc := widget.NewLabel("0")
	cv := widget.NewLabel("0")

	tv.Alignment = fyne.TextAlignTrailing
	tc.Alignment = fyne.TextAlignTrailing
	cv.Alignment = fyne.TextAlignTrailing

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
			OnTapped:   func() {},},
		&widget.Button{
			Text:       "直前の値を削除",
			Importance: widget.DangerImportance,
			OnTapped:   func() {},},
		),
	)

	w.ShowAndRun()
}
