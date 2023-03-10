// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package ast

import (
	"fmt"
	"go/token"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that ast.Object_starGenType is mocked
// completely
var _ Object_starGenType = (*MoqObject_starGenType_mock)(nil)

// Object_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Object_starGenType interface {
	Pos() token.Pos
}

// MoqObject_starGenType holds the state of a moq of the Object_starGenType
// type
type MoqObject_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqObject_starGenType_mock

	ResultsByParams_Pos []MoqObject_starGenType_Pos_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Pos struct{}
		}
	}
}

// MoqObject_starGenType_mock isolates the mock interface of the
// Object_starGenType type
type MoqObject_starGenType_mock struct {
	Moq *MoqObject_starGenType
}

// MoqObject_starGenType_recorder isolates the recorder interface of the
// Object_starGenType type
type MoqObject_starGenType_recorder struct {
	Moq *MoqObject_starGenType
}

// MoqObject_starGenType_Pos_params holds the params of the Object_starGenType
// type
type MoqObject_starGenType_Pos_params struct{}

// MoqObject_starGenType_Pos_paramsKey holds the map key params of the
// Object_starGenType type
type MoqObject_starGenType_Pos_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqObject_starGenType_Pos_resultsByParams contains the results for a given
// set of parameters for the Object_starGenType type
type MoqObject_starGenType_Pos_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqObject_starGenType_Pos_paramsKey]*MoqObject_starGenType_Pos_results
}

// MoqObject_starGenType_Pos_doFn defines the type of function needed when
// calling AndDo for the Object_starGenType type
type MoqObject_starGenType_Pos_doFn func()

// MoqObject_starGenType_Pos_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Object_starGenType type
type MoqObject_starGenType_Pos_doReturnFn func() token.Pos

// MoqObject_starGenType_Pos_results holds the results of the
// Object_starGenType type
type MoqObject_starGenType_Pos_results struct {
	Params  MoqObject_starGenType_Pos_params
	Results []struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqObject_starGenType_Pos_doFn
		DoReturnFn MoqObject_starGenType_Pos_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqObject_starGenType_Pos_fnRecorder routes recorded function calls to the
// MoqObject_starGenType moq
type MoqObject_starGenType_Pos_fnRecorder struct {
	Params    MoqObject_starGenType_Pos_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqObject_starGenType_Pos_results
	Moq       *MoqObject_starGenType
}

// MoqObject_starGenType_Pos_anyParams isolates the any params functions of the
// Object_starGenType type
type MoqObject_starGenType_Pos_anyParams struct {
	Recorder *MoqObject_starGenType_Pos_fnRecorder
}

// NewMoqObject_starGenType creates a new moq of the Object_starGenType type
func NewMoqObject_starGenType(scene *moq.Scene, config *moq.Config) *MoqObject_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqObject_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqObject_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Pos struct{}
			}
		}{ParameterIndexing: struct {
			Pos struct{}
		}{
			Pos: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Object_starGenType type
func (m *MoqObject_starGenType) Mock() *MoqObject_starGenType_mock { return m.Moq }

func (m *MoqObject_starGenType_mock) Pos() (result1 token.Pos) {
	m.Moq.Scene.T.Helper()
	params := MoqObject_starGenType_Pos_params{}
	var results *MoqObject_starGenType_Pos_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Pos {
		paramsKey := m.Moq.ParamsKey_Pos(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Pos(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Pos(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Pos(params))
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

// OnCall returns the recorder implementation of the Object_starGenType type
func (m *MoqObject_starGenType) OnCall() *MoqObject_starGenType_recorder {
	return &MoqObject_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqObject_starGenType_recorder) Pos() *MoqObject_starGenType_Pos_fnRecorder {
	return &MoqObject_starGenType_Pos_fnRecorder{
		Params:   MoqObject_starGenType_Pos_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqObject_starGenType_Pos_fnRecorder) Any() *MoqObject_starGenType_Pos_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Pos(r.Params))
		return nil
	}
	return &MoqObject_starGenType_Pos_anyParams{Recorder: r}
}

func (r *MoqObject_starGenType_Pos_fnRecorder) Seq() *MoqObject_starGenType_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Pos(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqObject_starGenType_Pos_fnRecorder) NoSeq() *MoqObject_starGenType_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Pos(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqObject_starGenType_Pos_fnRecorder) ReturnResults(result1 token.Pos) *MoqObject_starGenType_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqObject_starGenType_Pos_doFn
		DoReturnFn MoqObject_starGenType_Pos_doReturnFn
	}{
		Values: &struct {
			Result1 token.Pos
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqObject_starGenType_Pos_fnRecorder) AndDo(fn MoqObject_starGenType_Pos_doFn) *MoqObject_starGenType_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqObject_starGenType_Pos_fnRecorder) DoReturnResults(fn MoqObject_starGenType_Pos_doReturnFn) *MoqObject_starGenType_Pos_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 token.Pos
		}
		Sequence   uint32
		DoFn       MoqObject_starGenType_Pos_doFn
		DoReturnFn MoqObject_starGenType_Pos_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqObject_starGenType_Pos_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqObject_starGenType_Pos_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Pos {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqObject_starGenType_Pos_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqObject_starGenType_Pos_paramsKey]*MoqObject_starGenType_Pos_results{},
		}
		r.Moq.ResultsByParams_Pos = append(r.Moq.ResultsByParams_Pos, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Pos) {
			copy(r.Moq.ResultsByParams_Pos[insertAt+1:], r.Moq.ResultsByParams_Pos[insertAt:0])
			r.Moq.ResultsByParams_Pos[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Pos(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqObject_starGenType_Pos_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqObject_starGenType_Pos_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqObject_starGenType_Pos_fnRecorder {
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
					Result1 token.Pos
				}
				Sequence   uint32
				DoFn       MoqObject_starGenType_Pos_doFn
				DoReturnFn MoqObject_starGenType_Pos_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqObject_starGenType) PrettyParams_Pos(params MoqObject_starGenType_Pos_params) string {
	return fmt.Sprintf("Pos()")
}

func (m *MoqObject_starGenType) ParamsKey_Pos(params MoqObject_starGenType_Pos_params, anyParams uint64) MoqObject_starGenType_Pos_paramsKey {
	m.Scene.T.Helper()
	return MoqObject_starGenType_Pos_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqObject_starGenType) Reset() { m.ResultsByParams_Pos = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqObject_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Pos {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Pos(results.Params))
			}
		}
	}
}
