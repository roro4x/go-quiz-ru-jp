package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123123"
	dbname   = "postgres"
)

var dbc *sql.DB
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

func addNewWord(w http.ResponseWriter, r *http.Request) {
	var word word
	_ = json.NewDecoder(r.Body).Decode(&word)
	id := 0
	err := dbc.QueryRow(sqlInsertWord, word.LessonID, word.RuWord, word.JpWord).Scan(&id)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
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
	i := 1
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
	qWord = wordPair(m[1]).Ru
	trWord = wordPair(m[1]).Jp
	rand.Shuffle(len(m), func(i, j int) {
		m[i], m[j] = m[j], m[i]
	})
	t := task{
		QWord:  qWord,
		TrWord: trWord,
		Word1:  wordPair(m[1]).Jp,
		Word2:  wordPair(m[2]).Jp,
		Word3:  wordPair(m[3]).Jp,
		Word4:  wordPair(m[4]).Jp,
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Content-Type", "application/json")
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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func getLessons(w http.ResponseWriter, r *http.Request) {
	var aid []int
	rows, err := dbc.Query("SELECT DISTINCT(lesson_id) FROM dictionary;")
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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	json.NewEncoder(w).Encode(aid)
}

func dbConnect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

func main() {
	dbc = dbConnect()
	r := mux.NewRouter()
	r.HandleFunc("/api/word", addNewWord).Methods("POST")
	r.HandleFunc("/api/task", getTask).Methods("POST")
	r.HandleFunc("/api/check", checkTask).Methods("POST")
	r.HandleFunc("/api/lessons", getLessons).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}
