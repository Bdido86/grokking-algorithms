[Оглавление](../../README.md) > [Глава 4. Быстрая сортировка](../README.md) > Алгоритм Евклида

# Алгоритм Евклида

Предназначен для поиска наибольшего общего делителя двух чисел (НОД).

> НОД - наибольшее число, на которое оба числа делятся без остатка.

## Алгоритм с делением

1. Большее число делим на меньшее.
2. Если делится без остатка, то меньшее число и есть НОД (следует выйти из цикла).
3. Если есть остаток, то большее число заменяем на остаток от деления.
4. Переходим к пункту 1.

## Алгоритм с вычитанием

1. Из большего числа вычитаем меньшее.
2. Если получается 0, то значит, что числа равны друг другу и являются НОД (следует выйти из цикла).
3. Если результат вычитания не равен 0, то большее число заменяем на результат вычитания.
4. Переходим к пункту 1.