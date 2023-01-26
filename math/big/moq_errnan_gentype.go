// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package big

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that big.ErrNaN_genType is mocked
// completely
var _ ErrNaN_genType = (*MoqErrNaN_genType_mock)(nil)

// ErrNaN_genType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type ErrNaN_genType interface {
	Error() string
}

// MoqErrNaN_genType holds the state of a moq of the ErrNaN_genType type
type MoqErrNaN_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqErrNaN_genType_mock

	ResultsByParams_Error []MoqErrNaN_genType_Error_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Error struct{}
		}
	}
}

// MoqErrNaN_genType_mock isolates the mock interface of the ErrNaN_genType
// type
type MoqErrNaN_genType_mock struct {
	Moq *MoqErrNaN_genType
}

// MoqErrNaN_genType_recorder isolates the recorder interface of the
// ErrNaN_genType type
type MoqErrNaN_genType_recorder struct {
	Moq *MoqErrNaN_genType
}

// MoqErrNaN_genType_Error_params holds the params of the ErrNaN_genType type
type MoqErrNaN_genType_Error_params struct{}

// MoqErrNaN_genType_Error_paramsKey holds the map key params of the
// ErrNaN_genType type
type MoqErrNaN_genType_Error_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqErrNaN_genType_Error_resultsByParams contains the results for a given set
// of parameters for the ErrNaN_genType type
type MoqErrNaN_genType_Error_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqErrNaN_genType_Error_paramsKey]*MoqErrNaN_genType_Error_results
}

// MoqErrNaN_genType_Error_doFn defines the type of function needed when
// calling AndDo for the ErrNaN_genType type
type MoqErrNaN_genType_Error_doFn func()

// MoqErrNaN_genType_Error_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ErrNaN_genType type
type MoqErrNaN_genType_Error_doReturnFn func() string

// MoqErrNaN_genType_Error_results holds the results of the ErrNaN_genType type
type MoqErrNaN_genType_Error_results struct {
	Params  MoqErrNaN_genType_Error_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqErrNaN_genType_Error_doFn
		DoReturnFn MoqErrNaN_genType_Error_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqErrNaN_genType_Error_fnRecorder routes recorded function calls to the
// MoqErrNaN_genType moq
type MoqErrNaN_genType_Error_fnRecorder struct {
	Params    MoqErrNaN_genType_Error_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqErrNaN_genType_Error_results
	Moq       *MoqErrNaN_genType
}

// MoqErrNaN_genType_Error_anyParams isolates the any params functions of the
// ErrNaN_genType type
type MoqErrNaN_genType_Error_anyParams struct {
	Recorder *MoqErrNaN_genType_Error_fnRecorder
}

// NewMoqErrNaN_genType creates a new moq of the ErrNaN_genType type
func NewMoqErrNaN_genType(scene *moq.Scene, config *moq.Config) *MoqErrNaN_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqErrNaN_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqErrNaN_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Error struct{}
			}
		}{ParameterIndexing: struct {
			Error struct{}
		}{
			Error: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ErrNaN_genType type
func (m *MoqErrNaN_genType) Mock() *MoqErrNaN_genType_mock { return m.Moq }

func (m *MoqErrNaN_genType_mock) Error() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqErrNaN_genType_Error_params{}
	var results *MoqErrNaN_genType_Error_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Error {
		paramsKey := m.Moq.ParamsKey_Error(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Error(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Error(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Error(params))
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

// OnCall returns the recorder implementation of the ErrNaN_genType type
func (m *MoqErrNaN_genType) OnCall() *MoqErrNaN_genType_recorder {
	return &MoqErrNaN_genType_recorder{
		Moq: m,
	}
}

func (m *MoqErrNaN_genType_recorder) Error() *MoqErrNaN_genType_Error_fnRecorder {
	return &MoqErrNaN_genType_Error_fnRecorder{
		Params:   MoqErrNaN_genType_Error_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqErrNaN_genType_Error_fnRecorder) Any() *MoqErrNaN_genType_Error_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	return &MoqErrNaN_genType_Error_anyParams{Recorder: r}
}

func (r *MoqErrNaN_genType_Error_fnRecorder) Seq() *MoqErrNaN_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqErrNaN_genType_Error_fnRecorder) NoSeq() *MoqErrNaN_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Error(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqErrNaN_genType_Error_fnRecorder) ReturnResults(result1 string) *MoqErrNaN_genType_Error_fnRecorder {
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
		DoFn       MoqErrNaN_genType_Error_doFn
		DoReturnFn MoqErrNaN_genType_Error_doReturnFn
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

func (r *MoqErrNaN_genType_Error_fnRecorder) AndDo(fn MoqErrNaN_genType_Error_doFn) *MoqErrNaN_genType_Error_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqErrNaN_genType_Error_fnRecorder) DoReturnResults(fn MoqErrNaN_genType_Error_doReturnFn) *MoqErrNaN_genType_Error_fnRecorder {
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
		DoFn       MoqErrNaN_genType_Error_doFn
		DoReturnFn MoqErrNaN_genType_Error_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqErrNaN_genType_Error_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqErrNaN_genType_Error_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Error {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqErrNaN_genType_Error_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqErrNaN_genType_Error_paramsKey]*MoqErrNaN_genType_Error_results{},
		}
		r.Moq.ResultsByParams_Error = append(r.Moq.ResultsByParams_Error, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Error) {
			copy(r.Moq.ResultsByParams_Error[insertAt+1:], r.Moq.ResultsByParams_Error[insertAt:0])
			r.Moq.ResultsByParams_Error[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Error(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqErrNaN_genType_Error_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqErrNaN_genType_Error_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqErrNaN_genType_Error_fnRecorder {
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
				DoFn       MoqErrNaN_genType_Error_doFn
				DoReturnFn MoqErrNaN_genType_Error_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqErrNaN_genType) PrettyParams_Error(params MoqErrNaN_genType_Error_params) string {
	return fmt.Sprintf("Error()")
}

func (m *MoqErrNaN_genType) ParamsKey_Error(params MoqErrNaN_genType_Error_params, anyParams uint64) MoqErrNaN_genType_Error_paramsKey {
	m.Scene.T.Helper()
	return MoqErrNaN_genType_Error_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqErrNaN_genType) Reset() { m.ResultsByParams_Error = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqErrNaN_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Error {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Error(results.Params))
			}
		}
	}
}
