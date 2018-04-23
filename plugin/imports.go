package plugin

import (
	"sort"
	"strings"

	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

/* --------- Response file import cleaning -------- */

// Imports that are added by default but unneeded in GORM code
var unneededImports = []string{
	"import proto \"github.com/gogo/protobuf/proto\"\n",
	"import _ \"github.com/infobloxopen/protoc-gen-gorm/options\"\n",
	// if needed will be imported with an alias
	"import _ \"github.com/infobloxopen/protoc-gen-gorm/types\"\n",
	"var _ = proto.Marshal\n",
}

// CleanImports removes extraneous imports and lines from a proto response
// file Content
func CleanImports(pFileText *string) *string {
	if pFileText == nil {
		return pFileText
	}
	fileText := *pFileText
	for _, dep := range unneededImports {
		fileText = strings.Replace(fileText, dep, "", -1)
	}
	return &fileText
}

/* --------- Plugin level import handling --------- */

func (p *SqlxPlugin) resetImports() {
	//p.wktPkgName = ""
	//p.gormPkgName = ""
	//p.lftPkgName = ""
	p.usingUUID = false
	p.usingTime = false
	p.usingAuth = false
}

// GenerateImports writes out required imports for the generated files
func (p *SqlxPlugin) GenerateImports(file *generator.FileDescriptor) {
	var stdImports []string
	githubImports := map[string]string{}
	//if p.gormPkgName != "" {
	//	stdImports = append(stdImports, "context", "errors")
	//	githubImports[p.gormPkgName] = "github.com/jinzhu/gorm"
	//	githubImports[p.lftPkgName] = "github.com/infobloxopen/atlas-app-toolkit/op/gorm"
	//}
	//if p.usingUUID {
	//	githubImports["uuid"] = "github.com/satori/go.uuid"
	//	githubImports["gtypes"] = "github.com/infobloxopen/protoc-gen-gorm/types"
	//}
	//if p.usingTime {
	//	stdImports = append(stdImports, "time")
	//	githubImports["ptypes"] = "github.com/golang/protobuf/ptypes"
	//}


	p.PrintImport("log", "github.com/sirupsen/logrus")
	p.PrintImport("", "github.com/micro-grpc/mbox/lib")

	if p.isJSONB {
		p.PrintImport("", "encoding/json")
		p.PrintImport("", "database/sql/driver")
	}
	if p.isGORM {
		p.PrintImport("", "github.com/jinzhu/gorm")
		//if p.driver == "postgres" {
		//	githubImports["_"] = "github.com/jinzhu/gorm/dialects/postgres"
		//}
	}
	if p.isSqlx {
		p.PrintImport("", "github.com/jmoiron/sqlx")
		//p.PrintImport("", "github.com/jmoiron/sqlx/reflectx")
	}
	//if p.driver == "postgres" {
	//	githubImports["sql.postgres"] = "_ \"github.com/lib/pq\""
	//}

	//if p.usingAuth {
	//	githubImports["auth"] = "github.com/infobloxopen/atlas-app-toolkit/mw/auth"
	//}

	sort.Strings(stdImports)
	for _, dep := range stdImports {
		p.PrintImport(dep, dep)
	}
	p.P()
	var keys []string
	for k := range githubImports {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		p.PrintImport(key, githubImports[key])
	}
	p.P()
}