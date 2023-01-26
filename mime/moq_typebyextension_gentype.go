// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package mime

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// TypeByExtension_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type TypeByExtension_genType func(ext string) string

// MoqTypeByExtension_genType holds the state of a moq of the
// TypeByExtension_genType type
type MoqTypeByExtension_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqTypeByExtension_genType_mock

	ResultsByParams []MoqTypeByExtension_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Ext moq.ParamIndexing
		}
	}
}

// MoqTypeByExtension_genType_mock isolates the mock interface of the
// TypeByExtension_genType type
type MoqTypeByExtension_genType_mock struct {
	Moq *MoqTypeByExtension_genType
}

// MoqTypeByExtension_genType_params holds the params of the
// TypeByExtension_genType type
type MoqTypeByExtension_genType_params struct{ Ext string }

// MoqTypeByExtension_genType_paramsKey holds the map key params of the
// TypeByExtension_genType type
type MoqTypeByExtension_genType_paramsKey struct {
	Params struct{ Ext string }
	Hashes struct{ Ext hash.Hash }
}

// MoqTypeByExtension_genType_resultsByParams contains the results for a given
// set of parameters for the TypeByExtension_genType type
type MoqTypeByExtension_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqTypeByExtension_genType_paramsKey]*MoqTypeByExtension_genType_results
}

// MoqTypeByExtension_genType_doFn defines the type of function needed when
// calling AndDo for the TypeByExtension_genType type
type MoqTypeByExtension_genType_doFn func(ext string)

// MoqTypeByExtension_genType_doReturnFn defines the type of function needed
// when calling DoReturnResults for the TypeByExtension_genType type
type MoqTypeByExtension_genType_doReturnFn func(ext string) string

// MoqTypeByExtension_genType_results holds the results of the
// TypeByExtension_genType type
type MoqTypeByExtension_genType_results struct {
	Params  MoqTypeByExtension_genType_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqTypeByExtension_genType_doFn
		DoReturnFn MoqTypeByExtension_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqTypeByExtension_genType_fnRecorder routes recorded function calls to the
// MoqTypeByExtension_genType moq
type MoqTypeByExtension_genType_fnRecorder struct {
	Params    MoqTypeByExtension_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqTypeByExtension_genType_results
	Moq       *MoqTypeByExtension_genType
}

// MoqTypeByExtension_genType_anyParams isolates the any params functions of
// the TypeByExtension_genType type
type MoqTypeByExtension_genType_anyParams struct {
	Recorder *MoqTypeByExtension_genType_fnRecorder
}

// NewMoqTypeByExtension_genType creates a new moq of the
// TypeByExtension_genType type
func NewMoqTypeByExtension_genType(scene *moq.Scene, config *moq.Config) *MoqTypeByExtension_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqTypeByExtension_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqTypeByExtension_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Ext moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Ext moq.ParamIndexing
		}{
			Ext: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the TypeByExtension_genType type
func (m *MoqTypeByExtension_genType) Mock() TypeByExtension_genType {
	return func(ext string) string {
		m.Scene.T.Helper()
		moq := &MoqTypeByExtension_genType_mock{Moq: m}
		return moq.Fn(ext)
	}
}

func (m *MoqTypeByExtension_genType_mock) Fn(ext string) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqTypeByExtension_genType_params{
		Ext: ext,
	}
	var results *MoqTypeByExtension_genType_results
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
		result.DoFn(ext)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(ext)
	}
	return
}

func (m *MoqTypeByExtension_genType) OnCall(ext string) *MoqTypeByExtension_genType_fnRecorder {
	return &MoqTypeByExtension_genType_fnRecorder{
		Params: MoqTypeByExtension_genType_params{
			Ext: ext,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqTypeByExtension_genType_fnRecorder) Any() *MoqTypeByExtension_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqTypeByExtension_genType_anyParams{Recorder: r}
}

func (a *MoqTypeByExtension_genType_anyParams) Ext() *MoqTypeByExtension_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqTypeByExtension_genType_fnRecorder) Seq() *MoqTypeByExtension_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqTypeByExtension_genType_fnRecorder) NoSeq() *MoqTypeByExtension_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqTypeByExtension_genType_fnRecorder) ReturnResults(result1 string) *MoqTypeByExtension_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqTypeByExtension_genType_doFn
		DoReturnFn MoqTypeByExtension_genType_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqTypeByExtension_genType_fnRecorder) AndDo(fn MoqTypeByExtension_genType_doFn) *MoqTypeByExtension_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqTypeByExtension_genType_fnRecorder) DoReturnResults(fn MoqTypeByExtension_genType_doReturnFn) *MoqTypeByExtension_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqTypeByExtension_genType_doFn
		DoReturnFn MoqTypeByExtension_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqTypeByExtension_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqTypeByExtension_genType_resultsByParams
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
		results = &MoqTypeByExtension_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqTypeByExtension_genType_paramsKey]*MoqTypeByExtension_genType_results{},
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
		r.Results = &MoqTypeByExtension_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqTypeByExtension_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqTypeByExtension_genType_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqTypeByExtension_genType_doFn
				DoReturnFn MoqTypeByExtension_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqTypeByExtension_genType) PrettyParams(params MoqTypeByExtension_genType_params) string {
	return fmt.Sprintf("TypeByExtension_genType(%#v)", params.Ext)
}

func (m *MoqTypeByExtension_genType) ParamsKey(params MoqTypeByExtension_genType_params, anyParams uint64) MoqTypeByExtension_genType_paramsKey {
	m.Scene.T.Helper()
	var extUsed string
	var extUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Ext == moq.ParamIndexByValue {
			extUsed = params.Ext
		} else {
			extUsedHash = hash.DeepHash(params.Ext)
		}
	}
	return MoqTypeByExtension_genType_paramsKey{
		Params: struct{ Ext string }{
			Ext: extUsed,
		},
		Hashes: struct{ Ext hash.Hash }{
			Ext: extUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqTypeByExtension_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqTypeByExtension_genType) AssertExpectationsMet() {
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
