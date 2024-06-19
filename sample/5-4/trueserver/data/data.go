package data

import (
	"fmt"
	"math"
)

//構造体Memberとその関数・メソッド
//ほかの関数で使うので、名前は大文字で始める
type Member struct{
	//フィールドも大文字
	Name string
	Point int
	Coeff float64
}

func Effective(m Member)float64{
	return float64(m.Point)*m.Coeff
}

//Effectiveのメソッド版
func (member Member)EffectiveM() float64{
	return float64(member.Point)*member.Coeff
}

func Describe(m Member)string{
	return fmt.Sprintf("%sさんのポイントは%d点、有効ポイントは%.2f点",
	m.Name, m.Point, Effective(m))
}

//Describeのメソッド版
func(member Member)DescribeM() string{
	return fmt.Sprintf("%sさんのポイントは%d点、有効ポイントは%.2f点",
	member.Name, member.Point, member.EffectiveM())
}


func MaxPointMember(members []Member)Member{ //構造体のスライス
	mpm := members[0]
	for _, v := range members{
		if Effective(v) > Effective(mpm){ //面倒なので、等しい場合は考えないでおく
			mpm = v
		}
	}
	return mpm
}

func AddPoint(member **Member, p int){

	(**member).Point +=p //memberに星がついていることを明確にする

}

//AddPointのメソッド版。星が一個ですむ
func (member *Member)AddPointM(p int){ //(1)
	member.Point += p //中では星がなくても許される
}

func CreateFriendMember(member Member, name string )Member{ //引数は構造体の「コピー」
	member.Name = name//コピーのフィールドを変更
	return member
}

//構造体Travellerとそのメソッド
type Traveller struct{
	Name string
	X int
	Y int
	Record string
	}

//旅人を作成してスタート地点につける
func CreateTraveller(name string, x int, y int)Traveller{
	t:=Traveller{}
	t.Name=name
	t.X=x
	t.Y =y
	t.Record = fmt.Sprintf("%sさん(%d,%d)地点よりスタート\n",
	t.Name, t.X, t.Y )//旅人の旅程を文字列にしてつなげていく
	return t
}

//旅人が旅をするメソッド
func(t Traveller)Travel(x int, y int)Traveller{
	t.X=x
	t.Y=y
	t.Record += fmt.Sprintf("(%d, %d)へ移動\n", x, y )
	return t //自分自身（実はコピー）を返す
}

//旅人が到着するメソッド
func(t Traveller)Goal()Traveller{
	t.Record +="到着です\n"
	return t
}

//なぞのデータ型二つ
type Half float64
type Full int

//インターフェイスの定義
type Fraction interface{
	Value() string //メソッドの引数と戻り値を工夫したいところ
}

//HalfのレシーバでFractionのメソッドを実装
func(h Half)Value()string{
	return fmt.Sprintf("%.1f", float64(h))//強制型変換
}

//FullのレシーバでFractionのメソッドを実装
func(f Full)Value()string{
	return fmt.Sprintf("%d", int(f))
}

type Counter interface{
	DoCount()string
}

//Counterのメソッドを実装するべき二つの構造体
type CharCounter struct{
	Content string
}

type DigitCounter struct{
	Content int
}

//CharCounterのレシーバでCounterのメソッドを実装
func(counter CharCounter)DoCount()string{

	content := counter.Content //処理する値を取り出す
	s_string := fmt.Sprintf("「%s」は", content) //まず、内容を表示
	
	//文字列をUnicode記号の配列に変換する
	s_string += fmt.Sprintf("%d文字です", len([]rune(content)))

	return s_string
}

//DigitCounterのレシーバでCounterのメソッドを実装
func(counter DigitCounter)DoCount()string{
	content := counter.Content
	content_str := fmt.Sprintf("%d", content) //整数を文字列に変換

	s_string := fmt.Sprintf("「%d」は", content) //なんの桁数を数えたのか明示
	s_string += fmt.Sprintf("%d桁です", len([]rune(content_str)))

	return s_string
}

type MockReader interface{
	Read(content string) //引数文字列、戻り値なし
	Write() string //引数なし、戻り値文字列
}

//MockReaderのメソッドを実装するべき二つの構造体
type StringReader struct{
	Memory string //(1)
}

type IntReader struct{
	Memory []int //(2)
}

//MockReaderのメソッドReadの実装
//StringReaderのポインタレシーバによる
func(reader *StringReader)Read(content string){
	reader.Memory += content
	reader.Memory +="\n"
}

//IntReaderのポインタレシーバによる
func(reader *IntReader)Read(content string){

	digits:="0123456789"//数字のリストを文字列で実現

	for _, v := range content{ //引数contentについて
		for i, s := range digits{ //数字のリストについて
			if v == s{ //contentの文字が0から9のどれかに一致したら
				reader.Memory = append(reader.Memory, i)
			}
		}
	}
}

//MockReaderのメソッドWriteの実装
//StringReaderの値レシーバによる
func (reader StringReader)Write()string{
	s_string := "StringReaderインスタンスの中身は\n"
	s_string += "「"
	s_string += reader.Memory //StringReaderのフィールドMemoryは文字列
	s_string += "」"
	return s_string
}

//IntReaderの値レシーバによる
func (reader IntReader)Write()string{
	s_string := "IntReaderインスタンスの中身は\n"
	s_string += "["

	for _, v := range reader.Memory{ //IntReaderのフィールドMemoryは整数スライス
		s_string += fmt.Sprintf("%d ", v )
	}

	s_string += "]"
	return s_string
}

func (reader IntReader)Reader2Int()int{
	sum:=0
	memory := reader.Memory //何度も使うので変数に渡した

	lm :=len(memory) //スライスの要素数

	for i := 0; i<lm; i++ {
		mag := math.Pow10(lm-i) //本当はlm-i-1なのだがコードがキタナイ
		sum +=memory[i]*int(mag)
	}
	return sum/10 //最後に10で割る
}