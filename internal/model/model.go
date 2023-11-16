package model

type User struct {
	UserId     string `json:"user_id"`
	QuestionId int    `json:"question_id"`
}

type Question struct {
	QuestionId   int    `json:"question_id"`
	QuestionText string `json:"question_text"`
	Answer1      string `json:"answer_1"`
	Answer2      string `json:"answer_2"`
	Answer3      string `json:"answer_3"`
	Answer4      string `json:"answer_5"`
	NextQuestion int    `json:"next_question"`
}

type Progress struct {
	ProgressId string `json:"progress_id"`
	UserId     string `json:"user_id"`
	Q0Answer   int    `json:"q0_answer"`
	Q1Answer   int    `json:"q1_answer"`
	Q2Answer   int    `json:"q2_answer"`
	Q3Answer   int    `json:"q3_answer"`
	Q4Answer   int    `json:"q4_answer"`
	Q5Answer   int    `json:"q5_answer"`
	Q6Answer   int    `json:"q6_answer"`
	Q7Answer   int    `json:"q7_answer"`
	Q8Answer   int    `json:"q8_answer"`
	Q9Answer   int    `json:"q9_answer"`
	Q10Answer  int    `json:"q10_answer"`
	Q11Answer  int    `json:"q11_answer"`
	Q12Answer  int    `json:"q12_answer"`
	Q13Answer  int    `json:"q13_answer"`
	Q14Answer  int    `json:"q14_answer"`
	Q15Answer  int    `json:"q15_answer"`
	Q16Answer  int    `json:"q16_answer"`
	Q17Answer  int    `json:"q17_answer"`
	Q18Answer  int    `json:"q18_answer"`
	Q19Answer  int    `json:"q19_answer"`
	Q20Answer  int    `json:"q20_answer"`
	Q21Answer  int    `json:"q21_answer"`
}

type Results struct {
	Name  string
	Value int
}
