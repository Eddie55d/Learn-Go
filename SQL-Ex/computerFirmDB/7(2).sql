-- Найдите номера моделей и цены всех имеющихся в продаже
-- продуктов (любого типа) производителя B (латинская буква).
 -- Справка по теме:
--  Объединение
--  Явные операции соединения

SELECT Product.model,
       price
FROM Product
RIGHT JOIN PC ON PC.model = Product.model
WHERE maker='B'
UNION
SELECT Product.model,
       price
FROM Product
RIGHT JOIN Laptop ON Laptop.model = Product.model
WHERE maker='B'
UNION
SELECT Product.model,
       price
FROM Product
RIGHT JOIN Printer ON Printer.model = Product.model
WHERE maker='B'