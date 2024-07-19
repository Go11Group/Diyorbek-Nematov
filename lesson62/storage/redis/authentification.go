package redis

import (
	"students/models"

	"github.com/redis/go-redis/v9"
)

type AuthentificationRepo struct {
	RDB *redis.Client
}

func NewAuthentificationRepo(rdb *redis.Client) *AuthentificationRepo {
	return &AuthentificationRepo{
		RDB: rdb,
	}
}

func (repo *AuthentificationRepo) Register(user models.Register) (*models.Success, *models.Errors) {
	exists, err := repo.RDB.Exists(ctx, user.Username).Result()
	if err != nil {
		return nil, &models.Errors{
			Error: err.Error(),
		}
	}

	if exists == 1 {
		return nil, &models.Errors{
			Error: "this username already exists",
		}
	}

	err = repo.RDB.Set(ctx, user.Username, user.Password, 0).Err()

	if err != nil {
		return nil, &models.Errors{
			Error: err.Error(),
		}
	}

	return &models.Success{
		Message: "user registered successfully",
	}, nil
}

func (repo *AuthentificationRepo) Login(user models.Register) (string, *models.Errors) {
	exists, err := repo.RDB.Exists(ctx, user.Username).Result()
	if err != nil {
		return "", &models.Errors{
			Error: err.Error(),
		}
	}

	if exists != 1 {
		return "", &models.Errors{
			Error: "this username dose not exist",
		}
	}

	val, err := repo.RDB.Get(ctx, user.Username).Result()
	if err != nil {
		return "", &models.Errors{
			Error: err.Error(),
		}
	}

	return val, nil
}
