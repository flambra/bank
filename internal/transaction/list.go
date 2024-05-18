package transaction

import (
	"strconv"

	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var transactions []domain.Transaction
	transactionRepo := hRepository.New(hDb.Get(), &transactions, c)

	err = transactionRepo.FindAllWhere(map[string]interface{}{
		"payer_id":    id,
		"reciever_id": id,
	})
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	var transactionList []domain.TransactionResponse
	for _, transaction := range transactions {
		transactionList = append(transactionList, domain.TransactionResponse{
			Amount:      transaction.Amount,
			PayerID:     transaction.PayerID,
			RecieverID:  transaction.RecieverID,
			Description: transaction.Description,
		})
	}

	return hResp.SuccessResponse(c, transactionList)
}
