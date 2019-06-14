package sqlx

import (
	"fmt"
	"strings"
)

func strjoin(a []string, sep string, extra ...string) string {
	out := strings.Join(a, sep)
	switch len(extra) {
	case 1:
		return extra[0] + out
	case 2:
		return extra[0] + out + extra[1]
	default:
		return out
	}
}

// JoinType ...
type JoinType uint8

// Join type
const (
	InnerJoin JoinType = 1 + iota
	LeftJoin
	RightJoin
)

func (t JoinType) String() string {
	switch t {
	case InnerJoin:
		return " INNER JOIN "
	case RightJoin:
		return " RIGHT JOIN "
	case LeftJoin:
		return " LEFT JOIN "
	default:
		return " LEFT JOIN "
	}
}

// Direction ...
type Direction uint8

// Sort direction
const (
	ASC Direction = 1 + iota
	DESC
)

func (d Direction) String() string {
	switch d {
	case ASC:
		return " ASC"
	case DESC:
		return " DESC"
	default:
		return ""
	}
}

// Cmd ...
type Cmd struct {
	orderby bool
	from    bool
	fields  int
	values  bool
	sb      *strings.Builder
}

// NewCmd returns a mysql command builder.
func NewCmd() *Cmd {
	return &Cmd{
		orderby: false,
		from:    false,
		fields:  0,
		values:  false,
		sb:      new(strings.Builder),
	}
}

// Reset ...
func (c *Cmd) Reset() *Cmd {
	c.orderby = false
	c.from = false
	c.fields = 0
	c.values = false
	c.sb.Reset()
	return c
}

// Select build select part sql.
func (c *Cmd) Select(fields ...string) *Cmd {
	if c.sb.Len() > 0 {
		c.sb.WriteString(" SELECT ")
	} else {
		c.sb.WriteString("SELECT ")
	}

	if len(fields) > 0 {
		c.sb.WriteString(strjoin(fields, ","))
	} else {
		c.sb.WriteString("*")
	}

	c.from = false
	c.orderby = false
	return c
}

// Into ...
func (c *Cmd) Into(table string) *Cmd {
	c.sb.WriteString(" INTO ")
	c.sb.WriteString(table)
	return c
}

// From ...
func (c *Cmd) From(table ...string) *Cmd {
	if len(table) <= 0 {
		panic("at least one table")
	}

	if !c.from {
		c.sb.WriteString(" FROM ")
		c.from = true
	} else {
		c.sb.WriteByte(',')
	}

	c.sb.WriteString(strjoin(table, ","))
	return c
}

// SubQuery ...
func (c *Cmd) SubQuery(sub *Cmd, as ...string) *Cmd {
	out := "(" + sub.String() + ")"
	if len(as) > 0 {
		out = out + " AS " + as[0]
	}

	return c.From(out)
}

// Union  ...
func (c *Cmd) Union(other ...*Cmd) *Cmd {
	if len(other) > 0 {
		out := c.String()
		c.Reset()
		c.Parentheses(out).sb.WriteString(" UNION ")
		c.Parentheses(other[0].String())
	} else {
		c.sb.WriteString(" UNION")
	}

	return c
}

// Join ...
func (c *Cmd) Join(t JoinType, table ...string) *Cmd {
	size := len(table)
	if size < 0 {
		panic("at least one table")
	}

	c.sb.WriteString(t.String())

	if size == 1 {
		c.sb.WriteString(table[0])
	} else {
		c.sb.WriteString(strjoin(table, ",", "(", ")"))
	}

	return c
}

// InnerJoin ...
func (c *Cmd) InnerJoin(table ...string) *Cmd {
	return c.Join(InnerJoin, table...)
}

// LeftJoin ...
func (c *Cmd) LeftJoin(table ...string) *Cmd {
	return c.Join(LeftJoin, table...)
}

// RightJoin ...
func (c *Cmd) RightJoin(table ...string) *Cmd {
	return c.Join(RightJoin, table...)
}

// On ...
func (c *Cmd) On(condition ...string) *Cmd {
	size := len(condition)
	if size < 0 {
		panic("at least one condition")
	}

	c.sb.WriteString(" ON ")
	if size == 1 {
		c.sb.WriteString(condition[0])
	} else {
		c.sb.WriteString(strjoin(condition, " AND ", "(", ")"))
	}

	return c
}

// JoinOn ...
func (c *Cmd) JoinOn(t JoinType, table string, condition string) *Cmd {
	return c.Join(t, table).On(condition)
}

// LeftJoinOn ...
func (c *Cmd) LeftJoinOn(table string, condition string) *Cmd {
	return c.JoinOn(LeftJoin, table, condition)
}

// RightJoinOn ...
func (c *Cmd) RightJoinOn(table string, condition string) *Cmd {
	return c.JoinOn(RightJoin, table, condition)
}

// InnerJoinOn ...
func (c *Cmd) InnerJoinOn(table string, condition string) *Cmd {
	return c.JoinOn(InnerJoin, table, condition)
}

// Where ...
func (c *Cmd) Where(condition string) *Cmd {
	c.sb.WriteString(" WHERE ")
	c.sb.WriteString(condition)
	return c
}

// And ...
func (c *Cmd) And(condition ...string) *Cmd {
	if len(condition) <= 0 {
		panic("at least one condition")
	}

	c.sb.WriteString(strjoin(condition, " AND ", " AND "))

	return c
}

// Or ...
func (c *Cmd) Or(condition ...string) *Cmd {
	if len(condition) <= 0 {
		panic("at least one condition")
	}

	c.sb.WriteString(strjoin(condition, " OR ", " OR "))

	return c
}

// WhereAnd ...
func (c *Cmd) WhereAnd(condition string, others ...string) *Cmd {
	c.Where(condition)
	if len(others) > 0 {
		c.And(others...)
	}
	return c
}

// WhereOr ...
func (c *Cmd) WhereOr(condition string, others ...string) *Cmd {
	c.Where(condition)
	if len(others) > 0 {
		c.Or(others...)
	}

	return c
}

// Insert ...
func (c *Cmd) Insert(table string, fields ...string) *Cmd {
	c.sb.WriteString("INSERT")
	c.Into(table)
	c.fields = len(fields)
	if c.fields > 0 {
		c.sb.WriteString(strjoin(fields, ",", " (", ")"))
	}
	c.values = false

	return c
}

// InsertValues ...
func (c *Cmd) InsertValues(table string, fields ...string) *Cmd {
	return c.Insert(table, fields...).Values()
}

// Values ...
func (c *Cmd) Values(assignments ...string) *Cmd {
	if !c.values {
		c.sb.WriteString(" VALUES ")
		c.values = true
	} else {
		c.sb.WriteString(",")
	}

	if len(assignments) <= 0 {
		c.sb.WriteByte('(')
		c.sb.WriteString(strings.Repeat("?,", c.fields-1))
		c.sb.WriteString("?)")
	} else {
		c.sb.WriteString(strjoin(assignments, ",", "(", ")"))
	}
	return c
}

// Duplicate ...
func (c *Cmd) Duplicate(assignments ...string) *Cmd {
	c.sb.WriteString(" ON DUPLICATE KEY UPDATE ")
	c.sb.WriteString(strjoin(assignments, ","))
	return c
}

// DuplicateValues ...
func (c *Cmd) DuplicateValues(values ...string) *Cmd {
	out := make([]string, len(values))
	for i, x := range values {
		out[i] = fmt.Sprintf("%s=VALUES(%s)", x, x)
	}
	return c.Duplicate(out...)
}

// SetFields ...
func (c *Cmd) SetFields(fields ...string) *Cmd {
	out := make([]string, len(fields))

	for i, f := range fields {
		out[i] = fmt.Sprintf("%s=?", f)
	}

	return c.Set(out...)
}

// Set ...
func (c *Cmd) Set(assignments ...string) *Cmd {
	c.sb.WriteString(" SET")
	c.sb.WriteString(strjoin(assignments, ",", " "))
	return c
}

// Update ...
func (c *Cmd) Update(table string, others ...string) *Cmd {
	c.sb.WriteString("UPDATE ")
	c.sb.WriteString(table)
	if len(others) > 0 {
		c.sb.WriteString(strjoin(others, ",", ","))
	}

	return c
}

// Delete ...
func (c *Cmd) Delete(as ...string) *Cmd {
	c.sb.WriteString("DELETE")
	if len(as) > 0 {
		c.sb.WriteString(strjoin(as, ",", " "))
	}
	return c
}

// DeleteFrom ...
func (c *Cmd) DeleteFrom(table ...string) *Cmd {
	return c.Delete().From(table...)
}

// Replace ...
func (c *Cmd) Replace(table string, fields ...string) *Cmd {
	c.sb.WriteString("REPLACE")
	c.Into(table)

	c.fields = len(fields)
	c.values = false

	if c.fields > 0 {
		c.sb.WriteString(strjoin(fields, ",", "(", ")"))
	}

	return c
}

// Parentheses ...
func (c *Cmd) Parentheses(any string) *Cmd {
	c.sb.WriteString("(")
	c.sb.WriteString(any)
	c.sb.WriteString(")")
	return c
}

// Limit ...
func (c *Cmd) Limit(count int, offset ...int) *Cmd {

	if len(offset) > 0 {
		c.sb.WriteString(fmt.Sprintf(" LIMIT %d,%d", offset[0], count))
	} else {
		c.sb.WriteString(fmt.Sprintf(" LIMIT %d", count))
	}

	return c
}

// GroupBy ...
func (c *Cmd) GroupBy(fields ...string) *Cmd {
	c.sb.WriteString(" GROUP BY ")
	c.sb.WriteString(strjoin(fields, ","))
	return c
}

// OrderBy ...
func (c *Cmd) OrderBy(field string, direction ...Direction) *Cmd {
	dir := Direction(0)
	if len(direction) > 0 {
		dir = direction[0]
	}

	if !c.orderby {
		c.sb.WriteString(" ORDER BY ")
		c.orderby = true
	} else {
		c.sb.WriteString(", ")
	}

	c.sb.WriteString(field)
	c.sb.WriteString(dir.String())
	return c
}

// Having ...
func (c *Cmd) Having(condition string) *Cmd {
	c.sb.WriteString(" HAVING ")
	c.sb.WriteString(condition)
	return c
}

// Raw appends string to command
func (c *Cmd) Raw(any string) *Cmd {
	c.sb.WriteString(any)
	return c
}

func (c *Cmd) String() string {
	return c.sb.String()
}

// SQL ...
func (c *Cmd) SQL(sqlcmds ...string) string {
	if len(sqlcmds) > 0 {
		c.Reset()
		c.Raw(strings.Join(sqlcmds, ";"))
	}
	return c.String()
}
