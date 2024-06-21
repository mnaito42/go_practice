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

func main(){
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/algebra", algebra)
	http.HandleFunc("/math", mathfuncs)
	http.HandleFunc("/arrays", arrays)
	http.HandleFunc("/slices", slices)
	http.HandleFunc("/struct_members", struct_members)
	http.ListenAndServe(":8090", nil)
}

func algebra(writer http.ResponseWriter,
	req *http.Request){
	a, b := 7,8
	result := "%d+%d=%d\n"
	fmt.Fprintln(writer, "***整数の足し算***")
	fmt.Fprintf(writer, result, a, b, a+b)
	c := 15.0
	d := float64(a)
	result ="%.1f/%.1f=%.3f\n"
	fmt.Fprintln(writer, "\n***少数の割り算***")
	fmt.Fprintf(writer, result, c, d, c/d)
}

func mathfuncs(writer http.ResponseWriter,
	req *http.Request){
	pow28 := int(math.Pow(2, 8))

	fmt.Fprintf(writer, "%dの%d乗は%d\n", 2, 8, pow28)

	rad30 := 30.0*math.Pi/180.0
	fmt.Fprintf(writer, "\nsin%dは %.3f\n", 30, math.Sin(rad30))
}

func arrays(writer http.ResponseWriter,
	req *http.Request){

	fmt.Fprintln(writer, "***要素が５個の配列を定義***")
	arr1:=[5]int{2, 4, 6, 8, 10}
	fmt.Fprintln(writer, arr1)

	fmt.Fprintln(writer, "\n***要素5個で定義した配列に３つの要素しか定義しないとどうなる？***")
	arr2:=[5]int{1,3,5}
	fmt.Fprintln(writer, arr2)

	fmt.Fprintln(writer, "\n***要素の値は変更可？***")
	arr2[4]=99
	fmt.Fprintln(writer, arr2)

	fmt.Fprintln(writer, "\n***配列の一部を参照するスライス***")
	sl1 := arr1[1:3]
	sl2 := arr2[3:]
	fmt.Fprintln(writer, sl1)
	fmt.Fprintln(writer, sl2)

	fmt.Fprintln(writer, "\n***スライスの値を変更するとどうなる？***")
	sl1[1]=36
	fmt.Fprintln(writer, sl1)
	fmt.Fprintln(writer, arr1)
}

func slices(writer http.ResponseWriter,
	req *http.Request){

	sl :=[]int{30, 45, 60, 90, 180}

	var rad_v float64

	for _, v := range sl{
		rad_v = float64(v)*math.Pi/180.0
		fmt.Fprintf(writer, "sin%dは %.3f\n\n", v, math.Sin(rad_v))
	}

	fmt.Fprintln(writer, "\n***スライスなら要素を増やせる！***")
	sl=append(sl, 225, 275, 360)
	fmt.Fprintln(writer, sl)

	fmt.Fprintln(writer, "\n***スライスの一部を参照するスライス***")
	sl_sl := sl[2:5]
	fmt.Fprintln(writer, sl_sl)
}

type member struct {
	name string
	point int
	coeff float64
}

type vip struct {
	member
	vip_point int
}

func struct_members(writer http.ResponseWriter,
	req *http.Request){

	fmt.Fprintln(writer, "***構造体memberのインスタンス***")
	yumiko := member{"ゆみこ", 56, 1.24}

	toshio := member{}
	toshio.name= "トシオ"
	toshio.point=44
	toshio.coeff=0.98

	members := []member{yumiko,toshio}

	effective := "%sさんの有効ポイントは%.2f\n"

	for _, v := range members{
		fmt.Fprintf(writer, effective, v.name, float64(v.point)*v.coeff,
		)
	}

	fmt.Fprintln(writer, "\n***構造体を埋め込んだ構造体***")
	vip_yumiko := vip{yumiko, 30}

	vip_point := vip_yumiko.member.point+vip_yumiko.vip_point
	fmt.Fprintf(writer, "%sさんはVIPなのでポイントは%d点",
		vip_yumiko.member.name, vip_point)

	vip_effective_point := float64(vip_point)*vip_yumiko.member.coeff
	fmt.Fprintf(writer, "有効ポイントは%.2f点\n", vip_effective_point)
}