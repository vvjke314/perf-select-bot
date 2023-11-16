package repository

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/vvjke314/MPPR/lab1/internal/model"
)

type Repository interface {
}

type PostgresRepo struct {
	db *pgx.Conn
}

func NewPostgresRepo(ctx context.Context, connString string) (PostgresRepo, error) {
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		return PostgresRepo{}, err
	}

	return PostgresRepo{
		db: conn,
	}, nil
}

func (pg PostgresRepo) Close() {
	pg.db.Close(context.Background())
}

// GetQuestion gest question_text, answer_1, answer_2, answer_3, answer_4 from the question
func (pg PostgresRepo) GetQuestion(questId int) (model.Question, error) {
	question := model.Question{}
	if err := pg.db.QueryRow(context.Background(), "SELECT question_text, answer_1, answer_2, answer_3, answer_4 FROM question WHERE question_id=$1", questId).Scan(
		&question.QuestionText,
		&question.Answer1,
		&question.Answer2,
		&question.Answer3,
		&question.Answer4); err != nil {
		return model.Question{}, fmt.Errorf("[GetQuestion] Error occur while finding question: %w", err)
	}
	return question, nil
}

func (pg PostgresRepo) CheckUserExistence(userId string) (bool, error) {
	user := model.User{}
	if err := pg.db.QueryRow(context.Background(), "SELECT * FROM user_data WHERE user_id=$1", userId).Scan(&user.UserId, &user.QuestionId); err != nil {
		if err.Error() == "no rows in result set" {
			return false, nil
		}
		return false, fmt.Errorf("[GetUserById] Error occur hile finding user: %w", err)
	}

	return true, nil
}

// AddUser use user ChatID to add it to DB
func (pg PostgresRepo) AddUser(userId string) error {
	user := model.User{
		UserId:     userId,
		QuestionId: 0,
	}
	if _, err := pg.db.Exec(context.Background(), "INSERT INTO user_data(user_id, question_id) VALUES($1, $2)", user.UserId, user.QuestionId); err != nil {
		return fmt.Errorf("[AddUser] Unable to insert due to: %w", err)
	}
	if _, err := pg.db.Exec(context.Background(), "INSERT INTO progress(progress_id, user_id) VALUES($1, $2)", strconv.Itoa(rand.Int()), user.UserId); err != nil {
		return fmt.Errorf("[AddUser] Unable to insert due to: %w", err)
	}
	return nil
}

// GetState
func (pg PostgresRepo) GetState(userId string) (int, error) {
	var currState int
	if err := pg.db.QueryRow(context.Background(), "SELECT question_id FROM user_data WHERE user_id=$1", userId).Scan(&currState); err != nil {
		return 0, err
	}
	return currState, nil
}

// UpdateState
func (pg PostgresRepo) UpdateState(userId string, newState int) error {
	if _, err := pg.db.Exec(context.Background(), "UPDATE user_data SET question_id = $1 WHERE user_id = $2", newState, userId); err != nil {
		return fmt.Errorf("[UpdateState] Unable to update due to: %w", err)
	}
	return nil
}

// NextState
func (pg PostgresRepo) NextState(userId string) (int, error) {
	currState, err := pg.GetState(userId)
	if err != nil {
		return 0, fmt.Errorf("[NextState] Unable to get next state due to: %w", err)
	}

	var nextState int
	if err := pg.db.QueryRow(context.Background(), "SELECT next_question FROM question WHERE question_id=$1", currState).Scan(&nextState); err != nil {
		return 0, fmt.Errorf("[NextState] Unable to get next state due to: %w", err)
	}
	return nextState, nil
}

// UpdateProgress
func (pg PostgresRepo) UpdateProgress(userId string, state int, value string) error {
	switch state {
	case 1:
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q1_answer = $1 WHERE user_id = $2", value, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 2:
		var input int
		switch value {
		case "Подарок":
			input = 2
		case "Для себя":
			input = 1
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q2_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 3:
		var input int
		switch value {
		case "Счастье":
			input = 1
		case "Умиротворение":
			input = 2
		case "Вдохновение":
			input = 3
		case "Роскошь":
			input = 4
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q3_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 4:
		var input int
		switch value {
		case "Для молодежи":
			input = 1
		case "Взрослый человек":
			input = 2
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q4_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 5:
		var input int
		switch value {
		case "Девушке":
			input = 1
		case "Парню":
			input = 2
		case "Другу":
			input = 3
		case "Родственнику":
			input = 4
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q5_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 6:
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q6_answer = $1 WHERE user_id = $2", value, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 7:
		var input int
		switch value {
		case "Мужчина":
			input = 1
		case "Женщина":
			input = 2
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q7_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 8:
		var input int
		switch value {
		case "Жирная":
			input = 1
		case "Сухая":
			input = 2
		case "Комбинированная":
			input = 3
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q8_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 9:
		var input int
		switch value {
		case "Для повседневного пользования":
			input = 1
		case "На работе":
			input = 2
		case "Во время отдыха":
			input = 3
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q9_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 10:
		var input int
		switch value {
		case "Лето":
			input = 1
		case "Зима":
			input = 2
		case "Весна/Осень":
			input = 3
		case "В любое":
			input = 4
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q10_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 11:
		var input int
		switch value {
		case "Сладкий":
			input = 1
		case "Душистый":
			input = 2
		case "Крепкий":
			input = 3
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q11_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 12:
		var input int
		switch value {
		case "Легкие и свежие":
			input = 1
		case "Насыщенные и тяжелые":
			input = 2
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q12_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 13:
		var input int
		switch value {
		case "Удержание аромата":
			input = 1
		case "Интенсивность запаха":
			input = 2
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q13_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 14:
		var input int
		switch value {
		case "Романтическую":
			input = 1
		case "Элегантную":
			input = 2
		case "Энергичную":
			input = 3
		case "Провокационную":
			input = 4
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q14_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 15:
		var input int
		switch value {
		case "Классические ароматы":
			input = 1
		case "Современные и экзлюзивное":
			input = 2
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q15_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 16:
		var input int
		switch value {
		case "Цветочные":
			input = 1
		case "Древесные":
			input = 2
		case "Фруктовые":
			input = 3
		case "Ориентальные":
			input = 4
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q16_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 17:
		var input int
		switch value {
		case "Привлекательный":
			input = 1
		case "Ненавязчивый":
			input = 2
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q17_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 18:
		var input int
		switch value {
		case "Они дополняют атмосферу каждого сезона":
			input = 1
		case "Они не имеют большого значения":
			input = 2
		case "Они ограничивают свободу выбора":
			input = 3
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q18_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 19:
		var input int
		switch value {
		case "Сильная стойкость":
			input = 1
		case "Чтобы держались недолго":
			input = 2
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q19_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 20:
		var input int
		switch value {
		case "100 мл":
			input = 1
		case "50 мл":
			input = 2
		case "10 мл":
			input = 3
		}
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q20_answer = $1 WHERE user_id = $2", input, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}
	case 21:
		if _, err := pg.db.Exec(context.Background(), "UPDATE progress SET q21_answer = $1 WHERE user_id = $2", value, userId); err != nil {
			return fmt.Errorf("[UpdateProgress; CASE %v] Unable to update due to: %w", state, err)
		}

	}
	return nil
}

// GetResult
func (pg PostgresRepo) GetResult(userId string) (model.Progress, error) {
	var result model.Progress
	if err := pg.db.QueryRow(context.Background(), "SELECT q1_answer, q2_answer, q3_answer, q4_answer, q5_answer FROM progress WHERE user_id=$1", userId).Scan(
		&result.Q1Answer,
		&result.Q2Answer,
		&result.Q3Answer,
		&result.Q4Answer,
		&result.Q5Answer); err != nil {
		if err := pg.db.QueryRow(context.Background(), "SELECT q1_answer, q2_answer, q6_answer, q7_answer, q8_answer, q9_answer, q10_answer, q11_answer, q12_answer, q13_answer, q14_answer, q15_answer, q16_answer, q17_answer, q18_answer, q19_answer, q20_answer FROM progress WHERE user_id=$1", userId).Scan(
			&result.Q1Answer,
			&result.Q2Answer,
			&result.Q6Answer,
			&result.Q7Answer,
			&result.Q8Answer,
			&result.Q9Answer,
			&result.Q10Answer,
			&result.Q11Answer,
			&result.Q12Answer,
			&result.Q13Answer,
			&result.Q14Answer,
			&result.Q15Answer,
			&result.Q16Answer,
			&result.Q17Answer,
			&result.Q18Answer,
			&result.Q19Answer,
			&result.Q20Answer); err != nil {
			return model.Progress{}, fmt.Errorf("[GetResult] Error occur while getting data: %w", err)
		}
	}

	return result, nil
}

// DeleteUser
func (pg PostgresRepo) DeleteUser(userId string) error {
	_, err := pg.db.Exec(context.Background(), "DELETE FROM progress WHERE user_id = $1", userId)
	if err != nil {
		return fmt.Errorf("[DeleteUser] can't delete from progress due to: %w", err)
	}

	_, err = pg.db.Exec(context.Background(), "DELETE FROM user_data WHERE user_id = $1", userId)
	if err != nil {
		return fmt.Errorf("[DeleteUser] can't delete from user_data due to: %w", err)
	}

	return nil
}
