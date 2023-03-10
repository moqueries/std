// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package types

import (
	"fmt"
	"go/types"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Default_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Default_genType func(t types.Type) types.Type

// MoqDefault_genType holds the state of a moq of the Default_genType type
type MoqDefault_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDefault_genType_mock

	ResultsByParams []MoqDefault_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			T moq.ParamIndexing
		}
	}
}

// MoqDefault_genType_mock isolates the mock interface of the Default_genType
// type
type MoqDefault_genType_mock struct {
	Moq *MoqDefault_genType
}

// MoqDefault_genType_params holds the params of the Default_genType type
type MoqDefault_genType_params struct{ T types.Type }

// MoqDefault_genType_paramsKey holds the map key params of the Default_genType
// type
type MoqDefault_genType_paramsKey struct {
	Params struct{ T types.Type }
	Hashes struct{ T hash.Hash }
}

// MoqDefault_genType_resultsByParams contains the results for a given set of
// parameters for the Default_genType type
type MoqDefault_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDefault_genType_paramsKey]*MoqDefault_genType_results
}

// MoqDefault_genType_doFn defines the type of function needed when calling
// AndDo for the Default_genType type
type MoqDefault_genType_doFn func(t types.Type)

// MoqDefault_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Default_genType type
type MoqDefault_genType_doReturnFn func(t types.Type) types.Type

// MoqDefault_genType_results holds the results of the Default_genType type
type MoqDefault_genType_results struct {
	Params  MoqDefault_genType_params
	Results []struct {
		Values *struct {
			Result1 types.Type
		}
		Sequence   uint32
		DoFn       MoqDefault_genType_doFn
		DoReturnFn MoqDefault_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDefault_genType_fnRecorder routes recorded function calls to the
// MoqDefault_genType moq
type MoqDefault_genType_fnRecorder struct {
	Params    MoqDefault_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDefault_genType_results
	Moq       *MoqDefault_genType
}

// MoqDefault_genType_anyParams isolates the any params functions of the
// Default_genType type
type MoqDefault_genType_anyParams struct {
	Recorder *MoqDefault_genType_fnRecorder
}

// NewMoqDefault_genType creates a new moq of the Default_genType type
func NewMoqDefault_genType(scene *moq.Scene, config *moq.Config) *MoqDefault_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDefault_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDefault_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				T moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			T moq.ParamIndexing
		}{
			T: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Default_genType type
func (m *MoqDefault_genType) Mock() Default_genType {
	return func(t types.Type) types.Type {
		m.Scene.T.Helper()
		moq := &MoqDefault_genType_mock{Moq: m}
		return moq.Fn(t)
	}
}

func (m *MoqDefault_genType_mock) Fn(t types.Type) (result1 types.Type) {
	m.Moq.Scene.T.Helper()
	params := MoqDefault_genType_params{
		T: t,
	}
	var results *MoqDefault_genType_results
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
		result.DoFn(t)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(t)
	}
	return
}

func (m *MoqDefault_genType) OnCall(t types.Type) *MoqDefault_genType_fnRecorder {
	return &MoqDefault_genType_fnRecorder{
		Params: MoqDefault_genType_params{
			T: t,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqDefault_genType_fnRecorder) Any() *MoqDefault_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqDefault_genType_anyParams{Recorder: r}
}

func (a *MoqDefault_genType_anyParams) T() *MoqDefault_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqDefault_genType_fnRecorder) Seq() *MoqDefault_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDefault_genType_fnRecorder) NoSeq() *MoqDefault_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDefault_genType_fnRecorder) ReturnResults(result1 types.Type) *MoqDefault_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 types.Type
		}
		Sequence   uint32
		DoFn       MoqDefault_genType_doFn
		DoReturnFn MoqDefault_genType_doReturnFn
	}{
		Values: &struct {
			Result1 types.Type
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDefault_genType_fnRecorder) AndDo(fn MoqDefault_genType_doFn) *MoqDefault_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDefault_genType_fnRecorder) DoReturnResults(fn MoqDefault_genType_doReturnFn) *MoqDefault_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 types.Type
		}
		Sequence   uint32
		DoFn       MoqDefault_genType_doFn
		DoReturnFn MoqDefault_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDefault_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDefault_genType_resultsByParams
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
		results = &MoqDefault_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDefault_genType_paramsKey]*MoqDefault_genType_results{},
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
		r.Results = &MoqDefault_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDefault_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDefault_genType_fnRecorder {
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
					Result1 types.Type
				}
				Sequence   uint32
				DoFn       MoqDefault_genType_doFn
				DoReturnFn MoqDefault_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDefault_genType) PrettyParams(params MoqDefault_genType_params) string {
	return fmt.Sprintf("Default_genType(%#v)", params.T)
}

func (m *MoqDefault_genType) ParamsKey(params MoqDefault_genType_params, anyParams uint64) MoqDefault_genType_paramsKey {
	m.Scene.T.Helper()
	var tUsed types.Type
	var tUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.T == moq.ParamIndexByValue {
			tUsed = params.T
		} else {
			tUsedHash = hash.DeepHash(params.T)
		}
	}
	return MoqDefault_genType_paramsKey{
		Params: struct{ T types.Type }{
			T: tUsed,
		},
		Hashes: struct{ T hash.Hash }{
			T: tUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqDefault_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDefault_genType) AssertExpectationsMet() {
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
