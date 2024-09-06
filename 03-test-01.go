/*
package main

import (
	"fmt"
	"image"
)

func main() {
	// Попробуем создать переменную типа image.Image
	var img image.Image

	// Проверим наличие метода ModeNRGB
	// Это сгенерирует ошибку компиляции, если метода не существует
	_ = img.ModeNRGB

	fmt.Println("Метод ModeNRGB существует")
}
*/

// ============================================================
/* package main

import (
	"golang.org/x/image/draw"
	"image"

	//"image/draw"
	"image/jpeg"
	"os"
)

func main() {
	// Откройте исходное изображение
	file, err := os.Open("2.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Декодируйте изображение
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// Определите новый размер
	newWidth := 50
	newHeight := 300
	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	// Измените размер изображения
	draw.CatmullRom.Scale(newImg, newImg.Bounds(), img, img.Bounds(), draw.Over, nil)

	// Откройте файл для записи
	outFile, err := os.Create("output.jpg")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Сохраните измененное изображение
	jpeg.Encode(outFile, newImg, nil)
}
*/
//=======================================================================================
/*
package main

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"os"
)

func main() {
	imgIn, _ := os.Open("3.jpg")
	imgJpg, _ := jpeg.Decode(imgIn)
	imgIn.Close()

	// Изменение размера картинки на ширину 600 пикселей (высота автоматически подстраивается)
	imgJpg = resize.Resize(100, 200, imgJpg, resize.Bicubic)

	imgOut, _ := os.Create("test-out.jpg")
	jpeg.Encode(imgOut, imgJpg, nil)
	imgOut.Close()
}    */

// =================================================================
package main

import (
	"fmt"
	"image/jpeg"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {

	var f float32
	f = 0
	myApp := app.New()
	myWindow := myApp.NewWindow("Image Example")

	fileNames := []string{"1.jpg", "2.jpg", "3.jpg"}

	var images []fyne.CanvasObject
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file:", err)
			continue
		}
		defer file.Close()

		img, err := jpeg.Decode(file)
		if err != nil {
			fmt.Println("Error decoding image:", err)
			continue
		}

		imgCanvas := canvas.NewImageFromImage(img)
		imgCanvas.FillMode = canvas.ImageFillOriginal
		//imgCanvas.SetMinSize(fyne.NewSize(200, 200))

		imgCanvas.Resize(fyne.NewSize(100, 150))
		f = f + 30
		imgCanvas.Move(fyne.NewPos(30, f))

		images = append(images, imgCanvas)
	}

	// Используем оператор ... для передачи среза как списка аргументов
	stack := container.NewWithoutLayout(images...)

	myWindow.SetContent(stack)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
