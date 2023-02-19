// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package x509

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that x509.PublicKeyAlgorithm_genType is
// mocked completely
var _ PublicKeyAlgorithm_genType = (*MoqPublicKeyAlgorithm_genType_mock)(nil)

// PublicKeyAlgorithm_genType is the fabricated implementation type of this
// mock (emitted when mocking a collections of methods directly and not from an
// interface type)
type PublicKeyAlgorithm_genType interface {
	String() string
}

// MoqPublicKeyAlgorithm_genType holds the state of a moq of the
// PublicKeyAlgorithm_genType type
type MoqPublicKeyAlgorithm_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPublicKeyAlgorithm_genType_mock

	ResultsByParams_String []MoqPublicKeyAlgorithm_genType_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			String struct{}
		}
	}
}

// MoqPublicKeyAlgorithm_genType_mock isolates the mock interface of the
// PublicKeyAlgorithm_genType type
type MoqPublicKeyAlgorithm_genType_mock struct {
	Moq *MoqPublicKeyAlgorithm_genType
}

// MoqPublicKeyAlgorithm_genType_recorder isolates the recorder interface of
// the PublicKeyAlgorithm_genType type
type MoqPublicKeyAlgorithm_genType_recorder struct {
	Moq *MoqPublicKeyAlgorithm_genType
}

// MoqPublicKeyAlgorithm_genType_String_params holds the params of the
// PublicKeyAlgorithm_genType type
type MoqPublicKeyAlgorithm_genType_String_params struct{}

// MoqPublicKeyAlgorithm_genType_String_paramsKey holds the map key params of
// the PublicKeyAlgorithm_genType type
type MoqPublicKeyAlgorithm_genType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqPublicKeyAlgorithm_genType_String_resultsByParams contains the results
// for a given set of parameters for the PublicKeyAlgorithm_genType type
type MoqPublicKeyAlgorithm_genType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPublicKeyAlgorithm_genType_String_paramsKey]*MoqPublicKeyAlgorithm_genType_String_results
}

// MoqPublicKeyAlgorithm_genType_String_doFn defines the type of function
// needed when calling AndDo for the PublicKeyAlgorithm_genType type
type MoqPublicKeyAlgorithm_genType_String_doFn func()

// MoqPublicKeyAlgorithm_genType_String_doReturnFn defines the type of function
// needed when calling DoReturnResults for the PublicKeyAlgorithm_genType type
type MoqPublicKeyAlgorithm_genType_String_doReturnFn func() string

// MoqPublicKeyAlgorithm_genType_String_results holds the results of the
// PublicKeyAlgorithm_genType type
type MoqPublicKeyAlgorithm_genType_String_results struct {
	Params  MoqPublicKeyAlgorithm_genType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqPublicKeyAlgorithm_genType_String_doFn
		DoReturnFn MoqPublicKeyAlgorithm_genType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPublicKeyAlgorithm_genType_String_fnRecorder routes recorded function
// calls to the MoqPublicKeyAlgorithm_genType moq
type MoqPublicKeyAlgorithm_genType_String_fnRecorder struct {
	Params    MoqPublicKeyAlgorithm_genType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPublicKeyAlgorithm_genType_String_results
	Moq       *MoqPublicKeyAlgorithm_genType
}

// MoqPublicKeyAlgorithm_genType_String_anyParams isolates the any params
// functions of the PublicKeyAlgorithm_genType type
type MoqPublicKeyAlgorithm_genType_String_anyParams struct {
	Recorder *MoqPublicKeyAlgorithm_genType_String_fnRecorder
}

// NewMoqPublicKeyAlgorithm_genType creates a new moq of the
// PublicKeyAlgorithm_genType type
func NewMoqPublicKeyAlgorithm_genType(scene *moq.Scene, config *moq.Config) *MoqPublicKeyAlgorithm_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPublicKeyAlgorithm_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPublicKeyAlgorithm_genType_mock{},

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

// Mock returns the mock implementation of the PublicKeyAlgorithm_genType type
func (m *MoqPublicKeyAlgorithm_genType) Mock() *MoqPublicKeyAlgorithm_genType_mock { return m.Moq }

func (m *MoqPublicKeyAlgorithm_genType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqPublicKeyAlgorithm_genType_String_params{}
	var results *MoqPublicKeyAlgorithm_genType_String_results
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

// OnCall returns the recorder implementation of the PublicKeyAlgorithm_genType
// type
func (m *MoqPublicKeyAlgorithm_genType) OnCall() *MoqPublicKeyAlgorithm_genType_recorder {
	return &MoqPublicKeyAlgorithm_genType_recorder{
		Moq: m,
	}
}

func (m *MoqPublicKeyAlgorithm_genType_recorder) String() *MoqPublicKeyAlgorithm_genType_String_fnRecorder {
	return &MoqPublicKeyAlgorithm_genType_String_fnRecorder{
		Params:   MoqPublicKeyAlgorithm_genType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqPublicKeyAlgorithm_genType_String_fnRecorder) Any() *MoqPublicKeyAlgorithm_genType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqPublicKeyAlgorithm_genType_String_anyParams{Recorder: r}
}

func (r *MoqPublicKeyAlgorithm_genType_String_fnRecorder) Seq() *MoqPublicKeyAlgorithm_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPublicKeyAlgorithm_genType_String_fnRecorder) NoSeq() *MoqPublicKeyAlgorithm_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPublicKeyAlgorithm_genType_String_fnRecorder) ReturnResults(result1 string) *MoqPublicKeyAlgorithm_genType_String_fnRecorder {
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
		DoFn       MoqPublicKeyAlgorithm_genType_String_doFn
		DoReturnFn MoqPublicKeyAlgorithm_genType_String_doReturnFn
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

func (r *MoqPublicKeyAlgorithm_genType_String_fnRecorder) AndDo(fn MoqPublicKeyAlgorithm_genType_String_doFn) *MoqPublicKeyAlgorithm_genType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPublicKeyAlgorithm_genType_String_fnRecorder) DoReturnResults(fn MoqPublicKeyAlgorithm_genType_String_doReturnFn) *MoqPublicKeyAlgorithm_genType_String_fnRecorder {
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
		DoFn       MoqPublicKeyAlgorithm_genType_String_doFn
		DoReturnFn MoqPublicKeyAlgorithm_genType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPublicKeyAlgorithm_genType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPublicKeyAlgorithm_genType_String_resultsByParams
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
		results = &MoqPublicKeyAlgorithm_genType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPublicKeyAlgorithm_genType_String_paramsKey]*MoqPublicKeyAlgorithm_genType_String_results{},
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
		r.Results = &MoqPublicKeyAlgorithm_genType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPublicKeyAlgorithm_genType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPublicKeyAlgorithm_genType_String_fnRecorder {
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
				DoFn       MoqPublicKeyAlgorithm_genType_String_doFn
				DoReturnFn MoqPublicKeyAlgorithm_genType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPublicKeyAlgorithm_genType) PrettyParams_String(params MoqPublicKeyAlgorithm_genType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqPublicKeyAlgorithm_genType) ParamsKey_String(params MoqPublicKeyAlgorithm_genType_String_params, anyParams uint64) MoqPublicKeyAlgorithm_genType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqPublicKeyAlgorithm_genType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqPublicKeyAlgorithm_genType) Reset() { m.ResultsByParams_String = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPublicKeyAlgorithm_genType) AssertExpectationsMet() {
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
