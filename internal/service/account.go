package service

import (
	"context"
	"fmt"
	"github.com/answersuck/answersuck-backend/internal/dto"
	"github.com/answersuck/answersuck-backend/pkg/strings"
	"net/mail"
	"time"

	"github.com/answersuck/answersuck-backend/pkg/logging"

	"github.com/answersuck/answersuck-backend/internal/config"
	"github.com/answersuck/answersuck-backend/internal/domain"
	"github.com/answersuck/answersuck-backend/pkg/auth"
	"github.com/answersuck/answersuck-backend/pkg/blocklist"
	"github.com/answersuck/answersuck-backend/pkg/storage"
)

type accountService struct {
	cfg *config.Aggregate
	log logging.Logger

	repo    AccountRepo
	session Session
	email   Email

	token     auth.TokenManager
	storage   storage.Uploader
	blockList blocklist.Finder
}

const (
	verificationCodeLength   = 64
	passwordResetTokenLength = 64
)

func NewAccountService(cfg *config.Aggregate, l logging.Logger, r AccountRepo, s Session,
	t auth.TokenManager, e Email, u storage.Uploader, b blocklist.Finder) *accountService {
	return &accountService{
		cfg:       cfg,
		log:       l,
		repo:      r,
		token:     t,
		session:   s,
		email:     e,
		storage:   u,
		blockList: b,
	}
}

func (s *accountService) Create(ctx context.Context, a *domain.Account) (*domain.Account, error) {
	if s.blockList.Find(a.Username) {
		return nil, fmt.Errorf("accountService - Create - s.blockList.Find: %w", domain.ErrAccountForbiddenUsername)
	}

	if err := a.GeneratePasswordHash(); err != nil {
		return nil, fmt.Errorf("accountService - Create - acc.GeneratePasswordHash: %w", err)
	}

	a.SetDiceBearAvatar()

	if err := a.GenerateVerificationCode(verificationCodeLength); err != nil {
		return nil, fmt.Errorf("accountService - Create - a.GenerateVerificationCode: %w", err)
	}

	a, err := s.repo.Create(ctx, a)
	if err != nil {
		return nil, fmt.Errorf("accountService - Create - s.repo.Create: %w", err)
	}

	go func() {
		_ = s.email.SendAccountVerificationMail(ctx, a.Email, a.VerificationCode)
	}()

	return a, nil
}

func (s *accountService) GetById(ctx context.Context, aid string) (*domain.Account, error) {
	acc, err := s.repo.FindById(ctx, aid)
	if err != nil {
		return nil, fmt.Errorf("accountService - GetByID - s.repo.FindByID: %w", err)
	}

	return acc, nil
}

func (s *accountService) GetByEmail(ctx context.Context, email string) (*domain.Account, error) {
	acc, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("accountService - GetByEmail - s.repo.FindByEmail: %w", err)
	}

	return acc, nil
}

func (s *accountService) GetByUsername(ctx context.Context, username string) (*domain.Account, error) {
	acc, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("accountService - GetByUsername - s.repo.FindByUsername: %w", err)
	}

	return acc, nil
}

func (s *accountService) Delete(ctx context.Context, aid, sid string) error {
	if err := s.repo.Archive(ctx, aid, true, time.Now()); err != nil {
		return fmt.Errorf("accountService - Archive - s.repo.Archive: %w", err)
	}

	if err := s.session.Terminate(ctx, sid); err != nil {
		return fmt.Errorf("accountService - Archive - s.session.TerminateAll: %w", err)
	}

	return nil
}

func (s *accountService) RequestVerification(ctx context.Context, aid string) error {
	a, err := s.repo.FindVerification(ctx, aid)
	if err != nil {
		return fmt.Errorf("accountService - RequestVerification - s.repo.FindById: %w", err)
	}

	if a.Verified {
		return fmt.Errorf("accountService: %w", domain.ErrAccountAlreadyVerified)
	}

	go func() {
		_ = s.email.SendAccountVerificationMail(ctx, a.Email, a.Code)
	}()

	return nil
}

func (s *accountService) Verify(ctx context.Context, code string, verified bool) error {
	if err := s.repo.Verify(ctx, code, verified, time.Now()); err != nil {
		return fmt.Errorf("accountService - Verify - s.repo.Verify: %w", err)
	}

	return nil
}

func (s *accountService) RequestPasswordReset(ctx context.Context, login string) error {
	email := login

	if _, err := mail.ParseAddress(login); err != nil {

		a, err := s.GetByUsername(ctx, login)
		if err != nil {
			return fmt.Errorf("accountService - RequestPasswordReset - s.GetByUsername: %w", err)
		}

		email = a.Email
	}

	t, err := strings.NewUnique(passwordResetTokenLength)
	if err != nil {
		return fmt.Errorf("accountService - RequestPasswordReset - strings.NewUnique: %w", err)
	}

	if err = s.repo.InsertPasswordResetToken(ctx, email, t); err != nil {
		return fmt.Errorf("accountService - RequestPasswordReset - s.repo.InsertPasswordResetToken: %w", err)
	}

	if err = s.email.SendAccountPasswordResetMail(ctx, email, t); err != nil {
		return fmt.Errorf("accountService - RequestPasswordReset - s.email.SendAccountPasswordResetMail: %w", err)
	}

	return nil
}

func (s *accountService) PasswordReset(ctx context.Context, token, password string) error {
	t, err := s.repo.FindPasswordResetToken(ctx, token)
	if err != nil {
		return fmt.Errorf("accountService - PasswordReset - s.repo.FindPasswordResetToken: %w", err)
	}

	d := t.CreatedAt.Add(s.cfg.Password.ResetTokenExp)
	if time.Now().After(d) {
		return fmt.Errorf("accountService - PasswordReset: %w", domain.ErrAccountResetPasswordTokenExpired)
	}

	a := domain.Account{Password: password}
	if err = a.GeneratePasswordHash(); err != nil {
		return fmt.Errorf("accountService - PasswordReset - a.GeneratePassword: %w", err)
	}

	if err = s.repo.UpdatePasswordWithToken(ctx, dto.AccountUpdatePassword{
		Token:        t.Token,
		AccountId:    t.AccountId,
		PasswordHash: a.PasswordHash,
		UpdatedAt:    time.Now(),
	}); err != nil {
		return fmt.Errorf("accountService - PasswordReset - s.repo.UpdatePasswordWithToken: %w", err)
	}

	return nil
}
