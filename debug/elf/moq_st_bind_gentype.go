// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package elf

import (
	"debug/elf"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// ST_BIND_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type ST_BIND_genType func(info uint8) elf.SymBind

// MoqST_BIND_genType holds the state of a moq of the ST_BIND_genType type
type MoqST_BIND_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqST_BIND_genType_mock

	ResultsByParams []MoqST_BIND_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Info moq.ParamIndexing
		}
	}
}

// MoqST_BIND_genType_mock isolates the mock interface of the ST_BIND_genType
// type
type MoqST_BIND_genType_mock struct {
	Moq *MoqST_BIND_genType
}

// MoqST_BIND_genType_params holds the params of the ST_BIND_genType type
type MoqST_BIND_genType_params struct{ Info uint8 }

// MoqST_BIND_genType_paramsKey holds the map key params of the ST_BIND_genType
// type
type MoqST_BIND_genType_paramsKey struct {
	Params struct{ Info uint8 }
	Hashes struct{ Info hash.Hash }
}

// MoqST_BIND_genType_resultsByParams contains the results for a given set of
// parameters for the ST_BIND_genType type
type MoqST_BIND_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqST_BIND_genType_paramsKey]*MoqST_BIND_genType_results
}

// MoqST_BIND_genType_doFn defines the type of function needed when calling
// AndDo for the ST_BIND_genType type
type MoqST_BIND_genType_doFn func(info uint8)

// MoqST_BIND_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the ST_BIND_genType type
type MoqST_BIND_genType_doReturnFn func(info uint8) elf.SymBind

// MoqST_BIND_genType_results holds the results of the ST_BIND_genType type
type MoqST_BIND_genType_results struct {
	Params  MoqST_BIND_genType_params
	Results []struct {
		Values *struct {
			Result1 elf.SymBind
		}
		Sequence   uint32
		DoFn       MoqST_BIND_genType_doFn
		DoReturnFn MoqST_BIND_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqST_BIND_genType_fnRecorder routes recorded function calls to the
// MoqST_BIND_genType moq
type MoqST_BIND_genType_fnRecorder struct {
	Params    MoqST_BIND_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqST_BIND_genType_results
	Moq       *MoqST_BIND_genType
}

// MoqST_BIND_genType_anyParams isolates the any params functions of the
// ST_BIND_genType type
type MoqST_BIND_genType_anyParams struct {
	Recorder *MoqST_BIND_genType_fnRecorder
}

// NewMoqST_BIND_genType creates a new moq of the ST_BIND_genType type
func NewMoqST_BIND_genType(scene *moq.Scene, config *moq.Config) *MoqST_BIND_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqST_BIND_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqST_BIND_genType_mock{},

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

// Mock returns the moq implementation of the ST_BIND_genType type
func (m *MoqST_BIND_genType) Mock() ST_BIND_genType {
	return func(info uint8) elf.SymBind {
		m.Scene.T.Helper()
		moq := &MoqST_BIND_genType_mock{Moq: m}
		return moq.Fn(info)
	}
}

func (m *MoqST_BIND_genType_mock) Fn(info uint8) (result1 elf.SymBind) {
	m.Moq.Scene.T.Helper()
	params := MoqST_BIND_genType_params{
		Info: info,
	}
	var results *MoqST_BIND_genType_results
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

func (m *MoqST_BIND_genType) OnCall(info uint8) *MoqST_BIND_genType_fnRecorder {
	return &MoqST_BIND_genType_fnRecorder{
		Params: MoqST_BIND_genType_params{
			Info: info,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqST_BIND_genType_fnRecorder) Any() *MoqST_BIND_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqST_BIND_genType_anyParams{Recorder: r}
}

func (a *MoqST_BIND_genType_anyParams) Info() *MoqST_BIND_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqST_BIND_genType_fnRecorder) Seq() *MoqST_BIND_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqST_BIND_genType_fnRecorder) NoSeq() *MoqST_BIND_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqST_BIND_genType_fnRecorder) ReturnResults(result1 elf.SymBind) *MoqST_BIND_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 elf.SymBind
		}
		Sequence   uint32
		DoFn       MoqST_BIND_genType_doFn
		DoReturnFn MoqST_BIND_genType_doReturnFn
	}{
		Values: &struct {
			Result1 elf.SymBind
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqST_BIND_genType_fnRecorder) AndDo(fn MoqST_BIND_genType_doFn) *MoqST_BIND_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqST_BIND_genType_fnRecorder) DoReturnResults(fn MoqST_BIND_genType_doReturnFn) *MoqST_BIND_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 elf.SymBind
		}
		Sequence   uint32
		DoFn       MoqST_BIND_genType_doFn
		DoReturnFn MoqST_BIND_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqST_BIND_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqST_BIND_genType_resultsByParams
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
		results = &MoqST_BIND_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqST_BIND_genType_paramsKey]*MoqST_BIND_genType_results{},
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
		r.Results = &MoqST_BIND_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqST_BIND_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqST_BIND_genType_fnRecorder {
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
					Result1 elf.SymBind
				}
				Sequence   uint32
				DoFn       MoqST_BIND_genType_doFn
				DoReturnFn MoqST_BIND_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqST_BIND_genType) PrettyParams(params MoqST_BIND_genType_params) string {
	return fmt.Sprintf("ST_BIND_genType(%#v)", params.Info)
}

func (m *MoqST_BIND_genType) ParamsKey(params MoqST_BIND_genType_params, anyParams uint64) MoqST_BIND_genType_paramsKey {
	m.Scene.T.Helper()
	var infoUsed uint8
	var infoUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Info == moq.ParamIndexByValue {
			infoUsed = params.Info
		} else {
			infoUsedHash = hash.DeepHash(params.Info)
		}
	}
	return MoqST_BIND_genType_paramsKey{
		Params: struct{ Info uint8 }{
			Info: infoUsed,
		},
		Hashes: struct{ Info hash.Hash }{
			Info: infoUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqST_BIND_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqST_BIND_genType) AssertExpectationsMet() {
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
