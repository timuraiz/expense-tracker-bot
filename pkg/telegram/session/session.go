package session

type State struct {
	name string
}

func (s *State) GetName() string {
	return s.name
}

func NewState(name string) State {
	return State{
		name: name,
	}
}

type UserSession struct {
	UserID int64 // Telegram user ID
	State  State
	Data   map[string]interface{} // Store additional data needed for the conversation
}

func (s *UserSession) SetState(state State) {
	s.State = state
}

func NewUserSession(userID int64) *UserSession {
	return &UserSession{
		UserID: userID,
		State:  State{},
		Data:   make(map[string]interface{}),
	}
}

type Session interface {
	GetSession(userID int64) (*UserSession, error)
	SaveSession(session *UserSession) error
	// DeleteSession(userID int64) // Optional: Implement if you need to explicitly delete sessions
}
