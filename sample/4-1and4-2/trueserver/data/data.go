package data

import "fmt"

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


