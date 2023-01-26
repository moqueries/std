// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package http

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that http.Cookie_starGenType is mocked
// completely
var _ Cookie_starGenType = (*MoqCookie_starGenType_mock)(nil)

// Cookie_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Cookie_starGenType interface {
	String() string
}

// MoqCookie_starGenType holds the state of a moq of the Cookie_starGenType
// type
type MoqCookie_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCookie_starGenType_mock

	ResultsByParams_String []MoqCookie_starGenType_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			String struct{}
		}
	}
}

// MoqCookie_starGenType_mock isolates the mock interface of the
// Cookie_starGenType type
type MoqCookie_starGenType_mock struct {
	Moq *MoqCookie_starGenType
}

// MoqCookie_starGenType_recorder isolates the recorder interface of the
// Cookie_starGenType type
type MoqCookie_starGenType_recorder struct {
	Moq *MoqCookie_starGenType
}

// MoqCookie_starGenType_String_params holds the params of the
// Cookie_starGenType type
type MoqCookie_starGenType_String_params struct{}

// MoqCookie_starGenType_String_paramsKey holds the map key params of the
// Cookie_starGenType type
type MoqCookie_starGenType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqCookie_starGenType_String_resultsByParams contains the results for a
// given set of parameters for the Cookie_starGenType type
type MoqCookie_starGenType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCookie_starGenType_String_paramsKey]*MoqCookie_starGenType_String_results
}

// MoqCookie_starGenType_String_doFn defines the type of function needed when
// calling AndDo for the Cookie_starGenType type
type MoqCookie_starGenType_String_doFn func()

// MoqCookie_starGenType_String_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Cookie_starGenType type
type MoqCookie_starGenType_String_doReturnFn func() string

// MoqCookie_starGenType_String_results holds the results of the
// Cookie_starGenType type
type MoqCookie_starGenType_String_results struct {
	Params  MoqCookie_starGenType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqCookie_starGenType_String_doFn
		DoReturnFn MoqCookie_starGenType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCookie_starGenType_String_fnRecorder routes recorded function calls to
// the MoqCookie_starGenType moq
type MoqCookie_starGenType_String_fnRecorder struct {
	Params    MoqCookie_starGenType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCookie_starGenType_String_results
	Moq       *MoqCookie_starGenType
}

// MoqCookie_starGenType_String_anyParams isolates the any params functions of
// the Cookie_starGenType type
type MoqCookie_starGenType_String_anyParams struct {
	Recorder *MoqCookie_starGenType_String_fnRecorder
}

// NewMoqCookie_starGenType creates a new moq of the Cookie_starGenType type
func NewMoqCookie_starGenType(scene *moq.Scene, config *moq.Config) *MoqCookie_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCookie_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCookie_starGenType_mock{},

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

// Mock returns the mock implementation of the Cookie_starGenType type
func (m *MoqCookie_starGenType) Mock() *MoqCookie_starGenType_mock { return m.Moq }

func (m *MoqCookie_starGenType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqCookie_starGenType_String_params{}
	var results *MoqCookie_starGenType_String_results
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

// OnCall returns the recorder implementation of the Cookie_starGenType type
func (m *MoqCookie_starGenType) OnCall() *MoqCookie_starGenType_recorder {
	return &MoqCookie_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqCookie_starGenType_recorder) String() *MoqCookie_starGenType_String_fnRecorder {
	return &MoqCookie_starGenType_String_fnRecorder{
		Params:   MoqCookie_starGenType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqCookie_starGenType_String_fnRecorder) Any() *MoqCookie_starGenType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqCookie_starGenType_String_anyParams{Recorder: r}
}

func (r *MoqCookie_starGenType_String_fnRecorder) Seq() *MoqCookie_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCookie_starGenType_String_fnRecorder) NoSeq() *MoqCookie_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCookie_starGenType_String_fnRecorder) ReturnResults(result1 string) *MoqCookie_starGenType_String_fnRecorder {
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
		DoFn       MoqCookie_starGenType_String_doFn
		DoReturnFn MoqCookie_starGenType_String_doReturnFn
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

func (r *MoqCookie_starGenType_String_fnRecorder) AndDo(fn MoqCookie_starGenType_String_doFn) *MoqCookie_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCookie_starGenType_String_fnRecorder) DoReturnResults(fn MoqCookie_starGenType_String_doReturnFn) *MoqCookie_starGenType_String_fnRecorder {
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
		DoFn       MoqCookie_starGenType_String_doFn
		DoReturnFn MoqCookie_starGenType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCookie_starGenType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCookie_starGenType_String_resultsByParams
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
		results = &MoqCookie_starGenType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCookie_starGenType_String_paramsKey]*MoqCookie_starGenType_String_results{},
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
		r.Results = &MoqCookie_starGenType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCookie_starGenType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCookie_starGenType_String_fnRecorder {
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
				DoFn       MoqCookie_starGenType_String_doFn
				DoReturnFn MoqCookie_starGenType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCookie_starGenType) PrettyParams_String(params MoqCookie_starGenType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqCookie_starGenType) ParamsKey_String(params MoqCookie_starGenType_String_params, anyParams uint64) MoqCookie_starGenType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqCookie_starGenType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqCookie_starGenType) Reset() { m.ResultsByParams_String = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCookie_starGenType) AssertExpectationsMet() {
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
