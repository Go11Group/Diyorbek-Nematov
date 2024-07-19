package redis

import (
	"encoding/json"
	"students/models"

	"github.com/redis/go-redis/v9"
)

type SubjectRepo struct {
	RDB *redis.Client
}

func NewSubjectRepo(rdb redis.Client) *SubjectRepo {
	return &SubjectRepo{
		RDB: &rdb,
	}
}

func (repo *SubjectRepo) CreateSubject(subject *models.Subject) (*models.Success, error) {
	jsonData, err := json.Marshal(subject)
	if err != nil {
		return nil, err
	}

	if err := repo.RDB.Set(ctx, "subject:"+subject.ID, jsonData, 0).Err(); err != nil {
		return nil, err
	}

	err = repo.RDB.SAdd(ctx, "subjects", subject.ID).Err()
	if err != nil {
		return nil, err
	}

	return &models.Success{
		Message: "Subject created successfully",
	}, nil
}

func (repo *SubjectRepo) GetSubjects() ([]models.Subject, error) {
	// barcha student ID-larni olish
	subjectIDs, err := repo.RDB.SMembers(ctx, "subjects").Result()
	if err != nil {
		return nil, err
	}

	// barcha subjectlarni olish
	var allSubjects []models.Subject

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

		allSubjects = append(allSubjects, subject)
	}

	return allSubjects, nil
}

func (repo *SubjectRepo) GetSubject(id string) (*models.Subject, error) {
	val, err := repo.RDB.Get(ctx, "subject:"+id).Result()
	if err != nil {
		return nil, err
	}

	var subject models.Subject

	err = json.Unmarshal([]byte(val), &subject)
	if err != nil {
		return nil, err
	}

	return &subject, nil
}

func (repo *SubjectRepo) UpdateSubject(subject models.Subject) (*models.Success, error) {
	jsonData, err := json.Marshal(subject)
	if err != nil {
		return nil, err
	}

	if err := repo.RDB.Set(ctx, "subject:"+subject.ID, string(jsonData), 0).Err(); err != nil {
		return nil, err
	}

	return &models.Success{
		Message: "Subject updated successfully",
	}, nil
}

func (repo *StudentRepo) DeleteSubject(id string) (*models.Success, error) {
	err := repo.RDB.Del(ctx, "subject:"+id).Err()
	if err != nil {
		return nil, err
	}

	err = repo.RDB.SRem(ctx, "subjects", id).Err()
	if err != nil {
		return nil, err
	}

	return &models.Success{
		Message: "Subject deleted successfully",
	}, nil
}

func (repo *SubjectRepo) AssociateStudentWithSubject(associate models.Associate) (*models.Success, error) {
    err := repo.RDB.SAdd(ctx, "student:"+associate.StudentID+":subjects", associate.SubjectID).Err()
    if err != nil {
        return nil, err
    }

    err = repo.RDB.SAdd(ctx, "subject:"+associate.SubjectID+":students", associate.StudentID).Err()
    if err != nil {
        return nil, err
    }

	return &models.Success{
		Message: "Sutudent added to subject successfully",
	}, nil
}

func (repo *SubjectRepo) GetStudentsForSubject(subjectID string) ([]models.Student, error) {
    studentIDs, err := repo.RDB.SMembers(ctx, "subject:"+subjectID+":students").Result()
    if err != nil {
        return nil, err
    }

    var students []models.Student
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

        students = append(students, student)
    }

    return students, nil
}