package data

import "fmt"

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

func Describe(m Member)string{
	return fmt.Sprintf("%sさんのポイントは%d点、有効ポイントは%.2f点",
	m.Name, m.Point, Effective(m))
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

func CreateFriendMember(member Member, name string )Member{ //引数は構造体の「コピー」
	member.Name = name//コピーのフィールドを変更
	return member
}


