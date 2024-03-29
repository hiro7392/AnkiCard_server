DROP TABLE IF EXISTS cards;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS users;



CREATE TABLE users (
user_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
user_name VARCHAR(100) NOT NULL,
user_email VARCHAR(100) NOT NULL,
user_password VARCHAR(255) NOT NULL,
user_level INT NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
PRIMARY KEY(user_id)
);

CREATE TABLE tags (
tag_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
created_user_id INT UNSIGNED NOT NULL,
tag_name VARCHAR(100) NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
PRIMARY KEY(tag_id),
FOREIGN KEY (created_user_id) REFERENCES users(user_id)
);

CREATE TABLE cards (
card_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
tag_id INT UNSIGNED DEFAULT 1,
tag_name VARCHAR(100) NULL,
created_user_id INT UNSIGNED NOT NULL,
learning_level INT NOT NULL,
question_text VARCHAR(200) NOT NULL,
answer_text VARCHAR(200) NOT NULL,
created_at DATETIME NOT NULL,
updated_at DATETIME NOT NULL,
PRIMARY KEY(card_id),
FOREIGN KEY (tag_id) REFERENCES tags(tag_id),
FOREIGN KEY (created_user_id) REFERENCES users(user_id)
);
