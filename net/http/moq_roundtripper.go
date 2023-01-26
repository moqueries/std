// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package http

import (
	"fmt"
	"math/bits"
	"net/http"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that http.RoundTripper is mocked
// completely
var _ http.RoundTripper = (*MoqRoundTripper_mock)(nil)

// MoqRoundTripper holds the state of a moq of the RoundTripper type
type MoqRoundTripper struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqRoundTripper_mock

	ResultsByParams_RoundTrip []MoqRoundTripper_RoundTrip_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			RoundTrip struct {
				Param1 moq.ParamIndexing
			}
		}
	}
	// MoqRoundTripper_mock isolates the mock interface of the RoundTripper type
}

type MoqRoundTripper_mock struct {
	Moq *MoqRoundTripper
}

// MoqRoundTripper_recorder isolates the recorder interface of the RoundTripper
// type
type MoqRoundTripper_recorder struct {
	Moq *MoqRoundTripper
}

// MoqRoundTripper_RoundTrip_params holds the params of the RoundTripper type
type MoqRoundTripper_RoundTrip_params struct{ Param1 *http.Request }

// MoqRoundTripper_RoundTrip_paramsKey holds the map key params of the
// RoundTripper type
type MoqRoundTripper_RoundTrip_paramsKey struct {
	Params struct{ Param1 *http.Request }
	Hashes struct{ Param1 hash.Hash }
}

// MoqRoundTripper_RoundTrip_resultsByParams contains the results for a given
// set of parameters for the RoundTripper type
type MoqRoundTripper_RoundTrip_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRoundTripper_RoundTrip_paramsKey]*MoqRoundTripper_RoundTrip_results
}

// MoqRoundTripper_RoundTrip_doFn defines the type of function needed when
// calling AndDo for the RoundTripper type
type MoqRoundTripper_RoundTrip_doFn func(*http.Request)

// MoqRoundTripper_RoundTrip_doReturnFn defines the type of function needed
// when calling DoReturnResults for the RoundTripper type
type MoqRoundTripper_RoundTrip_doReturnFn func(*http.Request) (*http.Response, error)

// MoqRoundTripper_RoundTrip_results holds the results of the RoundTripper type
type MoqRoundTripper_RoundTrip_results struct {
	Params  MoqRoundTripper_RoundTrip_params
	Results []struct {
		Values *struct {
			Result1 *http.Response
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRoundTripper_RoundTrip_doFn
		DoReturnFn MoqRoundTripper_RoundTrip_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRoundTripper_RoundTrip_fnRecorder routes recorded function calls to the
// MoqRoundTripper moq
type MoqRoundTripper_RoundTrip_fnRecorder struct {
	Params    MoqRoundTripper_RoundTrip_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRoundTripper_RoundTrip_results
	Moq       *MoqRoundTripper
}

// MoqRoundTripper_RoundTrip_anyParams isolates the any params functions of the
// RoundTripper type
type MoqRoundTripper_RoundTrip_anyParams struct {
	Recorder *MoqRoundTripper_RoundTrip_fnRecorder
}

// NewMoqRoundTripper creates a new moq of the RoundTripper type
func NewMoqRoundTripper(scene *moq.Scene, config *moq.Config) *MoqRoundTripper {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqRoundTripper{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqRoundTripper_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				RoundTrip struct {
					Param1 moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			RoundTrip struct {
				Param1 moq.ParamIndexing
			}
		}{
			RoundTrip: struct {
				Param1 moq.ParamIndexing
			}{
				Param1: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the RoundTripper type
func (m *MoqRoundTripper) Mock() *MoqRoundTripper_mock { return m.Moq }

func (m *MoqRoundTripper_mock) RoundTrip(param1 *http.Request) (result1 *http.Response, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqRoundTripper_RoundTrip_params{
		Param1: param1,
	}
	var results *MoqRoundTripper_RoundTrip_results
	for _, resultsByParams := range m.Moq.ResultsByParams_RoundTrip {
		paramsKey := m.Moq.ParamsKey_RoundTrip(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_RoundTrip(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_RoundTrip(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_RoundTrip(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(param1)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(param1)
	}
	return
}

// OnCall returns the recorder implementation of the RoundTripper type
func (m *MoqRoundTripper) OnCall() *MoqRoundTripper_recorder {
	return &MoqRoundTripper_recorder{
		Moq: m,
	}
}

func (m *MoqRoundTripper_recorder) RoundTrip(param1 *http.Request) *MoqRoundTripper_RoundTrip_fnRecorder {
	return &MoqRoundTripper_RoundTrip_fnRecorder{
		Params: MoqRoundTripper_RoundTrip_params{
			Param1: param1,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqRoundTripper_RoundTrip_fnRecorder) Any() *MoqRoundTripper_RoundTrip_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RoundTrip(r.Params))
		return nil
	}
	return &MoqRoundTripper_RoundTrip_anyParams{Recorder: r}
}

func (a *MoqRoundTripper_RoundTrip_anyParams) Param1() *MoqRoundTripper_RoundTrip_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqRoundTripper_RoundTrip_fnRecorder) Seq() *MoqRoundTripper_RoundTrip_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RoundTrip(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRoundTripper_RoundTrip_fnRecorder) NoSeq() *MoqRoundTripper_RoundTrip_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RoundTrip(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRoundTripper_RoundTrip_fnRecorder) ReturnResults(result1 *http.Response, result2 error) *MoqRoundTripper_RoundTrip_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *http.Response
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRoundTripper_RoundTrip_doFn
		DoReturnFn MoqRoundTripper_RoundTrip_doReturnFn
	}{
		Values: &struct {
			Result1 *http.Response
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqRoundTripper_RoundTrip_fnRecorder) AndDo(fn MoqRoundTripper_RoundTrip_doFn) *MoqRoundTripper_RoundTrip_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRoundTripper_RoundTrip_fnRecorder) DoReturnResults(fn MoqRoundTripper_RoundTrip_doReturnFn) *MoqRoundTripper_RoundTrip_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *http.Response
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRoundTripper_RoundTrip_doFn
		DoReturnFn MoqRoundTripper_RoundTrip_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRoundTripper_RoundTrip_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRoundTripper_RoundTrip_resultsByParams
	for n, res := range r.Moq.ResultsByParams_RoundTrip {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqRoundTripper_RoundTrip_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRoundTripper_RoundTrip_paramsKey]*MoqRoundTripper_RoundTrip_results{},
		}
		r.Moq.ResultsByParams_RoundTrip = append(r.Moq.ResultsByParams_RoundTrip, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_RoundTrip) {
			copy(r.Moq.ResultsByParams_RoundTrip[insertAt+1:], r.Moq.ResultsByParams_RoundTrip[insertAt:0])
			r.Moq.ResultsByParams_RoundTrip[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_RoundTrip(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqRoundTripper_RoundTrip_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRoundTripper_RoundTrip_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRoundTripper_RoundTrip_fnRecorder {
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
					Result1 *http.Response
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqRoundTripper_RoundTrip_doFn
				DoReturnFn MoqRoundTripper_RoundTrip_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRoundTripper) PrettyParams_RoundTrip(params MoqRoundTripper_RoundTrip_params) string {
	return fmt.Sprintf("RoundTrip(%#v)", params.Param1)
}

func (m *MoqRoundTripper) ParamsKey_RoundTrip(params MoqRoundTripper_RoundTrip_params, anyParams uint64) MoqRoundTripper_RoundTrip_paramsKey {
	m.Scene.T.Helper()
	var param1Used *http.Request
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.RoundTrip.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqRoundTripper_RoundTrip_paramsKey{
		Params: struct{ Param1 *http.Request }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqRoundTripper) Reset() { m.ResultsByParams_RoundTrip = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqRoundTripper) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_RoundTrip {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_RoundTrip(results.Params))
			}
		}
	}
}
