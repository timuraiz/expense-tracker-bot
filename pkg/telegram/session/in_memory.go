package session

type InMemorySessionStorage struct {
	sessions map[int64]*UserSession
}

func NewInMemorySessionStorage() *InMemorySessionStorage {
	return &InMemorySessionStorage{
		sessions: make(map[int64]*UserSession),
	}
}

func (s *InMemorySessionStorage) GetSession(userID int64) (*UserSession, error) {
	if session, exists := s.sessions[userID]; exists {
		return session, nil
	}
	session := NewUserSession(userID)
	s.SaveSession(session)
	return session, nil
}

func (s *InMemorySessionStorage) SaveSession(session *UserSession) error {
	s.sessions[session.UserID] = session
	return nil
}
