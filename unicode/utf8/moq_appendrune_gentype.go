// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package utf8

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// AppendRune_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type AppendRune_genType func(p []byte, r rune) []byte

// MoqAppendRune_genType holds the state of a moq of the AppendRune_genType
// type
type MoqAppendRune_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAppendRune_genType_mock

	ResultsByParams []MoqAppendRune_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			P      moq.ParamIndexing
			Param2 moq.ParamIndexing
		}
	}
}

// MoqAppendRune_genType_mock isolates the mock interface of the
// AppendRune_genType type
type MoqAppendRune_genType_mock struct {
	Moq *MoqAppendRune_genType
}

// MoqAppendRune_genType_params holds the params of the AppendRune_genType type
type MoqAppendRune_genType_params struct {
	P      []byte
	Param2 rune
}

// MoqAppendRune_genType_paramsKey holds the map key params of the
// AppendRune_genType type
type MoqAppendRune_genType_paramsKey struct {
	Params struct{ Param2 rune }
	Hashes struct {
		P      hash.Hash
		Param2 hash.Hash
	}
}

// MoqAppendRune_genType_resultsByParams contains the results for a given set
// of parameters for the AppendRune_genType type
type MoqAppendRune_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAppendRune_genType_paramsKey]*MoqAppendRune_genType_results
}

// MoqAppendRune_genType_doFn defines the type of function needed when calling
// AndDo for the AppendRune_genType type
type MoqAppendRune_genType_doFn func(p []byte, r rune)

// MoqAppendRune_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the AppendRune_genType type
type MoqAppendRune_genType_doReturnFn func(p []byte, r rune) []byte

// MoqAppendRune_genType_results holds the results of the AppendRune_genType
// type
type MoqAppendRune_genType_results struct {
	Params  MoqAppendRune_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
		}
		Sequence   uint32
		DoFn       MoqAppendRune_genType_doFn
		DoReturnFn MoqAppendRune_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAppendRune_genType_fnRecorder routes recorded function calls to the
// MoqAppendRune_genType moq
type MoqAppendRune_genType_fnRecorder struct {
	Params    MoqAppendRune_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAppendRune_genType_results
	Moq       *MoqAppendRune_genType
}

// MoqAppendRune_genType_anyParams isolates the any params functions of the
// AppendRune_genType type
type MoqAppendRune_genType_anyParams struct {
	Recorder *MoqAppendRune_genType_fnRecorder
}

// NewMoqAppendRune_genType creates a new moq of the AppendRune_genType type
func NewMoqAppendRune_genType(scene *moq.Scene, config *moq.Config) *MoqAppendRune_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAppendRune_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAppendRune_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				P      moq.ParamIndexing
				Param2 moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			P      moq.ParamIndexing
			Param2 moq.ParamIndexing
		}{
			P:      moq.ParamIndexByHash,
			Param2: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the AppendRune_genType type
func (m *MoqAppendRune_genType) Mock() AppendRune_genType {
	return func(p []byte, param2 rune) []byte {
		m.Scene.T.Helper()
		moq := &MoqAppendRune_genType_mock{Moq: m}
		return moq.Fn(p, param2)
	}
}

func (m *MoqAppendRune_genType_mock) Fn(p []byte, param2 rune) (result1 []byte) {
	m.Moq.Scene.T.Helper()
	params := MoqAppendRune_genType_params{
		P:      p,
		Param2: param2,
	}
	var results *MoqAppendRune_genType_results
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
		result.DoFn(p, param2)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(p, param2)
	}
	return
}

func (m *MoqAppendRune_genType) OnCall(p []byte, param2 rune) *MoqAppendRune_genType_fnRecorder {
	return &MoqAppendRune_genType_fnRecorder{
		Params: MoqAppendRune_genType_params{
			P:      p,
			Param2: param2,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqAppendRune_genType_fnRecorder) Any() *MoqAppendRune_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqAppendRune_genType_anyParams{Recorder: r}
}

func (a *MoqAppendRune_genType_anyParams) P() *MoqAppendRune_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqAppendRune_genType_anyParams) Param2() *MoqAppendRune_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqAppendRune_genType_fnRecorder) Seq() *MoqAppendRune_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAppendRune_genType_fnRecorder) NoSeq() *MoqAppendRune_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAppendRune_genType_fnRecorder) ReturnResults(result1 []byte) *MoqAppendRune_genType_fnRecorder {
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
		DoFn       MoqAppendRune_genType_doFn
		DoReturnFn MoqAppendRune_genType_doReturnFn
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

func (r *MoqAppendRune_genType_fnRecorder) AndDo(fn MoqAppendRune_genType_doFn) *MoqAppendRune_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAppendRune_genType_fnRecorder) DoReturnResults(fn MoqAppendRune_genType_doReturnFn) *MoqAppendRune_genType_fnRecorder {
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
		DoFn       MoqAppendRune_genType_doFn
		DoReturnFn MoqAppendRune_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAppendRune_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAppendRune_genType_resultsByParams
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
		results = &MoqAppendRune_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAppendRune_genType_paramsKey]*MoqAppendRune_genType_results{},
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
		r.Results = &MoqAppendRune_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAppendRune_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAppendRune_genType_fnRecorder {
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
				DoFn       MoqAppendRune_genType_doFn
				DoReturnFn MoqAppendRune_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAppendRune_genType) PrettyParams(params MoqAppendRune_genType_params) string {
	return fmt.Sprintf("AppendRune_genType(%#v, %#v)", params.P, params.Param2)
}

func (m *MoqAppendRune_genType) ParamsKey(params MoqAppendRune_genType_params, anyParams uint64) MoqAppendRune_genType_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	var param2Used rune
	var param2UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Param2 == moq.ParamIndexByValue {
			param2Used = params.Param2
		} else {
			param2UsedHash = hash.DeepHash(params.Param2)
		}
	}
	return MoqAppendRune_genType_paramsKey{
		Params: struct{ Param2 rune }{
			Param2: param2Used,
		},
		Hashes: struct {
			P      hash.Hash
			Param2 hash.Hash
		}{
			P:      pUsedHash,
			Param2: param2UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqAppendRune_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAppendRune_genType) AssertExpectationsMet() {
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
