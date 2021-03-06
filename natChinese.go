package main

import (
	"fmt"
)

func chineseNames() []string {
	// fmt.Println("start1")
	names := []string{"NO CHINESE NAMES ARRAY"}
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	switch genderIndex {
	case 0:
		names = []string{
			//MALE
			"AI",
			"AIGUO",
			"BAI",
			"BINGWEN",
			"BO",
			"BOHAI",
			"BOJING",
			"BOLIN",
			"BOQIN",
			"CHANG",
			"CHANGMING",
			"CHANGPU",
			"CHAO",
			"CHEN",
			"CHENG",
			"CHONGLIN",
			"CHUANLI",
			"DA",
			"DELUN",
			"DEMING",
			"DINGXIANG",
			"DONG",
			"DONGHAI",
			"DUYI",
			"ENLAI",
			"FA",
			"FAN",
			"FANG",
			"FENG",
			"FENGGE",
			"FU",
			"GANG",
			"GEMING",
			"GEN",
			"GUANG",
			"GUANGLI",
			"GUI",
			"GUIREN",
			"GUOLIANG",
			"GUOWEI",
			"GUOZHI",
			"HAI",
			"HAN",
			"HE",
			"HENG",
			"HONG",
			"HONGHUI",
			"HUAN",
			"HUANG",
			"HUI",
			"HUILIANG",
			"HUIQING",
			"HUIZHONG",
			"JIAN",
			"JIANG",
			"JIANGUO",
			"JIANJUN",
			"JIANYU",
			"JIAYI",
			"JING",
			"JINGGUO",
			"JINHAI",
			"JUNJIE",
			"KANG",
			"LEI",
			"LI",
			"LIANG",
			"LING",
			"LIQIN",
			"LIU",
			"LIWEI",
			"LONGWEI",
			"MEILIN",
			"MENGYAO",
			"MINGLI",
			"MINGYU",
			"NIANZU",
			"NIU",
			"PEIZHI",
			"PENG",
			"PENGFEI",
			"PING",
			"QI",
			"QIANG",
			"QIAO",
			"QING",
			"QINGSHAN",
			"QINGSHENG",
			"QIQIANG",
			"QIU",
			"QUAN",
			"RENSHU",
			"RONG",
			"RU",
			"SHAN",
			"SHANYUAN",
			"SHAOQING",
			"SHEN",
			"SHI",
			"SHIHONG",
			"SHIRONG",
			"SHOUSHAN",
			"SHUN",
			"SHUNYUAN",
			"SONG",
			"TAO",
			"TENGFEI",
			"TINGGUANG",
			"TINGFENG",
			"WEI",
			"WEIMIN",
			"WEISHENG",
			"WEIYUAN",
			"WEIZHE",
			"WEN",
			"WENCHENG",
			"WENYAN",
			"XIANG",
			"XIANLIANG",
			"XIAOBO",
			"XIAODAN",
			"XIAOJIAN",
			"XIAOLI",
			"XIAOSHENG",
			"XIAOTONG",
			"XIAOSI",
			"XIAOWEN",
			"XIN",
			"XING",
			"XIU",
			"XUE",
			"XUEQIN",
			"YAN",
			"YANG",
			"YANLIN",
			"YAOTING",
			"YAOZU",
			"YE",
			"YI",
			"YING",
			"YINGJIE",
			"YONG",
			"YONGLIANG",
			"YONGNIAN",
			"YONGRUI",
			"YONGZHENG",
			"YOU",
			"YUAN",
			"YUANJUN",
			"YUN",
			"YUNRU",
			"YUSHENG",
			"ZEDONG",
			"ZEMIN",
			"ZHAOHUI",
			"ZHEN",
			"ZHENGSHENG",
			"ZHENGZHONG",
			"ZHIQIANG",
			"ZHONG",
			"ZHU",
			"ZIAN",
			"ZIHAO",
			"ZIXIN",
		}
		//////////////////////////////////////////////////////////////////////////////////////////////////////
	case 1:
		names = []string{
			//FEMALE
			"AI",
			"BAOZHAI",
			"BIYU",
			"CHANGCHANG",
			"CHANGYING",
			"CHEN",
			"CHENGUANG",
			"CHUNHUA",
			"CHUNTAO",
			"CUIFEN",
			"DAIYU",
			"DANDAN",
			"DONGMEI",
			"EHUANG",
			"FAN",
			"FANG",
			"FENFANG",
			"GUANG",
			"HONG",
			"HUALING",
			"HUAN",
			"HUIAN",
			"HUIFANG",
			"HUILING",
			"HUIQING",
			"JIA",
			"JIAO",
			"JIAYI",
			"JIAYING",
			"JIE",
			"JINGFEI",
			"JINGHUA",
			"JINJING",
			"JU",
			"JUAN",
			"LAN",
			"LANFEN",
			"LANYING",
			"LI",
			"LIFEN",
			"LIHUA",
			"LIJUAN",
			"LILING",
			"LIN",
			"LING",
			"LIQIN",
			"LIQIU",
			"LIU",
			"LULI",
			"MEI",
			"MEIFEN",
			"MEIFENG",
			"MEIHUI",
			"MEILI",
			"MEILIN",
			"MEIRONG",
			"MEIXIANG",
			"MEIXIU",
			"MINGXIA",
			"MINGZHU",
			"NGO - KWANG",
			"NING",
			"NIU",
			"NÜYING",
			"O - HUANG",
			"PEIZHI",
			"QI",
			"QIANG",
			"QIAO",
			"QIAOLIAN",
			"QING",
			"QINGGE",
			"QINGLING",
			"QINGZHAO",
			"QIU",
			"QIUYUE",
			"RONG",
			"RUOLAN",
			"SHAN",
			"SHU",
			"SHUANG",
			"SHUCHUN",
			"SONG",
			"SUYIN",
			"TING",
			"TUNGMEI",
			"WEN",
			"WENLING",
			"WENQIAN",
			"XIA",
			"XIANG",
			"XIAODAN",
			"XIAOFAN",
			"XIAOHUI",
			"XIAOLI",
			"XIAOLIAN",
			"XIAOLING",
			"XIAOQING",
			"XIFENG",
			"XINGJUAN",
			"XIU",
			"XIULAN",
			"XIURONG",
			"XIUYING",
			"XUE",
			"YA",
			"YAN",
			"YANMEI",
			"YANYU",
			"YING",
			"YINGTAI",
			"YU",
			"YUAN",
			"YUANJUN",
			"YUE",
			"YUN",
			"ZHENZHEN",
			"ZHILAN",
			"ZHU",
			"ZONGYING",
		}
		//////////////////////////////////////////////////////////////////////////////////////////////////////
	case 2:
		names = []string{
			//UNISEX
			"Ah",
			"An",
			"Bai",
			"Bao",
			"Bo",
			"Chang",
			"Chen",
			"Cheng",
			"Chin",
			"Chun",
			"Fu",
			"Guo",
			"Hai",
			"Heng",
			"Huan",
			"Hui",
			"Jian",
			"Jiang",
			"Jun",
			"Kun",
			"Lei",
			"Lim",
			"Liu",
			"Min",
			"Ming",
			"Mu",
			"Ping",
			"Qiu",
			"Rong",
			"Ru",
			"Shi",
			"Shui",
			"Su",
			"Tai",
			"Tu",
			"Wei",
			"Wen",
			"Wu",
			"Xiaodan",
			"Xiaojian",
			"Xiaosheng",
			"Xun",
			"Yi",
			"Yun",
			"Yusheng",
			"Zan",
			"Zheng",
			"Zhi",
			"Zhong",
			"Zhou",
		}
		//////////////////////////////////////////////////////////////////////////////////////////////////////
	default:
		fmt.Println("Error in CHINESE Names Array")
		return nil
	}
	// fmt.Println("finish1")
	return names
}

//############################################################################################################################################################################################
func chineseSurNames() []string {
	surNames := []string{"NO CHINESE SURNAMES ARRAY"}
	surNames = []string{
		"BAI",
		"CHAN",
		"CHEN",
		"CHEUNG",
		"CHONG",
		"CHOU",
		"CHOW",
		"CHU",
		"DU",
		"FAN",
		"GUAN",
		"GUO",
		"HAN",
		"HOU",
		"HSU",
		"HU",
		"HUANG",
		"JIANG",
		"JIN",
		"KUANG",
		"KWAN",
		"KWOK",
		"LAM",
		"LAU",
		"LEE",
		"LI",
		"LIAO",
		"LIM",
		"MIN",
		"LIN",
		"LIU",
		"LU",
		"MA",
		"MAH",
		"MAN",
		"NG",
		"PAN",
		"RUAN",
		"SONG",
		"SUN",
		"SUNG",
		"TAN",
		"TANG",
		"WANG",
		"WEN",
		"WONG",
		"WU",
		"XU",
		"XUN",
		"YANG",
		"YEUNG",
		"YU",
		"YUEN",
		"ZHANG",
		"ZHAO",
		"ZHENG",
		"ZHOU",
		"ZHU",
	}
	return surNames
}
