package postgres

import (
	"database/sql"
	"fmt"
	"learning_app/models"
	"learning_app/pkg"
	"time"
)

type LessonRepo struct {
	DB *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{DB: db}
}

func (l *LessonRepo) CreateLesson(lesson models.Lesson) error {
	_, err := l.DB.Exec(`
		INSERT INTO lessons (
			lesson_id,
			course_id,
			title,
			content,
			created_at,
			updated_at
		) VALUES($1, $2, $3, $4, $5, $6)
	`, lesson.ID, lesson.CourseID, lesson.Title, lesson.Content, time.Now(), time.Now())

	return err
}

func (l *LessonRepo) GetLessonByID(id string) (models.Lesson, error) {
	var lesson models.Lesson

	err := l.DB.QueryRow(`
		SELECT
			lesson_id,
			course_id,
			title,
			content
		FROM lessons
		WHERE deleted_at = 0 AND lesson_id = $1 
	`, id).Scan(&lesson.ID, &lesson.CourseID, &lesson.Title, &lesson.Content)

	return lesson, err
}

func (l *LessonRepo) GetLessons() ([]models.Lesson, error) {
	var lessons []models.Lesson

	rows, err := l.DB.Query(`
	SELECT
		lesson_id,
		course_id,
		title,
		content
	FROM lessons
	WHERE deleted_at = 0
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var lesson models.Lesson

		err = rows.Scan(&lesson.ID, &lesson.CourseID, &lesson.Title, &lesson.Content)
		if err != nil {
			return nil, err
		}

		lessons = append(lessons, lesson)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lessons, nil
}

func (l *LessonRepo) UpdateLesson(lesson models.Lesson) error {
	params := make(map[string]interface{})
	var query = "UPDATE lessons SET "
	if lesson.CourseID != "" {
		query += "course_id = :course_id, "
		params["course_id"] = lesson.CourseID
	}
	if lesson.Title != "" {
		query += "title = :title, "
		params["title"] = lesson.Title
	}
	if lesson.Content != "" {
		query += "content = :content, "
		params["content"] = lesson.Content
	}
	query = "updated_at = CURRENT_TIMESTAMP WHERE id = :id AND deleted_at = 0"
	params["id"] = lesson.ID

	query, args := pkg.ReplaceQueryParams(query, params)

	_, err := l.DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (l *LessonRepo) DeleteLesson(id string) error {
	res, err := l.DB.Exec(`
		UPDATE lessons 
			SET deleted_at = DATE_PART('epoch', CURRENT_TIMESTAMP)::INT 
		WHERE id = $1 AND deleted_at = 0
	`, time.Now(), id)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected, user with id %s not found or already deleted", id)
	}

	return nil
}

func (l *LessonRepo) GetAllLessons(fLesson models.FilterLesson) ([]models.Lesson, error) {
	var (
		params = make(map[string]interface{})
		args    []interface{}
		filter string
	)

	query := "SELECT lesson_id, course_id, title, content FROM users WHERE deleted_at = 0 "

	if fLesson.CourseID != "" {
		params["course_id"] = fLesson.CourseID
		filter += "AND course_id = :course_id "
	}
	if fLesson.Title != "" {
		params["title"] = fLesson.Title
		filter += "AND title = :title "
	}
	if fLesson.Content != "" {
		params["birthday"] = fLesson.Content
		filter += "AND content = :content"
	}
	if fLesson.Offset > 0 {
		params["offset"] = fLesson.Offset
		filter += "AND OFFSET = :offset "
	}
	if fLesson.Limit > 0 {
		params["limit"] = fLesson.Limit
		filter += "AND LIMIT = :limit "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)

	rows, err := l.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var lessons []models.Lesson
	for rows.Next() {
		var lesson models.Lesson

		err = rows.Scan(&lesson.ID, &lesson.CourseID, &lesson.Title, &lesson.Content)

		if err != nil {
			return nil, err
		}

		lessons = append(lessons, lesson)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lessons, err
}