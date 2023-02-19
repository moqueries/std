// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package rsa

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that rsa.PublicKey_starGenType is
// mocked completely
var _ PublicKey_starGenType = (*MoqPublicKey_starGenType_mock)(nil)

// PublicKey_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type PublicKey_starGenType interface {
	Size() int
}

// MoqPublicKey_starGenType holds the state of a moq of the
// PublicKey_starGenType type
type MoqPublicKey_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPublicKey_starGenType_mock

	ResultsByParams_Size []MoqPublicKey_starGenType_Size_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Size struct{}
		}
	}
}

// MoqPublicKey_starGenType_mock isolates the mock interface of the
// PublicKey_starGenType type
type MoqPublicKey_starGenType_mock struct {
	Moq *MoqPublicKey_starGenType
}

// MoqPublicKey_starGenType_recorder isolates the recorder interface of the
// PublicKey_starGenType type
type MoqPublicKey_starGenType_recorder struct {
	Moq *MoqPublicKey_starGenType
}

// MoqPublicKey_starGenType_Size_params holds the params of the
// PublicKey_starGenType type
type MoqPublicKey_starGenType_Size_params struct{}

// MoqPublicKey_starGenType_Size_paramsKey holds the map key params of the
// PublicKey_starGenType type
type MoqPublicKey_starGenType_Size_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqPublicKey_starGenType_Size_resultsByParams contains the results for a
// given set of parameters for the PublicKey_starGenType type
type MoqPublicKey_starGenType_Size_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPublicKey_starGenType_Size_paramsKey]*MoqPublicKey_starGenType_Size_results
}

// MoqPublicKey_starGenType_Size_doFn defines the type of function needed when
// calling AndDo for the PublicKey_starGenType type
type MoqPublicKey_starGenType_Size_doFn func()

// MoqPublicKey_starGenType_Size_doReturnFn defines the type of function needed
// when calling DoReturnResults for the PublicKey_starGenType type
type MoqPublicKey_starGenType_Size_doReturnFn func() int

// MoqPublicKey_starGenType_Size_results holds the results of the
// PublicKey_starGenType type
type MoqPublicKey_starGenType_Size_results struct {
	Params  MoqPublicKey_starGenType_Size_params
	Results []struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqPublicKey_starGenType_Size_doFn
		DoReturnFn MoqPublicKey_starGenType_Size_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPublicKey_starGenType_Size_fnRecorder routes recorded function calls to
// the MoqPublicKey_starGenType moq
type MoqPublicKey_starGenType_Size_fnRecorder struct {
	Params    MoqPublicKey_starGenType_Size_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPublicKey_starGenType_Size_results
	Moq       *MoqPublicKey_starGenType
}

// MoqPublicKey_starGenType_Size_anyParams isolates the any params functions of
// the PublicKey_starGenType type
type MoqPublicKey_starGenType_Size_anyParams struct {
	Recorder *MoqPublicKey_starGenType_Size_fnRecorder
}

// NewMoqPublicKey_starGenType creates a new moq of the PublicKey_starGenType
// type
func NewMoqPublicKey_starGenType(scene *moq.Scene, config *moq.Config) *MoqPublicKey_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPublicKey_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPublicKey_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Size struct{}
			}
		}{ParameterIndexing: struct {
			Size struct{}
		}{
			Size: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the PublicKey_starGenType type
func (m *MoqPublicKey_starGenType) Mock() *MoqPublicKey_starGenType_mock { return m.Moq }

func (m *MoqPublicKey_starGenType_mock) Size() (result1 int) {
	m.Moq.Scene.T.Helper()
	params := MoqPublicKey_starGenType_Size_params{}
	var results *MoqPublicKey_starGenType_Size_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Size {
		paramsKey := m.Moq.ParamsKey_Size(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Size(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Size(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Size(params))
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

// OnCall returns the recorder implementation of the PublicKey_starGenType type
func (m *MoqPublicKey_starGenType) OnCall() *MoqPublicKey_starGenType_recorder {
	return &MoqPublicKey_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqPublicKey_starGenType_recorder) Size() *MoqPublicKey_starGenType_Size_fnRecorder {
	return &MoqPublicKey_starGenType_Size_fnRecorder{
		Params:   MoqPublicKey_starGenType_Size_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqPublicKey_starGenType_Size_fnRecorder) Any() *MoqPublicKey_starGenType_Size_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Size(r.Params))
		return nil
	}
	return &MoqPublicKey_starGenType_Size_anyParams{Recorder: r}
}

func (r *MoqPublicKey_starGenType_Size_fnRecorder) Seq() *MoqPublicKey_starGenType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Size(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPublicKey_starGenType_Size_fnRecorder) NoSeq() *MoqPublicKey_starGenType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Size(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPublicKey_starGenType_Size_fnRecorder) ReturnResults(result1 int) *MoqPublicKey_starGenType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqPublicKey_starGenType_Size_doFn
		DoReturnFn MoqPublicKey_starGenType_Size_doReturnFn
	}{
		Values: &struct {
			Result1 int
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPublicKey_starGenType_Size_fnRecorder) AndDo(fn MoqPublicKey_starGenType_Size_doFn) *MoqPublicKey_starGenType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPublicKey_starGenType_Size_fnRecorder) DoReturnResults(fn MoqPublicKey_starGenType_Size_doReturnFn) *MoqPublicKey_starGenType_Size_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqPublicKey_starGenType_Size_doFn
		DoReturnFn MoqPublicKey_starGenType_Size_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPublicKey_starGenType_Size_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPublicKey_starGenType_Size_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Size {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqPublicKey_starGenType_Size_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPublicKey_starGenType_Size_paramsKey]*MoqPublicKey_starGenType_Size_results{},
		}
		r.Moq.ResultsByParams_Size = append(r.Moq.ResultsByParams_Size, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Size) {
			copy(r.Moq.ResultsByParams_Size[insertAt+1:], r.Moq.ResultsByParams_Size[insertAt:0])
			r.Moq.ResultsByParams_Size[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Size(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqPublicKey_starGenType_Size_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPublicKey_starGenType_Size_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPublicKey_starGenType_Size_fnRecorder {
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
					Result1 int
				}
				Sequence   uint32
				DoFn       MoqPublicKey_starGenType_Size_doFn
				DoReturnFn MoqPublicKey_starGenType_Size_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPublicKey_starGenType) PrettyParams_Size(params MoqPublicKey_starGenType_Size_params) string {
	return fmt.Sprintf("Size()")
}

func (m *MoqPublicKey_starGenType) ParamsKey_Size(params MoqPublicKey_starGenType_Size_params, anyParams uint64) MoqPublicKey_starGenType_Size_paramsKey {
	m.Scene.T.Helper()
	return MoqPublicKey_starGenType_Size_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqPublicKey_starGenType) Reset() { m.ResultsByParams_Size = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPublicKey_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Size {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Size(results.Params))
			}
		}
	}
}
