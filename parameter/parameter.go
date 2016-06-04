package parameter

// pluginでやりとりするパラメータだよ
// event/command共通 parameter構造体
// eventがcommandになってcomanndのレスポンスがeventのレスポンスになる

type Dest {
	DestName string // 正規表現マッチだよ
	DestArea DestArea     //frontendに投げるのか、backエンドに投げるのか？ allもありよ const定義
	DstAttrs map[string]string  // 自由に属性情報入れる
}

tyoe Src {
	SrcName string // inしたら入ってる
	SrcAttrs map[string]string  // 自由に属性情報入れる
}

tyoe Media {
	mimeType string // mime type
	mediaAttrs map[string]string  // 自由に属性情報入れる
}

tyoe Data {
	data []byte  // データの中身だよ
	dataAttrs map[string]string // 自由に属性情報入れられる
}

type ReqParam {
	*Dest
	*Src
	*Media
	*Data
}

type ResParam {
	*Dest
	*Src
	*Media
	*Data
}

DstAttr よく使われるkey
userName : ユーザ名
SrcAttr よく使われるkey
userName : ユーザ名
MediaAttr よく使われるkey
textLineType : one|multi

DataAttr  使われ方例
コマンドライン引数を入れる
