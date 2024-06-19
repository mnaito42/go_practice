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

func DescribeMaxPointMember(members []data.Member)string{
	s_string := "有効ポイント最大の方は\n"

	mpm := data.MaxPointMember(members) //引数として受け取ったmember

	s_string += fmt.Sprintf("%sさん", mpm.Name)
	s_string += "\n"

	return s_string
}