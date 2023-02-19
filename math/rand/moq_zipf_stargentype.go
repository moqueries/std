// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package rand

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that rand.Zipf_starGenType is mocked
// completely
var _ Zipf_starGenType = (*MoqZipf_starGenType_mock)(nil)

// Zipf_starGenType is the fabricated implementation type of this mock (emitted
// when mocking a collections of methods directly and not from an interface
// type)
type Zipf_starGenType interface {
	Uint64() uint64
}

// MoqZipf_starGenType holds the state of a moq of the Zipf_starGenType type
type MoqZipf_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqZipf_starGenType_mock

	ResultsByParams_Uint64 []MoqZipf_starGenType_Uint64_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Uint64 struct{}
		}
	}
}

// MoqZipf_starGenType_mock isolates the mock interface of the Zipf_starGenType
// type
type MoqZipf_starGenType_mock struct {
	Moq *MoqZipf_starGenType
}

// MoqZipf_starGenType_recorder isolates the recorder interface of the
// Zipf_starGenType type
type MoqZipf_starGenType_recorder struct {
	Moq *MoqZipf_starGenType
}

// MoqZipf_starGenType_Uint64_params holds the params of the Zipf_starGenType
// type
type MoqZipf_starGenType_Uint64_params struct{}

// MoqZipf_starGenType_Uint64_paramsKey holds the map key params of the
// Zipf_starGenType type
type MoqZipf_starGenType_Uint64_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqZipf_starGenType_Uint64_resultsByParams contains the results for a given
// set of parameters for the Zipf_starGenType type
type MoqZipf_starGenType_Uint64_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqZipf_starGenType_Uint64_paramsKey]*MoqZipf_starGenType_Uint64_results
}

// MoqZipf_starGenType_Uint64_doFn defines the type of function needed when
// calling AndDo for the Zipf_starGenType type
type MoqZipf_starGenType_Uint64_doFn func()

// MoqZipf_starGenType_Uint64_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Zipf_starGenType type
type MoqZipf_starGenType_Uint64_doReturnFn func() uint64

// MoqZipf_starGenType_Uint64_results holds the results of the Zipf_starGenType
// type
type MoqZipf_starGenType_Uint64_results struct {
	Params  MoqZipf_starGenType_Uint64_params
	Results []struct {
		Values *struct {
			Result1 uint64
		}
		Sequence   uint32
		DoFn       MoqZipf_starGenType_Uint64_doFn
		DoReturnFn MoqZipf_starGenType_Uint64_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqZipf_starGenType_Uint64_fnRecorder routes recorded function calls to the
// MoqZipf_starGenType moq
type MoqZipf_starGenType_Uint64_fnRecorder struct {
	Params    MoqZipf_starGenType_Uint64_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqZipf_starGenType_Uint64_results
	Moq       *MoqZipf_starGenType
}

// MoqZipf_starGenType_Uint64_anyParams isolates the any params functions of
// the Zipf_starGenType type
type MoqZipf_starGenType_Uint64_anyParams struct {
	Recorder *MoqZipf_starGenType_Uint64_fnRecorder
}

// NewMoqZipf_starGenType creates a new moq of the Zipf_starGenType type
func NewMoqZipf_starGenType(scene *moq.Scene, config *moq.Config) *MoqZipf_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqZipf_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqZipf_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Uint64 struct{}
			}
		}{ParameterIndexing: struct {
			Uint64 struct{}
		}{
			Uint64: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Zipf_starGenType type
func (m *MoqZipf_starGenType) Mock() *MoqZipf_starGenType_mock { return m.Moq }

func (m *MoqZipf_starGenType_mock) Uint64() (result1 uint64) {
	m.Moq.Scene.T.Helper()
	params := MoqZipf_starGenType_Uint64_params{}
	var results *MoqZipf_starGenType_Uint64_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Uint64 {
		paramsKey := m.Moq.ParamsKey_Uint64(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Uint64(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Uint64(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Uint64(params))
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

// OnCall returns the recorder implementation of the Zipf_starGenType type
func (m *MoqZipf_starGenType) OnCall() *MoqZipf_starGenType_recorder {
	return &MoqZipf_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqZipf_starGenType_recorder) Uint64() *MoqZipf_starGenType_Uint64_fnRecorder {
	return &MoqZipf_starGenType_Uint64_fnRecorder{
		Params:   MoqZipf_starGenType_Uint64_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqZipf_starGenType_Uint64_fnRecorder) Any() *MoqZipf_starGenType_Uint64_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Uint64(r.Params))
		return nil
	}
	return &MoqZipf_starGenType_Uint64_anyParams{Recorder: r}
}

func (r *MoqZipf_starGenType_Uint64_fnRecorder) Seq() *MoqZipf_starGenType_Uint64_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Uint64(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqZipf_starGenType_Uint64_fnRecorder) NoSeq() *MoqZipf_starGenType_Uint64_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Uint64(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqZipf_starGenType_Uint64_fnRecorder) ReturnResults(result1 uint64) *MoqZipf_starGenType_Uint64_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 uint64
		}
		Sequence   uint32
		DoFn       MoqZipf_starGenType_Uint64_doFn
		DoReturnFn MoqZipf_starGenType_Uint64_doReturnFn
	}{
		Values: &struct {
			Result1 uint64
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqZipf_starGenType_Uint64_fnRecorder) AndDo(fn MoqZipf_starGenType_Uint64_doFn) *MoqZipf_starGenType_Uint64_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqZipf_starGenType_Uint64_fnRecorder) DoReturnResults(fn MoqZipf_starGenType_Uint64_doReturnFn) *MoqZipf_starGenType_Uint64_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 uint64
		}
		Sequence   uint32
		DoFn       MoqZipf_starGenType_Uint64_doFn
		DoReturnFn MoqZipf_starGenType_Uint64_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqZipf_starGenType_Uint64_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqZipf_starGenType_Uint64_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Uint64 {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqZipf_starGenType_Uint64_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqZipf_starGenType_Uint64_paramsKey]*MoqZipf_starGenType_Uint64_results{},
		}
		r.Moq.ResultsByParams_Uint64 = append(r.Moq.ResultsByParams_Uint64, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Uint64) {
			copy(r.Moq.ResultsByParams_Uint64[insertAt+1:], r.Moq.ResultsByParams_Uint64[insertAt:0])
			r.Moq.ResultsByParams_Uint64[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Uint64(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqZipf_starGenType_Uint64_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqZipf_starGenType_Uint64_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqZipf_starGenType_Uint64_fnRecorder {
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
					Result1 uint64
				}
				Sequence   uint32
				DoFn       MoqZipf_starGenType_Uint64_doFn
				DoReturnFn MoqZipf_starGenType_Uint64_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqZipf_starGenType) PrettyParams_Uint64(params MoqZipf_starGenType_Uint64_params) string {
	return fmt.Sprintf("Uint64()")
}

func (m *MoqZipf_starGenType) ParamsKey_Uint64(params MoqZipf_starGenType_Uint64_params, anyParams uint64) MoqZipf_starGenType_Uint64_paramsKey {
	m.Scene.T.Helper()
	return MoqZipf_starGenType_Uint64_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqZipf_starGenType) Reset() { m.ResultsByParams_Uint64 = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqZipf_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Uint64 {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Uint64(results.Params))
			}
		}
	}
}
