// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package macho

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that macho.LoadBytes_genType is mocked
// completely
var _ LoadBytes_genType = (*MoqLoadBytes_genType_mock)(nil)

// LoadBytes_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type LoadBytes_genType interface {
	Raw() []byte
}

// MoqLoadBytes_genType holds the state of a moq of the LoadBytes_genType type
type MoqLoadBytes_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqLoadBytes_genType_mock

	ResultsByParams_Raw []MoqLoadBytes_genType_Raw_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Raw struct{}
		}
	}
}

// MoqLoadBytes_genType_mock isolates the mock interface of the
// LoadBytes_genType type
type MoqLoadBytes_genType_mock struct {
	Moq *MoqLoadBytes_genType
}

// MoqLoadBytes_genType_recorder isolates the recorder interface of the
// LoadBytes_genType type
type MoqLoadBytes_genType_recorder struct {
	Moq *MoqLoadBytes_genType
}

// MoqLoadBytes_genType_Raw_params holds the params of the LoadBytes_genType
// type
type MoqLoadBytes_genType_Raw_params struct{}

// MoqLoadBytes_genType_Raw_paramsKey holds the map key params of the
// LoadBytes_genType type
type MoqLoadBytes_genType_Raw_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqLoadBytes_genType_Raw_resultsByParams contains the results for a given
// set of parameters for the LoadBytes_genType type
type MoqLoadBytes_genType_Raw_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqLoadBytes_genType_Raw_paramsKey]*MoqLoadBytes_genType_Raw_results
}

// MoqLoadBytes_genType_Raw_doFn defines the type of function needed when
// calling AndDo for the LoadBytes_genType type
type MoqLoadBytes_genType_Raw_doFn func()

// MoqLoadBytes_genType_Raw_doReturnFn defines the type of function needed when
// calling DoReturnResults for the LoadBytes_genType type
type MoqLoadBytes_genType_Raw_doReturnFn func() []byte

// MoqLoadBytes_genType_Raw_results holds the results of the LoadBytes_genType
// type
type MoqLoadBytes_genType_Raw_results struct {
	Params  MoqLoadBytes_genType_Raw_params
	Results []struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqLoadBytes_genType_Raw_doFn
		DoReturnFn MoqLoadBytes_genType_Raw_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqLoadBytes_genType_Raw_fnRecorder routes recorded function calls to the
// MoqLoadBytes_genType moq
type MoqLoadBytes_genType_Raw_fnRecorder struct {
	Params    MoqLoadBytes_genType_Raw_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqLoadBytes_genType_Raw_results
	Moq       *MoqLoadBytes_genType
}

// MoqLoadBytes_genType_Raw_anyParams isolates the any params functions of the
// LoadBytes_genType type
type MoqLoadBytes_genType_Raw_anyParams struct {
	Recorder *MoqLoadBytes_genType_Raw_fnRecorder
}

// NewMoqLoadBytes_genType creates a new moq of the LoadBytes_genType type
func NewMoqLoadBytes_genType(scene *moq.Scene, config *moq.Config) *MoqLoadBytes_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqLoadBytes_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqLoadBytes_genType_mock{},

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

// Mock returns the mock implementation of the LoadBytes_genType type
func (m *MoqLoadBytes_genType) Mock() *MoqLoadBytes_genType_mock { return m.Moq }

func (m *MoqLoadBytes_genType_mock) Raw() (result1 []byte) {
	m.Moq.Scene.T.Helper()
	params := MoqLoadBytes_genType_Raw_params{}
	var results *MoqLoadBytes_genType_Raw_results
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

// OnCall returns the recorder implementation of the LoadBytes_genType type
func (m *MoqLoadBytes_genType) OnCall() *MoqLoadBytes_genType_recorder {
	return &MoqLoadBytes_genType_recorder{
		Moq: m,
	}
}

func (m *MoqLoadBytes_genType_recorder) Raw() *MoqLoadBytes_genType_Raw_fnRecorder {
	return &MoqLoadBytes_genType_Raw_fnRecorder{
		Params:   MoqLoadBytes_genType_Raw_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqLoadBytes_genType_Raw_fnRecorder) Any() *MoqLoadBytes_genType_Raw_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Raw(r.Params))
		return nil
	}
	return &MoqLoadBytes_genType_Raw_anyParams{Recorder: r}
}

func (r *MoqLoadBytes_genType_Raw_fnRecorder) Seq() *MoqLoadBytes_genType_Raw_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Raw(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqLoadBytes_genType_Raw_fnRecorder) NoSeq() *MoqLoadBytes_genType_Raw_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Raw(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqLoadBytes_genType_Raw_fnRecorder) ReturnResults(result1 []byte) *MoqLoadBytes_genType_Raw_fnRecorder {
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
		DoFn       MoqLoadBytes_genType_Raw_doFn
		DoReturnFn MoqLoadBytes_genType_Raw_doReturnFn
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

func (r *MoqLoadBytes_genType_Raw_fnRecorder) AndDo(fn MoqLoadBytes_genType_Raw_doFn) *MoqLoadBytes_genType_Raw_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqLoadBytes_genType_Raw_fnRecorder) DoReturnResults(fn MoqLoadBytes_genType_Raw_doReturnFn) *MoqLoadBytes_genType_Raw_fnRecorder {
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
		DoFn       MoqLoadBytes_genType_Raw_doFn
		DoReturnFn MoqLoadBytes_genType_Raw_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqLoadBytes_genType_Raw_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqLoadBytes_genType_Raw_resultsByParams
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
		results = &MoqLoadBytes_genType_Raw_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqLoadBytes_genType_Raw_paramsKey]*MoqLoadBytes_genType_Raw_results{},
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
		r.Results = &MoqLoadBytes_genType_Raw_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqLoadBytes_genType_Raw_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqLoadBytes_genType_Raw_fnRecorder {
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
				DoFn       MoqLoadBytes_genType_Raw_doFn
				DoReturnFn MoqLoadBytes_genType_Raw_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqLoadBytes_genType) PrettyParams_Raw(params MoqLoadBytes_genType_Raw_params) string {
	return fmt.Sprintf("Raw()")
}

func (m *MoqLoadBytes_genType) ParamsKey_Raw(params MoqLoadBytes_genType_Raw_params, anyParams uint64) MoqLoadBytes_genType_Raw_paramsKey {
	m.Scene.T.Helper()
	return MoqLoadBytes_genType_Raw_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqLoadBytes_genType) Reset() { m.ResultsByParams_Raw = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqLoadBytes_genType) AssertExpectationsMet() {
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