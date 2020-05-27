// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/example.proto

package example

import "encoding/json"
import "database/sql/driver"
import "github.com/jmoiron/sqlx"
import log "github.com/sirupsen/logrus"
import "github.com/weiwolves/proto-gen-sqlx/lib"

import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/weiwolves/protoc-gen-sqlx/pb/sql"
import _ "github.com/mwitkow/go-proto-validators"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

////////////////////////// GLOBAL vars :))))))
var preOne = "SELECT to_jsonb(f0) AS data FROM (%s) AS f0"
var preMulti = "SELECT COALESCE(jsonb_agg(f0), '[]'::jsonb) AS data FROM (%s) AS f0"
var preRows = "SELECT to_jsonb(f0) AS data FROM (%s) AS f0"

type Result struct {
	Total int64
}

// ExampleItemType no comment was provided for message type
type ExampleItemType struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Name      string `db:"name"`
	// Skipping field: Mess
	MoveID int64 `db:"device_id"`
}
type QueryExampleItem struct {
	Verbose     int
	DB          *sqlx.DB
	driver      string
	debug       bool
	created     bool
	table       string
	field       string
	fields      []string
	defFields   map[string]interface{}
	query       string
	filter      string
	filterValue []interface{}
	sort        string
	order       string
	limit       int64
	offset      int64
	current     int64
	total       int64
}

// NewQueryExampleItem - initialize QueryExampleItem
func NewQueryExampleItem(verbose int) *QueryExampleItem {
	p := QueryExampleItem{
		Verbose:   verbose,
		fields:    ExampleItemFields(),
		defFields: ExampleItemContains(),
		limit:     1000,
		offset:    0,
		sort:      "id",
		order:     "ASC",
	}
	p.table = p.TableName()
	p.field = p.DefaultFields()
	return &p
}

// SetDB setting database
func (p *QueryExampleItem) SetDB(db *sqlx.DB) {
	p.DB = db
}

// Close closes the database, releasing any open resources.
func (p *QueryExampleItem) Close() error {
	return p.DB.Close()
}

// TableName overrides the default tablename generated by QueryExampleItem
func (p *QueryExampleItem) TableName() string {
	return "example_items"
}

// DefaultFields all fields generated by QueryExampleItem
func (p *QueryExampleItem) DefaultFields() string {
	return "id, created_at, updated_at, deleted_at, name, move_id"
}

// DefaultFields all fields generated by QueryExampleItem
func ExampleItemFields() []string {
	return []string{"id", "created_at", "updated_at", "deleted_at", "name", "move_id"}
}

func ExampleItemContains() map[string]interface{} {
	set := make(map[string]interface{})
	set["id"] = "id"
	set["created_at"] = "created_at"
	set["updated_at"] = "updated_at"
	set["deleted_at"] = "deleted_at"
	set["name"] = "name"
	set["move_id"] = "move_id"
	return set
}

func (p *QueryExampleItem) ContainsField(item string) bool {
	_, ok := p.defFields[item]
	return ok
}

func (p *QueryExampleItem) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), p)
}
func (p *QueryExampleItem) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// BuildOneQuery - return one row
func (p *QueryExampleItem) BuildOneQuery(in *Query, field string) (string, []interface{}) {
	return p.BuildQuery(in, field, false, true)
}

// BuildMultiQuery - return rows
func (p *QueryExampleItem) BuildMultiQuery(in *Query, field string) (string, []interface{}) {
	return p.BuildQuery(in, field, false, false)
}

func (p *QueryExampleItem) Count(in *Query) int64 {
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
	return result.Total
}

func (p *QueryExampleItem) BuildQuery(in *Query, f string, reserved bool, one bool) (string, []interface{}) {
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
	return fmt.Sprintf("%s ORDER BY %s %s LIMIT %d OFFSET %d", str, p.sort, order, limit, offset), filterValue
}

// ConvertExampleItemToType takes a pb object and returns an orm object
func ConvertExampleItemToType(from ExampleItem) (ExampleItemType, error) {
	to := ExampleItemType{}
	var err error
	to.ID = from.Id
	to.CreatedAt = from.CreatedAt
	to.UpdatedAt = from.UpdatedAt
	to.DeletedAt = from.DeletedAt
	to.Name = from.Name
	// Skipping field: Mess
	to.MoveId = from.MoveId
	return to, err
}

// ConvertExampleItemFromType takes an orm object and returns a pb object
func ConvertExampleItemFromType(from ExampleItemType) (ExampleItem, error) {
	to := ExampleItem{}
	var err error
	to.Id = from.ID
	to.CreatedAt = from.CreatedAt
	to.UpdatedAt = from.UpdatedAt
	to.DeletedAt = from.DeletedAt
	to.Name = from.Name
	// Skipping field: Mess
	to.MoveId = from.MoveId
	return to, err
}

// ExampleType no comment was provided for message type
type ExampleType struct {
	ID        int64
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
	Name      string `db:"name"`
	// Skipping field: Mess
	Product      string
	Organization string `db:"organization_id"`
	State        string
	Items        []*ExampleItem
}
type QueryExample struct {
	Verbose     int
	DB          *sqlx.DB
	driver      string
	debug       bool
	created     bool
	table       string
	field       string
	fields      []string
	defFields   map[string]interface{}
	query       string
	filter      string
	filterValue []interface{}
	sort        string
	order       string
	limit       int64
	offset      int64
	current     int64
	total       int64
}

// NewQueryExample - initialize QueryExample
func NewQueryExample(verbose int) *QueryExample {
	p := QueryExample{
		Verbose:   verbose,
		fields:    ExampleFields(),
		defFields: ExampleContains(),
		limit:     1000,
		offset:    0,
		sort:      "id",
		order:     "ASC",
	}
	p.table = p.TableName()
	p.field = p.DefaultFields()
	return &p
}

// SetDB setting database
func (p *QueryExample) SetDB(db *sqlx.DB) {
	p.DB = db
}

// Close closes the database, releasing any open resources.
func (p *QueryExample) Close() error {
	return p.DB.Close()
}

// TableName overrides the default tablename generated by QueryExample
func (p *QueryExample) TableName() string {
	return "examples"
}

// DefaultFields all fields generated by QueryExample
func (p *QueryExample) DefaultFields() string {
	return "id, created_at, updated_at, deleted_at, name, product, organization_id AS organization, state, items"
}

// DefaultFields all fields generated by QueryExample
func ExampleFields() []string {
	return []string{"id", "created_at", "updated_at", "deleted_at", "name", "product", "organization_id AS organization", "state", "items"}
}

func ExampleContains() map[string]interface{} {
	set := make(map[string]interface{})
	set["id"] = "id"
	set["created_at"] = "created_at"
	set["updated_at"] = "updated_at"
	set["deleted_at"] = "deleted_at"
	set["name"] = "name"
	set["product"] = "product"
	set["organization_id"] = "organization_id AS organization"
	set["state"] = "state"
	set["items"] = "items"
	return set
}

func (p *QueryExample) ContainsField(item string) bool {
	_, ok := p.defFields[item]
	return ok
}

func (p *QueryExample) Scan(val interface{}) error {
	return json.Unmarshal(val.([]byte), p)
}
func (p *QueryExample) Value() (driver.Value, error) {
	return json.Marshal(p)
}

// BuildOneQuery - return one row
func (p *QueryExample) BuildOneQuery(in *Query, field string) (string, []interface{}) {
	return p.BuildQuery(in, field, false, true)
}

// BuildMultiQuery - return rows
func (p *QueryExample) BuildMultiQuery(in *Query, field string) (string, []interface{}) {
	return p.BuildQuery(in, field, false, false)
}

func (p *QueryExample) Count(in *Query) int64 {
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
	return result.Total
}

func (p *QueryExample) BuildQuery(in *Query, f string, reserved bool, one bool) (string, []interface{}) {
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
	return fmt.Sprintf("%s ORDER BY %s %s LIMIT %d OFFSET %d", str, p.sort, order, limit, offset), filterValue
}

// ConvertExampleToType takes a pb object and returns an orm object
func ConvertExampleToType(from Example) (ExampleType, error) {
	to := ExampleType{}
	var err error
	to.ID = from.Id
	to.CreatedAt = from.CreatedAt
	to.UpdatedAt = from.UpdatedAt
	to.DeletedAt = from.DeletedAt
	to.Name = from.Name
	// Skipping field: Mess
	to.Product = from.Product
	to.Organization = from.Organization
	to.State = from.State
	for _, v := range from.Items {
		if v != nil {
			if tempItems, cErr := ConvertExampleItemToType(*v); cErr == nil {
				to.Items = append(to.Items, &tempItems)
			} else {
				return to, cErr
			}
		} else {
			to.Items = append(to.Items, nil)
		}
	}
	return to, err
}

// ConvertExampleFromType takes an orm object and returns a pb object
func ConvertExampleFromType(from ExampleType) (Example, error) {
	to := Example{}
	var err error
	to.Id = from.ID
	to.CreatedAt = from.CreatedAt
	to.UpdatedAt = from.UpdatedAt
	to.DeletedAt = from.DeletedAt
	to.Name = from.Name
	// Skipping field: Mess
	to.Product = from.Product
	to.Organization = from.Organization
	to.State = from.State
	for _, v := range from.Items {
		if v != nil {
			if tempItems, cErr := ConvertExampleItemFromType(*v); cErr == nil {
				to.Items = append(to.Items, &tempItems)
			} else {
				return to, cErr
			}
		} else {
			to.Items = append(to.Items, nil)
		}
	}
	return to, err
}

////////////////////////// CURDL for objects
func applyField(fields []string) string {
	field := ""
	for _, v := range fields {
		if len(field) == 0 {
			field = v
		} else {
			field = fmt.Sprintf("%s, %s", field, v)
		}
	}
	return field
}

func applyFiltering(filtering []*Filtering) (string, []interface{}) {
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
}
