// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package http

import (
	"bufio"
	"fmt"
	"math/bits"
	"net/http"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ReadRequest_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type ReadRequest_genType func(b *bufio.Reader) (*http.Request, error)

// MoqReadRequest_genType holds the state of a moq of the ReadRequest_genType
// type
type MoqReadRequest_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqReadRequest_genType_mock

	ResultsByParams []MoqReadRequest_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			B moq.ParamIndexing
		}
	}
}

// MoqReadRequest_genType_mock isolates the mock interface of the
// ReadRequest_genType type
type MoqReadRequest_genType_mock struct {
	Moq *MoqReadRequest_genType
}

// MoqReadRequest_genType_params holds the params of the ReadRequest_genType
// type
type MoqReadRequest_genType_params struct{ B *bufio.Reader }

// MoqReadRequest_genType_paramsKey holds the map key params of the
// ReadRequest_genType type
type MoqReadRequest_genType_paramsKey struct {
	Params struct{ B *bufio.Reader }
	Hashes struct{ B hash.Hash }
}

// MoqReadRequest_genType_resultsByParams contains the results for a given set
// of parameters for the ReadRequest_genType type
type MoqReadRequest_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqReadRequest_genType_paramsKey]*MoqReadRequest_genType_results
}

// MoqReadRequest_genType_doFn defines the type of function needed when calling
// AndDo for the ReadRequest_genType type
type MoqReadRequest_genType_doFn func(b *bufio.Reader)

// MoqReadRequest_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ReadRequest_genType type
type MoqReadRequest_genType_doReturnFn func(b *bufio.Reader) (*http.Request, error)

// MoqReadRequest_genType_results holds the results of the ReadRequest_genType
// type
type MoqReadRequest_genType_results struct {
	Params  MoqReadRequest_genType_params
	Results []struct {
		Values *struct {
			Result1 *http.Request
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReadRequest_genType_doFn
		DoReturnFn MoqReadRequest_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqReadRequest_genType_fnRecorder routes recorded function calls to the
// MoqReadRequest_genType moq
type MoqReadRequest_genType_fnRecorder struct {
	Params    MoqReadRequest_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqReadRequest_genType_results
	Moq       *MoqReadRequest_genType
}

// MoqReadRequest_genType_anyParams isolates the any params functions of the
// ReadRequest_genType type
type MoqReadRequest_genType_anyParams struct {
	Recorder *MoqReadRequest_genType_fnRecorder
}

// NewMoqReadRequest_genType creates a new moq of the ReadRequest_genType type
func NewMoqReadRequest_genType(scene *moq.Scene, config *moq.Config) *MoqReadRequest_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqReadRequest_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqReadRequest_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				B moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			B moq.ParamIndexing
		}{
			B: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the ReadRequest_genType type
func (m *MoqReadRequest_genType) Mock() ReadRequest_genType {
	return func(b *bufio.Reader) (*http.Request, error) {
		m.Scene.T.Helper()
		moq := &MoqReadRequest_genType_mock{Moq: m}
		return moq.Fn(b)
	}
}

func (m *MoqReadRequest_genType_mock) Fn(b *bufio.Reader) (result1 *http.Request, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqReadRequest_genType_params{
		B: b,
	}
	var results *MoqReadRequest_genType_results
	for _, resultsByParams := range m.Moq.ResultsByParams {
		paramsKey := m.Moq.ParamsKey(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(b)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(b)
	}
	return
}

func (m *MoqReadRequest_genType) OnCall(b *bufio.Reader) *MoqReadRequest_genType_fnRecorder {
	return &MoqReadRequest_genType_fnRecorder{
		Params: MoqReadRequest_genType_params{
			B: b,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqReadRequest_genType_fnRecorder) Any() *MoqReadRequest_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqReadRequest_genType_anyParams{Recorder: r}
}

func (a *MoqReadRequest_genType_anyParams) B() *MoqReadRequest_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqReadRequest_genType_fnRecorder) Seq() *MoqReadRequest_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqReadRequest_genType_fnRecorder) NoSeq() *MoqReadRequest_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqReadRequest_genType_fnRecorder) ReturnResults(result1 *http.Request, result2 error) *MoqReadRequest_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *http.Request
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReadRequest_genType_doFn
		DoReturnFn MoqReadRequest_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *http.Request
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqReadRequest_genType_fnRecorder) AndDo(fn MoqReadRequest_genType_doFn) *MoqReadRequest_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqReadRequest_genType_fnRecorder) DoReturnResults(fn MoqReadRequest_genType_doReturnFn) *MoqReadRequest_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *http.Request
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqReadRequest_genType_doFn
		DoReturnFn MoqReadRequest_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqReadRequest_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqReadRequest_genType_resultsByParams
	for n, res := range r.Moq.ResultsByParams {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqReadRequest_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqReadRequest_genType_paramsKey]*MoqReadRequest_genType_results{},
		}
		r.Moq.ResultsByParams = append(r.Moq.ResultsByParams, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams) {
			copy(r.Moq.ResultsByParams[insertAt+1:], r.Moq.ResultsByParams[insertAt:0])
			r.Moq.ResultsByParams[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqReadRequest_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqReadRequest_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqReadRequest_genType_fnRecorder {
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
					Result1 *http.Request
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqReadRequest_genType_doFn
				DoReturnFn MoqReadRequest_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqReadRequest_genType) PrettyParams(params MoqReadRequest_genType_params) string {
	return fmt.Sprintf("ReadRequest_genType(%#v)", params.B)
}

func (m *MoqReadRequest_genType) ParamsKey(params MoqReadRequest_genType_params, anyParams uint64) MoqReadRequest_genType_paramsKey {
	m.Scene.T.Helper()
	var bUsed *bufio.Reader
	var bUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.B == moq.ParamIndexByValue {
			bUsed = params.B
		} else {
			bUsedHash = hash.DeepHash(params.B)
		}
	}
	return MoqReadRequest_genType_paramsKey{
		Params: struct{ B *bufio.Reader }{
			B: bUsed,
		},
		Hashes: struct{ B hash.Hash }{
			B: bUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqReadRequest_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqReadRequest_genType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams(results.Params))
			}
		}
	}
}
