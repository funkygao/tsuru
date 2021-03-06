// Copyright 2012 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package app

// actions represents an action, with the given methods
// to forward and backward for the action.
type action interface {
	forward(app *App) error
	backward(app *App)

	// rollbackItself indicates whether backward should be called when forward
	// fail. If false, only previously executed actions will be rolled back on
	// forward failures.
	rollbackItself() bool
}

// execute runs an action list. If an errors ocourrs
// execute stops the execution for the actions and call
// the rollback for previous actions.
func execute(a *App, actions []action) error {
	for index, action := range actions {
		err := action.forward(a)
		if err != nil {
			if !action.rollbackItself() {
				index--
			}
			rollBack(a, actions, index)
			return err
		}
	}
	return nil
}

// rollBack runs the rollback for the given actions.
func rollBack(a *App, actions []action, index int) {
	for i := index; i >= 0; i-- {
		action := actions[i]
		action.backward(a)
	}
}
