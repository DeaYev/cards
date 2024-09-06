package main

import (
	"image/jpeg"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Image Example")

	// Открываем файл изображения
	file, err := os.Open("3.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Декодируем изображение
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	// Создаем canvas.Image из декодированного изображения
	imgCanvas := canvas.NewImageFromImage(img)
	imgCanvas.FillMode = canvas.ImageFillOriginal
	imgCanvas.Resize(fyne.NewSize(200, 200)) // Устанавливаем размер изображения

	// Создаем контейнер без раскладки
	container := container.NewWithoutLayout(imgCanvas)

	// Устанавливаем позицию изображения (например, смещаем на 50 пикселей вправо и 30 пикселей вниз)
	imgCanvas.Move(fyne.NewPos(50, 30))

	myWindow.SetContent(container)
	myWindow.Resize(fyne.NewSize(400, 400)) // Размер окна
	myWindow.ShowAndRun()
}
