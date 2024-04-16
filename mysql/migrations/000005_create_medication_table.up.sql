CREATE TABLE Medication (
    medication_id INT PRIMARY KEY AUTO_INCREMENT,
    patient_id INT,
    name VARCHAR(255),
    dosage VARCHAR(255),
    schedule VARCHAR(255),
    FOREIGN KEY (patient_id) REFERENCES Patient(patient_id)
);

