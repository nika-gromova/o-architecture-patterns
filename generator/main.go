package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

type adapterData struct {
	Name    string
	Imports []string
	Methods []methodData
}

type methodData struct {
	AdapterName  string
	Name         string
	Args         []paramData
	ReturnParams []paramData
	Body         string
	IsCommand    bool
}

type paramData struct {
	ParamName string
	ParamType string
}

var helpers template.FuncMap = map[string]interface{}{
	"notLast": func(index int, len int) bool {
		return index+1 != len
	},
}

const (
	adapterTemplate = `
type {{.Name}} struct {
	obj any
	ctx context.Context
}

func (a *{{.Name}}) SetObject (obj any) {
    a.obj = obj
}
func (a *{{.Name}}) SetContext (ctx context.Context) {
    a.ctx = ctx
}
{{range .Methods}}
{{$lenArgs := len .Args}}{{$lenReturn := len .ReturnParams}}
func (a *{{.AdapterName}}) {{.Name}}({{range $i, $item := .Args}}{{.ParamName}} {{.ParamType}}{{if (notLast $i $lenReturn)}}, {{end}}{{end}}) ({{range $i, $item := .ReturnParams}}{{.ParamType}}{{if (notLast $i $lenReturn)}}, {{end}}{{end}}) {
    {{.Body}}
}{{end}}`

	imports = `
import (
	"context"
	"fmt"

	"github.com/nika-gromova/o-architecture-patterns/game/base"
	"github.com/nika-gromova/o-architecture-patterns/game/ioc"
)`
)

var (
	inputFileName = flag.String("input", "test.go",
		"путь до файл, в котором содержится интерфейс, для которого необходимо сгенерировать адаптер")
	outputFileName = flag.String("output", "test_generated.go",
		"результирующий файл со сгенеренным адаптером")
	interfaceName = flag.String("interface", "", "интерфейс для генерации")
)

func main() {
	flag.Parse()
	process(*inputFileName, *interfaceName, *outputFileName)
}

func process(inputFile, interfaceName, outputFile string) {
	tmpl, err := template.New("adapter").Funcs(helpers).Parse(adapterTemplate)
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, inputFile, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	out, _ := os.Create(outputFile)
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(out)

	adapter := &adapterData{
		Name: fmt.Sprintf("%sAdapter", interfaceName),
	}

	for _, f := range node.Decls {
		genD, ok := f.(*ast.GenDecl)
		if !ok {
			continue
		}

		var neededInterface *ast.InterfaceType
		for _, spec := range genD.Specs {
			currType, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			currInterface, ok := currType.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}
			if currType.Name.Name != interfaceName {
				continue
			}
			neededInterface = currInterface
			break
		}
		if neededInterface == nil {
			continue
		}

		for _, method := range neededInterface.Methods.List {
			methodD := methodData{
				AdapterName: adapter.Name,
				Name:        method.Names[0].Name,
			}
			function := method.Type.(*ast.FuncType)

			for _, param := range function.Params.List {
				methodD.Args = append(methodD.Args, parseParam(param))
			}

			for _, result := range function.Results.List {
				methodD.ReturnParams = append(methodD.ReturnParams, parseParam(result))
			}

			adapter.Methods = append(adapter.Methods, methodD)
		}
	}

	adapterTxt, err := generateAdapter(tmpl, adapter)
	if err != nil {
		log.Fatal(err)
	}
	outputDir, _ := path.Split(*outputFileName)
	if outputDir == "" {
		outputDir = "main"
	}

	fmt.Fprintln(out, `package `+outputDir)
	fmt.Fprintln(out) // empty line
	fmt.Fprintln(out, imports)
	fmt.Fprintln(out) // empty line

	fmt.Fprintln(out, adapterTxt)
}

func parseParam(field *ast.Field) paramData {
	var (
		varName string
		varType = parseType(field.Type)
	)
	if len(field.Names) > 0 {
		varName = field.Names[0].Name
	}

	return paramData{varName, varType}
}

func parseType(expr ast.Expr) string {
	var varType string

	switch expr := expr.(type) {
	case *ast.Ident:
		varType = expr.String()
	case *ast.SelectorExpr:
		varType = expr.X.(*ast.Ident).Name + "." + expr.Sel.Name
	case *ast.ArrayType:
		varType = "[]" + parseType(expr.Elt)
	case *ast.MapType:
		varType = "map[" + parseType(expr.Key) + "]" + parseType(expr.Value)
	}

	return varType
}

func generateAdapter(tmpl *template.Template, data *adapterData) (string, error) {
	for i := range data.Methods {
		body := generateMethodBody(data.Methods[i])
		data.Methods[i].Body = body
	}
	var result strings.Builder
	err := tmpl.Execute(&result, data)
	if err != nil {
		return "", fmt.Errorf("не удалось сгенерировать адаптер по шаблону: %w", err)
	}
	return result.String(), nil
}

func generateIocResolve(data methodData) string {
	methodName := strings.ToLower(data.Name)
	baseName := strings.TrimPrefix(data.AdapterName, "Adapter")
	var (
		key  string
		args = generateParamNames(data.Args)
	)
	key = getKey(baseName, methodName, "get")
	if key != "" {
		key = getKey(baseName, methodName, "set")
	} else {
		key = generateKey(baseName, "", methodName)
	}
	return generateIoC(key, args)
}

func getKey(baseName, methodName string, checkString string) string {
	if strings.HasPrefix(methodName, checkString) {
		entityName := strings.TrimPrefix(methodName, checkString)
		return generateKey(baseName, entityName, checkString)
	}
	return ""
}

func generateKey(baseName, entityName, methodName string) string {
	if entityName == "" {
		return fmt.Sprintf("%s:%s", baseName, methodName)
	}
	return fmt.Sprintf("%s:%s.%s", baseName, entityName, methodName)
}

func isMethodCommand(data methodData) bool {
	methodName := strings.ToLower(data.Name)
	if strings.HasPrefix(methodName, "get") {
		return false
	}
	return true
}

func generateParamNames(params []paramData) string {
	return strings.Join(lo.Map(params, func(item paramData, _ int) string {
		return item.ParamName
	}), ", ")
}

func generateIoC(key, args string) string {
	return fmt.Sprintf(`ioc.Resolve(a.ctx, "%s", a.obj, %s)`, key, args)
}

func generateMethodBody(data methodData) string {
	iocResolve := generateIocResolve(data)
	isCommand := isMethodCommand(data)
	var returnBody = `return err`
	var returnArg *paramData
	if len(data.ReturnParams) > 1 {
		returnArg = &data.ReturnParams[0]
	}
	if returnArg != nil {
		returnBody = fmt.Sprintf(`var resultTyped %s
        return resultTyped, err`,
			returnArg.ParamType)
	}

	body := fmt.Sprintf(`result, err := %s
    if err != nil {
        %s
    }`, iocResolve, returnBody)
	if isCommand {
		body += `
	resultCommand, ok := result.(base.Command)
	if !ok {
		return fmt.Errorf("failed to convert command")
	}
	return resultCommand.Execute()`
	} else {
		if returnArg != nil {
			body += fmt.Sprintf(`
	resultTyped, ok := result.(%s)
	if !ok {
		return resultTyped, fmt.Errorf("type conversion error")
	}
	return resultTyped, nil`,
				returnArg.ParamType)
		}
	}
	return body
}
