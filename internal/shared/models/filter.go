package models

type Filter struct {
	Query    Query `json:"query"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}
type FieldComparisons struct {
	Field    string `json:"field"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}
type Expressions struct {
	FieldComparisons []FieldComparisons `json:"fieldComparisons"`
}
type Query struct {
	Expressions Expressions `json:"expressions"`
}

func (f *Filter) QueryToSqlWhere() string {
	filter := ""
	filtersSize := len(f.Query.Expressions.FieldComparisons)
	if f.Query.Expressions.FieldComparisons != nil || filtersSize > 0 {
		filter += " WHERE "

		for i, e := range f.Query.Expressions.FieldComparisons {
			filter += e.Field + " " + e.Operator + " '" + e.Value + "' "
			if i < filtersSize-1 {
				filter += " AND "
			}
		}

	}
	return filter
}
