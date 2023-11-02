package setup

import (
	"fmt"

	"github.com/muhrizqiardi/spendtracker/internal/database/model"
	"github.com/muhrizqiardi/spendtracker/internal/database/seed"
	"github.com/muhrizqiardi/spendtracker/internal/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupMigrateAndSeedMySQL(cfg util.Config, lg util.Logger) (*gorm.DB, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DB_Username,
		cfg.DB_Password,
		cfg.DB_Host,
		cfg.DB_Port,
		cfg.DB_Name,
	)

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		lg.Error("Failed to open connection", err)
		return nil, err
	}

	if err := db.AutoMigrate(
		&model.User{},
		&model.Account{},
		&model.Category{},
		&model.Currency{},
		&model.Expense{},
	); err != nil {
		lg.Error("Failed to migrate", err)
		return nil, err
	}

	if err := seed.Seed(db, lg); err != nil {
		return nil, err
	}

	return db, nil
}
