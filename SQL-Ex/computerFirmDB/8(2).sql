-- Найдите производителя, выпускающего ПК, но не ПК-блокноты.
 -- Справка по теме:
--  Пересечение и разность
--  Предикат IN

SELECT DISTINCT maker
FROM Product
WHERE type IN ('PC')
EXCEPT
SELECT maker
FROM Product
WHERE type IN ('Laptop')