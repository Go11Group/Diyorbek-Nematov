package postgres

import (
	"database/sql"
	"solution/models"
)

type ProblemRepo struct {
	DB *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{DB: db}
}

func (p *ProblemRepo) CreateProblem(problem models.Problem) error {
	_, err := p.DB.Exec(`
		INSERT INTO problems (id, user_id, title, description) 
		VALUES ($1, $2, $3, $4)
	`, problem.UserID, problem.Title, problem.Description)
	return err
}

func (p *ProblemRepo) GetAllProblems() ([]models.Problem, error) {
	var problems []models.Problem

	rows, err := p.DB.Query(`
		SELECT id, problem_id, user_id, solution FROM solved_problems
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var problem models.Problem
		
		err := rows.Scan(&problem.ID, &problem.UserID, &problem.Title, &problem.Description)

		if err != nil {
			return nil, err
		}

		problems = append(problems, problem)
	}

	return problems, nil
}

func (p *ProblemRepo) GetProblemByID(id int) (models.Problem, error) {
	var problem models.Problem

	err := p.DB.QueryRow(`
		SELECT id, user_id, title, description FROM problems WHERE id = $1
	`, id).Scan(&problem.ID, &problem.UserID, &problem.Title, &problem.Description)

	return problem, err
}

func (p *ProblemRepo) UpdateProblem(problem models.Problem) error {
	_, err := p.DB.Exec(`
		UPDATE problems SET user_id = $1, title = $2, description = $3 WHERE id = $4
	`, problem.UserID, problem.Title, problem.Description, problem.ID)
	return err
}

func (p *ProblemRepo) DeleteProblem(id int) error {
	_, err := p.DB.Exec(`DELETE FROM problems WHERE id = $1`, id)
	return err
}
