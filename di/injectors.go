package di

import (
	"os"
	"goa_starter/database"
	"goa_starter/domain/repository"
	"goa_starter/infra"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func InjectDb() (*gorm.DB, error) {
	// load .env
	stage := os.Getenv("STAGE")
	if stage == "local" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	cs, err := database.NewConfigsFromEnv(stage)
	if err != nil {
		return nil, err
	}

	db, dErr := cs.Open()
	if dErr != nil {
		return nil, dErr
	}

	return db, nil
}

var DB, _ = InjectDb()

func InjectTermLimitRepository() repository.TermLimitRepository {
	return infra.NewTermLimitRepository(DB)
}

// func InjectConversionPointRepository() repository.ConversionPointRepository {
// 	return infra.NewConversionPointRepository(DB)
// }

// func InjectEndUserRepository() repository.EndUserRepository {
// 	return infra.NewEndUserRepository(DB)
// }

// func InjectMediaUserRepository() repository.MediaUserRepository {
// 	return infra.NewMediaUserRepository(DB)
// }

// func InjectMediaRepository() repository.MediaRepository {
// 	return infra.NewMediaRepository(DB)
// }

// func InjectUserOfferRepository() repository.UserOfferRepository {
// 	return infra.NewUserOfferRepository(DB)
// }