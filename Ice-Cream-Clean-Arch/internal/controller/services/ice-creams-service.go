package services

import (
	icecreamrepo "ice-cream-app/internal/controller/repositories/ice-cream-repo"
	"ice-cream-app/internal/models"
)

type IceCreamService struct {
	IceCreamRepo icecreamrepo.IceCreamDbStore
}

func NewIceCreamService(repo icecreamrepo.IceCreamDbStore) IceCreamService {
	return IceCreamService{
		IceCreamRepo: repo,
	}
}

// get all ice creams
func (s IceCreamService) GetIceCreams(params models.QueryParams) ([]models.IceCream, error) {
	return s.IceCreamRepo.GetIceCreams(params)
}

// get ice cream by ID
func (s IceCreamService) GetIceCreamByID(paramID string) (models.IceCream, error) {
	return s.IceCreamRepo.GetIceCreamByID(paramID)
}

// detete ice cream
func (s IceCreamService) DeleteIceCream(paramID string) error {
	return s.IceCreamRepo.DeleteIceCream(paramID)
}

// post ice cream
func (s IceCreamService) PostIceCream(iceCream models.IceCreamPost) (int64, error) {
	err := iceCream.Validate()
	if err != nil {
		return 0, err
	}
	return s.IceCreamRepo.PostIceCream(iceCream)
}

// update ice cream
func (s IceCreamService) PutIceCream(paramID int, iceCream *models.IceCream, iceCreamPut *models.IceCreamPost) error {

	validIceCream, err := iceCreamPut.ValidUpdate(iceCream)
	if err != nil {
		return err
	}
	return s.IceCreamRepo.PutIceCream(paramID, *validIceCream)
}
