CREATE TABLE IF NOT EXISTS questions (
    q_id  INT NOT NULL AUTO_INCREMENT,
    question TEXT NOT NULL,
    created_at DATETIME NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL,
    PRIMARY KEY (q_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;