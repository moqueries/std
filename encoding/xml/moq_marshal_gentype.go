// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package xml

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Marshal_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Marshal_genType func(v any) ([]byte, error)

// MoqMarshal_genType holds the state of a moq of the Marshal_genType type
type MoqMarshal_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMarshal_genType_mock

	ResultsByParams []MoqMarshal_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			V moq.ParamIndexing
		}
	}
}

// MoqMarshal_genType_mock isolates the mock interface of the Marshal_genType
// type
type MoqMarshal_genType_mock struct {
	Moq *MoqMarshal_genType
}

// MoqMarshal_genType_params holds the params of the Marshal_genType type
type MoqMarshal_genType_params struct{ V any }

// MoqMarshal_genType_paramsKey holds the map key params of the Marshal_genType
// type
type MoqMarshal_genType_paramsKey struct {
	Params struct{ V any }
	Hashes struct{ V hash.Hash }
}

// MoqMarshal_genType_resultsByParams contains the results for a given set of
// parameters for the Marshal_genType type
type MoqMarshal_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMarshal_genType_paramsKey]*MoqMarshal_genType_results
}

// MoqMarshal_genType_doFn defines the type of function needed when calling
// AndDo for the Marshal_genType type
type MoqMarshal_genType_doFn func(v any)

// MoqMarshal_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Marshal_genType type
type MoqMarshal_genType_doReturnFn func(v any) ([]byte, error)

// MoqMarshal_genType_results holds the results of the Marshal_genType type
type MoqMarshal_genType_results struct {
	Params  MoqMarshal_genType_params
	Results []struct {
		Values *struct {
			Result1 []byte
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqMarshal_genType_doFn
		DoReturnFn MoqMarshal_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMarshal_genType_fnRecorder routes recorded function calls to the
// MoqMarshal_genType moq
type MoqMarshal_genType_fnRecorder struct {
	Params    MoqMarshal_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMarshal_genType_results
	Moq       *MoqMarshal_genType
}

// MoqMarshal_genType_anyParams isolates the any params functions of the
// Marshal_genType type
type MoqMarshal_genType_anyParams struct {
	Recorder *MoqMarshal_genType_fnRecorder
}

// NewMoqMarshal_genType creates a new moq of the Marshal_genType type
func NewMoqMarshal_genType(scene *moq.Scene, config *moq.Config) *MoqMarshal_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMarshal_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMarshal_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				V moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			V moq.ParamIndexing
		}{
			V: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Marshal_genType type
func (m *MoqMarshal_genType) Mock() Marshal_genType {
	return func(v any) ([]byte, error) {
		m.Scene.T.Helper()
		moq := &MoqMarshal_genType_mock{Moq: m}
		return moq.Fn(v)
	}
}

func (m *MoqMarshal_genType_mock) Fn(v any) (result1 []byte, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqMarshal_genType_params{
		V: v,
	}
	var results *MoqMarshal_genType_results
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
		result.DoFn(v)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(v)
	}
	return
}

func (m *MoqMarshal_genType) OnCall(v any) *MoqMarshal_genType_fnRecorder {
	return &MoqMarshal_genType_fnRecorder{
		Params: MoqMarshal_genType_params{
			V: v,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqMarshal_genType_fnRecorder) Any() *MoqMarshal_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqMarshal_genType_anyParams{Recorder: r}
}

func (a *MoqMarshal_genType_anyParams) V() *MoqMarshal_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqMarshal_genType_fnRecorder) Seq() *MoqMarshal_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMarshal_genType_fnRecorder) NoSeq() *MoqMarshal_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMarshal_genType_fnRecorder) ReturnResults(result1 []byte, result2 error) *MoqMarshal_genType_fnRecorder {
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
		DoFn       MoqMarshal_genType_doFn
		DoReturnFn MoqMarshal_genType_doReturnFn
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

func (r *MoqMarshal_genType_fnRecorder) AndDo(fn MoqMarshal_genType_doFn) *MoqMarshal_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMarshal_genType_fnRecorder) DoReturnResults(fn MoqMarshal_genType_doReturnFn) *MoqMarshal_genType_fnRecorder {
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
		DoFn       MoqMarshal_genType_doFn
		DoReturnFn MoqMarshal_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMarshal_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMarshal_genType_resultsByParams
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
		results = &MoqMarshal_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMarshal_genType_paramsKey]*MoqMarshal_genType_results{},
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
		r.Results = &MoqMarshal_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMarshal_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMarshal_genType_fnRecorder {
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
				DoFn       MoqMarshal_genType_doFn
				DoReturnFn MoqMarshal_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMarshal_genType) PrettyParams(params MoqMarshal_genType_params) string {
	return fmt.Sprintf("Marshal_genType(%#v)", params.V)
}

func (m *MoqMarshal_genType) ParamsKey(params MoqMarshal_genType_params, anyParams uint64) MoqMarshal_genType_paramsKey {
	m.Scene.T.Helper()
	var vUsed any
	var vUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.V == moq.ParamIndexByValue {
			vUsed = params.V
		} else {
			vUsedHash = hash.DeepHash(params.V)
		}
	}
	return MoqMarshal_genType_paramsKey{
		Params: struct{ V any }{
			V: vUsed,
		},
		Hashes: struct{ V hash.Hash }{
			V: vUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMarshal_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMarshal_genType) AssertExpectationsMet() {
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
