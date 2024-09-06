package main

//=======================================================
//  Вывод зеленого фона  окна и раздача карт партненрам
//=======================================================
import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {

	var imageL [10]*canvas.Image
	var imageP [10]*canvas.Image
	var imageI [10]*canvas.Image

	//var images []*canvas.Image
	var imagesL []fyne.CanvasObject
	var imagesP []fyne.CanvasObject
	var imagesI []fyne.CanvasObject

	//var k string
	var f float32
	//  var images []fyne.CanvasObject   второй вариант . Нужно попробывать
	// Создаем новое приложение
	myApp := app.New()
	myWindow := myApp.NewWindow("Раздача карт ")

	// Создаем прямоугольник для задания цвета фона
	background := canvas.NewRectangle(color.RGBA{105, 255, 0, 255}) // Жёлтый цвет

	f = 0
	for i := 0; i < 10; i++ {
		k := i + 1
		f = f + 30
		numFile := strconv.Itoa(k)
		fileName := numFile + ".png"

		imageL[i] = canvas.NewImageFromFile(fileName)
		imageL[i].Resize(fyne.NewSize(100, 150))
		imageL[i].Move(fyne.NewPos(30, f))

		imagesL = append(imagesL, imageL[i])
	}

	f = 0
	for i := 0; i < 10; i++ {
		k := i + 11
		f = f + 30
		numFile := strconv.Itoa(k)
		fileName := numFile + ".png"

		imageP[i] = canvas.NewImageFromFile(fileName)
		imageP[i].Resize(fyne.NewSize(100, 150))
		imageP[i].Move(fyne.NewPos(200, f))
		imagesP = append(imagesP, imageP[i])

	}
	f = 0
	for i := 0; i < 10; i++ {
		k := i + 21
		f = f + 30
		numFile := strconv.Itoa(k)
		fileName := numFile + ".png"

		imageI[i] = canvas.NewImageFromFile(fileName)
		imageI[i].Resize(fyne.NewSize(100, 150))
		imageI[i].Move(fyne.NewPos(450, f))
		imagesI = append(imagesI, imageI[i])
	}

	stack1 := container.NewWithoutLayout(imagesL...)
	stack2 := container.NewWithoutLayout(imagesP...)
	stack3 := container.NewWithoutLayout(imagesI...)

	// Создаем контейнер с фоном и текстом
	content := container.NewStack(background, stack1, stack2, stack3)

	// Добавляем контейнер на окно и отображаем окно
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(600, 700))
	myWindow.ShowAndRun()
}
