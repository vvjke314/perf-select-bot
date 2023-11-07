-- +goose Up
CREATE TABLE question (
    question_id int NOT NULL,
    question_text text,
    answer_1 text,
    answer_2 text,
    answer_3 text,
    answer_4 text,
    next_question int,
    PRIMARY KEY(question_id)
);

-- +goose Down
DROP TABLE question;
