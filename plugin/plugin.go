package plugin

import (
	"fmt"
	"github.com/gogo/protobuf/gogoproto"
	//google_protobuf "google/protobuf"
	//google_protobuf "code.google.com/p/gogoprotobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	jgorm "github.com/jinzhu/gorm"
	"github.com/jinzhu/inflection"
	"github.com/micro-grpc/protoc-gen-sqlx/pb/sql"
	"github.com/sirupsen/logrus"
	"strings"
)

var wellKnownTypes = map[string]string{
	"StringValue": "*string",
	"DoubleValue": "*double",
	"FloatValue":  "*float",
	"Int32Value":  "*int32",
	"Int64Value":  "*int64",
	"UInt32Value": "*uint32",
	"UInt64Value": "*uint64",
	"BoolValue":   "*bool",
	//  "BytesValue" : "*[]byte",
}

//types, only for existence checks
var convertibleTypes = make(map[string]struct{})
var isGormTypes = make(map[string]struct{})
var isJSONBTypes = make(map[string]struct{})

// All message objects
var typeNames = make(map[string]*generator.Descriptor)

// SqlxPlugin implements the plugin interface and creates SqlX code from .protos
type SqlxPlugin struct {
	*generator.Generator
	//wktPkgName  string
	//gormPkgName string
	//lftPkgName  string // 'Locally Famous Types', used for collection operators
	usingUUID bool
	usingTime bool
	usingAuth bool
	usingGRPC bool
	isGORM bool
	isSqlx bool
	isJSONB bool
	driver string

}

// Name identifies the plugin
func (p *SqlxPlugin) Name() string {
	return "sqlx"
}

// Init is called once after data structures are built but before
// code generation begins.
func (p *SqlxPlugin) Init(g *generator.Generator) {
	p.Generator = g
	p.driver = "postgres"
}

func (p *SqlxPlugin) Generate(file *generator.FileDescriptor) {
	p.resetImports()



	for _, msg := range file.Messages() {
		// We don't want to bother with the MapEntry stuff
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		unlintedTypeName := generator.CamelCaseSlice(msg.TypeName())
		typeNames[unlintedTypeName] = msg
		if opts := getMessageOptions(msg); opts != nil && opts.Jsonb {
			isJSONBTypes[unlintedTypeName] = struct{}{}
		}
		if opts := getMessageOptions(msg); opts != nil && opts.Orm {
			convertibleTypes[unlintedTypeName] = struct{}{}
			logrus.Warnln(unlintedTypeName)
		}
		if opts := getMessageOptions(msg); opts != nil && opts.Gorm {
			isGormTypes[unlintedTypeName] = struct{}{}
		}
	}

	//отключил пока ненадо
	for _, msg := range file.Messages() {
		unlintedTypeName := generator.CamelCaseSlice(msg.TypeName())
		if _, exists := convertibleTypes[unlintedTypeName]; exists {
			//logrus.Errorln("orm: TRUE")
			p.isSqlx = true
		}
	}

	if p.isSqlx {
		p.P()
		p.P(`////////////////////////// GLOBAL vars :))))))`)
		p.P(`var preOne = "SELECT to_jsonb(f0) AS data FROM (%s) AS f0"`)
		p.P(`var preMulti = "SELECT COALESCE(jsonb_agg(f0), '[]'::jsonb) AS data FROM (%s) AS f0"`)
		p.P(`var preRows = "SELECT to_jsonb(f0) AS data FROM (%s) AS f0"`)
		p.P()
		p.P(`type Result struct {`)
		p.P(`Total int64`)
		p.P(`}`)
		p.P()
		//p.generateDefaultHandlers(file)
		p.P()
	}

	for _, msg := range file.Messages() {
		sqlDriver := "sqlx.DB"
		// We don't want to bother with the MapEntry stuff
		if msg.DescriptorProto.GetOptions().GetMapEntry() {
			continue
		}
		unlintedTypeName := generator.CamelCaseSlice(msg.TypeName())
		if _, exists := convertibleTypes[unlintedTypeName]; !exists {
			continue
		}
		if _, exists := isJSONBTypes[unlintedTypeName]; exists {
			p.isJSONB =true
		}
		if _, exists := isGormTypes[unlintedTypeName]; exists {
			sqlDriver = "gorm.DB"
			p.isGORM = true
		}
		p.isSqlx = true

		logrus.Println("Generate ORM structure")
		p.generateMessages(msg)

		if _, exists := isGormTypes[unlintedTypeName]; exists {
			logrus.Println("Generate GORM TableName")
			p.generateTableNameFunction(msg)

		}

		logrus.Println("Generate mBox structure", sqlDriver)
		p.generateMBboxStructure(msg, sqlDriver)
		p.generateTableName(msg)
		p.generateMBboxFields(msg)

		if _, exists := isJSONBTypes[unlintedTypeName]; exists {
			logrus.Println("Generate sql.Scanner and driver.Valuer for JSONB")
			p.generateCovertJSONBFunction(msg)
		}

		p.generateMBboxMetods(msg)
		p.generateConvertFunctions(msg)
	}
	p.P()
	p.P(`////////////////////////// CURDL for objects`)
	if p.isSqlx {
		p.generateGlobalApplyFunction()
	}
	//p.generateDefaultHandlers(file)
	p.P()

	//p.generateDefaultServer(file)
}

func (p *SqlxPlugin) generateMessages(message *generator.Descriptor) {
	typeName := p.TypeName(message)
	logrus.Println("typeName:", typeName)
	p.generateMessageHead(message)
	for _, field := range message.Field {
		fieldName := generator.CamelCase(field.GetName())
		ormFieldName := fieldName
		fieldType, _ := p.GoType(message, field)
		var tagString string
		if field.Options != nil {
			vv, err := proto.GetExtension(field.Options, gogoproto.E_Customname)
			if err == nil {
				//ext := vv.(*string)
				//res := vv.(*descriptor.FieldDescriptorProto)
				//res := field.GetOptions()
				res := fmt.Sprintf("%+v", *vv.(*string))
				ormFieldName = res
				logrus.Warningln("custom name for:", fieldName, "to:", res)
				//ormFieldName = fieldName
				//v, _ := vv.(*google_protobuf.ExtensionDescriptor)
				//res := vv.(*generator.ExtensionDescriptor)
				//v := vv.(*descriptor.DescriptorProto)
				//if len(res.GetCustomname()) > 0 {

					//logrus.Warnln(field.GetOptions())
				//logrus.Warnln(res)
				//}
				//if ok {
				//	logrus.Warningln("customname:", v, "origin:", fieldName)
				//}
			}
			v, err := proto.GetExtension(field.Options, sql.E_Field)
			opts, valid := v.(*sql.SqlxFieldOptions)
			if err == nil && valid && opts != nil {
				if opts.GetDrop() {
					p.P(`// Skipping field: `, generator.CamelCase(field.GetName()))
					continue
				}
				if len(opts.GetName()) > 0 {
					//logrus.Println("Field name:", opts.GetName())
					ormFieldName = opts.GetName()
				}
				if len(opts.GetColname()) > 0 {
					tagString = fmt.Sprintf("db:\"%s\"", opts.GetColname())
				}
				tags := v.(*sql.SqlxFieldOptions).GetTags()
				if len(tags) > 0 {
					if len(tagString) > 0 {
						tagString = fmt.Sprintf("%s %s", tagString, tags)
					} else {
						tagString = fmt.Sprintf("%s", tags)
					}
				}
				if len(tagString) > 0 {
					tagString = fmt.Sprintf("`%s`", tagString)
				}
			}
			//logrus.Println("field.Options:", field.Options)
		}
		logrus.Println("ormFieldName:", ormFieldName, "fieldName:", fieldName, "fieldType:", fieldType, "tags:", tagString)
		p.P(ormFieldName, " ", fieldType, tagString)
	}
	p.P(`}`)

}

// generateTableNameFunction the function to set the gorm table name
// back to gorm default, removing "ORM" suffix
func (p *SqlxPlugin) generateTableNameFunction(message *generator.Descriptor) {
	typeName := p.TypeName(message)

	p.P(`// TableName overrides the default tablename generated by GORM`)
	p.P(`func (`, typeName, `Type) TableName() string {`)

	tableName := inflection.Plural(jgorm.ToDBName(message.GetName()))
	if opts := getMessageOptions(message); opts != nil && len(opts.GetTable()) > 0 {
		tableName = opts.GetTable()
	}
	p.P(`return "`, tableName, `"`)
	p.P(`}`)
	p.P()
}

func (p *SqlxPlugin) generateTableName(message *generator.Descriptor) {
	typeName := p.TypeName(message)

	tableName := inflection.Plural(jgorm.ToDBName(message.GetName()))
	if opts := getMessageOptions(message); opts != nil && len(opts.GetTable()) > 0 {
		tableName = opts.GetTable()
	}
	p.P(`// TableName overrides the default tablename generated by Query`, typeName)
	p.P(`func (p *Query`, typeName, `) TableName() string {`)
	p.P(`return "`, tableName, `"`)
	p.P(`}`)
	p.P()
}

func (p *SqlxPlugin) generateCovertJSONBFunction(message *generator.Descriptor) {
	typeName := p.TypeName(message)
	p.P(`func (p *Query`, typeName, `) Scan(val interface{}) error {`)
	p.P(`return json.Unmarshal(val.([]byte), p)`)
	p.P(`}`)
	p.P(`func (p *Query`, typeName, `) Value() (driver.Value, error) {`)
	p.P(`return json.Marshal(p)`)
	p.P(`}`)
}

func (p *SqlxPlugin) generateGlobalApplyFunction()  {
	p.P(`func applyField(fields []string) string {
	field := ""
	for _, v := range fields {
		if len(field) == 0 {
			field = v
		} else {
			field = fmt.Sprintf("%s, %s", field, v)
		}
	}
	return field
}`)
	p.P()
	p.P(`func applyFiltering(filtering []*Filtering) (string, []interface{}) {
	filter := ""
	var filterValue []interface{}
	for key, val := range filtering {
		if len(filter) > 0 {
			filter = fmt.Sprintf("%s AND %s%s", filter, val.Name, lib.FilteringMode(val.Mode.String(), key+1))
		} else {
			filter = fmt.Sprintf(" %s%s", val.Name, lib.FilteringMode(val.Mode.String(), key+1))
		}
		if val.Mode.String() != "IS_NULL" && val.Mode.String() != "NOT_NULL" {
			filterValue = append(filterValue, val.Value)
		}
	}
	return filter, filterValue
}`)
	p.P()
}

func (p *SqlxPlugin) generateMBboxStructure(message *generator.Descriptor, sqlDriver string) {
	typeName := p.TypeName(message)

	p.P(`type Query`, typeName, ` struct {
	Verbose     int
	DB          *`, sqlDriver, `
	driver      string
	debug       bool
	created     bool
	table       string
	field       string
	fields      []string
	defFields 	map[string]interface{}
	query       string
	filter      string
	filterValue []interface{}
	sort        string
	order       string
	limit       int64
	offset      int64
	current     int64
	total       int64
}`)
	p.P()

	p.P(`// NewQuery`, typeName, ` - initialize Query`, typeName)
	p.P(`func NewQuery`, typeName, `(verbose int) *Query`, typeName, ` {`)
	p.P(`p := Query`, typeName, `{
		Verbose: verbose,
		fields:   `, typeName, `Fields(),
		defFields: `, typeName, `Contains(),
		limit:   1000,
		offset:  0,
		sort:    "id",
		order:   "ASC",`)
	p.P(`}`)
	p.P(`p.table = p.TableName()`)
	p.P(`p.field = p.DefaultFields()`)
	p.P(`return &p`)
	p.P(`}`)
	p.P()

	p.P(`// SetDB setting database`)
	p.P(`func (p *Query`, typeName, `) SetDB(db *sqlx.DB) {`)
	p.P(`p.DB = db`)
	p.P(`}`)

	p.P()
	p.P(`// Close closes the database, releasing any open resources.`)
	p.P(`func (p *Query`, typeName, `) Close() error {`)
	p.P(`return p.DB.Close()`)
	p.P(`}`)
	p.P()
}

func (p *SqlxPlugin) generateMBboxMetods(message *generator.Descriptor) {
	typeName := p.TypeName(message)
	request := "Request"

	if opts := getMessageOptions(message); opts != nil && len(opts.GetRequest()) > 0 {
	request = opts.GetRequest()
	}

	p.P(`// BuildOneQuery - return one row`)
	p.P(`func (p *Query`, typeName, `) BuildOneQuery(in *`, request, `, field string) (string, []interface{}) {`)
	p.P(`return p.BuildQuery(in, field, false, true)`)
	p.P(`}`)
	p.P()

	p.P(`// BuildMultiQuery - return rows`)
	p.P(`func (p *Query`, typeName, `) BuildMultiQuery(in *`, request, `, field string) (string, []interface{}) {`)
	p.P(`return p.BuildQuery(in, field, false, false)`)
	p.P(`}`)
	p.P()

	p.P(`func (p *Query`, typeName, `) Count(in *`, request, `) (int64) {
  var result Result
  wh := ""
  filter, args := applyFiltering(in.Filter)
  if len(filter) > 0 {
    wh = fmt.Sprintf(" WHERE%s", filter)
  }
  str := fmt.Sprintf("SELECT COUNT(*) AS total FROM %s%s LIMIT 1", p.table, wh)
  if p.Verbose > 3 {
    log.Println("[Count] QUERY:", str, "params:", args)
  }
  err := p.DB.QueryRowx(str, args...).StructScan(&result)
  if err != nil {
    log.Errorln("[Count]", err)
  }
  return result.Total`)
	p.P(`}`)
	p.P()

	p.P(`func (p *Query`, typeName, `) BuildQuery(in *`, request, `, f string, reserved bool, one bool) (string, []interface{}) {
	wh := ""
  filter, filterValue := applyFiltering(in.Filter)
  field := p.field

  if len(f) > 0 {
    field = f
  } else if len(in.Field) > 0 {
    field = applyField(in.Field)
  }
	if len(filter) > 0 {
		wh = fmt.Sprintf(" WHERE%s", filter)
	}

  limit := p.limit
  log.Warningln("LIMIT", in.GetLimit())

	order := p.order
  offset := p.offset
	if in.GetPage() > 0 {
	  log.Warningln("TODO pagination BuildQuery")
  } else if in.GetFirst() > 0 {
    offset = 0
    order = "ASC"
	  limit = in.GetFirst()
  } else if in.GetLast() > 0 {
    offset = 0
    order = "DESC"
    limit = in.GetLast()
  }

	str := fmt.Sprintf("SELECT %s FROM %s%s", field, p.table, wh)
		if one {
			return fmt.Sprintf("%s ORDER BY %s %s LIMIT 1", str, p.sort, order), filterValue
		}
		return fmt.Sprintf("%s ORDER BY %s %s LIMIT %d OFFSET %d", str, p.sort, order, limit, offset), filterValue`)
	p.P(`}`)
	p.P()
}

// generateMessageComment pulls from the proto file comment or creates a
// default comment if none is present there, and writes the signature and
// fields from the proto file options
func (p *SqlxPlugin) generateMessageHead(message *generator.Descriptor) {
	typeName := p.TypeName(message)
	typeNameOrm := fmt.Sprintf("%sType", typeName)
	// Check for a comment, generate a default if none is provided
	comment := p.Comments(message.Path())
	commentStart := strings.Split(strings.Trim(comment, " "), " ")[0]
	if generator.CamelCase(commentStart) == typeName || commentStart == typeNameOrm {
		comment = strings.Replace(comment, commentStart, typeNameOrm, 1)
	} else if len(comment) == 0 {
		comment = fmt.Sprintf(" %s no comment was provided for message type", typeNameOrm)
	} else if len(strings.Replace(comment, " ", "", -1)) > 0 {
		comment = fmt.Sprintf(" %s %s", typeNameOrm, comment)
	} else {
		comment = fmt.Sprintf(" %s no comment provided", typeNameOrm)
	}
	p.P(`//`, comment)
	p.P(`type `, typeNameOrm, ` struct {`)
	// Checking for any ORM only fields specified by option (sql.opts).include
	if opts := getMessageOptions(message); opts != nil {
		for _, field := range opts.Include {
			tagString := ""
			if len(field.GetTags()) > 0 {
				tagString = fmt.Sprintf("`%s`", field.GetTags())
			}
			p.P(generator.CamelCase(field.Name), ` `, field.Type, ` `, tagString)
		}
		// специфические настройкидля bbox
		if opts.GetUser() {
			p.P("UserID int64")
			//p.P("User User")
		}
		if opts.GetProduct() {
			p.P("ProductID int64")
		}
	}
}

// generateMBboxFields
func (p *SqlxPlugin) generateMBboxFields(message *generator.Descriptor) {
	typeName := p.TypeName(message)
	fields := ""
	fieldsArray := ""
	fieldMap := []string{}
	for _, field := range message.Field {
		// Checking if field is skipped
		ormField := field.GetName()
		pbField := field.GetName()
		if field.Options != nil {
			v, err := proto.GetExtension(field.Options, sql.E_Field)
			opts, valid := v.(*sql.SqlxFieldOptions)
			if err == nil && valid && opts != nil {
				if opts.GetDrop() {
					continue
				}
				if len(opts.GetColname()) > 0 {
					//logrus.Println("Field name:", opts.GetName())
					ormField = opts.GetColname()
				}
			}
		}
		rederField := ormField
		if ormField != pbField {
			rederField = fmt.Sprintf("%s AS %s", ormField, pbField)
			fieldMap = append(fieldMap, fmt.Sprintf("set[\"%s\"] = \"%s AS %s\"",  ormField, ormField, pbField))
		} else {
			fieldMap = append(fieldMap, fmt.Sprintf("set[\"%s\"] = \"%s\"",  ormField, ormField))
		}
		if len(fields) > 0 {
			fields = fmt.Sprintf("%s, %s", fields, rederField)
			fieldsArray = fmt.Sprintf("%s, \"%s\"", fieldsArray, rederField)
		} else {
			fields = rederField
			fieldsArray = fmt.Sprintf("\"%s\"", ormField)
		}
	}
	if len(fields) > 0 {
		p.P(`// DefaultFields all fields generated by Query`, typeName)
		p.P(`func (p *Query`, typeName, `) DefaultFields() string  {`)
		p.P(`return "`, fields, `"`)
		p.P(`}`)
		p.P()

		p.P(`// DefaultFields all fields generated by Query`, typeName)
		p.P(`func `, typeName, `Fields() []string  {`)
		p.P(`return []string{`, fieldsArray, `}`)
		p.P(`}`)
		p.P()

		p.P(`func `, typeName, `Contains() (map[string]interface{}) {`)
		p.P(`set := make(map[string]interface{})`)
		for _, val := range fieldMap {
			p.P(val)
		}
		p.P(`return set`)
		p.P(`}`)
		p.P()

		p.P(`func (p *Query`, typeName, `) ContainsField(item string) bool {`)
		p.P(`_, ok := p.defFields[item]`)
		p.P(`return ok`)
		p.P(`}`)
		p.P()
	}
}

// generateMapFunctions creates the converter functions
func (p *SqlxPlugin) generateConvertFunctions(message *generator.Descriptor) {
	typeName := p.TypeName(message)

	///// To Orm
	p.P(`// Convert`, typeName, `ToType takes a pb object and returns an orm object`)
	p.P(`func Convert`, typeName, `ToType (from `,
		typeName, `) (`, typeName, `Type, error) {`)
	p.P(`to := `, typeName, `Type{}`)
	p.P(`var err error`)
	for _, field := range message.Field {
		// Checking if field is skipped
		ormFieldName := generator.CamelCase(field.GetName())
		if field.Options != nil {
			v, err := proto.GetExtension(field.Options, sql.E_Field)
			opts, valid := v.(*sql.SqlxFieldOptions)
			if err == nil && valid && opts != nil {
				if opts.GetDrop() {
					p.P(`// Skipping field: `, generator.CamelCase(field.GetName()))
					continue
				}
				if len(opts.GetName()) > 0 {
					//logrus.Println("Field name:", opts.GetName())
					ormFieldName = opts.GetName()
				}
			}
		}
		p.generateFieldConversion(message, field, ormFieldName, true)
	}
	p.P(`return to, err`)
	p.P(`}`)

	p.P()
	///// To Pb
	p.P(`// Convert`, typeName, `FromType takes an orm object and returns a pb object`)
	p.P(`func Convert`, typeName, `FromType (from `, typeName, `Type) (`,
		typeName, `, error) {`)
	p.P(`to := `, typeName, `{}`)
	p.P(`var err error`)
	for _, field := range message.Field {
		// Checking if field is skipped
		ormFieldName := generator.CamelCase(field.GetName())
		if field.Options != nil {
			v, err := proto.GetExtension(field.Options, sql.E_Field)
			opts, valid := v.(*sql.SqlxFieldOptions)
			if err == nil && valid && opts != nil {
				if opts.GetDrop() {
					p.P(`// Skipping field: `, generator.CamelCase(field.GetName()))
					continue
				}
				if len(opts.GetName()) > 0 {
					//logrus.Println("Field name:", opts.GetName())
					ormFieldName = opts.GetName()
				}
			}
		}
		p.generateFieldConversion(message, field, ormFieldName, false)
	}
	p.P(`return to, err`)
	p.P(`}`)
}

// Output code that will convert a field to/from orm.
func (p *SqlxPlugin) generateFieldConversion(message *generator.Descriptor, field *descriptor.FieldDescriptorProto, ormFieldName string, toORM bool) error {
	fieldName := generator.CamelCase(field.GetName())
	toField := ormFieldName
	fromField := fieldName
	fieldType, _ := p.GoType(message, field)
	if field.IsRepeated() { // Repeated Object ----------------------------------
		if _, exists := convertibleTypes[strings.Trim(fieldType, "[]*")]; exists { // Repeated ORMable type
			fieldType = strings.Trim(fieldType, "[]*")
			dir := "From"
			if toORM {
				dir = "To"
			}

			p.P(`for _, v := range from.`, fieldName, ` {`)
			p.P(`if v != nil {`)
			p.P(`if temp`, fieldName, `, cErr := Convert`, fieldType, dir, `Type (*v); cErr == nil {`)
			p.P(`to.`, fieldName, ` = append(to.`, fieldName, `, &temp`, fieldName, `)`)
			p.P(`} else {`)
			p.P(`return to, cErr`)
			p.P(`}`)
			p.P(`} else {`)
			p.P(`to.`, fieldName, ` = append(to.`, fieldName, `, nil)`)
			p.P(`}`)
			p.P(`}`) // end repeated for
		} else {
			p.P(`// Repeated type `, fieldType, ` is not an ORMable message type`)
		}
	//} else if *(field.Type) == typeEnum { // Singular Enum, which is an int32 ---
	//	if toORM {
	//		p.P(`to.`, fieldName, ` = int32(from.`, fieldName, `)`)
	//	} else {
	//		p.P(`to.`, fieldName, ` = `, fieldType, `(from.`, fieldName, `)`)
	//	}
	//} else if *(field.Type) == typeMessage { // Singular Object -------------
		////Check for WKTs
		//parts := strings.Split(fieldType, ".")
		//coreType := parts[len(parts)-1]
		//// Type is a WKT, convert to/from as ptr to base type
		//if _, exists := wellKnownTypes[coreType]; exists { // Singular WKT -----
		//	if toORM {
		//		p.P(`if from.`, fieldName, ` != nil {`)
		//		p.P(`v := from.`, fieldName, `.Value`)
		//		p.P(`to.`, fieldName, ` = &v`)
		//		p.P(`}`)
		//	} else {
		//		p.P(`if from.`, fieldName, ` != nil {`)
		//		p.P(`to.`, fieldName, ` = &`, p.wktPkgName, ".", coreType,
		//			`{Value: *from.`, fieldName, `}`)
		//		p.P(`}`)
		//	}
		//} else if coreType == protoTypeUUID { // Singular UUID type ------------
		//	if toORM {
		//		p.P(`if from.`, fieldName, ` != nil {`)
		//		p.P(`tempUUID, uErr := uuid.FromString(from.`, fieldName, `.Value)`)
		//		p.P(`if uErr != nil {`)
		//		p.P(`return to, uErr`)
		//		p.P(`}`)
		//		p.P(`to.`, fieldName, ` = &tempUUID`)
		//		p.P(`}`)
		//	} else {
		//		p.P(`if from.`, fieldName, ` != nil {`)
		//		p.P(`to.`, fieldName, ` = &gtypes.UUIDValue{Value: from.`, fieldName, `.String()}`)
		//		p.P(`}`)
		//	}
		//} else if coreType == protoTypeTimestamp { // Singular WKT Timestamp ---
		//	if toORM {
		//		p.P(`if from.`, fieldName, ` != nil {`)
		//		p.P(`if to.`, fieldName, `, err = ptypes.Timestamp(from.`, fieldName, `); err != nil {`)
		//		p.P(`return to, err`)
		//		p.P(`}`)
		//		p.P(`}`)
		//	} else {
		//		p.P(`if to.`, fieldName, `, err = ptypes.TimestampProto(from.`, fieldName, `); err != nil {`)
		//		p.P(`return to, err`)
		//		p.P(`}`)
		//	}
		//} else if _, exists := convertibleTypes[strings.Trim(fieldType, "[]*")]; exists {
		//	// Not a WKT, but a type we're building converters for
		//	fieldType = strings.Trim(fieldType, "*")
		//	dir := "From"
		//	if toORM {
		//		dir = "To"
		//	}
		//	p.P(`if from.`, fieldName, ` != nil {`)
		//	p.P(`temp`, fieldType, `, err := Convert`, fieldType, dir, `ORM (*from.`, fieldName, `)`)
		//	p.P(`if err != nil {`)
		//	p.P(`return to, err`)
		//	p.P(`}`)
		//	p.P(`to.`, fieldName, ` = &temp`, fieldType)
		//	p.P(`}`)
		//}
	} else { // Singular raw ----------------------------------------------------
		if toORM {
			p.P(`to.`, toField, ` = from.`, fromField)
		} else {
			p.P(`to.`, fromField, ` = from.`, toField)
		}
	}
	return nil
}