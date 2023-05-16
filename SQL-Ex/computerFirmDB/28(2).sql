-- Используя таблицу Product, определить количество производителей, выпускающих
-- по одной модели.
-- Справка по теме:
--  Предложение GROUP BY
--  Предложение HAVING
--  Подзапросы

SELECT COUNT(qwe.maker)
FROM
  (SELECT maker,
          COUNT(model)
   FROM Product
   GROUP BY (maker)
   HAVING COUNT(model) = 1 ) AS qwe