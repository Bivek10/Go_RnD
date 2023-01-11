CREATE TABLE IF NOT EXISTS plants (
  id INT NOT NULL AUTO_INCREMENT,
  plant_id VARCHAR(45) NOT NULL,
  plant_name VARCHAR(200) NULL,
  scientific_name VARCHAR(200) NOT NULL,
  plant_type VARCHAR(200) NOT NULL,
  temperature   VARCHAR(10)  NOT NULL,
  description TEXT NOT NULL,
  image_url VARCHAR(250) NOT NULL,
  growing_season VARCHAR(250) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NULL,
  deleted_at DATETIME NULL,
  PRIMARY KEY (id)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

