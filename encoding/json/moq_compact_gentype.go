// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package json

import (
	"bytes"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Compact_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Compact_genType func(dst *bytes.Buffer, src []byte) error

// MoqCompact_genType holds the state of a moq of the Compact_genType type
type MoqCompact_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCompact_genType_mock

	ResultsByParams []MoqCompact_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Dst moq.ParamIndexing
			Src moq.ParamIndexing
		}
	}
}

// MoqCompact_genType_mock isolates the mock interface of the Compact_genType
// type
type MoqCompact_genType_mock struct {
	Moq *MoqCompact_genType
}

// MoqCompact_genType_params holds the params of the Compact_genType type
type MoqCompact_genType_params struct {
	Dst *bytes.Buffer
	Src []byte
}

// MoqCompact_genType_paramsKey holds the map key params of the Compact_genType
// type
type MoqCompact_genType_paramsKey struct {
	Params struct{ Dst *bytes.Buffer }
	Hashes struct {
		Dst hash.Hash
		Src hash.Hash
	}
}

// MoqCompact_genType_resultsByParams contains the results for a given set of
// parameters for the Compact_genType type
type MoqCompact_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCompact_genType_paramsKey]*MoqCompact_genType_results
}

// MoqCompact_genType_doFn defines the type of function needed when calling
// AndDo for the Compact_genType type
type MoqCompact_genType_doFn func(dst *bytes.Buffer, src []byte)

// MoqCompact_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Compact_genType type
type MoqCompact_genType_doReturnFn func(dst *bytes.Buffer, src []byte) error

// MoqCompact_genType_results holds the results of the Compact_genType type
type MoqCompact_genType_results struct {
	Params  MoqCompact_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqCompact_genType_doFn
		DoReturnFn MoqCompact_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCompact_genType_fnRecorder routes recorded function calls to the
// MoqCompact_genType moq
type MoqCompact_genType_fnRecorder struct {
	Params    MoqCompact_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCompact_genType_results
	Moq       *MoqCompact_genType
}

// MoqCompact_genType_anyParams isolates the any params functions of the
// Compact_genType type
type MoqCompact_genType_anyParams struct {
	Recorder *MoqCompact_genType_fnRecorder
}

// NewMoqCompact_genType creates a new moq of the Compact_genType type
func NewMoqCompact_genType(scene *moq.Scene, config *moq.Config) *MoqCompact_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCompact_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCompact_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Dst moq.ParamIndexing
				Src moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Dst moq.ParamIndexing
			Src moq.ParamIndexing
		}{
			Dst: moq.ParamIndexByHash,
			Src: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Compact_genType type
func (m *MoqCompact_genType) Mock() Compact_genType {
	return func(dst *bytes.Buffer, src []byte) error {
		m.Scene.T.Helper()
		moq := &MoqCompact_genType_mock{Moq: m}
		return moq.Fn(dst, src)
	}
}

func (m *MoqCompact_genType_mock) Fn(dst *bytes.Buffer, src []byte) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqCompact_genType_params{
		Dst: dst,
		Src: src,
	}
	var results *MoqCompact_genType_results
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
		result.DoFn(dst, src)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(dst, src)
	}
	return
}

func (m *MoqCompact_genType) OnCall(dst *bytes.Buffer, src []byte) *MoqCompact_genType_fnRecorder {
	return &MoqCompact_genType_fnRecorder{
		Params: MoqCompact_genType_params{
			Dst: dst,
			Src: src,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqCompact_genType_fnRecorder) Any() *MoqCompact_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqCompact_genType_anyParams{Recorder: r}
}

func (a *MoqCompact_genType_anyParams) Dst() *MoqCompact_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqCompact_genType_anyParams) Src() *MoqCompact_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqCompact_genType_fnRecorder) Seq() *MoqCompact_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCompact_genType_fnRecorder) NoSeq() *MoqCompact_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCompact_genType_fnRecorder) ReturnResults(result1 error) *MoqCompact_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqCompact_genType_doFn
		DoReturnFn MoqCompact_genType_doReturnFn
	}{
		Values: &struct {
			Result1 error
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCompact_genType_fnRecorder) AndDo(fn MoqCompact_genType_doFn) *MoqCompact_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCompact_genType_fnRecorder) DoReturnResults(fn MoqCompact_genType_doReturnFn) *MoqCompact_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqCompact_genType_doFn
		DoReturnFn MoqCompact_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCompact_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCompact_genType_resultsByParams
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
		results = &MoqCompact_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCompact_genType_paramsKey]*MoqCompact_genType_results{},
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
		r.Results = &MoqCompact_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCompact_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCompact_genType_fnRecorder {
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
					Result1 error
				}
				Sequence   uint32
				DoFn       MoqCompact_genType_doFn
				DoReturnFn MoqCompact_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCompact_genType) PrettyParams(params MoqCompact_genType_params) string {
	return fmt.Sprintf("Compact_genType(%#v, %#v)", params.Dst, params.Src)
}

func (m *MoqCompact_genType) ParamsKey(params MoqCompact_genType_params, anyParams uint64) MoqCompact_genType_paramsKey {
	m.Scene.T.Helper()
	var dstUsed *bytes.Buffer
	var dstUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Dst == moq.ParamIndexByValue {
			dstUsed = params.Dst
		} else {
			dstUsedHash = hash.DeepHash(params.Dst)
		}
	}
	var srcUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Src == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The src parameter can't be indexed by value")
		}
		srcUsedHash = hash.DeepHash(params.Src)
	}
	return MoqCompact_genType_paramsKey{
		Params: struct{ Dst *bytes.Buffer }{
			Dst: dstUsed,
		},
		Hashes: struct {
			Dst hash.Hash
			Src hash.Hash
		}{
			Dst: dstUsedHash,
			Src: srcUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCompact_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCompact_genType) AssertExpectationsMet() {
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
