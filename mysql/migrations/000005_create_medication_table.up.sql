CREATE TABLE IF NOT EXISTS medication (
    medication_id INT PRIMARY KEY AUTO_INCREMENT,
    patient_id INT,
    name VARCHAR(255),
    dosage VARCHAR(255),
    schedule VARCHAR(255),
    FOREIGN KEY (patient_id) REFERENCES patient(patient_id)
);

