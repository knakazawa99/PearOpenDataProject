package presenter

import (
	"api/domain/entity"
	"api/http/response"
)

type PearVersion interface {
	OutPutPearVersions([]entity.Pear) []response.PearDataVersionOutput
	OutPutPearAdminVersions([]entity.Pear) []response.PearAdminDataVersionOutput
}

type pearVersionImplement struct {
}

func (p pearVersionImplement) OutPutPearVersions(pears []entity.Pear) []response.PearDataVersionOutput {
	pearDataVersionOutputs := make([]response.PearDataVersionOutput, len(pears))
	for i := range pears {
		pearDataVersionOutputs[i] = response.PearDataVersionOutput{
			Version:     pears[i].Version,
			ReleaseNote: pears[i].ReleaseNote,
			CreatedAt:   pears[i].CreatedAt,
		}
	}
	return pearDataVersionOutputs
}

func (p pearVersionImplement) OutPutPearAdminVersions(pears []entity.Pear) []response.PearAdminDataVersionOutput {
	pearDataVersionOutputs := make([]response.PearAdminDataVersionOutput, len(pears))
	for i := range pears {
		pearDataVersionOutputs[i] = response.PearAdminDataVersionOutput{
			ID:             pears[i].ID,
			FilePath:       pears[i].FilePath,
			Version:        pears[i].Version,
			ReleaseNote:    pears[i].ReleaseNote,
			ReleaseComment: pears[i].ReleaseComment,
			ReleaseFlag:    pears[i].ReleaseFlag,
			CreatedAt:      pears[i].CreatedAt,
		}
	}
	return pearDataVersionOutputs
}

func NewPearVersion() PearVersion {
	return pearVersionImplement{}
}
