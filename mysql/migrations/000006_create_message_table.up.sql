CREATE TABLE message (
    message_id INT PRIMARY KEY AUTO_INCREMENT,
    sender_id INT NOT NULL,
    recipient_id INT NOT NULL,
    datetime DATETIME NOT NULL,
    content TEXT NOT NULL,
    FOREIGN KEY (sender_id) REFERENCES user(user_id),
    FOREIGN KEY (recipient_id) REFERENCES user(user_id) 
  )
