package main

import (
	"net/http"
	"html/template" //新しい
	"math"
	"fmt"
)

type Page struct{ //形式は自由だが、とにかく構造体
	//フィールドは大文字
	Name string
	Answer string
}

func makepage(name string)Page{
	return Page{name, "さぁ始めましょう"}
}

//HTMLページを表示させる
func render(writer http.ResponseWriter, pg Page){ 
	t, _ := template.ParseFiles("form.html")
	t.Execute(writer, pg) //構造体を引数に渡す
}

//文字列を受け取って数字に変更する
func string2int(content string)int{

	memory := []int{}

	//数字であることを確かめて、スライスmemoryの要素にする
	digits:="0123456789"
	for _, v := range content{
		for i, s := range digits{
			if v == s{
				memory =append(memory, i)
			}
		}
   }

   //memoryの要素から整数を作成
	sum:=0
	lm :=len(memory)
	for i := 0; i<lm; i++ {
		mag := math.Pow10(lm-i)
		sum +=memory[i]*int(mag)
	}
	return sum/10		
}


//URLget_stringで呼ばれる関数
func get_string(writer http.ResponseWriter, req *http.Request){
	pg := makepage("get_string")
	if req.Method == "POST" {
		gs := req.FormValue("input")
		pg.Answer = gs //送信された値をそのまま返す
	}
	render(writer, pg)
}

//URLadd2で呼ばれる関数
func add2(writer http.ResponseWriter, req *http.Request){
	pg := makepage("add2")
	if req.Method == "POST" {
		gs := req.FormValue("input")
		a2 := string2int(gs)+2
		pg.Answer=fmt.Sprintf("%s+2=%d", gs, a2)
		
	}
	render(writer, pg)
}
func main(){
	
	http.HandleFunc("/get_string", get_string)
	http.HandleFunc("/add2", add2)
	http.ListenAndServe(":8090", nil)
}