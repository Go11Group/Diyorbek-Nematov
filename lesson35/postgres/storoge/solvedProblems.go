package postgres

import (
	"database/sql"
	"solution/models"
)

type SolvedProblemRepo struct {
	DB *sql.DB
}

func NewSolvedProblemRepo(db *sql.DB) *SolvedProblemRepo {
	return &SolvedProblemRepo{DB: db}
}

func (s *SolvedProblemRepo) CreateSolvedProblem(solvedProblem models.SolvedProblem) error {
	_, err := s.DB.Exec(`
		INSERT INTO solved_problems (id, problem_id, user_id, solution) 
		VALUES ($1, $2, $3, $4)
	`, solvedProblem.ProblemID, solvedProblem.UserID, solvedProblem.Solution)
	return err
}

func (s *SolvedProblemRepo) GetAllSolvedProblems() ([]models.Solution, error) {
	var solutions []models.Solution

	rows, err := s.DB.Query(`
		SELECT 
			sp.id, 
			username,
			title as problemTitle, 
			description as problemDescription, 
			solution as problemSolution 
		FROM 
			users as u
		LEFT JOIN
			solved_problems as sp ON u.id = sp.user_id
		INNER JOIN
			problems as p ON sp.problem_id = p.id
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var solution models.Solution

		err := rows.Scan(&solution.ID, &solution.Username,  &solution.ProplemTitle, &solution.ProblemDescription, &solution.ProblemSolution)

		if err != nil {
			return nil, err
		}

		solutions = append(solutions, solution)
	}
	return solutions, err
}

func (s *SolvedProblemRepo) GetSolvedProblemById(userId, problemId int) (models.Solution, error) {
	var solution models.Solution
	err := s.DB.QueryRow(`
	SELECT 
		sp.id, 
		username,
		title as problemTitle, 
		description as problemDescription, 
		solution as problemSolution 
	FROM 
		users as u
	LEFT JOIN
		solved_problems as sp ON u.id = sp.user_id
	INNER JOIN
		problems as p ON sp.problem_id = p.id
	WHERE
		u.id=$1 AND p.id=$2
	`, userId, problemId).Scan(&solution.ID, &solution.Username,  &solution.ProplemTitle, &solution.ProblemDescription, &solution.ProblemSolution)

	return solution, err
}

func (s *SolvedProblemRepo) UpdateSolvedProblem(solvedProblem models.SolvedProblem) error {
	_, err := s.DB.Exec(`
		UPDATE solved_problems SET problem_id=$1, user_id=$2, solution=$3 WHERE id=$4
	`, solvedProblem.ProblemID, solvedProblem.UserID, solvedProblem.Solution, solvedProblem.ID)
	return err
}

func (s *SolvedProblemRepo) DeleteSolvedProblem(id int) error {
	_, err := s.DB.Exec(`DELETE FROM solved_problems WHERE id=$1`, id)
	return err
}