package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

/////////////////
// HashElement //
/////////////////

type HashElement interface {
	js.Object
}
type hashElement struct{ js.Object }

// NewHashElement is the constructor for HashElement.
func NewHashElement() HashElement {
	return &hashElement{pcc.Get("HashElement").New()}
}

///////////////////
// ActionManager //
///////////////////

type ActionManager interface {
	js.Object
	AddAction(Action, Node, bool)
	RemoveAllActions()
	RemoveAllActionsFromTarget(Node, bool)
	RemoveAction(Action)
	RemoveActionByTag(int, Node)
	GetActionByTag(int, Node) Action
	NumberOfRunningActionsInTarget(Node) int
	PauseTarget(Node)
	ResumeTarget(Node)
	PauseAllRunningActions() []Node
	ResumeTargets([]Node)
	PurgeSharedManager()
	Update(float64)
}
type actionManager struct{ js.Object }

// NewActionManager is the constructor for ActionManager.
func NewActionManager() ActionManager {
	return &actionManager{pcc.Get("ActionManager").New()}
}

// AddAction adds an action with a target.
// If the target is already present, then the action will be added to the existing target.
// If the target is not present, a new instance of this target will be created either paused or not, and the action will be added to the newly created target.
// When the target is paused, the queued actions won't be 'ticked'.
func (am *actionManager) AddAction(action Action, node Node, paused bool) {
	am.Call("addAction", action, node, paused)
}

// RemoveAllActions removes all actions from all the targets.
func (am *actionManager) RemoveAllActions() {
	am.Call("removeAllActions")
}

// RemoveAllActionsFromTarget removes all actions from a certain target.
// All the actions that belongs to the target will be removed.
func (am *actionManager) RemoveAllActionsFromTarget(target Node, forceDelete bool) {
	am.Call("removeAllActionsFromTarget", target, forceDelete)
}

// RemoveAction removes an action given an action reference.
func (am *actionManager) RemoveAction(action Action) {
	am.Call("removeAction", action)
}

// RemoveActionByTag removes an action given its tag and the target.
func (am *actionManager) RemoveActionByTag(tag int, target Node) {
	am.Call("removeActionByTag", tag, target)
}

// GetActionByTag gets an action given its tag an a target.
func (am *actionManager) GetActionByTag(tag int, target Node) Action {
	return &action{am.Call("getActionByTag", tag, target)}
}

// NumberOfRunningActionsInTarget returns the numbers of actions that are running in a certain target.
// Composable actions are counted as 1 action.
func (am *actionManager) NumberOfRunningActionsInTarget(target Node) int {
	return am.Call("numberOfRunningActionsInTarget", target).Int()
}

// PauseTarget pauses the target: all running actions and newly added actions will be paused.
func (am *actionManager) PauseTarget(target Node) {
	am.Call("pauseTarget", target)
}

// ResumeTarget resumes the target. All queued actions will be resumed.
func (am *actionManager) ResumeTarget(target Node) {
	am.Call("resumeTarget", target)
}

// PauseAllRunningActions pauses all running actions, returning a list of targets whose actions were paused.
func (am *actionManager) PauseAllRunningActions() []Node {
	array := am.Call("pauseAllRunningActions")
	length := array.Length()
	result := make([]Node, length)

	for i := 0; i < length; i++ {
		result[i] = &node{array.Index(i)}
	}

	return result
}

// ResumeTargets resume a set of targets (convenience function to reverse a pauseAllRunningActions call)
func (am *actionManager) ResumeTargets(targetsToResume []Node) {
	am.Call("resumeTargets", targetsToResume)
}

// PurgeSharedManager purges the shared action manager. It releases the retained instance.
// because it uses this, so it can not be static.
func (am *actionManager) PurgeSharedManager() {
	am.Call("purgeSharedManager")
}

// Update: dt delta time in seconds
func (am *actionManager) Update(dt float64) {
	am.Call("update", dt)
}
