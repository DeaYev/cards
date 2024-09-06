package main

//===========================================================
//   Вывод карты с частичным наложением  обрезанной карты   - работает
//===============================================================

import (
	"image"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	// Открываем основное изображение
	baseImageFile, err := os.Open("1.png")
	if err != nil {
		panic(err)
	}
	defer baseImageFile.Close()
	baseImage, err := png.Decode(baseImageFile)
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

	// Создаем новое изображение, куда будем рисовать
	outputImage := image.NewRGBA(baseImage.Bounds())

	// Копируем основное изображение на выходное изображение
	draw.Draw(outputImage, baseImage.Bounds(), baseImage, image.Point{0, 0}, draw.Src)

	// Определяем точку, в которой будет размещено накладываемое изображение
	offset := image.Point{X: 50, Y: 50}

	// Накладываем изображение поверх основного
	draw.Draw(outputImage, overlayImage.Bounds().Add(offset), overlayImage, image.Point{0, 0}, draw.Over)

	// Сохраняем результат в файл
	outputFile, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
	err = png.Encode(outputFile, outputImage)
	if err != nil {
		panic(err)
	}
}
