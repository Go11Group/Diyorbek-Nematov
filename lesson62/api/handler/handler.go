package handler

import (
	rdb "students/storage/redis"

	"github.com/redis/go-redis/v9"
)

type Handler struct {
	StudentRepo *rdb.StudentRepo
	SubjectRepo *rdb.SubjectRepo
}

func NewHandler(db redis.Client) *Handler {
	return &Handler{
		StudentRepo: rdb.NewStudentRepo(db),
		SubjectRepo: rdb.NewSubjectRepo(db),
	}
}
