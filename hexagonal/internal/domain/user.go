package domain

// Pure business object
// No references to infrastructure code
// No references to external code
// No references to other business objects
// No DB tags, no JSON tags
type User struct {
	ID    string
	Name  string
	Email string
}
