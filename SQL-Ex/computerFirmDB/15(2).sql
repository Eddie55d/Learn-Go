-- Найдите размеры жестких дисков, совпадающих у двух и более PC. Вывести: HD
-- Справка по теме:
--  Предложение GROUP BY
--  Предложение HAVING

SELECT hd
FROM PC
GROUP BY hd
HAVING COUNT(model)>=2