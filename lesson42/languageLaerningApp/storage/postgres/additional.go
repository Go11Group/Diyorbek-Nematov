package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"learning_app/models"
	"learning_app/pkg"
)

type AdditionalRepo struct {
	DB *sql.DB
}

// Yangi AdditionalRepo yaratish
func NewAdditionalRepo(db *sql.DB) *AdditionalRepo {
	return &AdditionalRepo{DB: db}
}

// Ma'lum bir foydalanuvchiga tegishli barcha kurslarni olish.
func (a *AdditionalRepo) GetCoursesbyUser(id string) (*models.CourseByUser, error) {
	var coursesByUser models.CourseByUser

	// databasesga so'rov
	rows, err := a.DB.Query(`
		SELECT
			u.user_id,
			c.course_id,
			c.title,
			c.description
		FROM
			users AS u 
		INNER JOIN 
			enrollments AS e ON u.user_id = e.user_id
		INNER JOIN
			courses AS c ON e.course_id = c.course_id
		WHERE 
			u.deleted_at = 0 AND c.deleted_at = 0 AND e.deleted_at = 0  AND u.user_id = $1
	`, id)

	// agar databasesdan o'qishda xatolik bo'lsa tekshirish
	if err != nil {
		return nil, err
	}

	// o'zgaruchilarni qabul qilish
	var courses []models.Course
	for rows.Next() {
		var course models.Course

		err = rows.Scan(&coursesByUser.UserID, &course.ID, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	coursesByUser.Courses = courses
	return &coursesByUser, err
}

//Ma'lum bir kursga tegishli barcha darslarni olish.
func (a *AdditionalRepo) GetLessonsByCourse(id string) (*models.LessonsByCourse, error) {
	var lessonsByCourse models.LessonsByCourse

	// databasesga query 
	rows, err := a.DB.Query(`
		SELECT 
			c.course_id,
			lesson_id,
			l.title,
			l.content
		FROM
			courses AS c 
		INNER JOIN 
			lessons as l ON c.course_id = l.course_id
		WHERE 
			c.deleted_at = 0 AND l.deleted_at = 0 AND c.course_id = $1
	`, id)

	if err != nil {
		return nil, err
	}

	// 	o'zgaruvchilarni qabul qilish
	var addLessons []models.AdditialLesson
	for rows.Next() {
		var addLesson models.AdditialLesson

		err = rows.Scan(&lessonsByCourse.CourseID, &addLesson.ID,
			&addLesson.Title, &addLesson.Content)
		if err != nil {
			return nil, err
		}

		addLessons = append(addLessons, addLesson)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	lessonsByCourse.Lessons = addLessons
	return &lessonsByCourse, nil
}

//Ma'lum bir kursga ro'yxatdan o'tgan barcha foydalanuvchilarni olish.
func (a *AdditionalRepo) GetEnrolledUsersbyCourse(id string) ([]models.EnrolledUsersByCourse, error) {
	var enrollUsersByCourse []models.EnrolledUsersByCourse
	// databasedan bir kursga royxatdan o'tgan barcha foydalanuvchilarni olish uchu so'rov
	rows, err := a.DB.Query(`
		SELECT 
			c.course_id,
			u.user_id,
			u.name,
			u.email
		FROM
			courses AS c 
		INNER JOIN 
			enrollments AS e ON c.course_id = e.course_id
		INNER JOIN 
			users AS u ON e.user_id = u.user_id
		WHERE 
			c.deleted_at = 0 AND u.deleted_at = 0 AND e.deleted_at = 0 AND c.course_id = $1
	`, id)

	if err != nil {
		return nil, err
	}
	// rows dan o'zgaruvchilarni scan qilish
	for rows.Next() {
		var eUBCourse models.EnrolledUsersByCourse

		err = rows.Scan(&eUBCourse.CourseID, &eUBCourse.EnrolledUsers.ID, &eUBCourse.EnrolledUsers.Name,
			&eUBCourse.EnrolledUsers.Email)

		if err != nil {
			return nil, err
		}

		enrollUsersByCourse = append(enrollUsersByCourse, eUBCourse)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return enrollUsersByCourse, nil
}

// Foydalanuvchilarni ismi, yoshi yoki email bo'yicha qidirish. Yosh oralig’i bo’yicha qidirish
func (a *AdditionalRepo) SearchUsers(fSearch models.SearchUser) (*models.Result, error) {
	var (
		params = make(map[string]interface{})
		args   []interface{}
		filter string
	)

	query := "SELECT user_id, name, email FROM users WHERE deleted_at = 0 "

	// o'zgaruvchilarni zero valuega tekshirish filterlash uchun kerakli o'zgaruvchilarni tanlash
	if fSearch.Name != "" {
		params["name"] = fSearch.Name
		filter += "AND name = :name "
	}
	if fSearch.Email != "" {
		params["email"] = fSearch.Email
		filter += "AND email = :email "
	}
	if fSearch.AgeTo > 0 && fSearch.AgeFrom > 0 {
		params["age_to"] = fSearch.AgeTo
		params["age_from"] = fSearch.AgeFrom

		filter += "AND EXTRACT(YEAR FROM age(birthday)) between :age_from AND :age_to "
	}

	query += filter

	query, args = pkg.ReplaceQueryParams(query, params)
	fmt.Println(query)

	// databasesga so'rov yuborish
	rows, err := a.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

		// databasedan kelgan ma'lumotlarni o'zgaruchilarga olish
	var results []models.AdditialUser
	for rows.Next() {
		var result models.AdditialUser

		err = rows.Scan(&result.ID, &result.Name, &result.Email)

		if err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &models.Result{Results: results}, err
}

// Ma'lum bir muddat davomida eng ko'p ro'yxatdan o'tilgan kurslarni olish.
func (a *AdditionalRepo) GetMostPopularCourses(startDate, endDate string) (models.Goal, error) {
	// Agar malum muddat uchun boshlang'ich va oxirgi kunlarning ikkalasi ham kelmasa error qaytarish
	if startDate == "" || endDate == "" {
		return models.Goal{}, errors.New("start_date or end_date not found")
	}

	// query yuborish
	rows, err := a.DB.Query(`
		SELECT 
			c.course_id,
			c.title,
			COUNT(c.course_id) AS enrollment_count
		FROM 
			courses AS c 
		INNER JOIN
			enrollments AS e ON c.course_id = e.course_id
		WHERE 
			e.enrollment_date BETWEEN $1 AND $2
		GROUP BY
			c.course_id, c.title
		HAVING
			COUNT(c.course_id) = (
				SELECT 
					MAX(enroll_count) 
				FROM (
					SELECT 
						COUNT(e.course_id) AS enroll_count
					FROM 
						courses AS c 
					INNER JOIN
						enrollments AS e ON c.course_id = e.course_id
					GROUP BY
						c.course_id
				) AS counts
			)
	`, startDate, endDate)

	// error ga tekshirish
	if err != nil {
		return models.Goal{}, nil
	}
	// o'zgaruvchilarni qabul qilish
	var popularCourses []models.PopularCourse

	for rows.Next() {
		var popularCourse models.PopularCourse
		err = rows.Scan(&popularCourse.CourseID, &popularCourse.CourseTitle, &popularCourse.EnrollmentsCount)
		if err != nil {
			return models.Goal{}, nil
		}

		popularCourses = append(popularCourses, popularCourse)
	}

	var timeRange = models.TimeRange{
		StartDate: startDate,
		EndDate:   endDate,
	}

	return models.Goal{TimePeriod: timeRange, PopularCourse: popularCourses}, nil
}
