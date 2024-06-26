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