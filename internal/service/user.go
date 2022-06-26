package service

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (s *Service) GetUserInfo(selfId uint64, UserID uint64) *User {
	return nil
}
