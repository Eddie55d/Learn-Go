-- Найдите производителей принтеров, которые производят ПК с наименьшим объемом RAM и с
-- самым быстрым процессором среди всех ПК, имеющих наименьший объем RAM.
-- Вывести: Maker
-- Справка по теме:
--  Явные операции соединения
--  Подзапросы
--  Получение итоговых значений
--  Предикат IN

SELECT DISTINCT maker
FROM Product
RIGHT JOIN PC ON Product.model = PC.model
WHERE maker IN
    (SELECT maker
     FROM Product
     WHERE type='Printer')
  AND ram =
    (SELECT MIN(ram)
     FROM PC)
  AND speed =
    (SELECT MAX(speed)
     FROM PC
     WHERE ram =
         (SELECT MIN(ram)
          FROM PC))