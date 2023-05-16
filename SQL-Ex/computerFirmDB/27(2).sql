-- Найдите средний размер диска ПК каждого из тех производителей, которые
-- выпускают и принтеры.
-- Вывести: maker, средний размер HD.
-- Справка по теме:
--  Получение итоговых значений
--  Явные операции соединения
--  Предикат IN

SELECT maker,
       AVG(hd)
FROM PC
LEFT JOIN Product ON Product.model = PC.model
WHERE maker IN
    (SELECT maker
     FROM Product
     WHERE type IN ('Printer'))
GROUP BY (maker)