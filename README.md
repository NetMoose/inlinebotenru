# inlinebotenru

Тестовый, но в то же время работающий inline-бот для телеграм, с неким более-менее полезным функционалом.

Бот, принимает строчку текста и делает следующие вещи:

1. Пытается исправить неправильно набранную строку (если не была переключена раскладка с английского и строка была набрана на русском)
2. То же самое, но с русской раскладки на английскую.
3. Транслитерация en -> ru
4. Транслитерация ru -> en
5. Подстановка ника пишущего и сама строка в полужирном начертании
6. Подстановка ника пишущего и сама строка в италике

Все это выдается в виде меню при вызове данного бота как *@swp_bot какаятострока* в любом чате с начала строки ввода. После чего можно выбрать нужный вариант из появляющегося меню и отправить его в этот чат.

Никаких других команд бот не принимает и ни на какие команды при непосредственном общении не отвечает (он чисто инлайновый и должен использоваться в **других чатах**)

[NetMoose](https://t.me/netmoose)
