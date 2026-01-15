package photo

type Criteria struct {
	Filters Filters
}

type Filters []Filter

type Field string

const (
	FieldID   Field = "id"
	FieldName Field = "name"
)

type Filter struct {
	Field Field
	Op    Operator
	Value any
}

type Operator string

const (
	OpEq       Operator = "="
	OpContains Operator = "CONTAINS"
)
