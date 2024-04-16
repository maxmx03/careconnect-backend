CREATE TABLE Patient (
    patient_id INT PRIMARY KEY AUTO_INCREMENT,
    cpf VARCHAR(255),
    date_of_birth DATE,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES User(user_id)
);

