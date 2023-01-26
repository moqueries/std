// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// PtraceDetach_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type PtraceDetach_genType func(pid int) (err error)

// MoqPtraceDetach_genType holds the state of a moq of the PtraceDetach_genType
// type
type MoqPtraceDetach_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPtraceDetach_genType_mock

	ResultsByParams []MoqPtraceDetach_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Pid moq.ParamIndexing
		}
	}
}

// MoqPtraceDetach_genType_mock isolates the mock interface of the
// PtraceDetach_genType type
type MoqPtraceDetach_genType_mock struct {
	Moq *MoqPtraceDetach_genType
}

// MoqPtraceDetach_genType_params holds the params of the PtraceDetach_genType
// type
type MoqPtraceDetach_genType_params struct{ Pid int }

// MoqPtraceDetach_genType_paramsKey holds the map key params of the
// PtraceDetach_genType type
type MoqPtraceDetach_genType_paramsKey struct {
	Params struct{ Pid int }
	Hashes struct{ Pid hash.Hash }
}

// MoqPtraceDetach_genType_resultsByParams contains the results for a given set
// of parameters for the PtraceDetach_genType type
type MoqPtraceDetach_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPtraceDetach_genType_paramsKey]*MoqPtraceDetach_genType_results
}

// MoqPtraceDetach_genType_doFn defines the type of function needed when
// calling AndDo for the PtraceDetach_genType type
type MoqPtraceDetach_genType_doFn func(pid int)

// MoqPtraceDetach_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the PtraceDetach_genType type
type MoqPtraceDetach_genType_doReturnFn func(pid int) (err error)

// MoqPtraceDetach_genType_results holds the results of the
// PtraceDetach_genType type
type MoqPtraceDetach_genType_results struct {
	Params  MoqPtraceDetach_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqPtraceDetach_genType_doFn
		DoReturnFn MoqPtraceDetach_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPtraceDetach_genType_fnRecorder routes recorded function calls to the
// MoqPtraceDetach_genType moq
type MoqPtraceDetach_genType_fnRecorder struct {
	Params    MoqPtraceDetach_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPtraceDetach_genType_results
	Moq       *MoqPtraceDetach_genType
}

// MoqPtraceDetach_genType_anyParams isolates the any params functions of the
// PtraceDetach_genType type
type MoqPtraceDetach_genType_anyParams struct {
	Recorder *MoqPtraceDetach_genType_fnRecorder
}

// NewMoqPtraceDetach_genType creates a new moq of the PtraceDetach_genType
// type
func NewMoqPtraceDetach_genType(scene *moq.Scene, config *moq.Config) *MoqPtraceDetach_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPtraceDetach_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPtraceDetach_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Pid moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Pid moq.ParamIndexing
		}{
			Pid: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the PtraceDetach_genType type
func (m *MoqPtraceDetach_genType) Mock() PtraceDetach_genType {
	return func(pid int) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqPtraceDetach_genType_mock{Moq: m}
		return moq.Fn(pid)
	}
}

func (m *MoqPtraceDetach_genType_mock) Fn(pid int) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqPtraceDetach_genType_params{
		Pid: pid,
	}
	var results *MoqPtraceDetach_genType_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		paramsKey := m.Moq.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(pid)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(pid)
	}
	return
}

func (m *MoqPtraceDetach_genType) OnCall(pid int) *MoqPtraceDetach_genType_fnRecorder {
	return &MoqPtraceDetach_genType_fnRecorder{
		Params: MoqPtraceDetach_genType_params{
			Pid: pid,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqPtraceDetach_genType_fnRecorder) Any() *MoqPtraceDetach_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqPtraceDetach_genType_anyParams{Recorder: r}
}

func (a *MoqPtraceDetach_genType_anyParams) Pid() *MoqPtraceDetach_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqPtraceDetach_genType_fnRecorder) Seq() *MoqPtraceDetach_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPtraceDetach_genType_fnRecorder) NoSeq() *MoqPtraceDetach_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPtraceDetach_genType_fnRecorder) ReturnResults(err error) *MoqPtraceDetach_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqPtraceDetach_genType_doFn
		DoReturnFn MoqPtraceDetach_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPtraceDetach_genType_fnRecorder) AndDo(fn MoqPtraceDetach_genType_doFn) *MoqPtraceDetach_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPtraceDetach_genType_fnRecorder) DoReturnResults(fn MoqPtraceDetach_genType_doReturnFn) *MoqPtraceDetach_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqPtraceDetach_genType_doFn
		DoReturnFn MoqPtraceDetach_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPtraceDetach_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPtraceDetach_genType_resultsByParams
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqPtraceDetach_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPtraceDetach_genType_paramsKey]*MoqPtraceDetach_genType_results{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqPtraceDetach_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPtraceDetach_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPtraceDetach_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults or DoReturnResults must be called before calling Repeat")
		return nil
	}
	r.Results.Repeat.Repeat(r.Moq.Scene.T, repeaters)
	last := r.Results.Results[len(r.Results.Results)-1]
	for n := 0; n < r.Results.Repeat.ResultCount-1; n++ {
		if r.Sequence {
			last = struct {
				Values     *struct{ Err error }
				Sequence   uint32
				DoFn       MoqPtraceDetach_genType_doFn
				DoReturnFn MoqPtraceDetach_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPtraceDetach_genType) PrettyParams(params MoqPtraceDetach_genType_params) string {
	return fmt.Sprintf("PtraceDetach_genType(%#v)", params.Pid)
}

func (m *MoqPtraceDetach_genType) ParamsKey(params MoqPtraceDetach_genType_params, anyParams uint64) MoqPtraceDetach_genType_paramsKey {
	m.Scene.T.Helper()
	var pidUsed int
	var pidUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Pid == moq.ParamIndexByValue {
			pidUsed = params.Pid
		} else {
			pidUsedHash = hash.DeepHash(params.Pid)
		}
	}
	return MoqPtraceDetach_genType_paramsKey{
		Params: struct{ Pid int }{
			Pid: pidUsed,
		},
		Hashes: struct{ Pid hash.Hash }{
			Pid: pidUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqPtraceDetach_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPtraceDetach_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams(results.Params))
			}
		}
	}
}
