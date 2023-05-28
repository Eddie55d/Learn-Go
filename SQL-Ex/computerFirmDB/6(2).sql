-- Для каждого производителя, выпускающего ПК-блокноты c объёмом жесткого
-- диска не менее 10 Гбайт, найти скорости таких ПК-блокнотов.
-- Вывод: производитель, скорость.

SELECT DISTINCT maker,
                Laptop.speed
FROM Product
LEFT JOIN Laptop ON Laptop.model = Product.model
WHERE Laptop.hd >= 10