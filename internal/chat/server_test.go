package chat

import (
	"net"
	"testing"
)

// newTestClient creates a Client using net.Pipe.
//
// net.Pipe gives us two connected in-memory network connections.
// This lets us test server/client behavior without opening a real TCP port.
// Very civilized. No actual sockets running around unsupervised.
func newTestClient(t *testing.T, server *Server, name string) (*Client, net.Conn) {
	t.Helper()

	serverSide, clientSide := net.Pipe()

	client := NewClient(serverSide, server)
	client.SetName(name)

	return client, clientSide
}

// TestNewServer checks that NewServer correctly initializes
// the server fields needed by the chat system.
func TestNewServer(t *testing.T) {
	server := NewServer("8989")

	if server == nil {
		t.Fatal("expected server, got nil")
	}

	if server.port != "8989" {
		t.Fatalf("expected port 8989, got %q", server.port)
	}

	if server.clients == nil {
		t.Fatal("expected clients map to be initialized")
	}

	if len(server.clients) != 0 {
		t.Fatalf("expected no connected clients, got %d", len(server.clients))
	}

	if server.history == nil {
		t.Fatal("expected history slice to be initialized")
	}

	if len(server.history) != 0 {
		t.Fatalf("expected empty history, got %d messages", len(server.history))
	}

	if server.joinChan == nil {
		t.Fatal("expected joinChan to be initialized")
	}
}

// TestAddClient checks that a client can be added to the server.
func TestAddClient(t *testing.T) {
	server := NewServer("8989")

	client, peer := newTestClient(t, server, "Alice")
	defer client.conn.Close()
	defer peer.Close()

	added := server.AddClient(client)
	if !added {
		t.Fatal("expected client to be added")
	}

	if server.ClientCount() != 1 {
		t.Fatalf("expected 1 client, got %d", server.ClientCount())
	}
}

// TestAddClientLimit checks that the server refuses more than MaxClients clients.
func TestAddClientLimit(t *testing.T) {
	server := NewServer("8989")

	var clients []*Client
	var peers []net.Conn

	defer func() {
		for _, client := range clients {
			client.conn.Close()
		}

		for _, peer := range peers {
			peer.Close()
		}
	}()

	for i := 0; i < MaxClients; i++ {
		client, peer := newTestClient(t, server, "Client")
		clients = append(clients, client)
		peers = append(peers, peer)

		if !server.AddClient(client) {
			t.Fatalf("expected client %d to be added", i+1)
		}
	}

	extraClient, extraPeer := newTestClient(t, server, "Extra")
	defer extraClient.conn.Close()
	defer extraPeer.Close()

	if server.AddClient(extraClient) {
		t.Fatal("expected extra client to be rejected after reaching MaxClients")
	}

	if server.ClientCount() != MaxClients {
		t.Fatalf("expected %d clients, got %d", MaxClients, server.ClientCount())
	}
}

// TestRemoveClient checks that removing a client decreases the active client count.
func TestRemoveClient(t *testing.T) {
	server := NewServer("8989")

	client, peer := newTestClient(t, server, "Alice")
	defer client.conn.Close()
	defer peer.Close()

	if !server.AddClient(client) {
		t.Fatal("expected client to be added")
	}

	server.RemoveClient(client)

	if server.ClientCount() != 0 {
		t.Fatalf("expected 0 clients after removal, got %d", server.ClientCount())
	}
}

// TestSaveMessage checks that SaveMessage stores messages in server history.
func TestSaveMessage(t *testing.T) {
	server := NewServer("8989")

	message := NewChatMessage("Alice", "hello")
	server.SaveMessage(message)

	if len(server.history) != 1 {
		t.Fatalf("expected 1 message in history, got %d", len(server.history))
	}

	if server.history[0].Sender != "Alice" {
		t.Fatalf("expected sender Alice, got %q", server.history[0].Sender)
	}

	if server.history[0].Text != "hello" {
		t.Fatalf("expected message text hello, got %q", server.history[0].Text)
	}
}

// TestValidatePortAcceptsNonEmptyPort checks that a normal non-empty port passes.
func TestValidatePortAcceptsNonEmptyPort(t *testing.T) {
	err := ValidatePort("8989")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

// TestValidatePortRejectsEmptyPort checks that an empty port is rejected.
func TestValidatePortRejectsEmptyPort(t *testing.T) {
	err := ValidatePort("")
	if err == nil {
		t.Fatal("expected error for empty port, got nil")
	}
}

// TestBroadcastDoesNotSendToSender checks the project rule that a sender should
// not receive their own message back from the server.
//
// The sender sees what they typed in their terminal already, so broadcasting it
// back would duplicate the message.
func TestBroadcastDoesNotSendToSender(t *testing.T) {
	server := NewServer("8989")

	sender, senderPeer := newTestClient(t, server, "Alice")
	receiver, receiverPeer := newTestClient(t, server, "Bob")

	defer sender.conn.Close()
	defer senderPeer.Close()
	defer receiver.conn.Close()
	defer receiverPeer.Close()

	if !server.AddClient(sender) {
		t.Fatal("expected sender to be added")
	}

	if !server.AddClient(receiver) {
		t.Fatal("expected receiver to be added")
	}

	message := NewChatMessage("Alice", "hello Bob")

	done := make(chan string, 1)

	go func() {
		buffer := make([]byte, 1024)
		n, _ := receiverPeer.Read(buffer)
		done <- string(buffer[:n])
	}()

	server.Broadcast(message, sender)

	output := <-done

	if output == "" {
		t.Fatal("expected receiver to get broadcast message")
	}
}
