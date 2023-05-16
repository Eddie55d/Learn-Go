-- Найдите максимальную цену ПК, выпускаемых каждым производителем, у которого есть модели в таблице PC.
-- Вывести: maker, максимальная цена.
-- Справка по теме:
--  Предложение GROUP BY
--  Явные операции соединения

SELECT maker,
       MAX(price)
FROM Product
RIGHT JOIN PC ON Product.model = PC.model
GROUP BY maker