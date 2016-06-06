package parameter

// pluginでやりとりするパラメータだよ
// event/command共通 parameter構造体
// eventがcommandになってcomanndのレスポンスがeventのレスポンスになる

type Dest struct {
	DestName string            // 正規表現マッチだよ
	DestArea DestArea          //frontendに投げるのか、backエンドに投げるのか？ allもありよ const定義
	DstAttrs map[string]string // 自由に属性情報入れる
}

type Src struct {
	SrcName  string            // inしたら入ってる
	SrcAttrs map[string]string // 自由に属性情報入れる
}

type Media struct {
	mimeType   string            // mime type
	mediaAttrs map[string]string // 自由に属性情報入れる
}

type Data struct {
	data      []byte            // データの中身だよ
	dataAttrs map[string]string // 自由に属性情報入れられる
}

type ReqParam struct {
	*Dest
	*Src
	*Media
	*Data
}

type ResParam struct {
	*Dest
	*Src
	*Media
	*Data
}

/*
DstAttr よく使われるkey
userName : ユーザ名
SrcAttr よく使われるkey
userName : ユーザ名
MediaAttr よく使われるkey
textLineType : one|multi

DataAttr  使われ方例
コマンドライン引数を入れる
*/
