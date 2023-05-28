-- Для каждого производителя, у которого присутствуют модели хотя бы в одной из
-- таблиц PC, Laptop или Printer, определить максимальную цену на его продукцию.
-- Вывод: имя производителя, если среди цен на продукцию данного производителя
-- присутствует NULL, то выводить для этого производителя NULL,
-- иначе максимальную цену.

SELECT qwe.maker,
       MAX(qwe.price)
FROM
  (SELECT DISTINCT maker,
                   price,
                   MAX(PC.price)
   FROM Product
   LEFT JOIN PC ON Product.model = PC.model
   WHERE product.type = 'PC'
   GROUP BY (maker,
             price)
   UNION SELECT DISTINCT maker,
                         price,
                         MAX(Laptop.price)
   FROM Product
   LEFT JOIN Laptop ON Product.model = Laptop.model
   WHERE product.type = 'Laptop'
   GROUP BY (maker,
             price)
   UNION SELECT DISTINCT maker,
                         price,
                         MAX(Printer.price)
   FROM Product
   LEFT JOIN Printer ON Product.model = Printer.model
   WHERE product.type = 'Printer'
   GROUP BY (maker,
             price)) AS qwe
GROUP BY (maker)