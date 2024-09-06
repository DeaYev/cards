package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func main() {
	// Открываем файл изображения
	file, err := os.Open("2.png")
	if err != nil {
		log.Fatalf("Не удалось открыть файл: %v", err)
	}
	defer file.Close()

	// Декодируем изображение
	srcImg, err := png.Decode(file)
	if err != nil {
		log.Fatalf("Не удалось декодировать изображение: %v", err)
	}

	// Получаем границы исходного изображения
	bounds := srcImg.Bounds()

	// Создаем новое изображение с такими же границами
	dstImg := image.NewRGBA(bounds)

	// Заливаем новое изображение белым цветом
	draw.Draw(dstImg, bounds, &image.Uniform{color.White}, image.Point{}, draw.Src)

	// Копируем исходное изображение на новое
	draw.Draw(dstImg, bounds, srcImg, bounds.Min, draw.Over)

	// Сохраняем новое изображение в файл
	outFile, err := os.Create("new_image.png")
	if err != nil {
		log.Fatalf("Не удалось создать файл: %v", err)
	}
	defer outFile.Close()

	err = jpeg.Encode(outFile, dstImg, nil)
	if err != nil {
		log.Fatalf("Не удалось сохранить изображение: %v", err)
	}

	log.Println("Изображение успешно сохранено.")
}
