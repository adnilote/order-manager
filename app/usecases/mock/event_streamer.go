package mock

// Code generated by http://github.com/gojuno/minimock (3.0.8). DO NOT EDIT.

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// EventStreamerMock implements usecases.EventStreamer
type EventStreamerMock struct {
	t minimock.Tester

	funcRegister          func(wsuuid string) (ch1 chan interface{})
	inspectFuncRegister   func(wsuuid string)
	afterRegisterCounter  uint64
	beforeRegisterCounter uint64
	RegisterMock          mEventStreamerMockRegister

	funcSendToAll          func(data interface{}) (err error)
	inspectFuncSendToAll   func(data interface{})
	afterSendToAllCounter  uint64
	beforeSendToAllCounter uint64
	SendToAllMock          mEventStreamerMockSendToAll

	funcUnregister          func(wsuuid string)
	inspectFuncUnregister   func(wsuuid string)
	afterUnregisterCounter  uint64
	beforeUnregisterCounter uint64
	UnregisterMock          mEventStreamerMockUnregister
}

// NewEventStreamerMock returns a mock for usecases.EventStreamer
func NewEventStreamerMock(t minimock.Tester) *EventStreamerMock {
	m := &EventStreamerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.RegisterMock = mEventStreamerMockRegister{mock: m}
	m.RegisterMock.callArgs = []*EventStreamerMockRegisterParams{}

	m.SendToAllMock = mEventStreamerMockSendToAll{mock: m}
	m.SendToAllMock.callArgs = []*EventStreamerMockSendToAllParams{}

	m.UnregisterMock = mEventStreamerMockUnregister{mock: m}
	m.UnregisterMock.callArgs = []*EventStreamerMockUnregisterParams{}

	return m
}

type mEventStreamerMockRegister struct {
	mock               *EventStreamerMock
	defaultExpectation *EventStreamerMockRegisterExpectation
	expectations       []*EventStreamerMockRegisterExpectation

	callArgs []*EventStreamerMockRegisterParams
	mutex    sync.RWMutex
}

// EventStreamerMockRegisterExpectation specifies expectation struct of the EventStreamer.Register
type EventStreamerMockRegisterExpectation struct {
	mock    *EventStreamerMock
	params  *EventStreamerMockRegisterParams
	results *EventStreamerMockRegisterResults
	Counter uint64
}

// EventStreamerMockRegisterParams contains parameters of the EventStreamer.Register
type EventStreamerMockRegisterParams struct {
	wsuuid string
}

// EventStreamerMockRegisterResults contains results of the EventStreamer.Register
type EventStreamerMockRegisterResults struct {
	ch1 chan interface{}
}

// Expect sets up expected params for EventStreamer.Register
func (mmRegister *mEventStreamerMockRegister) Expect(wsuuid string) *mEventStreamerMockRegister {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("EventStreamerMock.Register mock is already set by Set")
	}

	if mmRegister.defaultExpectation == nil {
		mmRegister.defaultExpectation = &EventStreamerMockRegisterExpectation{}
	}

	mmRegister.defaultExpectation.params = &EventStreamerMockRegisterParams{wsuuid}
	for _, e := range mmRegister.expectations {
		if minimock.Equal(e.params, mmRegister.defaultExpectation.params) {
			mmRegister.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmRegister.defaultExpectation.params)
		}
	}

	return mmRegister
}

// Inspect accepts an inspector function that has same arguments as the EventStreamer.Register
func (mmRegister *mEventStreamerMockRegister) Inspect(f func(wsuuid string)) *mEventStreamerMockRegister {
	if mmRegister.mock.inspectFuncRegister != nil {
		mmRegister.mock.t.Fatalf("Inspect function is already set for EventStreamerMock.Register")
	}

	mmRegister.mock.inspectFuncRegister = f

	return mmRegister
}

// Return sets up results that will be returned by EventStreamer.Register
func (mmRegister *mEventStreamerMockRegister) Return(ch1 chan interface{}) *EventStreamerMock {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("EventStreamerMock.Register mock is already set by Set")
	}

	if mmRegister.defaultExpectation == nil {
		mmRegister.defaultExpectation = &EventStreamerMockRegisterExpectation{mock: mmRegister.mock}
	}
	mmRegister.defaultExpectation.results = &EventStreamerMockRegisterResults{ch1}
	return mmRegister.mock
}

//Set uses given function f to mock the EventStreamer.Register method
func (mmRegister *mEventStreamerMockRegister) Set(f func(wsuuid string) (ch1 chan interface{})) *EventStreamerMock {
	if mmRegister.defaultExpectation != nil {
		mmRegister.mock.t.Fatalf("Default expectation is already set for the EventStreamer.Register method")
	}

	if len(mmRegister.expectations) > 0 {
		mmRegister.mock.t.Fatalf("Some expectations are already set for the EventStreamer.Register method")
	}

	mmRegister.mock.funcRegister = f
	return mmRegister.mock
}

// When sets expectation for the EventStreamer.Register which will trigger the result defined by the following
// Then helper
func (mmRegister *mEventStreamerMockRegister) When(wsuuid string) *EventStreamerMockRegisterExpectation {
	if mmRegister.mock.funcRegister != nil {
		mmRegister.mock.t.Fatalf("EventStreamerMock.Register mock is already set by Set")
	}

	expectation := &EventStreamerMockRegisterExpectation{
		mock:   mmRegister.mock,
		params: &EventStreamerMockRegisterParams{wsuuid},
	}
	mmRegister.expectations = append(mmRegister.expectations, expectation)
	return expectation
}

// Then sets up EventStreamer.Register return parameters for the expectation previously defined by the When method
func (e *EventStreamerMockRegisterExpectation) Then(ch1 chan interface{}) *EventStreamerMock {
	e.results = &EventStreamerMockRegisterResults{ch1}
	return e.mock
}

// Register implements usecases.EventStreamer
func (mmRegister *EventStreamerMock) Register(wsuuid string) (ch1 chan interface{}) {
	mm_atomic.AddUint64(&mmRegister.beforeRegisterCounter, 1)
	defer mm_atomic.AddUint64(&mmRegister.afterRegisterCounter, 1)

	if mmRegister.inspectFuncRegister != nil {
		mmRegister.inspectFuncRegister(wsuuid)
	}

	mm_params := &EventStreamerMockRegisterParams{wsuuid}

	// Record call args
	mmRegister.RegisterMock.mutex.Lock()
	mmRegister.RegisterMock.callArgs = append(mmRegister.RegisterMock.callArgs, mm_params)
	mmRegister.RegisterMock.mutex.Unlock()

	for _, e := range mmRegister.RegisterMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ch1
		}
	}

	if mmRegister.RegisterMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmRegister.RegisterMock.defaultExpectation.Counter, 1)
		mm_want := mmRegister.RegisterMock.defaultExpectation.params
		mm_got := EventStreamerMockRegisterParams{wsuuid}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmRegister.t.Errorf("EventStreamerMock.Register got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmRegister.RegisterMock.defaultExpectation.results
		if mm_results == nil {
			mmRegister.t.Fatal("No results are set for the EventStreamerMock.Register")
		}
		return (*mm_results).ch1
	}
	if mmRegister.funcRegister != nil {
		return mmRegister.funcRegister(wsuuid)
	}
	mmRegister.t.Fatalf("Unexpected call to EventStreamerMock.Register. %v", wsuuid)
	return
}

// RegisterAfterCounter returns a count of finished EventStreamerMock.Register invocations
func (mmRegister *EventStreamerMock) RegisterAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegister.afterRegisterCounter)
}

// RegisterBeforeCounter returns a count of EventStreamerMock.Register invocations
func (mmRegister *EventStreamerMock) RegisterBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmRegister.beforeRegisterCounter)
}

// Calls returns a list of arguments used in each call to EventStreamerMock.Register.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmRegister *mEventStreamerMockRegister) Calls() []*EventStreamerMockRegisterParams {
	mmRegister.mutex.RLock()

	argCopy := make([]*EventStreamerMockRegisterParams, len(mmRegister.callArgs))
	copy(argCopy, mmRegister.callArgs)

	mmRegister.mutex.RUnlock()

	return argCopy
}

// MinimockRegisterDone returns true if the count of the Register invocations corresponds
// the number of defined expectations
func (m *EventStreamerMock) MinimockRegisterDone() bool {
	for _, e := range m.RegisterMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RegisterMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRegisterCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegister != nil && mm_atomic.LoadUint64(&m.afterRegisterCounter) < 1 {
		return false
	}
	return true
}

// MinimockRegisterInspect logs each unmet expectation
func (m *EventStreamerMock) MinimockRegisterInspect() {
	for _, e := range m.RegisterMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to EventStreamerMock.Register with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.RegisterMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterRegisterCounter) < 1 {
		if m.RegisterMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to EventStreamerMock.Register")
		} else {
			m.t.Errorf("Expected call to EventStreamerMock.Register with params: %#v", *m.RegisterMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcRegister != nil && mm_atomic.LoadUint64(&m.afterRegisterCounter) < 1 {
		m.t.Error("Expected call to EventStreamerMock.Register")
	}
}

type mEventStreamerMockSendToAll struct {
	mock               *EventStreamerMock
	defaultExpectation *EventStreamerMockSendToAllExpectation
	expectations       []*EventStreamerMockSendToAllExpectation

	callArgs []*EventStreamerMockSendToAllParams
	mutex    sync.RWMutex
}

// EventStreamerMockSendToAllExpectation specifies expectation struct of the EventStreamer.SendToAll
type EventStreamerMockSendToAllExpectation struct {
	mock    *EventStreamerMock
	params  *EventStreamerMockSendToAllParams
	results *EventStreamerMockSendToAllResults
	Counter uint64
}

// EventStreamerMockSendToAllParams contains parameters of the EventStreamer.SendToAll
type EventStreamerMockSendToAllParams struct {
	data interface{}
}

// EventStreamerMockSendToAllResults contains results of the EventStreamer.SendToAll
type EventStreamerMockSendToAllResults struct {
	err error
}

// Expect sets up expected params for EventStreamer.SendToAll
func (mmSendToAll *mEventStreamerMockSendToAll) Expect(data interface{}) *mEventStreamerMockSendToAll {
	if mmSendToAll.mock.funcSendToAll != nil {
		mmSendToAll.mock.t.Fatalf("EventStreamerMock.SendToAll mock is already set by Set")
	}

	if mmSendToAll.defaultExpectation == nil {
		mmSendToAll.defaultExpectation = &EventStreamerMockSendToAllExpectation{}
	}

	mmSendToAll.defaultExpectation.params = &EventStreamerMockSendToAllParams{data}
	for _, e := range mmSendToAll.expectations {
		if minimock.Equal(e.params, mmSendToAll.defaultExpectation.params) {
			mmSendToAll.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSendToAll.defaultExpectation.params)
		}
	}

	return mmSendToAll
}

// Inspect accepts an inspector function that has same arguments as the EventStreamer.SendToAll
func (mmSendToAll *mEventStreamerMockSendToAll) Inspect(f func(data interface{})) *mEventStreamerMockSendToAll {
	if mmSendToAll.mock.inspectFuncSendToAll != nil {
		mmSendToAll.mock.t.Fatalf("Inspect function is already set for EventStreamerMock.SendToAll")
	}

	mmSendToAll.mock.inspectFuncSendToAll = f

	return mmSendToAll
}

// Return sets up results that will be returned by EventStreamer.SendToAll
func (mmSendToAll *mEventStreamerMockSendToAll) Return(err error) *EventStreamerMock {
	if mmSendToAll.mock.funcSendToAll != nil {
		mmSendToAll.mock.t.Fatalf("EventStreamerMock.SendToAll mock is already set by Set")
	}

	if mmSendToAll.defaultExpectation == nil {
		mmSendToAll.defaultExpectation = &EventStreamerMockSendToAllExpectation{mock: mmSendToAll.mock}
	}
	mmSendToAll.defaultExpectation.results = &EventStreamerMockSendToAllResults{err}
	return mmSendToAll.mock
}

//Set uses given function f to mock the EventStreamer.SendToAll method
func (mmSendToAll *mEventStreamerMockSendToAll) Set(f func(data interface{}) (err error)) *EventStreamerMock {
	if mmSendToAll.defaultExpectation != nil {
		mmSendToAll.mock.t.Fatalf("Default expectation is already set for the EventStreamer.SendToAll method")
	}

	if len(mmSendToAll.expectations) > 0 {
		mmSendToAll.mock.t.Fatalf("Some expectations are already set for the EventStreamer.SendToAll method")
	}

	mmSendToAll.mock.funcSendToAll = f
	return mmSendToAll.mock
}

// When sets expectation for the EventStreamer.SendToAll which will trigger the result defined by the following
// Then helper
func (mmSendToAll *mEventStreamerMockSendToAll) When(data interface{}) *EventStreamerMockSendToAllExpectation {
	if mmSendToAll.mock.funcSendToAll != nil {
		mmSendToAll.mock.t.Fatalf("EventStreamerMock.SendToAll mock is already set by Set")
	}

	expectation := &EventStreamerMockSendToAllExpectation{
		mock:   mmSendToAll.mock,
		params: &EventStreamerMockSendToAllParams{data},
	}
	mmSendToAll.expectations = append(mmSendToAll.expectations, expectation)
	return expectation
}

// Then sets up EventStreamer.SendToAll return parameters for the expectation previously defined by the When method
func (e *EventStreamerMockSendToAllExpectation) Then(err error) *EventStreamerMock {
	e.results = &EventStreamerMockSendToAllResults{err}
	return e.mock
}

// SendToAll implements usecases.EventStreamer
func (mmSendToAll *EventStreamerMock) SendToAll(data interface{}) (err error) {
	mm_atomic.AddUint64(&mmSendToAll.beforeSendToAllCounter, 1)
	defer mm_atomic.AddUint64(&mmSendToAll.afterSendToAllCounter, 1)

	if mmSendToAll.inspectFuncSendToAll != nil {
		mmSendToAll.inspectFuncSendToAll(data)
	}

	mm_params := &EventStreamerMockSendToAllParams{data}

	// Record call args
	mmSendToAll.SendToAllMock.mutex.Lock()
	mmSendToAll.SendToAllMock.callArgs = append(mmSendToAll.SendToAllMock.callArgs, mm_params)
	mmSendToAll.SendToAllMock.mutex.Unlock()

	for _, e := range mmSendToAll.SendToAllMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSendToAll.SendToAllMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSendToAll.SendToAllMock.defaultExpectation.Counter, 1)
		mm_want := mmSendToAll.SendToAllMock.defaultExpectation.params
		mm_got := EventStreamerMockSendToAllParams{data}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSendToAll.t.Errorf("EventStreamerMock.SendToAll got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSendToAll.SendToAllMock.defaultExpectation.results
		if mm_results == nil {
			mmSendToAll.t.Fatal("No results are set for the EventStreamerMock.SendToAll")
		}
		return (*mm_results).err
	}
	if mmSendToAll.funcSendToAll != nil {
		return mmSendToAll.funcSendToAll(data)
	}
	mmSendToAll.t.Fatalf("Unexpected call to EventStreamerMock.SendToAll. %v", data)
	return
}

// SendToAllAfterCounter returns a count of finished EventStreamerMock.SendToAll invocations
func (mmSendToAll *EventStreamerMock) SendToAllAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendToAll.afterSendToAllCounter)
}

// SendToAllBeforeCounter returns a count of EventStreamerMock.SendToAll invocations
func (mmSendToAll *EventStreamerMock) SendToAllBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSendToAll.beforeSendToAllCounter)
}

// Calls returns a list of arguments used in each call to EventStreamerMock.SendToAll.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSendToAll *mEventStreamerMockSendToAll) Calls() []*EventStreamerMockSendToAllParams {
	mmSendToAll.mutex.RLock()

	argCopy := make([]*EventStreamerMockSendToAllParams, len(mmSendToAll.callArgs))
	copy(argCopy, mmSendToAll.callArgs)

	mmSendToAll.mutex.RUnlock()

	return argCopy
}

// MinimockSendToAllDone returns true if the count of the SendToAll invocations corresponds
// the number of defined expectations
func (m *EventStreamerMock) MinimockSendToAllDone() bool {
	for _, e := range m.SendToAllMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendToAllMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendToAllCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendToAll != nil && mm_atomic.LoadUint64(&m.afterSendToAllCounter) < 1 {
		return false
	}
	return true
}

// MinimockSendToAllInspect logs each unmet expectation
func (m *EventStreamerMock) MinimockSendToAllInspect() {
	for _, e := range m.SendToAllMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to EventStreamerMock.SendToAll with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SendToAllMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSendToAllCounter) < 1 {
		if m.SendToAllMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to EventStreamerMock.SendToAll")
		} else {
			m.t.Errorf("Expected call to EventStreamerMock.SendToAll with params: %#v", *m.SendToAllMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSendToAll != nil && mm_atomic.LoadUint64(&m.afterSendToAllCounter) < 1 {
		m.t.Error("Expected call to EventStreamerMock.SendToAll")
	}
}

type mEventStreamerMockUnregister struct {
	mock               *EventStreamerMock
	defaultExpectation *EventStreamerMockUnregisterExpectation
	expectations       []*EventStreamerMockUnregisterExpectation

	callArgs []*EventStreamerMockUnregisterParams
	mutex    sync.RWMutex
}

// EventStreamerMockUnregisterExpectation specifies expectation struct of the EventStreamer.Unregister
type EventStreamerMockUnregisterExpectation struct {
	mock   *EventStreamerMock
	params *EventStreamerMockUnregisterParams

	Counter uint64
}

// EventStreamerMockUnregisterParams contains parameters of the EventStreamer.Unregister
type EventStreamerMockUnregisterParams struct {
	wsuuid string
}

// Expect sets up expected params for EventStreamer.Unregister
func (mmUnregister *mEventStreamerMockUnregister) Expect(wsuuid string) *mEventStreamerMockUnregister {
	if mmUnregister.mock.funcUnregister != nil {
		mmUnregister.mock.t.Fatalf("EventStreamerMock.Unregister mock is already set by Set")
	}

	if mmUnregister.defaultExpectation == nil {
		mmUnregister.defaultExpectation = &EventStreamerMockUnregisterExpectation{}
	}

	mmUnregister.defaultExpectation.params = &EventStreamerMockUnregisterParams{wsuuid}
	for _, e := range mmUnregister.expectations {
		if minimock.Equal(e.params, mmUnregister.defaultExpectation.params) {
			mmUnregister.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmUnregister.defaultExpectation.params)
		}
	}

	return mmUnregister
}

// Inspect accepts an inspector function that has same arguments as the EventStreamer.Unregister
func (mmUnregister *mEventStreamerMockUnregister) Inspect(f func(wsuuid string)) *mEventStreamerMockUnregister {
	if mmUnregister.mock.inspectFuncUnregister != nil {
		mmUnregister.mock.t.Fatalf("Inspect function is already set for EventStreamerMock.Unregister")
	}

	mmUnregister.mock.inspectFuncUnregister = f

	return mmUnregister
}

// Return sets up results that will be returned by EventStreamer.Unregister
func (mmUnregister *mEventStreamerMockUnregister) Return() *EventStreamerMock {
	if mmUnregister.mock.funcUnregister != nil {
		mmUnregister.mock.t.Fatalf("EventStreamerMock.Unregister mock is already set by Set")
	}

	if mmUnregister.defaultExpectation == nil {
		mmUnregister.defaultExpectation = &EventStreamerMockUnregisterExpectation{mock: mmUnregister.mock}
	}

	return mmUnregister.mock
}

//Set uses given function f to mock the EventStreamer.Unregister method
func (mmUnregister *mEventStreamerMockUnregister) Set(f func(wsuuid string)) *EventStreamerMock {
	if mmUnregister.defaultExpectation != nil {
		mmUnregister.mock.t.Fatalf("Default expectation is already set for the EventStreamer.Unregister method")
	}

	if len(mmUnregister.expectations) > 0 {
		mmUnregister.mock.t.Fatalf("Some expectations are already set for the EventStreamer.Unregister method")
	}

	mmUnregister.mock.funcUnregister = f
	return mmUnregister.mock
}

// Unregister implements usecases.EventStreamer
func (mmUnregister *EventStreamerMock) Unregister(wsuuid string) {
	mm_atomic.AddUint64(&mmUnregister.beforeUnregisterCounter, 1)
	defer mm_atomic.AddUint64(&mmUnregister.afterUnregisterCounter, 1)

	if mmUnregister.inspectFuncUnregister != nil {
		mmUnregister.inspectFuncUnregister(wsuuid)
	}

	mm_params := &EventStreamerMockUnregisterParams{wsuuid}

	// Record call args
	mmUnregister.UnregisterMock.mutex.Lock()
	mmUnregister.UnregisterMock.callArgs = append(mmUnregister.UnregisterMock.callArgs, mm_params)
	mmUnregister.UnregisterMock.mutex.Unlock()

	for _, e := range mmUnregister.UnregisterMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmUnregister.UnregisterMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmUnregister.UnregisterMock.defaultExpectation.Counter, 1)
		mm_want := mmUnregister.UnregisterMock.defaultExpectation.params
		mm_got := EventStreamerMockUnregisterParams{wsuuid}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmUnregister.t.Errorf("EventStreamerMock.Unregister got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmUnregister.funcUnregister != nil {
		mmUnregister.funcUnregister(wsuuid)
		return
	}
	mmUnregister.t.Fatalf("Unexpected call to EventStreamerMock.Unregister. %v", wsuuid)

}

// UnregisterAfterCounter returns a count of finished EventStreamerMock.Unregister invocations
func (mmUnregister *EventStreamerMock) UnregisterAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUnregister.afterUnregisterCounter)
}

// UnregisterBeforeCounter returns a count of EventStreamerMock.Unregister invocations
func (mmUnregister *EventStreamerMock) UnregisterBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmUnregister.beforeUnregisterCounter)
}

// Calls returns a list of arguments used in each call to EventStreamerMock.Unregister.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmUnregister *mEventStreamerMockUnregister) Calls() []*EventStreamerMockUnregisterParams {
	mmUnregister.mutex.RLock()

	argCopy := make([]*EventStreamerMockUnregisterParams, len(mmUnregister.callArgs))
	copy(argCopy, mmUnregister.callArgs)

	mmUnregister.mutex.RUnlock()

	return argCopy
}

// MinimockUnregisterDone returns true if the count of the Unregister invocations corresponds
// the number of defined expectations
func (m *EventStreamerMock) MinimockUnregisterDone() bool {
	for _, e := range m.UnregisterMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UnregisterMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUnregisterCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUnregister != nil && mm_atomic.LoadUint64(&m.afterUnregisterCounter) < 1 {
		return false
	}
	return true
}

// MinimockUnregisterInspect logs each unmet expectation
func (m *EventStreamerMock) MinimockUnregisterInspect() {
	for _, e := range m.UnregisterMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to EventStreamerMock.Unregister with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.UnregisterMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterUnregisterCounter) < 1 {
		if m.UnregisterMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to EventStreamerMock.Unregister")
		} else {
			m.t.Errorf("Expected call to EventStreamerMock.Unregister with params: %#v", *m.UnregisterMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcUnregister != nil && mm_atomic.LoadUint64(&m.afterUnregisterCounter) < 1 {
		m.t.Error("Expected call to EventStreamerMock.Unregister")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *EventStreamerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockRegisterInspect()

		m.MinimockSendToAllInspect()

		m.MinimockUnregisterInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *EventStreamerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *EventStreamerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockRegisterDone() &&
		m.MinimockSendToAllDone() &&
		m.MinimockUnregisterDone()
}
