package network

import (
	"bufio"
	"fmt"
	"net"
)

type TCPClient struct {
	conn net.Conn
}

func NewTCPClient(address string) (*TCPClient, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &TCPClient{conn: conn}, nil
}

func (c *TCPClient) Send(message string) error {
	_, err := fmt.Fprint(c.conn, message+"\n")
	return err
}

func (c *TCPClient) Listen(handler func(string)) {
	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		handler(scanner.Text())
	}
}

func (c *TCPClient) Close() {
	c.conn.Close()
}

func (c *TCPClient) Identification(clientName string) {
	_, err := fmt.Fprintf(c.conn, "%s\n", clientName)
	if err != nil {
		fmt.Println("Erro ao enviar identificação:", err)
		return
	}
}
