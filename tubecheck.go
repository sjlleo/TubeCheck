package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time" 
	"io/ioutil"
	"os"
	"net"
)

func ParseIP(s string) int {
	ip := net.ParseIP(s)
	if ip == nil {
		return 0
	}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '.':
			return 4
		case ':':
			return 6
		}
	}
	return 0
}

func RequestIP(requrl string, ip string) (int,string, string, string) {

	urlValue, err := url.Parse(requrl)
	if err != nil {
		return 400,"","",""
	}
	host := urlValue.Host
	if ip == "" {
		ip = host
	}
	newrequrl := strings.Replace(requrl, host, ip, 1)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{ServerName: host},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       5 * time.Second,
	}
	req, err := http.NewRequest("GET", newrequrl, nil)
	if err != nil {
		return 400,"","",""
	}
	req.Host = host
	req.Header.Set("USER-AGENT", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return 400,"","",""
	}
	defer resp.Body.Close()

	s, _ := ioutil.ReadAll(resp.Body)
	response := string(s)
	EndLocation := strings.Index(response,"\n")
	response = response[1:EndLocation]
	EndLocation = strings.Index(response,"=> ")
	response = response[EndLocation+3:]
	EndLocation = strings.Index(response," ")
	response = response[:EndLocation]
	EndLocation = strings.Index(response,"-")


	if(EndLocation == -1) {
		method := "Youtube Video Server"
		EndLocation = strings.Index(response,".")
		airCode := response[EndLocation+1:]
		return 200,method,"",airCode
	} else {
		method := "Google Global CacheCDN (ISP Cooperation)"
		isp := response[:EndLocation]
		airCode := response[EndLocation+1:]
		return 200,method,isp,airCode
	}
}

func RequestIPRegion(requrl string, ip string) string {

	urlValue, err := url.Parse(requrl)
	if err != nil {
		return "error"
	}
	host := urlValue.Host
	if ip == "" {
		ip = host
	}
	newrequrl := strings.Replace(requrl, host, ip, 1)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{ServerName: host},
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse },
		Timeout:       5 * time.Second,
	}
	req, err := http.NewRequest("GET", newrequrl, nil)
	if err != nil {
		return "error"
	}
	req.Host = host
	req.Header.Set("USER-AGENT", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return "error"
	}
	defer resp.Body.Close()

	s, _ := ioutil.ReadAll(resp.Body)
	response := string(s)
	EndLocation := strings.Index(response,"\"countryCode\"")
	if (EndLocation != -1){
		return response[EndLocation+15:EndLocation+17]
	}
	return "null"
}

func RegionCheck(module string) string{
	var ipv4, ipv6 string
	dns:= "www.youtube.com"
	url:= "https://www.youtube.com/red"
	ns, err := net.LookupHost(dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return "error"
	}

	switch {
	case len(ns) != 0:
		for _, n := range ns {

			if ParseIP(n) == 4 {
				ipv4 = n
			}
			if ParseIP(n) == 6 {
				ipv6 = "[" + n + "]"
			}
		}
	}

	switch {
	case module == "ipv4":
		return RequestIPRegion(url,ipv4)
	case module == "ipv6":
		return RequestIPRegion(url,ipv6)
	}
	return ""
}

func findAirCode(code string) string {
	i, v := 0, ""
	airPortName := []string{"日本 大阪","日本 东京","韩国 金浦", "加拿大 渥太华","加拿大 蒙特利尔","加拿大 温哥华","加拿大 卡尔加里","加拿大 埃德蒙顿","加拿大 多伦多","美国  华盛顿","美国  阿伦敦","美国  阿尔伯克斯","美国  阿特兰大","美国  奥斯汀","美国  卡拉马祖","美国  哈特福德","美国  伯明翰","美国  纳什维尔","美国  博伊西","美国  波士顿","美国  布朗斯韦尔","美国  巴吞鲁日","美国  巴特尔克里克","美国  布法罗","美国  巴尔的摩","美国  哥伦比亚","美国  阿克伦肯顿","美国  查塔诺加","美国  芝加哥","美国  查尔斯顿","美国  锡达拉皮兹","美国  克利夫兰","美国  夏洛特","美国  哥伦布","美国  科珀斯克里斯蒂","美国  辛辛那提","美国  代顿","美国  丹佛","美国  达拉斯","美国  得梅因","美国  底特律","美国  埃尔帕索","美国   伊利/伊利湖","美国  纽瓦克","美国  埃文斯韦尔","美国  劳德代尔堡","美国  弗林特","美国  育空堡","美国  大急流域","美国  斯波坎","美国  格林斯伯勒","美国  格林维尔","美国  格林贝","美国  哈里斯堡","美国  休斯敦","美国  亨茨维尔","美国  火奴鲁鲁（夏威夷州的首府)","美国  威奇托","美国  威尔明顿","美国  印第安纳波利斯","美国  杰克逊","美国  杰克逊威尔","美国  拉斯维加斯","美国  洛杉机","美国  列克星敦","美国  小石城","美国  林肯","美国  拉雷多","美国  堪萨斯城","美国  奥兰多","美国  孟菲斯","美国  麦卡伦","美国  迈阿密","美国  堪萨斯","美国  密尔沃基","美国  麦迪逊","美国  明利阿波利斯 ","美国  新奥尔良","美国  莫比尔","美国  纽约","美国  俄克拉荷马城","美国  奥马哈","美国  诺福克","美国  奥兰多","美国  西棕榈滩","美国  波特兰","美国  费城","美国  费尼克斯","美国  皮奥里亚","美国  匹兹堡","美国  彭萨科拉","美国  普罗维登斯","美国  达拉谟","美国  里士满","美国  里诺","美国  罗彻斯特","美国  圣迭戈","美国  圣安东尼奥","美国  萨凡纳","美国  南本德","美国  路易斯维尔","美国  西雅图","美国  西雅图","美国  旧金山","美国  斯普林菲尔德","美国  什里夫波特","美国  盐湖城","美国  萨克拉门托","美国  圣路易斯","美国  塔尔萨","美国  锡拉丘兹","美国  托莱多","美国  坦帕","美国  塔尔萨","美国  图森","美国  诺科斯韦尔","墨西哥墨西哥城","墨西哥瓜达拉哈拉","危地马拉危地马拉","洪都拉斯  特古西加尔巴","萨尔瓦多圣萨尔瓦多","尼加拉瓜   马拉瓜","哥斯达黎加圣何塞","巴拿马 巴拿马城","巴哈马拿骚","古巴哈瓦那","古巴圣地亚哥","牙买加金斯敦","海地太子港","多米尼加圣多明各","波多黎各 圣胡安","多米尼克罗索","格林纳达圣乔治","巴巴多斯 布里奇顿","特立尼达和多巴哥  西班牙港","哥伦比亚 圣菲波哥达","委内瑞拉加拉加斯","圭亚那乔治敦","苏里南帕拉马里博","法属圭那亚卡宴","巴西 巴西利亚","巴西 库里蒂巴","巴西 阿雷格里港","巴西 马卤斯","巴西 里约热内卢","巴西 圣保罗","厄瓜多尔基多","厄瓜多尔瓜尔基尔","秘鲁 利马","玻利维亚   苏克雷","巴拉圭亚松森","乌拉圭蒙得维的亚","阿根廷布宜诺斯艾利斯","智利 安托法加斯塔","智利 圣地亚哥","拉丁美洲瓜德罗普","英国 伦敦","英国 阿伯丁","英国 伯明翰","英国 伯恩茅斯","英国 布里斯托尔","英国 加地夫","英国 爱丁堡","英国 埃克塞特","英国 格拉斯哥 ","英国 利物浦","英国 曼彻斯特","英国 诺里奇","英国 普利茅斯","英国 南安普敦","英国 布里斯托尔","英国 克里伊登","英国 考文垂","英国 利兹","英国 朴茨茅斯","英国 纽卡斯尔","英国 亨伯赛德郡","英国 格拉斯哥 ","英国 诺丁汉","北爱尔兰 贝尔法斯特","爱尔兰 都柏林 ","爱尔兰 科克","爱尔兰 香农","比利时布鲁塞尔","比利时安特卫普","比利时奥斯坦德","卢森堡卢森堡","荷兰阿姆斯特丹","荷兰鹿特丹","荷兰爱恩德霍芬","荷兰恩斯赫德","丹麦 哥本哈根","丹麦 阿尔伯格","丹麦 奥胡斯","丹麦 比灵顿","德国柏林","德国慕尼黑","德国不来梅","德国汉诺威","德国杜塞尔多夫","德国法兰克福","德国莱比锡","德国杜伊斯堡","德国斯图加特","德国汉堡","德国爱尔福特","德国明斯特","德国纽伦堡","德国德累斯顿","德国萨尔布吕肯","德国科隆","德国多特蒙德","德国比勒费尔德","德国开姆尼茨","德国埃森","德国波恩","德国重聚","法国 巴黎","法国 马塞","法国 里昂","法国 波尔多","法国 里尔","法国 图卢兹","法国 南特","法国 牟罗兹","法国 蒙彼利埃","法国 格勒诺布尔","法国 鲁昂","法国 尼斯","法国 斯特拉斯堡","法国 凡尔塞","法属波利尼帕皮提 ","摩纳哥摩纳哥","瑞士伯尔尼","瑞士日内瓦","瑞士苏黎世","瑞士巴塞尔","安道尔安道尔","西班牙马德里","西班牙阿利坎特","西班牙巴塞罗那","西班牙华伦西亚(巴伦比亚)","西班牙塞尔维亚","西班牙马拉加","西班牙巴利亚多利德","葡萄牙里斯本","葡萄牙波尔图","意大利罗马","意大利安齐奥","意大利安科纳","意大利布林迪西","意大利博洛尼亚","意大利巴里","意大利热那亚","意大利米兰","意大利米兰市郊","意大利拿坡里","意大利威尼斯","意大利佛罗伦萨","意大利都灵","意大利的里雅斯特","意大利卡塔尼亚","意大利塔兰托","意大利比萨","意大利墨西拿","意大利维罗纳","希腊雅典","希腊塞萨洛尼基","奥地利维也纳","奥地利林茨","奥地利格拉茨","奥地利萨尔斯堡","奥地利因斯布鲁克","捷克 布拉格","芬兰赫尔辛基","瑞典斯德哥尔摩","瑞典海尔辛堡","瑞典哥德堡","瑞典马尔默","瑞典北雪平","挪威奥斯陆","阿尔巴尼亚地拉那","马其顿斯科普里","保加利亚索非亚","南斯拉夫 贝尔格莱德","罗马尼亚 布加勒斯特","摩尔多瓦 基希纳乌","克罗地亚 萨格勒布","斯洛文尼亚卢布尔雅那","匈牙利布达佩斯","斯洛伐克布拉迪斯拉发","波兰 华沙","波兰 克拉克夫","波兰 格但斯克","立陶宛维尔纽斯","拉托维亚里加","爱沙尼亚 塔林","冰岛雷克亚未克","俄罗斯 莫斯科","俄罗斯 圣彼得堡(列宁格勒)","白俄罗斯明斯克","乌克兰基辅","波黑萨拉热窝","伊朗 德黑兰","伊朗 阿巴达","阿富汗喀布尔","科威特 科威特","沙特阿拉伯利雅得","沙特阿拉伯吉达","沙特阿拉伯达曼","也门萨那","也门亚丁","伊拉克巴格达","黎巴嫩 贝鲁特","黎巴嫩 巴林","阿联酋阿布扎比","阿联酋迪拜","阿联酋沙加","卡塔尔 多哈","以色列  耶路撒冷","以色列  特拉维夫","叙利亚大马士革","约旦 安曼","土耳其安卡拉","土耳其阿达那","土耳其布尔萨","土耳其伊兹密尔","土耳其伊斯坦布尔","巴林  巴林","塞浦路斯尼科西亚","塞浦路斯拉纳卡","阿塞拜疆 巴库","亚美尼亚 埃里温","格鲁吉亚 第比利斯","阿曼  马斯喀特","土库曼斯坦阿什哈巴德","塔吉克斯坦杜尚别","哈萨克斯坦卡拉干达","吉尔吉斯斯坦比什凯克","乌兹别克斯坦 塔什干","埃及开罗","苏丹 喀土穆","苏丹 马斯喀特阿曼","埃塞俄比亚亚的斯亚贝巴","吉布提 吉布提","肯尼亚 内罗毕","利比亚 的黎波里","阿尔及利亚  阿尔及尔","阿尔及利亚  安纳巴","突尼斯  突尼斯","摩洛哥拉巴特","摩洛哥卡萨布兰卡","乍得 恩贾梅纳","尼日尔 尼亚美","尼日利亚 阿布贾","尼日利亚 拉各斯","尼日利亚 哈科特港 ","马里 巴马科","布基纳法索瓦加杜古","贝宁 科托努","多哥 洛美","加纳  阿克拉","科特迪瓦阿穆苏克罗","科特迪瓦阿比让","塞拉利昂弗里敦","利比里亚蒙罗维亚","几内亚科纳克里","塞内加尔 达喀尔","冈比亚 班珠尔","马里塔尼亚努瓦克肖特","中非共和国班吉","喀麦隆雅温得","赤道几内亚马拉博","乌干达坎帕拉","卢旺达 基加利","坦桑尼亚达累斯萨拉姆","布隆迪布琼布拉","刚果布拉柴维尔","加蓬 利伯维尔","圣多美和普林西比圣多美","莫桑比克马普托","马拉维利隆圭","赞比亚卢萨卡","津巴布韦哈拉雷","安哥拉罗安达","博茨瓦纳哈伯罗内","纳米比亚温得和克","南非约翰内斯堡","南非德班","南非开普敦","毛里求斯 毛里求斯","马达加斯加塔那那利佛","科摩罗莫罗尼","马埃塞舌尔群岛","努瓦克肖特毛里塔尼亚","中国香港","中国台北","中国高雄","朝鲜 平壤","韩国 汉城仁川","韩国 釜山","日本 东京","日本 大阪","日本 名古屋","日本 福冈","日本 横滨","日本 广岛","日本 冲绳岛","日本 仙台","日本 札幌","菲律宾马尼拉","菲律宾宿务","菲律宾达沃","马来西亚吉隆坡","马来西亚槟城","马来西亚凌家卫岛","马来西亚哥大基纳巴卢","马来西亚古晋","马来西亚怡宝","马来西亚新山","马来西亚哥打巴鲁","马来西亚诗巫","马来西亚山打根","文莱 斯里巴加湾市","新加坡 新加坡/樟宜  ","印度尼西亚 雅加达","印度尼西亚 棉兰","印度尼西亚 泗水","印度尼西亚 登巴萨","印度尼西亚 乌戒潘当","印度尼西亚 坤甸","东帝汶帝力","越南胡志明市","越南河内","越南海防","老挝 万象","泰国曼谷","泰国清迈","泰国合艾","泰国普吉","泰国","缅甸 仰光","缅甸 曼德勒","柬埔寨 金边","孟加拉达卡","孟加拉吉大港","印度 新德里","印度 孟买","印度 加尔各答","印度 马德拉斯","印度 班加罗尔","印度 荷兰安得列斯群岛","印度 海德拉巴","尼泊尔加德满都","巴基斯坦伊斯兰堡","巴基斯坦卡拉奇","巴基斯坦拉合尔","巴基斯坦白沙瓦","斯里兰卡 科伦坡","马尔代夫 马累","蒙古乌兰巴托","澳大利亚 堪培拉","澳大利亚 墨尔本","澳大利亚 阿德莱得","澳大利亚 达尔文","澳大利亚 凯恩斯","澳大利亚 布里斯班","澳大利亚 珀斯","澳大利亚 悉尼","新西兰惠灵顿","新西兰奥克兰 ","新西兰利特尔顿(基督城)","巴布亚新几内亚莫尔兹比港","斐济群岛 苏瓦","基里巴斯 塔拉瓦","所罗门群岛霍尼亚拉","汤加 努瓦阿洛法","萨摩亚  阿皮亚","图瓦卢富纳富提","密克罗尼西亚帕利基尔","瓦努阿图维拉港"}
	airPortCode := []string{"KIX","NRT","GMP","YOW","YMQ/YUL","YVR","YYC","YEG","YTO/YYZ","WAS/IAD","ABE","ABQ","ATL","AUS","AZO","BDL","BHM","BNA","BOI","BOS","BRO","BTR","BTL","BUF","BWI","CAE","CAK","CHA","CHI/ORD","CHS","CID","CLE","CLT","CMH","CRP","CVG","DAY","DEN","DFW","DSM","DTW","ELP","ERI","EWR","EVV","FLL","FNT","FWA","GRR","GEG","GSO","GSP","GRB","HAR","HOU/IAH","HSV","HNL","ICT","ILM","IND","JAN","JAX","LAS","LAX","LEX","LIT","LNK","LRD","MCI","MCO","MEM","MFE","MIA","MKC","MKE","MSN","MSP","MSY","MOB","NYC/JFK","OKC","OMA","ORF","ORL","PBI","PDX","PHL/PHA","PHX","PIA","PIT","PNS","PVD","RDU","RIC","RNO","ROC","SAN","SAT","SAV","SBN","SDF","SEA","BFI","SFO","SGF","SHV","SLC","SMF","STL","TUL","SYR","TOL","TPA","TUL","TUS","TYS","MEX","GDL/MEX","GUA","TGU","SAL","MGA","SJO","PTY","NAS","HAV","SCU","KIN","PAP","SDQ","SJU","ROX","GND","BGI","POS","BOG","CCS","GEO","PBM","CAY","BSB","CWB","POA","MAO","RIO","SAO","UIO","GYE","LIM","SRE","ASU","MVD","BUE","ANF","SCL","PTP","LON/LHR","ABZ","BHX","BOH","BRS","CWL","EDI","EXT","GLA","LPL","MAN","NWI","PLH","SOU","BRS","CDQ","CVT","LBA","PME","NCL","HUY","PIK","EMA","BFS","DUB","ORK","SNN","BRU","ANR","OST","LUX","AMS","RTM","EIN","ENS","CPH","ALL","AAR","BLL","BER/TXL","MUC","BRE","HAJ","DUS","FRA","LEJ","DUI","STR","HAM","ERF","FMO","NUE","DRS","SCN","CGN","DTM","BFE","ZTZ","ESS","BON","RUN","PAR/CDG","MRS","LYS","BOD","LIL","TLS","NTE","MLH","MPL","GNB","URO","NCE","SXB","XVE","PPT","XMM/GRZ","BRN","GVA","ZRH","BSL","ALV","MAD","ALC","BCN","VLC","SVQ","AGP","VLL","LIS","OPO","ROM","AHO","AOI","BDS","BLQ","BRI","GOA","MIL/MXP","SWK","NAP","VCE","FLR","TRN","TRS","CTA","TAR","PSA","QME","VRN","ATH","SKG","VIE","LNZ","GRZ","SZG","INN","PRG","HEL","STO/ARN","AGH","GOT","MMA/MMX","NRK","OSL","TIA","SKP","SOF","BEG","BUH","KIV","ZAG","LJU","BUD","BTS","WAW","KRK","GDN","VNO","RIX","TLL","REK","MOW","LED","MSQ","IEV/KBP","SJJ","THR","ABD","KBL","KWI","RUH","JED","DMM","SAH","ADE","BGW","BEY","BAH","AUH","DXB","SHJ","DOH","JRD","TLV","DAM","AMM","ANK","ADA","BTZ","IZM","IST","BAH","NIC","LCA","BAK","EVN","TBS","MSH","ASB","DYU","KGF","FRU","TAS","CAI","KRT","MCT","ADD","JIB","NBO","TIP","ALG","AAE","TUN","RBA","CAS","NDJ","NIM","ABV","LOS","PHC","BKO","OUA","COO","LFW","ACC","ASK","ABJ","HGS","MLW","CKF","DKR","BJL","KLA","BGF","YAO","SSG","KLA","KGL","DAR","BJM","BZV","LBV","TMS","MPM","LLW","LUN","HRE","LAD","GBE","WDH","JNB","DUR","CPT","MRU","TNR","YVA","SEZ","NKC","HKG","TPE","KHH","FNJ","SEL/ICN","PUS","TYONRT","KIX/OSA","NGO","FUK","YOK","HIJ","OKA","SDJ","SPA","MNL","HEB","DVO","KUL","PEN","LGK","BKI","KCH","IPH","JHB","KBR","SBW","SDK","BWN","SIN","JKT","MES","SUB","DPS","UPG","PNK","DIL","SGN","HAN","HPH","VTE","BKK","CEI","HDY","HKT","NSI","RGN","MDL","PNH","DAC","CGP","DEL","BOM","CCU","MAA","BLR","SXM","HYD","KTM","ISB","KHI","LHE","PEW","CMB","MLE","ULN","CBR","MEL","ADL","DRN","CNS","BNE","PER","SYD","WLG","AKL","CHC","POM","SUV","TRW","HIR","TBU","APW","FUN","KSA","VLI"}
	codeTune := code[:]
	for ; i < len(codeTune); i++ {
		if(codeTune[i] >= '0' && codeTune[i] <= '9') {
			break
		}
	}
	code = strings.ToUpper(code[:i])
	for i,v = range airPortCode {
		if strings.Contains(code,v) {
			return airPortName[i]
		}
	}
	return code
}

func FindCountry(Code string) string {
	countryName := []string{"无信息","美国", "阿富汗", "奥兰群岛", "阿尔巴尼亚", "阿尔及利亚", "美属萨摩亚", "安道尔", "安哥拉", "安圭拉", "南极洲", "安提瓜和巴布达", "阿根廷", "亚美尼亚", "阿鲁巴", "澳大利亚", "奥地利", "阿塞拜疆", "巴哈马", "巴林", "孟加拉国", "巴巴多斯", "白俄罗斯", "比利时", "伯利兹", "贝宁", "百慕大", "不丹", "玻利维亚", "波黑", "博茨瓦纳", "布维岛", "巴西", "英属印度洋领地", "文莱", "保加利亚", "布基纳法索", "布隆迪", "柬埔寨", "喀麦隆", "加拿大", "佛得角", "开曼群岛", "中非", "乍得", "智利", "中国", "圣诞岛", "科科斯（基林）群岛", "哥伦比亚", "科摩罗", "刚果（布）", "刚果（金）", "库克群岛", "哥斯达黎加", "科特迪瓦", "克罗地亚", "古巴", "塞浦路斯", "捷克", "丹麦", "吉布提", "多米尼克", "多米尼加", "厄瓜多尔", "埃及", "萨尔瓦多", "赤道几内亚", "厄立特里亚", "爱沙尼亚", "埃塞俄比亚", "福克兰群岛（马尔维纳斯）", "法罗群岛", "斐济", "芬兰", "法国", "法属圭亚那", "法属波利尼西亚", "法属南部领地", "加蓬", "冈比亚", "格鲁吉亚", "德国", "加纳", "直布罗陀", "希腊", "格陵兰", "格林纳达", "瓜德罗普", "关岛", "危地马拉", "格恩西岛", "几内亚", "几内亚比绍", "圭亚那", "海地", "赫德岛和麦克唐纳岛", "梵蒂冈", "洪都拉斯", "香港", "匈牙利", "冰岛", "印度", "印度尼西亚", "伊朗", "伊拉克", "爱尔兰", "英国属地曼岛", "以色列", "意大利", "牙买加", "日本", "泽西岛", "约旦", "哈萨克斯坦", "肯尼亚", "基里巴斯", "朝鲜", "韩国", "科威特", "吉尔吉斯斯坦", "老挝", "拉脱维亚", "黎巴嫩", "莱索托", "利比里亚", "利比亚", "列支敦士登", "立陶宛", "卢森堡", "澳门", "前南马其顿", "马达加斯加", "马拉维", "马来西亚", "马尔代夫", "马里", "马耳他", "马绍尔群岛", "马提尼克", "毛利塔尼亚", "毛里求斯", "马约特", "墨西哥", "密克罗尼西亚联邦", "摩尔多瓦", "摩纳哥", "蒙古", "黑山", "蒙特塞拉特", "摩洛哥", "莫桑比克", "缅甸", "纳米比亚", "瑙鲁", "尼泊尔", "荷兰", "荷属安的列斯", "新喀里多尼亚", "新西兰", "尼加拉瓜", "尼日尔", "尼日利亚", "纽埃", "诺福克岛", "北马里亚纳", "挪威", "阿曼", "巴基斯坦", "帕劳", "巴勒斯坦", "巴拿马", "巴布亚新几内亚", "巴拉圭", "秘鲁", "菲律宾", "皮特凯恩", "波兰", "葡萄牙", "波多黎各", "卡塔尔", "留尼汪", "罗马尼亚", "俄罗斯联邦", "卢旺达", "圣赫勒拿", "圣基茨和尼维斯", "圣卢西亚", "圣皮埃尔和密克隆", "圣文森特和格林纳丁斯", "萨摩亚", "圣马力诺", "圣多美和普林西比", "沙特阿拉伯", "塞内加尔", "塞尔维亚", "塞舌尔", "塞拉利昂", "新加坡", "斯洛伐克", "斯洛文尼亚", "所罗门群岛", "索马里", "南非", "南乔治亚岛和南桑德韦奇岛", "西班牙", "斯里兰卡", "苏丹", "苏里南", "斯瓦尔巴岛和扬马延岛", "斯威士兰", "瑞典", "瑞士", "叙利亚", "台湾", "塔吉克斯坦", "坦桑尼亚", "泰国", "东帝汶", "多哥", "托克劳", "汤加", "特立尼达和多巴哥", "突尼斯", "土耳其", "土库曼斯坦", "特克斯和凯科斯群岛", "图瓦卢", "乌干达", "乌克兰", "阿联酋", "英国", "美国本土外小岛屿", "乌拉圭", "乌兹别克斯坦", "瓦努阿图", "委内瑞拉", "越南", "英属维尔京群岛", "美属维尔京群岛", "瓦利斯和富图纳", "西撒哈拉", "也门", "赞比亚", "津巴布韦"}
	countryCode := []string{"null","us", "af", "ax", "al", "dz", "as", "ad", "ao", "ai", "aq", "ag", "ar", "am", "aw", "au", "at", "az", "bs", "bh", "bd", "bb", "by", "be", "bz", "bj", "bm", "bt", "bo", "ba", "bw", "bv", "br", "io", "bn", "bg", "bf", "bi", "kh", "cm", "ca", "cv", "ky", "cf", "td", "cl", "cn", "cx", "cc", "co", "km", "cg", "cd", "ck", "cr", "ci", "hr", "cu", "cy", "cz", "dk", "dj", "dm", "do", "ec", "eg", "sv", "gq", "er", "ee", "et", "fk", "fo", "fj", "fi", "fr", "gf", "pf", "tf", "ga", "gm", "ge", "de", "gh", "gi", "gr", "gl", "gd", "gp", "gu", "gt", "gg", "gn", "gw", "gy", "ht", "hm", "va", "hn", "hk", "hu", "is", "in", "id", "ir", "iq", "ie", "im", "il", "it", "jm", "jp", "je", "jo", "kz", "ke", "ki", "kp", "kr", "kw", "kg", "la", "lv", "lb", "ls", "lr", "ly", "li", "lt", "lu", "mo", "mk", "mg", "mw", "my", "mv", "ml", "mt", "mh", "mq", "mr", "mu", "yt", "mx", "fm", "md", "mc", "mn", "me", "ms", "ma", "mz", "mm", "na", "nr", "np", "nl", "an", "nc", "nz", "ni", "ne", "ng", "nu", "nf", "mp", "no", "om", "pk", "pw", "ps", "pa", "pg", "py", "pe", "ph", "pn", "pl", "pt", "pr", "qa", "re", "ro", "ru", "rw", "sh", "kn", "lc", "pm", "vc", "ws", "sm", "st", "sa", "sn", "rs", "sc", "sl", "sg", "sk", "si", "sb", "so", "za", "gs", "es", "lk", "sd", "sr", "sj", "sz", "se", "ch", "sy", "tw", "tj", "tz", "th", "tl", "tg", "tk", "to", "tt", "tn", "tr", "tm", "tc", "tv", "ug", "ua", "ae", "gb", "um", "uy", "uz", "vu", "ve", "vn", "vg", "vi", "wf", "eh", "ye", "zm", "zw"}
	for i, v := range countryCode {
		if strings.Contains(Code, v) {
			return countryName[i]
		}
	}
	return Code
}

func main() {
	var ipv4, ipv6 string
	var NextLineSignal bool = false
	dns:= "redirector.googlevideo.com"

	fmt.Println("** Youtube 检测小工具 v1.0 Beta By \033[1;36m@sjlleo\033[0m **")

	ns, err := net.LookupHost(dns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	switch {
	case len(ns) != 0:
		for _, n := range ns {

			if ParseIP(n) == 4 {
				ipv4 = n
			}
			if ParseIP(n) == 6 {
				ipv6 = "[" + n + "]"
			}
		}
	}

	responseCode,method,isp,airCode := RequestIP("https://redirector.googlevideo.com/report_mapping",ipv4)
	if (responseCode == 200) {
		NextLineSignal = true
		fmt.Printf("\033[0;36m[IPv4]\033[0m\n连接方式: %s\n",method)
		if (isp != "") {
			fmt.Printf("ISP运营商: %s\n",strings.ToUpper(isp))
		}
		fmt.Printf("视频缓存节点地域: %s(%s)\n",findAirCode(airCode),strings.ToUpper(airCode))
		RegionCode := RegionCheck("ipv4")
		fmt.Printf("Youtube识别地域: %s(%s)\n",FindCountry(strings.ToLower(RegionCode)),RegionCode)
	}

	responseCode,method,isp,airCode = RequestIP("https://redirector.googlevideo.com/report_mapping",ipv6)
	if (responseCode == 200) {
		if NextLineSignal {
			fmt.Print("\n")
		}
		fmt.Printf("\033[0;36m[IPv6]\033[0m\n连接方式: %s\n",method)
		if (isp != "") {
			fmt.Printf("ISP运营商: %s\n",strings.ToUpper(isp))
		}
		fmt.Printf("视频缓存节点地域: %s(%s)\n",findAirCode(airCode),strings.ToUpper(airCode))
		RegionCode := RegionCheck("ipv6")
		fmt.Printf("Youtube识别地域: %s(%s)\n",FindCountry(strings.ToLower(RegionCode)),RegionCode)
	}

	return
}