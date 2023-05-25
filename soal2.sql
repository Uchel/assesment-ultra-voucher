CREATE TABLE parent(
id INT PRIMARY KEY,
parent_name VARCHAR(100)
);
INSERT INTO parent(id, parent_name) VALUES
(2, 'Ilham'),
(3, 'Irwan');

CREATE TABLE person (
	id INT,
	name VARCHAR(100),
	parent_id INT
);

INSERT INTO person(id,name,parent_id) VALUES
(1,'Zaki',2),
(2,'Ilham',NULL),
(3,'Irwan',2),
(4,'Arka',3);

SELECT pe.id, pe.name, pa.parent_name FROM person AS pe 
LEFT JOIN parent AS pa ON pe.parent_id =pa.id;

