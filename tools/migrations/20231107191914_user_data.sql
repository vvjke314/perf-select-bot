-- +goose Up
CREATE TABLE user_data (
    user_id text NOT NULL,
    question_id int,
    PRIMARY KEY(user_id),
    FOREIGN KEY(question_id)
    REFERENCES question(question_id)
);

-- +goose Down
DROP TABLE user_data;
