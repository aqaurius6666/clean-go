package gentity

type Extend[T E] struct {
	Entity   T
	ExFields *ExtendFields
}
type ExtendFields struct {
	Offset    *int
	Limit     *int
	Fields    []string
	OrderBy   []string
	OrderType []OrderType
	Filters   []Filter `validate:"dive"`
	Debug     bool
	Joins     []string
}
type Filter struct {
	Field    string      `json:"f" validate:"required"`
	Operator string      `json:"o" validate:"eq==|eq=<>|eq=>|eq=<|eq=>=|eq=<="`
	Value    interface{} `json:"v" validate:"required"`
}
type Options interface {
	apply(*ExtendFields)
}

func WithExtend[T E](entity T, ex *ExtendFields, opts ...Options) Extend[T] {
	if ex == nil {
		ex = &ExtendFields{}
	}
	for _, opt := range opts {
		opt.apply(ex)
	}
	return Extend[T]{
		Entity:   entity,
		ExFields: ex,
	}
}

type OrderType string
type Operator string

var (
	ASC  OrderType = "ASC"
	DESC OrderType = "DESC"

	EQ Operator = "eq"
	NE Operator = "ne"
	GT Operator = "gt"
	GE Operator = "ge"
	LT Operator = "lt"
	LE Operator = "le"
)

func (s Operator) ToOp() string {
	switch s {
	case EQ:
		return "="
	case NE:
		return "!="
	case GT:
		return ">"
	case GE:
		return ">="
	case LT:
		return "<"
	case LE:
		return "<="
	default:
		panic("invalid operator")
	}
}

// func ToOperator(a string) Operator {
// 	switch a {
// 	case "eq":
// 		return EQ
// 	case "ne":
// 		return NE
// 	case "gt":
// 		return GT
// 	case "ge":
// 		return GE
// 	case "lt":
// 		return LT
// 	case "le":
// 		return LE
// 	default:
// 		return EQ
// 	}
// }

type offset struct {
	value int
}

func Offset(o int) Options {
	return offset{value: o}
}

func (s offset) apply(e *ExtendFields) {
	e.Offset = &s.value
}

type limit struct {
	value int
}

func Limit(o int) Options {
	return limit{value: o}
}

func (s limit) apply(e *ExtendFields) {
	e.Limit = &s.value
}

type fields struct {
	value []string
}

func Fields(f []string) Options {
	return fields{value: f}
}

func (s fields) apply(e *ExtendFields) {
	e.Fields = s.value
}

type debug struct {
}

func Debug() Options {
	return debug{}
}

func (s debug) apply(e *ExtendFields) {
	e.Debug = true
}

type order struct {
	orderBy   []string
	orderType []OrderType
}

func Order(orderBy []string, orderType []OrderType) Options {
	return order{orderBy: orderBy, orderType: orderType}
}

func (s order) apply(e *ExtendFields) {
	e.OrderBy = s.orderBy
	e.OrderType = s.orderType
}
