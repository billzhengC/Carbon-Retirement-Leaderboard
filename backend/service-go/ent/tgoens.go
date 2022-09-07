// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"toucan-leaderboard/ent/tgoens"

	"entgo.io/ent/dialect/sql"
)

// TGoEns is the model entity for the TGoEns schema.
type TGoEns struct {
	config `json:"-"`
	// ID of the ent.
	// auto increment primary key
	ID uint64 `json:"id,omitempty"`
	// wallet public key
	WalletPub string `json:"wallet_pub,omitempty"`
	// creator of the certificate
	Ens string `json:"ens,omitempty"`
	// modify time
	Mtime time.Time `json:"mtime,omitempty"`
	// create time
	Ctime time.Time `json:"ctime,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TGoEns) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case tgoens.FieldID:
			values[i] = new(sql.NullInt64)
		case tgoens.FieldWalletPub, tgoens.FieldEns:
			values[i] = new(sql.NullString)
		case tgoens.FieldMtime, tgoens.FieldCtime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TGoEns", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TGoEns fields.
func (te *TGoEns) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case tgoens.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			te.ID = uint64(value.Int64)
		case tgoens.FieldWalletPub:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wallet_pub", values[i])
			} else if value.Valid {
				te.WalletPub = value.String
			}
		case tgoens.FieldEns:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ens", values[i])
			} else if value.Valid {
				te.Ens = value.String
			}
		case tgoens.FieldMtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field mtime", values[i])
			} else if value.Valid {
				te.Mtime = value.Time
			}
		case tgoens.FieldCtime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ctime", values[i])
			} else if value.Valid {
				te.Ctime = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TGoEns.
// Note that you need to call TGoEns.Unwrap() before calling this method if this TGoEns
// was returned from a transaction, and the transaction was committed or rolled back.
func (te *TGoEns) Update() *TGoEnsUpdateOne {
	return (&TGoEnsClient{config: te.config}).UpdateOne(te)
}

// Unwrap unwraps the TGoEns entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (te *TGoEns) Unwrap() *TGoEns {
	_tx, ok := te.config.driver.(*txDriver)
	if !ok {
		panic("ent: TGoEns is not a transactional entity")
	}
	te.config.driver = _tx.drv
	return te
}

// String implements the fmt.Stringer.
func (te *TGoEns) String() string {
	var builder strings.Builder
	builder.WriteString("TGoEns(")
	builder.WriteString(fmt.Sprintf("id=%v, ", te.ID))
	builder.WriteString("wallet_pub=")
	builder.WriteString(te.WalletPub)
	builder.WriteString(", ")
	builder.WriteString("ens=")
	builder.WriteString(te.Ens)
	builder.WriteString(", ")
	builder.WriteString("mtime=")
	builder.WriteString(te.Mtime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("ctime=")
	builder.WriteString(te.Ctime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// TGoEnsSlice is a parsable slice of TGoEns.
type TGoEnsSlice []*TGoEns

func (te TGoEnsSlice) config(cfg config) {
	for _i := range te {
		te[_i].config = cfg
	}
}
