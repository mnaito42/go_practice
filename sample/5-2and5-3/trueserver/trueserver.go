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

	fmt.Fprintln(writer, "\n***インターフェイスの練習***")
	fractions := []data.Fraction{ //Half型とFull型を一緒のスライスに置ける
		data.Half(1.5), data.Full(2), data.Half(2.5),
		data.Full(3), data.Half(3.5) , //閉じカッコが次の行なのでカンマ必要
	}
	fmt.Fprintln(writer, functions.ShowFractions(fractions))

	fmt.Fprintln(writer, "\n***もっとそれらしいインターフェイス***")
	counters := []data.Counter{
		data.CharCounter{"Let's count!"},
		data.CharCounter{"一二三四五六七八九"},
		data.DigitCounter{2500},
		data.DigitCounter{1963061},
		data.CharCounter{"以上!"}, //閉じ括弧の前で改行するのでカンマ
	}
	fmt.Fprintln(writer, functions.CountAll(counters))

	fmt.Fprintln(writer, "\n***ポインタもインターフェイスを実装できる***")
	var reader data.MockReader //型をインターフェイスに指定

	reader = &data.StringReader{} //アドレスはStringReader指定でよかった(Go1.19)
	reader.Read("2023年1月1日")
	reader.Read("Goのインターフェイスを学習した")
	reader.Read("難しかった")
	fmt.Fprintln(writer, reader.Write())

	reader = &data.IntReader{}
	reader.Read("21")
	reader.Read("abc") //読み飛ばされる
	reader.Read("75")
	reader.Read("へ3") //一部読み飛ばされる
	fmt.Fprintln(writer, reader.Write())

	fmt.Fprintln(writer, "\n***実装しないメソッドを使う***")
	int_reader := reader.(*data.IntReader)//型情報を変換
	fmt.Fprintln(writer, functions.IntReader2Int(*int_reader))//中身を渡す

}

func flows(writer http.ResponseWriter,
	req *http.Request){
	fmt.Fprintln(writer, "***while文に相当するfor***")
	fmt.Fprintln(writer, functions.While10(6)) //numが10より小さい
	fmt.Fprintln(writer, functions.While10(13)) //numが10より大きい
	fmt.Fprintln(writer, functions.While10(10)) //numがちょうど10

	fmt.Fprintln(writer, "\n***forを用いた無限ループ***")
	fmt.Fprintln(writer, functions.Forever(3))
	fmt.Fprintln(writer, functions.Forever(10000))

	fmt.Fprintln(writer, "\n***Switch文***")
	for i:=0; i<7; i++{ //(1)
		fmt.Fprintln(writer, functions.Div3(i))
	}

	fmt.Fprintln(writer, "\n***多様な条件のSwitch文***")
	for i:=0; i<10; i++{
		fmt.Fprintln(writer, functions.DivBy3(i))
	}
}

func generics(writer http.ResponseWriter, req *http.Request){
	//引数に渡してみる2種類のスライス
	sl_int :=[]int{0,1,2,3,4}
	sl_str := []string{"花","鳥","風","月","猫", "蛙", "春"}
	fmt.Fprintln(writer, sl_int)
	fmt.Fprintln(writer, sl_str)

	fmt.Fprintln(writer, "\n***ジェネリック***")
	//インデックス3の要素を除去する
	sl_3_int, g_str := functions.RemoveByIndex(sl_int, 3)
	fmt.Fprintln(writer, g_str)
	fmt.Fprintln(writer, sl_3_int)

	//インデックス0の要素を除去する
	sl_0_str, g_str := functions.RemoveByIndex(sl_str, 0)
	fmt.Fprintln(writer, g_str)
	fmt.Fprintln(writer, sl_0_str)
}

func goroutine(writer http.ResponseWriter, req *http.Request){
	fmt.Fprintln(writer, "***Goroutineによるスレッド管理***")

	//(1)Goroutine開始
	go func(){
		for i :=0; i<3; i++{
			fmt.Fprintln(writer,functions.Record("Hello",i,100))
		}
		fmt.Fprintln(writer,"Hello完了です")
	}()

	//(2)そのあとに始まる処理
	for i :=0; i<3; i++{
		fmt.Fprintln(writer, functions.Record("World",i,100))
	}
	fmt.Fprintln(writer,"World完了です")
}

func a_channel(writer http.ResponseWriter, req *http.Request){

	ch := make(chan string)//make関数でチャンネル作成

	//２つのGoroutine中でチャンネルに送る
	go functions.InChannel("最初はグー", 0, 200, ch)
	go functions.InChannel("ジャンケンポン", 0, 100, ch)

	x1 :=<-ch //先にチャンネルに送られてきたもの
	x2 :=<-ch //次にチャンネルに送られてきたもの

	fmt.Fprintln(writer, x1)
	fmt.Fprintln(writer, x2)
}

func b_channel(writer http.ResponseWriter, req *http.Request){
	fmt.Fprintln(writer, "***バッファつきチャンネル***")

	ch4 := make(chan string, 4) //バッファ4個のチャンネル
	go func(){//Goroutine
		for i :=0; i<3; i++{
			functions.InChannel("Hello",i,200, ch4)
		} //ch4に値を3個送る
		ch4<-"Hello完了です"//4個目を送る
	}()

	for i :=0; i<3; i++{
		fmt.Fprintf(writer, "World_%d\n", i)
	} //チャンネルに贈らず直接即時書き出す
	fmt.Fprintln(writer,"World完了です")

	//チャンネルが待っていてくれれば書き出せる！
	for i := 0; i<4; i++ {
		x:=<-ch4 //チャンネルから値を取り出す
		fmt.Fprintln(writer, x)
	}
}

func main(){
	http.HandleFunc("/add", add)
	http.HandleFunc("/sub", sub)
	http.HandleFunc("/with_slices", with_slices)
	http.HandleFunc("/with_structs", with_structs)
	http.HandleFunc("/with_pointers", with_pointers)
	http.HandleFunc("/with_methods", with_methods)
	http.HandleFunc("/flows", flows)
	http.HandleFunc("/generics", generics)
	http.HandleFunc("/goroutine", goroutine)
	http.HandleFunc("/a_channel", a_channel)
	http.HandleFunc("/b_channel", b_channel)

	http.ListenAndServe(":8090", nil)
}