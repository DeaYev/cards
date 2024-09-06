package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func main() {
	myApp := app.New()

	// Создаем изображения
	image1 := canvas.NewImageFromFile("01.png")
	image2 := canvas.NewImageFromFile("05.png")
	image3 := canvas.NewImageFromFile("03.png")

	// Устанавливаем размеры изображений
	image1.SetMinSize(theme.IconInlineSize(), theme.IconInlineSize())
	image2.SetMinSize(theme.IconInlineSize(), theme.IconInlineSize())
	image3.SetMinSize(theme.IconInlineSize(), theme.IconInlineSize())

	// Создаем контейнер для изображений
	imagesContainer := container.NewGridWrap(container.NewMax(image1, image2, image3))

	// Создаем окно приложения
	myWindow := myApp.NewWindow("Multiple Images Example")
	myWindow.SetContent(imagesContainer)

	// Показываем окно и запускаем приложение
	myWindow.ShowAndRun()
}
