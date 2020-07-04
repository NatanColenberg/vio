package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WsMessage struct {
	Type string
	Data string
}

type WsStudiesMessage struct {
	Type string
	Data []Study
}

type WeImgMessage struct {
	Type string
	Data []byte
}

type Connection struct {
	Addr string
	Type string
	WS   *websocket.Conn
}

// WebSocket
// var wsConnections = make(map[string]*websocket.Conn)
var wsConnections = []Connection{}
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {

	// App Constance
	const buildPath string = "build/"
	const port int = 8080

	// Router
	router := mux.NewRouter()

	// WebSocket Endpoint
	router.HandleFunc("/ws/studyList", wsStudyListEndpoint)
	router.HandleFunc("/ws/viewer", wsViewerEndpoint)

	// File Server
	router.PathPrefix("/").Handler(http.FileServer(http.Dir(buildPath)))

	// CORS Headers
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "POST", "DELETE"}),
	)

	// Register Middleware
	router.Use(loggingMiddleware)

	// Run Server
	srv := &http.Server{
		Handler: cors(router),
		Addr:    ":" + strconv.Itoa(port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	color.New(color.BgHiGreen, color.FgBlack, color.Bold).
		Println("Server is Running on PORT " + strconv.Itoa(port))

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

// Handlers

func getStudiesHandler(remoteAddr string) {
	for _, conn := range wsConnections {
		if conn.Addr == remoteAddr {
			msg := WsStudiesMessage{Type: "studyListUpdate", Data: Studies}
			jsonMsg, err := json.Marshal(msg)
			if err != nil {
				log.Fatal("Failed to Marshal message")
			}
			conn.WS.WriteMessage(1, jsonMsg)
		}
	}
}

func selectedStudyChangedHandler(studyID string) {

	if studyID == "next" {

	} else if studyID == "prev" {

	} else {
		// Change Selected Study
		for index, study := range Studies {
			if study.Accession == studyID {
				Studies[index].Selected = true
			} else {
				Studies[index].Selected = false
			}
		}
	}

	// Send Updates
	for _, conn := range wsConnections {
		switch conn.Type {
		case "viewer":
			getSelectedStudyImage(conn)
			break

		case "studyList":
			getStudiesHandler(conn.Addr)
		}

	}
}

// Helper Methods

func getSelectedStudyImage(conn Connection) {
	studyID := getSelectedStudyID(0)
	imgName := studyID + ".jpg"
	img, err := ioutil.ReadFile("./img/" + imgName) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	msg := WeImgMessage{Type: "selectedStudyChanged", Data: img}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Fatal("Failed to Marshal message")
	}
	conn.WS.WriteMessage(1, jsonMsg)
}

func getSelectedStudyID(indexDeviation int) string {
	for index, study := range Studies {
		if study.Selected {
			if (index+indexDeviation) < 0 || index+indexDeviation > (len(Studies)-1) {
				return Studies[index].Accession
			} else {
				return Studies[index+indexDeviation].Accession
			}

		}
	}
	return ""
}

// Middleware Methods

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mes := "URI: " + r.RequestURI + ", RemoteAddr: " + r.RemoteAddr + ", Method:" + r.Method
		log.Println(mes)
		next.ServeHTTP(w, r)
	})
}

// WebSocket
func wsStudyListEndpoint(w http.ResponseWriter, r *http.Request) {
	wsEndpoint(w, r, "studyList")
}
func wsViewerEndpoint(w http.ResponseWriter, r *http.Request) {
	wsEndpoint(w, r, "viewer")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request, connType string) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	conn := Connection{Addr: ws.RemoteAddr().String(), Type: connType, WS: ws}
	wsConnections = append(wsConnections, conn)

	log.Printf("Client [%s - %s] Successfully Connected...", r.RemoteAddr, connType)

	reader(conn)
}

func reader(conn Connection) {
	for {
		// read in a message
		messageType, p, err := conn.WS.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print out that message for clarity
		var userMessage WsMessage
		unmarshalErr := json.Unmarshal([]byte(p), &userMessage)
		if unmarshalErr != nil {
			log.Println(unmarshalErr)
		}
		logMessage := "Message from [%v] : [%d](%s) %s\n"
		log.Printf(logMessage, conn.WS.RemoteAddr().String(), messageType, userMessage.Type, userMessage.Data)

		switch userMessage.Type {
		case "getStudies":
			getStudiesHandler(conn.WS.RemoteAddr().String())
			break
		case "getStudyImage":
			getSelectedStudyImage(conn)
			break
		case "selectedStudyChanged":
			selectedStudyChangedHandler(userMessage.Data)
			break
		case "nextSelectedStudy":
			studyID := getSelectedStudyID(1)
			selectedStudyChangedHandler(studyID)
			break
		case "prevSelectedStudy":
			studyID := getSelectedStudyID(-1)
			selectedStudyChangedHandler(studyID)
			break
		}
	}
}
