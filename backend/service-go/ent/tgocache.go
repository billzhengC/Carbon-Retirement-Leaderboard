// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"toucan-leaderboard/ent/tgocache"

	"entgo.io/ent/dialect/sql"
)

// TGoCache is the model entity for the TGoCache schema.
type TGoCache struct {
	config `json:"-"`
	// ID of the ent.
	// auto increment primary key
	ID uint64 `json:"id,omitempty"`
	// cache key
	CacheKey string `json:"cache_key,omitempty"`
	// cached value
	CacheValue string `json:"cache_value,omitempty"`
	// modify time
	Mtime time.Time `json:"mtime,omitempty"`
	// create time
	Ctime time.Time `json:"ctime,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TGoCache) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tgocache.FieldID:
			values[i] = new(sql.NullInt64)
		case tgocache.FieldCacheKey, tgocache.FieldCacheValue:
			values[i] = new(sql.NullString)
		case tgocache.FieldMtime, tgocache.FieldCtime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TGoCache", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TGoCache fields.
func (tc *TGoCache) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tgocache.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tc.ID = uint64(value.Int64)
		case tgocache.FieldCacheKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cache_key", values[i])
			} else if value.Valid {
				tc.CacheKey = value.String
			}
		case tgocache.FieldCacheValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cache_value", values[i])
			} else if value.Valid {
				tc.CacheValue = value.String
			}
		case tgocache.FieldMtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mtime", values[i])
			} else if value.Valid {
				tc.Mtime = value.Time
			}
		case tgocache.FieldCtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ctime", values[i])
			} else if value.Valid {
				tc.Ctime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TGoCache.
// Note that you need to call TGoCache.Unwrap() before calling this method if this TGoCache
// was returned from a transaction, and the transaction was committed or rolled back.
func (tc *TGoCache) Update() *TGoCacheUpdateOne {
	return (&TGoCacheClient{config: tc.config}).UpdateOne(tc)
}

// Unwrap unwraps the TGoCache entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tc *TGoCache) Unwrap() *TGoCache {
	_tx, ok := tc.config.driver.(*txDriver)
	if !ok {
		panic("ent: TGoCache is not a transactional entity")
	}
	tc.config.driver = _tx.drv
	return tc
}

// String implements the fmt.Stringer.
func (tc *TGoCache) String() string {
	var builder strings.Builder
	builder.WriteString("TGoCache(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tc.ID))
	builder.WriteString("cache_key=")
	builder.WriteString(tc.CacheKey)
	builder.WriteString(", ")
	builder.WriteString("cache_value=")
	builder.WriteString(tc.CacheValue)
	builder.WriteString(", ")
	builder.WriteString("mtime=")
	builder.WriteString(tc.Mtime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ctime=")
	builder.WriteString(tc.Ctime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TGoCaches is a parsable slice of TGoCache.
type TGoCaches []*TGoCache

func (tc TGoCaches) config(cfg config) {
	for _i := range tc {
		tc[_i].config = cfg
	}
}
