package service

import (
	"fmt"

	"github.com/muhrizqiardi/spendtracker/internal/repository"
)

const Prompt string = "Given the maximum of 20 expenses consists of name, description, and the amount, give me a financial advice based on that, in two sentence maximum."

type AdviceService interface {
	GetAdvice(userID int) (string, error)
}

type adviceService struct {
	es  ExpenseService
	oar repository.OpenAIRepository
}

func NewAdviceService(es ExpenseService, oar repository.OpenAIRepository) *adviceService {
	return &adviceService{es, oar}
}

func (ads *adviceService) GetAdvice(userID int) (string, error) {
	expenses, err := ads.es.GetManyBelongedToUser(userID, 20, 1)
	if err != nil {
		return "", err
	}

	message := `My last expenses were:
`
	for _, e := range expenses {
		message += fmt.Sprintf(`- name: %s, description: %s, amount: %d`, e.Name, e.Description, e.Amount)
	}

	response, err := ads.oar.GetResponse(Prompt, message)
	if err != nil {
		return "", err
	}

	return response, nil
}
