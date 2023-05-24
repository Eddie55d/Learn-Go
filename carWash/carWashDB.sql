CREATE TYPE genders AS ENUM('male', 'female', 'other');

CREATE TABLE IF NOT EXISTS employees
(
	employee_id smallserial,
	first_name varchar(15) NOT NULL,
	last_name varchar(25) NOT NULL,
	gender genders NOT NULL,
  performance_factor real NOT NULL,
	work_days smallint[] NOT NULL,

  CONSTRAINT pk_employees_employee_id PRIMARY KEY(employee_id)
);

CREATE TABLE IF NOT EXISTS machine_categories 
(
  machine_id smallserial,
  machine_type varchar(15) NOT NULL,
  wash_time time NOT NULL,

  CONSTRAINT pk_machine_categories_machine_id PRIMARY KEY(machine_id)
);

CREATE TABLE IF NOT EXISTS car_wash_schedule
(
  wash_id serial,
  wash_date timestamp NOT NULL,
  expected_wash_start_time time NOT NULL,
  machine_id smallint NOT NULL,
  employee_id smallint NOT NULL,
  actual_wash_start_time time(0),
  actual_wash_end_time time(0),

  CONSTRAINT pk_car_wash_schedule_wash_id PRIMARY KEY(wash_id),
  CONSTRAINT fk_car_wash_schedule_machine_id FOREIGN KEY(machine_id) REFERENCES machine_categories,
  CONSTRAINT fk_car_wash_schedule_employee_id FOREIGN KEY(employee_id) REFERENCES employees
);

-- создание мойщика
INSERT INTO employees (first_name, last_name, gender, performance_factor, work_days)
VALUES
('Name1', 'Surname1', 'male', 1.2, ARRAY[1, 2, 5, 6]),
('Name2', 'Surname2', 'male', 1.2, ARRAY[3, 4, 7]),
('Name3', 'Surname3', 'female', 1, ARRAY[1, 3, 5]),
('Name4', 'Surname4', 'male', 1.5, ARRAY[2, 3, 6, 7]),
('Name5', 'Surname5', 'male', 1.3, ARRAY[4, 5, 6]),
('Name5', 'Surname5', 'female', 1.4, ARRAY[1, 4, 5])

-- обновление данных мойщика
UPDATE employees SET first_name = 'Name6', last_name = 'Surname6'
WHERE employee_id = 6

-- получение мойщика
SELECT first_name, last_name, performance_factor FROM employees
ORDER BY performance_factor

-- удаление мойщика
DELETE FROM employees
WHERE employee_id = 6

-- создание категории машины
 INSERT INTO machine_categories (machine_type, wash_time)
 VALUES
 ('car', '00:30:00'),
 ('van', '00:45:00'),
 ('truck', '01:15:00'),
 ('moto', '00:00:15')

-- обновление данных категории машины
 UPDATE machine_categories SET wash_time = '01:00:00'
 WHERE machine_id = 3
 RETURNING wash_time

-- получение категории машины
 SELECT machine_type FROM machine_categories
 ORDER BY machine_id

-- удаление категории машины
DELETE FROM machine_categories
WHERE machine_id = 4

--создание ожидаемой записи мытья машины
INSERT INTO car_wash_schedule 
(wash_date, expected_wash_start_time, machine_id, employee_id, actual_wash_start_time, actual_wash_end_time)
VALUES
('2023-01-10 10:00:05', '12:00:00', 1, 1, NULL, NULL),
('2023-01-10 12:20:10', '13:00:00', 1, 4, '13:30:00', '13:50:00'),
('2023-01-10 13:00:20', '15:00:00', 2, 4, '15:00:00', '15:30:00'),
('2023-01-10 13:30:15', '17:00:00', 3, 1, '17:00:00', '17:50:00'),

('2023-01-11 09:30:07', '11:00:00', 2, 2, '11:00:00', '11:37:30'),
('2023-01-11 11:10:05', '14:30:00', 1, 3, '14:30:00', '15:00:00'),
('2023-01-11 14:23:34', '19:00:00', 1, 4, '19:00:00', '19:20:00'),

('2023-01-12 09:10:10', '10:30:00', 1, 5, '10:30:00', '10:53:05'),
('2023-01-12 10:00:05', '10:45:00', 1, 2, '10:45:00', '11:10:00'),
('2023-01-12 11:00:18', '13:10:00', 3, 2, '13:10:00', '14:00:00'),
('2023-01-12 11:20:00', '12:50:00', 1, 5, '12:50:00', '13:13:05'),
('2023-01-12 15:00:15', '17:15:00', 2, 5, '17:15:00', '17:49:37'),

('2023-01-13 13:30:08', '14:10:00', 2, 5, '14:10:00', '14:44:37'),
('2023-01-13 15:05:55', '19:00:00', 1, 3, '19:00:00', '19:30:00'),
('2023-01-13 16:37:23', '18:30:00', 1, 1, '18:30:00', '18:55:00'),

('2023-01-14 09:45:39', '12:00:00', 1, 1, NULL, NULL),
('2023-01-14 10:10:05', '12:00:00', 1, 4, NULL, NULL),
('2023-01-14 11:07:33', '13:40:00', 1, 5, NULL, NULL)

-- обновление записи мытья машины
UPDATE car_wash_schedule SET expected_wash_start_time = '13:30:00'
WHERE wash_id = 2
RETURNING expected_wash_start_time

-- получение всех записей мытья, по входных параметрам - временной 
-- диапазон (с такого числа, часа, минуты, секунды, по такой-то)
SELECT wash_id, wash_date, expected_wash_start_time 
FROM car_wash_schedule
WHERE (DATE(wash_date) >= '2023-01-12' AND  expected_wash_start_time >='15:00:00')
AND (DATE(wash_date) <= '2023-01-14' AND  expected_wash_start_time <='12:00:00')

-- проставление даты фактического начала мойки
UPDATE car_wash_schedule SET actual_wash_start_time = '12:00:00'
WHERE wash_id = 1

-- проставление даты фактического окончания мойки
UPDATE car_wash_schedule SET actual_wash_end_time = (
SELECT (wash_time / performance_factor) + actual_wash_start_time FROM car_wash_schedule
LEFT JOIN machine_categories USING(machine_id)
LEFT JOIN employees USING(employee_id)
WHERE wash_id = 1)
WHERE wash_id = 1

-- удаление записи мытья машины
DELETE FROM car_wash_schedule
WHERE wash_id = 18


-- получение самого подходящего мойщика, по входным параметрам - дата, 
-- ожидаемое время мойки, id категории машины

-- не учитывал категорию машины потому-что все моют любую категорию машины 
-- согласно одному коэфф.
SELECT employee_id FROM employees
WHERE performance_factor = (SELECT MAX(performance_factor) FROM  employees 
						                WHERE employee_id IN (SELECT free_emp.employee_id FROM (SELECT employee_id FROM employees
                            WHERE EXTRACT(ISODOW FROM TIMESTAMP '2023-01-13') = ANY(work_days)
                            EXCEPT
                            SELECT employee_id FROM public.car_wash_schedule
                            WHERE DATE(wash_date) = '2023-01-13' AND expected_wash_start_time = '20:30:00') AS free_emp))

-- для получения мойщиков в конкретный день
SELECT employee_id FROM employees
WHERE EXTRACT(ISODOW FROM TIMESTAMP '2001-02-14') = ANY(work_days)

-- для получения свободного мойщика на конкретное время
SELECT employee_id FROM employees
WHERE EXTRACT(ISODOW FROM TIMESTAMP '2023-01-14') = ANY(work_days)
EXCEPT
SELECT employee_id FROM public.car_wash_schedule
WHERE DATE(wash_date) = '2023-01-14' AND expected_wash_start_time = '12:00:00'

-- получение лучшего  мойщика по критерию наибольшего времени работы, 
-- с наименьшем количеством оконченных моек с опозданием

-- получение мойщика по максимальному времени моек
-- не брал опоздания
SELECT employee_id, (SELECT SUM(actual_wash_end_time - actual_wash_start_time)) AS total
FROM public.car_wash_schedule
GROUP BY employee_id 
ORDER BY total DESC
LIMIT 1