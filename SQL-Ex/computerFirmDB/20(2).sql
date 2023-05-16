-- Найдите производителей, выпускающих по меньшей мере три различных модели ПК. Вывести: Maker, число моделей ПК.
-- Справка по теме:
--  Предложение GROUP BY
--  Предложение HAVING

SELECT maker,
       COUNT(model)
FROM Product
WHERE type = 'PC'
GROUP BY maker
HAVING COUNT(model) > 2