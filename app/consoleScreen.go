package uvs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Testing screen
func consoleScreen(uv *UV_Station) fyne.CanvasObject {
	m := 0

	textArea := widget.NewEntry()
	textArea.Keyboard()

	txt := container.NewVBox()

	form := &widget.Form{
		BaseWidget: widget.BaseWidget{},
		Items:      []*widget.FormItem{},

		OnSubmit: func() {

			m++
			if m > 4 {
				m = 0
				txt.RemoveAll()
			}

			text2 := canvas.NewText(" "+textArea.Text, theme.ForegroundColor())
			txt.Add(text2)
			textArea.SetText("")
			txt.Refresh()
		},
		SubmitText: uv.T.Send,
	}

	form.Append("", textArea)

	screen := container.NewBorder(nil, form, nil, nil, txt)
	return screen
}
