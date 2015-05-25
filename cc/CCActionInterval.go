package cc

////////////////////
// ActionInterval //
////////////////////

type ActionInterval interface {
	FiniteTimeAction
	GetElapsed() float64
	InitWithDuration(float64) bool
	//IsDone() bool defined as part of Action
	//Clone() Action  defined as part of Action
	//TODO Easing(ActionEase)
	//Step(float64) defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Reverse() Action defined as part of Action
	SetAmplitudeRate(float64)
	GetAmplitudeRate() float64
	Speed() Action
	GetSpeed() float64
	SetSpeed(float64)
	Repeat(int)
	RepeatForever()
}

type actionInterval struct{ finiteTimeAction }

func NewActionInterval(duration float64) ActionInterval {
	return &actionInterval{finiteTimeAction{action{pcc.Call("actionInterval", duration)}}}
}

func (a *actionInterval) GetElapsed() float64 {
	return a.Call("getElapsed").Float()
}

func (a *actionInterval) InitWithDuration(duration float64) bool {
	return a.Call("initWithDuration", duration).Bool()
}

func (a *actionInterval) IsDone() bool {
	return a.Call("isDone").Bool()
}

func (a *actionInterval) Clone() Action {
	return &actionInterval{finiteTimeAction{action{a.Call("clone")}}}
}

func (a *actionInterval) Step(dt float64) {
	a.Call("step", dt)
}

func (a *actionInterval) StartWithTarget(target Node) {
	a.Call("startWithTarget", target)
}

func (a *actionInterval) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{a.Call("reverse")}}
}

func (a *actionInterval) SetAmplitudeRate(amp float64) {
	a.Call("setAmplitudeRate", amp)
}

func (a *actionInterval) GetAmplitudeRate() float64 {
	return a.Call("getAmplitudeRate").Float()
}

func (a *actionInterval) Speed() Action {
	return &action{a.Call("speed")}
}

func (a *actionInterval) GetSpeed() float64 {
	return a.Call("getSpeed").Float()
}

func (a *actionInterval) SetSpeed(speed float64) {
	a.Call("setSpeed", speed)
}

func (a *actionInterval) Repeat(times int) {
	a.Call("repeat", times)
}

func (a *actionInterval) RepeatForever() {
	a.Call("repeatForever")
}

//////////////
// Sequence //
//////////////

type Sequence interface {
	ActionInterval
	InitWithTwoActions(FiniteTimeAction, FiniteTimeAction) bool
	//Clone() Action defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Stop() defined as part of Action
	//Update(float64) defined as part of Action
	//Reverse() FiniteTimeAction defined as part of Action
}

type sequence struct{ actionInterval }

func NewSequence(actions ...FiniteTimeAction) Sequence {
	return &sequence{actionInterval{finiteTimeAction{action{pcc.Call("sequence", actions)}}}}
}

func (s *sequence) InitWithTwoActions(actionOne FiniteTimeAction, actionTwo FiniteTimeAction) bool {
	return s.Call("initWithTwoActions", actionOne, actionTwo).Bool()
}

func (s *sequence) Clone() Action {
	return &sequence{actionInterval{finiteTimeAction{action{s.Call("clone")}}}}
}

func (s *sequence) StartWithTarget(target Node) {
	s.Call("startWithTarget", target)
}

func (s *sequence) Stop() {
	s.Call("stop")
}

func (s *sequence) Update(dt float64) {
	s.Call("update", dt)
}

func (s *sequence) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{s.Call("reverse")}}
}

////////////
// Repeat //
////////////

///////////////////
// RepeatForever //
///////////////////

///////////
// Spawn //
///////////

//////////////
// RotateTo //
//////////////

type RotateTo interface {
	ActionInterval
	InitWithDurationAndDeltaAngle(float64, float64, float64) bool
	//Clone() Action defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Update(float64) defined as part of Action
	//Reverse() FiniteTimeAction defined as part of Action
}

type rotateTo struct{ actionInterval }

func NewRotateTo(duration float64, deltaAngleX float64, deltaAngleY float64) ActionInterval {
	return &rotateTo{actionInterval{finiteTimeAction{action{pcc.Call("rotateTo", duration, deltaAngleX, deltaAngleY)}}}}
}

func (r *rotateTo) InitWithDurationAndDeltaAngle(duration float64, deltaAngleX float64, deltaAngleY float64) bool {
	return r.Call("initWithDuration", duration, deltaAngleX, deltaAngleY).Bool()
}

func (r *rotateTo) Clone() Action {
	return &rotateTo{actionInterval{finiteTimeAction{action{r.Call("clone")}}}}
}

func (r *rotateTo) StartWithTarget(target Node) {
	r.Call("startWithTarget", target)
}

func (r *rotateTo) Update(dt float64) {
	r.Call("update", dt)
}

func (r *rotateTo) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{r.Call("reverse")}}
}

//////////////
// RotateBy //
//////////////

type RotateBy interface {
	ActionInterval
	InitWithDurationAndDeltaAngle(float64, float64, float64) bool
	//Clone() Action defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Update(float64) defined as part of Action
	//Reverse() FiniteTimeAction defined as part of Action
}

type rotateBy struct{ actionInterval }

func NewRotateBy(duration float64, deltaAngleX float64, deltaAngleY float64) ActionInterval {
	return &rotateBy{actionInterval{finiteTimeAction{action{pcc.Call("rotateBy", duration, deltaAngleX, deltaAngleY)}}}}
}

func (r *rotateBy) InitWithDurationAndDeltaAngle(duration float64, deltaAngleX float64, deltaAngleY float64) bool {
	return r.Call("initWithDuration", duration, deltaAngleX, deltaAngleY).Bool()
}

func (r *rotateBy) Clone() Action {
	return &rotateBy{actionInterval{finiteTimeAction{action{r.Call("clone")}}}}
}

func (r *rotateBy) StartWithTarget(target Node) {
	r.Call("startWithTarget", target)
}

func (r *rotateBy) Update(dt float64) {
	r.Call("update", dt)
}

func (r *rotateBy) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{r.Call("reverse")}}
}

////////////
// MoveBy //
////////////

////////////
// MoveTo //
////////////

////////////
// SkewTo //
////////////

////////////
// SkewBy //
////////////

////////////
// JumpBy //
////////////

////////////
// JumpTo //
////////////

//////////////
// BezierBy //
//////////////

//////////////
// BezierTo //
//////////////

/////////////
// ScaleTo //
/////////////

type ScaleTo interface {
	ActionInterval
	InitWithDurationAndScale(float64, float64, float64) bool
	//Clone() Action defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Update(float64) defined as part of Action
	//Reverse() FiniteTimeAction defined as part of Action
}

type scaleTo struct{ actionInterval }

func NewScaleTo(duration float64, sx float64, sy float64) ActionInterval {
	return &scaleTo{actionInterval{finiteTimeAction{action{pcc.Call("scaleTo", duration, sx, sy)}}}}
}

func (s *scaleTo) InitWithDurationAndScale(duration float64, sx float64, sy float64) bool {
	return s.Call("initWithDuration", duration, sx, sy).Bool()
}

func (s *scaleTo) Clone() Action {
	return &scaleTo{actionInterval{finiteTimeAction{action{s.Call("clone")}}}}
}

func (s *scaleTo) StartWithTarget(target Node) {
	s.Call("startWithTarget", target)
}

func (s *scaleTo) Update(dt float64) {
	s.Call("update", dt)
}

func (s *scaleTo) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{s.Call("reverse")}}
}

/////////////
// ScaleBy //
/////////////

type ScaleBy interface {
	ActionInterval
	InitWithDurationAndScale(float64, float64, float64) bool
	//Clone() Action defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Update(float64) defined as part of Action
	//Reverse() FiniteTimeAction defined as part of Action
}

type scaleBy struct{ actionInterval }

func NewScaleBy(duration float64, sx float64, sy float64) ActionInterval {
	return &scaleBy{actionInterval{finiteTimeAction{action{pcc.Call("scaleBy", duration, sx, sy)}}}}
}

func (s *scaleBy) InitWithDurationAndScale(duration float64, sx float64, sy float64) bool {
	return s.Call("initWithDuration", duration, sx, sy).Bool()
}

func (s *scaleBy) Clone() Action {
	return &scaleBy{actionInterval{finiteTimeAction{action{s.Call("clone")}}}}
}

func (s *scaleBy) StartWithTarget(target Node) {
	s.Call("startWithTarget", target)
}

func (s *scaleBy) Update(dt float64) {
	s.Call("update", dt)
}

func (s *scaleBy) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{s.Call("reverse")}}
}

///////////
// Blink //
///////////

////////////
// FadeTo //
////////////

////////////
// FadeIn //
////////////

/////////////
// FadeOut //
/////////////

////////////
// TintTo //
////////////

////////////
// TintBy //
////////////

///////////////
// DelayTime //
///////////////

/////////////////
// ReverseTime //
/////////////////

/////////////
// Animate //
/////////////

////////////////////
// TargetedAction //
////////////////////
