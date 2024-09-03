package main

import (
    "flag"
    "fmt"
    "io"
    "log"
    "net"
)

func main() {
    // Define flags
    port := flag.Int("port", 1234, "Port to listen on")
    exec := flag.String("exec", "", "Command to execute")
    token := flag.String("token", "", "Optional token for authentication")
    flag.Parse()

    // Start handling TCP connections
    err := handleTCP(*port, *exec, *token)
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
}

// handleTCP starts a TCP server that listens on the specified port and handles connections.
func handleTCP(port int, exec string, token string) error {
    address := fmt.Sprintf(":%d", port)
    listener, err := net.Listen("tcp", address)
    if err != nil {
        return fmt.Errorf("error starting TCP server: %v", err)
    }
    defer listener.Close()

    log.Printf("TCP server listening on port %d", port)

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("error accepting connection: %v", err)
            continue
        }
        go handleConnection(conn, exec, token)
    }
}

// handleConnection manages the individual TCP connection.
func handleConnection(conn net.Conn, exec string, token string) {
    defer conn.Close()

    log.Println("TCP connection established")

    // For demonstration, just echo data received from the connection
    buffer := make([]byte, 1024)
    for {
        n, err := conn.Read(buffer)
        if err != nil {
            if err != io.EOF {
                log.Printf("error reading from connection: %v", err)
            }
            break
        }

        data := buffer[:n]
        log.Printf("Received data: %s", string(data))

        // Example: Execute a command if specified
        if exec != "" {
            log.Printf("Command to execute: %s", exec)
            // Implement command execution logic here if needed
        }

        // Echo the data back to the client
        _, err = conn.Write(data)
        if err != nil {
            log.Printf("error writing to connection: %v", err)
            break
        }
    }
}
