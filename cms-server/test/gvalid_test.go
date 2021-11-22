package utils

func gvalid_test() {
	//vRequired()
	//vRequiredIf()
	//vRequiredUnless()
	//vRequiredWith()
	//vRequiredWithAll()
	//vRequiredWithout()
	//vRequiredWithoutAll()
	//vDate()
	//vDateFormat()
	//vEmail()
	//vPhone()
	//vTelephone()
	//vPassport()
	//vPassword()
	//vPassword2()
	//vPassword3()
	//vPostcode()
	//vIdNumber()
	//vLuhn()
	//vQQ()
	//vIp()
	//vIpv4()
	//vIpv6()
	//vMac()
	//vUrl()
	//vDomain()
	//vLength()
	//vMinLength()
	//vMaxLength()
	//vBetween()
	//vMin()
	//vMax()
	//vJSON()
	//vInteger()
	//vFloat()
	//vBoolean()
	//vSame()
	//vDifferent()
	//vIn()
	//vNotIn()
	//vRegex()
	//vMap()
	//vMsg()
}

/*
// 必需参数
// required
// 格式：required
func vRequired() {
	if err := gvalid.Check("", "required", nil); err != nil {
		fmt.Println(err.Strings())
		fmt.Println(err.Maps())
	}
}

// 给定字段值与所给值相等时
// required-if
// 格式: required-if:field,value,...
// 说明：必需参数(当任意所给定字段值与所给值相等时，即：当field字段的值为value时，当前验证字段为必须参数)
func vRequiredIf() {
	type User struct {
		A string
		B string `valid:"b@required-if:A,1"`
	}

	var user = User{"1", ""}
	if err := gvalid.CheckStruct(&user, nil); err != nil {
		fmt.Println(err.Strings())
		fmt.Println(err.Maps())
	}
}

// 给定字段值与所给值都不相等时
// required-unless
// 格式: required-unless:field,value,...
// 说明：必需参数(当所给定字段值与所给值都不相等时，即：当field字段的值不为value时，当前验证字段为必须参数)
func vRequiredUnless() {
	type User struct {
		A string
		B string `valid:"b@required-unless:A,1"`
	}

	var user = User{"2", ""}
	if err := gvalid.CheckStruct(&user, nil); err != nil {
		fmt.Println(err.Strings())
		fmt.Println(err.Maps())
	}
}

// 给定任意字段值不为空时
// required-with
// 格式: required-with:field1,field2,...
// 说明：必需参数(当所给定任意字段值不为空时)
func vRequiredWith() {
	type User struct {
		A string
		B string `valid:"b@required-with:A,C"`
		C string
	}

	var user = User{"a", "", "c"}
	if err := gvalid.CheckStruct(&user, nil); err != nil {
		fmt.Println(err.Strings())
		fmt.Println(err.Maps())
	}
}

// 给定所有字段值都不为空时
// required-with-all
// 格式: required-with-all:field1,field2,...
// 说明：必须参数(当所给定所有字段值都不为空时)
func vRequiredWithAll() {
	type User struct {
		A string
		B string `valid:"b@required-with-all:A,C"`
		C string
	}

	var user = User{"a", "", "c"}
	if err := gvalid.CheckStruct(&user, nil); err != nil {
		fmt.Println(err.Strings())
		fmt.Println(err.Maps())
	}
}

// 给定任意字段值为空时
// required-without
// 格式: required-without:field1,field2,...
// 说明：必需参数(当所给定任意字段值为空时)
func vRequiredWithout() {
	type User struct {
		A string
		B string `valid:"b@required-without:A,C"`
		C string
	}

	var user = User{"", "", "c"}
	if err := gvalid.CheckStruct(&user, nil); err != nil {
		fmt.Println(err.Strings())
		fmt.Println(err.Maps())
	}
}

// 给定所有字段值都为空时
// required-without-all
// 格式: required-without-all:field1,field2,...
// 说明：必须参数(当所给定所有字段值都为空时)
func vRequiredWithoutAll() {
	type User struct {
		A string
		B string `valid:"b@required-without-all:A,C"`
		C string
	}

	var user = User{"", "", ""}
	if err := gvalid.CheckStruct(&user, nil); err != nil {
		fmt.Println(err.Strings())
		fmt.Println(err.Maps())
	}
}

// 日期类型
// date
// 格式: date
// 说明：参数为常用日期类型，日期之间支持的连接符号-或/或.，也支持不带连接符号的8位长度日期，
// 格式如： 2006-01-02, 2006/01/02, 2006.01.02, 20060102, 01-Nov-2018, 01.Nov.018, 01/Nov/2018
func vDate() {
	rule := `date`
	if err := gvalid.Check("2020-01-07", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 自定义时间日期验证
// 格式：date-format:format
func vDateFormat() {
	rule := `date-format:2006-01-02 15:04:05`
	if err := gvalid.Check("2020-01-07:14:02:00", rule, nil); err != nil {
		fmt.Println(err.Strings())
		fmt.Println(err.Maps())
	}
}

// EMAIL邮箱地址
// email
// 格式: email
// 说明：EMAIL邮箱地址
func vEmail() {
	rule := `email`
	if err := gvalid.Check("xx@xx.xx", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 手机号
// phone
// 格式: phone
// 说明：手机号
func vPhone() {
	rule := `phone`
	if err := gvalid.Check("13888888888", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 国内座机电话号码
// telephone
// 格式: telephone
// 说明：国内座机电话号码，"XXXX-XXXXXXX"、"XXXX-XXXXXXXX"、"XXX-XXXXXXX"、"XXX-XXXXXXXX"、"XXXXXXX"、"XXXXXXXX"
func vTelephone() {
	rule := `telephone`
	if err := gvalid.Check("0111-12345678", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 通用帐号规则
// passport
// 格式: passport
// 说明：通用帐号规则(字母开头，只能包含字母、数字和下划线，长度在6~18之间)
func vPassport() {
	rule := `passport`
	if err := gvalid.Check("admin", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 通用密码
// password
// 格式: password
// 说明：通用密码(任意可见字符，长度在6~18之间)
func vPassword() {
	rule := `password`
	if err := gvalid.Check("123", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 中等强度密码
// password2
// 格式: password2
// 说明：中等强度密码(在弱密码的基础上，必须包含大小写字母和数字)
func vPassword2() {
	rule := `password2`
	if err := gvalid.Check("123Abc", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 强等强度密码
// password3
// 格式: password3
// 说明：强等强度密码(在弱密码的基础上，必须包含大小写字母、数字和特殊字符)
func vPassword3() {
	rule := `password3`
	if err := gvalid.Check("123Abc@", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 中国邮政编码
// postcode
// 格式: postcode
// 说明：中国邮政编码
func vPostcode() {
	rule := `postcode`
	if err := gvalid.Check("1234124", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 公民身份证号码
// id-number
// 格式: id-number
// 说明：公民身份证号码
func vIdNumber() {
	rule := `id-number`
	if err := gvalid.Check("340622199601016631", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 银行卡号校验
// luhn
// 格式: luhn
// 说明：银行卡号校验
func vLuhn() {
	rule := `luhn`
	if err := gvalid.Check("621226202001015588", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 腾讯QQ号码
// qq
// 格式: qq
// 说明：腾讯QQ号码
func vQQ() {
	rule := `qq`
	if err := gvalid.Check("1227716207", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// IPv4/IPv6地址
// ip
// 格式: ip
// 说明：IPv4/IPv6地址
func vIp() {
	rule := `ip`
	if err := gvalid.Check("192.168.1.1", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// IPv4地址
// ipv4
// 格式: ipv4
// 说明：IPv4地址
func vIpv4() {
	rule := `ipv4`
	if err := gvalid.Check("192.168.1.1", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// IPv6地址
// ipv6
// 格式: ipv6
// 说明：IPv6地址
func vIpv6() {
	rule := `ipv6`
	if err := gvalid.Check("ae80::ebc8:9f54:20ac:b32b", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// MAC地址
// mac
// 格式: mac
// 说明：MAC地址
func vMac() {
	rule := `mac`
	if err := gvalid.Check("a1:00:ec:21:32:22", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// url
// 格式: url
// 说明：URL
func vUrl() {
	rule := `url`
	if err := gvalid.Check("http://127.0.0.1:8080/get_user", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 域名
// domain
// 格式: domain
// 说明：域名
func vDomain() {
	rule := `domain`
	if err := gvalid.Check("a.net", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 参数长度为min到max
// length
// 格式: length:min,max
// 说明：参数长度为min到max(长度参数为整形)，注意中文一个汉字占3字节
func vLength() {
	rule := `length:6,10`
	if err := gvalid.Check("hello", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 参数长度最小为min
// min-length
// 格式: min-length:min
// 说明：参数长度最小为min(长度参数为整形)，注意中文一个汉字占3字节
func vMinLength() {
	rule := `min-length:6`
	if err := gvalid.Check("hello2", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 参数长度最大为max
// max-length
// 格式: max-length:max
// 说明：参数长度最大为max(长度参数为整形)，注意中文一个汉字占3字节
func vMaxLength() {
	rule := `max-length:6`
	if err := gvalid.Check("hello22", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 参数大小为min到max
// between
// 格式: between:min,max
// 说明：参数大小为min到max(支持整形和浮点类型参数)
func vBetween() {
	rule := `between:1,5`
	if err := gvalid.Check("4", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 参数大小最小为min
// min
// 格式: min:min
// 说明：参数大小最小为min(支持整形和浮点类型参数)
func vMin() {
	rule := `min:3`
	if err := gvalid.Check("4", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 参数大小最大为max
// max
// 格式: max:max
// 说明：参数大小最大为max(支持整形和浮点类型参数)
func vMax() {
	rule := `max:3`
	if err := gvalid.Check("2", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// JSON格式校验
// json
// 格式: json
// 说明：判断数据格式为JSON
func vJSON() {
	rule := `json`
	if err := gvalid.Check(`{"name": "admin"}`, rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// integer
// 格式: integer
// 说明：整数
func vInteger() {
	rule := `integer`
	if err := gvalid.Check("a12341234123412", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// float
// 格式: float
// 说明：浮点数
func vFloat() {
	rule := `float`
	if err := gvalid.Check("2323.11", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// boolean
// 格式: boolean
// 说明：布尔值(1,true,on,yes为true | 0,false,off,no,""为false)
func vBoolean() {
	rule := `boolean`
	if err := gvalid.Check("false", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 参数值必需与field参数的值相同
// same
// 格式: same:field
// 说明：参数值必需与field参数的值相同
func vSame() {
	type User struct {
		Name      string
		Password1 string
		Password2 string `valid:"password@same:Password1"`
	}

	var user = User{"admin", "111111", "222222"}

	if err := gvalid.CheckStruct(user, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 参数值不能与field参数的值相同
// different
// 格式: different:field
// 说明：参数值不能与field参数的值相同
func vDifferent() {
	type User struct {
		Name      string
		Password1 string
		Password2 string `valid:"password@different:Password1"`
	}

	var user = User{"admin", "111111", "111111"}

	if err := gvalid.CheckStruct(user, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// in
// 格式: in:value1,value2,...
// 说明：参数值应该在value1,value2,...中(字符串匹配)
func vIn() {
	rule := `in:admin,superadmin`
	if err := gvalid.Check("superadmin", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// not-in
// 格式: not-in:value1,value2,...
// 说明：参数值不应该在value1,value2,...中(字符串匹配)
func vNotIn() {
	rule := `not-in:admin,superadmin`
	if err := gvalid.Check("admin", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 自定义正则匹配规则
// regex
// 格式: regex:pattern
// 说明：参数值应当满足正则匹配规则pattern
func vRegex() {
	rule := `regex:^\d{3}`
	if err := gvalid.Check("133", rule, nil); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// map验证 自定义错误提示
func vMap() {
	data := map[string]string{
		"username": "root",
		"password": "123Abc@",
	}

	rules := map[string]string{
		"username": "required",
		"password": "password3",
	}

	if err := gvalid.CheckMap(data, rules); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}

// 自定义错误提示
func vMsg() {
	rule := `required|length:6,10`

	msgs := map[string]string{
		"required": "老铁，这个不能空啊！",
		"length":   "兄得，范围在6-10哥字符",
	}

	if err := gvalid.Check("abc", rule, msgs); err != nil {
		fmt.Println(err.Maps())
		fmt.Println(err.Strings())
	}
}
*/
