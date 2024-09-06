package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/theme"
	//"github.com/nfnt/resize"
)

func main() {
	myApp := app.New()

	// Создаем изображения
	image1 := canvas.NewImageFromFile("1.png")
	image2 := canvas.NewImageFromFile("2.png")
	image3 := canvas.NewImageFromFile("3.png")
	// Устанавливаем размеры изображений
	//image1 = resize.Resize(100, 200, image1, resize.Bicubic)

	// Устанавливаем размеры изображений
	/*image1.SetMinSize(theme.IconInlineSize(), theme.IconInlineSize())
	image2.SetMinSize(theme.IconInlineSize(), theme.IconInlineSize())
	image3.SetMinSize(theme.IconInlineSize(), theme.IconInlineSize())  */

	// Создаем контейнер для изображений
	imagesContainer := container.NewGridWrap(fyne.NewSize(400, 200), image1, image2, image3)

	// Создаем окно приложения
	myWindow := myApp.NewWindow("Multiple Images Example")
	myWindow.SetContent(imagesContainer)

	// Показываем окно и запускаем приложение
	myWindow.ShowAndRun()
}
