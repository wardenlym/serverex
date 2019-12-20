package codegen

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

// 去掉c#类名中的"xxxResponse"后缀
func trim_response_name_suffix(name string) string {
	if strings.HasSuffix(name, "Response") {
		return strings.TrimSuffix(name, "Response")
	}
	return name
}

// 去掉去掉c#类字段名中的"odm.xxx"前缀
func trim_pkg_and_dot(s string) string {
	if strings.Contains(s, ".") {
		return strings.Trim(filepath.Ext(s), ".")
	}
	return s
}

// 去掉指针符号
func trim_star(s string) string {
	if strings.Contains(s, "*") {
		return strings.Trim(s, "*")
	}
	return s
}

// 小写第一个字母. C# 和 mongodb 中, 都使用首字母小写驼峰命名
func lowercase_first(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

// 根据go的结构体生成c#的类 只有简单类型. 初始化了List 和 Dictionary
func parse_go_struct_field_to_csharp_class_field(typename, fieldname string) string {
	with_initer := ""

	if strings.HasPrefix(typename, "float32") {
		typename = "float"
	} else if strings.HasPrefix(typename, "int64") {
		typename = "long"
	} else if strings.HasPrefix(typename, "[]") {
		typename = "List<" + trim_pkg_and_dot(trim_star(typename[2:])) + ">"
		with_initer = "new " + typename + "()"

	} else if strings.HasPrefix(typename, "map[") {
		r := regexp.MustCompile(`map\[(\w+)\]([\w.]*)`)
		if r.MatchString(typename) {
			ss := r.FindStringSubmatch(typename)
			typename = fmt.Sprintf("Dictionary<%s, %s>", trim_pkg_and_dot(trim_star(ss[1])), trim_pkg_and_dot(trim_star(ss[2])))
			with_initer = "new " + typename + "()"
		}
	} else {
		// nop
	}

	if with_initer != "" {
		return fmt.Sprintf("\tpublic %s %s = %s;", trim_star(trim_pkg_and_dot(typename)), lowercase_first(fieldname), with_initer)
	} else {
		return fmt.Sprintf("\tpublic %s %s;", trim_star(trim_pkg_and_dot(typename)), lowercase_first(fieldname))
	}
}

// 为了用继承添加一个类型标记
func parse_inherit(name string) string {
	if strings.HasSuffix(name, "Request") {
		return "request"
	} else {
		r := strings.TrimSuffix(name, "Response")
		return "response_of<" + r + "Request" + ">"
	}
}

// 为了生成一个用于取msgcode的字段
func parse_override1(name string) string {
	r := ""
	if strings.HasSuffix(name, "Request") {
		r = strings.TrimSuffix(name, "Request")

	} else {
		r = strings.TrimSuffix(name, "Response")
	}
	return "public override Msg.Code _msgcode() { return Msg.Code." + r + ";}"
}

// 主要是为了对齐c#的枚举类型
func format_enum_for_csharp(list EnumDef) EnumDef {
	maxlen := 0
	for _, v := range list.Fields {
		maxlen = func() int {
			if len(v.Name) > maxlen {
				return len(v.Name)
			}
			return maxlen
		}()
	}

	for i, v := range list.Fields {
		fmt_s := "%-" + fmt.Sprint(maxlen) + "s = %d,"
		if v.Comment != "" {
			fmt_s = fmt_s + " // " + v.Comment
		}
		if i+1 < len(list.Fields) && list.Fields[i+1].Code != v.Code+1 {
			fmt_s = fmt_s + "\n"
		}
		list.Fields[i].FormatedForCsharp = fmt.Sprintf(fmt_s, v.Name, v.Code)

	}
	return list
}
