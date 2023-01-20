CREATE TABLE IF NOT EXISTS choices(
    c_id  INT NOT NULL AUTO_INCREMENT,
    choice TEXT NOT NULL,
    is_correct INT NOT NULL,
    q_id INT NOT NULL,
    created_at DATETIME NULL,
    updated_at DATETIME    NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (c_id),
    CONSTRAINT questions_q_id_fk FOREIGN KEY (q_id) REFERENCES questions (q_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
