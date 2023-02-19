// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package os

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// TempDir_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type TempDir_genType func() string

// MoqTempDir_genType holds the state of a moq of the TempDir_genType type
type MoqTempDir_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqTempDir_genType_mock

	ResultsByParams []MoqTempDir_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct{}
	}
}

// MoqTempDir_genType_mock isolates the mock interface of the TempDir_genType
// type
type MoqTempDir_genType_mock struct {
	Moq *MoqTempDir_genType
}

// MoqTempDir_genType_params holds the params of the TempDir_genType type
type MoqTempDir_genType_params struct{}

// MoqTempDir_genType_paramsKey holds the map key params of the TempDir_genType
// type
type MoqTempDir_genType_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqTempDir_genType_resultsByParams contains the results for a given set of
// parameters for the TempDir_genType type
type MoqTempDir_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqTempDir_genType_paramsKey]*MoqTempDir_genType_results
}

// MoqTempDir_genType_doFn defines the type of function needed when calling
// AndDo for the TempDir_genType type
type MoqTempDir_genType_doFn func()

// MoqTempDir_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the TempDir_genType type
type MoqTempDir_genType_doReturnFn func() string

// MoqTempDir_genType_results holds the results of the TempDir_genType type
type MoqTempDir_genType_results struct {
	Params  MoqTempDir_genType_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqTempDir_genType_doFn
		DoReturnFn MoqTempDir_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqTempDir_genType_fnRecorder routes recorded function calls to the
// MoqTempDir_genType moq
type MoqTempDir_genType_fnRecorder struct {
	Params    MoqTempDir_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqTempDir_genType_results
	Moq       *MoqTempDir_genType
}

// MoqTempDir_genType_anyParams isolates the any params functions of the
// TempDir_genType type
type MoqTempDir_genType_anyParams struct {
	Recorder *MoqTempDir_genType_fnRecorder
}

// NewMoqTempDir_genType creates a new moq of the TempDir_genType type
func NewMoqTempDir_genType(scene *moq.Scene, config *moq.Config) *MoqTempDir_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqTempDir_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqTempDir_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct{}
		}{ParameterIndexing: struct{}{}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the TempDir_genType type
func (m *MoqTempDir_genType) Mock() TempDir_genType {
	return func() string { m.Scene.T.Helper(); moq := &MoqTempDir_genType_mock{Moq: m}; return moq.Fn() }
}

func (m *MoqTempDir_genType_mock) Fn() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqTempDir_genType_params{}
	var results *MoqTempDir_genType_results
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
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqTempDir_genType) OnCall() *MoqTempDir_genType_fnRecorder {
	return &MoqTempDir_genType_fnRecorder{
		Params:   MoqTempDir_genType_params{},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqTempDir_genType_fnRecorder) Any() *MoqTempDir_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqTempDir_genType_anyParams{Recorder: r}
}

func (r *MoqTempDir_genType_fnRecorder) Seq() *MoqTempDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqTempDir_genType_fnRecorder) NoSeq() *MoqTempDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqTempDir_genType_fnRecorder) ReturnResults(result1 string) *MoqTempDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqTempDir_genType_doFn
		DoReturnFn MoqTempDir_genType_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqTempDir_genType_fnRecorder) AndDo(fn MoqTempDir_genType_doFn) *MoqTempDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqTempDir_genType_fnRecorder) DoReturnResults(fn MoqTempDir_genType_doReturnFn) *MoqTempDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqTempDir_genType_doFn
		DoReturnFn MoqTempDir_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqTempDir_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqTempDir_genType_resultsByParams
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
		results = &MoqTempDir_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqTempDir_genType_paramsKey]*MoqTempDir_genType_results{},
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
		r.Results = &MoqTempDir_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqTempDir_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqTempDir_genType_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqTempDir_genType_doFn
				DoReturnFn MoqTempDir_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqTempDir_genType) PrettyParams(params MoqTempDir_genType_params) string {
	return fmt.Sprintf("TempDir_genType()")
}

func (m *MoqTempDir_genType) ParamsKey(params MoqTempDir_genType_params, anyParams uint64) MoqTempDir_genType_paramsKey {
	m.Scene.T.Helper()
	return MoqTempDir_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqTempDir_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqTempDir_genType) AssertExpectationsMet() {
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