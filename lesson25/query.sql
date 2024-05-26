--Subquery, SQL so'rovlari ichida boshqa so'rovlar yozish uchun ishlatiladi. 
--Uning maqsadi odatda asosiy so'rovning boshqa qismini tayyorlash va unga ma'lumot ta'minlashdir. 
--Bu ma'lumot asosiy so'rovni bajarish uchun kerak bo'lgan ma'lumotlardan kelib chiqadi.


--SELECT column1, column2, ...
--FROM table1
--WHERE column1 = (SELECT column1 FROM table2 WHERE condition);


--SELECT column1, column2, ...
--FROM table1
--WHERE column1 IN (SELECT column1 FROM table2 WHERE condition);


SELECT 
    name,
    (SELECT COUNT(*) FROM users_cars WHERE user_id = users.id) AS car_count
FROM 
    users;

