-- Перечислите номера моделей любых типов, имеющих самую высокую цену по всей имеющейся
-- в базе данных продукции.
-- Справка по теме:
--  Объединение
--  Получение итоговых значений
--  Подзапросы
--  CTE
 WITH maxPrice AS
  (SELECT model,
          price
   FROM PC
   UNION SELECT model,
                price
   FROM Laptop
   UNION SELECT model,
                price
   FROM Printer)
SELECT model
FROM maxPrice
WHERE price =
    (SELECT MAX(Price)
     FROM maxPrice)