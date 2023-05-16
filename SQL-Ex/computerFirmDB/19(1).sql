-- Для каждого производителя, имеющего модели в таблице Laptop, найдите
-- средний размер экрана выпускаемых им ПК-блокнотов.
-- Вывести: maker, средний размер экрана.
-- Справка по теме:
--  Предложение GROUP BY
--  Явные операции соединения

SELECT maker,
       AVG(screen)
FROM Laptop
LEFT JOIN Product ON Laptop.model = Product.model
GROUP BY (maker)