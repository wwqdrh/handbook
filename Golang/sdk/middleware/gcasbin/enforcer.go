package gcasbin

import (
	"github.com/casbin/casbin"
)

// Use the Model file and default FileAdapter:
func BasicEnforcer(model, policy string) (enforce *casbin.Enforcer, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	enforce = casbin.NewEnforcer(model, policy)
	return
}
