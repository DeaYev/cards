package main

//=======================
//  Вывод желтого окна с надписью
//==========================
import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	// Создаем новое приложение
	myApp := app.New()
	myWindow := myApp.NewWindow("Window with Background Color")

	// Создаем прямоугольник для задания цвета фона
	background := canvas.NewRectangle(color.RGBA{255, 255, 0, 255}) // Жёлтый цвет

	// Создаем текстовый объект
	text := canvas.NewText("Hello, Fyne!", color.Black)

	// Создаем контейнер с фоном и текстом
	content := container.NewStack(background, text)

	// Добавляем контейнер на окно и отображаем окно
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
