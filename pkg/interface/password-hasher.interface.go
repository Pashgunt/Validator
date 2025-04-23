package pkginterface

type PasswordHasherInterface interface {
	GetPasswordHash() string
}
