-- Найдите среднюю цену ПК и ПК-блокнотов, выпущенных производителем A (латинская буква).
-- Вывести: одна общая средняя цена.
-- Справка по теме:
--  Явные операции соединения
--  Подзапросы
--  Объединение
--  Получение итоговых значений

SELECT AVG(price) AS AVG_price
FROM
  (SELECT price
   FROM PC
   LEFT JOIN Product ON PC.model = Product.model
   WHERE maker = 'A'
   UNION ALL SELECT price
   FROM Laptop
   LEFT JOIN Product ON Laptop.model = Product.model
   WHERE maker = 'A' ) as avg