package main

//=======================
//  Вывод желтого окна с надписью и карты в левом верхнем углу
//    нужного размера
//==========================
import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {

	var image [10]*canvas.Image
	//var images []*canvas.Image
	var images []fyne.CanvasObject
	//var k string
	var f float32
	//  var images []fyne.CanvasObject   второй вариант . Нужно попробывать
	// Создаем новое приложение
	myApp := app.New()
	myWindow := myApp.NewWindow("Раздача карт ")

	// Создаем прямоугольник для задания цвета фона
	background := canvas.NewRectangle(color.RGBA{105, 255, 0, 255}) // Жёлтый цвет

	//    Ввод изображения без canvas
	/*imgIn, _ := os.Open("3.jpg")
	imgJpg, _ := jpeg.Decode(imgIn)
	imgIn.Close()*/
	f = 0
	for i := 0; i < 10; i++ {
		k := i + 1
		numFile := strconv.Itoa(k)
		fileName := numFile + ".png"

		image[i] = canvas.NewImageFromFile(fileName)
		image[i].Resize(fyne.NewSize(100, 150))
		f = f + 30
		image[i].Move(fyne.NewPos(30, f))

		images = append(images, image[i])
	}

	stack := container.NewWithoutLayout(images...)

	// Создаем контейнер с фоном и текстом
	content := container.NewStack(background, stack)

	// Добавляем контейнер на окно и отображаем окно
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(600, 800))
	myWindow.ShowAndRun()
}
