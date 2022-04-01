package h001

import (
	"bytes"
	"flag"
	"go/ast"
	"go/token"
	"io/ioutil"
	"strings"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"golang.org/x/tools/go/analysis"
)

const analyzerName = "h001"

const doc = `check for schema arguments that do not exist in Acceptance Test Checks.
The h001 analyzer reports cases of schema arguments that do not exist in Acceptance Test Checks.`

var (
	fields string
)

func parseFlags() flag.FlagSet {
	var flags = flag.NewFlagSet(analyzerName, flag.ExitOnError)
	flags.StringVar(&fields, "fields", "enterprise_project_id", "will check whether these fields exist in the test file")
	return *flags
}

var Analyzer = &analysis.Analyzer{
	Name:  analyzerName,
	Flags: parseFlags(),
	Doc:   doc,
	Requires: []*analysis.Analyzer{
		commentignore.Analyzer,
	},
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	ignorer := pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	checkFields := strings.Split(fields, ",")

	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			smap, ok := n.(*ast.CompositeLit)
			if !ok {
				return true
			}

			if !schema.IsMapStringSchema(smap, pass.TypesInfo) {
				return true
			}

			if ignorer.ShouldIgnore(analyzerName, smap) {
				return true
			}

			//测试字段是否在测试用例中覆盖，只要存在即可
			testFilePath := getTestFilePath(getFilePathOfNode(render(pass.Fset, file.Pos())))

			input, err := ioutil.ReadFile(testFilePath)
			if err != nil {
				return true
			}

			for _, attributeName := range schema.GetSchemaMapAttributeNames(smap) {
				switch t := attributeName.(type) {
				default:
					continue
				case *ast.BasicLit:
					value := strings.Trim(t.Value, `"`)
					if isNeedCheck(checkFields, value) && !bytes.Contains(input, []byte(value)) {
						pass.Reportf(t.Pos(), "%s: schema argument %q should be used in test file", analyzerName, value)
					}
				}
			}
			return true
		})
	}

	return nil, nil
}

func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := ast.Fprint(&buf, fset, x, ast.NotNilFilter); err != nil {
		panic(err)
	}
	return buf.String()
}

//格式 "  0  /sss/sss/aa.go:1:1"
func getFilePathOfNode(nodePos string) string {
	tmp := strings.Split(strings.TrimSpace(nodePos), " ")
	if len(tmp) < 1 {
		return nodePos
	}
	path := strings.Split(tmp[len(tmp)-1], ":")[0]
	return path
}

func getTestFilePath(path string) string {
	tmp := strings.Replace(path, ".go", "_test.go", 1)
	if strings.Contains(tmp, "services") {
		tmp = strings.Replace(tmp, `/services/`, `/services/acceptance/`, -1)
	}
	return tmp
}

func isNeedCheck(checkFieldConfig []string, value string) bool {
	flag := false
	for _, v := range checkFieldConfig {
		if v == value {
			flag = true
		}
	}
	return flag
}
