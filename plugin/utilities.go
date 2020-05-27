package plugin

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
	"github.com/weiwolves/protoc-gen-sqlx/pb/sql"
)

// retrieves the SqlxMessageOptions from a message
func getMessageOptions(message *generator.Descriptor) *sql.SqlxMessageOptions {
	if message.Options == nil {
		return nil
	}
	v, err := proto.GetExtension(message.Options, sql.E_Opts)
	if err != nil {
		return nil
	}
	opts, ok := v.(*sql.SqlxMessageOptions)
	if !ok {
		return nil
	}
	return opts
}
