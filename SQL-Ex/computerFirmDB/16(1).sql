-- Найдите пары моделей PC, имеющих одинаковые скорость и RAM. В результате каждая
-- пара указывается только один раз, т.е. (i,j), но не (j,i),
-- Порядок вывода: модель с большим номером, модель с меньшим номером, скорость и RAM.
-- Справка по теме:
--  Использование в запросе нескольких источников записей

SELECT PC.model,
       PC2.model,
       PC.speed,
       PC.ram
FROM PC,

  (SELECT model,
          speed,
          ram
   FROM PC) AS PC2
WHERE PC.model > PC2.model
  AND (PC.speed = PC2.speed
       AND PC.ram = PC2.ram)
GROUP BY (PC.model,
          PC2.model,
          PC.speed,
          PC.ram)