package transaction

import (
	"github.com/flambra/bank/internal/balance"
	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var transaction domain.Transaction
	var request domain.TransactionCreateRequest
	transactionRepo := hRepository.New(hDb.Get(), &transaction, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	transaction = domain.Transaction{
		Amount:      request.Amount,
		PayerID:     request.PayerID,
		RecieverID:  request.RecieverID,
		Description: request.Description,
	}

	if err := balance.Update(request.PayerID, -transaction.Amount); err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	if err := balance.Update(request.RecieverID, transaction.Amount); err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	err := transactionRepo.Create()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessCreated(c, &transaction)
}
