package functions //main以外のパッケージ名

import (
	"fmt"
	"trueserver/data"
)

func Add(a int, b int) int{ //自分で定義する関数
	return a+b
}

func Sub(a,b int) (string, int){ //戻り値が二つ、異なる型
	return "%d-%dは%dなのだ", a-b //二つの戻り値はカンマで列記
}

func AddAll(sl []int, a int){ //スライスの型指定
	for i:=0; i<len(sl); i++{ //要素数を取得
		sl[i] += a
	}
}

func AddAndCopy(sl []int, a int)[]int{ //スライスを戻す

	sl_cp := []int{} //戻すためのスライス
	for i:=0; i<len(sl); i++{ //コピー元スライスの要素について繰り返す
		sl_cp = append(sl_cp, sl[i]+a) //コピー先スライスの要素を増やす
	}
	return sl_cp
}

func Describe(member data.Member)string{
	s_string :=
	fmt.Sprintf(data.Describe(member))
	s_string +="\n" //改行を追加しておくとラク
	return s_string
}

func DescribeAllMembers(members []data.Member)string{

	s_string := ""

	for _, v := range members {
		s_string +=Describe(v) //同じfunctions.goの関数「Describe」
	}

	return s_string
}

//関数の中でメソッドを呼び出す
func DescribeM_AllMembers(members []data.Member)string{
	s_string := "メソッドを使って書き出しても\n"
	for _, v := range members {
	s_string += v.DescribeM()//メソッド呼び出し
	s_string += "\n"
	}
	return s_string
	}

func DescribeMaxPointMember(members []data.Member)string{
	s_string := "有効ポイント最大の方は\n"

	mpm := data.MaxPointMember(members) //引数として受け取ったmember

	s_string += fmt.Sprintf("%sさん", mpm.Name)
	s_string += "\n"

	return s_string
}

func DescribeMockStruct(
	mockmemory []int,mockaddress int )string{

	s_string := fmt.Sprintf("名前は%dさん、",
		mockmemory[mockaddress])

	s_string += fmt.Sprintf("年齢%d歳、",
		mockmemory[mockaddress+1])

	s_string += fmt.Sprintf("身長は%dcm",
		mockmemory[mockaddress+2])
		
	return s_string
}

func UpdateOrCopy(a int, b *int) int{
	a += 3
	*b += 3
	return a
}

func AddPointAndReport(member **data.Member, p int)string{ //星2個で参照すると構造体になるヨ

	data.AddPoint(member, p) //アドレスのアドレスをそのまま、願いをこめて渡す

	s_string:= "<<得点アップサービス>>\n"

	s_string +=fmt.Sprintf("%sさんのポイント%d点アップ\n",
	(**member).Name, p) //アドレスのアドレスからたどりいた構造体のフィールドName

	s_string += "\n"

	return s_string //完成した説明文を戻す
}

//data.goのメソッドAddPointMを呼び出す
func AddPointMAndReport(member *data.Member, p int)string{ //星1個でよい
	
	member.AddPointM( p ) //ポインタレシーバなのでそのまま

	s_string:= "<<メソッドによる得点アップサービス>>\n"

	s_string +=fmt.Sprintf("%sさんのポイント%d点アップ",
		member.Name, p) //参照なのに星がいらない！

	return s_string
}



func CreateFriendAndReport(member data.Member,
	friend_name string)(data.Member, string){ //変更したインスタンスと説明文を戻す

	friend := data.CreateFriendMember(member, friend_name) //引数を受け取ったまま渡す

	s_string :=fmt.Sprintf(
		"%sさんの紹介で、お友達%sさんが加わりました",
		member.Name, friend_name)//この関数で受け取ったままの値と、さらにほかの関数に渡してもどってきた値

	s_string +="\n"

	return friend, s_string //コピーして変更したインスタンス
}

//インターフェイスを引数にとる関数
func ShowFractions(fractions[]data.Fraction)string{ //引数のデータ型に注目
	s_string := "スケール的には\n"
	for _, v := range fractions{
		s_string += fmt.Sprintf("%s倍か", v.Value()) //FractionのメソッドValue
		s_string += "\n"
	}
	s_string += "というところでしょうか\n"

	return s_string
}

func CountAll(counters []data.Counter)string{
	s_string :="<<data.Counterインターフェイス>>\n"

	for _, v := range counters{
		s_string += v.DoCount()
		s_string += "\n"
	}

	return s_string
}

//Reader2Intに説明をつけるだけ
func IntReader2Int(reader data.IntReader)string{
	s_string :="IntReaderから構成される整数は\n"
	s_string += fmt.Sprintf(
		"%d", reader.Reader2Int(), //次の行が閉じ括弧なのでカンマをつけている
	)
	s_string += "\n"
	return s_string
}

//forでwhile条件を表す
func While10(num int)string{
	s_string := fmt.Sprintf("最初は%d\n", num)
	if num <10{ // while num <10　に相当
		for num <10{
			num++
			s_string += fmt.Sprintf("%d\n", num)
		}
	}else if num >10{//閉じカッコとelseを離さない
		for num >10{ // while num >10　に相当
			num--
		s_string += fmt.Sprintf("%d\n", num)
		}
	}else{ 
		s_string +=fmt.Sprintf("今も%d\n", num)
	}
	return s_string
}

func Forever(limit int)string{
	i:=0
	for{ //無限ループ
		i++
		if i > limit{ //上限より1大きい整数でやめる
			return fmt.Sprintf("%dでやめました", i)
		}
	}
}

func Div3(num int)string{
	s_string := "3は"

	switch num{ //numについてswitchする
	case 0: //numの値が0の場合
		s_string += "0では割れません"
	case 1:
		s_string += "1で割る意味はあまりない"
	case 2:
		s_string += "2で割ると1と1/2"
	case 3:
		s_string += "3で割るとちょうど1"
	default://どこにもひっかからなかった場合
		if num%3==0 { //約分できる場合
			s_string += fmt.Sprintf("%dで割ると1/%d", num,num/3)
		}else{
			s_string += fmt.Sprintf("%dで割ると3/%d", num,num)
		}
	}

	return s_string
}

func DivBy3(num int)string{
	s_string := fmt.Sprintf("%dを3で割る", num)
	m := num%3 //3で割った余り。条件に入ってくる

	switch { //条件を変数ひとつの値に絞らない
	case num <1://「式」をcase文のあとに置ける
		s_string += "のは考えない"
	case num< 3: //ここまではnumの条件
		s_string += fmt.Sprintf("と%d/3", num)
	case m==0: //ここだけがmの条件
		s_string += fmt.Sprintf("と%d", num/3)
	default:
		s_string += fmt.Sprintf("と%dと%d/3", (num-m)/3, m)
	}

	return s_string
}