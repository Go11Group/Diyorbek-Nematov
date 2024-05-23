-- Update -> buyrug'i orqali jadvaldagi ma'lumotlarni o'zgartirish mumkin
-- UPDATE tableName SET columnName = 'new'  WHERE columnName = 'old'

UPDATE employees SET position = 'Software Developer'  WHERE position = 'Developer';

-- DELETE -> buyrug'i orqali jadvaldagi malumotlarni o'chirish uchun qo'llaniladi
-- DELETE FROM tableName WHERE (shart)

DELETE FROM employees WHERE postion = 'Tester';

-- Ager DELETE FROM tableName query berilsa bunda jdvaldagi barcha ma'lumotlarni o'chirib yuboradi


-- ORDER BY -> buyrug'i berilgan columnlar asosidida jadvalni tartiblab beradi
-- ORDER BY DESC bilan esa jadvalni teskari tartibda tartiblab beradi
-- SELECT column1, column2 FROM tableName ORDER BY column1, column2, ... ASC bu yerda ASC O'sish tartibda tartiblab beradi

SELECT * FROM employees WHERE position = 'Project Manajer' ORDER BY name;

SELECT * FROM employees WHERE position = 'Developer' ORDER BY id DESC;

-- birinchi queryda position columni Project Mangerga teng bo'lgani name bo'yicha o'sish tartibda tartiblab chiqilmoqda
-- Ikkinchi queryda position columni Developer ga teng bo'lgani id bo'yicha kamayish tartibda tartiblab chiqariladi

-- GROUP BY -> operatori bir xil ma'lumotlarni guruhlash uchun ishlatiladi
-- GROUP BY operatori odatda agregat funksiyalar bilan qo'llaniladi masalan (COUNT(), SUM(), MAX(), MIN()) bilan ishlatiladi

-- SELECT column1, column2, FROM tableName GROUP BY column1, column2

SELECT position, COUNT(id) FROM employees
GROUP BY position

-- GROUP BY operatori HAVING Kalit so'zi bilan ham qo'llaniladi HAVING Gruuhlangan natijalarni filtirlash uchun ishlatiladi


-- JOINlar ikkita yoki undan ortiq jadvaldan bog‘langan ustunlar asosida ma'lumotlarni olish uchun ishlatiladi.
-- Join turlari quyidagilardan iborat:

-- INNER JOIN ikkita ikkita jadvalda ham mos keluvchi yozuvlar mavjud bo'lgan ma'lumotlarni qaytaradi
-- LEFT JOIN chap tarafdagi barcha ma'lumotlarni va o'ng jadvaldan mos keluvchi ya'ni chap jadvalda bor ma'lumotlarni chiqaradi
-- RIGHT JOIN o'ng jadvaldagi barcha ma'lumotlarni va chap jadvaldan o'ng jadvalda borlarini mos keluvchi ma'lumotlarni chiqaradi
-- FULL JOIN Ikkala jadvalda ham barcha ma'lumotlarni qaytaradi agar mos malumot topilmasa bo'shliq qaytadi
-- CROS JOIN ikki jadvalda hamma mumkin bo'lgan kombinatsiyalarni qaytaradi ya'ni bitta rowga ikkinchi jadvaldagi
-- barcha rowlarni mos qo'yib chiqadi va boshqa rowlar uchun ham shunday qiladi
-- SELF JOIN bir jadvalni o'zi bilan join qilishdir


--      INNER JOIN
-- SELECT table1.column1, table2.column2...
-- FROM table1
-- INNER JOIN table2
-- ON table1.common_column = table2.common_column;

-- LEFT JOIN
-- SELECT table1.column1, table2.column2...
-- FROM table1
-- LEFT JOIN table2
-- ON table1.common_column = table2.common_column;

-- RIGHT JOIN
-- SELECT table1.column1, table2.column2...
-- FROM table1
-- RIGHT JOIN table2
-- ON table1.common_column = table2.common_column;

-- FULL JOIN
-- SELECT table1.column1, table2.column2...
-- FROM table1
-- FULL JOIN table2
-- ON table1.common_column = table2.common_column;


-- CROSS JOIN
-- SELECT table1.column1, table2.column2..
-- FROM table1
-- CROSS JOIN table2;

-- SELF JOIN 
-- SELECT a.column_name, b.column_name...
-- FROM table_name AS a, table_name AS b
-- WHERE condition;

