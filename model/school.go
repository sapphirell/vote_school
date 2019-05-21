package model

import (
    "database/sql"
    "log"
)

const school_table_name = "school"

type SchoolModel struct {
    linker     DbLinker
    Db         *sql.DB
}

type SchoolRow struct {
    Id         int `json:"id"`
    SchoolName string `json:"school_name"`
    MottoCh    string `json:"motto_ch"`
    MottoEn    string `json:"motto_en"`
    Ticket     int `json:"ticket"`
    Name       string `json:"name"`
    Img        string `json:"img"`
}

func (m *SchoolModel) constructSchoolModel() {
    d := DbLinker{}
    d.Init()
    m.linker, m.Db = d, d.DB
}

func (m *SchoolModel) GetSchoolRank()  ([8]SchoolRow,error) {
    m.constructSchoolModel()

    data := [8]SchoolRow{};

    queryString := "SELECT id,school_name,motto_ch,motto_en,ticket,name,img FROM " + school_table_name + " Order By ticket"
    rows, err := m.Db.Query(queryString)

    i := 0
    for rows.Next() {
        row := SchoolRow{}
        err := rows.Scan(&row.Id, &row.SchoolName, &row.MottoCh, &row.MottoEn, &row.Ticket, &row.Name, &row.Img)
        if err != nil {
            log.Fatal(err)
        }
        //log.Println(row)
        data[i] = row
        i++
    }
    return data,err
}

