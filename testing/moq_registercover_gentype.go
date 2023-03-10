// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package testing

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"testing"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// RegisterCover_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type RegisterCover_genType func(c testing.Cover)

// MoqRegisterCover_genType holds the state of a moq of the
// RegisterCover_genType type
type MoqRegisterCover_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqRegisterCover_genType_mock

	ResultsByParams []MoqRegisterCover_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			C moq.ParamIndexing
		}
	}
}

// MoqRegisterCover_genType_mock isolates the mock interface of the
// RegisterCover_genType type
type MoqRegisterCover_genType_mock struct {
	Moq *MoqRegisterCover_genType
}

// MoqRegisterCover_genType_params holds the params of the
// RegisterCover_genType type
type MoqRegisterCover_genType_params struct{ C testing.Cover }

// MoqRegisterCover_genType_paramsKey holds the map key params of the
// RegisterCover_genType type
type MoqRegisterCover_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ C hash.Hash }
}

// MoqRegisterCover_genType_resultsByParams contains the results for a given
// set of parameters for the RegisterCover_genType type
type MoqRegisterCover_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRegisterCover_genType_paramsKey]*MoqRegisterCover_genType_results
}

// MoqRegisterCover_genType_doFn defines the type of function needed when
// calling AndDo for the RegisterCover_genType type
type MoqRegisterCover_genType_doFn func(c testing.Cover)

// MoqRegisterCover_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the RegisterCover_genType type
type MoqRegisterCover_genType_doReturnFn func(c testing.Cover)

// MoqRegisterCover_genType_results holds the results of the
// RegisterCover_genType type
type MoqRegisterCover_genType_results struct {
	Params  MoqRegisterCover_genType_params
	Results []struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqRegisterCover_genType_doFn
		DoReturnFn MoqRegisterCover_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRegisterCover_genType_fnRecorder routes recorded function calls to the
// MoqRegisterCover_genType moq
type MoqRegisterCover_genType_fnRecorder struct {
	Params    MoqRegisterCover_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRegisterCover_genType_results
	Moq       *MoqRegisterCover_genType
}

// MoqRegisterCover_genType_anyParams isolates the any params functions of the
// RegisterCover_genType type
type MoqRegisterCover_genType_anyParams struct {
	Recorder *MoqRegisterCover_genType_fnRecorder
}

// NewMoqRegisterCover_genType creates a new moq of the RegisterCover_genType
// type
func NewMoqRegisterCover_genType(scene *moq.Scene, config *moq.Config) *MoqRegisterCover_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqRegisterCover_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqRegisterCover_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				C moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			C moq.ParamIndexing
		}{
			C: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the RegisterCover_genType type
func (m *MoqRegisterCover_genType) Mock() RegisterCover_genType {
	return func(c testing.Cover) { m.Scene.T.Helper(); moq := &MoqRegisterCover_genType_mock{Moq: m}; moq.Fn(c) }
}

func (m *MoqRegisterCover_genType_mock) Fn(c testing.Cover) {
	m.Moq.Scene.T.Helper()
	params := MoqRegisterCover_genType_params{
		C: c,
	}
	var results *MoqRegisterCover_genType_results
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
		result.DoFn(c)
	}

	if result.DoReturnFn != nil {
		result.DoReturnFn(c)
	}
	return
}

func (m *MoqRegisterCover_genType) OnCall(c testing.Cover) *MoqRegisterCover_genType_fnRecorder {
	return &MoqRegisterCover_genType_fnRecorder{
		Params: MoqRegisterCover_genType_params{
			C: c,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqRegisterCover_genType_fnRecorder) Any() *MoqRegisterCover_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqRegisterCover_genType_anyParams{Recorder: r}
}

func (a *MoqRegisterCover_genType_anyParams) C() *MoqRegisterCover_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqRegisterCover_genType_fnRecorder) Seq() *MoqRegisterCover_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRegisterCover_genType_fnRecorder) NoSeq() *MoqRegisterCover_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRegisterCover_genType_fnRecorder) ReturnResults() *MoqRegisterCover_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqRegisterCover_genType_doFn
		DoReturnFn MoqRegisterCover_genType_doReturnFn
	}{
		Values:   &struct{}{},
		Sequence: sequence,
	})
	return r
}

func (r *MoqRegisterCover_genType_fnRecorder) AndDo(fn MoqRegisterCover_genType_doFn) *MoqRegisterCover_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRegisterCover_genType_fnRecorder) DoReturnResults(fn MoqRegisterCover_genType_doReturnFn) *MoqRegisterCover_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{}
		Sequence   uint32
		DoFn       MoqRegisterCover_genType_doFn
		DoReturnFn MoqRegisterCover_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRegisterCover_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRegisterCover_genType_resultsByParams
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
		results = &MoqRegisterCover_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRegisterCover_genType_paramsKey]*MoqRegisterCover_genType_results{},
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
		r.Results = &MoqRegisterCover_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRegisterCover_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRegisterCover_genType_fnRecorder {
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
				DoFn       MoqRegisterCover_genType_doFn
				DoReturnFn MoqRegisterCover_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRegisterCover_genType) PrettyParams(params MoqRegisterCover_genType_params) string {
	return fmt.Sprintf("RegisterCover_genType(%#v)", params.C)
}

func (m *MoqRegisterCover_genType) ParamsKey(params MoqRegisterCover_genType_params, anyParams uint64) MoqRegisterCover_genType_paramsKey {
	m.Scene.T.Helper()
	var cUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.C == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The c parameter can't be indexed by value")
		}
		cUsedHash = hash.DeepHash(params.C)
	}
	return MoqRegisterCover_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ C hash.Hash }{
			C: cUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqRegisterCover_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqRegisterCover_genType) AssertExpectationsMet() {
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
