// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package flate

import (
	"compress/flate"
	"fmt"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that flate.Resetter is mocked
// completely
var _ flate.Resetter = (*MoqResetter_mock)(nil)

// MoqResetter holds the state of a moq of the Resetter type
type MoqResetter struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqResetter_mock

	ResultsByParams_Reset []MoqResetter_Reset_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Reset struct {
				Param1 moq.ParamIndexing
				Dict   moq.ParamIndexing
			}
		}
	}
	// MoqResetter_mock isolates the mock interface of the Resetter type
}

type MoqResetter_mock struct {
	Moq *MoqResetter
}

// MoqResetter_recorder isolates the recorder interface of the Resetter type
type MoqResetter_recorder struct {
	Moq *MoqResetter
}

// MoqResetter_Reset_params holds the params of the Resetter type
type MoqResetter_Reset_params struct {
	Param1 io.Reader
	Dict   []byte
}

// MoqResetter_Reset_paramsKey holds the map key params of the Resetter type
type MoqResetter_Reset_paramsKey struct {
	Params struct{ Param1 io.Reader }
	Hashes struct {
		Param1 hash.Hash
		Dict   hash.Hash
	}
}

// MoqResetter_Reset_resultsByParams contains the results for a given set of
// parameters for the Resetter type
type MoqResetter_Reset_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqResetter_Reset_paramsKey]*MoqResetter_Reset_results
}

// MoqResetter_Reset_doFn defines the type of function needed when calling
// AndDo for the Resetter type
type MoqResetter_Reset_doFn func(r io.Reader, dict []byte)

// MoqResetter_Reset_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Resetter type
type MoqResetter_Reset_doReturnFn func(r io.Reader, dict []byte) error

// MoqResetter_Reset_results holds the results of the Resetter type
type MoqResetter_Reset_results struct {
	Params  MoqResetter_Reset_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqResetter_Reset_doFn
		DoReturnFn MoqResetter_Reset_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqResetter_Reset_fnRecorder routes recorded function calls to the
// MoqResetter moq
type MoqResetter_Reset_fnRecorder struct {
	Params    MoqResetter_Reset_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqResetter_Reset_results
	Moq       *MoqResetter
}

// MoqResetter_Reset_anyParams isolates the any params functions of the
// Resetter type
type MoqResetter_Reset_anyParams struct {
	Recorder *MoqResetter_Reset_fnRecorder
}

// NewMoqResetter creates a new moq of the Resetter type
func NewMoqResetter(scene *moq.Scene, config *moq.Config) *MoqResetter {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqResetter{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqResetter_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Reset struct {
					Param1 moq.ParamIndexing
					Dict   moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Reset struct {
				Param1 moq.ParamIndexing
				Dict   moq.ParamIndexing
			}
		}{
			Reset: struct {
				Param1 moq.ParamIndexing
				Dict   moq.ParamIndexing
			}{
				Param1: moq.ParamIndexByHash,
				Dict:   moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Resetter type
func (m *MoqResetter) Mock() *MoqResetter_mock { return m.Moq }

func (m *MoqResetter_mock) Reset(param1 io.Reader, dict []byte) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqResetter_Reset_params{
		Param1: param1,
		Dict:   dict,
	}
	var results *MoqResetter_Reset_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Reset {
		paramsKey := m.Moq.ParamsKey_Reset(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Reset(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Reset(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Reset(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(param1, dict)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(param1, dict)
	}
	return
}

// OnCall returns the recorder implementation of the Resetter type
func (m *MoqResetter) OnCall() *MoqResetter_recorder {
	return &MoqResetter_recorder{
		Moq: m,
	}
}

func (m *MoqResetter_recorder) Reset(param1 io.Reader, dict []byte) *MoqResetter_Reset_fnRecorder {
	return &MoqResetter_Reset_fnRecorder{
		Params: MoqResetter_Reset_params{
			Param1: param1,
			Dict:   dict,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqResetter_Reset_fnRecorder) Any() *MoqResetter_Reset_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Reset(r.Params))
		return nil
	}
	return &MoqResetter_Reset_anyParams{Recorder: r}
}

func (a *MoqResetter_Reset_anyParams) Param1() *MoqResetter_Reset_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqResetter_Reset_anyParams) Dict() *MoqResetter_Reset_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqResetter_Reset_fnRecorder) Seq() *MoqResetter_Reset_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Reset(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqResetter_Reset_fnRecorder) NoSeq() *MoqResetter_Reset_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Reset(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqResetter_Reset_fnRecorder) ReturnResults(result1 error) *MoqResetter_Reset_fnRecorder {
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
		DoFn       MoqResetter_Reset_doFn
		DoReturnFn MoqResetter_Reset_doReturnFn
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

func (r *MoqResetter_Reset_fnRecorder) AndDo(fn MoqResetter_Reset_doFn) *MoqResetter_Reset_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqResetter_Reset_fnRecorder) DoReturnResults(fn MoqResetter_Reset_doReturnFn) *MoqResetter_Reset_fnRecorder {
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
		DoFn       MoqResetter_Reset_doFn
		DoReturnFn MoqResetter_Reset_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqResetter_Reset_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqResetter_Reset_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Reset {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqResetter_Reset_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqResetter_Reset_paramsKey]*MoqResetter_Reset_results{},
		}
		r.Moq.ResultsByParams_Reset = append(r.Moq.ResultsByParams_Reset, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Reset) {
			copy(r.Moq.ResultsByParams_Reset[insertAt+1:], r.Moq.ResultsByParams_Reset[insertAt:0])
			r.Moq.ResultsByParams_Reset[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Reset(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqResetter_Reset_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqResetter_Reset_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqResetter_Reset_fnRecorder {
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
				DoFn       MoqResetter_Reset_doFn
				DoReturnFn MoqResetter_Reset_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqResetter) PrettyParams_Reset(params MoqResetter_Reset_params) string {
	return fmt.Sprintf("Reset(%#v, %#v)", params.Param1, params.Dict)
}

func (m *MoqResetter) ParamsKey_Reset(params MoqResetter_Reset_params, anyParams uint64) MoqResetter_Reset_paramsKey {
	m.Scene.T.Helper()
	var param1Used io.Reader
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Reset.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	var dictUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Reset.Dict == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The dict parameter of the Reset function can't be indexed by value")
		}
		dictUsedHash = hash.DeepHash(params.Dict)
	}
	return MoqResetter_Reset_paramsKey{
		Params: struct{ Param1 io.Reader }{
			Param1: param1Used,
		},
		Hashes: struct {
			Param1 hash.Hash
			Dict   hash.Hash
		}{
			Param1: param1UsedHash,
			Dict:   dictUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqResetter) Reset() { m.ResultsByParams_Reset = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqResetter) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Reset {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Reset(results.Params))
			}
		}
	}
}
