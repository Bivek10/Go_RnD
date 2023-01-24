CREATE TABLE IF NOT EXISTS quizhistory(
    quiz_id INT NOT NULL,
    user_id VARCHAR(250) NOT NULL,
    score INT NOT NULL,
    created_at DATETIME NULL,
    updated_at DATETIME    NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (quiz_id, user_id),
    CONSTRAINT quizs_table_quiz_id_fk FOREIGN KEY (quiz_id) REFERENCES quizs (id),
    CONSTRAINT users_table_user_id_fk FOREIGN KEY (user_id) REFERENCES users (id)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;