CREATE TABLE IF NOT EXISTS auth_users (
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    auth_information_id INT,
    organization CHAR(64),
    name CHAR(64),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (auth_information_id) REFERENCES auth_information(id)
);
