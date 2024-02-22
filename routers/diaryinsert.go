package routers

import (
	"eq/db"
	"net/http"
)

func Diaryinsert(w http.ResponseWriter, req *http.Request) {
	//var r xbdb.ReInfo
	params := postparams(req)
	diary := db.NewDiary()
	diary.Add(params)
}
