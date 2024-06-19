package main
import (
	"fmt"
	"net/http"
	"math"
)

func hello(writer http.ResponseWriter,
	req *http.Request){
	fmt.Fprintln(writer, "レッツゴー\n") 
}

func algebra(writer http.ResponseWriter,  //引数は上記のhelloと全く同じ
	req *http.Request){

		a, b := 7,8
		result := "%d+%d=%d\n" //Fprintfの最初の引数に渡す
		fmt.Fprintln(writer, "***整数の足し算***") //これはただの見出し
		fmt.Fprintf(writer, result, a, b, a+b) //整数を文字列に埋め込んで出力

		c := 15.0
		d :=float64(a) //小数に変換
		result ="%.1f/%.1f=%.3f\n" //表示桁数を指定する小数の書式
		fmt.Fprintln(writer, "\n***小数の割り算***")
		fmt.Fprintf(writer, result, c, d, c/d)
	}

func mathfuncs(writer http.ResponseWriter,
	req *http.Request){

	pow28 :=int(math.Pow(2,8)) //mathパッケージのPow関数
	fmt.Fprintf(writer, "%dの%d乗は%d\n", 2, 8, pow28)

	rad30 := 30.0*math.Pi/180.0 //度をラジアンに変換
	fmt.Fprintf(writer, "\nsin%d°は %.3f\n",
		30, math.Sin(rad30))
}

func main(){
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/algebra", algebra)
	http.HandleFunc("/math", mathfuncs)

	http.ListenAndServe(":8090", nil) //必ず最後に置く
}