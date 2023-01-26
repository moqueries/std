// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package template

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// HTMLEscaper_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type HTMLEscaper_genType func(args ...interface{}) string

// MoqHTMLEscaper_genType holds the state of a moq of the HTMLEscaper_genType
// type
type MoqHTMLEscaper_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqHTMLEscaper_genType_mock

	ResultsByParams []MoqHTMLEscaper_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Args moq.ParamIndexing
		}
	}
}

// MoqHTMLEscaper_genType_mock isolates the mock interface of the
// HTMLEscaper_genType type
type MoqHTMLEscaper_genType_mock struct {
	Moq *MoqHTMLEscaper_genType
}

// MoqHTMLEscaper_genType_params holds the params of the HTMLEscaper_genType
// type
type MoqHTMLEscaper_genType_params struct{ Args []interface{} }

// MoqHTMLEscaper_genType_paramsKey holds the map key params of the
// HTMLEscaper_genType type
type MoqHTMLEscaper_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ Args hash.Hash }
}

// MoqHTMLEscaper_genType_resultsByParams contains the results for a given set
// of parameters for the HTMLEscaper_genType type
type MoqHTMLEscaper_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqHTMLEscaper_genType_paramsKey]*MoqHTMLEscaper_genType_results
}

// MoqHTMLEscaper_genType_doFn defines the type of function needed when calling
// AndDo for the HTMLEscaper_genType type
type MoqHTMLEscaper_genType_doFn func(args ...interface{})

// MoqHTMLEscaper_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the HTMLEscaper_genType type
type MoqHTMLEscaper_genType_doReturnFn func(args ...interface{}) string

// MoqHTMLEscaper_genType_results holds the results of the HTMLEscaper_genType
// type
type MoqHTMLEscaper_genType_results struct {
	Params  MoqHTMLEscaper_genType_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqHTMLEscaper_genType_doFn
		DoReturnFn MoqHTMLEscaper_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqHTMLEscaper_genType_fnRecorder routes recorded function calls to the
// MoqHTMLEscaper_genType moq
type MoqHTMLEscaper_genType_fnRecorder struct {
	Params    MoqHTMLEscaper_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqHTMLEscaper_genType_results
	Moq       *MoqHTMLEscaper_genType
}

// MoqHTMLEscaper_genType_anyParams isolates the any params functions of the
// HTMLEscaper_genType type
type MoqHTMLEscaper_genType_anyParams struct {
	Recorder *MoqHTMLEscaper_genType_fnRecorder
}

// NewMoqHTMLEscaper_genType creates a new moq of the HTMLEscaper_genType type
func NewMoqHTMLEscaper_genType(scene *moq.Scene, config *moq.Config) *MoqHTMLEscaper_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqHTMLEscaper_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqHTMLEscaper_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Args moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Args moq.ParamIndexing
		}{
			Args: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the HTMLEscaper_genType type
func (m *MoqHTMLEscaper_genType) Mock() HTMLEscaper_genType {
	return func(args ...interface{}) string {
		m.Scene.T.Helper()
		moq := &MoqHTMLEscaper_genType_mock{Moq: m}
		return moq.Fn(args...)
	}
}

func (m *MoqHTMLEscaper_genType_mock) Fn(args ...interface{}) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqHTMLEscaper_genType_params{
		Args: args,
	}
	var results *MoqHTMLEscaper_genType_results
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
		result.DoFn(args...)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(args...)
	}
	return
}

func (m *MoqHTMLEscaper_genType) OnCall(args ...interface{}) *MoqHTMLEscaper_genType_fnRecorder {
	return &MoqHTMLEscaper_genType_fnRecorder{
		Params: MoqHTMLEscaper_genType_params{
			Args: args,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqHTMLEscaper_genType_fnRecorder) Any() *MoqHTMLEscaper_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqHTMLEscaper_genType_anyParams{Recorder: r}
}

func (a *MoqHTMLEscaper_genType_anyParams) Args() *MoqHTMLEscaper_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqHTMLEscaper_genType_fnRecorder) Seq() *MoqHTMLEscaper_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqHTMLEscaper_genType_fnRecorder) NoSeq() *MoqHTMLEscaper_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqHTMLEscaper_genType_fnRecorder) ReturnResults(result1 string) *MoqHTMLEscaper_genType_fnRecorder {
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
		DoFn       MoqHTMLEscaper_genType_doFn
		DoReturnFn MoqHTMLEscaper_genType_doReturnFn
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

func (r *MoqHTMLEscaper_genType_fnRecorder) AndDo(fn MoqHTMLEscaper_genType_doFn) *MoqHTMLEscaper_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqHTMLEscaper_genType_fnRecorder) DoReturnResults(fn MoqHTMLEscaper_genType_doReturnFn) *MoqHTMLEscaper_genType_fnRecorder {
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
		DoFn       MoqHTMLEscaper_genType_doFn
		DoReturnFn MoqHTMLEscaper_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqHTMLEscaper_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqHTMLEscaper_genType_resultsByParams
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
		results = &MoqHTMLEscaper_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqHTMLEscaper_genType_paramsKey]*MoqHTMLEscaper_genType_results{},
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
		r.Results = &MoqHTMLEscaper_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqHTMLEscaper_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqHTMLEscaper_genType_fnRecorder {
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
				DoFn       MoqHTMLEscaper_genType_doFn
				DoReturnFn MoqHTMLEscaper_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqHTMLEscaper_genType) PrettyParams(params MoqHTMLEscaper_genType_params) string {
	return fmt.Sprintf("HTMLEscaper_genType(%#v)", params.Args)
}

func (m *MoqHTMLEscaper_genType) ParamsKey(params MoqHTMLEscaper_genType_params, anyParams uint64) MoqHTMLEscaper_genType_paramsKey {
	m.Scene.T.Helper()
	var argsUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Args == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The args parameter can't be indexed by value")
		}
		argsUsedHash = hash.DeepHash(params.Args)
	}
	return MoqHTMLEscaper_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Args hash.Hash }{
			Args: argsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqHTMLEscaper_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqHTMLEscaper_genType) AssertExpectationsMet() {
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
