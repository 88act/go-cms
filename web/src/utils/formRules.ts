//必填项
const required = {
	required: true,
	message: '请填写此项内容',
	trigger: 'blur',
}

function requiredMsg(msg) {
	return {
		required: true,
		message: msg,
		trigger: 'blur'
	}
}


const username2 = {
	pattern: /^(?![0-9]*$)(?![a-zA-Z]*$)[a-zA-Z0-9]{6,24}$/,
	message: '用户名必须为6-24位字母和数字组合',
	trigger: 'blur',
}

//用户名 以字母开头、可带数字_.的字串
const username = {
	pattern: /^^[a-zA-Z]{1}([a-zA-Z0-9]){5,23}$/,
	message: '用户名必须以字母开头,6-24位字母或数字',
	trigger: 'blur',
}


//校验登录名：只能输入5-20个以字母开头、可带数字、“_”、“.”的字串
const username3 = {
	pattern: /^[a-zA-Z]{1}([a-zA-Z0-9]|[._]){4,19}$/,
	message: '用户名必须为6-24位字母和数字组合',
	trigger: 'blur',

}

//弱：纯数字，纯字母，纯特殊字符
const password = {
	pattern: /^(?:\d+|[a-zA-Z]+|[!@#$%^&*]+){6,24}$/,
	message: '密码6-24位,字母/数字/符号(除空格)',
	trigger: 'blur',
}
//密码校验 （6-20位英文字母、数字或者符号(除空格)，且字母、数字和标点符号至少包含两种）
const password2 = {
	pattern: /^(?![\d]+$)(?![a-zA-Z]+$)(?![^\da-zA-Z]+$)([^\u4e00-\u9fa5\s]){6,24}$/,
	message: '密码6-24位,字母/数字/符号(除空格),且至少包含两种',
	trigger: 'blur',
}
//数字范围校验
function range(min = 0, max = 99999999.99, msg) {
	const validator = (rule, val, cb) => {
		val = +val
		if (!(val >= min && val <= max)) {
			cb(new Error(rule.message))
		}
		cb()
	}

	if (!msg || msg == "") {
		return {
			validator,
			message: `数字范围在 ${min} 到 ${max}`,
			trigger: 'blur'
		}
	} else {
		return {
			validator,
			message: msg,
			trigger: 'blur'
		}
	}
}

//字符长度校验
function len(min = 0, max = 255) {
	return {
		min,
		max,
		message: `长度在 ${min} 到 ${max} 个字符`,
		trigger: 'blur'
	}
}

//手机号校验
const mobile= {
  pattern: /^1[1-9][0-9]{9}$/,
  message: '请输入正确的手机号',
  trigger: 'blur'
}
//电话
const phone = {
		pattern: /^1[1-9][0-9]{9}$/,
		message: '请输入正确的电话号',
		trigger: 'blur',
}


//数字范围校验 1~ 100
function rangeInt(max) {
	return {
		pattern: /^([1-9][0-9]{0,1}|max)$/,
		message: '请输入 1~${min}范围整数 ',
		trigger: 'blur',
	}
}

//只含有汉字字母数字
function okChat()  {
	return {
		pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]+$/,
		message: '请输入汉字字母数字',
		trigger: 'blur',
	}
}




//邮箱校验
const email = {
	type: 'email',
	message: '请输入正确的邮箱地址',
	trigger: 'blur'
}
//纯整数校验
const integer = {
	pattern: /^(\-|\+)?\d+?$/,
	message: '请输入正确的数字',
	trigger: 'blur'
}
//纯整数或小数校验
const float = {
	pattern: /^(\-|\+)?\d+(\.\d+)?$/,
	message: '请输入正确的数字（整数或小数）',
	trigger: 'blur'
}
//日期校验
const date = {
	type: 'date',
	message: '请选择日期',
	trigger: 'blur'
}
//url网址
const url = {
	type: 'url',
	message: '请输入正确的url网址',
	trigger: 'blur'
}
//身份证号码
const idcard = {
	pattern: /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/,
	message: '请输入正确的邮箱地址',
	trigger: 'blur'
}

//微信号校验
const wecheat = {
	pattern: /^[a-zA-Z][a-zA-Z0-9_-]{5,19}$/,
	message: '请输入正确的微信号',
	trigger: 'blur'
};
//营业执照校验
const businessLicense = {
	pattern: /(^(?:(?![IOZSV])[\dA-Z]){2}\d{6}(?:(?![IOZSV])[\dA-Z]){10}$)|(^\d{15}$)/,
	message: '请输入正确的营业执照',
	trigger: 'blur'
};

//银行卡号校验
function bankAccount() {
	const validator = (rule, value, cb) => {
		const strBin =
			'10,18,30,35,37,40,41,42,43,44,45,46,47,48,49,50,51,52,53,54,55,56,58,60,62,65,68,69,84,87,88,94,95,98,99'
		if (!value) {
			cb(new Error('收款账户不能为空'))
		} else if (!Number.isInteger(+value)) {
			cb(new Error('银行卡号必须全为数字'))
		} else if (value.trim().length < 12 || value.trim().length > 19) {
			cb(new Error('银行卡号长度必须在12到19之间'))
		} else if (strBin.indexOf(value.substring(0, 2)) === -1) {
			cb(new Error('银行卡号开头6位不符合规范'))
		} else {
			cb();
		}
	};
	return {
		validator,
		trigger: 'blur'
	}
}

//中文校验
const chinese = {
	pattern: /^[\u0391-\uFFE5A-Za-z]+$/,
	message: '请输入中文',
	trigger: 'blur'
}



export {
	required,
	requiredMsg,
	username,
	password,
	len,
	integer,
	range,
	rangeInt,
	float,
	phone,
  mobile,
	email,
	date,
	url,
	idcard,
	wecheat,
	businessLicense,
	bankAccount,
	chinese,
	okChat
}



// 1、是否合法IP地址:
// pattern:/^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/
// 2、是否手机号码或者固话
// pattern:/^((0\d{2,3}-\d{7,8})|(1[34578]\d{9}))$/
// 3、 是否身份证号码
// pattern:/(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/
// 4、是否邮箱
// pattern:/^([a-zA-Z0-9]+[-_\.]?)+@[a-zA-Z0-9]+\.[a-z]+$/
// 5、整数填写
// pattern:/^-?[1-9]\d*$/
// 6、正整数填写
// pattern:/^[1-9]\d*$/
// 7、小写字母
// pattern:/^[a-z]+$/
// 8、大写字母
// pattern:/^[A-Z]+$/
// 9、大小写混合
// pattern:/^[A-Za-z]+$/
// 10、数字加英文，不包含特殊字符
// pattern:/^[a-zA-Z0-9]+$/
// 11、前两位是数字后一位是英文
// pattern:/^\d{2}[a-zA-Z]+$/
// 12、密码校验（6-20位英文字母、数字或者符号（除空格），且字母、数字和标点符号至少包含两种）
// pattern:/^(?![\d]+$)(?![a-zA-Z]+$)(?![^\da-zA-Z]+$)([^\u4e00-\u9fa5\s]){6,20}$/
// 13、中文校验
// pattern:/^[\u0391-\uFFE5A-Za-z]+$/
// 14、1-100的数字
// pattern: /^([1-9][0-9]{0,1}|100)$/
