package session

import (
	"errors"
	"sync"
)

//-----------------------------------------------------------------------------
type MemorySession struct {
	sessionId string
	data      map[string]interface{}
	dataLock  sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}

	return s
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.dataLock.Lock()
	defer m.dataLock.Unlock()
	m.data[key] = value
	return
}

func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.dataLock.RLock()
	defer m.dataLock.RUnlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

func (m *MemorySession) Del(key string) (err error) {
	m.dataLock.Lock()
	defer m.dataLock.Unlock()
	delete(m.data, key)
	return
}

func (m *MemorySession) Save() (err error) {
	return
}
