package main

import (
	"fmt"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	arg := os.Args[1]
	bot, err := tgbotapi.NewBotAPI(arg) // create new bot
	if err != nil {
		fmt.Println(err)
	}
	//bot.Debug = true

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.InlineQuery == nil { // if no inline query, ignore it
			continue
		}

		texter := engru(update.InlineQuery.Query)
		textre := rueng(update.InlineQuery.Query)
		textter := translitenru(update.InlineQuery.Query)
		texttre := translitruen(update.InlineQuery.Query)
		user := update.InlineQuery.From.UserName
		article := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Swap keyboard en -> ru", texter)
		article.Description = texter
		article0 := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID+"_0", "Swap keyboard ru -> en", textre)
		article0.Description = textre
		article1 := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID+"_1", "Translit en -> ru", textter)
		article1.Description = textter
		article2 := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID+"_2", "Translit ru -> en", texttre)
		article2.Description = texttre
		article3 := tgbotapi.NewInlineQueryResultArticleHTML(update.InlineQuery.ID+"_3", "Think italic", "@"+user+" <i>"+update.InlineQuery.Query+"</i>")
		article3.Description = "@" + user + " <i>" + update.InlineQuery.Query + "</i>"
		article4 := tgbotapi.NewInlineQueryResultArticleHTML(update.InlineQuery.ID+"_4", "Think bold", "@"+user+" <b>"+update.InlineQuery.Query+"</b>")
		article4.Description = "@" + user + " <b>" + update.InlineQuery.Query + "</b>"
		fmt.Printf("User:%s Send %s Recode en->ru:%s Recode ru->en:%s Translit en->ru %s Translit ru->en %s Think italic %s Think bold %s\n", user, update.InlineQuery.Query, texter, textre, textter, texttre, "@"+user+" <i>"+update.InlineQuery.Query+"</i>", "@"+user+" <b>"+update.InlineQuery.Query+"</b>")

		var t []interface{}
		t = append(t, article, article0, article1, article2, article3, article4)
		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    false,
			CacheTime:     0,
			Results:       t,
		}

		if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
			fmt.Println(err)
		}
	}
}

func engru(s string) string {
	r := strings.NewReplacer(
		"Q", "Й",
		"W", "Ц",
		"E", "У",
		"R", "К",
		"T", "Е",
		"Y", "Н",
		"U", "Г",
		"I", "Ш",
		"O", "Щ",
		"P", "З",
		"[", "Х",
		"]", "Ъ",
		"A", "Ф",
		"S", "Ы",
		"D", "В",
		"F", "А",
		"G", "П",
		"H", "Р",
		"J", "О",
		"K", "Л",
		"L", "Д",
		";", "Ж",
		"\"", "Э",
		"Z", "Я",
		"X", "Ч",
		"C", "С",
		"V", "М",
		"B", "И",
		"N", "Т",
		"M", "Ь",
		"<", "Б",
		">", "Ю",
		"q", "й",
		"w", "ц",
		"e", "у",
		"r", "к",
		"t", "е",
		"y", "н",
		"u", "г",
		"i", "ш",
		"o", "щ",
		"p", "з",
		"[", "х",
		"]", "ъ",
		"a", "ф",
		"s", "ы",
		"d", "в",
		"f", "а",
		"g", "п",
		"h", "р",
		"j", "о",
		"k", "л",
		"l", "д",
		";", "ж",
		"'", "э",
		"z", "я",
		"x", "ч",
		"c", "с",
		"v", "м",
		"b", "и",
		"n", "т",
		"m", "ь",
		",", "б",
		".", "ю",
		"?", ",",
		"~", "Ё",
		"`", "ё",
		"^", ",")
	return r.Replace(s)
}

func rueng(s string) string {
	r := strings.NewReplacer(
		"Й", "Q",
		"Ц", "W",
		"У", "E",
		"К", "R",
		"Е", "T",
		"Н", "Y",
		"Г", "U",
		"Ш", "I",
		"Щ", "O",
		"З", "P",
		"Х", "[",
		"Ъ", "]",
		"Ф", "A",
		"Ы", "S",
		"В", "D",
		"А", "F",
		"П", "G",
		"Р", "H",
		"О", "J",
		"Л", "K",
		"Д", "L",
		"Ж", ";",
		"Э", "\"",
		"Я", "Z",
		"Ч", "X",
		"С", "C",
		"М", "V",
		"И", "B",
		"Т", "N",
		"Ь", "M",
		"Б", "<",
		"Ю", ">",
		"й", "q",
		"ц", "w",
		"у", "e",
		"к", "r",
		"е", "t",
		"н", "y",
		"г", "u",
		"ш", "i",
		"щ", "o",
		"з", "p",
		"х", "[",
		"ъ", "]",
		"ф", "a",
		"ы", "s",
		"в", "d",
		"а", "f",
		"п", "g",
		"р", "h",
		"о", "j",
		"л", "k",
		"д", "l",
		"ж", ";",
		"э", "'",
		"я", "z",
		"ч", "x",
		"с", "c",
		"м", "v",
		"и", "b",
		"т", "n",
		"ь", "m",
		"б", ",",
		"ю", ".",
		",", "?",
		"Ё", "~",
		"ё", "`",
		",", "^")
	return r.Replace(s)
}

func translitenru(s string) string {
	r := strings.NewReplacer(
		"SCH", "Щ",
		"ZH", "Ж",
		"TC", "Ц",
		"CH", "Ч",
		"SH", "Ш",
		"YE", "Э",
		"YU", "Ю",
		"YA", "Я",
		"A", "А",
		"B", "Б",
		"V", "В",
		"G", "Г",
		"D", "Д",
		"E", "Е",
		"~", "Ё",
		"Z", "З",
		"I", "И",
		"J", "Й",
		"K", "К",
		"L", "Л",
		"M", "М",
		"N", "Н",
		"O", "О",
		"P", "П",
		"R", "Р",
		"S", "С",
		"T", "Т",
		"U", "У",
		"F", "Ф",
		"H", "Х",
		"`", "Ъ",
		"'", "Ь",
		"Y", "Ы",

		"sch", "щ",
		"zh", "ж",
		"tc", "ц",
		"ch", "ч",
		"sh", "ш",
		"ye", "э",
		"yu", "ю",
		"ya", "я",

		"a", "а",
		"b", "б",
		"v", "в",
		"g", "г",
		"d", "д",
		"e", "е",
		"`", "ё",
		"z", "з",
		"i", "и",
		"j", "й",
		"k", "к",
		"l", "л",
		"m", "м",
		"n", "н",
		"o", "о",
		"p", "п",
		"r", "р",
		"s", "с",
		"t", "т",
		"u", "у",
		"f", "ф",
		"h", "х",
		"`", "ъ",
		"y", "ы",
		"'", "ь")
	return r.Replace(s)
}

func translitruen(s string) string {
	r := strings.NewReplacer(
		"Щ", "SCH",
		"Ж", "ZH",
		"Ц", "TC",
		"Ч", "CH",
		"Ш", "SH",
		"Э", "YE",
		"Ю", "YU",
		"Я", "YA",
		"А", "A",
		"Б", "B",
		"В", "V",
		"Г", "G",
		"Д", "D",
		"Е", "E",
		"Ё", "E",
		"З", "Z",
		"И", "I",
		"Й", "J",
		"К", "K",
		"Л", "L",
		"М", "M",
		"Н", "N",
		"О", "O",
		"П", "P",
		"Р", "R",
		"С", "S",
		"Т", "T",
		"У", "U",
		"Ф", "F",
		"Х", "H",
		"Ъ", "`",
		"Ь", "'",
		"Ы", "Y",

		"щ", "sch",
		"ж", "zh",
		"ц", "tc",
		"ч", "ch",
		"ш", "sh",
		"э", "ye",
		"ю", "yu",
		"я", "ya",

		"а", "a",
		"б", "b",
		"в", "v",
		"г", "g",
		"д", "d",
		"е", "e",
		"ё", "e",
		"з", "z",
		"и", "i",
		"й", "j",
		"к", "k",
		"л", "l",
		"м", "m",
		"н", "n",
		"о", "o",
		"п", "p",
		"р", "r",
		"с", "s",
		"т", "t",
		"у", "u",
		"ф", "f",
		"х", "h",
		"ъ", "`",
		"ы", "y",
		"ь", "'")
	return r.Replace(s)
}

//
