CREATE TABLE IF NOT EXISTS questions (
    q_id  INT NOT NULL AUTO_INCREMENT,
    question TEXT NOT NULL,
    quiz_id INT NOT NULL, 
    created_at DATETIME NULL,
    updated_at DATETIME NULL,
    deleted_at DATETIME NULL,
    PRIMARY KEY (q_id),
    CONSTRAINT quizs_quiz_id_fk FOREIGN KEY (quiz_id) REFERENCES quizs(quiz_id)

) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;