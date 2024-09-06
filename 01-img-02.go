package main

//===========================================================
//   Наложение картвы на карту не правильно создает фон и обрезает карты
//=====================================================================

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	// Открываем основное изображение
	cardImageFile, err := os.Open("1.png")
	if err != nil {
		panic(err)
	}
	defer cardImageFile.Close()
	cardImage, err := png.Decode(cardImageFile)
	if err != nil {
		panic(err)
	}

	// Открываем накладываемое изображение
	overlayImageFile, err := os.Open("20.png")
	if err != nil {
		panic(err)
	}
	defer overlayImageFile.Close()
	overlayImage, err := png.Decode(overlayImageFile)
	if err != nil {
		panic(err)
	}

	// Создаем новое изображение, куда будем рисовать и изменяем фон

	bounds := cardImage.Bounds()
	fmt.Printf("Границы изображения: %v\n", bounds) //--------------------------------
	fmt.Printf("Размеры изображения: ширина = %d, высота = %d\n", bounds.Dx(), bounds.Dy())

	outputImage := image.NewRGBA(bounds)

	// Копируем основное изображение на выходное изображение
	draw.Draw(outputImage, cardImage.Bounds(), cardImage, image.Point{30, 30}, draw.Src)

	// Определяем точку, в которой будет размещено накладываемое изображение
	offset := image.Point{X: 0, Y: 60}

	// Накладываем изображение поверх основного
	draw.Draw(outputImage, overlayImage.Bounds().Add(offset), overlayImage, image.Point{10, 10}, draw.Over)

	//  Сохраняем результат в файл
	outputFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	err = png.Encode(outputFile, outputImage)
	if err != nil {
		panic(err)
	}

	showNewWindow1(outputImage)

}

func showNewWindow1(outputImage image.Image) {
	myApp := app.New()

	myWindow := myApp.NewWindow("Вы выбрали карту")

	background1 := canvas.NewRectangle(color.RGBA{45, 125, 0, 155}) // Зеленый цвет

	//imageCanvas := canvas.NewImageFromImage(outputImage)
	//imageCanvas.FillMode = canvas.ImageFillOriginal

	outputImageCanvas := canvas.NewImageFromImage(outputImage)
	outputImageCanvas.FillMode = canvas.ImageFillOriginal

	//background := canvas.NewRectangle(color.RGBA{25, 255, 0, 255}) // Жёлтый цвет

	content := container.NewStack(background1, outputImageCanvas /*imageCanvas*/)

	// Добавляем изображение на контейнер и отображаем окно

	myWindow.Resize(fyne.NewSize(300, 200))
	background1.Resize(fyne.NewSize(300, 200))
	outputImageCanvas.Resize(outputImageCanvas.MinSize())
	outputImageCanvas.Move(fyne.NewPos(
		(myWindow.Canvas().Size().Width-outputImageCanvas.MinSize().Width)/2,
		(myWindow.Canvas().Size().Height-outputImageCanvas.MinSize().Height)/2,
	))

	myWindow.SetContent(container.NewCenter(content))
	myWindow.ShowAndRun()

	myWindow.Show()

}
