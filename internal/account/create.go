package account

import (
	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	var account domain.Account
	var request domain.AccountCreateRequest
	accountRepo := hRepository.New(hDb.Get(), &account, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	account = domain.Account{
		Account:       request.Account,
		Balance:       request.Balance,
		Agency:        request.Agency,
		Digit:         request.Digit,
		Favorite:      request.Favorite,
		Owner:         request.Owner,
		TypeAccount:   request.TypeAccount,
		InstitutionID: request.InstitutionID,
	}

	err := accountRepo.Create()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessCreated(c, &account)
}
