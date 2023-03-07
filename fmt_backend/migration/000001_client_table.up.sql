CREATE TABLE
    IF NOT EXISTS clients (
        id INT NOT NULL AUTO_INCREMENT,
        first_name VARCHAR(200) NULL,
        last_name VARCHAR(200) NULL,
        email VARCHAR(100) NOT NULL,
        address VARCHAR(10) NOT NULL,
        password VARCHAR(250) NOT NULL,
        profile_photo VARCHAR(250) NULL,
        created_at DATETIME NOT NULL,
        updated_at DATETIME NULL,
        deleted_at DATETIME NULL,
        PRIMARY KEY (id),
        CONSTRAINT `UQ_admins_email` UNIQUE (email)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;