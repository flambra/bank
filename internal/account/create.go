package account

import (
	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

type AccountCreateRequest struct {
	Account       int    `json:"account"`
	Balance       int    `json:"balance"`
	Agency        int    `json:"agency"`
	Digit         int    `json:"digit"`
	Favorite      bool   `json:"favorite"`
	Owner         string `json:"owner"`
	TypeAccount   string `json:"type_account"`
	InstitutionID int    `json:"institution_id"`
}

func Create(c *fiber.Ctx) error {

	var account domain.Account
	var request AccountCreateRequest
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
