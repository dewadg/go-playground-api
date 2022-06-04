package resolver

type resolver struct{}

func New() *resolver {
	return &resolver{}
}

func (r *resolver) Ping() (string, error) {
	return "pong", nil
}
