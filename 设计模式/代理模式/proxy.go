package proxy

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"strings"
)

func generate(file string) (string, error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		return "", nil
	}

	//获取代理需要的数据
	data := proxyData{
		Package: f.Name.Name,
	}

	cmap := ast.NewCommentMap(fset, f, f.Comments)
	for node, group := range cmap {
		// 从注释 @proxy 接口名，获取接口名称
		name := getProxyInterfaceName(group)
		if name == "" {
			continue
		}

		// 获取代理的类名
		data.ProxyStructName = node.(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Name.Name

		// 从文件中查找接口
		obj := f.Scope.Lookup(name)

		t := obj.Decl.(*ast.TypeSpec).Type.(*ast.InterfaceType)

		for _, field := range t.Methods.List {
			fc := field.Type.(*ast.FuncType)

			method := &proxyMethod{
				Name: field.Names[0].Name,
			}

			// 获取方法的参数和返回值
			method.Params, method.ParamNames = getParamsOrResults(fc.Params)

			data.Methods = append(data.Methods, method)
		}
	}

	// 生成文件
	tpl, err := template.New("").Parse(proxyTpl)
	if err != nil {
		return "", nil
	}

	buf := &bytes.Buffer{}
	if err := tpl.Execute(buf, data); err != nil {
		return "", nil
	}

	// 使用 go fmt 对生成的代码进行格式化
	src, err := format.Source(buf.Bytes())
	if err != nil {
		return "", nil
	}

	return string(src), nil
}

func getParamsOrResults(fields *ast.FieldList) (string, string) {
	var (
		params    []string
		paraNames []string
	)

	for i, param := range fields.List {
		var names []string
		for _, name := range param.Names {
			names = append(names, name.Name)
		}

		if len(names) == 0 {
			names = append(names, fmt.Sprintf("r%d", i))
		}

		paraNames = append(paraNames, names...)

		// 参数名加参数类型组成完整的参数
		param := fmt.Sprintf("%s %s",
			strings.Join(names, ","),
			param.Type.(*ast.Ident).Name,
		)
		params = append(params, strings.TrimSpace(param))
	}

	return strings.Join(params, ","), strings.Join(paraNames, ",")
}

func getProxyInterfaceName(groups []*ast.CommentGroup) string {
	for _, cgroup := range groups {
		for _, comment := range cgroup.List {
			if strings.Contains(comment.Text, "@proxy") {
				interfaceName := strings.TrimLeft(comment.Text, "// @proxy")
				return strings.TrimSpace(interfaceName)
			}
		}
	}
	return ""
}

const proxyTpl = `
package {{.Package}}

type {{ .ProxyStructName }}Proxy struct {
	child *{{ .ProxyStructName }}
}

func New{{ .ProxyStructName }}Proxy(child *{{ .ProxyStructName }}) *{{ .ProxyStructName }}Proxy {
	return &{{ .ProxyStructName }}Proxy{child: child}
}

{{ range .Methods }}
func (p *{{$.ProxyStructName}}) {{ .Name }} ({{ .Params }}) ({{ .Results }}) {
	start := time.Now()

	{{ .ResultNames }} = p.child.{{ .Name }}({{ .ParamNames }})

	log.Printf("user login cost time: %s", time.Now().Sub(start)

	return {{ .ResultNames }}
}

{{ end }}
`

type proxyData struct {
	// 包名
	Package string

	// 需要代理的类名
	ProxyStructName string

	// 需要代理的方法
	Methods []*proxyMethod
}

type proxyMethod struct {
	// 方法名
	Name string

	// 参数、含参数类型
	Params string

	// 参数名
	ParamNames string

	// 返回值
	Results string

	// 返回值名
	ResultNames string
}
