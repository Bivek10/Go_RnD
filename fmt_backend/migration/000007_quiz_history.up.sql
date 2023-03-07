CREATE TABLE IF NOT EXISTS quizhistory(
    quiz_id INT NOT NULL,
    client_id INT NOT NULL,
    score INT NOT NULL,
    created_at DATETIME NULL,
    updated_at DATETIME    NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (quiz_id, client_id),
    CONSTRAINT quizs_table_quiz_id_fk FOREIGN KEY (quiz_id) REFERENCES quizs (id),
    CONSTRAINT client_table_client_id_fk FOREIGN KEY (client_id) REFERENCES clients (id)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
