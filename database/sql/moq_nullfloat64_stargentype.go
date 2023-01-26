// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sql

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that sql.NullFloat64_starGenType is
// mocked completely
var _ NullFloat64_starGenType = (*MoqNullFloat64_starGenType_mock)(nil)

// NullFloat64_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type NullFloat64_starGenType interface {
	Scan(value interface{}) error
}

// MoqNullFloat64_starGenType holds the state of a moq of the
// NullFloat64_starGenType type
type MoqNullFloat64_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNullFloat64_starGenType_mock

	ResultsByParams_Scan []MoqNullFloat64_starGenType_Scan_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Scan struct {
				Value moq.ParamIndexing
			}
		}
	}
	// MoqNullFloat64_starGenType_mock isolates the mock interface of the
}

// NullFloat64_starGenType type
type MoqNullFloat64_starGenType_mock struct {
	Moq *MoqNullFloat64_starGenType
}

// MoqNullFloat64_starGenType_recorder isolates the recorder interface of the
// NullFloat64_starGenType type
type MoqNullFloat64_starGenType_recorder struct {
	Moq *MoqNullFloat64_starGenType
}

// MoqNullFloat64_starGenType_Scan_params holds the params of the
// NullFloat64_starGenType type
type MoqNullFloat64_starGenType_Scan_params struct{ Value interface{} }

// MoqNullFloat64_starGenType_Scan_paramsKey holds the map key params of the
// NullFloat64_starGenType type
type MoqNullFloat64_starGenType_Scan_paramsKey struct {
	Params struct{ Value interface{} }
	Hashes struct{ Value hash.Hash }
}

// MoqNullFloat64_starGenType_Scan_resultsByParams contains the results for a
// given set of parameters for the NullFloat64_starGenType type
type MoqNullFloat64_starGenType_Scan_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNullFloat64_starGenType_Scan_paramsKey]*MoqNullFloat64_starGenType_Scan_results
}

// MoqNullFloat64_starGenType_Scan_doFn defines the type of function needed
// when calling AndDo for the NullFloat64_starGenType type
type MoqNullFloat64_starGenType_Scan_doFn func(value interface{})

// MoqNullFloat64_starGenType_Scan_doReturnFn defines the type of function
// needed when calling DoReturnResults for the NullFloat64_starGenType type
type MoqNullFloat64_starGenType_Scan_doReturnFn func(value interface{}) error

// MoqNullFloat64_starGenType_Scan_results holds the results of the
// NullFloat64_starGenType type
type MoqNullFloat64_starGenType_Scan_results struct {
	Params  MoqNullFloat64_starGenType_Scan_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqNullFloat64_starGenType_Scan_doFn
		DoReturnFn MoqNullFloat64_starGenType_Scan_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNullFloat64_starGenType_Scan_fnRecorder routes recorded function calls to
// the MoqNullFloat64_starGenType moq
type MoqNullFloat64_starGenType_Scan_fnRecorder struct {
	Params    MoqNullFloat64_starGenType_Scan_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNullFloat64_starGenType_Scan_results
	Moq       *MoqNullFloat64_starGenType
}

// MoqNullFloat64_starGenType_Scan_anyParams isolates the any params functions
// of the NullFloat64_starGenType type
type MoqNullFloat64_starGenType_Scan_anyParams struct {
	Recorder *MoqNullFloat64_starGenType_Scan_fnRecorder
}

// NewMoqNullFloat64_starGenType creates a new moq of the
// NullFloat64_starGenType type
func NewMoqNullFloat64_starGenType(scene *moq.Scene, config *moq.Config) *MoqNullFloat64_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNullFloat64_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNullFloat64_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Scan struct {
					Value moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Scan struct {
				Value moq.ParamIndexing
			}
		}{
			Scan: struct {
				Value moq.ParamIndexing
			}{
				Value: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the NullFloat64_starGenType type
func (m *MoqNullFloat64_starGenType) Mock() *MoqNullFloat64_starGenType_mock { return m.Moq }

func (m *MoqNullFloat64_starGenType_mock) Scan(value interface{}) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqNullFloat64_starGenType_Scan_params{
		Value: value,
	}
	var results *MoqNullFloat64_starGenType_Scan_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Scan {
		paramsKey := m.Moq.ParamsKey_Scan(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Scan(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Scan(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Scan(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(value)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(value)
	}
	return
}

// OnCall returns the recorder implementation of the NullFloat64_starGenType
// type
func (m *MoqNullFloat64_starGenType) OnCall() *MoqNullFloat64_starGenType_recorder {
	return &MoqNullFloat64_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqNullFloat64_starGenType_recorder) Scan(value interface{}) *MoqNullFloat64_starGenType_Scan_fnRecorder {
	return &MoqNullFloat64_starGenType_Scan_fnRecorder{
		Params: MoqNullFloat64_starGenType_Scan_params{
			Value: value,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqNullFloat64_starGenType_Scan_fnRecorder) Any() *MoqNullFloat64_starGenType_Scan_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Scan(r.Params))
		return nil
	}
	return &MoqNullFloat64_starGenType_Scan_anyParams{Recorder: r}
}

func (a *MoqNullFloat64_starGenType_Scan_anyParams) Value() *MoqNullFloat64_starGenType_Scan_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqNullFloat64_starGenType_Scan_fnRecorder) Seq() *MoqNullFloat64_starGenType_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Scan(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNullFloat64_starGenType_Scan_fnRecorder) NoSeq() *MoqNullFloat64_starGenType_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Scan(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNullFloat64_starGenType_Scan_fnRecorder) ReturnResults(result1 error) *MoqNullFloat64_starGenType_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqNullFloat64_starGenType_Scan_doFn
		DoReturnFn MoqNullFloat64_starGenType_Scan_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNullFloat64_starGenType_Scan_fnRecorder) AndDo(fn MoqNullFloat64_starGenType_Scan_doFn) *MoqNullFloat64_starGenType_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNullFloat64_starGenType_Scan_fnRecorder) DoReturnResults(fn MoqNullFloat64_starGenType_Scan_doReturnFn) *MoqNullFloat64_starGenType_Scan_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqNullFloat64_starGenType_Scan_doFn
		DoReturnFn MoqNullFloat64_starGenType_Scan_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNullFloat64_starGenType_Scan_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNullFloat64_starGenType_Scan_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Scan {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqNullFloat64_starGenType_Scan_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNullFloat64_starGenType_Scan_paramsKey]*MoqNullFloat64_starGenType_Scan_results{},
		}
		r.Moq.ResultsByParams_Scan = append(r.Moq.ResultsByParams_Scan, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Scan) {
			copy(r.Moq.ResultsByParams_Scan[insertAt+1:], r.Moq.ResultsByParams_Scan[insertAt:0])
			r.Moq.ResultsByParams_Scan[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Scan(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqNullFloat64_starGenType_Scan_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNullFloat64_starGenType_Scan_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNullFloat64_starGenType_Scan_fnRecorder {
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
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqNullFloat64_starGenType_Scan_doFn
				DoReturnFn MoqNullFloat64_starGenType_Scan_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNullFloat64_starGenType) PrettyParams_Scan(params MoqNullFloat64_starGenType_Scan_params) string {
	return fmt.Sprintf("Scan(%#v)", params.Value)
}

func (m *MoqNullFloat64_starGenType) ParamsKey_Scan(params MoqNullFloat64_starGenType_Scan_params, anyParams uint64) MoqNullFloat64_starGenType_Scan_paramsKey {
	m.Scene.T.Helper()
	var valueUsed interface{}
	var valueUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Scan.Value == moq.ParamIndexByValue {
			valueUsed = params.Value
		} else {
			valueUsedHash = hash.DeepHash(params.Value)
		}
	}
	return MoqNullFloat64_starGenType_Scan_paramsKey{
		Params: struct{ Value interface{} }{
			Value: valueUsed,
		},
		Hashes: struct{ Value hash.Hash }{
			Value: valueUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNullFloat64_starGenType) Reset() { m.ResultsByParams_Scan = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNullFloat64_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Scan {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Scan(results.Params))
			}
		}
	}
}
