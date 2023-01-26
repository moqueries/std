// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package hex

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// DecodeString_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type DecodeString_genType func(s string) ([]byte, error)

// MoqDecodeString_genType holds the state of a moq of the DecodeString_genType
// type
type MoqDecodeString_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDecodeString_genType_mock

	ResultsByParams []MoqDecodeString_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			S moq.ParamIndexing
		}
	}
}

// MoqDecodeString_genType_mock isolates the mock interface of the
// DecodeString_genType type
type MoqDecodeString_genType_mock struct {
	Moq *MoqDecodeString_genType
}

// MoqDecodeString_genType_params holds the params of the DecodeString_genType
// type
type MoqDecodeString_genType_params struct{ S string }

// MoqDecodeString_genType_paramsKey holds the map key params of the
// DecodeString_genType type
type MoqDecodeString_genType_paramsKey struct {
	Params struct{ S string }
	Hashes struct{ S hash.Hash }
}

// MoqDecodeString_genType_resultsByParams contains the results for a given set
// of parameters for the DecodeString_genType type
type MoqDecodeString_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDecodeString_genType_paramsKey]*MoqDecodeString_genType_results
}

// MoqDecodeString_genType_doFn defines the type of function needed when
// calling AndDo for the DecodeString_genType type
type MoqDecodeString_genType_doFn func(s string)

// MoqDecodeString_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the DecodeString_genType type
type MoqDecodeString_genType_doReturnFn func(s string) ([]byte, error)

// MoqDecodeString_genType_results holds the results of the
// DecodeString_genType type
type MoqDecodeString_genType_results struct {
	Params  MoqDecodeString_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDecodeString_genType_doFn
		DoReturnFn MoqDecodeString_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDecodeString_genType_fnRecorder routes recorded function calls to the
// MoqDecodeString_genType moq
type MoqDecodeString_genType_fnRecorder struct {
	Params    MoqDecodeString_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDecodeString_genType_results
	Moq       *MoqDecodeString_genType
}

// MoqDecodeString_genType_anyParams isolates the any params functions of the
// DecodeString_genType type
type MoqDecodeString_genType_anyParams struct {
	Recorder *MoqDecodeString_genType_fnRecorder
}

// NewMoqDecodeString_genType creates a new moq of the DecodeString_genType
// type
func NewMoqDecodeString_genType(scene *moq.Scene, config *moq.Config) *MoqDecodeString_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDecodeString_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDecodeString_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				S moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			S moq.ParamIndexing
		}{
			S: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the DecodeString_genType type
func (m *MoqDecodeString_genType) Mock() DecodeString_genType {
	return func(s string) ([]byte, error) {
		m.Scene.T.Helper()
		moq := &MoqDecodeString_genType_mock{Moq: m}
		return moq.Fn(s)
	}
}

func (m *MoqDecodeString_genType_mock) Fn(s string) (result1 []byte, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqDecodeString_genType_params{
		S: s,
	}
	var results *MoqDecodeString_genType_results
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
		result.DoFn(s)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(s)
	}
	return
}

func (m *MoqDecodeString_genType) OnCall(s string) *MoqDecodeString_genType_fnRecorder {
	return &MoqDecodeString_genType_fnRecorder{
		Params: MoqDecodeString_genType_params{
			S: s,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqDecodeString_genType_fnRecorder) Any() *MoqDecodeString_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqDecodeString_genType_anyParams{Recorder: r}
}

func (a *MoqDecodeString_genType_anyParams) S() *MoqDecodeString_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqDecodeString_genType_fnRecorder) Seq() *MoqDecodeString_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDecodeString_genType_fnRecorder) NoSeq() *MoqDecodeString_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDecodeString_genType_fnRecorder) ReturnResults(result1 []byte, result2 error) *MoqDecodeString_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDecodeString_genType_doFn
		DoReturnFn MoqDecodeString_genType_doReturnFn
	}{
		Values: &struct {
			Result1 []byte
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDecodeString_genType_fnRecorder) AndDo(fn MoqDecodeString_genType_doFn) *MoqDecodeString_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDecodeString_genType_fnRecorder) DoReturnResults(fn MoqDecodeString_genType_doReturnFn) *MoqDecodeString_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqDecodeString_genType_doFn
		DoReturnFn MoqDecodeString_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDecodeString_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDecodeString_genType_resultsByParams
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
		results = &MoqDecodeString_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDecodeString_genType_paramsKey]*MoqDecodeString_genType_results{},
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
		r.Results = &MoqDecodeString_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDecodeString_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDecodeString_genType_fnRecorder {
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
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqDecodeString_genType_doFn
				DoReturnFn MoqDecodeString_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDecodeString_genType) PrettyParams(params MoqDecodeString_genType_params) string {
	return fmt.Sprintf("DecodeString_genType(%#v)", params.S)
}

func (m *MoqDecodeString_genType) ParamsKey(params MoqDecodeString_genType_params, anyParams uint64) MoqDecodeString_genType_paramsKey {
	m.Scene.T.Helper()
	var sUsed string
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			sUsed = params.S
		} else {
			sUsedHash = hash.DeepHash(params.S)
		}
	}
	return MoqDecodeString_genType_paramsKey{
		Params: struct{ S string }{
			S: sUsed,
		},
		Hashes: struct{ S hash.Hash }{
			S: sUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqDecodeString_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDecodeString_genType) AssertExpectationsMet() {
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