-- Найдите среднюю скорость ПК, выпущенных производителем A.
-- Справка по теме:
--  Получение итоговых значений
--  Явные операции соединения

SELECT AVG(speed)
FROM PC
RIGHT JOIN Product ON Product.model = PC.model
WHERE maker='A'