package main

import "fmt"

//Реализовать паттерн «адаптер» на любом примере.

//основной интерфейс, чтобы воспроизвести музыкальный файл
type player interface {
	playOGG()
}

//сервис который воспроизводит только ogg формат
type walkman struct {}
func (w *walkman)playOGG()  {
	fmt.Println("Воспроизвожу OGG файл..")
}

//есть сервис который воспроизводит только mp3 формат
type iRiver struct {}
func (i *iRiver)playMP3()  {
	fmt.Println("Воспроизвожу MP3 файл..")
}

//адаптер переводит ogg в mp3 формат
type musicAdapter struct {
	OGGtoMP3 *iRiver
}
func (m *musicAdapter)playOGG()  {
	fmt.Println("Адаптер переводит ogg в mp3 формат..")
	m.OGGtoMP3.playMP3()
}

//наш клиент который запустит ogg файл
type human struct {}
func (h *human) playOGG(play player)  {
	fmt.Println("Открываю файл OGG формате..")
	play.playOGG()
	}

func main()  {
	human:= &human{} //клиент который запустит ogg файл
	walkman := &walkman{} //сервис который воспроизводит ogg формат

	human.playOGG(walkman) //успешное открытие ogg файла

	OGGtoMP3 := &iRiver{} //сервис который воспроизводит mp3 формат
	mp3iRiverAdapter := &musicAdapter{ //адаптер конвертирует ogg в mp3 формат
		OGGtoMP3,
	}
human.playOGG(mp3iRiverAdapter) //успешное воспроизведение

}