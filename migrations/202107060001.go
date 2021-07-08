package migrations

import (
	"fmt"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/guregu/null"
	"gorm.io/gorm"
)

func init() {
	Migrations = append(Migrations, &gormigrate.Migration{
		ID: "202107060001",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(
				&LogsMigrations{},
				&LogsFoundMigrations{},
				&MatchesMigrations{},
				&SchemaMigrations{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("people")
		},
	})
}

// SchemaMigrations struct is a row record of the schema_migrations table in the main database
type SchemaMigrations struct {
	Version int64 `gorm:"primary_key;column:version;type:INT8;" json:"version"`
	Dirty   bool  `gorm:"column:dirty;type:BOOL;" json:"dirty"`
}

var schema_migrationsTableInfo = &TableInfo{
	Name: "schema_migrations",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "version",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "INT8",
			DatabaseTypePretty: "INT8",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "INT8",
			ColumnLength:       -1,
			GoFieldName:        "Version",
			GoFieldType:        "int64",
			JSONFieldName:      "version",
			ProtobufFieldName:  "version",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "dirty",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "BOOL",
			DatabaseTypePretty: "BOOL",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "BOOL",
			ColumnLength:       -1,
			GoFieldName:        "Dirty",
			GoFieldType:        "bool",
			JSONFieldName:      "dirty",
			ProtobufFieldName:  "dirty",
			ProtobufType:       "bool",
			ProtobufPos:        2,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SchemaMigrations) TableName() string {
	return "schema_migrations"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SchemaMigrations) BeforeSave(*gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SchemaMigrations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SchemaMigrations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SchemaMigrations) TableInfo() *TableInfo {
	return schema_migrationsTableInfo
}

// LogsMigrations struct is a row record of the schema_migrations table in the main database
type LogsMigrations struct {
	ID             string    `gorm:"primary_key;column:id;type:VARCHAR;" json:"id"`
	LogLine        string    `gorm:"column:log_line;type:VARCHAR;" json:"log_line"`
	LastError      null.Time `gorm:"column:last_error;type:DATETIME;" json:"last_error"`
	UpdatedAt      null.Time `gorm:"column:updated_at;type:DATETIME;" json:"updated_at"`
	MatchingString string    `gorm:"column:matching_string;type:VARCHAR;" json:"matching_string"`
}

var logsTableInfo = &TableInfo{
	Name: "logs",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "log_line",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "LogLine",
			GoFieldType:        "string",
			JSONFieldName:      "log_line",
			ProtobufFieldName:  "log_line",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "last_error",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "DATETIME",
			DatabaseTypePretty: "DATETIME",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "DATETIME",
			ColumnLength:       -1,
			GoFieldName:        "LastError",
			GoFieldType:        "null.Time",
			JSONFieldName:      "last_error",
			ProtobufFieldName:  "last_error",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "updated_at",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "DATETIME",
			DatabaseTypePretty: "DATETIME",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "DATETIME",
			ColumnLength:       -1,
			GoFieldName:        "UpdateAt",
			GoFieldType:        "null.Time",
			JSONFieldName:      "updated_at",
			ProtobufFieldName:  "updated_at",
			ProtobufType:       "uint64",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "matching_string",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "MatchingString",
			GoFieldType:        "string",
			JSONFieldName:      "matching_string",
			ProtobufFieldName:  "matching_string",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *LogsMigrations) TableName() string {
	return "logs"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *LogsMigrations) BeforeSave(*gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *LogsMigrations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *LogsMigrations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *LogsMigrations) TableInfo() *TableInfo {
	return logsTableInfo
}

// LogsFoundMigrations struct is a row record of the schema_migrations table in the main database
type LogsFoundMigrations struct {
	LogsID    string    `gorm:"column:logs_id;type:VARCHAR;" json:"logs_id"`
	TimeStart null.Time `gorm:"column:time_start;type:DATETIME;" json:"time_start"`
	TimeEnd   null.Time `gorm:"column:time_end;type:DATETIME;" json:"time_end"`
}

var logsFoundTableInfo = &TableInfo{
	Name: "logs_found",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "logs_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "LogsID",
			GoFieldType:        "string",
			JSONFieldName:      "logs_id",
			ProtobufFieldName:  "logs_id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "time_start",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "DATETIME",
			DatabaseTypePretty: "DATETIME",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "DATETIME",
			ColumnLength:       -1,
			GoFieldName:        "TimeStart",
			GoFieldType:        "null.Time",
			JSONFieldName:      "time_start",
			ProtobufFieldName:  "time_start",
			ProtobufType:       "uint64",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "time_end",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "DATETIME",
			DatabaseTypePretty: "DATETIME",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "DATETIME",
			ColumnLength:       -1,
			GoFieldName:        "TimeEnd",
			GoFieldType:        "null.Time",
			JSONFieldName:      "time_end",
			ProtobufFieldName:  "time_end",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *LogsFoundMigrations) TableName() string {
	return "logs_found"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *LogsFoundMigrations) BeforeSave(*gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *LogsFoundMigrations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *LogsFoundMigrations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *LogsFoundMigrations) TableInfo() *TableInfo {
	return logsFoundTableInfo
}

// LogsFoundMigrations struct is a row record of the schema_migrations table in the main database
type MatchesMigrations struct {
	ID             string `gorm:"primary_key;column:id;type:VARCHAR;" json:"id"`
	MatchingString string `gorm:"column:matching_string;type:VARCHAR;" json:"matching_string"`
	Apps           string `gorm:"column:apps;type:VARCHAR;" json:"apps"`
	Name           string `gorm:"column:name;type:VARCHAR;" json:"name"`
	Description    string `gorm:"column:description;type:VARCHAR;" json:"description"`
}

var matchesTableInfo = &TableInfo{
	Name: "matches",
	Columns: []*ColumnInfo{

		{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "string",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		{
			Index:              1,
			Name:               "matching_string",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "MatchingString",
			GoFieldType:        "string",
			JSONFieldName:      "matching_string",
			ProtobufFieldName:  "matching_string",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		{
			Index:              2,
			Name:               "apps",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "Apps",
			GoFieldType:        "string",
			JSONFieldName:      "apps",
			ProtobufFieldName:  "apps",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		{
			Index:              3,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		{
			Index:              4,
			Name:               "description",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "VARCHAR",
			DatabaseTypePretty: "VARCHAR",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "VARCHAR",
			ColumnLength:       -1,
			GoFieldName:        "Description",
			GoFieldType:        "string",
			JSONFieldName:      "description",
			ProtobufFieldName:  "description",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MatchesMigrations) TableName() string {
	return "matches"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MatchesMigrations) BeforeSave(*gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MatchesMigrations) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MatchesMigrations) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MatchesMigrations) TableInfo() *TableInfo {
	return matchesTableInfo
}

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)

	tables map[string]*TableInfo
)

func init() {
	tables = make(map[string]*TableInfo)

	tables["logs"] = logsTableInfo
	tables["logs_found"] = logsFoundTableInfo
	tables["matches"] = matchesTableInfo
	tables["schema_migrations"] = schema_migrationsTableInfo
}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave(*gorm.DB) error
	Prepare()
	Validate(action Action) error
	TableInfo() *TableInfo
}

// TableInfo describes a table in the database
type TableInfo struct {
	Name    string        `json:"name"`
	Columns []*ColumnInfo `json:"columns"`
}

// ColumnInfo describes a column in the database table
type ColumnInfo struct {
	Index              int    `json:"index"`
	GoFieldName        string `json:"go_field_name"`
	GoFieldType        string `json:"go_field_type"`
	JSONFieldName      string `json:"json_field_name"`
	ProtobufFieldName  string `json:"protobuf_field_name"`
	ProtobufType       string `json:"protobuf_field_type"`
	ProtobufPos        int    `json:"protobuf_field_pos"`
	Comment            string `json:"comment"`
	Notes              string `json:"notes"`
	Name               string `json:"name"`
	Nullable           bool   `json:"is_nullable"`
	DatabaseTypeName   string `json:"database_type_name"`
	DatabaseTypePretty string `json:"database_type_pretty"`
	IsPrimaryKey       bool   `json:"is_primary_key"`
	IsAutoIncrement    bool   `json:"is_auto_increment"`
	IsArray            bool   `json:"is_array"`
	ColumnType         string `json:"column_type"`
	ColumnLength       int64  `json:"column_length"`
	DefaultValue       string `json:"default_value"`
}

// GetTableInfo retrieve TableInfo for a table
func GetTableInfo(name string) (*TableInfo, bool) {
	val, ok := tables[name]
	return val, ok
}
