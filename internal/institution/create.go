package institution

import (
	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

type InstitutionCreateRequest struct {
	Code int    `json:"code"`
	Name string `json:"name"`
	CNPJ string `json:"cnpj"`
	Logo string `json:"logo"`
}

func Create(c *fiber.Ctx) error {

	var institution domain.Institution
	var request InstitutionCreateRequest
	institutionRepo := hRepository.New(hDb.Get(), &institution, c)

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var count int64
	db := hDb.Get()
	db.Model(&domain.Institution{}).Where("code = ? or cnpj = ?", request.Code, request.CNPJ).Count(&count)
	if count > 0 {
		return hResp.StatusConflict(c, &institution, "Code or CNPJ already in use")
	}

	institution = domain.Institution{
		Code: request.Code,
		Name: request.Name,
		CNPJ: request.CNPJ,
		Logo: request.Logo,
	}

	err := institutionRepo.Create()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessCreated(c, &institution)
}
