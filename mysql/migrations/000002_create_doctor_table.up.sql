CREATE TABLE Doctor (
    doctor_id INT PRIMARY KEY AUTO_INCREMENT,
    crm VARCHAR(255),
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES User(user_id)
);

