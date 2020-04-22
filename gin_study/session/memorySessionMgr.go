package session

import (
	"sync"

	uuid "github.com/satori/go.uuid"
)

//-----------------------------------------------------------------------------
type MemorySessionMgr struct {
	sessionMap map[string]Session
	lock       sync.RWMutex
}

func NewMemorySessionMgr() *MemorySessionMgr {
	mgr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return mgr
}

func (m *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

func (m *MemorySessionMgr) CreateSession() (session Session, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	id := uuid.NewV4()
	sessionId := id.String()
	session = NewMemorySession(sessionId)
	return
}
