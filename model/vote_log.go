package model

import (
    "database/sql"
    "fmt"
    "log"
    "time"

    //"fmt"
    //"strconv"
)

const vote_log_table = "vote_log"

type VoteLogModel struct {
    linker DbLinker
    Db     *sql.DB
}

type VoteLogRow struct {
    Id       int    `json:"id"`
    TelPhone string `json:"tel_phone"`
    Sid      int    `json:"sid"`
    Time     string `json:"time"` //time.Now().Format("2006-01-02 15:04:05")
}

func (m *VoteLogModel) constructVoteLogModel() {
    d := DbLinker{}
    d.Init()
    m.linker, m.Db = d, d.DB
}
func (m *VoteLogModel) Exec(query string, args ...interface{}) (sql.Result, error) {
    m.constructVoteLogModel()
    var res sql.Result
    var err error
    if args == nil {
        res, err = m.linker.DB.Exec(query)
    } else {
        res, err = m.linker.DB.Exec(query, args)
    }

    return res, err
}

//获取该手机号都投过谁
func (m *VoteLogModel) GetVoteLog(telPhone string) ([]VoteLogRow, error) {
    m.constructVoteLogModel()

    queryString := "SELECT id,tel_phone,sid,time FROM " + vote_log_table + " WHERE tel_phone = ?"
    rows, err := m.Db.Query(queryString, telPhone)
    data := []VoteLogRow{}
    i := 0
    for rows.Next() {
        row := VoteLogRow{}
        err := rows.Scan(&row.Id, &row.TelPhone, &row.Sid, &row.Time)
        if err != nil {
            log.Fatal(err)
        }
        //log.Println(row)
        data = append(data, row)
        //data[i] = row
        i++
    }
    return data, err

}

func (m *VoteLogModel) VoteMySchool(sidArr []int, telPhone string) (int) {
    //获取用户所有的点赞记录
    userVoteLog, _ := m.GetVoteLog(telPhone)
    fmt.Println("已经投票过：",len(userVoteLog))
    if len(userVoteLog) >= 3 {
        return -1
    }
    for _, value := range userVoteLog {
        for _, sid := range sidArr {
            if value.Sid == sid {
                return value.Sid; //已经投过票了,并且把投过票的sid返回
            }
        }
    }
    var insertString,updateString string
    insertString += "Insert into " + vote_log_table + "(tel_phone, sid, time) values (?,?,?)"
    updateString += "UPDATE " + school_table_name + " SET `ticket` = ticket + 1 where id = ?"
    stmt, err := m.Db.Prepare(insertString)
    dateTime := time.Now().Format("2006-01-02 15:04:05")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(sidArr)
    for _, sid := range sidArr {
        _, err := stmt.Exec(telPhone, sid, dateTime) //逐条插入点赞记录
        if err != nil {
            fmt.Println(err)
        }
        m.Db.Exec(updateString,sid) //修改点赞数量

    }

    //修改点赞数量
    return 0
}

//func (m *CprOrdersModel) SaveOrderResult(renderTime float64, resultVideoUrl string, notifyResult int) {
//   m.construct()
//   var queryString string
//
//   queryString = "UPDATE " + table + " SET `render_time` = '" + strconv.FormatFloat(renderTime, 'f', -1, 64) + "' "
//   queryString += " ,`result_video_url` = '" + resultVideoUrl + "' "
//   queryString += " ,`notify_result` = '" + strconv.FormatInt(int64(notifyResult), 10) + "' "
//
//   _, err := m.Db.Exec(queryString)
//   if err != nil {
//       fmt.Println(err)
//   }
//}
