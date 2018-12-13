package eew

import (
	"fmt"
	"time"
	"strings"
	"strconv"
)

type Telegram struct {
	Code_type    string
	Section      string
	Msg_type     string
	Warn_time    string
	Command_code string
	Eq_time      string
	Eq_id        string
	Warn_type    string
	Warn_code    string
	Warn_num     string
	Shinou       string
	Shinou_lat   float64
	Shinou_lng   float64
	Shinou_dpth  int64
	Magnitude    float64
	Shindo       string
	Ebis         []Ebi
}
type Ebi struct {
	Shinou  string
	Shindo1 string
	Shindo2 string
	Time    string
	Arrive  string
}

var Code_type = map[string]string{
	"35": "最大予測震度のみ",
	"36": "Ｍ、最大予測震度及び主要動到達予測時刻",
	"37": "Ｍ、最大予測度及び主要動到達予測時刻（確度大）",
	"39": "キャンセル報",
	"47": "一般向け速報",
	"48": "キャンセル報",
}

//発信官署
var Section = map[string]string{
	"01": "札幌",
	"02": "仙台",
	"03": "東京",
	"04": "大阪",
	"05": "福岡",
	"06": "沖縄",
}

//訓練等の識別符
var Msg_type = map[string]string{
	"00": "通常",
	"01": "訓練",
	"10": "取り消し",
	"11": "訓練取り消し",
	"20": "参考/テスト",
	"30": "コードのみ配信試験",
}

//震度
var Shindo = map[string]string{
	"01": "震度1",
	"02": "震度2",
	"03": "震度3",
	"04": "震度4",
	"5-": "震度5弱",
	"5+": "震度5強",
	"6-": "震度6弱",
	"6+": "震度6強",
	"07": "震度7",
	"//": "不明",
}

//データの確からしさ
var Rk1 = map[string]string{
	"1": "P 波／S 波レベル越え、またはテリトリー法（１点）",
	"2": "テリトリー法（2 点）",
	"3": "グリッドサーチ法（3 点／4 点）",
	"4": "グリッドサーチ法（5 点）",
	"5": "防災科研システム（4 点以下、または精度情報なし）",
	"6": "防災科研システム（5 点以上）",
	"7": "EPOS（海域〔観測網外〕）",
	"8": "EPOS（内陸〔観測網内〕）",
	"9": "予備",
	"/": "不明、未設定時、キャンセル時",
}

//震源の深さの確からしさ
var Rk2 = map[string]string{
	"1":  "P 波／S 波レベル越え、またはテリトリー法（1 点）",
	"2":  "テリトリー法（2 点）",
	"3":  "グリッドサーチ法（3 点／4 点）",
	"4":  "グリッドサーチ法（5 点）",
	"5":  "防災科研システム（4 点以下、または精度情報なし）",
	"6":  "防災科研システム（5 点以上）",
	"7":  "EPOS（海域〔観測網外〕）",
	"8":  "EPOS（内陸〔観測網内〕）",
	"9":  "予備",
	"/ ": "不明、未設定時、キャンセル時",
}

//マグニチュードの確からしさ
var Rk3 = map[string]string{
	"1": "未定義",
	"2": "防災科研システム",
	"3": "全点（最大5 点）P 相",
	"4": "P 相／全相混在",
	"5": "全点（最大5 点）全相",
	"6": "EPOS",
	"7": "未定義",
	"8": "P 波／S 波レベル越え",
	"9": "予備",
	"/": "不明、未設定時、キャンセル時",
}

//地震の発生場所
var Rt1 = map[string]string{
	"0": "陸",
	"1": "海",
	"/": "不明",
}

//最大予測震度の変化
var Rc1 = map[string]string{
	"0": "ほとんど変化なし",
	"1": "最大予測震度が1.0 以上大きくなった。",
	"2": "最大予測震度が1.0 以上小さくなった。",
	"/": "不明、未設定時、キャンセル時",
}

//最大予測震度の変化の理由
var Rc2 = map[string]string{
	"0": "変化なし",
	"1": "主としてＭが変化したため(1.0 以上)。",
	"2": "主として震源位置が変化したため(10km 以上)。",
	"3": "Ｍ及び震源位置が変化したため(1 と2 の複合条件)。",
	"4": "震源の深さが変化したため(上記のいずれにもあてはまらず、30km 以上の変化)。",
	"/": "不明、未設定時、キャンセル時",
}

//到達状況
var Ebiyy = map[string]string{
	"00": "未到達",
	"01": "既に到着と予想",
}

//タイプ
var Warn_type = map[string]string{
	"NCN":  "高度利用者向け",
	"NPN":  "一般向け",
	"NCPN": "一般向け",
	"PRC":  "強い揺れが推定される地域の追加",
	"EBI":  "予想震度と到達時間",
	"CAI":  "新たに追加された強い揺れが推定される地方",
	"CPI":  "新たに追加された強い揺れが推定される県",
	"CBI":  "新たに追加された強い揺れが推定される地域",
	"PAI":  "強い揺れが推定される地方",
	"PPI":  "強い揺れが推定される県",
	"PBI":  "強い揺れが推定される地域",
}

//震央コード
var Shinou_code = map[string]string{
	"000": "",
	"100": "石狩支庁北部",
	"101": "石狩支庁中部",
	"102": "石狩支庁南部",
	"105": "渡島支庁北部",
	"106": "渡島支庁東部",
	"107": "渡島支庁西部",
	"110": "檜山支庁",
	"115": "後志支庁北部",
	"116": "後志支庁東部",
	"117": "後志支庁西部",
	"120": "空知支庁北部",
	"121": "空知支庁中部",
	"122": "空知支庁南部",
	"125": "上川支庁北部",
	"126": "上川支庁中部",
	"127": "上川支庁南部",
	"130": "留萌支庁中北部",
	"131": "留萌支庁南部",
	"135": "宗谷支庁北部",
	"136": "宗谷支庁南部",
	"140": "網走支庁網走地方",
	"141": "網走支庁北見地方",
	"142": "網走支庁紋別地方",
	"145": "胆振支庁西部",
	"146": "胆振支庁中東部",
	"150": "日高支庁西部",
	"151": "日高支庁中部",
	"152": "日高支庁東部",
	"155": "十勝支庁北部",
	"156": "十勝支庁中部",
	"157": "十勝支庁南部",
	"160": "釧路支庁北部",
	"161": "釧路支庁中南部",
	"165": "根室支庁北部",
	"166": "根室支庁中部",
	"167": "根室支庁南部",
	"180": "北海道南西沖",
	"181": "北海道西方沖",
	"182": "石狩湾",
	"183": "北海道北西沖",
	"184": "宗谷海峡",
	"186": "国後島付近",
	"187": "択捉島付近",
	"188": "北海道東方沖",
	"189": "根室半島南東沖",
	"190": "釧路沖",
	"191": "十勝沖",
	"192": "浦河沖",
	"193": "苫小牧沖",
	"194": "内浦湾",
	"195": "宗谷東方沖",
	"196": "網走沖",
	"197": "択捉島南東沖",
	"200": "青森県津軽北部",
	"201": "青森県津軽南部",
	"202": "青森県三八上北地方",
	"203": "青森県下北地方",
	"210": "岩手県沿岸北部",
	"211": "岩手県沿岸南部",
	"212": "岩手県内陸北部",
	"213": "岩手県内陸南部",
	"220": "宮城県北部",
	"221": "宮城県南部",
	"222": "宮城県中部",
	"230": "秋田県沿岸北部",
	"231": "秋田県沿岸南部",
	"232": "秋田県内陸北部",
	"233": "秋田県内陸南部",
	"240": "山形県庄内地方",
	"241": "山形県最上地方",
	"242": "山形県村山地方",
	"243": "山形県置賜地方",
	"250": "福島県中通り",
	"251": "福島県浜通り",
	"252": "福島県会津",
	"280": "津軽海峡",
	"281": "山形県沖",
	"282": "秋田県沖",
	"283": "青森県西方沖",
	"284": "陸奥湾",
	"285": "青森県東方沖",
	"286": "岩手県沖",
	"287": "宮城県沖",
	"288": "三陸沖",
	"289": "福島県沖",
	"300": "茨城県北部",
	"301": "茨城県南部",
	"309": "千葉県南東沖",
	"310": "栃木県北部",
	"311": "栃木県南部",
	"320": "群馬県北部",
	"321": "群馬県南部",
	"330": "埼玉県北部",
	"331": "埼玉県南部",
	"332": "埼玉県秩父地方",
	"340": "千葉県北東部",
	"341": "千葉県北西部",
	"342": "千葉県南部",
	"349": "房総半島南方沖",
	"350": "東京都２３区",
	"351": "東京都多摩東部",
	"352": "東京都多摩西部",
	"360": "神奈川県東部",
	"361": "神奈川県西部",
	"370": "新潟県上越地方",
	"371": "新潟県中越地方",
	"372": "新潟県下越地方",
	"378": "新潟県下越沖",
	"379": "新潟県上中越沖",
	"380": "富山県東部",
	"381": "富山県西部",
	"390": "石川県能登地方",
	"391": "石川県加賀地方",
	"400": "福井県嶺北",
	"401": "福井県嶺南",
	"411": "山梨県中・西部",
	"412": "山梨県東部・富士五湖",
	"420": "長野県北部",
	"421": "長野県中部",
	"422": "長野県南部",
	"430": "岐阜県飛騨地方",
	"431": "岐阜県美濃東部",
	"432": "岐阜県美濃中西部",
	"440": "静岡県伊豆地方",
	"441": "静岡県東部",
	"442": "静岡県中部",
	"443": "静岡県西部",
	"450": "愛知県東部",
	"451": "愛知県西部",
	"460": "三重県北部",
	"461": "三重県中部",
	"462": "三重県南部",
	"469": "三重県南東沖",
	"471": "茨城県沖",
	"472": "関東東方沖",
	"473": "千葉県東方沖",
	"475": "八丈島東方沖",
	"476": "八丈島近海",
	"477": "東京湾",
	"478": "相模湾",
	"480": "伊豆大島近海",
	"481": "伊豆半島東方沖",
	"482": "三宅島近海",
	"483": "新島・神津島近海",
	"485": "駿河湾",
	"486": "駿河湾南方沖",
	"487": "遠州灘",
	"489": "三河湾",
	"490": "伊勢湾",
	"492": "若狭湾",
	"493": "福井県沖",
	"494": "石川県西方沖",
	"495": "能登半島沖",
	"497": "富山湾",
	"498": "佐渡付近",
	"499": "東海道南方沖",
	"500": "滋賀県北部",
	"501": "滋賀県南部",
	"510": "京都府北部",
	"511": "京都府南部",
	"520": "大阪府北部",
	"521": "大阪府南部",
	"530": "兵庫県北部",
	"531": "兵庫県南東部",
	"532": "兵庫県南西部",
	"540": "奈良県",
	"550": "和歌山県北部",
	"551": "和歌山県南部",
	"560": "鳥取県東部",
	"562": "鳥取県中部",
	"563": "鳥取県西部",
	"570": "島根県東部",
	"571": "島根県西部",
	"580": "岡山県北部",
	"581": "岡山県南部",
	"590": "広島県北部",
	"591": "広島県南東部",
	"592": "広島県南西部",
	"600": "徳島県北部",
	"601": "徳島県南部",
	"610": "香川県東部",
	"611": "香川県西部",
	"620": "愛媛県東予",
	"621": "愛媛県中予",
	"622": "愛媛県南予",
	"630": "高知県東部",
	"631": "高知県中部",
	"632": "高知県西部",
	"673": "土佐湾",
	"674": "紀伊水道",
	"675": "大阪湾",
	"676": "播磨灘",
	"677": "瀬戸内海中部",
	"678": "安芸灘",
	"679": "周防灘",
	"680": "伊予灘",
	"681": "豊後水道",
	"682": "山口県北西沖",
	"683": "島根県沖",
	"684": "鳥取県沖",
	"685": "隠岐島近海",
	"686": "兵庫県北方沖",
	"687": "京都府沖",
	"688": "淡路島付近",
	"689": "和歌山県南方沖",
	"700": "山口県北部",
	"701": "山口県東部",
	"702": "山口県西部",
	"710": "福岡県福岡地方",
	"711": "福岡県北九州地方",
	"712": "福岡県筑豊地方",
	"713": "福岡県筑後地方",
	"720": "佐賀県北部",
	"721": "佐賀県南部",
	"730": "長崎県北部",
	"731": "長崎県南西部",
	"732": "長崎県島原半島",
	"740": "熊本県阿蘇地方",
	"741": "熊本県熊本地方",
	"742": "熊本県球磨地方",
	"743": "熊本県天草・芦北地方",
	"750": "大分県北部",
	"751": "大分県中部",
	"752": "大分県南部",
	"753": "大分県西部",
	"760": "宮崎県北部平野部",
	"761": "宮崎県北部山沿い",
	"762": "宮崎県南部平野部",
	"763": "宮崎県南部山沿い",
	"770": "鹿児島県薩摩地方",
	"771": "鹿児島県大隅地方",
	"783": "五島列島近海",
	"784": "天草灘",
	"785": "有明海",
	"786": "橘湾",
	"787": "鹿児島湾",
	"790": "種子島近海",
	"791": "日向灘",
	"793": "奄美大島近海",
	"795": "壱岐・対馬近海",
	"796": "福岡県北西沖",
	"797": "薩摩半島西方沖",
	"798": "トカラ列島近海",
	"799": "奄美大島北西沖",
	"820": "大隅半島東方沖",
	"821": "九州地方南東沖",
	"822": "種子島南東沖",
	"823": "奄美大島北東沖",
	"850": "沖縄本島近海",
	"851": "南大東島近海",
	"852": "沖縄本島南方沖",
	"853": "宮古島近海",
	"854": "石垣島近海",
	"855": "石垣島南方沖",
	"856": "西表島付近",
	"857": "与那国島近海",
	"858": "沖縄本島北西沖",
	"859": "宮古島北西沖",
	"860": "石垣島北西沖",
	"900": "台湾付近",
	"901": "東シナ海",
	"902": "四国沖",
	"903": "鳥島近海",
	"904": "鳥島東方沖",
	"905": "オホーツク海南部",
	"906": "サハリン西方沖",
	"907": "日本海北部",
	"908": "日本海中部",
	"909": "日本海西部",
	"911": "父島近海",
	"912": "千島列島",
	"913": "千島列島南東沖",
	"914": "北海道南東沖",
	"915": "東北地方東方沖",
	"916": "小笠原諸島西方沖",
	"917": "硫黄島近海",
	"918": "小笠原諸島東方沖",
	"919": "南海道南方沖",
	"920": "薩南諸島東方沖",
	"921": "本州南方沖",
	"922": "サハリン南部付近",
	"930": "北西太平洋",
	"932": "マリアナ諸島",
	"933": "黄海",
	"934": "朝鮮半島南部",
	"935": "朝鮮半島北部",
	"936": "中国東北部",
	"937": "ウラジオストク付近",
	"938": "シベリア南部",
	"939": "サハリン近海",
	"940": "アリューシャン列島",
	"941": "カムチャツカ半島付近",
	"942": "北米西部",
	"943": "北米中部",
	"944": "北米東部",
	"945": "中米",
	"946": "南米西部",
	"947": "南米中部",
	"948": "南米東部",
	"949": "北東太平洋",
	"950": "南太平洋",
	"951": "インドシナ半島付近",
	"952": "フィリピン付近",
	"953": "インドネシア付近",
	"954": "グアム付近",
	"955": "ニューギニア付近",
	"956": "ニュージーランド付近",
	"957": "オーストラリア付近",
	"958": "シベリア付近",
	"959": "ロシア西部",
	"960": "ロシア中部",
	"961": "ロシア東部",
	"962": "中央アジア",
	"963": "中国西部",
	"964": "中国中部",
	"965": "中国東部",
	"966": "インド付近",
	"967": "インド洋",
	"968": "中東",
	"969": "ヨーロッパ西部",
	"970": "ヨーロッパ中部",
	"971": "ヨーロッパ東部",
	"972": "地中海",
	"973": "アフリカ西部",
	"974": "アフリカ中部",
	"975": "アフリカ東部",
	"976": "北大西洋",
	"977": "南大西洋",
	"978": "北極付近",
	"979": "南極付近",
}

func Decoder(str string) Telegram {
	var to,from Telegram
	from = Reader(str)

	to.Code_type=Code_type[from.Code_type]
	to.Section=Section[from.Section]
	to.Msg_type=Msg_type[from.Msg_type]
	to.Warn_type=Warn_type[from.Warn_type]

	wt,_ := strconv.ParseInt(from.Warn_time,10,32)
	to.Warn_time=fmt.Sprintln(time.Unix(wt, 0).Format("2006-01-02 15:04:05"))
	to.Command_code=from.Command_code
	et,_ := strconv.ParseInt(from.Eq_time,10,32)
	to.Eq_time=fmt.Sprintln(time.Unix(et, 0).Format("2006-01-02 15:04:05"))
	to.Eq_id=from.Eq_id
	to.Warn_num=from.Warn_num
	
	to.Shinou=Shinou_code[from.Shinou]
	to.Shindo = Shindo[from.Shindo]

	to.Shinou_lat=from.Shinou_lat
	to.Shinou_lng=from.Shinou_lng
	to.Shinou_dpth=from.Shinou_dpth
	to.Magnitude=from.Magnitude
	
	return to
}

func Reader(str string) Telegram {
	var msg Telegram
	for _, line := range strings.Split(string(str), "\n") {
		offset := 0
		length := len(line)
		vec := []byte(line)
		if length == 0 {
			//nope
		} else if length == 25 {
			// 1 : [37 03 00 110311144702 C11](25)
			msg.Code_type = string(vec[0:2])
			msg.Section = string(vec[3:5])
			msg.Msg_type = string(vec[6:8])
			msg.Warn_time = string(vec[9:21])
			msg.Command_code = string(vec[23:25])
		} else if length == 12 {
			// 2 : [110311144616](12)
			msg.Eq_time = line
		} else if length == 46 {
			// 3 : [ND20110311144640 NCN009 JD////////////// JN///](46)
			msg.Eq_id = string(vec[2:16])
			msg.Warn_type = string(vec[17:20])
			msg.Warn_num = string(vec[20:23])
		} else if length == 48 {
			// 4 : [288 N381 E1429 010 76 5- RK66444 RT11/// RC0////](48)
			msg.Shinou = string(vec[0:3])
			lat,_ := strconv.ParseFloat(string(vec[5:8]),64)
			msg.Shinou_lat = lat/10
			lng,_ := strconv.ParseFloat(string(vec[10:14]),64)
			msg.Shinou_lng = lng/10
			dpth,_ := strconv.ParseInt(string(vec[15:18]),10,32)
			msg.Shinou_dpth = dpth
			mag,_ := strconv.ParseFloat(string(vec[19:21]),64)
			msg.Magnitude = mag/10
			msg.Shindo = string(vec[22:24])
		} else if length == 63 {
			// 5 : [EBI 222 S5-04 ////// 11 220 S5-04 ////// 11 211 S5-04 ////// 11](63)
			offset = 4
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
			offset = offset + 20
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
			offset = offset + 20
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
		} else if length == 59 {
			// 6 : [210 S5-04 144703 10 221 S5-04 144703 10 213 S0404 ////// 11](59)
			offset = 0
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
			offset = offset + 20
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
			offset = offset + 20
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
		} else if length == 43 {
			// [EBI 251 S0404 ////// 01 250 S0404 ////// 01](43)
			offset = 4
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
			offset = offset + 20
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
		} else if length == 39 {
			//10 : [340 S0403 144739 00 331 S0403 144748 00](39)
			offset = 0
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
			offset = offset + 20
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
		} else if length == 23 {
			//  [EBI 341 S0403 ////// 01](23)
			offset = 4
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
		} else if length == 19 {
			//11 : [340 S0403 144739 00](19)
			offset = 0
			msg.Ebis = append(msg.Ebis, Ebi{string(vec[offset+0 : offset+3]), string(vec[offset+5 : offset+7]), string(vec[offset+7 : offset+9]), string(vec[offset+10 : offset+16]), string(vec[offset+17 : offset+19])})
		} else if length == 5 {
			//13 : [9999=](5)
			return msg
		} else {
			fmt.Printf("Error: unkown line [%s](%d)\n", line, len(line))
		}
	}
	return msg
}
