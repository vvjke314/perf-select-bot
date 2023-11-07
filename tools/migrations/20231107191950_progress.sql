-- +goose Up
CREATE TABLE progress (
    progress_id text NOT NULL,
    user_id text,
    q0_answer int,
    q1_answer int,
    q2_answer int,
    q3_answer int,
    q4_answer int,
    q5_answer int,
    q6_answer int,
    q7_answer int,
    q8_answer int,
    q9_answer int,
    q10_answer int,
    q11_answer int,
    q12_answer int,
    q13_answer int,
    q14_answer int,
    q15_answer int,
    q16_answer int,
    q17_answer int,
    q18_answer int,
    q19_answer int,
    q20_answer int,
    q21_answer int,
    PRIMARY KEY(progress_id),
    FOREIGN KEY(user_id)
    REFERENCES user_data(user_id)
);

-- +goose Down
DROP TABLE progress;
