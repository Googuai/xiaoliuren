package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"xiaoliuren/solarlunar"
)

func main() {
	fmt.Println(`
	─┐ ┬┬┌─┐┌─┐┬  ┬┬ ┬┬─┐┌─┐┌┐┌
	┌┴┬┘│├─┤│ ││  ││ │├┬┘├┤ │││
	┴ └─┴┴ ┴└─┘┴─┘┴└─┘┴└─└─┘┘└┘
	                           author:googuai											  
	`)
	list := []string{"运势", "财富", "感情", "事业", "身体", "神鬼", "行人"}
	flag.Parse()
	s := flag.Arg(0)

	if !IsContain(list, s) {
		// fmt.Println("请从下列选项中认真选择所求之事：运势 财富 感情 事业 身体 神鬼 行人")

		fmt.Println("Error！！！输入错误！！！")
		fmt.Println("-------------------------------")
		fmt.Println("./xiaoliuren.exe 运势/财富/感情/事业/身体/神鬼/行人")
		return
	}

	var zhi int
	// 获取当前时间
	nowString := time.Now().Format("2006-01-02 15:04:05")
	// fmt.Println(nowString)
	currentDate := strings.Split(nowString, " ")[0]
	currentTime := strings.Split(nowString, " ")[1]
	// fmt.Println("公历时日：", currentDate, currentTime)

	lunarDate := solarlunar.SolarToChineseLunar(currentDate)
	fmt.Println(lunarDate)

	_, lunarMonth, lunarDay, _, _ := solarlunar.GetCalculateLunar()(currentDate)
	zhi += lunarMonth
	zhi += lunarDay
	lunarHour := solarlunar.CalculateTime(strings.Split(currentTime, ":")[0])
	zhi += lunarHour

	zhi = zhi % 6
	// fmt.Println(zhi)

	// 解
	switch zhi {
	case 1:
		fmt.Println(`
大安
身未动时，属木青龙．谋事一五七，贵人西南，冲犯东方，大人青面阴神，小孩婆祖六畜惊。
解曰：大安事事昌，求财在坤方，失物去不远，宅舍保安康，行人身未动，病者主无妨，将军回田野，仔细更推详。
玥沣解：
卜到大安，属吉卦，凡事都可以得到安康，但是此为静卦，宜静不宜动。
-----------------------------------
		`)
		getDetail(0, s)
	case 2:
		fmt.Println(`
留连
卒未归时，属水元武，谋事二八十，贵人南方，冲犯北方，大人乌面夫人，小孩游路亡魂。
解曰：留连事难成，求谋日未明，官事凡宜缓，去者未回程，失物南方见，急讨方心称，更须防口舌，人口且平平。
玥沣解：
卜到留连，属凶卦，代表凡事阻碍、迟滞，此卦更不宜有过大动作，凡事宜守。
-----------------------------------
		`)
		getDetail(1, s)
	case 3:
		fmt.Println(`
速喜
人便至时，属火未雀，谋事三六九，贵人西南，冲犯南方，大心火箭将军，小孩婆祖动物惊。 
解曰：速喜喜来临，求财向南行，失物申未午，逢人路上寻，官事有福德，病者无祸侵，田宅六畜吉，行人有信音。
玥沣解：
卜到速喜为吉卦，代表凡事皆有喜讯，而且很快就会到来。
-----------------------------------
		`)
		getDetail(2, s)
	case 4:
		fmt.Println(`
赤口
官事凶时，属金白虎，谋事四七十，贵人东方，冲犯西方，大人金神七煞，小孩迷魂童子。
解曰：赤口主口舌，官非切宜防，失物速速讨，行人有惊慌，六畜多作怪，病者出西方，更须防咀咒，诚恐染瘟。
玥沣解：
卜到赤口为凶卦，代表运势多舛，而且诸多纷争亦有口舌之祸。
-----------------------------------
		`)
		getDetail(3, s)
	case 5:
		fmt.Println(`
小吉
人来喜时，属水六合，谋事一五七，贵人西南，冲犯东方，大人无主家神，小孩婆祖六畜惊。
解曰：小吉最吉昌，路上好商量，阴人来报喜，失物在坤方， 
	行人即便至，交关甚是强，凡事皆和合，病者叩穹苍。
玥沣解：
卜到小吉为吉卦，代表凡事皆吉，但是不如大安的安稳也不如速喜快速，而是介于两者中间。
-----------------------------------
		`)
		getDetail(4, s)
	case 0:
		fmt.Println(`
空亡	
音信稀时，属土勾陈，主事三六九，贵人北方，冲犯厝地，大人土压夫人，小孩土瘟神煞。
解曰：空亡事不详，阴人多乖张，求财无利益，行人有灾殃，失物寻不见，官事有刑伤，病人逢暗鬼，解禳保安康。
玥沣解：
卜到空亡为凶卦，代表凡事秽暗不明，内心不安，运途起伏。
-----------------------------------
		`)
		getDetail(5, s)
	}

}

func getDetail(zhi int, s string) {
	details := []map[string]string{{"运势": "目前运势还不错，有稳定成长的情况，但不宜躁进。", "财富": "求财可，但是目前不宜扩张，只能够守住旧业。", "感情": "若为女子问则好，感情顺遂。若为男子问则较差，感情虽稳，但是以无新鲜感，会出现点小问题。", "事业": "目前工作稳定，可得上司赏识，但切勿锋芒太露。", "身体": "身体没有大病，但须注意病由口入，或因过度操劳而得病。", "神鬼": "大安为解灾之神，鬼神之事问题不大，若是小孩为自身惊吓所致，若是大人则为冲犯东方之煞神或犯土煞。", "行人": "人平安，但目前不愿与自身连络。"}, {"运势": "目前运势低迷，心情不开朗，凡事受阻。", "财富": "求财不可得，此为破财之卦，且有被人影响破财之现象。", "感情": "双方沟通不良、冷战、或者一方过于强势，感情不得平衡。", "事业": "被上司盯或者被人扯后腿，小人之卦。", "身体": "肠胃不舒服或者精神压力太大所得之病。", "神鬼": "小孩子主要被过路游魂所煞到，大人为冲犯女性鬼神。", "行人": "人平安，但目前仍流连忘返。"}, {"运势": "目前运势渐开，要积极的行动就可以如愿。", "财富": "求财可得，但有先破财而后得财或者先得财后破财之兆，若得到钱财就必须赶快脱身。", "感情": "若是刚开始的感情，则为热恋。若是已经持续一段时间，则为口舌。", "事业": "工作得利，但须注意文件上的疏失。", "身体": "心脏、血液循环有问题或者头部、脑压的问题，但是问题不大。", "神鬼": "小孩子被动物吓到或者被女性阴神冲犯，大人为冲犯男性鬼神。", "行人": "人已经快到了。"}, {"运势": "目前运势不明，若有大计划就要赶快实施、不要拖延，则可成功。若卜小事则不成。", "财富": "大起大落之财，求财不易。", "感情": "感情纷争多，或女方身体有疾病。", "事业": "若为武职或者粗重行业则顺，若为文职则不顺。", "身体": "胸口、支气管，或者有血光之灾，且赤口也有流行疾病的意义。", "神鬼": "犯到择日凶神，或者被人诅咒索害。", "行人": "所问之人目前有困难或者有事情纠缠。"}, {"运势": "目前运势不错，保持目前状况就会越来越好。", "财富": "求财可得，而且有因人得财之兆。 ", "感情": "若没有感情，则可因他人介绍而得。若有感情，则恋情顺利。", "事业": "工作不错，但须注意处理公司财务之事，以及与下属沟通之事。", "身体": "肝胆之疾病和消化系统，但是问题不大。", "神鬼": "小孩子被动物吓到或者被女性阴神冲犯，大人为冲犯家中祖先。", "行人": "人已经快到了。"}, {"运势": "目前运势不佳，自身拿不定主意，无所适从，可多听取他人之意见，切莫随意就做判断。", "财富": "求财难得，保守为要。", "感情": "双方争执多，且有因他人问题或者介入而争执之事。", "事业": "工作失利，容易被人陷害或者暗中耳语，或者因他人问题而让自己工作失利。", "身体": "脾胃出毛病，或者神经系统出问题，也有因灵界而生病之兆。", "神鬼": "家中阳宅或者阴宅出问题，导致冲犯。", "行人": "人在途中遇到困难或灾厄而难到。"}}
	fmt.Println(details[zhi][s])
}

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
