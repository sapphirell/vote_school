package main

import (
    "encoding/json"
    "fmt"
    "github.com/tidwall/gjson"
    "io/ioutil"
    "log"
    "net/http"

    "./model"
)

func schoolList(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json; charset=UTF-8")             //返回数据格式是json

    school := model.SchoolModel{}
    data,_ := school.GetSchoolRank()
    jsonByte, _ := json.Marshal(data)

    n ,_ := fmt.Fprintf(w, string(jsonByte))
    fmt.Println(n)
}

func voteMySchool(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "voteMySchool")
}

func voteLog(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
    w.Header().Set("content-type", "application/json; charset=UTF-8")  //返回数据格式是json

    result,_ := ioutil.ReadAll(req.Body)
    requestJson := string(result)
    telPhone := gjson.Get(requestJson, "tel_phone").String()

    fmt.Println(telPhone)
    voteLog := model.VoteLogModel{}
    data,_ := voteLog.GetVoteLog(telPhone)

    fmt.Println(data)

    resp ,_:= json.Marshal(data)
    fmt.Fprintf(w, string(resp))
    //jsonByte, _ := json.Marshal(data)
    //
    //n ,_ := fmt.Fprintf(w, string(jsonByte))
}
func main() {
    fmt.Println("starting...")
    http.HandleFunc("/", schoolList)
    http.HandleFunc("/vote", voteMySchool)
    http.HandleFunc("/log", voteLog)
    err := http.ListenAndServe(":8080", nil)
    fmt.Println("success")
    if err != nil {
       log.Fatal("ListenAndServe: ", err.Error())
    }
}
