// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package reflect

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that reflect.ChanDir_genType is mocked
// completely
var _ ChanDir_genType = (*MoqChanDir_genType_mock)(nil)

// ChanDir_genType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type ChanDir_genType interface {
	String() string
}

// MoqChanDir_genType holds the state of a moq of the ChanDir_genType type
type MoqChanDir_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqChanDir_genType_mock

	ResultsByParams_String []MoqChanDir_genType_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			String struct{}
		}
	}
}

// MoqChanDir_genType_mock isolates the mock interface of the ChanDir_genType
// type
type MoqChanDir_genType_mock struct {
	Moq *MoqChanDir_genType
}

// MoqChanDir_genType_recorder isolates the recorder interface of the
// ChanDir_genType type
type MoqChanDir_genType_recorder struct {
	Moq *MoqChanDir_genType
}

// MoqChanDir_genType_String_params holds the params of the ChanDir_genType
// type
type MoqChanDir_genType_String_params struct{}

// MoqChanDir_genType_String_paramsKey holds the map key params of the
// ChanDir_genType type
type MoqChanDir_genType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqChanDir_genType_String_resultsByParams contains the results for a given
// set of parameters for the ChanDir_genType type
type MoqChanDir_genType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqChanDir_genType_String_paramsKey]*MoqChanDir_genType_String_results
}

// MoqChanDir_genType_String_doFn defines the type of function needed when
// calling AndDo for the ChanDir_genType type
type MoqChanDir_genType_String_doFn func()

// MoqChanDir_genType_String_doReturnFn defines the type of function needed
// when calling DoReturnResults for the ChanDir_genType type
type MoqChanDir_genType_String_doReturnFn func() string

// MoqChanDir_genType_String_results holds the results of the ChanDir_genType
// type
type MoqChanDir_genType_String_results struct {
	Params  MoqChanDir_genType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqChanDir_genType_String_doFn
		DoReturnFn MoqChanDir_genType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqChanDir_genType_String_fnRecorder routes recorded function calls to the
// MoqChanDir_genType moq
type MoqChanDir_genType_String_fnRecorder struct {
	Params    MoqChanDir_genType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqChanDir_genType_String_results
	Moq       *MoqChanDir_genType
}

// MoqChanDir_genType_String_anyParams isolates the any params functions of the
// ChanDir_genType type
type MoqChanDir_genType_String_anyParams struct {
	Recorder *MoqChanDir_genType_String_fnRecorder
}

// NewMoqChanDir_genType creates a new moq of the ChanDir_genType type
func NewMoqChanDir_genType(scene *moq.Scene, config *moq.Config) *MoqChanDir_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqChanDir_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqChanDir_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				String struct{}
			}
		}{ParameterIndexing: struct {
			String struct{}
		}{
			String: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ChanDir_genType type
func (m *MoqChanDir_genType) Mock() *MoqChanDir_genType_mock { return m.Moq }

func (m *MoqChanDir_genType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqChanDir_genType_String_params{}
	var results *MoqChanDir_genType_String_results
	for _, resultsByParams := range m.Moq.ResultsByParams_String {
		paramsKey := m.Moq.ParamsKey_String(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_String(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_String(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_String(params))
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

// OnCall returns the recorder implementation of the ChanDir_genType type
func (m *MoqChanDir_genType) OnCall() *MoqChanDir_genType_recorder {
	return &MoqChanDir_genType_recorder{
		Moq: m,
	}
}

func (m *MoqChanDir_genType_recorder) String() *MoqChanDir_genType_String_fnRecorder {
	return &MoqChanDir_genType_String_fnRecorder{
		Params:   MoqChanDir_genType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqChanDir_genType_String_fnRecorder) Any() *MoqChanDir_genType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqChanDir_genType_String_anyParams{Recorder: r}
}

func (r *MoqChanDir_genType_String_fnRecorder) Seq() *MoqChanDir_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqChanDir_genType_String_fnRecorder) NoSeq() *MoqChanDir_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqChanDir_genType_String_fnRecorder) ReturnResults(result1 string) *MoqChanDir_genType_String_fnRecorder {
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
		DoFn       MoqChanDir_genType_String_doFn
		DoReturnFn MoqChanDir_genType_String_doReturnFn
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

func (r *MoqChanDir_genType_String_fnRecorder) AndDo(fn MoqChanDir_genType_String_doFn) *MoqChanDir_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqChanDir_genType_String_fnRecorder) DoReturnResults(fn MoqChanDir_genType_String_doReturnFn) *MoqChanDir_genType_String_fnRecorder {
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
		DoFn       MoqChanDir_genType_String_doFn
		DoReturnFn MoqChanDir_genType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqChanDir_genType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqChanDir_genType_String_resultsByParams
	for n, res := range r.Moq.ResultsByParams_String {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqChanDir_genType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqChanDir_genType_String_paramsKey]*MoqChanDir_genType_String_results{},
		}
		r.Moq.ResultsByParams_String = append(r.Moq.ResultsByParams_String, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_String) {
			copy(r.Moq.ResultsByParams_String[insertAt+1:], r.Moq.ResultsByParams_String[insertAt:0])
			r.Moq.ResultsByParams_String[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_String(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqChanDir_genType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqChanDir_genType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqChanDir_genType_String_fnRecorder {
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
				DoFn       MoqChanDir_genType_String_doFn
				DoReturnFn MoqChanDir_genType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqChanDir_genType) PrettyParams_String(params MoqChanDir_genType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqChanDir_genType) ParamsKey_String(params MoqChanDir_genType_String_params, anyParams uint64) MoqChanDir_genType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqChanDir_genType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqChanDir_genType) Reset() { m.ResultsByParams_String = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqChanDir_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_String {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_String(results.Params))
			}
		}
	}
}
