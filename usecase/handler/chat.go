package handler

import (
	"database/sql"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var (
	upgrader      = websocket.Upgrader{}
	clients       = make(map[string]*websocket.Conn) // map of users to WebSocket connections
	mutex         = &sync.Mutex{}
	messageBuffer = make([]string, 0)
)

func NewChat(router *mux.Router, db *sql.DB) *mux.Router {

	router.HandleFunc("/ws/{sender}/{recipient}", websocketHandler(db))

	return router
}

func websocketHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		sender := params["sender"]
		recipient := params["recipient"]

		// Upgrade the connection to a WebSocket
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Info().Err(err)
			return
		}
		defer conn.Close()

		// Register the client with the sender and recipient usernames
		mutex.Lock()
		clients[sender+"_"+recipient] = conn
		mutex.Unlock()

		// Read existing messages from the database (if any)
		messages, err := getMessages(db, sender, recipient)
		if err != nil {
			log.Info().Err(err)
			return
		}

		// Send existing messages to the new client
		for _, message := range messages {
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Info().Err(err)
				return
			}
		}

		// Listen for new messages
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Info().Err(err)
				break
			}

			log.Debug().Str("sender", sender).Str("recipient", recipient).Str("message", string(msg)).Msg("Incoming message")

			// Save the message to the database
			err = saveMessage(db, sender, recipient, string(msg))
			if err != nil {
				log.Info().Err(err)
				return
			}

			// Broadcast the message to both sender and recipient WebSocket connections
			broadcast(sender, recipient, string(msg))
		}
	}
}

func broadcast(sender, recipient, message string) {
	mutex.Lock()
	defer mutex.Unlock()

	// Add the message to the buffer
	messageBuffer = append(messageBuffer, message)

	// Send the message to the sender's WebSocket connections
	if senderConn, ok := clients[sender+"_"+recipient]; ok {
		err := senderConn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Info().Err(err)
			delete(clients, sender+"_"+recipient)
		}
	}

	// Send the message to the recipient's WebSocket connections
	if recipientConn, ok := clients[recipient+"_"+sender]; ok {
		err := recipientConn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Info().Err(err)
			delete(clients, recipient+"_"+sender)
		}
	}
}

func saveMessage(db *sql.DB, sender, recipient, message string) error {
	_, err := db.Exec("INSERT INTO messages (sender, recipient, message) VALUES (?, ?, ?)", sender, recipient, message)
	return err
}

func getMessages(db *sql.DB, sender, recipient string) ([]string, error) {
	rows, err := db.Query("SELECT message FROM messages WHERE (sender = ? AND recipient = ?) OR (sender = ? AND recipient = ?) ORDER BY id", sender, recipient, recipient, sender)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []string
	for rows.Next() {
		var message string
		err := rows.Scan(&message)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
