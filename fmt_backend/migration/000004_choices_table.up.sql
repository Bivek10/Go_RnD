CREATE TABLE IF NOT EXISTS choices(
    id  INT NOT NULL AUTO_INCREMENT,
    choice TEXT NOT NULL,
    is_correct INT NOT NULL,
    question_id INT NOT NULL,
    created_at DATETIME NULL,
    updated_at DATETIME    NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (id),
    CONSTRAINT questions_q_id_fk FOREIGN KEY (question_id) REFERENCES questions (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
