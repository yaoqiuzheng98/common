package environment

type Environment string

const (
	Development Environment = "dev"
	Production  Environment = "prod"
	Test        Environment = "test"
)

func (receiver Environment) String() string {
	return string(receiver)
}
