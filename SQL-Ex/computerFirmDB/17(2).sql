-- Найдите модели ПК-блокнотов, скорость которых меньше скорости каждого из ПК.
-- Вывести: type, model, speed
-- Справка по теме:
--  Использование в запросе нескольких источников записей
--  Использование ключевых слов ANY, ALL с предикатами сравнения

SELECT DISTINCT type,
                Laptop.model,
                Laptop.speed
FROM Product,
     Laptop
WHERE speed < ALL
    (SELECT speed
     FROM PC)
  AND type='Laptop'