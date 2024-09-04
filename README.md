# Netpher

![Go Version](https://img.shields.io/badge/Go-1.19-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-brightgreen)

## Overview

Netpher is a simple TCP utility written in Go, inspired by netcat. It can operate as both a server and a client, allowing for basic data transmission over a network. Netpher also supports remote command execution on the server side.

## Features

- TCP server mode
- TCP client mode
- Remote command execution (server mode)
- Easy to use, lightweight, and efficient

## Usage

### Running the Server

To start a server on a specified port:

```bash
./netpher -mode server -port 1234
```

To start a server on a specified port and execute a command:

```bash
./netpher -mode server -port 1234 -exec "bash"
```

### Running the Client

To connect to a server:

```bash
./netpher -mode client -address 127.0.0.1 -port 1234
```


### Building the Project

Ensure you have Go installed, then build the project with:

```bash
go build
```

### Installation

To install netpher globally on your system, run:

```bash
go install github.com/ginozza/netpher@latest
```


## License

This project is licensed under the MIT License. See the LICENSE file for more details.

## Contributing

Feel free to submit issues or pull requests.

## Disclaimer

This tool is for educational purposes only. Use it responsibly and only in environments where you have permission to operate.




