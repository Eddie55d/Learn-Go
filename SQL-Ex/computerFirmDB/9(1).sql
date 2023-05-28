-- Найдите производителей ПК с процессором не менее 450 Мгц. Вывести: Maker
-- Справка по теме:
--  Явные операции соединения
--  Простой оператор SELECT

SELECT DISTINCT maker
FROM Product
JOIN PC ON PC.model = Product.model
WHERE speed >= 450