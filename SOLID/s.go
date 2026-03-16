package main


// S — Single Responsibility Principle
// A type should have one reason to change. In Go, this often means splitting structs and keeping interfaces narrow.

// BAD — one struct doing too much
type UserService struct{}
func (s *UserService) SaveUser(u User) error { ... }
func (s *UserService) SendWelcomeEmail(u User) error { ... }
func (s *UserService) GenerateReport() []byte { ... }

// GOOD — each type has one job
type UserRepo struct{ db *sql.DB }
func (r *UserRepo) Save(u User) error { ... }

type Mailer struct{ client SMTPClient }
func (m *Mailer) SendWelcome(u User) error { ... }

type Reporter struct{ client ReportClient }
func (r *Reporter) Generate() []byte { ... }

type UserService struct {
	repo UserRepo
	mail Mailer
	rep  Reporter
}
func (s *UserService) SaveUser(u User) error {
	return s.repo.Save(u)
}
func (s *UserService) SendWelcomeEmail(u User) error {
	return s.mail.SendWelcome(u)
}
func (s *UserService) GenerateReport() []byte {
	return s.rep.Generate()
}
