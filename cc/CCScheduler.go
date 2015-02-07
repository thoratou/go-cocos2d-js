package cc

import (
	"github.com/gopherjs/gopherjs/js"
)

const (
	PRIORITY_SYSTEM     = (-2147483647 - 1)
	PRIORITY_NON_SYSTEM = PRIORITY_SYSTEM + 1
)

///////////////
// ListEntry //
///////////////

/////////////////////
// HashUpdateEntry //
/////////////////////

////////////////////
// HashTimerEntry //
////////////////////

///////////
// Timer //
///////////

///////////////
// Scheduler //
///////////////

type Scheduler interface {
	js.Object
	SetTimeScale(float64)
	GetTimeScale() float64
	Update(float64)
	ScheduleCallbackForTarget(js.Object, func(...interface{}), int, int, float64, bool)
	ScheduleUpdateForTarget(js.Object, int, bool)
	UnscheduleCallbackForTarget(js.Object, func(...interface{}))
	UnscheduleUpdateForTarget(js.Object)
	UnscheduleAllCallbacksForTarget(js.Object)
	UnscheduleAllCallbacks()
	UnscheduleAllCallbacksWithMinPriority(int)
	PauseAllTargets()
	PauseAllTargetsWithMinPriority(int)
	ResumeTargets([]js.Object)
	PauseTarget(js.Object)
	ResumeTarget(js.Object)
	IsTargetPaused(js.Object) bool
}

type scheduler struct{ js.Object }

func NewScheduler() Scheduler {
	return &scheduler{pcc.Get("Scheduler").New()}
}

// SetTimeScale modifies the time of all scheduled callbacks.
// You can use this property to create a 'slow motion' or 'fast forward' effect.
// Default is 1.0. To create a 'slow motion' effect, use values below 1.0.
// To create a 'fast forward' effect, use values higher than 1.0.
func (s *scheduler) SetTimeScale(timeScale float64) {
	s.Call("setTimeScale", timeScale)
}

// GetTimeScale returns time scale of scheduler
func (s *scheduler) GetTimeScale() float64 {
	return s.Call("getTimeScale").Float()
}

// Update updates the scheduler. (You should NEVER call this method, unless you know what you are doing).
func (s *scheduler) Update(dt float64) {
	s.Call("update", dt)
}

// The scheduled method will be called every 'interval' seconds.
// If paused is YES, then it won't be called until it is resumed.
// If 'interval' is 0, it will be called every frame, but if so, it recommended to use 'scheduleUpdateForTarget:' instead.
// If the callback function is already scheduled, then only the interval parameter will be updated without re-scheduling it again.
// repeat let the action be repeated repeat + 1 times, use cc.REPEAT_FOREVER to let the action run continuously.
// delay is the amount of time the action will wait before it'll start.
func (s *scheduler) ScheduleCallbackForTarget(target js.Object, callbackFn func(...interface{}), interval int, repeat int, delay float64, paused bool) {
	s.Call("scheduleCallbackForTarget", target, callbackFn, interval, repeat, delay, paused)
}

// ScheduleUpdateForTarget schedules the 'update' callback_fn for a given target with a given priority.
// The 'update' callback_fn will be called every frame.
// The lower the priority, the earlier it is called.
func (s *scheduler) ScheduleUpdateForTarget(target js.Object, priority int, paused bool) {
	s.Call("scheduleUpdateForTarget", target, priority, paused)
}

// UnscheduleCallbackForTarget unschedules a callback function for a given target.
// If you want to unschedule the "update", use UnscheduleUpdateForTarget.
func (s *scheduler) UnscheduleCallbackForTarget(target js.Object, callbackFn func(...interface{})) {
	s.Call("unscheduleCallbackForTarget", target, callbackFn)
}

// UnscheduleUpdateForTarget unschedules the update callback function for a given target.
func (s *scheduler) UnscheduleUpdateForTarget(target js.Object) {
	s.Call("unscheduleUpdateForTarget", target)
}

// UnscheduleAllCallbacksForTarget unschedules all function callbacks for a given target.
// This also includes the "update" callback function.
func (s *scheduler) UnscheduleAllCallbacksForTarget(target js.Object) {
	s.Call("unscheduleAllCallbacksForTarget", target)
}

// UnscheduleAllCallbacks unschedules all function callbacks from all targets.
// You should NEVER call this method, unless you know what you are doing.
func (s *scheduler) UnscheduleAllCallbacks() {
	s.Call("unscheduleAllCallbacks")
}

// UnscheduleAllCallbacksWithMinPriority unschedules all function callbacks from all targets with a minimum priority.
// You should only call this with kCCPriorityNonSystemMin or higher.
func (s *scheduler) UnscheduleAllCallbacksWithMinPriority(minPriority int) {
	s.Call("unscheduleAllCallbacksWithMinPriority", minPriority)
}

// PauseAllTargets pauses all selectors from all targets.
// You should NEVER call this method, unless you know what you are doing.
func (s *scheduler) PauseAllTargets() {
	s.Call("pauseAllTargets")
}

// PauseAllTargetsWithMinPriority pauses all selectors from all targets with a minimum priority.
// You should only call this with kCCPriorityNonSystemMin or higher.
func (s *scheduler) PauseAllTargetsWithMinPriority(minPriority int) {
	s.Call("pauseAllTargetsWithMinPriority", minPriority)
}

// ResumeTargets resumes selectors on a set of targets.
// This can be useful for undoing a call to pauseAllCallbacks.
func (s *scheduler) ResumeTargets(targetsToResume []js.Object) {
	s.Call("resumeTargets", targetsToResume)
}

// PauseTarget Pauses the target.
// All scheduled selectors/update for a given target won't be 'ticked' until the target is resumed.
// If the target is not present, nothing happens.
func (s *scheduler) PauseTarget(target js.Object) {
	s.Call("pauseTarget", target)
}

// ResumeTarget resumes the target.
// The 'target' will be unpaused, so all schedule selectors/update will be 'ticked' again.
// If the target is not present, nothing happens.
func (s *scheduler) ResumeTarget(target js.Object) {
	s.Call("resumeTarget", target)
}

// IsTargetPaused returns whether or not the target is paused
func (s *scheduler) IsTargetPaused(target js.Object) bool {
	return s.Call("isTargetPaused", target).Bool()
}
