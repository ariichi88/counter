package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type japaneseTheme struct{}

func (*japaneseTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	if s.Bold {
		if s.Italic {
			return theme.DefaultTheme().Font(s)
		}
		return resourceMplus1cBoldTtf
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	return resourceMplus1cRegularTtf
}

func (*japaneseTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*japaneseTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*japaneseTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
