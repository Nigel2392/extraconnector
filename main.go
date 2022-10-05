package extraconnector

import (
	"encoding/json"
	"errors"
	"net"
	"strconv"
)

type Server struct {
	IP              string `json:"addr"`
	PORT            int    `json:"port"`
	CONN            net.Conn
	Current_Channel int
}

type Message struct {
	Channel_ID int         `json:"channel_id"`
	Type       string      `json:"type"`
	Key        string      `json:"key"`
	Val        interface{} `json:"val"`
	TTL        int         `json:"ttl"`
}

type Extra_Data struct {
	Data       map[string]any `json:"DATA"`
	Channel_ID int            `json:"CHANNEL"`
	STATUS     string         `json:"STATUS"`
}

func (s *Server) Connect() {
	str_port := strconv.Itoa(s.PORT)
	addr := s.IP + ":" + str_port
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		err = errors.New("connection failed, check your config.json. Address: " + addr)
		panic(err)
	}
	s.CONN = conn
}

func (s *Server) Disconnect() {
	s.CONN.Close()
}

func (s *Server) Send(msg *Message) (Extra_Data, error) {
	json_msg, err := json.Marshal(msg)
	if err != nil {
		return Extra_Data{}, err
	}
	_, err = s.CONN.Write(json_msg)
	if err != nil {
		return Extra_Data{}, err
	}
	return s.Read()
}

func (s *Server) Read() (Extra_Data, error) {
	var data Extra_Data
	buffer := make([]byte, 8096)
	n, err := s.CONN.Read(buffer)
	if err != nil {
		return Extra_Data{}, err
	}
	err = json.Unmarshal(buffer[:n], &data)
	if err != nil {
		return Extra_Data{}, err
	}
	return data, nil
}

func (s *Server) MessageSet(key string, value any, ttl int) Message {
	msg := Message{
		Channel_ID: s.Current_Channel,
		Type:       "SET",
		Key:        key,
		Val:        value,
		TTL:        ttl,
	}
	return msg
}

func (s *Server) MessageGet(key string) Message {
	msg := Message{
		Channel_ID: s.Current_Channel,
		Type:       "GET",
		Key:        key,
	}
	return msg
}

func (s *Server) MessageDel(key string) Message {
	msg := Message{
		Channel_ID: s.Current_Channel,
		Type:       "DEL",
		Key:        key,
	}
	return msg
}

func (s *Server) MessageHasKey(key string) Message {
	msg := Message{
		Channel_ID: s.Current_Channel,
		Type:       "HASKEY",
		Key:        key,
	}
	return msg
}
func (s *Server) MessageSize() Message {
	msg := Message{
		Channel_ID: s.Current_Channel,
		Type:       "SIZE",
	}
	return msg

}
func (s *Server) MessageSizeAll() Message {
	msg := Message{
		Channel_ID: s.Current_Channel,
		Type:       "SIZEALL",
	}
	return msg
}

func (s *Server) MessageKeys() Message {
	msg := Message{
		Channel_ID: s.Current_Channel,
		Type:       "KEYS",
	}
	return msg
}

func (s *Server) MessageSetChannel(channel int) {
	s.Current_Channel = channel
}

func (s *Server) Cache_Set(key string, value any, ttl int) (Extra_Data, error) {
	msg := s.MessageSet(key, value, ttl)
	return s.Send(&msg)
}

func (s *Server) Cache_Get(key string) (Extra_Data, error) {
	msg := s.MessageGet(key)
	return s.Send(&msg)
}

func (s *Server) Cache_Del(key string) (Extra_Data, error) {
	msg := s.MessageDel(key)
	return s.Send(&msg)
}

func (s *Server) Cache_HasKey(key string) (Extra_Data, error) {
	msg := s.MessageHasKey(key)
	return s.Send(&msg)
}

func (s *Server) Cache_Size() (Extra_Data, error) {
	msg := s.MessageSize()
	return s.Send(&msg)
}

func (s *Server) Cache_SizeAll() (Extra_Data, error) {
	msg := s.MessageSizeAll()
	return s.Send(&msg)
}

func (s *Server) Cache_Keys() (Extra_Data, error) {
	msg := s.MessageKeys()
	return s.Send(&msg)
}

func (s *Server) Cache_SetChannel(channel int) {
	s.MessageSetChannel(channel)
}
