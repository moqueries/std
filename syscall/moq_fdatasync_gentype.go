// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Fdatasync_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type Fdatasync_genType func(fd int) (err error)

// MoqFdatasync_genType holds the state of a moq of the Fdatasync_genType type
type MoqFdatasync_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqFdatasync_genType_mock

	ResultsByParams []MoqFdatasync_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fd moq.ParamIndexing
		}
	}
}

// MoqFdatasync_genType_mock isolates the mock interface of the
// Fdatasync_genType type
type MoqFdatasync_genType_mock struct {
	Moq *MoqFdatasync_genType
}

// MoqFdatasync_genType_params holds the params of the Fdatasync_genType type
type MoqFdatasync_genType_params struct{ Fd int }

// MoqFdatasync_genType_paramsKey holds the map key params of the
// Fdatasync_genType type
type MoqFdatasync_genType_paramsKey struct {
	Params struct{ Fd int }
	Hashes struct{ Fd hash.Hash }
}

// MoqFdatasync_genType_resultsByParams contains the results for a given set of
// parameters for the Fdatasync_genType type
type MoqFdatasync_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqFdatasync_genType_paramsKey]*MoqFdatasync_genType_results
}

// MoqFdatasync_genType_doFn defines the type of function needed when calling
// AndDo for the Fdatasync_genType type
type MoqFdatasync_genType_doFn func(fd int)

// MoqFdatasync_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Fdatasync_genType type
type MoqFdatasync_genType_doReturnFn func(fd int) (err error)

// MoqFdatasync_genType_results holds the results of the Fdatasync_genType type
type MoqFdatasync_genType_results struct {
	Params  MoqFdatasync_genType_params
	Results []struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqFdatasync_genType_doFn
		DoReturnFn MoqFdatasync_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqFdatasync_genType_fnRecorder routes recorded function calls to the
// MoqFdatasync_genType moq
type MoqFdatasync_genType_fnRecorder struct {
	Params    MoqFdatasync_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqFdatasync_genType_results
	Moq       *MoqFdatasync_genType
}

// MoqFdatasync_genType_anyParams isolates the any params functions of the
// Fdatasync_genType type
type MoqFdatasync_genType_anyParams struct {
	Recorder *MoqFdatasync_genType_fnRecorder
}

// NewMoqFdatasync_genType creates a new moq of the Fdatasync_genType type
func NewMoqFdatasync_genType(scene *moq.Scene, config *moq.Config) *MoqFdatasync_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqFdatasync_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqFdatasync_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fd moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fd moq.ParamIndexing
		}{
			Fd: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Fdatasync_genType type
func (m *MoqFdatasync_genType) Mock() Fdatasync_genType {
	return func(fd int) (_ error) {
		m.Scene.T.Helper()
		moq := &MoqFdatasync_genType_mock{Moq: m}
		return moq.Fn(fd)
	}
}

func (m *MoqFdatasync_genType_mock) Fn(fd int) (err error) {
	m.Moq.Scene.T.Helper()
	params := MoqFdatasync_genType_params{
		Fd: fd,
	}
	var results *MoqFdatasync_genType_results
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
		result.DoFn(fd)
	}

	if result.Values != nil {
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		err = result.DoReturnFn(fd)
	}
	return
}

func (m *MoqFdatasync_genType) OnCall(fd int) *MoqFdatasync_genType_fnRecorder {
	return &MoqFdatasync_genType_fnRecorder{
		Params: MoqFdatasync_genType_params{
			Fd: fd,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqFdatasync_genType_fnRecorder) Any() *MoqFdatasync_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqFdatasync_genType_anyParams{Recorder: r}
}

func (a *MoqFdatasync_genType_anyParams) Fd() *MoqFdatasync_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqFdatasync_genType_fnRecorder) Seq() *MoqFdatasync_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqFdatasync_genType_fnRecorder) NoSeq() *MoqFdatasync_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqFdatasync_genType_fnRecorder) ReturnResults(err error) *MoqFdatasync_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqFdatasync_genType_doFn
		DoReturnFn MoqFdatasync_genType_doReturnFn
	}{
		Values: &struct{ Err error }{
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqFdatasync_genType_fnRecorder) AndDo(fn MoqFdatasync_genType_doFn) *MoqFdatasync_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqFdatasync_genType_fnRecorder) DoReturnResults(fn MoqFdatasync_genType_doReturnFn) *MoqFdatasync_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Err error }
		Sequence   uint32
		DoFn       MoqFdatasync_genType_doFn
		DoReturnFn MoqFdatasync_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqFdatasync_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqFdatasync_genType_resultsByParams
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
		results = &MoqFdatasync_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqFdatasync_genType_paramsKey]*MoqFdatasync_genType_results{},
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
		r.Results = &MoqFdatasync_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqFdatasync_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqFdatasync_genType_fnRecorder {
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
				DoFn       MoqFdatasync_genType_doFn
				DoReturnFn MoqFdatasync_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqFdatasync_genType) PrettyParams(params MoqFdatasync_genType_params) string {
	return fmt.Sprintf("Fdatasync_genType(%#v)", params.Fd)
}

func (m *MoqFdatasync_genType) ParamsKey(params MoqFdatasync_genType_params, anyParams uint64) MoqFdatasync_genType_paramsKey {
	m.Scene.T.Helper()
	var fdUsed int
	var fdUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Fd == moq.ParamIndexByValue {
			fdUsed = params.Fd
		} else {
			fdUsedHash = hash.DeepHash(params.Fd)
		}
	}
	return MoqFdatasync_genType_paramsKey{
		Params: struct{ Fd int }{
			Fd: fdUsed,
		},
		Hashes: struct{ Fd hash.Hash }{
			Fd: fdUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqFdatasync_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqFdatasync_genType) AssertExpectationsMet() {
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