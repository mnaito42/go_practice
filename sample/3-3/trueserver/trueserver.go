package main

import (
	"fmt"
	"net/http"
	//自分で作ったモジュールをインポート
	"trueserver/functions"
	"trueserver/data"
)

func add(writer http.ResponseWriter, req *http.Request){
	fmt.Fprintln(writer, "***はじめての関数***")
	
	//自分で作ったモジュールを使った処理結果を出力
	result := functions.Add(5,7)
	fmt.Fprintf(writer,"5+7=%d", result )
}

func sub(writer http.ResponseWriter, req *http.Request){
	a,b := 5,7 //引数に渡す値を変数に持たせておく
	output, result := functions.Sub(a,b) //変数から引数に値をリレー
	fmt.Fprintln(writer, "***複数の値を戻す関数***")
	fmt.Fprintf(writer, output, a, b, result )
}

func with_slices(writer http.ResponseWriter,
	req *http.Request){
	
	sl_1 := []int{1, 2, 3, 4}

	fmt.Fprintln(writer,
		"***スライスそのものを書き換える***")
	//処理前の値を確かめる
	fmt.Fprintln(writer,"\nsl_1は")
	fmt.Fprintln(writer, sl_1)//最初に定義されたスライス

	functions.AddAll(sl_1, 9) //sl_1を引数に渡す

	fmt.Fprintln(writer,"\nいまやsl_1は")
	fmt.Fprintln(writer, sl_1) //処理後のsl_1を確認

	fmt.Fprintln(
		writer, "\n***スライスのコピーを書き換える***")
	sl_2 := functions.AddAndCopy(sl_1, 100) 

	//新しくできたsl_2
	fmt.Fprintln(writer,"\nsl_2は")
	fmt.Fprintln(writer, sl_2)

	//もとのsl_1
	fmt.Fprintln(writer,"\n一方sl_1は")
	fmt.Fprintln(writer, sl_1)
}

func with_structs(writer http.ResponseWriter,
	req *http.Request){
	
	members := []data.Member{
		data.Member{"ゆみこ", 56, 1.24},
		data.Member{"トシオ",44, 0.98},
		data.Member{"かをる", 70, 1.02}, //カンマ必要
	}
	fmt.Fprintln(writer, "***構造体を用いた関数***")
	fmt.Fprintln(writer,
		functions.DescribeAllMembers(members))

	fmt.Fprintln(writer, "***構造体を戻す関数***")
	fmt.Fprintln(writer,
		functions.DescribeMaxPointMember(members))
}

func main(){
	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/with_slices", with_slices)
	http.HandleFunc("/with_structs", with_structs)

	http.ListenAndServe(":8090", nil)
}