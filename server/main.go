package main

import (
	"context"
	"log"
	"net"
	"os"
	"sync"
	"time"
	chittychat "your_project/proto"

	"google.golang.org/grpc"
)

type server struct {
	chittychat.UnimplementedChittyChatServer
	mu          sync.Mutex
	clients     map[string]chittychat.ChittyChat_JoinChatServer
	lamportTime int64
	logger      *log.Logger // Add a logger field
}

func (s *server) JoinChat(req *chittychat.JoinRequest, stream chittychat.ChittyChat_JoinChatServer) error {
	clientID := req.GetClientId()

	s.mu.Lock()
	s.clients[clientID] = stream
	s.lamportTime++
	joinMessage := &chittychat.BroadcastMessage{
		ClientId:    "system",
		Message:     "Participant " + clientID + " joined Chitty-Chat",
		LamportTime: s.lamportTime,
	}
	s.broadcast(joinMessage)

	// Use the logger to log the client joining along with the Lamport timestamp
	logMessage := "Participant %s joined the chat with Lamport time %d"
	s.logger.Printf(logMessage, clientID, s.lamportTime)

	s.mu.Unlock()

	for {
		time.Sleep(time.Second)
	}
}

func (s *server) PublishMessage(ctx context.Context, req *chittychat.PublishRequest) (*chittychat.PublishResponse, error) {
	clientID := req.GetClientId()
	message := req.GetMessage()

	// Make sure that the message is under 128 chars
	if len(message) > 128 {
		return &chittychat.PublishResponse{Success: false, Error: "Message too long"}, nil
	}

	s.mu.Lock()
	s.lamportTime++
	broadcastMessage := &chittychat.BroadcastMessage{
		ClientId:    clientID,
		Message:     message,
		LamportTime: s.lamportTime,
	}
	s.broadcast(broadcastMessage)
	s.mu.Unlock()

	// Use the logger to log the published message along with the Lamport timestamp
	logMessage := "Message from %s: %s (Lamport time: %d)"
	s.logger.Printf(logMessage, clientID, message, s.lamportTime)
	return &chittychat.PublishResponse{Success: true}, nil
}

func (s *server) LeaveChat(ctx context.Context, req *chittychat.LeaveRequest) (*chittychat.LeaveResponse, error) {
	clientID := req.GetClientId()

	s.mu.Lock()
	delete(s.clients, clientID)
	s.lamportTime++
	leaveMessage := &chittychat.BroadcastMessage{
		ClientId:    "system",
		Message:     "Participant " + clientID + " left Chitty-Chat",
		LamportTime: s.lamportTime,
	}
	s.broadcast(leaveMessage)
	s.mu.Unlock()

	// Use the logger to log the client leaving along with the Lamport timestamp
	logMessage := "Participant %s left the chat (Lamport time: %d)"
	s.logger.Printf(logMessage, clientID, s.lamportTime)
	return &chittychat.LeaveResponse{Success: true}, nil
}

// The server broadcasts the message to all clients
func (s *server) broadcast(msg *chittychat.BroadcastMessage) {
	for _, client := range s.clients {
		client.Send(msg)
	}
}

func main() {
	// Open a file for logging
	logFile, err := os.OpenFile("serverLogUser.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening log file: %v", err)
	}
	defer logFile.Close()

	// Create a new logger that writes to the file
	logger := log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	s := &server{
		clients:     make(map[string]chittychat.ChittyChat_JoinChatServer),
		lamportTime: 0,
		logger:      logger, // Set the logger in the server struct
	}
	chittychat.RegisterChittyChatServer(grpcServer, s)

	logger.Println("Server started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}
}
