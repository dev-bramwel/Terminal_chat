package chat

import (
	"errors"
	"fmt"
	"log"
	"net"
	"sync"
)

const MaxClients = 10

type Server struct {
	port     string
	clients  map[*Client]bool
	history  []Message
	joinChan chan net.Conn
	mu       sync.Mutex
}

func NewServer(port string) *Server {
	return &Server{
		port:     port,
		clients:  make(map[*Client]bool),
		history:  make([]Message, 0),
		joinChan: make(chan net.Conn),
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Println("Listening on the port :" + s.port)
	go s.acceptConnections(listener)

	for conn := range s.joinChan {
		client := NewClient(conn, s)
		go client.Handle()
	}

	return nil
}

func (s *Server) acceptConnections(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		s.joinChan <- conn
	}
}

func (s *Server) AddClient(client *Client) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.clients) >= MaxClients {
		return false
	}

	s.clients[client] = true
	return true
}

func (s *Server) RemoveClient(client *Client) {
	s.mu.Lock()
	_, exists := s.clients[client]
	if exists {
		delete(s.clients, client)
	}
	s.mu.Unlock()

	if exists {
		s.Broadcast(NewSystemMessage(client.Name()+" has left our chat."), client)
	}
}

func (s *Server) SaveMessage(message Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.history = append(s.history, message)
}

func (s *Server) SendHistory(client *Client) {
	s.mu.Lock()
	history := make([]Message, len(s.history))
	copy(history, s.history)
	s.mu.Unlock()

	for _, message := range history {
		client.WriteLine(message.Format())
	}
}

func (s *Server) Broadcast(message Message, sender *Client) {
	s.mu.Lock()
	clients := make([]*Client, 0, len(s.clients))
	for client := range s.clients {
		if client != sender {
			clients = append(clients, client)
		}
	}
	s.mu.Unlock()

	for _, client := range clients {
		client.WriteLine("\n" + message.Format())
		client.Write(Prompt(client.Name()))
	}
}

func (s *Server) ClientCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.clients)
}

func ValidatePort(port string) error {
	if port == "" {
		return errors.New("empty port")
	}
	return nil
}
