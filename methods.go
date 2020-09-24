package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strings"
)

var sqlInsertWord = `INSERT INTO dictionary VALUES($1, $2, $3, nextval('dictionary_id')) RETURNING dictionary_id;`

type word struct {
	LessonID int    `json:"lesson_id"`
	RuWord   string `json:"ru_word"`
	JpWord   string `json:"jp_word"`
}

type wordPair struct {
	Ru string
	Jp string
}

type task struct {
	QWord  string `json:"question_word"`
	TrWord string `json:"right_answer"`
	Word1  string `json:"word1"`
	Word2  string `json:"word2"`
	Word3  string `json:"word3"`
	Word4  string `json:"word4"`
}

func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
}

func addNewWord(w http.ResponseWriter, r *http.Request) {
	var word word
	_ = json.NewDecoder(r.Body).Decode(&word)
	id := 0
	err := dbc.QueryRow(sqlInsertWord, word.LessonID, word.RuWord, word.JpWord).Scan(&id)
	if err != nil {
		panic(err)
	}
	setHeaders(w)
	json.NewEncoder(w).Encode(id)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	type LessonID struct {
		LessonsID []string `json:"lessons_id"`
	}
	var lID LessonID
	m := make(map[int]wordPair)
	var qWord string
	var trWord string
	_ = json.NewDecoder(r.Body).Decode(&lID)
	lIDstr := strings.Join(lID.LessonsID, ",")
	rows, err := dbc.Query("SELECT ru_word, jp_word FROM dictionary WHERE lesson_id in (" + lIDstr + ") ORDER BY random() LIMIT 4;")
	if err != nil {
		panic(err)
	}
	i := 0
	for rows.Next() {
		var ru string
		var jp string
		err = rows.Scan(&ru, &jp)
		if err != nil {
			panic(err)
		}
		m[i] = wordPair{
			Ru: ru,
			Jp: jp,
		}
		i++
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	qWord = wordPair(m[0]).Ru
	trWord = wordPair(m[0]).Jp
	rand.Shuffle(len(m), func(i, j int) {
		m[i], m[j] = m[j], m[i]
	})
	t := task{
		QWord:  qWord,
		TrWord: trWord,
		Word1:  wordPair(m[0]).Jp,
		Word2:  wordPair(m[1]).Jp,
		Word3:  wordPair(m[2]).Jp,
		Word4:  wordPair(m[3]).Jp,
	}
	setHeaders(w)
	json.NewEncoder(w).Encode(t)
}

func checkTask(w http.ResponseWriter, r *http.Request) {
	var word word
	var res string
	_ = json.NewDecoder(r.Body).Decode(&word)
	err := dbc.QueryRow("SELECT CASE WHEN EXISTS (SELECT * FROM dictionary WHERE ru_word = '" + word.RuWord + "' and jp_word = '" + word.JpWord + "') THEN 'RIGHT' ELSE 'WRONG' END").Scan(&res)
	if err != nil {
		panic(err)
	}
	setHeaders(w)
	json.NewEncoder(w).Encode(res)
}

func getLessons(w http.ResponseWriter, r *http.Request) {
	var aid []int
	rows, err := dbc.Query("SELECT DISTINCT(lesson_id) FROM dictionary ORDER BY lesson_id ASC;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		aid = append(aid, id)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	setHeaders(w)
	json.NewEncoder(w).Encode(aid)
}
