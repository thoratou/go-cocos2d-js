package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

////////////
// Action //
////////////

type Action interface {
	Clone() Action
	IsDone() bool
	StartWithTarget(Node)
	Stop()
	SetStep(func(float64))
	Step(float64)
	StepSuper(float64)
	SetUpdate(func(float64))
	Update(float64)
	UpdateSuper(float64)
	GetTarget() Node
	SetTarget(Node)
	GetOriginalTarget() Node
	SetOriginalTarget(Node)
	GetTag() int
	SetTag(int)
	Retain()
	Release()
}

type action struct{ *js.Object }

// NewAction is the constructor for Action.
func NewAction() Action {
	return &action{pcc.Call("action")}
}

//TODO: tmp
func NewActionJs(o *js.Object) Action {
	return &action{o}
}

func (a *action) Clone() Action {
	return &action{a.Call("clone")}
}

func (a *action) IsDone() bool {
	return a.Call("isDone").Bool()
}

func (a *action) StartWithTarget(target Node) {
	a.Call("startWithTarget", target)
}

func (a *action) Stop() {
	a.Call("stop")
}

func (a *action) SetStep(step func(float64)) {
	BackupFunc(a.Object, "step")
	a.Set("step", step)
}

func (a *action) Step(dt float64) {
	a.Call("step", dt)
}

func (a *action) StepSuper(dt float64) {
	SuperCall(a.Object, "step", dt)
}

func (a *action) SetUpdate(update func(float64)) {
	BackupFunc(a.Object, "update")
	a.Set("update", update)
}

func (a *action) Update(dt float64) {
	a.Call("update", dt)
}

func (a *action) UpdateSuper(dt float64) {
	SuperCall(a.Object, "update", dt)
}

func (a *action) GetTarget() Node {
	return &node{a.Call("getTarget")}
}

func (a *action) SetTarget(target Node) {
	a.Call("setTarget", target)
}

func (a *action) GetOriginalTarget() Node {
	return &node{a.Call("getOriginalTarget")}
}

func (a *action) SetOriginalTarget(target Node) {
	a.Call("setOriginalTarget", target)
}

func (a *action) GetTag() int {
	return a.Call("getTag").Int()
}

func (a *action) SetTag(tag int) {
	a.Call("setTag", tag)
}

func (a *action) Retain() {
	a.Call("retain")
}

func (a *action) Release() {
	a.Call("release")
}

//////////////////////
// FiniteTimeAction //
//////////////////////

type FiniteTimeAction interface {
	Action
	GetDuration() float64
	SetDuration(float64)
	Reverse() FiniteTimeAction
	//Clone() Action defined as part of Action
}

type finiteTimeAction struct{ action }

func (f *finiteTimeAction) GetDuration() float64 {
	return f.Call("getDuration").Float()
}

func (f *finiteTimeAction) SetDuration(duration float64) {
	f.Call("setDuration", duration)
}

func (f *finiteTimeAction) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{f.Call("reverse")}}
}

func (f *finiteTimeAction) Clone() Action {
	return &finiteTimeAction{action{f.Call("clone")}}
}

///////////
// Speed //
///////////

//TODO Speed

////////////
// Follow //
////////////

//TODO Follow
