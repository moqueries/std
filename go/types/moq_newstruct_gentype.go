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

// NewStruct_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type NewStruct_genType func(fields []*types.Var, tags []string) *types.Struct

// MoqNewStruct_genType holds the state of a moq of the NewStruct_genType type
type MoqNewStruct_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNewStruct_genType_mock

	ResultsByParams []MoqNewStruct_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fields moq.ParamIndexing
			Tags   moq.ParamIndexing
		}
	}
}

// MoqNewStruct_genType_mock isolates the mock interface of the
// NewStruct_genType type
type MoqNewStruct_genType_mock struct {
	Moq *MoqNewStruct_genType
}

// MoqNewStruct_genType_params holds the params of the NewStruct_genType type
type MoqNewStruct_genType_params struct {
	Fields []*types.Var
	Tags   []string
}

// MoqNewStruct_genType_paramsKey holds the map key params of the
// NewStruct_genType type
type MoqNewStruct_genType_paramsKey struct {
	Params struct{}
	Hashes struct {
		Fields hash.Hash
		Tags   hash.Hash
	}
}

// MoqNewStruct_genType_resultsByParams contains the results for a given set of
// parameters for the NewStruct_genType type
type MoqNewStruct_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNewStruct_genType_paramsKey]*MoqNewStruct_genType_results
}

// MoqNewStruct_genType_doFn defines the type of function needed when calling
// AndDo for the NewStruct_genType type
type MoqNewStruct_genType_doFn func(fields []*types.Var, tags []string)

// MoqNewStruct_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the NewStruct_genType type
type MoqNewStruct_genType_doReturnFn func(fields []*types.Var, tags []string) *types.Struct

// MoqNewStruct_genType_results holds the results of the NewStruct_genType type
type MoqNewStruct_genType_results struct {
	Params  MoqNewStruct_genType_params
	Results []struct {
		Values *struct {
			Result1 *types.Struct
		}
		Sequence   uint32
		DoFn       MoqNewStruct_genType_doFn
		DoReturnFn MoqNewStruct_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNewStruct_genType_fnRecorder routes recorded function calls to the
// MoqNewStruct_genType moq
type MoqNewStruct_genType_fnRecorder struct {
	Params    MoqNewStruct_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNewStruct_genType_results
	Moq       *MoqNewStruct_genType
}

// MoqNewStruct_genType_anyParams isolates the any params functions of the
// NewStruct_genType type
type MoqNewStruct_genType_anyParams struct {
	Recorder *MoqNewStruct_genType_fnRecorder
}

// NewMoqNewStruct_genType creates a new moq of the NewStruct_genType type
func NewMoqNewStruct_genType(scene *moq.Scene, config *moq.Config) *MoqNewStruct_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNewStruct_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNewStruct_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fields moq.ParamIndexing
				Tags   moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Fields moq.ParamIndexing
			Tags   moq.ParamIndexing
		}{
			Fields: moq.ParamIndexByHash,
			Tags:   moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the NewStruct_genType type
func (m *MoqNewStruct_genType) Mock() NewStruct_genType {
	return func(fields []*types.Var, tags []string) *types.Struct {
		m.Scene.T.Helper()
		moq := &MoqNewStruct_genType_mock{Moq: m}
		return moq.Fn(fields, tags)
	}
}

func (m *MoqNewStruct_genType_mock) Fn(fields []*types.Var, tags []string) (result1 *types.Struct) {
	m.Moq.Scene.T.Helper()
	params := MoqNewStruct_genType_params{
		Fields: fields,
		Tags:   tags,
	}
	var results *MoqNewStruct_genType_results
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
		result.DoFn(fields, tags)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(fields, tags)
	}
	return
}

func (m *MoqNewStruct_genType) OnCall(fields []*types.Var, tags []string) *MoqNewStruct_genType_fnRecorder {
	return &MoqNewStruct_genType_fnRecorder{
		Params: MoqNewStruct_genType_params{
			Fields: fields,
			Tags:   tags,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqNewStruct_genType_fnRecorder) Any() *MoqNewStruct_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqNewStruct_genType_anyParams{Recorder: r}
}

func (a *MoqNewStruct_genType_anyParams) Fields() *MoqNewStruct_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqNewStruct_genType_anyParams) Tags() *MoqNewStruct_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqNewStruct_genType_fnRecorder) Seq() *MoqNewStruct_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNewStruct_genType_fnRecorder) NoSeq() *MoqNewStruct_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNewStruct_genType_fnRecorder) ReturnResults(result1 *types.Struct) *MoqNewStruct_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Struct
		}
		Sequence   uint32
		DoFn       MoqNewStruct_genType_doFn
		DoReturnFn MoqNewStruct_genType_doReturnFn
	}{
		Values: &struct {
			Result1 *types.Struct
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqNewStruct_genType_fnRecorder) AndDo(fn MoqNewStruct_genType_doFn) *MoqNewStruct_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNewStruct_genType_fnRecorder) DoReturnResults(fn MoqNewStruct_genType_doReturnFn) *MoqNewStruct_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Struct
		}
		Sequence   uint32
		DoFn       MoqNewStruct_genType_doFn
		DoReturnFn MoqNewStruct_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNewStruct_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNewStruct_genType_resultsByParams
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
		results = &MoqNewStruct_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNewStruct_genType_paramsKey]*MoqNewStruct_genType_results{},
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
		r.Results = &MoqNewStruct_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNewStruct_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNewStruct_genType_fnRecorder {
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
					Result1 *types.Struct
				}
				Sequence   uint32
				DoFn       MoqNewStruct_genType_doFn
				DoReturnFn MoqNewStruct_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNewStruct_genType) PrettyParams(params MoqNewStruct_genType_params) string {
	return fmt.Sprintf("NewStruct_genType(%#v, %#v)", params.Fields, params.Tags)
}

func (m *MoqNewStruct_genType) ParamsKey(params MoqNewStruct_genType_params, anyParams uint64) MoqNewStruct_genType_paramsKey {
	m.Scene.T.Helper()
	var fieldsUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Fields == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The fields parameter can't be indexed by value")
		}
		fieldsUsedHash = hash.DeepHash(params.Fields)
	}
	var tagsUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Tags == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The tags parameter can't be indexed by value")
		}
		tagsUsedHash = hash.DeepHash(params.Tags)
	}
	return MoqNewStruct_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct {
			Fields hash.Hash
			Tags   hash.Hash
		}{
			Fields: fieldsUsedHash,
			Tags:   tagsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNewStruct_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNewStruct_genType) AssertExpectationsMet() {
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