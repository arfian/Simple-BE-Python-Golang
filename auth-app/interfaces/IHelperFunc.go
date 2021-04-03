package interfaces

type IHelperFunc interface {
	GeneratePass() (string)
	GenerateUsername(name string) (string)
}