package transaction

import (
	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Page(c *fiber.Ctx) error {
	var response []domain.TransactionPageResponse
	var transactions []domain.Transaction
	var filter domain.TransactionPageFilter

	paginator := hRepository.BuildPaginator(&response)
	transactionRepo := hRepository.New(hDb.Get(), &transactions, c)

	err := c.QueryParser(paginator)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = c.QueryParser(&filter)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	err = transactionRepo.FindAllPaginating(&filter, paginator)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, paginator)
}
