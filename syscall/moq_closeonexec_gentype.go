// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package syscall

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// CloseOnExec_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type CloseOnExec_genType func(fd int)

// MoqCloseOnExec_genType holds the state of a moq of the CloseOnExec_genType
// type
type MoqCloseOnExec_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCloseOnExec_genType_mock

	ResultsByParams []MoqCloseOnExec_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fd moq.ParamIndexing
		}
	}
}

// MoqCloseOnExec_genType_mock isolates the mock interface of the
// CloseOnExec_genType type
type MoqCloseOnExec_genType_mock struct {
	Moq *MoqCloseOnExec_genType
}

// MoqCloseOnExec_genType_params holds the params of the CloseOnExec_genType
// type
type MoqCloseOnExec_genType_params struct{ Fd int }

// MoqCloseOnExec_genType_paramsKey holds the map key params of the
// CloseOnExec_genType type
type MoqCloseOnExec_genType_paramsKey struct {
	Params struct{ Fd int }
	Hashes struct{ Fd hash.Hash }
}

// MoqCloseOnExec_genType_resultsByParams contains the results for a given set
// of parameters for the CloseOnExec_genType type
type MoqCloseOnExec_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCloseOnExec_genType_paramsKey]*MoqCloseOnExec_genType_results
}

// MoqCloseOnExec_genType_doFn defines the type of function needed when calling
// AndDo for the CloseOnExec_genType type
type MoqCloseOnExec_genType_doFn func(fd int)

// MoqCloseOnExec_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the CloseOnExec_genType type
type MoqCloseOnExec_genType_doReturnFn func(fd int)

// MoqCloseOnExec_genType_results holds the results of the CloseOnExec_genType
// type
type MoqCloseOnExec_genType_results struct {
	Params  MoqCloseOnExec_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCloseOnExec_genType_doFn
		DoReturnFn MoqCloseOnExec_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCloseOnExec_genType_fnRecorder routes recorded function calls to the
// MoqCloseOnExec_genType moq
type MoqCloseOnExec_genType_fnRecorder struct {
	Params    MoqCloseOnExec_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCloseOnExec_genType_results
	Moq       *MoqCloseOnExec_genType
}

// MoqCloseOnExec_genType_anyParams isolates the any params functions of the
// CloseOnExec_genType type
type MoqCloseOnExec_genType_anyParams struct {
	Recorder *MoqCloseOnExec_genType_fnRecorder
}

// NewMoqCloseOnExec_genType creates a new moq of the CloseOnExec_genType type
func NewMoqCloseOnExec_genType(scene *moq.Scene, config *moq.Config) *MoqCloseOnExec_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCloseOnExec_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCloseOnExec_genType_mock{},

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

// Mock returns the moq implementation of the CloseOnExec_genType type
func (m *MoqCloseOnExec_genType) Mock() CloseOnExec_genType {
	return func(fd int) { m.Scene.T.Helper(); moq := &MoqCloseOnExec_genType_mock{Moq: m}; moq.Fn(fd) }
}

func (m *MoqCloseOnExec_genType_mock) Fn(fd int) {
	m.Moq.Scene.T.Helper()
	params := MoqCloseOnExec_genType_params{
		Fd: fd,
	}
	var results *MoqCloseOnExec_genType_results
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

	if result.DoReturnFn != nil {
		result.DoReturnFn(fd)
	}
	return
}

func (m *MoqCloseOnExec_genType) OnCall(fd int) *MoqCloseOnExec_genType_fnRecorder {
	return &MoqCloseOnExec_genType_fnRecorder{
		Params: MoqCloseOnExec_genType_params{
			Fd: fd,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqCloseOnExec_genType_fnRecorder) Any() *MoqCloseOnExec_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqCloseOnExec_genType_anyParams{Recorder: r}
}

func (a *MoqCloseOnExec_genType_anyParams) Fd() *MoqCloseOnExec_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqCloseOnExec_genType_fnRecorder) Seq() *MoqCloseOnExec_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCloseOnExec_genType_fnRecorder) NoSeq() *MoqCloseOnExec_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCloseOnExec_genType_fnRecorder) ReturnResults() *MoqCloseOnExec_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCloseOnExec_genType_doFn
		DoReturnFn MoqCloseOnExec_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCloseOnExec_genType_fnRecorder) AndDo(fn MoqCloseOnExec_genType_doFn) *MoqCloseOnExec_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCloseOnExec_genType_fnRecorder) DoReturnResults(fn MoqCloseOnExec_genType_doReturnFn) *MoqCloseOnExec_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqCloseOnExec_genType_doFn
		DoReturnFn MoqCloseOnExec_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCloseOnExec_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCloseOnExec_genType_resultsByParams
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
		results = &MoqCloseOnExec_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCloseOnExec_genType_paramsKey]*MoqCloseOnExec_genType_results{},
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
		r.Results = &MoqCloseOnExec_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCloseOnExec_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCloseOnExec_genType_fnRecorder {
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
				Values     *struct{}
				Sequence   uint32
				DoFn       MoqCloseOnExec_genType_doFn
				DoReturnFn MoqCloseOnExec_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCloseOnExec_genType) PrettyParams(params MoqCloseOnExec_genType_params) string {
	return fmt.Sprintf("CloseOnExec_genType(%#v)", params.Fd)
}

func (m *MoqCloseOnExec_genType) ParamsKey(params MoqCloseOnExec_genType_params, anyParams uint64) MoqCloseOnExec_genType_paramsKey {
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
	return MoqCloseOnExec_genType_paramsKey{
		Params: struct{ Fd int }{
			Fd: fdUsed,
		},
		Hashes: struct{ Fd hash.Hash }{
			Fd: fdUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCloseOnExec_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCloseOnExec_genType) AssertExpectationsMet() {
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
