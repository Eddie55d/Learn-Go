-- Найдите модели принтеров, имеющих самую высокую цену. Вывести: model, price
-- Справка по теме:
--  Подзапросы
--  Получение итоговых значений

SELECT model,
       price
FROM Printer
WHERE price =
    (SELECT MAX(price)
     FROM Printer)