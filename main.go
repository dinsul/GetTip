package main

import (
    "fmt" // пакет для форматированного ввода вывода
    "net/http" // пакет для поддержки HTTP протокола
    "html/template" // для динамической генерации HTML
    "path/filepath"
    "strings" // пакет для работы с  UTF-8 строками
    "log" // пакет для логирования
)

type Team string

const (
    Black Team = "000000"
    Red Team = "FF0000"
    Blue Team = "0000FF"
    Green Team = "00FF00"
)

type Card struct {
    //Id int
    Word string
    Team Team
    //IsClicked bool
    //IsOpened bool
}

type game struct {
    Title string
    Cards [5][5]Card
}

func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm() //анализ аргументов,
    fmt.Println(r.Form)  // ввод информации о форме на стороне сервера
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    var name string
    var question string
    var answer string

    for k, v := range r.Form {
        switch k {
            case "name": name = strings.Join(v, ", ")
            case "question": question = strings.Join(v, ", ")
            case "answer": answer = strings.Join(v, ", ")
        }
    }

    fmt.Fprintf(w, "Привет, %s!\n", name) // отправляем данные на клиентскую сторону

        if question == "q1" {
            switch answer {
            case "a1":
                fmt.Fprintf(w, "Нет. Неверно.")
            case "a2":
                fmt.Fprintf(w, "Правильный ответ.")
            case "a3":
                fmt.Fprintf(w, "Ты идиот?!")
            }
        }
}

func TestRange(rw http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    for k, v := range r.Form {
        fmt.Println(k, ": ", v)
    }

    main_screen := filepath.Join("index.html");

    tmpl, err := template.ParseFiles(main_screen);
    if err != nil {
        http.Error(rw, err.Error(), 400)
		return
	}

    g := game {
        Title : "Hello my game",
    }

    g.Cards[0] = [5]Card {{"Hello", Red}, {"world", Blue}, {"!", Blue}, {"!", Green}, {"!", Blue}}
    g.Cards[1] = [5]Card {{"Привет", Red}, {"world", Blue}, {"!", Blue}, {"!", Green}, {"!", Blue}}
    g.Cards[2] = [5]Card {{"Hello", Red}, {"мир", Blue}, {"!", Blue}, {"!", Green}, {"!", Blue}}
    g.Cards[3] = [5]Card {{"Hello", Red}, {"world", Red}, {"!", Blue}, {"!", Green}, {"!", Blue}}
    g.Cards[4] = [5]Card {{"Hello", Green}, {"world", Blue}, {"!", Blue}, {"!", Green}, {"!", Blue}}

    err = tmpl.ExecuteTemplate(rw, "words", g)
    if err != nil {
        http.Error(rw, err.Error(), 400)
		return
	}
}

func IndexHandler(rw http.ResponseWriter, r *http.Request) {
    http.ServeFile(rw, r, r.URL.Path[1:])
}

func main() {
    //http.HandleFunc("/hello", HomeRouterHandler) // установим роутер
    http.HandleFunc("/test", HomeRouterHandler) // установим роутер
    http.HandleFunc("/", IndexHandler) // установим роутер
    http.HandleFunc("/game", TestRange)
    err := http.ListenAndServe(":8081", nil) // задаем слушать порт
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
