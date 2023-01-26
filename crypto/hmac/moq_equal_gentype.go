// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package hmac

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Equal_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Equal_genType func(mac1, mac2 []byte) bool

// MoqEqual_genType holds the state of a moq of the Equal_genType type
type MoqEqual_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqEqual_genType_mock

	ResultsByParams []MoqEqual_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Mac1 moq.ParamIndexing
			Mac2 moq.ParamIndexing
		}
	}
}

// MoqEqual_genType_mock isolates the mock interface of the Equal_genType type
type MoqEqual_genType_mock struct {
	Moq *MoqEqual_genType
}

// MoqEqual_genType_params holds the params of the Equal_genType type
type MoqEqual_genType_params struct{ Mac1, Mac2 []byte }

// MoqEqual_genType_paramsKey holds the map key params of the Equal_genType
// type
type MoqEqual_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ Mac1, Mac2 hash.Hash }
}

// MoqEqual_genType_resultsByParams contains the results for a given set of
// parameters for the Equal_genType type
type MoqEqual_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqEqual_genType_paramsKey]*MoqEqual_genType_results
}

// MoqEqual_genType_doFn defines the type of function needed when calling AndDo
// for the Equal_genType type
type MoqEqual_genType_doFn func(mac1, mac2 []byte)

// MoqEqual_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Equal_genType type
type MoqEqual_genType_doReturnFn func(mac1, mac2 []byte) bool

// MoqEqual_genType_results holds the results of the Equal_genType type
type MoqEqual_genType_results struct {
	Params  MoqEqual_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqEqual_genType_doFn
		DoReturnFn MoqEqual_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqEqual_genType_fnRecorder routes recorded function calls to the
// MoqEqual_genType moq
type MoqEqual_genType_fnRecorder struct {
	Params    MoqEqual_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqEqual_genType_results
	Moq       *MoqEqual_genType
}

// MoqEqual_genType_anyParams isolates the any params functions of the
// Equal_genType type
type MoqEqual_genType_anyParams struct {
	Recorder *MoqEqual_genType_fnRecorder
}

// NewMoqEqual_genType creates a new moq of the Equal_genType type
func NewMoqEqual_genType(scene *moq.Scene, config *moq.Config) *MoqEqual_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqEqual_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqEqual_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Mac1 moq.ParamIndexing
				Mac2 moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Mac1 moq.ParamIndexing
			Mac2 moq.ParamIndexing
		}{
			Mac1: moq.ParamIndexByHash,
			Mac2: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Equal_genType type
func (m *MoqEqual_genType) Mock() Equal_genType {
	return func(mac1, mac2 []byte) bool {
		m.Scene.T.Helper()
		moq := &MoqEqual_genType_mock{Moq: m}
		return moq.Fn(mac1, mac2)
	}
}

func (m *MoqEqual_genType_mock) Fn(mac1, mac2 []byte) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqEqual_genType_params{
		Mac1: mac1,
		Mac2: mac2,
	}
	var results *MoqEqual_genType_results
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
		result.DoFn(mac1, mac2)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(mac1, mac2)
	}
	return
}

func (m *MoqEqual_genType) OnCall(mac1, mac2 []byte) *MoqEqual_genType_fnRecorder {
	return &MoqEqual_genType_fnRecorder{
		Params: MoqEqual_genType_params{
			Mac1: mac1,
			Mac2: mac2,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqEqual_genType_fnRecorder) Any() *MoqEqual_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqEqual_genType_anyParams{Recorder: r}
}

func (a *MoqEqual_genType_anyParams) Mac1() *MoqEqual_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqEqual_genType_anyParams) Mac2() *MoqEqual_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqEqual_genType_fnRecorder) Seq() *MoqEqual_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqEqual_genType_fnRecorder) NoSeq() *MoqEqual_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqEqual_genType_fnRecorder) ReturnResults(result1 bool) *MoqEqual_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqEqual_genType_doFn
		DoReturnFn MoqEqual_genType_doReturnFn
	}{
		Values: &struct {
			Result1 bool
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqEqual_genType_fnRecorder) AndDo(fn MoqEqual_genType_doFn) *MoqEqual_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqEqual_genType_fnRecorder) DoReturnResults(fn MoqEqual_genType_doReturnFn) *MoqEqual_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqEqual_genType_doFn
		DoReturnFn MoqEqual_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqEqual_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqEqual_genType_resultsByParams
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
		results = &MoqEqual_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqEqual_genType_paramsKey]*MoqEqual_genType_results{},
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
		r.Results = &MoqEqual_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqEqual_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqEqual_genType_fnRecorder {
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
					Result1 bool
				}
				Sequence   uint32
				DoFn       MoqEqual_genType_doFn
				DoReturnFn MoqEqual_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqEqual_genType) PrettyParams(params MoqEqual_genType_params) string {
	return fmt.Sprintf("Equal_genType(%#v, %#v)", params.Mac1, params.Mac2)
}

func (m *MoqEqual_genType) ParamsKey(params MoqEqual_genType_params, anyParams uint64) MoqEqual_genType_paramsKey {
	m.Scene.T.Helper()
	var mac1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Mac1 == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The mac1 parameter can't be indexed by value")
		}
		mac1UsedHash = hash.DeepHash(params.Mac1)
	}
	var mac2UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Mac2 == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The mac2 parameter can't be indexed by value")
		}
		mac2UsedHash = hash.DeepHash(params.Mac2)
	}
	return MoqEqual_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Mac1, Mac2 hash.Hash }{
			Mac1: mac1UsedHash,
			Mac2: mac2UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqEqual_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqEqual_genType) AssertExpectationsMet() {
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
