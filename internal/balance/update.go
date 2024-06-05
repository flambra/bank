package balance

import (
	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
)

func Update(accountID int, amount int) error {
	var balance domain.BalanceCurrent
	balanceRepo := hRepository.New(hDb.Get(), &balance, nil)

	err := balanceRepo.GetById(accountID)
	if err != nil {
		return err
	}

	balance.Amount += amount
	err = balanceRepo.Create()
	if err != nil {
		return hResp.InternalServerErrorResponse(nil, err.Error())
	}

	err = saveBalanceHistorical(&balance)
	if err != nil {
		return hResp.InternalServerErrorResponse(nil, err.Error())
	}

	return nil
}

func saveBalanceHistorical(balanceCurrent *domain.BalanceCurrent) error {
	balanceHistorical := domain.BalanceHistorical{
		BalanceID: balanceCurrent.ID,
		AccountID: balanceCurrent.AccountID,
		Amount:    balanceCurrent.Amount,
		CreatedAt: balanceCurrent.UpdatedAt,
	}

	historyRepo := hRepository.New(hDb.Get(), &balanceHistorical, nil)
	err := historyRepo.Create()
	if err != nil {
		return err
	}

	return nil
}
