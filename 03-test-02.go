package main

import (
	"image"
	"image/jpeg"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
)

type DraggableImage struct {
	*canvas.Image
	dragging         bool
	offsetX, offsetY float32
}

func NewDraggableImage(img image.Image) *DraggableImage {
	imgCanvas := canvas.NewImageFromImage(img)
	imgCanvas.FillMode = canvas.ImageFillOriginal
	imgCanvas.SetMinSize(fyne.NewSize(200, 200))

	return &DraggableImage{
		Image: imgCanvas,
	}
}

func (di *DraggableImage) Dragged(e *fyne.DragEvent) {
	if di.dragging {
		di.Move(fyne.NewPos(e.PointEvent.Position.X-di.offsetX, e.PointEvent.Position.Y-di.offsetY))
		di.Refresh()
	}
}

func (di *DraggableImage) DragEnd() {
	di.dragging = false
}

func (di *DraggableImage) MouseDown(e *desktop.MouseEvent) {
	di.dragging = true
	di.offsetX = e.Position.X - di.Position().X
	di.offsetY = e.Position.Y - di.Position().Y
}

func (di *DraggableImage) MouseUp(e *desktop.MouseEvent) {
	di.dragging = false
}

func (di *DraggableImage) Tapped(e *fyne.PointEvent) {}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Draggable Image Example")

	// Открываем файл изображения
	file, err := os.Open("2.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Декодируем изображение
	img, err := jpeg.Decode(file)
	if err != nil {
		panic(err)
	}

	draggableImage := NewDraggableImage(img)

	container := container.NewWithoutLayout(draggableImage)
	myWindow.SetContent(container)
	myWindow.Resize(fyne.NewSize(800, 600))

	draggableImage.Resize(fyne.NewSize(200, 200)) // Устанавливаем размер изображения
	draggableImage.Move(fyne.NewPos(50, 50))      // Устанавливаем начальную позицию изображения

	myWindow.ShowAndRun()
}
