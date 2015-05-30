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

type Spawn interface {
	ActionInterval
	InitWithTwoActions(FiniteTimeAction, FiniteTimeAction) bool
	//Clone() Action defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Stop() defined as part of Action
	//Update(float64) defined as part of Action
	//Reverse() FiniteTimeAction defined as part of Action
}

type spawn struct{ actionInterval }

func NewSpawn(actions ...FiniteTimeAction) Spawn {
	return &spawn{actionInterval{finiteTimeAction{action{pcc.Call("spawn", actions)}}}}
}

func (s *spawn) InitWithTwoActions(actionOne FiniteTimeAction, actionTwo FiniteTimeAction) bool {
	return s.Call("initWithTwoActions", actionOne, actionTwo).Bool()
}

func (s *spawn) Clone() Action {
	return &spawn{actionInterval{finiteTimeAction{action{s.Call("clone")}}}}
}

func (s *spawn) StartWithTarget(target Node) {
	s.Call("startWithTarget", target)
}

func (s *spawn) Stop() {
	s.Call("stop")
}

func (s *spawn) Update(dt float64) {
	s.Call("update", dt)
}

func (s *spawn) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{s.Call("reverse")}}
}

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

func NewRotateTo(duration float64, deltaAngleX float64, deltaAngleY float64) RotateTo {
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

func NewRotateBy(duration float64, deltaAngleX float64, deltaAngleY float64) RotateBy {
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

type MoveBy interface {
	ActionInterval
	InitWithDurationAndDeltaPosition(float64, Point) bool
	//Clone() Action defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Update(float64) defined as part of Action
	//Reverse() FiniteTimeAction defined as part of Action
}

type moveBy struct{ actionInterval }

func NewMoveBy(duration float64, position Point) MoveBy {
	return &moveBy{actionInterval{finiteTimeAction{action{pcc.Call("moveBy", duration, position)}}}}
}

func (m *moveBy) InitWithDurationAndDeltaPosition(duration float64, position Point) bool {
	return m.Call("initWithDuration", duration, position).Bool()
}

func (m *moveBy) Clone() Action {
	return &moveBy{actionInterval{finiteTimeAction{action{m.Call("clone")}}}}
}

func (m *moveBy) StartWithTarget(target Node) {
	m.Call("startWithTarget", target)
}

func (m *moveBy) Update(dt float64) {
	m.Call("update", dt)
}

func (m *moveBy) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{m.Call("reverse")}}
}

////////////
// MoveTo //
////////////

type MoveTo interface {
	ActionInterval
	InitWithDurationAndDeltaPosition(float64, Point) bool
	//Clone() Action defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Update(float64) defined as part of Action
	//Reverse() FiniteTimeAction defined as part of Action
}

type moveTo struct{ actionInterval }

func NewMoveTo(duration float64, position Point) MoveTo {
	return &moveTo{actionInterval{finiteTimeAction{action{pcc.Call("moveTo", duration, position)}}}}
}

func (m *moveTo) InitWithDurationAndDeltaPosition(duration float64, position Point) bool {
	return m.Call("initWithDuration", duration, position).Bool()
}

func (m *moveTo) Clone() Action {
	return &moveTo{actionInterval{finiteTimeAction{action{m.Call("clone")}}}}
}

func (m *moveTo) StartWithTarget(target Node) {
	m.Call("startWithTarget", target)
}

func (m *moveTo) Update(dt float64) {
	m.Call("update", dt)
}

func (m *moveTo) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{m.Call("reverse")}}
}

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

func NewScaleTo(duration float64, sx float64, sy float64) ScaleTo {
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

func NewScaleBy(duration float64, sx float64, sy float64) ScaleBy {
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

type TintTo interface {
	ActionInterval
	InitWithDurationAndColor(float64, int, int, int) bool
	//Clone() Action defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Update(float64) defined as part of Action
	//Reverse() FiniteTimeAction defined as part of Action
}

type tintTo struct{ actionInterval }

func NewTintTo(duration float64, red int, green int, blue int) TintTo {
	return &tintTo{actionInterval{finiteTimeAction{action{pcc.Call("tintTo", duration, red, green, blue)}}}}
}

func (t *tintTo) InitWithDurationAndColor(duration float64, red int, green int, blue int) bool {
	return t.Call("initWithDuration", duration, red, green, blue).Bool()
}

func (t *tintTo) Clone() Action {
	return &tintTo{actionInterval{finiteTimeAction{action{t.Call("clone")}}}}
}

func (t *tintTo) StartWithTarget(target Node) {
	t.Call("startWithTarget", target)
}

func (t *tintTo) Update(dt float64) {
	t.Call("update", dt)
}

func (t *tintTo) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{t.Call("reverse")}}
}

////////////
// TintBy //
////////////

type TintBy interface {
	ActionInterval
	InitWithDurationAndColor(float64, int, int, int) bool
	//Clone() Action defined as part of Action
	//StartWithTarget(Node) defined as part of Action
	//Update(float64) defined as part of Action
	//Reverse() FiniteTimeAction defined as part of Action
}

type tintBy struct{ actionInterval }

func NewTintBy(duration float64, red int, green int, blue int) TintBy {
	return &tintBy{actionInterval{finiteTimeAction{action{pcc.Call("tintBy", duration, red, green, blue)}}}}
}

func (t *tintBy) InitWithDurationAndColor(duration float64, red int, green int, blue int) bool {
	return t.Call("initWithDuration", duration, red, green, blue).Bool()
}

func (t *tintBy) Clone() Action {
	return &tintTo{actionInterval{finiteTimeAction{action{t.Call("clone")}}}}
}

func (t *tintBy) StartWithTarget(target Node) {
	t.Call("startWithTarget", target)
}

func (t *tintBy) Update(dt float64) {
	t.Call("update", dt)
}

func (t *tintBy) Reverse() FiniteTimeAction {
	return &finiteTimeAction{action{t.Call("reverse")}}
}

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
