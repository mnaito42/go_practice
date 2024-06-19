package main
import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, //ポインタ記号なし
	req *http.Request){ //ポインタ記号アリ
	fmt.Fprintln(writer, "レッツゴー\n") //ブラウザに出力する内容
}


func main(){
	http.HandleFunc("/hello", hello) // helloというURLで、helloという関数が呼ばれる
	http.ListenAndServe(":8090", nil) //ポート番号は8080よりちょっと大きいくらいで
}