**3. Стеки**

**Что такое:**  
Стек — это абстрактный тип данных, представляющий собой коллекцию элементов, с двумя основными принципами действий: push (добавить элемент на вершину) и pop (удалить верхний элемент).

**Математическая модель:**  
Можно представить стек как список или последовательность \(s_1, s_2, ..., s_n\), где \(s_n\) — это верхний элемент стека.

**Области применения:**
- Обратная Польская нотация и вычисление арифметических выражений.
- Поддержка операций undo в приложениях.
- Обработка рекурсии в программировании.

**Когда лучше применять:**
- Когда требуется доступ к последнему добавленному элементу (LIFO — Last In, First Out).
- Когда нужно временно хранить элементы и извлекать их в обратном порядке.

**Когда выбрать другую структуру:**  
Если требуется доступ к элементам в середине коллекции или добавление/удаление из середины, стек не подойдет. В этом случае могут подойти списки или другие структуры данных.

**Пример:** При анализе выражений типа "2 + 3 * 4", можно использовать стек для временного хранения операндов и операторов.

**Пример кода на Go:**

**Ресурсы для дополнительного изучения:**
- "Введение в алгоритмы" Кормена — содержит информацию о стеках и их применении.
- [Golang Docs](https://golang.org/doc/) — официальная документация Go.
- Сайты для практики задач по программированию, такие как LeetCode или HackerRank, содержат задачи, связанные со стеками, которые можно решать на Go.