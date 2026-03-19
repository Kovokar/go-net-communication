package network

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type RemoteStr string

type PlayerInfo struct {
	Remote RemoteStr
	Name   string
}

type TCPServer struct {
	Addr        string
	connections map[net.Conn]bool
	mu          sync.Mutex // Importante para evitar race conditions ao mexer no map/slice
	playerMap   map[RemoteStr]PlayerInfo
}

func NewTCPServer(port string) *TCPServer {
	return &TCPServer{
		Addr:        ":" + port,
		connections: make(map[net.Conn]bool),
		playerMap:   make(map[RemoteStr]PlayerInfo, 0),
	}
}

func (s *TCPServer) Start() {
	listener, err := net.Listen("tcp", s.Addr)
	if err != nil {
		panic(fmt.Sprintf("Erro ao iniciar servidor: %v", err))
	}
	defer listener.Close()

	fmt.Printf("Servidor rodando na porta %s\n", s.Addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Erro ao aceitar conexão:", err)
			continue
		}

		if len(s.playerMap) == 2 {
			fmt.Println(s.connections)
			fmt.Println("Nova conexão rejeitada: limite atingido")
			conn.Write([]byte("Servidor cheio. Tente novamente mais tarde.\n"))
			conn.Close()
			continue
		}

		// Registra a conexão de forma segura
		s.mu.Lock()
		s.connections[conn] = true
		s.mu.Unlock()

		s.showConnections()

		// Inicia o tratamento da conexão
		go s.handleConnection(conn)
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {
	remoteStr := RemoteStr(conn.RemoteAddr().String())
	defer func() {
		s.mu.Lock()
		delete(s.connections, conn)
		delete(s.playerMap, remoteStr)
		s.mu.Unlock()

		conn.Close()
		fmt.Println("Cliente desconectado:", conn.RemoteAddr())
		s.showConnections()
	}()

	scanner := bufio.NewScanner(conn)

	// A primeira mensagem é tratada como identificação (Nome)
	if scanner.Scan() {
		name := scanner.Text()
		// name := "pedro"
		s.mu.Lock()
		newPlayer := PlayerInfo{remoteStr, name}
		s.playerMap[newPlayer.Remote] = newPlayer
		s.mu.Unlock()
		fmt.Printf("👤 Jogador identificado: (%s)\n", name)
	}

	// Loop de mensagens subsequentes
	for scanner.Scan() {
		msg := scanner.Text()
		name := s.playerMap[remoteStr].Name

		fmt.Printf("Mensagem de [%s]: %s\n", name, msg)
	}
}

func (s *TCPServer) showConnections() {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Printf("Conexões ativas (%d): \n", len(s.connections))
	for conn := range s.connections {
		fmt.Println("-", conn.RemoteAddr())
	}
	fmt.Println()
}
