package service

import (
	"errors"
	"testing"

	"github.com/guide-backend/internal/model"
	"github.com/stretchr/testify/assert"
)

// --- Mock Implementations ---

type mockUserRepo struct {
	fromUser model.User
	toUser   model.User
	failCase string
}

func (m *mockUserRepo) GetUserByID(id uint) (model.User, error) {
	if m.failCase == "sender_not_found" {
		return model.User{}, errors.New("not found")
	}
	return m.fromUser, nil
}

func (m *mockUserRepo) GetUserByEmail(email string) (model.User, error) {
	if m.failCase == "receiver_not_found" {
		return model.User{}, errors.New("not found")
	}
	return m.toUser, nil
}

func (m *mockUserRepo) CreateUser(user model.User) (model.User, error) {
	return model.User{}, nil
}

func (m *mockUserRepo) GetAllUsers() ([]model.User, error) {
	return nil, nil
}

type mockTransferRepo struct {
	failCase string
}

func (m *mockTransferRepo) TransferBalanceAtomic(fromID, toID uint, amount int) error {
	if m.failCase == "insufficient_balance" {
		return errors.New("insufficient balance")
	}
	return nil
}

// --- Test Case ---

const testEmail = "bob@example.com"

func TestTransferSuccess(t *testing.T) {
	s := NewTransferService(
		&mockTransferRepo{},
		&mockUserRepo{
			fromUser: model.User{
				ID:      1,
				Name:    "Alice",
				Balance: 1000,
			},
			toUser: model.User{
				ID:   2,
				Name: "Bob",
			},
		},
	)

	req := TransferRequest{
		FromUserID: 1,
		ToEmail:    testEmail,
		Amount:     100,
	}

	resp, err := s.Transfer(req)

	assert.NoError(t, err)
	assert.Equal(t, "Alice", resp.FromName)
	assert.Equal(t, "Bob", resp.ToName)
	assert.Equal(t, 100, resp.Amount)
	assert.Equal(t, "success", resp.Status)
}

func TestTransferInvalidAmount(t *testing.T) {
	s := NewTransferService(&mockTransferRepo{}, &mockUserRepo{})

	req := TransferRequest{FromUserID: 1, ToEmail: "testEmail", Amount: 0}

	resp, err := s.Transfer(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "amount must be greater than 0")
	assert.Empty(t, resp.FromName)
	assert.Empty(t, resp.ToName)
	assert.Equal(t, 0, resp.Amount)
	assert.Empty(t, resp.Status)
}

func TestTransferSenderNotFound(t *testing.T) {
	s := NewTransferService(
		&mockTransferRepo{},
		&mockUserRepo{failCase: "sender_not_found"},
	)

	req := TransferRequest{FromUserID: 99, ToEmail: "testEmail", Amount: 100}

	resp, err := s.Transfer(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "sender not found")
	assert.Empty(t, resp.FromName)
	assert.Empty(t, resp.ToName)
	assert.Equal(t, 0, resp.Amount)
	assert.Empty(t, resp.Status)
}

func TestTransferReceiverNotFound(t *testing.T) {
	s := NewTransferService(
		&mockTransferRepo{},
		&mockUserRepo{failCase: "receiver_not_found"},
	)

	req := TransferRequest{FromUserID: 1, ToEmail: testEmail, Amount: 100}

	resp, err := s.Transfer(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "receiver not found")
	assert.Empty(t, resp.FromName)
	assert.Empty(t, resp.ToName)
	assert.Equal(t, 0, resp.Amount)
	assert.Empty(t, resp.Status)
}

func TestTransferSelfTransfer(t *testing.T) {
	user := model.User{ID: 1, Name: "Alice"}

	s := NewTransferService(
		&mockTransferRepo{},
		&mockUserRepo{
			fromUser: user,
			toUser:   user,
		},
	)

	req := TransferRequest{FromUserID: 1, ToEmail: testEmail, Amount: 100}

	resp, err := s.Transfer(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot transfer to self")
	assert.Empty(t, resp.FromName)
	assert.Empty(t, resp.ToName)
	assert.Equal(t, 0, resp.Amount)
	assert.Empty(t, resp.Status)
}

func TestTransferInsufficientBalance(t *testing.T) {
	s := NewTransferService(
		&mockTransferRepo{failCase: "insufficient_balance"},
		&mockUserRepo{
			fromUser: model.User{ID: 1, Name: "Alice", Balance: 50},
			toUser:   model.User{ID: 2, Name: "Bob"},
		},
	)

	req := TransferRequest{FromUserID: 1, ToEmail: testEmail, Amount: 100}

	resp, err := s.Transfer(req)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "insufficient balance")
	assert.Empty(t, resp.FromName)
	assert.Empty(t, resp.ToName)
	assert.Equal(t, 0, resp.Amount)
	assert.Empty(t, resp.Status)
}
