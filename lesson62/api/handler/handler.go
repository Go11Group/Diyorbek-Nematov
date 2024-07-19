package handler

import (
	rdb "students/storage/redis"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	StudentRepo      *rdb.StudentRepo
	SubjectRepo      *rdb.SubjectRepo
	Authentification *rdb.AuthentificationRepo
	Enforcer         *casbin.Enforcer
}

func NewHandler(db *redis.Client, enforcer *casbin.Enforcer) *Handler {
	return &Handler{
		StudentRepo:      rdb.NewStudentRepo(db),
		SubjectRepo:      rdb.NewSubjectRepo(db),
		Authentification: rdb.NewAuthentificationRepo(db),
		Enforcer:         enforcer,
	}
}
