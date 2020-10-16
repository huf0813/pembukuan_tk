package usecase

import (
	"github.com/huf0813/pembukuan_tk/repository"
)

type UserUseCase struct {
	UserRepo repository.UserRepo
}

type UserUseCaseInterface interface {
	FetchWithoutID() (interface{}, error)
}

func (uus *UserUseCase) FetchWithoutID() (interface{}, error) {
	result, err := uus.UserRepo.FetchAll()
	if err != nil {
		return nil, err
	}

	resultsWithoutID := []struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		{},
	}

	for _, r := range result {
		resultsWithoutID = append(resultsWithoutID, struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{Name: r.Name, Age: r.Age})
	}

	return resultsWithoutID[1:], nil
}
