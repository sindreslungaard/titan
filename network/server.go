package network

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"titan/program"
	"titan/protocol"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type ConnectionHandler func(*Socket)

type Server struct {
	next              uint64
	connectionhandler ConnectionHandler
}

func NewServer() *Server {
	server := &Server{
		connectionhandler: func(s *Socket) {
			log.Warn().Msg("Default connection handler is being used")
		},
	}
	return server
}

func (s *Server) OnConnect(f ConnectionHandler) *Server {
	s.connectionhandler = f
	return s
}

func (s *Server) Listen(port int) error {
	http.HandleFunc("/", s.connect)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", s.connect)

	log.Info().Int("port", port).Msg("Server started")
	return http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  ReadBufferSize,
	WriteBufferSize: WriteBufferSize,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (s *Server) connect(w http.ResponseWriter, r *http.Request) {
	defer program.Recover()

	log.Debug().Msg("New connection")

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Debug().Err(err).Msg("")
		return
	}

	socket := &Socket{
		id:     atomic.AddUint64(&s.next, 1),
		conn:   conn,
		closed: 0,
		write:  make(chan []byte),
		receive: func(b protocol.Buffer) {
			log.Warn().Msg("Default socket receive handler is being used")
		},
		onclose: func() {
			log.Warn().Msg("Default socket close handler is being used")
		},
	}

	s.connectionhandler(socket)

	go socket.readproc()
	go socket.writeproc()
}
