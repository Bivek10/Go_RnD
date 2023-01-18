CREATE TABLE IF NOT EXISTS quizs(
    quiz_id  INT NOT NULL AUTO_INCREMENT,
    quiz TEXT NOT NULL,
    created_at DATETIME NULL,
    updated_at DATETIME    NULL,
    deleted_at DATETIME    NULL,
    PRIMARY KEY (quiz_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
