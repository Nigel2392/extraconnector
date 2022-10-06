# ExtraConnector (GO)
Connector for Cache, written in Golang

To use, initialise the server struct at first:
```
var S = extraconnector.Server{
    IP: "127.0.0.1",
    PORT: 3239,
    Current_Channel: 0
}
```

```
// Set Channel for server
func (s *Server) MessageSetChannel(channel int) {
	s.Current_Channel = channel
}

// Set value with key in current cache
func (s *Server) Cache_Set(key string, value any, ttl int) (Extra_Data, error) {
	msg := s.MessageSet(key, value, ttl)
	return s.Send(&msg)
}

// Get value from current cache
func (s *Server) Cache_Get(key string) (Extra_Data, error) {
	msg := s.MessageGet(key)
	return s.Send(&msg)
}

// Delete value from current cache.
func (s *Server) Cache_Del(key string) (Extra_Data, error) {
	msg := s.MessageDel(key)
	return s.Send(&msg)
}

// Check if current cache has key
func (s *Server) Cache_HasKey(key string) (Extra_Data, error) {
	msg := s.MessageHasKey(key)
	return s.Send(&msg)
}

// Check amount of keys in current cache
func (s *Server) Cache_Size() (Extra_Data, error) {
	msg := s.MessageSize()
	return s.Send(&msg)
}

// Check amount of keys in all caches
func (s *Server) Cache_SizeAll() (Extra_Data, error) {
	msg := s.MessageSizeAll()
	return s.Send(&msg)
}

// Get all keys for the current cache
func (s *Server) Cache_Keys() (Extra_Data, error) {
	msg := s.MessageKeys()
	return s.Send(&msg)
}

// Set channel for the current cache
func (s *Server) Cache_SetChannel(channel int) {
	s.MessageSetChannel(channel)
}
```
