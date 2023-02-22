// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package asn1

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Unmarshal_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type Unmarshal_genType func(b []byte, val any) (rest []byte, err error)

// MoqUnmarshal_genType holds the state of a moq of the Unmarshal_genType type
type MoqUnmarshal_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqUnmarshal_genType_mock

	ResultsByParams []MoqUnmarshal_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			B   moq.ParamIndexing
			Val moq.ParamIndexing
		}
	}
}

// MoqUnmarshal_genType_mock isolates the mock interface of the
// Unmarshal_genType type
type MoqUnmarshal_genType_mock struct {
	Moq *MoqUnmarshal_genType
}

// MoqUnmarshal_genType_params holds the params of the Unmarshal_genType type
type MoqUnmarshal_genType_params struct {
	B   []byte
	Val any
}

// MoqUnmarshal_genType_paramsKey holds the map key params of the
// Unmarshal_genType type
type MoqUnmarshal_genType_paramsKey struct {
	Params struct{ Val any }
	Hashes struct {
		B   hash.Hash
		Val hash.Hash
	}
}

// MoqUnmarshal_genType_resultsByParams contains the results for a given set of
// parameters for the Unmarshal_genType type
type MoqUnmarshal_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqUnmarshal_genType_paramsKey]*MoqUnmarshal_genType_results
}

// MoqUnmarshal_genType_doFn defines the type of function needed when calling
// AndDo for the Unmarshal_genType type
type MoqUnmarshal_genType_doFn func(b []byte, val any)

// MoqUnmarshal_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Unmarshal_genType type
type MoqUnmarshal_genType_doReturnFn func(b []byte, val any) (rest []byte, err error)

// MoqUnmarshal_genType_results holds the results of the Unmarshal_genType type
type MoqUnmarshal_genType_results struct {
	Params  MoqUnmarshal_genType_params
	Results []struct {
		Values *struct {
			Rest []byte
			Err  error
		}
		Sequence   uint32
		DoFn       MoqUnmarshal_genType_doFn
		DoReturnFn MoqUnmarshal_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqUnmarshal_genType_fnRecorder routes recorded function calls to the
// MoqUnmarshal_genType moq
type MoqUnmarshal_genType_fnRecorder struct {
	Params    MoqUnmarshal_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqUnmarshal_genType_results
	Moq       *MoqUnmarshal_genType
}

// MoqUnmarshal_genType_anyParams isolates the any params functions of the
// Unmarshal_genType type
type MoqUnmarshal_genType_anyParams struct {
	Recorder *MoqUnmarshal_genType_fnRecorder
}

// NewMoqUnmarshal_genType creates a new moq of the Unmarshal_genType type
func NewMoqUnmarshal_genType(scene *moq.Scene, config *moq.Config) *MoqUnmarshal_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqUnmarshal_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqUnmarshal_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				B   moq.ParamIndexing
				Val moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			B   moq.ParamIndexing
			Val moq.ParamIndexing
		}{
			B:   moq.ParamIndexByHash,
			Val: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Unmarshal_genType type
func (m *MoqUnmarshal_genType) Mock() Unmarshal_genType {
	return func(b []byte, val any) (_ []byte, _ error) {
		m.Scene.T.Helper()
		moq := &MoqUnmarshal_genType_mock{Moq: m}
		return moq.Fn(b, val)
	}
}

func (m *MoqUnmarshal_genType_mock) Fn(b []byte, val any) (rest []byte, err error) {
	m.Moq.Scene.T.Helper()
	params := MoqUnmarshal_genType_params{
		B:   b,
		Val: val,
	}
	var results *MoqUnmarshal_genType_results
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
		result.DoFn(b, val)
	}

	if result.Values != nil {
		rest = result.Values.Rest
		err = result.Values.Err
	}
	if result.DoReturnFn != nil {
		rest, err = result.DoReturnFn(b, val)
	}
	return
}

func (m *MoqUnmarshal_genType) OnCall(b []byte, val any) *MoqUnmarshal_genType_fnRecorder {
	return &MoqUnmarshal_genType_fnRecorder{
		Params: MoqUnmarshal_genType_params{
			B:   b,
			Val: val,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqUnmarshal_genType_fnRecorder) Any() *MoqUnmarshal_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqUnmarshal_genType_anyParams{Recorder: r}
}

func (a *MoqUnmarshal_genType_anyParams) B() *MoqUnmarshal_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqUnmarshal_genType_anyParams) Val() *MoqUnmarshal_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqUnmarshal_genType_fnRecorder) Seq() *MoqUnmarshal_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqUnmarshal_genType_fnRecorder) NoSeq() *MoqUnmarshal_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqUnmarshal_genType_fnRecorder) ReturnResults(rest []byte, err error) *MoqUnmarshal_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Rest []byte
			Err  error
		}
		Sequence   uint32
		DoFn       MoqUnmarshal_genType_doFn
		DoReturnFn MoqUnmarshal_genType_doReturnFn
	}{
		Values: &struct {
			Rest []byte
			Err  error
		}{
			Rest: rest,
			Err:  err,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqUnmarshal_genType_fnRecorder) AndDo(fn MoqUnmarshal_genType_doFn) *MoqUnmarshal_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqUnmarshal_genType_fnRecorder) DoReturnResults(fn MoqUnmarshal_genType_doReturnFn) *MoqUnmarshal_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Rest []byte
			Err  error
		}
		Sequence   uint32
		DoFn       MoqUnmarshal_genType_doFn
		DoReturnFn MoqUnmarshal_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqUnmarshal_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqUnmarshal_genType_resultsByParams
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
		results = &MoqUnmarshal_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqUnmarshal_genType_paramsKey]*MoqUnmarshal_genType_results{},
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
		r.Results = &MoqUnmarshal_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqUnmarshal_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqUnmarshal_genType_fnRecorder {
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
					Rest []byte
					Err  error
				}
				Sequence   uint32
				DoFn       MoqUnmarshal_genType_doFn
				DoReturnFn MoqUnmarshal_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqUnmarshal_genType) PrettyParams(params MoqUnmarshal_genType_params) string {
	return fmt.Sprintf("Unmarshal_genType(%#v, %#v)", params.B, params.Val)
}

func (m *MoqUnmarshal_genType) ParamsKey(params MoqUnmarshal_genType_params, anyParams uint64) MoqUnmarshal_genType_paramsKey {
	m.Scene.T.Helper()
	var bUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.B == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The b parameter can't be indexed by value")
		}
		bUsedHash = hash.DeepHash(params.B)
	}
	var valUsed any
	var valUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Val == moq.ParamIndexByValue {
			valUsed = params.Val
		} else {
			valUsedHash = hash.DeepHash(params.Val)
		}
	}
	return MoqUnmarshal_genType_paramsKey{
		Params: struct{ Val any }{
			Val: valUsed,
		},
		Hashes: struct {
			B   hash.Hash
			Val hash.Hash
		}{
			B:   bUsedHash,
			Val: valUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqUnmarshal_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqUnmarshal_genType) AssertExpectationsMet() {
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
