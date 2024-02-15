package session

type InMemorySessionStorage struct {
	sessions map[int64]*UserSession
}

func NewInMemorySessionStorage() *InMemorySessionStorage {
	return &InMemorySessionStorage{
		sessions: make(map[int64]*UserSession),
	}
}

func (s *InMemorySessionStorage) GetSession(userID int64) *UserSession {
	if session, exists := s.sessions[userID]; exists {
		return session
	}
	session := NewUserSession(userID)
	s.sessions[userID] = session
	return session
}

func (s *InMemorySessionStorage) SaveSession(session *UserSession) {
	s.sessions[session.UserID] = session
}
