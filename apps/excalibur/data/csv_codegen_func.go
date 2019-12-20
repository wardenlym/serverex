package data

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"strconv"
	"strings"
)

func genCSVtoMapDefine(typename string, lines []string, filter_use func(string, string) bool) string {
	fieldname := strings.Split(strings.Trim(lines[1], "\r\n"), ",")
	types := strings.Split(strings.Trim(lines[2], "\r\n"), ",")
	kvsrc := ""
	for i := 3; i < len(lines); i++ {
		line := strings.Trim(lines[i], "\r\n")
		if len(line) == 0 {
			continue
		}
		fields := strings.Split(line, ",")
		src1 := ""
		for j, fieldv := range fields {

			if !filter_use(typename, fieldname[j]) {
				continue
			}

			src1 = src1 + genStructFieldValue(types[j], fieldname[j], fieldv)
		}

		kvsrc = kvsrc + genIntKeyMapFieldDefine(fields[0], typename, src1)
	}

	mapsrc := genMapVarDefine("Table_"+typename, "int", typename, kvsrc)
	// fmt.Println(mapsrc)
	return mapsrc
}

func genStructFieldValue2(typename string, fieldname string, fieldvalue string) string {
	src := ""
	if strings.Contains(typename, "[]") {
		if typename == "int[]" || typename == "int[:]" {
			members := strings.Split(fieldvalue, "|")
			s := ""
			for _, member := range members {
				if strings.Contains(member, ":") {
					id := strings.Split(member, ":")[0]
					cnt, e := strconv.Atoi(strings.Split(member, ":")[1])
					if e != nil {
						panic(e)
					}
					for k := 0; k < cnt; k++ {
						s = s + id + ","
					}
				} else {
					s = s + member + ","
				}
			}
			src = fmt.Sprintf("%s:[]int{%s},\n", strings.Title(fieldname), strings.TrimSuffix(s, ","))
		}
	} else if strings.Contains(typename, "string") {
		src = fmt.Sprintf("%s:\"%s\",\n", strings.Title(fieldname), fieldvalue)
	} else {
		src = fmt.Sprintf("%s:%s,\n", strings.Title(fieldname), fieldvalue)
	}
	return src
}

func parse_intn(fieldname string, fieldvalue string) string {
	members := strings.Split(fieldvalue, "|")
	s := ""
	for _, member := range members {
		if member == "" {
			continue
		}
		if strings.Contains(member, ":") {
			id := strings.Split(member, ":")[0]
			if _, e := strconv.Atoi(id); e != nil {
				panic(fieldname + " " + fieldvalue + " " + e.Error())
			}
			cnt, e := strconv.Atoi(strings.Split(member, ":")[1])
			if e != nil {
				panic(fieldname + " " + fieldvalue + " " + e.Error())
			}
			for k := 0; k < cnt; k++ {
				s = s + id + ","
			}
		} else {
			if _, e := strconv.Atoi(member); e != nil {
				panic(fieldname + " " + fieldvalue + " " + e.Error())
			}
			s = s + member + ","
		}
	}
	src := fmt.Sprintf("%s:[]int{%s},\n", strings.Title(fieldname), strings.TrimSuffix(s, ","))
	return src
}

func parse_stringn(fieldname string, fieldvalue string) string {
	members := strings.Split(fieldvalue, "|")
	s := ""
	for _, member := range members {
		if _, e := strconv.ParseFloat(member, 32); e != nil {
			return fmt.Sprintf("%s:[]string{\"%s\"},\n", strings.Title(fieldname), fieldvalue)
		}
		s = s + member + ","
	}
	src := fmt.Sprintf("%s:[]string{\"%s\"},\n", strings.Title(fieldname), strings.TrimSuffix(s, ","))
	return src
}

func parse_floatn(fieldname string, fieldvalue string) string {
	members := strings.Split(fieldvalue, "|")
	s := ""
	for _, member := range members {
		s = s + "\"" + member + "\"" + ","
	}
	src := fmt.Sprintf("%s:[]string{%s},\n", strings.Title(fieldname), strings.TrimSuffix(s, ","))
	return src
}

func genStructFieldValue(typename string, fieldname string, fieldvalue string) string {
	src := ""
	switch typename {
	case "int":
		var i int = 0
		var e error = nil
		if fieldvalue != "" {
			i, e = strconv.Atoi(fieldvalue)
		}

		if e != nil {

			panic(typename + " " + fieldname + " " + fieldvalue + " " + e.Error())
		}
		src = fmt.Sprintf("%s:%d,\n", strings.Title(fieldname), i)
	case "string":
		src = fmt.Sprintf("%s:\"%s\",\n", strings.Title(fieldname), fieldvalue)
	case "float":
		_, e := strconv.ParseFloat(fieldvalue, 32)
		if e != nil && fieldvalue != "" {
			panic(typename + " " + fieldname + " " + fieldvalue + " " + e.Error())
		}
		src = fmt.Sprintf("%s:%s,\n", strings.Title(fieldname), fieldvalue)
	case "float[]":
		panic(typename + " " + fieldname + " " + fieldvalue + " ")
	case "string[]":
		src = parse_stringn(fieldname, fieldvalue)
	case "int[]", "int[:]":
		src = parse_intn(fieldname, fieldvalue)
	default:
		panic("unknow type :" + typename)
	}
	return src
}

func writeGolangFile(filename string, content string) error {
	dst, err := format.Source([]byte(content))
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, dst, 0644)
	if err != nil {
		return err
	}
	return nil
}

func genGolangFile(packagename string, content string) string {
	return fmt.Sprintf("package %s\n\n%s\n", packagename, content)
}

func genMapVarDefine(name string, typekey string, typevalue string, fieldsrc string) string {
	src := fmt.Sprintf("var %s = map[%s]%s{}\n", name, typekey, typevalue) +
		fmt.Sprintf("func Load_%s() {\n", name) +
		fmt.Sprintf("%s = map[%s]%s {\n%s\n}", name, typekey, typevalue, fieldsrc) +
		"}\n"
	return src
}

func genIntKeyMapFieldDefine(k string, valuetype string, src string) string {
	return fmt.Sprintf("%s : %s {\n%s\n},\n", k, valuetype, src)
}
