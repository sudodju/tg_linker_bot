package main

import (
	"flag"
	"log"

	"main.go/telegram_bot_yt_lesson/clients/telegram"
)

const (
	tgBotHost = "api.telegram.org"
)

func main() {
	// передавать токен будем с помощью флага запуска программы
	// token = flags.Get(token)

	tgClient = telegram.New(tgBotHost, mustToken())

	//fetcher = fetcher.New() - отпправляет запросы к api tg и получает события
	// fetcher и processor - общаются с api tg
	//processor = processor.New() - после обработки сам отправляет нам сообщения

	//запуск consumer.Start(fetcher, processor) - получает события и обрабатывает их
	// для получения использует fethcer, а для обработки processor
}

// токен - строка, если пустой то ошибка
// принудительно завершить программу
func mustToken() string {
	// во время запуска указываем в виде: bot -tg-bot-token 'my token'
	// здесь будет лежать не значение, а ссылка на значение
	token := flag.String("token-bot-token", "", "token for access to tg bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token in not specified")
	}

	return *token
}
