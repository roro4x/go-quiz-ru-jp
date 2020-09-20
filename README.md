# go-quiz-ru-jp

*Hello!*
*It's a quiz game for learning japanese language.*

### Exaples JSONs

**`/task`** - get new quiz task
```
{"lessons_id": ["1","2","4","5"]}
```
*Response example :*  
```
{
    "question_word": "y",
    "right_answer": "n",
    "wrod1": "n",
    "wrod2": "n",
    "wrod3": "n",
    "wrod4": "n"
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