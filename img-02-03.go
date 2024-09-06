package main

//=======================
//  Вывод желтого окна с надписью и карты в левом верхнем углу
//    нужного размера
//==========================
import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {

	var image [3]*canvas.Image
	// Создаем новое приложение
	myApp := app.New()
	myWindow := myApp.NewWindow("Window with Background Color")

	// Создаем прямоугольник для задания цвета фона
	background := canvas.NewRectangle(color.RGBA{255, 255, 0, 255}) // Жёлтый цвет

	//    Ввод изображения без canvas
	/*imgIn, _ := os.Open("3.jpg")
	imgJpg, _ := jpeg.Decode(imgIn)
	imgIn.Close()*/

	// Создаем canvas.Image
	image[0] = canvas.NewImageFromFile("1.png")
	image[1] = canvas.NewImageFromFile("2.png")
	image[2] = canvas.NewImageFromFile("3.png")

	//imgCanvas.FillMode = canvas.ImageFillOriginal

	// Устанавливаем желаемый размер для изображения

	image[0].Resize(fyne.NewSize(100, 150)) // Устанавливаем размер изображения
	image[1].Resize(fyne.NewSize(100, 150))
	image[2].Resize(fyne.NewSize(100, 150))

	// Устанавливаем позицию изображения (например, смещаем на 50 пикселей вправо и 30 пикселей вниз)
	image[0].Move(fyne.NewPos(50, 30))
	image[1].Move(fyne.NewPos(50, 80))
	image[2].Move(fyne.NewPos(50, 130))

	// Создаем контейнер Stack с изображением

	stack := container.NewWithoutLayout(image[0], image[1], image[2])

	// Создаем контейнер с фоном и текстом
	content := container.NewStack(background, stack)

	// Добавляем контейнер на окно и отображаем окно
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(600, 800))
	myWindow.ShowAndRun()
}
