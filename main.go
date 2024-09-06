package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Change Background and Display Image")

	// Изменение фона окна
	background := canvas.NewRectangle(color.RGBA{R: 145, G: 55, B: 55, A: 255})

	// Загрузка картинки
	image := canvas.NewImageFromFile( /*D:\\Go\\GoProg\\2024\\Наложение карт\\*/ "5.png")
	image.FillMode = canvas.ImageFillOriginal
	// Создание контейнера с фоном и картинкой
	content := container.NewStack(background, image)

	// Обновление размера картинки при изменении размера окна
	myWindow.Resize(fyne.NewSize(800, 600))
	image.Resize(image.MinSize())
	image.Move(fyne.NewPos(
		(myWindow.Canvas().Size().Width-image.MinSize().Width)/30,
		(myWindow.Canvas().Size().Height-image.MinSize().Height)/30,
	))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
