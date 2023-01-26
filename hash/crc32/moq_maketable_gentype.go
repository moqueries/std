// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package crc32

import (
	"fmt"
	"hash/crc32"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// MakeTable_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type MakeTable_genType func(poly uint32) *crc32.Table

// MoqMakeTable_genType holds the state of a moq of the MakeTable_genType type
type MoqMakeTable_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqMakeTable_genType_mock

	ResultsByParams []MoqMakeTable_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Poly moq.ParamIndexing
		}
	}
}

// MoqMakeTable_genType_mock isolates the mock interface of the
// MakeTable_genType type
type MoqMakeTable_genType_mock struct {
	Moq *MoqMakeTable_genType
}

// MoqMakeTable_genType_params holds the params of the MakeTable_genType type
type MoqMakeTable_genType_params struct{ Poly uint32 }

// MoqMakeTable_genType_paramsKey holds the map key params of the
// MakeTable_genType type
type MoqMakeTable_genType_paramsKey struct {
	Params struct{ Poly uint32 }
	Hashes struct{ Poly hash.Hash }
}

// MoqMakeTable_genType_resultsByParams contains the results for a given set of
// parameters for the MakeTable_genType type
type MoqMakeTable_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqMakeTable_genType_paramsKey]*MoqMakeTable_genType_results
}

// MoqMakeTable_genType_doFn defines the type of function needed when calling
// AndDo for the MakeTable_genType type
type MoqMakeTable_genType_doFn func(poly uint32)

// MoqMakeTable_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the MakeTable_genType type
type MoqMakeTable_genType_doReturnFn func(poly uint32) *crc32.Table

// MoqMakeTable_genType_results holds the results of the MakeTable_genType type
type MoqMakeTable_genType_results struct {
	Params  MoqMakeTable_genType_params
	Results []struct {
		Values *struct {
			Result1 *crc32.Table
		}
		Sequence   uint32
		DoFn       MoqMakeTable_genType_doFn
		DoReturnFn MoqMakeTable_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqMakeTable_genType_fnRecorder routes recorded function calls to the
// MoqMakeTable_genType moq
type MoqMakeTable_genType_fnRecorder struct {
	Params    MoqMakeTable_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqMakeTable_genType_results
	Moq       *MoqMakeTable_genType
}

// MoqMakeTable_genType_anyParams isolates the any params functions of the
// MakeTable_genType type
type MoqMakeTable_genType_anyParams struct {
	Recorder *MoqMakeTable_genType_fnRecorder
}

// NewMoqMakeTable_genType creates a new moq of the MakeTable_genType type
func NewMoqMakeTable_genType(scene *moq.Scene, config *moq.Config) *MoqMakeTable_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqMakeTable_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqMakeTable_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Poly moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Poly moq.ParamIndexing
		}{
			Poly: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the MakeTable_genType type
func (m *MoqMakeTable_genType) Mock() MakeTable_genType {
	return func(poly uint32) *crc32.Table {
		m.Scene.T.Helper()
		moq := &MoqMakeTable_genType_mock{Moq: m}
		return moq.Fn(poly)
	}
}

func (m *MoqMakeTable_genType_mock) Fn(poly uint32) (result1 *crc32.Table) {
	m.Moq.Scene.T.Helper()
	params := MoqMakeTable_genType_params{
		Poly: poly,
	}
	var results *MoqMakeTable_genType_results
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
		result.DoFn(poly)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(poly)
	}
	return
}

func (m *MoqMakeTable_genType) OnCall(poly uint32) *MoqMakeTable_genType_fnRecorder {
	return &MoqMakeTable_genType_fnRecorder{
		Params: MoqMakeTable_genType_params{
			Poly: poly,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqMakeTable_genType_fnRecorder) Any() *MoqMakeTable_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqMakeTable_genType_anyParams{Recorder: r}
}

func (a *MoqMakeTable_genType_anyParams) Poly() *MoqMakeTable_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqMakeTable_genType_fnRecorder) Seq() *MoqMakeTable_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqMakeTable_genType_fnRecorder) NoSeq() *MoqMakeTable_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqMakeTable_genType_fnRecorder) ReturnResults(result1 *crc32.Table) *MoqMakeTable_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *crc32.Table
		}
		Sequence   uint32
		DoFn       MoqMakeTable_genType_doFn
		DoReturnFn MoqMakeTable_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *crc32.Table
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqMakeTable_genType_fnRecorder) AndDo(fn MoqMakeTable_genType_doFn) *MoqMakeTable_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqMakeTable_genType_fnRecorder) DoReturnResults(fn MoqMakeTable_genType_doReturnFn) *MoqMakeTable_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *crc32.Table
		}
		Sequence   uint32
		DoFn       MoqMakeTable_genType_doFn
		DoReturnFn MoqMakeTable_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqMakeTable_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqMakeTable_genType_resultsByParams
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
		results = &MoqMakeTable_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqMakeTable_genType_paramsKey]*MoqMakeTable_genType_results{},
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
		r.Results = &MoqMakeTable_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqMakeTable_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqMakeTable_genType_fnRecorder {
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
					Result1 *crc32.Table
				}
				Sequence   uint32
				DoFn       MoqMakeTable_genType_doFn
				DoReturnFn MoqMakeTable_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqMakeTable_genType) PrettyParams(params MoqMakeTable_genType_params) string {
	return fmt.Sprintf("MakeTable_genType(%#v)", params.Poly)
}

func (m *MoqMakeTable_genType) ParamsKey(params MoqMakeTable_genType_params, anyParams uint64) MoqMakeTable_genType_paramsKey {
	m.Scene.T.Helper()
	var polyUsed uint32
	var polyUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Poly == moq.ParamIndexByValue {
			polyUsed = params.Poly
		} else {
			polyUsedHash = hash.DeepHash(params.Poly)
		}
	}
	return MoqMakeTable_genType_paramsKey{
		Params: struct{ Poly uint32 }{
			Poly: polyUsed,
		},
		Hashes: struct{ Poly hash.Hash }{
			Poly: polyUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqMakeTable_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqMakeTable_genType) AssertExpectationsMet() {
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