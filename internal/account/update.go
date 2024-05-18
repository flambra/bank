package account

import (
	"strconv"

	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

type AccountUpdateRequest struct {
	Account       int    `json:"account"`
	Balance       int    `json:"balance"`
	Agency        int    `json:"agency"`
	Digit         int    `json:"digit"`
	Favorite      bool   `json:"favorite"`
	Owner         string `json:"owner"`
	TypeAccount   string `json:"type_account"`
	InstitutionID int    `json:"institution_id"`
}

func Update(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var account domain.Account
	accountRepo := hRepository.New(hDb.Get(), &account, c)

	err = accountRepo.GetById(id)
	if err != nil {
		return hResp.NotFoundResponse(c, err.Error(), "account not found")
	}

	var request AccountUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	updateAccount(&account, request)

	err = accountRepo.Save()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, &account)
}

func updateAccount(account *domain.Account, request AccountUpdateRequest) {
	if request.Account != 0 {
		account.Account = request.Account
	}
	if request.Balance != 0 {
		account.Balance = request.Balance
	}
	if request.Agency != 0 {
		account.Agency = request.Agency
	}
	if request.Digit != 0 {
		account.Digit = request.Digit
	}
	if request.Favorite {
		account.Favorite = request.Favorite
	}
	if request.Owner != "" {
		account.Owner = request.Owner
	}
	if request.TypeAccount != "" {
		account.TypeAccount = request.TypeAccount
	}
	if request.InstitutionID != 0 {
		account.InstitutionID = request.InstitutionID
	}
}
