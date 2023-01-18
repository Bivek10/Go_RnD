CREATE TABLE IF NOT EXISTS users (
  user_id VARCHAR(45) NOT NULL,
  full_name VARCHAR(45) NULL,
  email VARCHAR(100) NOT NULL,
  phone_number VARCHAR(10) NOT NULL,
  address   VARCHAR(100)  NOT NULL,
  designation VARCHAR(100) NOT NULL,
  password  VARCHAR(250) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NULL,
  deleted_at DATETIME NULL,
  PRIMARY KEY (user_id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

