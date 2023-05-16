-- Найдите производителей самых дешевых цветных принтеров. Вывести: maker, price
-- Справка по теме:
--  Явные операции соединения
--  Получение итоговых значений
--  Подзапросы

SELECT DISTINCT maker,
                price
FROM Printer
LEFT JOIN Product ON Printer.model = Product.model
WHERE color='y'
  AND price =
    (SELECT MIN(price)
     FROM Printer
     WHERE color='y')