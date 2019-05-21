package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "./model"
)

func schoolList(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json")             //返回数据格式是json

    school := model.SchoolModel{}
    data,_ := school.GetSchoolRank()
    jsonByte, _ := json.Marshal(data)

    n ,_ := fmt.Fprintf(w, string(jsonByte))
    fmt.Println(n)
}

func voteMySchool(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "voteMySchool")
}

func rangeSchool(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "voteMySchool")
}
func main() {
    http.HandleFunc("/", schoolList)
    http.HandleFunc("/vote", voteMySchool)
    err := http.ListenAndServe("localhost:8888", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
