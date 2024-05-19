package institution

import (
	"github.com/flambra/bank/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/flambra/helpers/hRepository"
	"github.com/flambra/helpers/hResp"
	"github.com/gofiber/fiber/v2"
)

func List(c *fiber.Ctx) error {
	var institutions []domain.Institution
	institutionRepo := hRepository.New(hDb.Get(), &institutions, c)

	err := institutionRepo.FindAllWhere(nil)
	if err != nil {
		return hResp.InternalServerErrorResponse(c, err.Error())
	}

	var response []domain.InstitutionListResponse
	for _, institution := range institutions {
		response = append(response, domain.InstitutionListResponse{
			ID:   institution.ID,
			Name: institution.Name,
			Code: institution.Code,
			Logo: institution.Logo,
		})
	}

	return hResp.SuccessResponse(c, response)
}
