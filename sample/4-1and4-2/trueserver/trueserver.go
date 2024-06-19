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
	
	fmt.Fprintln(writer, "***構造体のポインタを用いる関数***")
	member_add := &members[0] //members[0]のアドレス
	fmt.Fprintf(writer,
		functions.AddPointAndReport(&member_add, 12)) //members[0]のアドレスのアドレス
	
	//どこかのコピーではなくmembers[0]そのものが変更されたことを確認する
	fmt.Fprintln(writer, functions.Describe(members[0]))

	friend, s_string :=
		functions.CreateFriendAndReport(members[1],"エミコ")//members[1]のコピーを作る
	fmt.Fprintln(writer, s_string) //まず説明文を表示
	fmt.Fprintln(writer, functions.Describe(friend)) //戻ってきたコピーの値を確認

	fmt.Fprintln(writer, functions.Describe(members[1])) //ここのmembers[1]の値は変わっていないことを確認

	members = append(members, friend)//作成したインスタンスをスライスの要素に加える
	fmt.Fprintln(writer, "\n***メソッドの使用***")
	fmt.Fprintln(writer,
		functions.DescribeM_AllMembers(members))//この関数の使用法を確認する

	//メソッドの使用によって不思議簡略化された関数functions.AddPointMAndReportを使ってみる
	fmt.Fprintln(writer, "\n<<お友達紹介特典>>")
	fmt.Fprintln(writer,
		functions.AddPointMAndReport(&members[1], 20))//アドレスを渡せばよい！	
	
	//本当に変更されたのか？	
	fmt.Fprintln(writer, functions.Describe(members[1]))	
}

func with_pointers(writer http.ResponseWriter,
	req *http.Request){

	//模擬メモリ空間
	mockmemory := []int{
		325, 14, 160, 440, 16, 175, //カンマ必要
	}

	fmt.Fprintln(writer, "\n<<アドレス「0」指定>>")
	fmt.Fprintln(writer,
		functions.DescribeMockStruct(mockmemory, 0))

	fmt.Fprintln(writer, "\n<<アドレス「3」指定>>")
	fmt.Fprintln(writer,
	functions.DescribeMockStruct(mockmemory, 3))

	fmt.Fprintln(writer, "\n***ポインタを使う意味***")
	a, b := 10,10
	aa := functions.UpdateOrCopy(a, &b)//(1)
	fmt.Fprintf(writer, "a=%d, b=%d, aa=%d", a, b, aa)

}

func with_methods(writer http.ResponseWriter,
	req *http.Request){
	fmt.Fprintln(writer, "***連続処理***")
	
	marco := data.CreateTraveller("マルコ", 0, 0)//Recordに最初の文字列が与えられる
	marco = marco.Travel(2,3). //Travellerインスタンスが戻る
	Travel(12,24).
	Travel(45,78).
	Goal()//この戻り値がmatcoに代入される

	fmt.Fprintln(writer, marco.Record) //Recordに書き溜められた値を書き出してみよう
}

func main(){
	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/with_slices", with_slices)
	http.HandleFunc("/with_structs", with_structs)
	http.HandleFunc("/with_pointers", with_pointers)
	http.HandleFunc("/with_methods", with_methods)

	http.ListenAndServe(":8090", nil)
}