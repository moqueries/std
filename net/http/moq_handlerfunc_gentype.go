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

// The following type assertion assures that http.HandlerFunc_genType is mocked
// completely
var _ HandlerFunc_genType = (*MoqHandlerFunc_genType_mock)(nil)

// HandlerFunc_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type HandlerFunc_genType interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// MoqHandlerFunc_genType holds the state of a moq of the HandlerFunc_genType
// type
type MoqHandlerFunc_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqHandlerFunc_genType_mock

	ResultsByParams_ServeHTTP []MoqHandlerFunc_genType_ServeHTTP_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			ServeHTTP struct {
				W      moq.ParamIndexing
				Param2 moq.ParamIndexing
			}
		}
	}
	// MoqHandlerFunc_genType_mock isolates the mock interface of the
}

// HandlerFunc_genType type
type MoqHandlerFunc_genType_mock struct {
	Moq *MoqHandlerFunc_genType
}

// MoqHandlerFunc_genType_recorder isolates the recorder interface of the
// HandlerFunc_genType type
type MoqHandlerFunc_genType_recorder struct {
	Moq *MoqHandlerFunc_genType
}

// MoqHandlerFunc_genType_ServeHTTP_params holds the params of the
// HandlerFunc_genType type
type MoqHandlerFunc_genType_ServeHTTP_params struct {
	W      http.ResponseWriter
	Param2 *http.Request
}

// MoqHandlerFunc_genType_ServeHTTP_paramsKey holds the map key params of the
// HandlerFunc_genType type
type MoqHandlerFunc_genType_ServeHTTP_paramsKey struct {
	Params struct {
		W      http.ResponseWriter
		Param2 *http.Request
	}
	Hashes struct {
		W      hash.Hash
		Param2 hash.Hash
	}
}

// MoqHandlerFunc_genType_ServeHTTP_resultsByParams contains the results for a
// given set of parameters for the HandlerFunc_genType type
type MoqHandlerFunc_genType_ServeHTTP_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqHandlerFunc_genType_ServeHTTP_paramsKey]*MoqHandlerFunc_genType_ServeHTTP_results
}

// MoqHandlerFunc_genType_ServeHTTP_doFn defines the type of function needed
// when calling AndDo for the HandlerFunc_genType type
type MoqHandlerFunc_genType_ServeHTTP_doFn func(w http.ResponseWriter, r *http.Request)

// MoqHandlerFunc_genType_ServeHTTP_doReturnFn defines the type of function
// needed when calling DoReturnResults for the HandlerFunc_genType type
type MoqHandlerFunc_genType_ServeHTTP_doReturnFn func(w http.ResponseWriter, r *http.Request)

// MoqHandlerFunc_genType_ServeHTTP_results holds the results of the
// HandlerFunc_genType type
type MoqHandlerFunc_genType_ServeHTTP_results struct {
	Params  MoqHandlerFunc_genType_ServeHTTP_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqHandlerFunc_genType_ServeHTTP_doFn
		DoReturnFn MoqHandlerFunc_genType_ServeHTTP_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqHandlerFunc_genType_ServeHTTP_fnRecorder routes recorded function calls
// to the MoqHandlerFunc_genType moq
type MoqHandlerFunc_genType_ServeHTTP_fnRecorder struct {
	Params    MoqHandlerFunc_genType_ServeHTTP_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqHandlerFunc_genType_ServeHTTP_results
	Moq       *MoqHandlerFunc_genType
}

// MoqHandlerFunc_genType_ServeHTTP_anyParams isolates the any params functions
// of the HandlerFunc_genType type
type MoqHandlerFunc_genType_ServeHTTP_anyParams struct {
	Recorder *MoqHandlerFunc_genType_ServeHTTP_fnRecorder
}

// NewMoqHandlerFunc_genType creates a new moq of the HandlerFunc_genType type
func NewMoqHandlerFunc_genType(scene *moq.Scene, config *moq.Config) *MoqHandlerFunc_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqHandlerFunc_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqHandlerFunc_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				ServeHTTP struct {
					W      moq.ParamIndexing
					Param2 moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			ServeHTTP struct {
				W      moq.ParamIndexing
				Param2 moq.ParamIndexing
			}
		}{
			ServeHTTP: struct {
				W      moq.ParamIndexing
				Param2 moq.ParamIndexing
			}{
				W:      moq.ParamIndexByHash,
				Param2: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the HandlerFunc_genType type
func (m *MoqHandlerFunc_genType) Mock() *MoqHandlerFunc_genType_mock { return m.Moq }

func (m *MoqHandlerFunc_genType_mock) ServeHTTP(w http.ResponseWriter, param2 *http.Request) {
	m.Moq.Scene.T.Helper()
	params := MoqHandlerFunc_genType_ServeHTTP_params{
		W:      w,
		Param2: param2,
	}
	var results *MoqHandlerFunc_genType_ServeHTTP_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ServeHTTP {
		paramsKey := m.Moq.ParamsKey_ServeHTTP(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ServeHTTP(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ServeHTTP(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ServeHTTP(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(w, param2)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(w, param2)
	}
	return
}

// OnCall returns the recorder implementation of the HandlerFunc_genType type
func (m *MoqHandlerFunc_genType) OnCall() *MoqHandlerFunc_genType_recorder {
	return &MoqHandlerFunc_genType_recorder{
		Moq: m,
	}
}

func (m *MoqHandlerFunc_genType_recorder) ServeHTTP(w http.ResponseWriter, param2 *http.Request) *MoqHandlerFunc_genType_ServeHTTP_fnRecorder {
	return &MoqHandlerFunc_genType_ServeHTTP_fnRecorder{
		Params: MoqHandlerFunc_genType_ServeHTTP_params{
			W:      w,
			Param2: param2,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqHandlerFunc_genType_ServeHTTP_fnRecorder) Any() *MoqHandlerFunc_genType_ServeHTTP_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ServeHTTP(r.Params))
		return nil
	}
	return &MoqHandlerFunc_genType_ServeHTTP_anyParams{Recorder: r}
}

func (a *MoqHandlerFunc_genType_ServeHTTP_anyParams) W() *MoqHandlerFunc_genType_ServeHTTP_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqHandlerFunc_genType_ServeHTTP_anyParams) Param2() *MoqHandlerFunc_genType_ServeHTTP_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqHandlerFunc_genType_ServeHTTP_fnRecorder) Seq() *MoqHandlerFunc_genType_ServeHTTP_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ServeHTTP(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqHandlerFunc_genType_ServeHTTP_fnRecorder) NoSeq() *MoqHandlerFunc_genType_ServeHTTP_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ServeHTTP(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqHandlerFunc_genType_ServeHTTP_fnRecorder) ReturnResults() *MoqHandlerFunc_genType_ServeHTTP_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqHandlerFunc_genType_ServeHTTP_doFn
		DoReturnFn MoqHandlerFunc_genType_ServeHTTP_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqHandlerFunc_genType_ServeHTTP_fnRecorder) AndDo(fn MoqHandlerFunc_genType_ServeHTTP_doFn) *MoqHandlerFunc_genType_ServeHTTP_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqHandlerFunc_genType_ServeHTTP_fnRecorder) DoReturnResults(fn MoqHandlerFunc_genType_ServeHTTP_doReturnFn) *MoqHandlerFunc_genType_ServeHTTP_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqHandlerFunc_genType_ServeHTTP_doFn
		DoReturnFn MoqHandlerFunc_genType_ServeHTTP_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqHandlerFunc_genType_ServeHTTP_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqHandlerFunc_genType_ServeHTTP_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ServeHTTP {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqHandlerFunc_genType_ServeHTTP_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqHandlerFunc_genType_ServeHTTP_paramsKey]*MoqHandlerFunc_genType_ServeHTTP_results{},
		}
		r.Moq.ResultsByParams_ServeHTTP = append(r.Moq.ResultsByParams_ServeHTTP, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ServeHTTP) {
			copy(r.Moq.ResultsByParams_ServeHTTP[insertAt+1:], r.Moq.ResultsByParams_ServeHTTP[insertAt:0])
			r.Moq.ResultsByParams_ServeHTTP[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ServeHTTP(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqHandlerFunc_genType_ServeHTTP_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqHandlerFunc_genType_ServeHTTP_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqHandlerFunc_genType_ServeHTTP_fnRecorder {
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
				Values     *struct{}
				Sequence   uint32
				DoFn       MoqHandlerFunc_genType_ServeHTTP_doFn
				DoReturnFn MoqHandlerFunc_genType_ServeHTTP_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqHandlerFunc_genType) PrettyParams_ServeHTTP(params MoqHandlerFunc_genType_ServeHTTP_params) string {
	return fmt.Sprintf("ServeHTTP(%#v, %#v)", params.W, params.Param2)
}

func (m *MoqHandlerFunc_genType) ParamsKey_ServeHTTP(params MoqHandlerFunc_genType_ServeHTTP_params, anyParams uint64) MoqHandlerFunc_genType_ServeHTTP_paramsKey {
	m.Scene.T.Helper()
	var wUsed http.ResponseWriter
	var wUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.ServeHTTP.W == moq.ParamIndexByValue {
			wUsed = params.W
		} else {
			wUsedHash = hash.DeepHash(params.W)
		}
	}
	var param2Used *http.Request
	var param2UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.ServeHTTP.Param2 == moq.ParamIndexByValue {
			param2Used = params.Param2
		} else {
			param2UsedHash = hash.DeepHash(params.Param2)
		}
	}
	return MoqHandlerFunc_genType_ServeHTTP_paramsKey{
		Params: struct {
			W      http.ResponseWriter
			Param2 *http.Request
		}{
			W:      wUsed,
			Param2: param2Used,
		},
		Hashes: struct {
			W      hash.Hash
			Param2 hash.Hash
		}{
			W:      wUsedHash,
			Param2: param2UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqHandlerFunc_genType) Reset() { m.ResultsByParams_ServeHTTP = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqHandlerFunc_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_ServeHTTP {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ServeHTTP(results.Params))
			}
		}
	}
}
