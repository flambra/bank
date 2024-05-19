package institution

import (
	"strconv"

	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	rawId := c.Params("id")
	if rawId == "" {
		return hResp.BadRequestResponse(c, "inform id")
	}

	id, err := strconv.Atoi(rawId)
	if err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	var institution domain.Institution
	var request domain.InstitutionUpdateRequest
	institutionRepo := hRepository.New(hDb.Get(), &institution, c)

	err = institutionRepo.GetById(id)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	if err := c.BodyParser(&request); err != nil {
		return hResp.BadRequestResponse(c, err.Error())
	}

	updateInstitution(&institution, request)

	err = institutionRepo.Save()
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	return hResp.SuccessResponse(c, &institution)
}

func updateInstitution(institution *domain.Institution, request domain.InstitutionUpdateRequest) {
	if request.Code != 0 {
		institution.Code = request.Code
	}
	if request.Name != "" {
		institution.Name = request.Name
	}
	if request.CNPJ != "" {
		institution.CNPJ = request.CNPJ
	}
	if request.Logo != "" {
		institution.Logo = request.Logo
	}
}
