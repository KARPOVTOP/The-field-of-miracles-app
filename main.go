package main

import (
	"math/rand"
	"time"
)

type questions struct { //Структура для создания образа вопросов
	quest  string //Вопрос
	answer string // Ответ на вопрос
}

var hp, money int = 3, 0 // Жизни игрока и деньги
var isGame = true        // Идет ли игра или закончилась

func createQuest() []questions { //Создание среза с вопросами и ответами
	slice := []questions{
		{quest: "Что использовали в Китае для глажки белья вместо утюга?", answer: "сковорода"},
		{quest: "Как у западных и южных славян назывались селение, деревня, курень?", answer: "жупа"},
		{quest: "Польский ученый-математик Гуго Дионисий Штейнгауз, прославившийся также своими афоризмами, говорил: «Комплимент женщине должен быть правдивее, чем...»", answer: "правда"},
		{quest: "В Австралии на парковках возле некоторых торговых центров по ночам и вечерам включают классическую музыку, чтобы отпугнуть кого-то. Кого?", answer: "подросток"},
		{quest: "Пельмени издавна заготавливают в форме ушек. Что символизируют такие пельмени?", answer: "послушание"},
	}
	return slice
}

func generationQuest(slice []questions) (string, string) { // Генерация случайного вопроса и ответа
	i := rand.Intn(len(slice)) //Генерация индекса для выбора случайного элемента среза
	return slice[i].quest, slice[i].answer
}

func main() {
	rand.Seed(time.Now().Unix()) //Опора для генератора чисел

	//for isGame { //Бесконечный цикл ,пока идет игра
	quest, answer := generationQuest(createQuest())

	//}
}