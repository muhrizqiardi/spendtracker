package seed

import (
	_ "embed"
	"encoding/csv"
	"io"
	"strings"

	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/util"
	"gorm.io/gorm"
)

//go:embed seed_currencies.csv
var currenciesCSV string

func Seed(db *gorm.DB, lg util.Logger) error {
	r := csv.NewReader(strings.NewReader(currenciesCSV))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			lg.Error("Parsing currencies CSV failed", err)
			return err
		}

		currency := model.Currency{
			Model: gorm.Model{},
			Code:  record[1],
		}
		if err := db.Save(&currency).Error; err != nil {
			lg.Error("Inserting currency failed", err)
			return err
		}
	}

	return nil
}
