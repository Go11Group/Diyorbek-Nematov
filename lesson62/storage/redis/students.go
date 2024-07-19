package redis

import (
	"context"
	"encoding/json"
	"students/models"

	"github.com/redis/go-redis/v9"
)

type StudentRepo struct {
	RDB *redis.Client
}

func NewStudentRepo(rdb redis.Client) *StudentRepo {
	return &StudentRepo{
		RDB: &rdb,
	}
}

var ctx = context.Background()

func (repo *StudentRepo) CreateStudent(student models.Student) (*models.Success, error) {
	jsonData, err := json.Marshal(student)
	if err != nil {
		return nil, err
	}

	if err := repo.RDB.Set(ctx, "student:"+student.ID, string(jsonData), 0).Err(); err != nil {
		return nil, err
	}

	err = repo.RDB.SAdd(ctx, "students", student.ID).Err()
	if err != nil {
		return nil, err
	}

	return &models.Success{
		Message: "Student created successfully",
	}, nil
}

func (repo *StudentRepo) GetStudents() ([]models.Student, error) {
	// barcha student ID-larni olish
	studentIDs, err := repo.RDB.SMembers(ctx, "students").Result()
	if err != nil {
		return nil, err
	}

	// barcha studentlarni olish
	var allStudents []models.Student

	for _, id := range studentIDs {
		val, err := repo.RDB.Get(ctx, "student:"+id).Result()
		if err != nil {
			return nil, err
		}

		var student models.Student

		err = json.Unmarshal([]byte(val), &student)
		if err != nil {
			return nil, err
		}

		allStudents = append(allStudents, student)
	}

	return allStudents, nil
}

func (repo *StudentRepo) GetStudent(id string) (*models.Student, error) {
	val, err := repo.RDB.Get(ctx, "studenet:"+id).Result()
	if err != nil {
		return nil, err
	}

	var student models.Student

	err = json.Unmarshal([]byte(val), &student)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (repo *StudentRepo) UpdateStudent(student models.Student) (*models.Success, error) {
	jsonData, err := json.Marshal(student)
	if err != nil {
		return nil, err
	}

	if err := repo.RDB.Set(ctx, "student:"+student.ID, string(jsonData), 0).Err(); err != nil {
		return nil, err
	}

	return &models.Success{
		Message: "Student updated successfully",
	}, nil
}

func (repo *StudentRepo) DeleteStudent(id string) (*models.Success, error) {
	err := repo.RDB.Del(ctx, "student:"+id).Err()
	if err != nil {
		return nil, err
	}

	err = repo.RDB.SRem(ctx, "students", id).Err()
	if err != nil {
		return nil, err
	}

	return &models.Success{
		Message: "Student deleted successfully",
	}, nil
}

func (repo *StudentRepo) GetSubjectsForStudent(studentID string) ([]models.Subject, error) {
    subjectIDs, err := repo.RDB.SMembers(ctx, "student:"+studentID+":subjects").Result()
    if err != nil {
        return nil, err
    }

    var subjects []models.Subject
    for _, id := range subjectIDs {
        val, err := repo.RDB.Get(ctx, "subject:"+id).Result()
        if err != nil {
            return nil, err
        }

        var subject models.Subject
        err = json.Unmarshal([]byte(val), &subject)
        if err != nil {
            return nil, err
        }

        subjects = append(subjects, subject)
    }

    return subjects, nil
}