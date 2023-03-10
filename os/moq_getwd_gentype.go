// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package os

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// Getwd_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Getwd_genType func() (dir string, err error)

// MoqGetwd_genType holds the state of a moq of the Getwd_genType type
type MoqGetwd_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqGetwd_genType_mock

	ResultsByParams []MoqGetwd_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct{}
	}
}

// MoqGetwd_genType_mock isolates the mock interface of the Getwd_genType type
type MoqGetwd_genType_mock struct {
	Moq *MoqGetwd_genType
}

// MoqGetwd_genType_params holds the params of the Getwd_genType type
type MoqGetwd_genType_params struct{}

// MoqGetwd_genType_paramsKey holds the map key params of the Getwd_genType
// type
type MoqGetwd_genType_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqGetwd_genType_resultsByParams contains the results for a given set of
// parameters for the Getwd_genType type
type MoqGetwd_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqGetwd_genType_paramsKey]*MoqGetwd_genType_results
}

// MoqGetwd_genType_doFn defines the type of function needed when calling AndDo
// for the Getwd_genType type
type MoqGetwd_genType_doFn func()

// MoqGetwd_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Getwd_genType type
type MoqGetwd_genType_doReturnFn func() (dir string, err error)

// MoqGetwd_genType_results holds the results of the Getwd_genType type
type MoqGetwd_genType_results struct {
	Params  MoqGetwd_genType_params
	Results []struct {
		Values *struct {
			Dir string
			Err error
		}
		Sequence   uint32
		DoFn       MoqGetwd_genType_doFn
		DoReturnFn MoqGetwd_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqGetwd_genType_fnRecorder routes recorded function calls to the
// MoqGetwd_genType moq
type MoqGetwd_genType_fnRecorder struct {
	Params    MoqGetwd_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqGetwd_genType_results
	Moq       *MoqGetwd_genType
}

// MoqGetwd_genType_anyParams isolates the any params functions of the
// Getwd_genType type
type MoqGetwd_genType_anyParams struct {
	Recorder *MoqGetwd_genType_fnRecorder
}

// NewMoqGetwd_genType creates a new moq of the Getwd_genType type
func NewMoqGetwd_genType(scene *moq.Scene, config *moq.Config) *MoqGetwd_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqGetwd_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqGetwd_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct{}
		}{ParameterIndexing: struct{}{}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Getwd_genType type
func (m *MoqGetwd_genType) Mock() Getwd_genType {
	return func() (_ string, _ error) { m.Scene.T.Helper(); moq := &MoqGetwd_genType_mock{Moq: m}; return moq.Fn() }
}

func (m *MoqGetwd_genType_mock) Fn() (dir string, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqGetwd_genType_params{}
	var results *MoqGetwd_genType_results
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
		dir = result.Values.Dir
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		dir, err = result.DoReturnFn()
	}
	return
}

func (m *MoqGetwd_genType) OnCall() *MoqGetwd_genType_fnRecorder {
	return &MoqGetwd_genType_fnRecorder{
		Params:   MoqGetwd_genType_params{},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqGetwd_genType_fnRecorder) Any() *MoqGetwd_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqGetwd_genType_anyParams{Recorder: r}
}

func (r *MoqGetwd_genType_fnRecorder) Seq() *MoqGetwd_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqGetwd_genType_fnRecorder) NoSeq() *MoqGetwd_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqGetwd_genType_fnRecorder) ReturnResults(dir string, err error) *MoqGetwd_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Dir string
			Err error
		}
		Sequence   uint32
		DoFn       MoqGetwd_genType_doFn
		DoReturnFn MoqGetwd_genType_doReturnFn
	}{
		Values: &struct {
			Dir string
			Err error
		}{
			Dir: dir,
			Err: err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqGetwd_genType_fnRecorder) AndDo(fn MoqGetwd_genType_doFn) *MoqGetwd_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqGetwd_genType_fnRecorder) DoReturnResults(fn MoqGetwd_genType_doReturnFn) *MoqGetwd_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Dir string
			Err error
		}
		Sequence   uint32
		DoFn       MoqGetwd_genType_doFn
		DoReturnFn MoqGetwd_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqGetwd_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqGetwd_genType_resultsByParams
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
		results = &MoqGetwd_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqGetwd_genType_paramsKey]*MoqGetwd_genType_results{},
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
		r.Results = &MoqGetwd_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqGetwd_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqGetwd_genType_fnRecorder {
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
					Dir string
					Err error
				}
				Sequence   uint32
				DoFn       MoqGetwd_genType_doFn
				DoReturnFn MoqGetwd_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqGetwd_genType) PrettyParams(params MoqGetwd_genType_params) string {
	return fmt.Sprintf("Getwd_genType()")
}

func (m *MoqGetwd_genType) ParamsKey(params MoqGetwd_genType_params, anyParams uint64) MoqGetwd_genType_paramsKey {
	m.Scene.T.Helper()
	return MoqGetwd_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqGetwd_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqGetwd_genType) AssertExpectationsMet() {
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
