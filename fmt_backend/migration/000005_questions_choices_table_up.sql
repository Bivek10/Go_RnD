CREATE TABLE IF NOT EXISTS questionchoices(
    q_id  INT NOT NULL,
    c_id INT NOT NULL,
    created_at DATETIME NULL,
    updated_at DATETIME    NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (q_id, c_id),
    CONSTRAINT questions_q_id_fk FOREIGN KEY (q_id) REFERENCES questions (q_id),
    CONSTRAINT choices_q_id_fk FOREIGN KEY (c_id) REFERENCES choices (c_id)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;