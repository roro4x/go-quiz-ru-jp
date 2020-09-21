# go-quiz-ru-jp

*Hello!*
*It's a quiz game for learning japanese language.*

### Examples of requests

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