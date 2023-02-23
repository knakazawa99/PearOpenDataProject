package presenter

import (
	"api/domain/entity"
	"api/http/response"
)

type PearVersion interface {
	OutPutPearVersions([]entity.Pear) []response.PearDataVersionOutput
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

func NewPearVersion() PearVersion {
	return pearVersionImplement{}
}
