package query

type Operator string

const (
	OP_NOT       = "!"
	OP_EQUAL     = "eq"
	OP_NOT_EQUAL = OP_NOT + OP_EQUAL
)

// condition
type Condition struct {
	O Operator
	V interface{}
}

func match(o Operator, v interface{}) Condition {
	return Condition{
		O: o,
		V: v,
	}
}

func Eq(v interface{}) Condition {
	return match(OP_EQUAL, v)
}

func NotEq(v interface{}) Condition {
	return match(OP_NOT_EQUAL, v)
}
