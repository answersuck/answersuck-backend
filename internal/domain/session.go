package domain

import (
	"errors"
	"fmt"
	"github.com/answersuck/vault/pkg/strings"
	"net/netip"
	"time"
)

// Client errors
var (
	ErrSessionCannotBeTerminated = errors.New("current session cannot be terminated, use logout instead")
)

// System errors
var (
	ErrSessionContextNotFound = errors.New("session not found in context")
	ErrSessionDeviceMismatch  = errors.New("device doesn't match with device of current session")
)

type Session struct {
	Id        string    `json:"id"`
	AccountId string    `json:"accountId"`
	UserAgent string    `json:"userAgent"`
	IP        string    `json:"ip"`
	MaxAge    int       `json:"maxAge"`
	ExpiresAt int64     `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewSession(aid, ua, ip string, expiration time.Duration) (*Session, error) {
	// TODO: add useragent validation

	if _, err := netip.ParseAddr(ip); err != nil {
		return nil, fmt.Errorf("netip.ParseAddr: %w", err)
	}

	sid, err := strings.NewUnique(64)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	return &Session{
		Id:        sid,
		AccountId: aid,
		UserAgent: ua,
		IP:        ip,
		MaxAge:    int(expiration.Seconds()),
		ExpiresAt: now.Add(expiration).Unix(),
		CreatedAt: now,
	}, nil
}
