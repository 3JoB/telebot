package telebot

import (
	"github.com/3JoB/unsafeConvert"
	"github.com/goccy/go-json"
)

const dataCurrencies = `{"AED":{"code":"AED","title":"United Arab Emirates Dirham","symbol":"AED","native":"\u062f.\u0625.\u200f","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"367","max_amount":"3673200"},"AFN":{"code":"AFN","title":"Afghan Afghani","symbol":"AFN","native":"\u060b","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"7554","max_amount":"75540495"},"ALL":{"code":"ALL","title":"Albanian Lek","symbol":"ALL","native":"Lek","thousands_sep":".","decimal_sep":",","symbol_left":false,"space_between":false,"exp":2,"min_amount":"10908","max_amount":"109085036"},"AMD":{"code":"AMD","title":"Armenian Dram","symbol":"AMD","native":"\u0564\u0580.","thousands_sep":",","decimal_sep":".","symbol_left":false,"space_between":true,"exp":2,"min_amount":"48398","max_amount":"483984962"},"ARS":{"code":"ARS","title":"Argentine Peso","symbol":"ARS","native":"$","thousands_sep":".","decimal_sep":",","symbol_left":true,"space_between":true,"exp":2,"min_amount":"3720","max_amount":"37202998"},"AUD":{"code":"AUD","title":"Australian Dollar","symbol":"AU$","native":"$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"139","max_amount":"1392750"},"AZN":{"code":"AZN","title":"Azerbaijani Manat","symbol":"AZN","native":"\u043c\u0430\u043d.","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"170","max_amount":"1702500"},"BAM":{"code":"BAM","title":"Bosnia & Herzegovina Convertible Mark","symbol":"BAM","native":"KM","thousands_sep":".","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"171","max_amount":"1715550"},"BDT":{"code":"BDT","title":"Bangladeshi Taka","symbol":"BDT","native":"\u09f3","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"8336","max_amount":"83367500"},"BGN":{"code":"BGN","title":"Bulgarian Lev","symbol":"BGN","native":"\u043b\u0432.","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"171","max_amount":"1716850"},"BND":{"code":"BND","title":"Brunei Dollar","symbol":"BND","native":"$","thousands_sep":".","decimal_sep":",","symbol_left":true,"space_between":false,"exp":2,"min_amount":"134","max_amount":"1349850"},"BOB":{"code":"BOB","title":"Bolivian Boliviano","symbol":"BOB","native":"Bs","thousands_sep":".","decimal_sep":",","symbol_left":true,"space_between":true,"exp":2,"min_amount":"687","max_amount":"6877150"},"BRL":{"code":"BRL","title":"Brazilian Real","symbol":"R$","native":"R$","thousands_sep":".","decimal_sep":",","symbol_left":true,"space_between":true,"exp":2,"min_amount":"377","max_amount":"3775397"},"CAD":{"code":"CAD","title":"Canadian Dollar","symbol":"CA$","native":"$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"132","max_amount":"1321950"},"CHF":{"code":"CHF","title":"Swiss Franc","symbol":"CHF","native":"CHF","thousands_sep":"'","decimal_sep":".","symbol_left":false,"space_between":true,"exp":2,"min_amount":"99","max_amount":"993220"},"CLP":{"code":"CLP","title":"Chilean Peso","symbol":"CLP","native":"$","thousands_sep":".","decimal_sep":",","symbol_left":true,"space_between":true,"exp":0,"min_amount":"666","max_amount":"6665199"},"CNY":{"code":"CNY","title":"Chinese Renminbi Yuan","symbol":"CN\u00a5","native":"CN\u00a5","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"674","max_amount":"6747298"},"COP":{"code":"COP","title":"Colombian Peso","symbol":"COP","native":"$","thousands_sep":".","decimal_sep":",","symbol_left":true,"space_between":true,"exp":2,"min_amount":"315595","max_amount":"3155950000"},"CRC":{"code":"CRC","title":"Costa Rican Col\u00f3n","symbol":"CRC","native":"\u20a1","thousands_sep":".","decimal_sep":",","symbol_left":true,"space_between":false,"exp":2,"min_amount":"60113","max_amount":"601130282"},"CZK":{"code":"CZK","title":"Czech Koruna","symbol":"CZK","native":"K\u010d","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"2251","max_amount":"22510978"},"DKK":{"code":"DKK","title":"Danish Krone","symbol":"DKK","native":"kr","thousands_sep":"","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"654","max_amount":"6545403"},"DOP":{"code":"DOP","title":"Dominican Peso","symbol":"DOP","native":"$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"5032","max_amount":"50329504"},"DZD":{"code":"DZD","title":"Algerian Dinar","symbol":"DZD","native":"\u062f.\u062c.\u200f","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"11872","max_amount":"118729869"},"EGP":{"code":"EGP","title":"Egyptian Pound","symbol":"EGP","native":"\u062c.\u0645.\u200f","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"1791","max_amount":"17912012"},"EUR":{"code":"EUR","title":"Euro","symbol":"\u20ac","native":"\u20ac","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"87","max_amount":"877155"},"GBP":{"code":"GBP","title":"British Pound","symbol":"\u00a3","native":"\u00a3","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"75","max_amount":"757605"},"GEL":{"code":"GEL","title":"Georgian Lari","symbol":"GEL","native":"GEL","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"266","max_amount":"2663750"},"GTQ":{"code":"GTQ","title":"Guatemalan Quetzal","symbol":"GTQ","native":"Q","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"768","max_amount":"7689850"},"HKD":{"code":"HKD","title":"Hong Kong Dollar","symbol":"HK$","native":"$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"784","max_amount":"7845505"},"HNL":{"code":"HNL","title":"Honduran Lempira","symbol":"HNL","native":"L","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"2427","max_amount":"24277502"},"HRK":{"code":"HRK","title":"Croatian Kuna","symbol":"HRK","native":"kn","thousands_sep":".","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"650","max_amount":"6506302"},"HUF":{"code":"HUF","title":"Hungarian Forint","symbol":"HUF","native":"Ft","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"27844","max_amount":"278440341"},"IDR":{"code":"IDR","title":"Indonesian Rupiah","symbol":"IDR","native":"Rp","thousands_sep":".","decimal_sep":",","symbol_left":true,"space_between":false,"exp":2,"min_amount":"1406555","max_amount":"14065550000"},"ILS":{"code":"ILS","title":"Israeli New Sheqel","symbol":"\u20aa","native":"\u20aa","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"366","max_amount":"3668230"},"INR":{"code":"INR","title":"Indian Rupee","symbol":"\u20b9","native":"\u20b9","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"7090","max_amount":"70900503"},"ISK":{"code":"ISK","title":"Icelandic Kr\u00f3na","symbol":"ISK","native":"kr","thousands_sep":".","decimal_sep":",","symbol_left":false,"space_between":true,"exp":0,"min_amount":"119","max_amount":"1195599"},"JMD":{"code":"JMD","title":"Jamaican Dollar","symbol":"JMD","native":"$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"13153","max_amount":"131539958"},"JPY":{"code":"JPY","title":"Japanese Yen","symbol":"\u00a5","native":"\uffe5","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":0,"min_amount":"109","max_amount":"1095549"},"KES":{"code":"KES","title":"Kenyan Shilling","symbol":"KES","native":"Ksh","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"10032","max_amount":"100322011"},"KGS":{"code":"KGS","title":"Kyrgyzstani Som","symbol":"KGS","native":"KGS","thousands_sep":"\u00a0","decimal_sep":"-","symbol_left":false,"space_between":true,"exp":2,"min_amount":"6982","max_amount":"69820300"},"KRW":{"code":"KRW","title":"South Korean Won","symbol":"\u20a9","native":"\u20a9","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":0,"min_amount":"1119","max_amount":"11190001"},"KZT":{"code":"KZT","title":"Kazakhstani Tenge","symbol":"KZT","native":"\u20b8","thousands_sep":"\u00a0","decimal_sep":"-","symbol_left":true,"space_between":false,"exp":2,"min_amount":"37767","max_amount":"377674954"},"LBP":{"code":"LBP","title":"Lebanese Pound","symbol":"LBP","native":"\u0644.\u0644.\u200f","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"150080","max_amount":"1500802255"},"LKR":{"code":"LKR","title":"Sri Lankan Rupee","symbol":"LKR","native":"\u0dbb\u0dd4.","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"18078","max_amount":"180789638"},"MAD":{"code":"MAD","title":"Moroccan Dirham","symbol":"MAD","native":"\u062f.\u0645.\u200f","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"955","max_amount":"9554850"},"MDL":{"code":"MDL","title":"Moldovan Leu","symbol":"MDL","native":"MDL","thousands_sep":",","decimal_sep":".","symbol_left":false,"space_between":true,"exp":2,"min_amount":"1703","max_amount":"17038967"},"MNT":{"code":"MNT","title":"Mongolian T\u00f6gr\u00f6g","symbol":"MNT","native":"MNT","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":true,"space_between":false,"exp":2,"min_amount":"261750","max_amount":"2617500000"},"MUR":{"code":"MUR","title":"Mauritian Rupee","symbol":"MUR","native":"MUR","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"3438","max_amount":"34384499"},"MVR":{"code":"MVR","title":"Maldivian Rufiyaa","symbol":"MVR","native":"MVR","thousands_sep":",","decimal_sep":".","symbol_left":false,"space_between":true,"exp":2,"min_amount":"1550","max_amount":"15501063"},"MXN":{"code":"MXN","title":"Mexican Peso","symbol":"MX$","native":"$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"1898","max_amount":"18988704"},"MYR":{"code":"MYR","title":"Malaysian Ringgit","symbol":"MYR","native":"RM","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"412","max_amount":"4124501"},"MZN":{"code":"MZN","title":"Mozambican Metical","symbol":"MZN","native":"MTn","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"6188","max_amount":"61889913"},"NGN":{"code":"NGN","title":"Nigerian Naira","symbol":"NGN","native":"\u20a6","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"36174","max_amount":"361749532"},"NIO":{"code":"NIO","title":"Nicaraguan C\u00f3rdoba","symbol":"NIO","native":"C$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"3241","max_amount":"32415503"},"NOK":{"code":"NOK","title":"Norwegian Krone","symbol":"NOK","native":"kr","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":true,"space_between":true,"exp":2,"min_amount":"851","max_amount":"8510100"},"NPR":{"code":"NPR","title":"Nepalese Rupee","symbol":"NPR","native":"\u0928\u0947\u0930\u0942","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"11299","max_amount":"112995016"},"NZD":{"code":"NZD","title":"New Zealand Dollar","symbol":"NZ$","native":"$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"146","max_amount":"1461850"},"PAB":{"code":"PAB","title":"Panamanian Balboa","symbol":"PAB","native":"B\/.","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"99","max_amount":"995290"},"PEN":{"code":"PEN","title":"Peruvian Nuevo Sol","symbol":"PEN","native":"S\/.","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"333","max_amount":"3331250"},"PHP":{"code":"PHP","title":"Philippine Peso","symbol":"PHP","native":"\u20b1","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"5260","max_amount":"52602981"},"PKR":{"code":"PKR","title":"Pakistani Rupee","symbol":"PKR","native":"\u20a8","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"13921","max_amount":"139214990"},"PLN":{"code":"PLN","title":"Polish Z\u0142oty","symbol":"PLN","native":"z\u0142","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"376","max_amount":"3764026"},"PYG":{"code":"PYG","title":"Paraguayan Guaran\u00ed","symbol":"PYG","native":"\u20b2","thousands_sep":".","decimal_sep":",","symbol_left":true,"space_between":true,"exp":0,"min_amount":"6013","max_amount":"60134502"},"QAR":{"code":"QAR","title":"Qatari Riyal","symbol":"QAR","native":"\u0631.\u0642.\u200f","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"364","max_amount":"3641101"},"RON":{"code":"RON","title":"Romanian Leu","symbol":"RON","native":"RON","thousands_sep":".","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"417","max_amount":"4172003"},"RSD":{"code":"RSD","title":"Serbian Dinar","symbol":"RSD","native":"\u0434\u0438\u043d.","thousands_sep":".","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"10391","max_amount":"103910127"},"RUB":{"code":"RUB","title":"Russian Ruble","symbol":"RUB","native":"\u0440\u0443\u0431.","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"6598","max_amount":"65986027"},"SAR":{"code":"SAR","title":"Saudi Riyal","symbol":"SAR","native":"\u0631.\u0633.\u200f","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"373","max_amount":"3732650"},"SEK":{"code":"SEK","title":"Swedish Krona","symbol":"SEK","native":"kr","thousands_sep":".","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"904","max_amount":"9047896"},"SGD":{"code":"SGD","title":"Singapore Dollar","symbol":"SGD","native":"$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"135","max_amount":"1353897"},"THB":{"code":"THB","title":"Thai Baht","symbol":"\u0e3f","native":"\u0e3f","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"3156","max_amount":"31563499"},"TJS":{"code":"TJS","title":"Tajikistani Somoni","symbol":"TJS","native":"TJS","thousands_sep":"\u00a0","decimal_sep":";","symbol_left":false,"space_between":true,"exp":2,"min_amount":"938","max_amount":"9389950"},"TRY":{"code":"TRY","title":"Turkish Lira","symbol":"TRY","native":"TL","thousands_sep":".","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"526","max_amount":"5267200"},"TTD":{"code":"TTD","title":"Trinidad and Tobago Dollar","symbol":"TTD","native":"$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"675","max_amount":"6757850"},"TWD":{"code":"TWD","title":"New Taiwan Dollar","symbol":"NT$","native":"NT$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"3072","max_amount":"30722993"},"TZS":{"code":"TZS","title":"Tanzanian Shilling","symbol":"TZS","native":"TSh","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"230200","max_amount":"2302000188"},"UAH":{"code":"UAH","title":"Ukrainian Hryvnia","symbol":"UAH","native":"\u20b4","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":false,"space_between":false,"exp":2,"min_amount":"2764","max_amount":"27648991"},"UGX":{"code":"UGX","title":"Ugandan Shilling","symbol":"UGX","native":"USh","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":0,"min_amount":"3657","max_amount":"36575502"},"USD":{"code":"USD","title":"United States Dollar","symbol":"$","native":"$","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":false,"exp":2,"min_amount":"100","max_amount":1000000},"UYU":{"code":"UYU","title":"Uruguayan Peso","symbol":"UYU","native":"$","thousands_sep":".","decimal_sep":",","symbol_left":true,"space_between":true,"exp":2,"min_amount":"3246","max_amount":"32469503"},"UZS":{"code":"UZS","title":"Uzbekistani Som","symbol":"UZS","native":"UZS","thousands_sep":"\u00a0","decimal_sep":",","symbol_left":false,"space_between":true,"exp":2,"min_amount":"832759","max_amount":"8327599915"},"VND":{"code":"VND","title":"Vietnamese \u0110\u1ed3ng","symbol":"\u20ab","native":"\u20ab","thousands_sep":".","decimal_sep":",","symbol_left":false,"space_between":true,"exp":0,"min_amount":"23084","max_amount":"230840500"},"YER":{"code":"YER","title":"Yemeni Rial","symbol":"YER","native":"\u0631.\u064a.\u200f","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"25030","max_amount":"250301249"},"ZAR":{"code":"ZAR","title":"South African Rand","symbol":"ZAR","native":"R","thousands_sep":",","decimal_sep":".","symbol_left":true,"space_between":true,"exp":2,"min_amount":"1362","max_amount":"13620106"}}`

var SupportedCurrencies = make(map[string]Currency)

func init() {
	err := defaultJson.Unmarshal(unsafeConvert.ByteSlice(dataCurrencies), &SupportedCurrencies)
	if err != nil {
		panic(err)
	}
}
