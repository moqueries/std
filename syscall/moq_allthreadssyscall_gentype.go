// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"syscall"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// AllThreadsSyscall_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type AllThreadsSyscall_genType func(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)

// MoqAllThreadsSyscall_genType holds the state of a moq of the
// AllThreadsSyscall_genType type
type MoqAllThreadsSyscall_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAllThreadsSyscall_genType_mock

	ResultsByParams []MoqAllThreadsSyscall_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Trap moq.ParamIndexing
			A1   moq.ParamIndexing
			A2   moq.ParamIndexing
			A3   moq.ParamIndexing
		}
	}
}

// MoqAllThreadsSyscall_genType_mock isolates the mock interface of the
// AllThreadsSyscall_genType type
type MoqAllThreadsSyscall_genType_mock struct {
	Moq *MoqAllThreadsSyscall_genType
}

// MoqAllThreadsSyscall_genType_params holds the params of the
// AllThreadsSyscall_genType type
type MoqAllThreadsSyscall_genType_params struct{ Trap, A1, A2, A3 uintptr }

// MoqAllThreadsSyscall_genType_paramsKey holds the map key params of the
// AllThreadsSyscall_genType type
type MoqAllThreadsSyscall_genType_paramsKey struct {
	Params struct{ Trap, A1, A2, A3 uintptr }
	Hashes struct{ Trap, A1, A2, A3 hash.Hash }
}

// MoqAllThreadsSyscall_genType_resultsByParams contains the results for a
// given set of parameters for the AllThreadsSyscall_genType type
type MoqAllThreadsSyscall_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAllThreadsSyscall_genType_paramsKey]*MoqAllThreadsSyscall_genType_results
}

// MoqAllThreadsSyscall_genType_doFn defines the type of function needed when
// calling AndDo for the AllThreadsSyscall_genType type
type MoqAllThreadsSyscall_genType_doFn func(trap, a1, a2, a3 uintptr)

// MoqAllThreadsSyscall_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the AllThreadsSyscall_genType type
type MoqAllThreadsSyscall_genType_doReturnFn func(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)

// MoqAllThreadsSyscall_genType_results holds the results of the
// AllThreadsSyscall_genType type
type MoqAllThreadsSyscall_genType_results struct {
	Params  MoqAllThreadsSyscall_genType_params
	Results []struct {
		Values *struct {
			R1, R2 uintptr
			Err    syscall.Errno
		}
		Sequence   uint32
		DoFn       MoqAllThreadsSyscall_genType_doFn
		DoReturnFn MoqAllThreadsSyscall_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAllThreadsSyscall_genType_fnRecorder routes recorded function calls to
// the MoqAllThreadsSyscall_genType moq
type MoqAllThreadsSyscall_genType_fnRecorder struct {
	Params    MoqAllThreadsSyscall_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAllThreadsSyscall_genType_results
	Moq       *MoqAllThreadsSyscall_genType
}

// MoqAllThreadsSyscall_genType_anyParams isolates the any params functions of
// the AllThreadsSyscall_genType type
type MoqAllThreadsSyscall_genType_anyParams struct {
	Recorder *MoqAllThreadsSyscall_genType_fnRecorder
}

// NewMoqAllThreadsSyscall_genType creates a new moq of the
// AllThreadsSyscall_genType type
func NewMoqAllThreadsSyscall_genType(scene *moq.Scene, config *moq.Config) *MoqAllThreadsSyscall_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAllThreadsSyscall_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAllThreadsSyscall_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Trap moq.ParamIndexing
				A1   moq.ParamIndexing
				A2   moq.ParamIndexing
				A3   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Trap moq.ParamIndexing
			A1   moq.ParamIndexing
			A2   moq.ParamIndexing
			A3   moq.ParamIndexing
		}{
			Trap: moq.ParamIndexByValue,
			A1:   moq.ParamIndexByValue,
			A2:   moq.ParamIndexByValue,
			A3:   moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the AllThreadsSyscall_genType type
func (m *MoqAllThreadsSyscall_genType) Mock() AllThreadsSyscall_genType {
	return func(trap, a1, a2, a3 uintptr) (_, _ uintptr, _ syscall.Errno) {
		m.Scene.T.Helper()
		moq := &MoqAllThreadsSyscall_genType_mock{Moq: m}
		return moq.Fn(trap, a1, a2, a3)
	}
}

func (m *MoqAllThreadsSyscall_genType_mock) Fn(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) {
	m.Moq.Scene.T.Helper()
	params := MoqAllThreadsSyscall_genType_params{
		Trap: trap,
		A1:   a1,
		A2:   a2,
		A3:   a3,
	}
	var results *MoqAllThreadsSyscall_genType_results
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
		result.DoFn(trap, a1, a2, a3)
	}

	if result.Values != nil {
		r1 = result.Values.R1
		r2 = result.Values.R2
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		r1, r2, err = result.DoReturnFn(trap, a1, a2, a3)
	}
	return
}

func (m *MoqAllThreadsSyscall_genType) OnCall(trap, a1, a2, a3 uintptr) *MoqAllThreadsSyscall_genType_fnRecorder {
	return &MoqAllThreadsSyscall_genType_fnRecorder{
		Params: MoqAllThreadsSyscall_genType_params{
			Trap: trap,
			A1:   a1,
			A2:   a2,
			A3:   a3,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqAllThreadsSyscall_genType_fnRecorder) Any() *MoqAllThreadsSyscall_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqAllThreadsSyscall_genType_anyParams{Recorder: r}
}

func (a *MoqAllThreadsSyscall_genType_anyParams) Trap() *MoqAllThreadsSyscall_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqAllThreadsSyscall_genType_anyParams) A1() *MoqAllThreadsSyscall_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqAllThreadsSyscall_genType_anyParams) A2() *MoqAllThreadsSyscall_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqAllThreadsSyscall_genType_anyParams) A3() *MoqAllThreadsSyscall_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqAllThreadsSyscall_genType_fnRecorder) Seq() *MoqAllThreadsSyscall_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAllThreadsSyscall_genType_fnRecorder) NoSeq() *MoqAllThreadsSyscall_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAllThreadsSyscall_genType_fnRecorder) ReturnResults(r1, r2 uintptr, err syscall.Errno) *MoqAllThreadsSyscall_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			R1, R2 uintptr
			Err    syscall.Errno
		}
		Sequence   uint32
		DoFn       MoqAllThreadsSyscall_genType_doFn
		DoReturnFn MoqAllThreadsSyscall_genType_doReturnFn
	}{
		Values: &struct {
			R1, R2 uintptr
			Err    syscall.Errno
		}{
			R1:  r1,
			R2:  r2,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqAllThreadsSyscall_genType_fnRecorder) AndDo(fn MoqAllThreadsSyscall_genType_doFn) *MoqAllThreadsSyscall_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAllThreadsSyscall_genType_fnRecorder) DoReturnResults(fn MoqAllThreadsSyscall_genType_doReturnFn) *MoqAllThreadsSyscall_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			R1, R2 uintptr
			Err    syscall.Errno
		}
		Sequence   uint32
		DoFn       MoqAllThreadsSyscall_genType_doFn
		DoReturnFn MoqAllThreadsSyscall_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAllThreadsSyscall_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAllThreadsSyscall_genType_resultsByParams
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
		results = &MoqAllThreadsSyscall_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAllThreadsSyscall_genType_paramsKey]*MoqAllThreadsSyscall_genType_results{},
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
		r.Results = &MoqAllThreadsSyscall_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAllThreadsSyscall_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAllThreadsSyscall_genType_fnRecorder {
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
				Values *struct {
					R1, R2 uintptr
					Err    syscall.Errno
				}
				Sequence   uint32
				DoFn       MoqAllThreadsSyscall_genType_doFn
				DoReturnFn MoqAllThreadsSyscall_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAllThreadsSyscall_genType) PrettyParams(params MoqAllThreadsSyscall_genType_params) string {
	return fmt.Sprintf("AllThreadsSyscall_genType(%#v, %#v, %#v, %#v)", params.Trap, params.A1, params.A2, params.A3)
}

func (m *MoqAllThreadsSyscall_genType) ParamsKey(params MoqAllThreadsSyscall_genType_params, anyParams uint64) MoqAllThreadsSyscall_genType_paramsKey {
	m.Scene.T.Helper()
	var trapUsed uintptr
	var trapUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Trap == moq.ParamIndexByValue {
			trapUsed = params.Trap
		} else {
			trapUsedHash = hash.DeepHash(params.Trap)
		}
	}
	var a1Used uintptr
	var a1UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.A1 == moq.ParamIndexByValue {
			a1Used = params.A1
		} else {
			a1UsedHash = hash.DeepHash(params.A1)
		}
	}
	var a2Used uintptr
	var a2UsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.A2 == moq.ParamIndexByValue {
			a2Used = params.A2
		} else {
			a2UsedHash = hash.DeepHash(params.A2)
		}
	}
	var a3Used uintptr
	var a3UsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.A3 == moq.ParamIndexByValue {
			a3Used = params.A3
		} else {
			a3UsedHash = hash.DeepHash(params.A3)
		}
	}
	return MoqAllThreadsSyscall_genType_paramsKey{
		Params: struct{ Trap, A1, A2, A3 uintptr }{
			Trap: trapUsed,
			A1:   a1Used,
			A2:   a2Used,
			A3:   a3Used,
		},
		Hashes: struct{ Trap, A1, A2, A3 hash.Hash }{
			Trap: trapUsedHash,
			A1:   a1UsedHash,
			A2:   a2UsedHash,
			A3:   a3UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqAllThreadsSyscall_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAllThreadsSyscall_genType) AssertExpectationsMet() {
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
