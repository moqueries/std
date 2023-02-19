// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package macho

import (
	"debug/macho"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that macho.Load is mocked completely
var _ macho.Load = (*MoqLoad_mock)(nil)

// MoqLoad holds the state of a moq of the Load type
type MoqLoad struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqLoad_mock

	ResultsByParams_Raw []MoqLoad_Raw_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Raw struct{}
		}
	}
}

// MoqLoad_mock isolates the mock interface of the Load type
type MoqLoad_mock struct {
	Moq *MoqLoad
}

// MoqLoad_recorder isolates the recorder interface of the Load type
type MoqLoad_recorder struct {
	Moq *MoqLoad
}

// MoqLoad_Raw_params holds the params of the Load type
type MoqLoad_Raw_params struct{}

// MoqLoad_Raw_paramsKey holds the map key params of the Load type
type MoqLoad_Raw_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqLoad_Raw_resultsByParams contains the results for a given set of
// parameters for the Load type
type MoqLoad_Raw_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqLoad_Raw_paramsKey]*MoqLoad_Raw_results
}

// MoqLoad_Raw_doFn defines the type of function needed when calling AndDo for
// the Load type
type MoqLoad_Raw_doFn func()

// MoqLoad_Raw_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Load type
type MoqLoad_Raw_doReturnFn func() []byte

// MoqLoad_Raw_results holds the results of the Load type
type MoqLoad_Raw_results struct {
	Params  MoqLoad_Raw_params
	Results []struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqLoad_Raw_doFn
		DoReturnFn MoqLoad_Raw_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqLoad_Raw_fnRecorder routes recorded function calls to the MoqLoad moq
type MoqLoad_Raw_fnRecorder struct {
	Params    MoqLoad_Raw_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqLoad_Raw_results
	Moq       *MoqLoad
}

// MoqLoad_Raw_anyParams isolates the any params functions of the Load type
type MoqLoad_Raw_anyParams struct {
	Recorder *MoqLoad_Raw_fnRecorder
}

// NewMoqLoad creates a new moq of the Load type
func NewMoqLoad(scene *moq.Scene, config *moq.Config) *MoqLoad {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqLoad{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqLoad_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Raw struct{}
			}
		}{ParameterIndexing: struct {
			Raw struct{}
		}{
			Raw: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Load type
func (m *MoqLoad) Mock() *MoqLoad_mock { return m.Moq }

func (m *MoqLoad_mock) Raw() (result1 []byte) {
	m.Moq.Scene.T.Helper()
	params := MoqLoad_Raw_params{}
	var results *MoqLoad_Raw_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Raw {
		paramsKey := m.Moq.ParamsKey_Raw(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Raw(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Raw(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Raw(params))
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

// OnCall returns the recorder implementation of the Load type
func (m *MoqLoad) OnCall() *MoqLoad_recorder {
	return &MoqLoad_recorder{
		Moq: m,
	}
}

func (m *MoqLoad_recorder) Raw() *MoqLoad_Raw_fnRecorder {
	return &MoqLoad_Raw_fnRecorder{
		Params:   MoqLoad_Raw_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqLoad_Raw_fnRecorder) Any() *MoqLoad_Raw_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Raw(r.Params))
		return nil
	}
	return &MoqLoad_Raw_anyParams{Recorder: r}
}

func (r *MoqLoad_Raw_fnRecorder) Seq() *MoqLoad_Raw_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Raw(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqLoad_Raw_fnRecorder) NoSeq() *MoqLoad_Raw_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Raw(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqLoad_Raw_fnRecorder) ReturnResults(result1 []byte) *MoqLoad_Raw_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqLoad_Raw_doFn
		DoReturnFn MoqLoad_Raw_doReturnFn
	}{
		Values: &struct {
			Result1 []byte
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqLoad_Raw_fnRecorder) AndDo(fn MoqLoad_Raw_doFn) *MoqLoad_Raw_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqLoad_Raw_fnRecorder) DoReturnResults(fn MoqLoad_Raw_doReturnFn) *MoqLoad_Raw_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqLoad_Raw_doFn
		DoReturnFn MoqLoad_Raw_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqLoad_Raw_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqLoad_Raw_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Raw {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqLoad_Raw_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqLoad_Raw_paramsKey]*MoqLoad_Raw_results{},
		}
		r.Moq.ResultsByParams_Raw = append(r.Moq.ResultsByParams_Raw, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Raw) {
			copy(r.Moq.ResultsByParams_Raw[insertAt+1:], r.Moq.ResultsByParams_Raw[insertAt:0])
			r.Moq.ResultsByParams_Raw[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Raw(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqLoad_Raw_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqLoad_Raw_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqLoad_Raw_fnRecorder {
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
					Result1 []byte
				}
				Sequence   uint32
				DoFn       MoqLoad_Raw_doFn
				DoReturnFn MoqLoad_Raw_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqLoad) PrettyParams_Raw(params MoqLoad_Raw_params) string { return fmt.Sprintf("Raw()") }

func (m *MoqLoad) ParamsKey_Raw(params MoqLoad_Raw_params, anyParams uint64) MoqLoad_Raw_paramsKey {
	m.Scene.T.Helper()
	return MoqLoad_Raw_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqLoad) Reset() { m.ResultsByParams_Raw = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqLoad) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Raw {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Raw(results.Params))
			}
		}
	}
}