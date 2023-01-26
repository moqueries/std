// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package strconv

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// AppendQuote_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type AppendQuote_genType func(dst []byte, s string) []byte

// MoqAppendQuote_genType holds the state of a moq of the AppendQuote_genType
// type
type MoqAppendQuote_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAppendQuote_genType_mock

	ResultsByParams []MoqAppendQuote_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Dst moq.ParamIndexing
			S   moq.ParamIndexing
		}
	}
}

// MoqAppendQuote_genType_mock isolates the mock interface of the
// AppendQuote_genType type
type MoqAppendQuote_genType_mock struct {
	Moq *MoqAppendQuote_genType
}

// MoqAppendQuote_genType_params holds the params of the AppendQuote_genType
// type
type MoqAppendQuote_genType_params struct {
	Dst []byte
	S   string
}

// MoqAppendQuote_genType_paramsKey holds the map key params of the
// AppendQuote_genType type
type MoqAppendQuote_genType_paramsKey struct {
	Params struct{ S string }
	Hashes struct {
		Dst hash.Hash
		S   hash.Hash
	}
}

// MoqAppendQuote_genType_resultsByParams contains the results for a given set
// of parameters for the AppendQuote_genType type
type MoqAppendQuote_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAppendQuote_genType_paramsKey]*MoqAppendQuote_genType_results
}

// MoqAppendQuote_genType_doFn defines the type of function needed when calling
// AndDo for the AppendQuote_genType type
type MoqAppendQuote_genType_doFn func(dst []byte, s string)

// MoqAppendQuote_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the AppendQuote_genType type
type MoqAppendQuote_genType_doReturnFn func(dst []byte, s string) []byte

// MoqAppendQuote_genType_results holds the results of the AppendQuote_genType
// type
type MoqAppendQuote_genType_results struct {
	Params  MoqAppendQuote_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqAppendQuote_genType_doFn
		DoReturnFn MoqAppendQuote_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAppendQuote_genType_fnRecorder routes recorded function calls to the
// MoqAppendQuote_genType moq
type MoqAppendQuote_genType_fnRecorder struct {
	Params    MoqAppendQuote_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAppendQuote_genType_results
	Moq       *MoqAppendQuote_genType
}

// MoqAppendQuote_genType_anyParams isolates the any params functions of the
// AppendQuote_genType type
type MoqAppendQuote_genType_anyParams struct {
	Recorder *MoqAppendQuote_genType_fnRecorder
}

// NewMoqAppendQuote_genType creates a new moq of the AppendQuote_genType type
func NewMoqAppendQuote_genType(scene *moq.Scene, config *moq.Config) *MoqAppendQuote_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAppendQuote_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAppendQuote_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Dst moq.ParamIndexing
				S   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Dst moq.ParamIndexing
			S   moq.ParamIndexing
		}{
			Dst: moq.ParamIndexByHash,
			S:   moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the AppendQuote_genType type
func (m *MoqAppendQuote_genType) Mock() AppendQuote_genType {
	return func(dst []byte, s string) []byte {
		m.Scene.T.Helper()
		moq := &MoqAppendQuote_genType_mock{Moq: m}
		return moq.Fn(dst, s)
	}
}

func (m *MoqAppendQuote_genType_mock) Fn(dst []byte, s string) (result1 []byte) {
	m.Moq.Scene.T.Helper()
	params := MoqAppendQuote_genType_params{
		Dst: dst,
		S:   s,
	}
	var results *MoqAppendQuote_genType_results
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
		result.DoFn(dst, s)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(dst, s)
	}
	return
}

func (m *MoqAppendQuote_genType) OnCall(dst []byte, s string) *MoqAppendQuote_genType_fnRecorder {
	return &MoqAppendQuote_genType_fnRecorder{
		Params: MoqAppendQuote_genType_params{
			Dst: dst,
			S:   s,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqAppendQuote_genType_fnRecorder) Any() *MoqAppendQuote_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqAppendQuote_genType_anyParams{Recorder: r}
}

func (a *MoqAppendQuote_genType_anyParams) Dst() *MoqAppendQuote_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqAppendQuote_genType_anyParams) S() *MoqAppendQuote_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqAppendQuote_genType_fnRecorder) Seq() *MoqAppendQuote_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAppendQuote_genType_fnRecorder) NoSeq() *MoqAppendQuote_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAppendQuote_genType_fnRecorder) ReturnResults(result1 []byte) *MoqAppendQuote_genType_fnRecorder {
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
		DoFn       MoqAppendQuote_genType_doFn
		DoReturnFn MoqAppendQuote_genType_doReturnFn
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

func (r *MoqAppendQuote_genType_fnRecorder) AndDo(fn MoqAppendQuote_genType_doFn) *MoqAppendQuote_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAppendQuote_genType_fnRecorder) DoReturnResults(fn MoqAppendQuote_genType_doReturnFn) *MoqAppendQuote_genType_fnRecorder {
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
		DoFn       MoqAppendQuote_genType_doFn
		DoReturnFn MoqAppendQuote_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAppendQuote_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAppendQuote_genType_resultsByParams
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
		results = &MoqAppendQuote_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAppendQuote_genType_paramsKey]*MoqAppendQuote_genType_results{},
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
		r.Results = &MoqAppendQuote_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAppendQuote_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAppendQuote_genType_fnRecorder {
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
				DoFn       MoqAppendQuote_genType_doFn
				DoReturnFn MoqAppendQuote_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAppendQuote_genType) PrettyParams(params MoqAppendQuote_genType_params) string {
	return fmt.Sprintf("AppendQuote_genType(%#v, %#v)", params.Dst, params.S)
}

func (m *MoqAppendQuote_genType) ParamsKey(params MoqAppendQuote_genType_params, anyParams uint64) MoqAppendQuote_genType_paramsKey {
	m.Scene.T.Helper()
	var dstUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Dst == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The dst parameter can't be indexed by value")
		}
		dstUsedHash = hash.DeepHash(params.Dst)
	}
	var sUsed string
	var sUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			sUsed = params.S
		} else {
			sUsedHash = hash.DeepHash(params.S)
		}
	}
	return MoqAppendQuote_genType_paramsKey{
		Params: struct{ S string }{
			S: sUsed,
		},
		Hashes: struct {
			Dst hash.Hash
			S   hash.Hash
		}{
			Dst: dstUsedHash,
			S:   sUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqAppendQuote_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAppendQuote_genType) AssertExpectationsMet() {
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