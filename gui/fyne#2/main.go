package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("аперет калькулятора")
	w.Resize(fyne.NewSize(400, 320))

	label1 := widget.NewLabel("Введите первое число")
	num1 := widget.NewEntry()

	label2 := widget.NewLabel("Введите первое число")
	num2 := widget.NewEntry()

	answer := widget.NewLabel("")

	btn := widget.NewButton("Считаю", func() {
		n1, err1 := strconv.ParseFloat(num1.Text, 64)
		n2, err2 := strconv.ParseFloat(num2.Text, 64)

		if err1 != nil || err2 != nil {
			answer.SetText("Error")
		} else {
			sum := n1 + n2
			sub := n1 - n2
			mul := n1 * n2
			div := n1 / n2

			answer.SetText(fmt.Sprintf("(+) %f\n (-) %f\n(*) %f\n(/) %f", sum, sub, mul, div))
		}
	})

	w.SetContent(container.NewVBox(
		label1,
		num1,
		label2,
		num2,
		btn,
		answer,
	))

	w.ShowAndRun()
}
