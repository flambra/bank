package transaction

import (
	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func Page(c *fiber.Ctx) error {
	var response []domain.TransactionResponse
	var transactions []domain.Transaction
	var filter domain.TransactionPageFilter

	transactionRepo := hRepository.New(hDb.Get(), &transactions, c)
	paginator := hRepository.BuildPaginator(&response)

	err := c.QueryParser(paginator)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = c.QueryParser(&filter)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = transactionRepo.FindAllPaginating(&filter, paginator, clause.Associations)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, paginator)
}
