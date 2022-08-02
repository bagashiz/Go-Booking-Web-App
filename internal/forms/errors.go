package forms

// errors is a custom errors struct that embeds a map[string][]string.
type errors map[string][]string

// Add is a function that adds an error message for a given field.
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get is a function that returns the first error message for a given field.
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
