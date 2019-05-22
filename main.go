package main

import (
    "encoding/json"
    "fmt"
    "github.com/tidwall/gjson"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"

    "./model"
)

func schoolList(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")                //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type")    //header的类型
    w.Header().Set("content-type", "application/json; charset=UTF-8") //返回数据格式是json

    school := model.SchoolModel{}
    data, _ := school.GetSchoolRank()
    jsonByte, _ := json.Marshal(data)

    n, _ := fmt.Fprintf(w, string(jsonByte))
    fmt.Println(n)
}

func voteMySchool(w http.ResponseWriter, req *http.Request) {

    w.Header().Set("Access-Control-Allow-Origin", "*")                //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type")    //header的类型
    w.Header().Set("content-type", "application/json; charset=UTF-8") //返回数据格式是json

    type votePostData struct {
        telPhone string `json:"tel_phone"`
        sid      []int  `json:"sid"`
    }

    result, _ := ioutil.ReadAll(req.Body)
    requestJson := string(result)

    sidArr := gjson.Get(requestJson, "vote").Array()
    telPhone := gjson.Get(requestJson, "tel_phone").String()

    if len(sidArr) == 0 {
        fmt.Fprintf(w, "{\"error\":\"sid 为空\"}")
        return
    }
    if telPhone == "" {
        fmt.Fprintf(w, "{\"error\":\"tel_phone 为空\"}")
        return
    }
    //格式化数据
    sids := []int{}
    fmt.Println("sidArr", sidArr)
    for _, v := range sidArr {
        fmt.Println("sid:", v.Int())
        sids = append(sids, int(v.Int()))
    }
    //fmt.Println("sids:", sids)

    voteLog := model.VoteLogModel{}
    res := voteLog.VoteMySchool(sids, telPhone)
    if res > 0 {
        errStr := "{\"error\":\"已经点赞过了\",\"code\":1,\"sid\":"
        errStr += strconv.Itoa(res)
        errStr += "}"
        fmt.Fprintf(w, errStr)
        return
    }
    if res == -1 {
        fmt.Fprintf(w, "{\"error\":\"您只能投最多3票\",\"code\":2}")
        return
    }
    fmt.Fprintf(w, "{\"error\":\"\",\"code\":0}")

    //fmt.Println("返回值：", strconv.Itoa(res) )
}

func voteLog(w http.ResponseWriter, req *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")                //允许访问所有域
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type")    //header的类型
    w.Header().Set("content-type", "application/json; charset=UTF-8") //返回数据格式是json

    result, _ := ioutil.ReadAll(req.Body)
    requestJson := string(result)
    telPhone := gjson.Get(requestJson, "tel_phone").String()

    if telPhone == "" {
        fmt.Fprintf(w, "{\"error\":\"telPhone 为空\",\"code\":1}")
        return
    }

    fmt.Println(telPhone)
    voteLog := model.VoteLogModel{}
    data, _ := voteLog.GetVoteLog(telPhone)

    fmt.Println(data)

    resp, _ := json.Marshal(data)
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

    if err != nil {
        log.Fatal("ListenAndServe: ", err.Error())
    }
}
