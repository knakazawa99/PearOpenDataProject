package presenter

import (
	"time"

	"api/domain/entity"
)

type PearDataVersionOutput struct {
	Version     string    `json:"version"`
	ReleaseNote string    `json:"release_note"`
	CreatedAt   time.Time `json:"created_at"`
}

type PearVersion interface {
	OutPutPearVersions([]entity.Pear) []PearDataVersionOutput
}

type pearVersionImplement struct {
}

func (p pearVersionImplement) OutPutPearVersions(pears []entity.Pear) []PearDataVersionOutput {
	pearDataVersionOutputs := make([]PearDataVersionOutput, len(pears))
	for i := range pears {
		pearDataVersionOutputs[i] = PearDataVersionOutput{
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
