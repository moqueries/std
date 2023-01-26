// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package os

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// UserConfigDir_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type UserConfigDir_genType func() (string, error)

// MoqUserConfigDir_genType holds the state of a moq of the
// UserConfigDir_genType type
type MoqUserConfigDir_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqUserConfigDir_genType_mock

	ResultsByParams []MoqUserConfigDir_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct{}
	}
}

// MoqUserConfigDir_genType_mock isolates the mock interface of the
// UserConfigDir_genType type
type MoqUserConfigDir_genType_mock struct {
	Moq *MoqUserConfigDir_genType
}

// MoqUserConfigDir_genType_params holds the params of the
// UserConfigDir_genType type
type MoqUserConfigDir_genType_params struct{}

// MoqUserConfigDir_genType_paramsKey holds the map key params of the
// UserConfigDir_genType type
type MoqUserConfigDir_genType_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqUserConfigDir_genType_resultsByParams contains the results for a given
// set of parameters for the UserConfigDir_genType type
type MoqUserConfigDir_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUserConfigDir_genType_paramsKey]*MoqUserConfigDir_genType_results
}

// MoqUserConfigDir_genType_doFn defines the type of function needed when
// calling AndDo for the UserConfigDir_genType type
type MoqUserConfigDir_genType_doFn func()

// MoqUserConfigDir_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the UserConfigDir_genType type
type MoqUserConfigDir_genType_doReturnFn func() (string, error)

// MoqUserConfigDir_genType_results holds the results of the
// UserConfigDir_genType type
type MoqUserConfigDir_genType_results struct {
	Params  MoqUserConfigDir_genType_params
	Results []struct {
		Values *struct {
			Result1 string
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqUserConfigDir_genType_doFn
		DoReturnFn MoqUserConfigDir_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUserConfigDir_genType_fnRecorder routes recorded function calls to the
// MoqUserConfigDir_genType moq
type MoqUserConfigDir_genType_fnRecorder struct {
	Params    MoqUserConfigDir_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUserConfigDir_genType_results
	Moq       *MoqUserConfigDir_genType
}

// MoqUserConfigDir_genType_anyParams isolates the any params functions of the
// UserConfigDir_genType type
type MoqUserConfigDir_genType_anyParams struct {
	Recorder *MoqUserConfigDir_genType_fnRecorder
}

// NewMoqUserConfigDir_genType creates a new moq of the UserConfigDir_genType
// type
func NewMoqUserConfigDir_genType(scene *moq.Scene, config *moq.Config) *MoqUserConfigDir_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqUserConfigDir_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqUserConfigDir_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct{}
		}{ParameterIndexing: struct{}{}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the UserConfigDir_genType type
func (m *MoqUserConfigDir_genType) Mock() UserConfigDir_genType {
	return func() (string, error) {
		m.Scene.T.Helper()
		moq := &MoqUserConfigDir_genType_mock{Moq: m}
		return moq.Fn()
	}
}

func (m *MoqUserConfigDir_genType_mock) Fn() (result1 string, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqUserConfigDir_genType_params{}
	var results *MoqUserConfigDir_genType_results
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
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

func (m *MoqUserConfigDir_genType) OnCall() *MoqUserConfigDir_genType_fnRecorder {
	return &MoqUserConfigDir_genType_fnRecorder{
		Params:   MoqUserConfigDir_genType_params{},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqUserConfigDir_genType_fnRecorder) Any() *MoqUserConfigDir_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqUserConfigDir_genType_anyParams{Recorder: r}
}

func (r *MoqUserConfigDir_genType_fnRecorder) Seq() *MoqUserConfigDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUserConfigDir_genType_fnRecorder) NoSeq() *MoqUserConfigDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUserConfigDir_genType_fnRecorder) ReturnResults(result1 string, result2 error) *MoqUserConfigDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqUserConfigDir_genType_doFn
		DoReturnFn MoqUserConfigDir_genType_doReturnFn
	}{
		Values: &struct {
			Result1 string
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUserConfigDir_genType_fnRecorder) AndDo(fn MoqUserConfigDir_genType_doFn) *MoqUserConfigDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUserConfigDir_genType_fnRecorder) DoReturnResults(fn MoqUserConfigDir_genType_doReturnFn) *MoqUserConfigDir_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqUserConfigDir_genType_doFn
		DoReturnFn MoqUserConfigDir_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUserConfigDir_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUserConfigDir_genType_resultsByParams
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
		results = &MoqUserConfigDir_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUserConfigDir_genType_paramsKey]*MoqUserConfigDir_genType_results{},
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
		r.Results = &MoqUserConfigDir_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUserConfigDir_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUserConfigDir_genType_fnRecorder {
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
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqUserConfigDir_genType_doFn
				DoReturnFn MoqUserConfigDir_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUserConfigDir_genType) PrettyParams(params MoqUserConfigDir_genType_params) string {
	return fmt.Sprintf("UserConfigDir_genType()")
}

func (m *MoqUserConfigDir_genType) ParamsKey(params MoqUserConfigDir_genType_params, anyParams uint64) MoqUserConfigDir_genType_paramsKey {
	m.Scene.T.Helper()
	return MoqUserConfigDir_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqUserConfigDir_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUserConfigDir_genType) AssertExpectationsMet() {
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