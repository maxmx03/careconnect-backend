CREATE TABLE doctor (
    doctor_id INT PRIMARY KEY AUTO_INCREMENT,
    patient_id INT,
    crm VARCHAR(255),
    username VARCHAR(255),
    password VARCHAR(255)
    FOREIGN KEY (patient_id) REFERENCES doctor(patient_id),
);
