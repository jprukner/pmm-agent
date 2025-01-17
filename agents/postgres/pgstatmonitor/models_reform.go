// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package pgstatmonitor

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type pgStatDatabaseViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("pg_catalog").
func (v *pgStatDatabaseViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("pg_stat_database").
func (v *pgStatDatabaseViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *pgStatDatabaseViewType) Columns() []string {
	return []string{
		"datid",
		"datname",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *pgStatDatabaseViewType) NewStruct() reform.Struct {
	return new(pgStatDatabase)
}

// pgStatDatabaseView represents pg_stat_database view or table in SQL database.
var pgStatDatabaseView = &pgStatDatabaseViewType{
	s: parse.StructInfo{
		Type:      "pgStatDatabase",
		SQLSchema: "pg_catalog",
		SQLName:   "pg_stat_database",
		Fields: []parse.FieldInfo{
			{Name: "DatID", Type: "int64", Column: "datid"},
			{Name: "DatName", Type: "*string", Column: "datname"},
		},
		PKFieldIndex: -1,
	},
	z: new(pgStatDatabase).Values(),
}

// String returns a string representation of this struct or record.
func (s pgStatDatabase) String() string {
	res := make([]string, 2)
	res[0] = "DatID: " + reform.Inspect(s.DatID, true)
	res[1] = "DatName: " + reform.Inspect(s.DatName, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *pgStatDatabase) Values() []interface{} {
	return []interface{}{
		s.DatID,
		s.DatName,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *pgStatDatabase) Pointers() []interface{} {
	return []interface{}{
		&s.DatID,
		&s.DatName,
	}
}

// View returns View object for that struct.
func (s *pgStatDatabase) View() reform.View {
	return pgStatDatabaseView
}

// check interfaces
var (
	_ reform.View   = pgStatDatabaseView
	_ reform.Struct = (*pgStatDatabase)(nil)
	_ fmt.Stringer  = (*pgStatDatabase)(nil)
)

type pgUserViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("pg_catalog").
func (v *pgUserViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("pg_user").
func (v *pgUserViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *pgUserViewType) Columns() []string {
	return []string{
		"usesysid",
		"usename",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *pgUserViewType) NewStruct() reform.Struct {
	return new(pgUser)
}

// pgUserView represents pg_user view or table in SQL database.
var pgUserView = &pgUserViewType{
	s: parse.StructInfo{
		Type:      "pgUser",
		SQLSchema: "pg_catalog",
		SQLName:   "pg_user",
		Fields: []parse.FieldInfo{
			{Name: "UserID", Type: "int64", Column: "usesysid"},
			{Name: "UserName", Type: "*string", Column: "usename"},
		},
		PKFieldIndex: -1,
	},
	z: new(pgUser).Values(),
}

// String returns a string representation of this struct or record.
func (s pgUser) String() string {
	res := make([]string, 2)
	res[0] = "UserID: " + reform.Inspect(s.UserID, true)
	res[1] = "UserName: " + reform.Inspect(s.UserName, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *pgUser) Values() []interface{} {
	return []interface{}{
		s.UserID,
		s.UserName,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *pgUser) Pointers() []interface{} {
	return []interface{}{
		&s.UserID,
		&s.UserName,
	}
}

// View returns View object for that struct.
func (s *pgUser) View() reform.View {
	return pgUserView
}

// check interfaces
var (
	_ reform.View   = pgUserView
	_ reform.Struct = (*pgUser)(nil)
	_ fmt.Stringer  = (*pgUser)(nil)
)

type pgStatMonitorViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *pgStatMonitorViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("pg_stat_monitor").
func (v *pgStatMonitorViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *pgStatMonitorViewType) Columns() []string {
	return []string{
		"bucket",
		"bucket_start_time",
		"userid",
		"dbid",
		"queryid",
		"query",
		"total_calls",
		"total_time",
		"effected_rows",
		"shared_blks_hit",
		"shared_blks_read",
		"shared_blks_dirtied",
		"shared_blks_written",
		"local_blks_hit",
		"local_blks_read",
		"local_blks_dirtied",
		"local_blks_written",
		"temp_blks_read",
		"temp_blks_written",
		"blk_read_time",
		"blk_write_time",
		"client_ip",
		"resp_calls",
		"cpu_user_time",
		"cpu_sys_time",
		"tables_names",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *pgStatMonitorViewType) NewStruct() reform.Struct {
	return new(pgStatMonitor)
}

// pgStatMonitorView represents pg_stat_monitor view or table in SQL database.
var pgStatMonitorView = &pgStatMonitorViewType{
	s: parse.StructInfo{
		Type:    "pgStatMonitor",
		SQLName: "pg_stat_monitor",
		Fields: []parse.FieldInfo{
			{Name: "Bucket", Type: "int64", Column: "bucket"},
			{Name: "BucketStartTime", Type: "time.Time", Column: "bucket_start_time"},
			{Name: "UserID", Type: "int64", Column: "userid"},
			{Name: "DBID", Type: "int64", Column: "dbid"},
			{Name: "QueryID", Type: "string", Column: "queryid"},
			{Name: "Query", Type: "string", Column: "query"},
			{Name: "TotalCalls", Type: "int64", Column: "total_calls"},
			{Name: "TotalTime", Type: "float64", Column: "total_time"},
			{Name: "EffectedRows", Type: "int64", Column: "effected_rows"},
			{Name: "SharedBlksHit", Type: "int64", Column: "shared_blks_hit"},
			{Name: "SharedBlksRead", Type: "int64", Column: "shared_blks_read"},
			{Name: "SharedBlksDirtied", Type: "int64", Column: "shared_blks_dirtied"},
			{Name: "SharedBlksWritten", Type: "int64", Column: "shared_blks_written"},
			{Name: "LocalBlksHit", Type: "int64", Column: "local_blks_hit"},
			{Name: "LocalBlksRead", Type: "int64", Column: "local_blks_read"},
			{Name: "LocalBlksDirtied", Type: "int64", Column: "local_blks_dirtied"},
			{Name: "LocalBlksWritten", Type: "int64", Column: "local_blks_written"},
			{Name: "TempBlksRead", Type: "int64", Column: "temp_blks_read"},
			{Name: "TempBlksWritten", Type: "int64", Column: "temp_blks_written"},
			{Name: "BlkReadTime", Type: "float64", Column: "blk_read_time"},
			{Name: "BlkWriteTime", Type: "float64", Column: "blk_write_time"},
			{Name: "ClientIP", Type: "string", Column: "client_ip"},
			{Name: "RespCalls", Type: "pq.StringArray", Column: "resp_calls"},
			{Name: "CPUUserTime", Type: "float64", Column: "cpu_user_time"},
			{Name: "CPUSysTime", Type: "float64", Column: "cpu_sys_time"},
			{Name: "TablesNames", Type: "pq.StringArray", Column: "tables_names"},
		},
		PKFieldIndex: -1,
	},
	z: new(pgStatMonitor).Values(),
}

// String returns a string representation of this struct or record.
func (s pgStatMonitor) String() string {
	res := make([]string, 26)
	res[0] = "Bucket: " + reform.Inspect(s.Bucket, true)
	res[1] = "BucketStartTime: " + reform.Inspect(s.BucketStartTime, true)
	res[2] = "UserID: " + reform.Inspect(s.UserID, true)
	res[3] = "DBID: " + reform.Inspect(s.DBID, true)
	res[4] = "QueryID: " + reform.Inspect(s.QueryID, true)
	res[5] = "Query: " + reform.Inspect(s.Query, true)
	res[6] = "TotalCalls: " + reform.Inspect(s.TotalCalls, true)
	res[7] = "TotalTime: " + reform.Inspect(s.TotalTime, true)
	res[8] = "EffectedRows: " + reform.Inspect(s.EffectedRows, true)
	res[9] = "SharedBlksHit: " + reform.Inspect(s.SharedBlksHit, true)
	res[10] = "SharedBlksRead: " + reform.Inspect(s.SharedBlksRead, true)
	res[11] = "SharedBlksDirtied: " + reform.Inspect(s.SharedBlksDirtied, true)
	res[12] = "SharedBlksWritten: " + reform.Inspect(s.SharedBlksWritten, true)
	res[13] = "LocalBlksHit: " + reform.Inspect(s.LocalBlksHit, true)
	res[14] = "LocalBlksRead: " + reform.Inspect(s.LocalBlksRead, true)
	res[15] = "LocalBlksDirtied: " + reform.Inspect(s.LocalBlksDirtied, true)
	res[16] = "LocalBlksWritten: " + reform.Inspect(s.LocalBlksWritten, true)
	res[17] = "TempBlksRead: " + reform.Inspect(s.TempBlksRead, true)
	res[18] = "TempBlksWritten: " + reform.Inspect(s.TempBlksWritten, true)
	res[19] = "BlkReadTime: " + reform.Inspect(s.BlkReadTime, true)
	res[20] = "BlkWriteTime: " + reform.Inspect(s.BlkWriteTime, true)
	res[21] = "ClientIP: " + reform.Inspect(s.ClientIP, true)
	res[22] = "RespCalls: " + reform.Inspect(s.RespCalls, true)
	res[23] = "CPUUserTime: " + reform.Inspect(s.CPUUserTime, true)
	res[24] = "CPUSysTime: " + reform.Inspect(s.CPUSysTime, true)
	res[25] = "TablesNames: " + reform.Inspect(s.TablesNames, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *pgStatMonitor) Values() []interface{} {
	return []interface{}{
		s.Bucket,
		s.BucketStartTime,
		s.UserID,
		s.DBID,
		s.QueryID,
		s.Query,
		s.TotalCalls,
		s.TotalTime,
		s.EffectedRows,
		s.SharedBlksHit,
		s.SharedBlksRead,
		s.SharedBlksDirtied,
		s.SharedBlksWritten,
		s.LocalBlksHit,
		s.LocalBlksRead,
		s.LocalBlksDirtied,
		s.LocalBlksWritten,
		s.TempBlksRead,
		s.TempBlksWritten,
		s.BlkReadTime,
		s.BlkWriteTime,
		s.ClientIP,
		s.RespCalls,
		s.CPUUserTime,
		s.CPUSysTime,
		s.TablesNames,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *pgStatMonitor) Pointers() []interface{} {
	return []interface{}{
		&s.Bucket,
		&s.BucketStartTime,
		&s.UserID,
		&s.DBID,
		&s.QueryID,
		&s.Query,
		&s.TotalCalls,
		&s.TotalTime,
		&s.EffectedRows,
		&s.SharedBlksHit,
		&s.SharedBlksRead,
		&s.SharedBlksDirtied,
		&s.SharedBlksWritten,
		&s.LocalBlksHit,
		&s.LocalBlksRead,
		&s.LocalBlksDirtied,
		&s.LocalBlksWritten,
		&s.TempBlksRead,
		&s.TempBlksWritten,
		&s.BlkReadTime,
		&s.BlkWriteTime,
		&s.ClientIP,
		&s.RespCalls,
		&s.CPUUserTime,
		&s.CPUSysTime,
		&s.TablesNames,
	}
}

// View returns View object for that struct.
func (s *pgStatMonitor) View() reform.View {
	return pgStatMonitorView
}

// check interfaces
var (
	_ reform.View   = pgStatMonitorView
	_ reform.Struct = (*pgStatMonitor)(nil)
	_ fmt.Stringer  = (*pgStatMonitor)(nil)
)

type pgStatMonitorSettingsViewType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *pgStatMonitorSettingsViewType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("pg_stat_monitor_settings").
func (v *pgStatMonitorSettingsViewType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *pgStatMonitorSettingsViewType) Columns() []string {
	return []string{
		"name",
		"value",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *pgStatMonitorSettingsViewType) NewStruct() reform.Struct {
	return new(pgStatMonitorSettings)
}

// pgStatMonitorSettingsView represents pg_stat_monitor_settings view or table in SQL database.
var pgStatMonitorSettingsView = &pgStatMonitorSettingsViewType{
	s: parse.StructInfo{
		Type:    "pgStatMonitorSettings",
		SQLName: "pg_stat_monitor_settings",
		Fields: []parse.FieldInfo{
			{Name: "Name", Type: "string", Column: "name"},
			{Name: "Value", Type: "int64", Column: "value"},
		},
		PKFieldIndex: -1,
	},
	z: new(pgStatMonitorSettings).Values(),
}

// String returns a string representation of this struct or record.
func (s pgStatMonitorSettings) String() string {
	res := make([]string, 2)
	res[0] = "Name: " + reform.Inspect(s.Name, true)
	res[1] = "Value: " + reform.Inspect(s.Value, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *pgStatMonitorSettings) Values() []interface{} {
	return []interface{}{
		s.Name,
		s.Value,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *pgStatMonitorSettings) Pointers() []interface{} {
	return []interface{}{
		&s.Name,
		&s.Value,
	}
}

// View returns View object for that struct.
func (s *pgStatMonitorSettings) View() reform.View {
	return pgStatMonitorSettingsView
}

// check interfaces
var (
	_ reform.View   = pgStatMonitorSettingsView
	_ reform.Struct = (*pgStatMonitorSettings)(nil)
	_ fmt.Stringer  = (*pgStatMonitorSettings)(nil)
)

func init() {
	parse.AssertUpToDate(&pgStatDatabaseView.s, new(pgStatDatabase))
	parse.AssertUpToDate(&pgUserView.s, new(pgUser))
	parse.AssertUpToDate(&pgStatMonitorView.s, new(pgStatMonitor))
	parse.AssertUpToDate(&pgStatMonitorSettingsView.s, new(pgStatMonitorSettings))
}
