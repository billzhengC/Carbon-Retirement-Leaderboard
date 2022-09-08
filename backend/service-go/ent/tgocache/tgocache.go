// Code generated by ent, DO NOT EDIT.

package tgocache

import (
	"time"
)

const (
	// Label holds the string label denoting the tgocache type in the database.
	Label = "tgo_cache"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCacheKey holds the string denoting the cache_key field in the database.
	FieldCacheKey = "cache_key"
	// FieldCacheValue holds the string denoting the cache_value field in the database.
	FieldCacheValue = "cache_value"
	// FieldMtime holds the string denoting the mtime field in the database.
	FieldMtime = "mtime"
	// FieldCtime holds the string denoting the ctime field in the database.
	FieldCtime = "ctime"
	// Table holds the table name of the tgocache in the database.
	Table = "t_go_caches"
)

// Columns holds all SQL columns for tgocache fields.
var Columns = []string{
	FieldID,
	FieldCacheKey,
	FieldCacheValue,
	FieldMtime,
	FieldCtime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCacheValue holds the default value on creation for the "cache_value" field.
	DefaultCacheValue string
	// DefaultMtime holds the default value on creation for the "mtime" field.
	DefaultMtime time.Time
	// UpdateDefaultMtime holds the default value on update for the "mtime" field.
	UpdateDefaultMtime func() time.Time
	// DefaultCtime holds the default value on creation for the "ctime" field.
	DefaultCtime time.Time
)