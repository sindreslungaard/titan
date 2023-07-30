package network

import (
	"sync/atomic"
	"time"
	"titan/program"
	"titan/protocol"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

const (
	ReadBufferSize  = 1024
	WriteBufferSize = 2048
	WriteDeadline   = time.Second * 10
	PingPeriod      = 40 * time.Second
	PongWait        = 80 * time.Second
)

type MessageHandlerFunc func(protocol.Buffer)
type SocketClosedHandlerFunc func()

type Socket struct {
	id   uint64
	conn *websocket.Conn

	closed uint32

	write   chan []byte
	receive MessageHandlerFunc
	onclose SocketClosedHandlerFunc
}

func (s *Socket) readproc() {
	defer program.Recover()
	defer s.Close()

	s.conn.SetReadLimit(ReadBufferSize)
	s.conn.SetReadDeadline(time.Now().Add(PongWait))
	s.conn.SetPongHandler(func(string) error {
		s.conn.SetReadDeadline(time.Now().Add(PongWait))
		return nil
	})

	for {
		_, data, err := s.conn.ReadMessage()

		if err != nil && websocket.IsUnexpectedCloseError(
			err,
			websocket.CloseNormalClosure,
			websocket.CloseGoingAway,
		) {
			log.Debug().Uint64("id", s.id).Err(err).Msg("Failed to read message from client connection")
		} else if err != nil {
			// Underlying network connection is most likely closed, return and clean up
			return
		}

		if program.Config.Network.LogPackets {
			buf := make([]byte, len(data))
			copy(buf, data)
			b := protocol.BufferFrom(buf)
			b.ReadInt()
			header := b.ReadShort()
			log.Info().Int("header", int(header)).Bytes("payload", buf).Msg("IN ")
		}

		s.receive(protocol.BufferFrom(data))
	}
}

func (s *Socket) writeproc() {
	defer program.Recover()
	defer s.Close()

	ticker := time.NewTicker(PingPeriod)
	defer ticker.Stop()

	for {

		select {
		case message, ok := <-s.write:
			{
				err := s.conn.SetWriteDeadline(time.Now().Add(WriteDeadline))

				if err != nil || !ok {
					s.conn.WriteMessage(websocket.CloseMessage, []byte{})
					return
				}

				if err := s.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
					log.Debug().Err(err).Msg("Failed to write message to socket")
					return
				}
			}

		case <-ticker.C:
			{
				s.conn.SetWriteDeadline(time.Now().Add(WriteDeadline))
				if err := s.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					log.Debug().Err(err).Uint64("id", s.id).Msg("Failed to ping socket")
					return
				}
			}
		}
	}
}

func (s *Socket) Write(data []byte) {
	closed := atomic.LoadUint32(&s.closed) == 1

	if closed {
		log.Debug().Uint64("id", s.id).Msg("Can't write to closed socket")
		return
	}

	if program.Config.Network.LogPackets {
		b := protocol.BufferFrom(data)
		b.ReadInt()
		header := b.ReadShort()
		log.Info().Int("header", int(header)).Bytes("payload", data).Msg("OUT")
	}

	s.write <- data
}

func (s *Socket) Close() {
	closed := atomic.LoadUint32(&s.closed) == 1

	if closed {
		return
	}

	atomic.StoreUint32(&s.closed, 1)

	func() {
		defer program.Recover()
		s.onclose()
	}()

	s.conn.Close()
	close(s.write)

	log.Debug().Uint64("id", s.id).Msg("Closed connection")
}

func (s *Socket) OnReceive(f func(protocol.Buffer)) {
	s.receive = f
}

func (s *Socket) OnClose(cb func()) {
	s.onclose = cb
}
