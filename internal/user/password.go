package user

import "xcluster/pkg/argon2"

type Password string

func (p Password) Compare(other string) (bool, error) {
	wrapper, err := argon2.NewHashWrapperFromString(string(p))
	if err != nil {
		return false, err
	}
	return wrapper.CompareString(other), nil
}
