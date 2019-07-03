package controller

import (
	"new/roarapp-operator/pkg/controller/roarapp"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, roarapp.Add)
}
