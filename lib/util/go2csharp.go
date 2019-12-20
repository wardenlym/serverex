package util

import (
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
)

func MakeCSharpNamespace(using string, name string, funcs ...func() string) string {
	src := using + "\n\n"
	emit := func(s string) {
		src = src + s + "\n"
	}

	emit("namespace " + name + " \n{")

	for _, f := range funcs {
		emit(f())
	}

	emit("}")

	return src
}

func removeDotInName(s string) string {
	if strings.Contains(s, ".") {
		return strings.Trim(filepath.Ext(s), ".")
	}
	return s
}

func ParseStructToClass1(s interface{}) func() string {
	return func() string {
		src := ""
		emit := func(s string) {
			src = src + "\t" + s + "\n"
		}

		typ := reflect.TypeOf(s)

		emit("public class " + typ.Name() + "\n\t{")

		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)

			typename := ""
			init := ""
			if strings.HasPrefix(field.Type.String(), "float32") {
				typename = "float"
			} else if strings.HasPrefix(field.Type.String(), "int64") {
				typename = "long"
			} else if strings.HasPrefix(field.Type.String(), "[]") {
				typename = "List<" + removeDotInName(field.Type.String()[2:]) + ">"
				init = "new " + typename + "()"

			} else if strings.HasPrefix(field.Type.String(), "map[") {
				r := regexp.MustCompile(`map\[(\w+)\]([\w.]*)`)
				if r.MatchString(field.Type.String()) {
					ss := r.FindStringSubmatch(field.Type.String())
					typename = fmt.Sprintf("Dictionary<%s, %s>", removeDotInName(ss[1]), removeDotInName(ss[2]))
					init = "new " + typename + "()"
				}
			} else {
				typename = removeDotInName(field.Type.String())
			}

			name := field.Tag.Get("json")
			if init != "" {
				emit(fmt.Sprintf("\tpublic %s %s = %s;", typename, name, init))
			} else {
				emit(fmt.Sprintf("\tpublic %s %s;", typename, name))
			}
		}

		emit("}")
		return src
	}
}

func ParseStructToClass(s interface{}) func() string {
	return func() string {
		src := ""
		emit := func(s string) {
			src = src + "\t" + s + "\n"
		}

		typ := reflect.TypeOf(s)

		name := typ.Name()

		if strings.HasSuffix(name, "Request") {
			emit("public class " + typ.Name() + " : request\n\t{")
			r := strings.TrimSuffix(name, "Request")
			emit("\tpublic override Msg.Code _msgcode() { return Msg.Code." + r + "; }\n")
		} else {
			r := strings.TrimSuffix(name, "Response")
			emit("public class " + typ.Name() + " : response_of<" + r + "Request" + ">\n\t{")
			emit("\tpublic override Msg.Code _msgcode() { return Msg.Code." + r + "; }\n")
		}

		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)

			typename := ""
			init := ""
			if strings.HasPrefix(field.Type.String(), "float32") {
				typename = "float"
			} else if strings.HasPrefix(field.Type.String(), "int64") {
				typename = "long"
			} else if strings.HasPrefix(field.Type.String(), "[]") {
				typename = "List<" + removeDotInName(field.Type.String()[2:]) + ">"
				init = "new " + typename + "()"

			} else if strings.HasPrefix(field.Type.String(), "map[") {
				r := regexp.MustCompile(`map\[(\w+)\]([\w.]*)`)
				if r.MatchString(field.Type.String()) {
					ss := r.FindStringSubmatch(field.Type.String())
					typename = fmt.Sprintf("Dictionary<%s, %s>", removeDotInName(ss[1]), removeDotInName(ss[2]))
					init = "new " + typename + "()"
				}
			} else {
				typename = removeDotInName(field.Type.String())
			}

			name := field.Tag.Get("json")
			if init != "" {
				emit(fmt.Sprintf("\tpublic %s %s = %s;", typename, name, init))
			} else {
				emit(fmt.Sprintf("\tpublic %s %s;", typename, name))
			}
		}

		emit("}")
		return src
	}
}
