package forms

type errors map[string][]string

//Add error message for given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

//Get return first error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
