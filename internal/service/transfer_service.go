package service

import (
	"errors"

	"github.com/guide-backend/internal/repository"
)

type TransferService struct {
	transferRepo repository.TransferRepo
	userRepo     repository.UserRepo
}

func NewTransferService(transferRepo repository.TransferRepo, userRepo repository.UserRepo) TransferService {
	return TransferService{
		transferRepo: transferRepo,
		userRepo:     userRepo,
	}
}

func (s TransferService) Transfer(req TransferRequest) (TransferResponse, error) {
	if req.Amount <= 0 {
		return TransferResponse{}, errors.New("amount must be greater than 0")
	}

	fromUser, err := s.userRepo.GetUserByID(req.FromUserID)
	if err != nil {
		return TransferResponse{}, errors.New("sender not found")
	}

	toUser, err := s.userRepo.GetUserByEmail(req.ToEmail)
	if err != nil {
		return TransferResponse{}, errors.New("receiver not found")
	}

	if fromUser.ID == toUser.ID {
		return TransferResponse{}, errors.New("cannot transfer to self")
	}

	if fromUser.Balance < req.Amount {
		return TransferResponse{}, errors.New("insufficient balance")
	}

	err = s.transferRepo.TransferBalanceAtomic(fromUser.ID, toUser.ID, req.Amount)
	if err != nil {
		return TransferResponse{}, err
	}

	return TransferResponse{
		FromName: fromUser.Name,
		ToName:   toUser.Name,
		Amount:   req.Amount,
		Status:   "success",
	}, nil
}
