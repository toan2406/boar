package networklog

import (
	"github.com/rivo/tview"
)

func Render() {
	box := tview.NewBox().SetBorder(true).SetTitle("Boar Proxy")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
