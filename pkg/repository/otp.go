package repository

import "errors"

type OtpStorage interface {
	Put(otp, userID string)
	Get(otp string) (string, error)
}

type InMemoryOTP struct {
	storage map[string]string
}

func NewInMemoryOTP() *OtpStorage {
	var storage OtpStorage
	storage = &InMemoryOTP{
		storage: make(map[string]string),
	}
	return &storage
}

func (s *InMemoryOTP) Put(otp, userID string) {
	s.storage[otp] = userID
}

func (s *InMemoryOTP) Get(otp string) (string, error) {
	userId, ok := s.storage[otp]
	if ok {
		delete(s.storage, otp)
		return userId, nil
	}
	return "", errors.New("no such otp")
}
