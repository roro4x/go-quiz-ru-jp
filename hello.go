package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var dictionary map[int]map[string]string
var questionWord string
var answerWord string
var answerVar1 string
var answerVar2 string
var answerVar3 string
var answerVar4 string

var lessonString string = "Я - わたし/ Ты - あなた/Он, Она - あのひと/Учитель - せんせい/Студент - がくせい/Врач - いしゃ/Ученый - けんきゅうしゃ/Инженер - エンジニア"

type indexPageVariables struct {
	Rus     string
	Jap     string
	Qword   string
	Answer1 string
	Answer2 string
	Answer3 string
	Answer4 string
}

type resultPage struct {
	ResMessage string
	RAnswer    string
}

func lessonParser(s string) []string { return strings.Split(s, "/") }

func dictionaryParser(s string) []string { return strings.Split(s, "-") }

func createDictionary(s string) (m map[int]map[string]string) {
	m = make(map[int]map[string]string)
	for i := 0; i < len(lessonParser(lessonString)); i++ {
		mm := make(map[string]string)
		mm["Rus"] = strings.TrimSpace(dictionaryParser(lessonParser(lessonString)[i])[0])
		mm["Jap"] = strings.TrimSpace(dictionaryParser(lessonParser(lessonString)[i])[1])
		m[len(m)+1] = mm
	}
	return
}

func getPair(m map[int]map[string]string, i int) map[string]string {
	return m[i]
}

func getRandomNums(m map[int]map[string]string) (nums []int) {
	nums = []int{0, 0, 0, 0}
	rand.Seed(time.Now().UTC().UnixNano())
	i := 0
	for i < 4 {
		num := rand.Intn(len(m))
		if !contains(nums, num) {
			nums[i] = num
			i++
		}
	}
	return
}

func contains(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func generatePageVariables() {
	task := getRandomNums(dictionary)
	getPair(dictionary, task[0])
	var m map[string]string = getPair(dictionary, task[0])
	var question []string
	for i := 0; i < 4; i++ {
		question = append(question, getPair(dictionary, task[i])["Jap"])
	}
	rand.Shuffle(len(question), func(i, j int) {
		question[i], question[j] = question[j], question[i]
	})
	questionWord = m["Rus"]
	answerWord = m["Jap"]
	answerVar1 = question[0]
	answerVar2 = question[1]
	answerVar3 = question[2]
	answerVar4 = question[3]
}

func index(w http.ResponseWriter, r *http.Request) {
	generatePageVariables()
	QuizVariables := indexPageVariables{
		Jap:     answerWord,
		Rus:     questionWord,
		Qword:   questionWord,
		Answer1: answerVar1,
		Answer2: answerVar2,
		Answer3: answerVar3,
		Answer4: answerVar4,
	}
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, QuizVariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func answerPage(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("a")
	var PageVariables resultPage
	if s == answerWord {
		PageVariables = resultPage{
			ResMessage: "Answer is right",
			RAnswer:    answerWord,
		}
	}
	if !(s == answerWord) {
		PageVariables = resultPage{
			ResMessage: "Answer is wrong",
			RAnswer:    answerWord,
		}
	}
	t, err := template.ParseFiles("answerPage.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, PageVariables)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

func main() {
	dictionary = createDictionary(lessonString)
	http.HandleFunc("/", index)
	http.HandleFunc("/answer/", answerPage)
	http.ListenAndServe(":8080", nil)
}
