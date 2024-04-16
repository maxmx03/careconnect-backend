CREATE TABLE doctor (
    doctor_id INT PRIMARY KEY AUTO_INCREMENT,
    crm VARCHAR(255),
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);

