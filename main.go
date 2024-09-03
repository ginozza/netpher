package main

import (
    "flag"
    "fmt"
    "io"
    "log"
    "net"
    "os"
    "os/exec"
)

func main() {
    mode := flag.String("mode", "server", "Mode to run: server or client")
    port := flag.Int("port", 1234, "Port to listen on or connect to")
    execCmd := flag.String("exec", "", "Command to execute on the server side")
    address := flag.String("address", "127.0.0.1", "Address to connect to (client mode only)")
    flag.Parse()

    if *mode == "server" {
        err := handleTCPServer(*port, *execCmd)
        if err != nil {
            log.Fatalf("Error: %v", err)
        }
    } else if *mode == "client" {
        err := handleTCPClient(*address, *port)
        if err != nil {
            log.Fatalf("Error: %v", err)
        }
    } else {
        log.Fatalf("Invalid mode specified: %s. Use 'server' or 'client'.", *mode)
    }
}

// handleTCPServer starts a TCP server that listens on the specified port and handles connections.
func handleTCPServer(port int, execCmd string) error {
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
        go handleConnection(conn, execCmd)
    }
}

// handleConnection manages the individual TCP connection.
func handleConnection(conn net.Conn, execCmd string) {
    defer conn.Close()
    log.Println("TCP connection established")

    // If a command is specified, execute it
    if execCmd != "" {
        log.Printf("Executing command: %s", execCmd)
        cmd := exec.Command(execCmd)
        cmd.Stdout = conn
        cmd.Stderr = conn
        cmd.Stdin = conn
        err := cmd.Start()
        if err != nil {
            log.Printf("error starting command: %v", err)
            return
        }
        err = cmd.Wait()
        if err != nil {
            log.Printf("error waiting for command to finish: %v", err)
        }
        return
    }

    // Echo data received from the connection
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

        _, err = conn.Write(data)
        if err != nil {
            log.Printf("error writing to connection: %v", err)
            break
        }
    }
}

// handleTCPClient connects to a TCP server and handles data exchange.
func handleTCPClient(address string, port int) error {
    serverAddress := fmt.Sprintf("%s:%d", address, port)
    conn, err := net.Dial("tcp", serverAddress)
    if err != nil {
        return fmt.Errorf("error connecting to server: %v", err)
    }
    defer conn.Close()

    log.Printf("Connected to server %s", serverAddress)

    // Forward data from stdin to the connection
    go io.Copy(conn, os.Stdin)
    // Forward data from the connection to stdout
    _, err = io.Copy(os.Stdout, conn)
    if err != nil {
        log.Printf("error copying data from connection: %v", err)
    }

    return nil
}

