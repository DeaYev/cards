package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	suits = []string{"Пика", "Трефа", "Бубна", "Черва"}
	cards = []string{"Туз", "Король", "Дама", "Валет", "10", "9", "8", "7"}
)

type Card struct {
	Suit   string
	Value  string
	Number int
}

func main() {

	deck := getDeck()

	suitList := make(map[string]int, len(cards))
	suitName := ""

	myApp := app.New()
	myWindow := myApp.NewWindow("Выбор карты")

	background := canvas.NewRectangle(color.RGBA{105, 255, 0, 255}) // Жёлтый цвет

	// Создаем группу радиокнопок
	radioGroup := widget.NewRadioGroup(suits, func(check string) {
		suitList = deck[check]
		suitName = check

	})
	radioGroup.Resize((fyne.NewSize(600, 850))) //---------

	CloseButton := widget.NewButton("Выберете масть и закройте окно", func() {

		myWindow.Close() // Закрываем новое окно
		openNewWindow1(myApp, suitName, suitList)
	})

	radioGroup1 := container.NewStack(background,
		widget.NewLabel("Выберите карту :"),
		radioGroup,
		CloseButton,
	)

	// Создаем контейнер для отображения компонентов

	myWindow.SetContent(radioGroup1)

	myWindow.ShowAndRun()
	myWindow.Close()

}

// ==============================================================================
// Функция для отображения второго окна
func openNewWindow1(myApp fyne.App, suitName string, suitList map[string]int) {
	var cardID int
	var cardName string

	myWindow1 := myApp.NewWindow("Выбор карты")

	Group1 := widget.NewRadioGroup(cards, func(check string) {
		cardID = suitList[check]
		cardName = fmt.Sprintf("%s %s", suitName, cardName)

	})

	CloseButt := widget.NewButton("Выберете карту и закройте окно", func() {
		myWindow1.Close() // Закрываем новое окно
		fileName := ""    //"f://Go//GoProg//2024//0102//"
		openNewWindow2(myApp, cardName, cardID, fileName)
	})

	Group2 := container.NewHBox(
		widget.NewLabel("Выберите карту :"),
		Group1,
		CloseButt,
	)

	// Устанавливаем содержимое окна

	myWindow1.SetContent(Group2)

	myWindow1.Show()

}

//====================================================================================

func openNewWindow2(myApp fyne.App, cardName string, cardID int, fileNm string) {
	myWindow1 := myApp.NewWindow("Вы выбрали карту")

	fileName := fmt.Sprintf("%d.png", cardID)

	//fileName := fmt.Sprintf(%sfileName, fileNm)
	fileName = fileNm + fileName

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Fatalf("Файл %s не существует", fileName)
	}

	myWindow1.Resize(fyne.NewSize(500, 500))
	img := canvas.NewImageFromFile(fileName)
	img.FillMode = canvas.ImageFillContain
	CloseButt1 := widget.NewButton("Хотите продолжить выбор?", func() {
		//myWindow1.Close()
	})
	CloseButt2 := widget.NewButton("Или   закройте окно", func() {
		newImage := canvas.NewImageFromFile(fileName)
		newImage.FillMode = canvas.ImageFillContain
		myWindow1.SetContent(container.NewVBox(
			newImage,
			/* reloadButton,   */

		))
		myWindow1.Close()
	})
	CloseButt3 := widget.NewButton("Или раздать карты игрокам", func() {
		Sort(myApp)

		myWindow1.Close() // Закрываем новое окно

	})

	Group2 := container.NewAdaptiveGrid(2,
		widget.NewLabel("Вы выбрали карту :"),
		widget.NewLabel(cardName),
		img,
		CloseButt1,
		CloseButt2,
		CloseButt3,
	)

	myWindow1.SetContent(Group2)
	myWindow1.Show()

}

func getDeck() map[string]map[string]int {
	index := 1
	deck := make(map[string]map[string]int)

	for _, s := range suits {
		suit := make(map[string]int)
		deck[s] = suit

		for _, c := range cards {
			deck[s][c] = index
			index++

		}
	}

	return deck
}

func Sort(myApp fyne.App) {
	fmt.Println("Раздача")

	// Создание колоды карт
	deck1 := createDeck()
	var mas [32]Card
	var p1 [10]Card
	var p2 [10]Card
	var p3 [10]Card
	var p4 [2]Card
	var i, j, n int32

	// Тасование колоды
	shuffledDeck := shuffleDeck(deck1)

	//  Раздача карт в преферансе
	i = 0
	for i < 32 {
		mas[i] = shuffledDeck[i]
		i++

	}

	j = 0
	i = 0
	for i < 2 {
		p1[i] = mas[j]

		i++
		j++

	}
	i = 0
	for i < 2 {
		p2[i] = mas[j]

		i++
		j++
	}
	i = 0
	for i < 2 {
		p3[i] = mas[j]

		i++
		j++
	}
	i = 0
	for i < 2 {
		p4[i] = mas[j]

		i++
		j++
	}

	for n < 8 {

		n = n + i

		i = 0
		for i < 2 {
			p1[n+i] = mas[j]

			i++
			j++

		}
		i = 0
		for i < 2 {
			p2[n+i] = mas[j]

			i++
			j++

		}
		i = 0
		for i < 2 {
			p3[n+i] = mas[j]
			i++
			j++
		}

	}

	//   Сортировка карт у игроков
	rt1 := srt(p1[:])

	rt2 := srt(p2[:])

	rt3 := srt(p3[:])

	var nrt1 [10]int
	var nrt2 [10]int
	var nrt3 [10]int
	var nrt4 [2]int

	for t1 := 0; t1 < 10; t1++ {
		nrt1[t1] = rt1[t1].Number
		nrt2[t1] = rt2[t1].Number
		nrt3[t1] = rt3[t1].Number

	}
	nrt4[0] = p4[0].Number
	nrt4[1] = p4[1].Number
	//--------------------------------------------------------------

	wind_Card(myApp, nrt1[:], nrt2[:], nrt3[:], nrt4[:])
} //----------------------------------------

// =============================================================================
func wind_Card(myApp fyne.App, rf1 []int, rf2 []int, rf3 []int, rf4 []int) {

	var imageL [10]*canvas.Image
	var imageP [10]*canvas.Image
	var imageI [10]*canvas.Image
	var imagePr [2]*canvas.Image

	//var images []*canvas.Image
	var imagesL []fyne.CanvasObject
	var imagesP []fyne.CanvasObject
	var imagesI []fyne.CanvasObject
	var imagesPr []fyne.CanvasObject

	var k float32
	var f float32
	fmt.Println(rf1, rf2, rf3, rf4) //---------------------------------------------------------

	//-------------------------------------------------
	// Сохраняем результат в файл
	file, err := os.Create("numbers.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, v := range rf1 {
		fmt.Fprintln(file, v)
	}
	for _, v := range rf2 {
		fmt.Fprintln(file, v)
	}
	for _, v := range rf3 {
		fmt.Fprintln(file, v)
	}
	for _, v := range rf4 {
		fmt.Fprintln(file, v)
	}
	//-----------------------------------------------------

	//  var images []fyne.CanvasObject   второй вариант . Нужно попробывать

	myWindow3 := myApp.NewWindow("Раздача карт ")

	// Создаем прямоугольник для задания цвета фона
	background := canvas.NewRectangle(color.RGBA{80, 150, 0, 255}) // зеленый цвет

	k = 40
	f = 0
	for i := 0; i < 10; i++ {

		f = f + k
		numFile := strconv.Itoa(rf1[i])
		fileName := numFile + ".png"

		imageL[i] = canvas.NewImageFromFile(fileName)
		imageL[i].Resize(fyne.NewSize(100, 150))
		imageL[i].Move(fyne.NewPos(30, f))

		imagesL = append(imagesL, imageL[i])
	}

	f = 0
	for i := 0; i < 10; i++ {

		f = f + k
		numFile := strconv.Itoa(rf2[i])
		fileName := numFile + ".png"

		imageP[i] = canvas.NewImageFromFile(fileName)
		imageP[i].Resize(fyne.NewSize(100, 150))
		imageP[i].Move(fyne.NewPos(200, f))
		imagesP = append(imagesP, imageP[i])

	}
	f = 0
	for i := 0; i < 10; i++ {

		f = f + k
		numFile := strconv.Itoa(rf3[i])
		fileName := numFile + ".png"

		imageI[i] = canvas.NewImageFromFile(fileName)
		imageI[i].Resize(fyne.NewSize(100, 150))
		imageI[i].Move(fyne.NewPos(370, f))
		imagesI = append(imagesI, imageI[i])
	}

	f = 0
	for i := 0; i < 2; i++ {

		f = f + k
		numFile := strconv.Itoa(rf4[i])
		fileName := numFile + ".png"

		imagePr[i] = canvas.NewImageFromFile(fileName)
		imagePr[i].Resize(fyne.NewSize(100, 150))
		imagePr[i].Move(fyne.NewPos(540, f))
		imagesPr = append(imagesPr, imagePr[i])

	}
	fmt.Println("f= ", f) //--------------------------------------------------

	stack1 := container.NewWithoutLayout(imagesL...)
	stack2 := container.NewWithoutLayout(imagesP...)
	stack3 := container.NewWithoutLayout(imagesI...)
	stack4 := container.NewWithoutLayout(imagesPr...)

	// Создаем контейнер с фоном и текстом
	content := container.NewStack(background, stack1, stack2, stack3, stack4)

	// Добавляем контейнер на окно и отображаем окно
	myWindow3.SetContent(content)
	myWindow3.Resize(fyne.NewSize(800, 700))
	//myWindow3.ShowAndRun()
	myWindow3.Show()
}

// ================================================================================
// Функция  сортировки  карт у игроков
func srt(p []Card) []Card {

	var dt Card
	var d, d1 int
	var i, j int32

	i = 0
	j = 0

	for i < 9 {
		j = i + 1

		for j < 10 {
			fCard := p[i]
			d = fCard.Number

			f1Card := p[j]
			d1 = f1Card.Number

			if d1 < d {
				dt = p[j]

				p[j] = p[i]
				p[i] = dt

			}
			j++
		}

		i++

	}

	return p

}

// Функция для создания новой колоды карт
func createDeck() []Card {
	var deck1 []Card

	suits := []string{"Пики", "Бубна", "Трефы", "Червы"}
	values := []string{"A", "K", "Q", "J", "10", "9", "8", "7"}
	number := 0

	for _, suit := range suits {
		for _, value := range values {
			number = number + 1
			card := Card{Suit: suit, Value: value, Number: number}
			deck1 = append(deck1, card)

		}
	}
	return deck1
}

// Функция для тасования колоды карт
func shuffleDeck(deck1 []Card) []Card {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел с использованием текущего времени

	for i := range deck1 {
		// Генерация случайного индекса для обмена местами текущей карты и карты на случайной позиции
		j := rand.Intn(len(deck1))
		deck1[i], deck1[j] = deck1[j], deck1[i]

	}

	return deck1
}
