// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package elf

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// R_SYM32_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type R_SYM32_genType func(info uint32) uint32

// MoqR_SYM32_genType holds the state of a moq of the R_SYM32_genType type
type MoqR_SYM32_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqR_SYM32_genType_mock

	ResultsByParams []MoqR_SYM32_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Info moq.ParamIndexing
		}
	}
}

// MoqR_SYM32_genType_mock isolates the mock interface of the R_SYM32_genType
// type
type MoqR_SYM32_genType_mock struct {
	Moq *MoqR_SYM32_genType
}

// MoqR_SYM32_genType_params holds the params of the R_SYM32_genType type
type MoqR_SYM32_genType_params struct{ Info uint32 }

// MoqR_SYM32_genType_paramsKey holds the map key params of the R_SYM32_genType
// type
type MoqR_SYM32_genType_paramsKey struct {
	Params struct{ Info uint32 }
	Hashes struct{ Info hash.Hash }
}

// MoqR_SYM32_genType_resultsByParams contains the results for a given set of
// parameters for the R_SYM32_genType type
type MoqR_SYM32_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqR_SYM32_genType_paramsKey]*MoqR_SYM32_genType_results
}

// MoqR_SYM32_genType_doFn defines the type of function needed when calling
// AndDo for the R_SYM32_genType type
type MoqR_SYM32_genType_doFn func(info uint32)

// MoqR_SYM32_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the R_SYM32_genType type
type MoqR_SYM32_genType_doReturnFn func(info uint32) uint32

// MoqR_SYM32_genType_results holds the results of the R_SYM32_genType type
type MoqR_SYM32_genType_results struct {
	Params  MoqR_SYM32_genType_params
	Results []struct {
		Values *struct {
			Result1 uint32
		}
		Sequence   uint32
		DoFn       MoqR_SYM32_genType_doFn
		DoReturnFn MoqR_SYM32_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqR_SYM32_genType_fnRecorder routes recorded function calls to the
// MoqR_SYM32_genType moq
type MoqR_SYM32_genType_fnRecorder struct {
	Params    MoqR_SYM32_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqR_SYM32_genType_results
	Moq       *MoqR_SYM32_genType
}

// MoqR_SYM32_genType_anyParams isolates the any params functions of the
// R_SYM32_genType type
type MoqR_SYM32_genType_anyParams struct {
	Recorder *MoqR_SYM32_genType_fnRecorder
}

// NewMoqR_SYM32_genType creates a new moq of the R_SYM32_genType type
func NewMoqR_SYM32_genType(scene *moq.Scene, config *moq.Config) *MoqR_SYM32_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqR_SYM32_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqR_SYM32_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Info moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Info moq.ParamIndexing
		}{
			Info: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the R_SYM32_genType type
func (m *MoqR_SYM32_genType) Mock() R_SYM32_genType {
	return func(info uint32) uint32 {
		m.Scene.T.Helper()
		moq := &MoqR_SYM32_genType_mock{Moq: m}
		return moq.Fn(info)
	}
}

func (m *MoqR_SYM32_genType_mock) Fn(info uint32) (result1 uint32) {
	m.Moq.Scene.T.Helper()
	params := MoqR_SYM32_genType_params{
		Info: info,
	}
	var results *MoqR_SYM32_genType_results
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
		result.DoFn(info)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(info)
	}
	return
}

func (m *MoqR_SYM32_genType) OnCall(info uint32) *MoqR_SYM32_genType_fnRecorder {
	return &MoqR_SYM32_genType_fnRecorder{
		Params: MoqR_SYM32_genType_params{
			Info: info,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqR_SYM32_genType_fnRecorder) Any() *MoqR_SYM32_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqR_SYM32_genType_anyParams{Recorder: r}
}

func (a *MoqR_SYM32_genType_anyParams) Info() *MoqR_SYM32_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqR_SYM32_genType_fnRecorder) Seq() *MoqR_SYM32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqR_SYM32_genType_fnRecorder) NoSeq() *MoqR_SYM32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqR_SYM32_genType_fnRecorder) ReturnResults(result1 uint32) *MoqR_SYM32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 uint32
		}
		Sequence   uint32
		DoFn       MoqR_SYM32_genType_doFn
		DoReturnFn MoqR_SYM32_genType_doReturnFn
	}{
		Values: &struct {
			Result1 uint32
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqR_SYM32_genType_fnRecorder) AndDo(fn MoqR_SYM32_genType_doFn) *MoqR_SYM32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqR_SYM32_genType_fnRecorder) DoReturnResults(fn MoqR_SYM32_genType_doReturnFn) *MoqR_SYM32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 uint32
		}
		Sequence   uint32
		DoFn       MoqR_SYM32_genType_doFn
		DoReturnFn MoqR_SYM32_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqR_SYM32_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqR_SYM32_genType_resultsByParams
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
		results = &MoqR_SYM32_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqR_SYM32_genType_paramsKey]*MoqR_SYM32_genType_results{},
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
		r.Results = &MoqR_SYM32_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqR_SYM32_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqR_SYM32_genType_fnRecorder {
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
					Result1 uint32
				}
				Sequence   uint32
				DoFn       MoqR_SYM32_genType_doFn
				DoReturnFn MoqR_SYM32_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqR_SYM32_genType) PrettyParams(params MoqR_SYM32_genType_params) string {
	return fmt.Sprintf("R_SYM32_genType(%#v)", params.Info)
}

func (m *MoqR_SYM32_genType) ParamsKey(params MoqR_SYM32_genType_params, anyParams uint64) MoqR_SYM32_genType_paramsKey {
	m.Scene.T.Helper()
	var infoUsed uint32
	var infoUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Info == moq.ParamIndexByValue {
			infoUsed = params.Info
		} else {
			infoUsedHash = hash.DeepHash(params.Info)
		}
	}
	return MoqR_SYM32_genType_paramsKey{
		Params: struct{ Info uint32 }{
			Info: infoUsed,
		},
		Hashes: struct{ Info hash.Hash }{
			Info: infoUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqR_SYM32_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqR_SYM32_genType) AssertExpectationsMet() {
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
