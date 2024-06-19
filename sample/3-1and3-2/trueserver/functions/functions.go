package functions //main以外のパッケージ名


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
