package protocol

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"unicode"
)

type Defs struct {
	Msg EnumDef
	Err EnumDef
	Odm []StructDef
	Rpc []StructDef
}

type EnumCodeDef struct {
	Code    int
	Name    string
	Comment string
}

type EnumDef struct {
	Namespace string
	Name      string
	Fields    []EnumCodeDef
}

type StructFieldDef struct {
	Name       string
	Type       string
	Tag        string
	TypeCsharp string
	Comment    string
}

type StructDef struct {
	Namespace string
	Name      string
	Fields    []StructFieldDef
}

func parse(src []byte) Defs {
	interp := &interp{
		ret: Defs{
			Msg: EnumDef{Fields: []EnumCodeDef{}},
			Err: EnumDef{Fields: []EnumCodeDef{}},
			Odm: []StructDef{},
			Rpc: []StructDef{},
		},
	}
	for i, v := range strings.Split(string(src), "\n") {
		interp.emit(i+1, v)
	}
	return interp.complie()
}

type interp struct {
	index       int
	namespace   string
	enum        string
	ret         Defs
	struct_name string

	in_enum   bool
	in_struct bool

	src_odm_golang string
	src_odm_csharp string
	src_msg_golang string
	src_msg_csharp string
	src_rpc_golang string
	src_rpc_csharp string
}

func (p *interp) emit(no int, line string) {

	// var empty_reg = regexp.MustCompile(`^(\S+).*$`)
	// var no_reg = regexp.MustCompile(`^}$`)
	var ns_reg = regexp.MustCompile(`^ns (\w+)$`)
	var enum_reg = regexp.MustCompile(`^enum (\w+) from (\w+) {$`)
	var enum_field_reg = regexp.MustCompile(`^\s+(\w+)$`)
	var enum_field_comment_reg = regexp.MustCompile(`^\s+(\w+)\s+(.*)$`)

	var struct_reg = regexp.MustCompile(`^type (\w+) struct {$`)
	var struct_field_reg = regexp.MustCompile(`^\s+(\w+)\s+(\S+)$`)
	var struct_field_comment_reg = regexp.MustCompile(`^\s+(\w+)\s+(\S+)\s+(.*)$`)

	if ns_reg.Match([]byte(line)) {
		ss := ns_reg.FindStringSubmatch(line)
		println(no, "ns_reg", ss[1])

		p.namespace = ss[1]

	} else if enum_reg.Match([]byte(line)) {
		ss := enum_reg.FindStringSubmatch(line)
		println(no, "enum_reg", ss[1], ss[2])
		from, _ := strconv.Atoi(ss[2])
		p.index = from
		p.in_enum = true
		if ss[1] == "Msg" {
			p.enum = "Msg"
			p.ret.Msg.Namespace = p.namespace
			p.ret.Msg.Name = "Msg"
			fmt.Printf("######## 1 %v\n", p.ret.Msg)
		} else if ss[1] == "Err" {
			p.enum = "Err"
			p.ret.Err.Namespace = p.namespace
			p.ret.Err.Name = "Err"
			fmt.Printf("######## 2 %v\n", p.ret.Msg)
		}

	} else if p.in_enum && enum_field_reg.Match([]byte(line)) {
		ss := enum_field_reg.FindStringSubmatch(line)
		println(no, "enum_field_reg", ss[1])

		if p.enum == "Msg" {
			p.ret.Msg.Fields = append(p.ret.Msg.Fields,
				EnumCodeDef{
					Code: p.index,
					Name: ss[1],
				})

			p.index++
		} else if p.enum == "Err" {
			p.ret.Err.Fields = append(p.ret.Err.Fields,
				EnumCodeDef{
					Code: p.index,
					Name: ss[1],
				})
			p.index++
		}
	} else if p.in_enum && enum_field_comment_reg.Match([]byte(line)) {
		ss := enum_field_comment_reg.FindStringSubmatch(line)
		println(no, "enum_field_comment_reg", ss[1], ss[2])

		if p.enum == "Msg" {
			p.ret.Msg.Fields = append(p.ret.Msg.Fields,
				EnumCodeDef{
					Code:    p.index,
					Name:    ss[1],
					Comment: ss[2],
				})

			p.index++
		} else if p.enum == "Err" {
			p.ret.Err.Fields = append(p.ret.Err.Fields,
				EnumCodeDef{
					Code:    p.index,
					Name:    ss[1],
					Comment: ss[2],
				})
			p.index++
		}
	} else if struct_reg.Match([]byte(line)) {
		ss := struct_reg.FindStringSubmatch(line)
		println(no, "struct_reg", ss[1])
		p.in_struct = true
		p.in_enum = false
		if p.namespace == "rpc" {
			p.ret.Rpc = append(p.ret.Rpc, StructDef{
				Namespace: p.namespace,
				Name:      ss[1],
				Fields:    []StructFieldDef{},
			})
		} else if p.namespace == "odm" {
			p.ret.Odm = append(p.ret.Odm, StructDef{
				Namespace: p.namespace,
				Name:      ss[1],
				Fields:    []StructFieldDef{},
			})
		}

	} else if p.in_struct && struct_field_reg.Match([]byte(line)) {
		ss := struct_field_reg.FindStringSubmatch(line)
		println(no, "struct_field_reg", ss[1], ss[2])

		if p.namespace == "rpc" {
			p.ret.Rpc[len(p.ret.Rpc)-1].Fields = append(p.ret.Rpc[len(p.ret.Rpc)-1].Fields,
				StructFieldDef{
					Name: ss[1],
					Type: ss[2],
				},
			)
		} else if p.namespace == "odm" {
			p.ret.Odm[len(p.ret.Odm)-1].Fields = append(p.ret.Odm[len(p.ret.Odm)-1].Fields,
				StructFieldDef{
					Name: ss[1],
					Type: ss[2],
				},
			)
		}

	} else if p.in_struct && struct_field_comment_reg.Match([]byte(line)) {
		ss := struct_field_comment_reg.FindStringSubmatch(line)
		println(no, "struct_field_comment_reg", ss[1], ss[2], ss[3])
		fmt.Println("---------------", ss[3])
		if p.namespace == "rpc" {
			p.ret.Rpc[len(p.ret.Rpc)-1].Fields = append(p.ret.Rpc[len(p.ret.Rpc)-1].Fields,
				StructFieldDef{
					Name:    ss[1],
					Type:    ss[2],
					Comment: ss[3],
				},
			)
		} else if p.namespace == "odm" {
			p.ret.Odm[len(p.ret.Odm)-1].Fields = append(p.ret.Odm[len(p.ret.Odm)-1].Fields,
				StructFieldDef{
					Name:    ss[1],
					Type:    ss[2],
					Comment: ss[3],
				},
			)
		}

	} else {
		//println(no, "miss:", line)
	}

}

func (p *interp) complie() Defs {

	{
		// msg golang
		tpl_txt := `/// Code generated by protocol.define; DO NOT EDIT.

package msg
	
type Code = int

const (
{{range .Msg.Fields}}    C_{{.Name}} = {{.Code}} {{.Comment}}
{{end}})

var Code_To_String = map[Code]string{
{{range .Msg.Fields}}    C_{{.Name}} : "{{.Name}}",
{{end}}}

const (
{{range .Err.Fields}}    {{.Name}} = {{.Code}}
{{end}})


`
		tpl, e := template.New("msg").Parse(tpl_txt)
		if e != nil {
			panic(e)
		}

		source := bytes.NewBuffer([]byte{})
		e = tpl.Execute(source, p.ret)
		fmt.Println(string(source.Bytes()))
		dst, err := format.Source([]byte(source.Bytes()))
		if err != nil {
			panic(err)
		}

		p.src_msg_golang = string(dst)

		ioutil.WriteFile("./msg/msg.go", dst, 0664)
	}

	{
		// msg csharp
		tpl_txt := `/// Code generated by protocol.define; DO NOT EDIT.
namespace ExcaliburNetwork.Protocol
{
	public partial class Msg
	{
		public enum Code : int {
{{range .Msg.Fields}}            {{.Name}} = {{.Code}},
{{end}}        }

		public enum Err : int {
{{range .Err.Fields}}            {{.Name}} = {{.Code}},
{{end}}        }
	}
}
`
		tpl, e := template.New("msg_csharp").Parse(tpl_txt)
		if e != nil {
			panic(e)
		}

		source := bytes.NewBuffer([]byte{})
		e = tpl.Execute(source, p.ret)
		fmt.Println(string(source.Bytes()))
		p.src_msg_csharp = string(source.Bytes())

		ioutil.WriteFile("./msg/Msg.cs", source.Bytes(), 0644)
	}

	{
		// odm golang
		for i, v := range p.ret.Odm {
			for j, w := range v.Fields {
				name := lowercaseFirst(w.Name)

				tag := "`" + `json:"` + name + `" bson:"` + name + `"` + "`"
				p.ret.Odm[i].Fields[j].Tag = tag
			}
		}

		tpl_txt := `/// Code generated by protocol.define; DO NOT EDIT.

package odm

{{range .}}
type {{.Name}} struct {
{{range .Fields}}    {{.Name}} {{.Type}} {{.Tag}} {{.Comment}}
{{end}}}
{{end}}
`
		tpl, e := template.New("odm").Parse(tpl_txt)
		if e != nil {
			panic(e)
		}

		source := bytes.NewBuffer([]byte{})
		e = tpl.Execute(source, p.ret.Odm)

		// fmt.Printf("%v\n", string(source.Bytes()))
		dst, err := format.Source([]byte(source.Bytes()))
		if err != nil {
			panic(err)
		}

		// fmt.Println(string(dst))
		p.src_odm_golang = string(dst)

		ioutil.WriteFile("./odm/odm.go", dst, 0644)
	}

	{
		// odm csharp
	}

	{
		// rpc golang
		for i, v := range p.ret.Rpc {
			for j, w := range v.Fields {
				name := lowercaseFirst(w.Name)

				tag := "`" + `json:"` + name + `"` + "`"
				p.ret.Rpc[i].Fields[j].Tag = tag
			}
		}

		tpl_txt := `/// Code generated by protocol.define; DO NOT EDIT.

package msg

import "gitlab.com/megatech/serverex/apps/excalibur/protocol/odm"

{{range .}}
type {{.Name}} struct {
{{range .Fields}}    {{.Name}} {{.Type}} {{.Tag}} {{.Comment}}
{{end}}}
{{end}}
`
		tpl, e := template.New("rpc").Parse(tpl_txt)
		if e != nil {
			panic(e)
		}

		source := bytes.NewBuffer([]byte{})
		e = tpl.Execute(source, p.ret.Rpc)

		// fmt.Printf("%v\n", string(source.Bytes()))
		dst, err := format.Source([]byte(source.Bytes()))
		if err != nil {
			panic(err)
		}

		// fmt.Println(string(dst))
		p.src_rpc_golang = string(dst)

		ioutil.WriteFile("./msg/request.go", dst, 0644)
	}

	{

		// rpc csharp
		tpl_txt := `/// Code generated by protocol.define; DO NOT EDIT.
namespace Network.sdk.Excalibur.Protocol
{
	public partial class Msg
	{
		public static void deserialize_call(Msg.Code code, object s, object ctx)
		{
			string json = s.ToString();
			switch (code)
			{
				{{range .Msg.Fields}}case Msg.Code.{{.Name}}:
					{
						var ret = Msg.FromJson<{{.Name}}>(json);
						var f = ctx as rpc_callback<Ret<{{.Name}}>>;
						f(ret);
						break;
					}
					{{end}}
				default:
					{
						break;
					}
			}
		}
	}
}
`
		tpl, e := template.New("msg").Parse(tpl_txt)
		if e != nil {
			panic(e)
		}

		source := bytes.NewBuffer([]byte{})
		e = tpl.Execute(source, p.ret)
		fmt.Println(string(source.Bytes()))

		p.src_rpc_csharp = string(source.Bytes())

		ioutil.WriteFile("./rpc/Rpc.cs", source.Bytes(), 0644)
		ioutil.WriteFile("E:\\Excalibur\\Clent\\Assets\\Network\\sdk\\Excalibur\\Protocol\\Rpc.cs", source.Bytes(), 0644)
	}

	return p.ret
}

func (p *interp) get_odm_golang() string {
	return p.src_odm_golang
}

func (p *interp) get_odm_csharp() string {
	return p.src_odm_csharp
}

func (p *interp) get_msg_golang() string {
	return p.src_msg_golang
}

func (p *interp) get_msg_csharp() string {
	return p.src_msg_csharp
}

func (p *interp) get_rpc_golang() string {
	return p.src_rpc_golang
}

func (p *interp) get_rpc_csharp() string {
	return p.src_rpc_csharp
}

func lowercaseFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
