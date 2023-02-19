// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package parse

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"text/template/parse"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// New_genType is the fabricated implementation type of this mock (emitted when
// mocking functions directly and not from a function type)
type New_genType func(name string, funcs ...map[string]interface{}) *parse.Tree

// MoqNew_genType holds the state of a moq of the New_genType type
type MoqNew_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNew_genType_mock

	ResultsByParams []MoqNew_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Name  moq.ParamIndexing
			Funcs moq.ParamIndexing
		}
	}
}

// MoqNew_genType_mock isolates the mock interface of the New_genType type
type MoqNew_genType_mock struct {
	Moq *MoqNew_genType
}

// MoqNew_genType_params holds the params of the New_genType type
type MoqNew_genType_params struct {
	Name  string
	Funcs []map[string]interface{}
}

// MoqNew_genType_paramsKey holds the map key params of the New_genType type
type MoqNew_genType_paramsKey struct {
	Params struct{ Name string }
	Hashes struct {
		Name  hash.Hash
		Funcs hash.Hash
	}
}

// MoqNew_genType_resultsByParams contains the results for a given set of
// parameters for the New_genType type
type MoqNew_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNew_genType_paramsKey]*MoqNew_genType_results
}

// MoqNew_genType_doFn defines the type of function needed when calling AndDo
// for the New_genType type
type MoqNew_genType_doFn func(name string, funcs ...map[string]interface{})

// MoqNew_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the New_genType type
type MoqNew_genType_doReturnFn func(name string, funcs ...map[string]interface{}) *parse.Tree

// MoqNew_genType_results holds the results of the New_genType type
type MoqNew_genType_results struct {
	Params  MoqNew_genType_params
	Results []struct {
		Values *struct {
			Result1 *parse.Tree
		}
		Sequence   uint32
		DoFn       MoqNew_genType_doFn
		DoReturnFn MoqNew_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNew_genType_fnRecorder routes recorded function calls to the
// MoqNew_genType moq
type MoqNew_genType_fnRecorder struct {
	Params    MoqNew_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNew_genType_results
	Moq       *MoqNew_genType
}

// MoqNew_genType_anyParams isolates the any params functions of the
// New_genType type
type MoqNew_genType_anyParams struct {
	Recorder *MoqNew_genType_fnRecorder
}

// NewMoqNew_genType creates a new moq of the New_genType type
func NewMoqNew_genType(scene *moq.Scene, config *moq.Config) *MoqNew_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNew_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNew_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Name  moq.ParamIndexing
				Funcs moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Name  moq.ParamIndexing
			Funcs moq.ParamIndexing
		}{
			Name:  moq.ParamIndexByValue,
			Funcs: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the New_genType type
func (m *MoqNew_genType) Mock() New_genType {
	return func(name string, funcs ...map[string]interface{}) *parse.Tree {
		m.Scene.T.Helper()
		moq := &MoqNew_genType_mock{Moq: m}
		return moq.Fn(name, funcs...)
	}
}

func (m *MoqNew_genType_mock) Fn(name string, funcs ...map[string]interface{}) (result1 *parse.Tree) {
	m.Moq.Scene.T.Helper()
	params := MoqNew_genType_params{
		Name:  name,
		Funcs: funcs,
	}
	var results *MoqNew_genType_results
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
		result.DoFn(name, funcs...)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(name, funcs...)
	}
	return
}

func (m *MoqNew_genType) OnCall(name string, funcs ...map[string]interface{}) *MoqNew_genType_fnRecorder {
	return &MoqNew_genType_fnRecorder{
		Params: MoqNew_genType_params{
			Name:  name,
			Funcs: funcs,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNew_genType_fnRecorder) Any() *MoqNew_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNew_genType_anyParams{Recorder: r}
}

func (a *MoqNew_genType_anyParams) Name() *MoqNew_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNew_genType_anyParams) Funcs() *MoqNew_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqNew_genType_fnRecorder) Seq() *MoqNew_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNew_genType_fnRecorder) NoSeq() *MoqNew_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNew_genType_fnRecorder) ReturnResults(result1 *parse.Tree) *MoqNew_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *parse.Tree
		}
		Sequence   uint32
		DoFn       MoqNew_genType_doFn
		DoReturnFn MoqNew_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *parse.Tree
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNew_genType_fnRecorder) AndDo(fn MoqNew_genType_doFn) *MoqNew_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNew_genType_fnRecorder) DoReturnResults(fn MoqNew_genType_doReturnFn) *MoqNew_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *parse.Tree
		}
		Sequence   uint32
		DoFn       MoqNew_genType_doFn
		DoReturnFn MoqNew_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNew_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNew_genType_resultsByParams
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
		results = &MoqNew_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNew_genType_paramsKey]*MoqNew_genType_results{},
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
		r.Results = &MoqNew_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNew_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNew_genType_fnRecorder {
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
					Result1 *parse.Tree
				}
				Sequence   uint32
				DoFn       MoqNew_genType_doFn
				DoReturnFn MoqNew_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNew_genType) PrettyParams(params MoqNew_genType_params) string {
	return fmt.Sprintf("New_genType(%#v, %#v)", params.Name, params.Funcs)
}

func (m *MoqNew_genType) ParamsKey(params MoqNew_genType_params, anyParams uint64) MoqNew_genType_paramsKey {
	m.Scene.T.Helper()
	var nameUsed string
	var nameUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Name == moq.ParamIndexByValue {
			nameUsed = params.Name
		} else {
			nameUsedHash = hash.DeepHash(params.Name)
		}
	}
	var funcsUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Funcs == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The funcs parameter can't be indexed by value")
		}
		funcsUsedHash = hash.DeepHash(params.Funcs)
	}
	return MoqNew_genType_paramsKey{
		Params: struct{ Name string }{
			Name: nameUsed,
		},
		Hashes: struct {
			Name  hash.Hash
			Funcs hash.Hash
		}{
			Name:  nameUsedHash,
			Funcs: funcsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNew_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNew_genType) AssertExpectationsMet() {
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
