# go-quiz-ru-jp

*Hello!*
*It's a quiz game for learning japanese language.*

### API
#### Examples of requests

**`/task`** - get new quiz task
```
{"lessons_id": ["1","2","4","5"]}
```
*Response example :*  
```
{
    "question_word": "Ты",
    "right_answer": "あなた",
    "word1": "わたし",
    "word2": "せんせい",
    "word3": "いしゃ",
    "word4": "これ"
}
```
**`/word`** - add new word in DB
```
{"lesson_id": 2,
"ru_word": "Ты",
"jp_word": "あなた"}
```
*Response example :*
dictionary_id
```
12
```
**`/check`** - check the answer
```
{"ru_word": "Ты",
"jp_word": "あなた"}
```
*Response example :* 
```
"WRONG"
```
**`/lessons`** - get numbers of lessons from DB
*Response example :* 
```
[1,3,4]
```
### SQL
#### Create Table
```
CREATE TABLE dictionary (
    dictionary_id   integer NOT NULL,
	lesson_id        integer NOT NULL,
    ru_word       text NOT NULL,
    jp_word         text NOT NULL,
    CONSTRAINT dictionary_pkey PRIMARY KEY (dictionary_id)
);
```
#### Create SEQUENCE
```
CREATE SEQUENCE dictionary_id START 1;
```
#### List of words in BD
```
INSERT INTO dictionary VALUES
(nextval('dictionary_id'), 1, 'Я', 'わたし'),
(nextval('dictionary_id'), 1, 'Мы', 'わたしたち'),
(nextval('dictionary_id'), 1, 'Вы', 'あなた'),
(nextval('dictionary_id'), 1, 'Он, Она', 'あのひと'),
(nextval('dictionary_id'), 1, 'Дамы и господа', 'みなさん'),
(nextval('dictionary_id'), 1, 'Учитель', 'せんせい'),
(nextval('dictionary_id'), 1, 'Преподаватель', 'きょうし'),
(nextval('dictionary_id'), 1, 'Студент', 'がくせい'),
(nextval('dictionary_id'), 1, 'Служащий в компании', 'かいちゃいん'),
(nextval('dictionary_id'), 1, 'Врач', 'いしゃ'),
(nextval('dictionary_id'), 1, 'Ученый', 'けんきゅうしゃ'),
(nextval('dictionary_id'), 1, 'Инженер', 'インジニア'),
(nextval('dictionary_id'), 1, 'Университет', 'だいがく'),
(nextval('dictionary_id'), 1, 'Это', 'これ'),
(nextval('dictionary_id'), 2, 'То(около собеседника)', 'それ'),
(nextval('dictionary_id'), 2, 'То(одинаково удалено)', 'あれ'),
(nextval('dictionary_id'), 2, 'Книга', 'ほん'),
(nextval('dictionary_id'), 2, 'Словарь', 'じしょ'),
(nextval('dictionary_id'), 2, 'Журнал', 'ざっし'),
(nextval('dictionary_id'), 3, 'Здесь', 'ここ'),
(nextval('dictionary_id'), 3, 'Там, у вас', 'そこ'),
(nextval('dictionary_id'), 3, 'Там', 'あそこ'),
(nextval('dictionary_id'), 3, 'Аудитория', 'きょうしつ'),
(nextval('dictionary_id'), 3, 'Столовая', 'じむしょ'),
(nextval('dictionary_id'), 4, 'Просыпаться', 'おきます'),
(nextval('dictionary_id'), 4, 'Спать', 'ねます'),
(nextval('dictionary_id'), 4, 'Работать', 'はたらきます'),
(nextval('dictionary_id'), 4, 'Отдыхать', 'やすみます'),
(nextval('dictionary_id'), 4, 'Учиться', 'べんきょうします');
```