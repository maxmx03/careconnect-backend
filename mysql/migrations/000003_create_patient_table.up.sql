CREATE TABLE patient (
    patient_id INT PRIMARY KEY AUTO_INCREMENT,
    doctor_id INT,
    cpf VARCHAR(255),
    date_of_birth DATE,
    username VARCHAR(255),
    password VARCHAR(255),
    FOREIGN KEY (doctor_id) REFERENCES doctor(doctor_id),
);

