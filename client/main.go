package main

import (
	"bufio"
	"context"
	"log"
	"os"

	chittychat "your_project/proto"

	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <client-id>", os.Args[0])
	}
	clientID := os.Args[1]

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := chittychat.NewChittyChatClient(conn)

	// Client request server to join the chat
	stream, err := client.JoinChat(context.Background(), &chittychat.JoinRequest{ClientId: clientID})
	if err != nil {
		log.Fatalf("could not join chat: %v", err)
	}

	// This goRoutine recieves messages from the chat
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error receiving message: %v", err)
			}
			log.Printf("Received message: %s [Lamport Time: %d]", msg.Message, msg.LamportTime)
		}
	}()

	// Scanner to take input from client
	scanner := bufio.NewScanner(os.Stdin)

	log.Println("You can start sending messages. Type 'exit' to leave the chat.")

	// Read and send messages until the user types "exit"
	for {

		log.Print("Enter message: ")

		if scanner.Scan() {
			message := scanner.Text()

			// type exit to leave the chat
			if message == "exit" {
				break
			}

			_, err = client.PublishMessage(context.Background(), &chittychat.PublishRequest{
				ClientId: clientID,
				Message:  message,
			})
			if err != nil {
				log.Printf("could not publish message: %v", err)
			}
		}
	}

	_, err = client.LeaveChat(context.Background(), &chittychat.LeaveRequest{ClientId: clientID})
	if err != nil {
		log.Fatalf("could not leave chat: %v", err)
	}

	log.Printf("%s left the chat", clientID)
}
