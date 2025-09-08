package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Product struct {
	ID    int
	Name  string
	Price float64
}

var (
	movies_by_genre = map[string][]string{
		"комедия":    {"иван васильевич меняет профессию", "операция ы", "бриллиантовая рука"},
		"драма":      {"побег из шоушенка", "зеленая миля", "шрек 2, но шрек 3 не смотри, лучше на второй остановиться."},
		"фантастика": {"матрица", "начало", "интерстеллар"},
		"боевик":     {"терминатор", "крепкий орешек", "джон уик"},
	}

	movies_by_mood = map[string][]string{
		"веселое":   {"мальчишник в вегасе", "одноклассники", "трудности перевода"},
		"грустное":  {"хатико", "зеленая миля", "достучаться до небес"},
		"романтическое": {"титаник", "записная книжка", "в джазе только девушки"},
	}

	music_by_genre = map[string][]string{
		"рок":       {"би 2 - варвара", "кино - хочу перемен", "ддт - что такое осень"},
		"поп":       {"anna asti - по барам", "мот - август это ты", "anna asti - верю в тебя"},
		"хип-хоп":   {"eminem - without me", "21 savage - a lot", "макс корж - вспоминай меня"},
	}

	music_by_mood = map[string][]string{
		"веселое":   {"the limba, morgenshtern - известным", "jony - love your voice", "instasamka - balance"},
		"грустное":  {"morgenshtern - красный флаг", "morgenshtern - повод", "morgenshtern - пустой вокзал"},
		"романтическое": {"скриптонит, charusha - космос", "chase atlantic - swim", "feduk - хлопья летят на верх"},
	}

	jokes = []string{
		"колобок повесился",
		"я мог бы рассказать анекдот про шагающую трость, но так мы далеко не уйдём",
		"может попробуешь другую команду?",
		"не, серьёзно. давай лучше музыку тебе подберём? напиши 'музыка жанр веселое'",
		"ладно.",
	}

	joke_index = 0

	merch = []Product{
		{ID: 1, Name: "футболка с лого", Price: 999.99},
		{ID: 2, Name: "кружка", Price: 1499.50},
		{ID: 3, Name: "наклейки", Price: 799.99},
		{ID: 4, Name: "толстовка", Price: 649.99},
		{ID: 5, Name: "батарейка", Price: 60.0},
		{ID: 6, Name: "фломастер с лого", Price: 60.99},
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	fmt.Println("привет! я твой дружелюбный помощник чатбот - нибери")
	fmt.Println("что хочешь?")
	fmt.Println("- фильм жанр [комедия, драма, фантастика, боевик]")
	fmt.Println("- фильм настроение [веселое, грустное, романтическое]") 
	fmt.Println("- музыка жанр [рок, поп, хип-хоп]")
	fmt.Println("- музыка настроение [веселое, грустное, романтическое]")
	fmt.Println("- анекдот (БОЛЬШАЯ РЕКОМЕНДАЦИЯ!)")
	fmt.Println("- мерч")
	fmt.Println("- выход")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if joke_index >= len(jokes) {break}

		fmt.Print("\nтвой выбор: ")
		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if input == "выход" {
			fmt.Println("пока! 👣 :(")
			break
		}

		handleCommand(input)
	}
}

func handleCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "фильм":
		if len(parts) < 3 {
			fmt.Println("нужно: фильм жанр [комедия, драма, фантастика, боевик] или фильм настроение [веселое, грустное, романтическое]")
			return
		}
		if parts[1] == "жанр" {
			suggestMovie("жанр", strings.Join(parts[2:], " "), movies_by_genre)
		} else if parts[1] == "настроение" {
			suggestMovie("настроение", strings.Join(parts[2:], " "), movies_by_mood)
		}

	case "музыка":
		if len(parts) < 3 {
			fmt.Println("нужно: музыка жанр [рок, поп, хип-хоп] или музыка настроение [веселое, грустное, романтическое]")
			return
		}
		if parts[1] == "жанр" {
			suggestMusic("жанр", strings.Join(parts[2:], " "), music_by_genre)
		} else if parts[1] == "настроение" {
			suggestMusic("настроение", strings.Join(parts[2:], " "), music_by_mood)
		}

	case "анекдот":

		fmt.Println("😄", jokes[joke_index])
		joke_index += 1

	case "мерч":
		fmt.Println("вот что есть:")
		for _, item := range merch {
			fmt.Printf("%d. %s - %.2f руб\n", item.ID, item.Name, item.Price)
		}

		fmt.Println("чтобы купить, используй команду - купить [номер товара]")

	case "купить":
		if len(parts) < 2 {
			fmt.Println("укажи номер товара")
			return
		}
		var id int
		fmt.Sscanf(parts[1], "%d", &id)
		buyItem(id)

	default:
		fmt.Println("не понял, попробуй еще раз")
	}
}

func suggestMovie(category, value string, source map[string][]string) {
	movies, ok := source[value]
	if !ok {
		fmt.Printf("не знаю таких %s %s\n", category, value)
		return
	}
	fmt.Printf("посмотри %s\n", movies[rand.Intn(len(movies))])
}

func suggestMusic(category, value string, source map[string][]string) {
	tracks, ok := source[value]
	if !ok {
		fmt.Printf("не знаю таких %s %s\n", category, value)
		return
	}
	fmt.Printf("послушай %s\n", tracks[rand.Intn(len(tracks))])
}

func buyItem(id int) {
	for _, item := range merch {
		if item.ID == id {
			fmt.Printf("купили %s за %.2f руб\n", item.Name, item.Price)
			fmt.Println("спасибо за покупку! 👍👍👍😍💰💋")
			return
		}
	}
	fmt.Printf("нет товара с номером %d\n", id)
}