package main

import (
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
	"os"
)

var pwd, _ = os.Getwd()
var ActiveClients = make(map[clientConn]int)

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		err := template.Must(template.ParseFiles(pwd+"/index.html")).Execute(writer, "localhost:8080")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.Handle("/sock", websocket.Handler(SockServer))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic("Не удалось запустить сервер : " + err.Error())
	}

}

type clientConn struct {
	websocket *websocket.Conn
	clientIP  string
	name      string
}

func SockServer(ws *websocket.Conn) {
	var err error
	var clientMessage string

	defer func() {
		if err = ws.Close(); err != nil {
			log.Println("не удалось закрыть", err.Error())
		}
	}()

	client := ws.Request().RemoteAddr
	log.Println("Подключился клиент", client)
	sockCli := clientConn{ws, client, "petrena"}
	ActiveClients[sockCli] = 0
	log.Println("Количество подключенных человек: ", len(ActiveClients))

	for {
		if err = websocket.Message.Receive(ws, &clientMessage); err != nil {
			log.Println("Соединение отключено", err.Error())

			delete(ActiveClients, sockCli)
			log.Println("Всего подключенных людей: ", len(ActiveClients))
			return
		}

		//clientMessage = sockCli.name + clientMessage
		for cs, _ := range ActiveClients {
			if err = websocket.Message.Send(cs.websocket, clientMessage); err != nil {
				log.Println("Не удалось отправить сообщение ", cs.clientIP, err.Error())
			}
		}
	}
}
