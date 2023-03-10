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

// IsEmptyTree_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type IsEmptyTree_genType func(n parse.Node) bool

// MoqIsEmptyTree_genType holds the state of a moq of the IsEmptyTree_genType
// type
type MoqIsEmptyTree_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqIsEmptyTree_genType_mock

	ResultsByParams []MoqIsEmptyTree_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			N moq.ParamIndexing
		}
	}
}

// MoqIsEmptyTree_genType_mock isolates the mock interface of the
// IsEmptyTree_genType type
type MoqIsEmptyTree_genType_mock struct {
	Moq *MoqIsEmptyTree_genType
}

// MoqIsEmptyTree_genType_params holds the params of the IsEmptyTree_genType
// type
type MoqIsEmptyTree_genType_params struct{ N parse.Node }

// MoqIsEmptyTree_genType_paramsKey holds the map key params of the
// IsEmptyTree_genType type
type MoqIsEmptyTree_genType_paramsKey struct {
	Params struct{ N parse.Node }
	Hashes struct{ N hash.Hash }
}

// MoqIsEmptyTree_genType_resultsByParams contains the results for a given set
// of parameters for the IsEmptyTree_genType type
type MoqIsEmptyTree_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqIsEmptyTree_genType_paramsKey]*MoqIsEmptyTree_genType_results
}

// MoqIsEmptyTree_genType_doFn defines the type of function needed when calling
// AndDo for the IsEmptyTree_genType type
type MoqIsEmptyTree_genType_doFn func(n parse.Node)

// MoqIsEmptyTree_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the IsEmptyTree_genType type
type MoqIsEmptyTree_genType_doReturnFn func(n parse.Node) bool

// MoqIsEmptyTree_genType_results holds the results of the IsEmptyTree_genType
// type
type MoqIsEmptyTree_genType_results struct {
	Params  MoqIsEmptyTree_genType_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIsEmptyTree_genType_doFn
		DoReturnFn MoqIsEmptyTree_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqIsEmptyTree_genType_fnRecorder routes recorded function calls to the
// MoqIsEmptyTree_genType moq
type MoqIsEmptyTree_genType_fnRecorder struct {
	Params    MoqIsEmptyTree_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqIsEmptyTree_genType_results
	Moq       *MoqIsEmptyTree_genType
}

// MoqIsEmptyTree_genType_anyParams isolates the any params functions of the
// IsEmptyTree_genType type
type MoqIsEmptyTree_genType_anyParams struct {
	Recorder *MoqIsEmptyTree_genType_fnRecorder
}

// NewMoqIsEmptyTree_genType creates a new moq of the IsEmptyTree_genType type
func NewMoqIsEmptyTree_genType(scene *moq.Scene, config *moq.Config) *MoqIsEmptyTree_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqIsEmptyTree_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqIsEmptyTree_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				N moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			N moq.ParamIndexing
		}{
			N: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the IsEmptyTree_genType type
func (m *MoqIsEmptyTree_genType) Mock() IsEmptyTree_genType {
	return func(n parse.Node) bool {
		m.Scene.T.Helper()
		moq := &MoqIsEmptyTree_genType_mock{Moq: m}
		return moq.Fn(n)
	}
}

func (m *MoqIsEmptyTree_genType_mock) Fn(n parse.Node) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqIsEmptyTree_genType_params{
		N: n,
	}
	var results *MoqIsEmptyTree_genType_results
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
		result.DoFn(n)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(n)
	}
	return
}

func (m *MoqIsEmptyTree_genType) OnCall(n parse.Node) *MoqIsEmptyTree_genType_fnRecorder {
	return &MoqIsEmptyTree_genType_fnRecorder{
		Params: MoqIsEmptyTree_genType_params{
			N: n,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqIsEmptyTree_genType_fnRecorder) Any() *MoqIsEmptyTree_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqIsEmptyTree_genType_anyParams{Recorder: r}
}

func (a *MoqIsEmptyTree_genType_anyParams) N() *MoqIsEmptyTree_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqIsEmptyTree_genType_fnRecorder) Seq() *MoqIsEmptyTree_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqIsEmptyTree_genType_fnRecorder) NoSeq() *MoqIsEmptyTree_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqIsEmptyTree_genType_fnRecorder) ReturnResults(result1 bool) *MoqIsEmptyTree_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIsEmptyTree_genType_doFn
		DoReturnFn MoqIsEmptyTree_genType_doReturnFn
	}{
		Values: &struct {
			Result1 bool
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqIsEmptyTree_genType_fnRecorder) AndDo(fn MoqIsEmptyTree_genType_doFn) *MoqIsEmptyTree_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqIsEmptyTree_genType_fnRecorder) DoReturnResults(fn MoqIsEmptyTree_genType_doReturnFn) *MoqIsEmptyTree_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqIsEmptyTree_genType_doFn
		DoReturnFn MoqIsEmptyTree_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqIsEmptyTree_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqIsEmptyTree_genType_resultsByParams
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
		results = &MoqIsEmptyTree_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqIsEmptyTree_genType_paramsKey]*MoqIsEmptyTree_genType_results{},
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
		r.Results = &MoqIsEmptyTree_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqIsEmptyTree_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqIsEmptyTree_genType_fnRecorder {
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
					Result1 bool
				}
				Sequence   uint32
				DoFn       MoqIsEmptyTree_genType_doFn
				DoReturnFn MoqIsEmptyTree_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqIsEmptyTree_genType) PrettyParams(params MoqIsEmptyTree_genType_params) string {
	return fmt.Sprintf("IsEmptyTree_genType(%#v)", params.N)
}

func (m *MoqIsEmptyTree_genType) ParamsKey(params MoqIsEmptyTree_genType_params, anyParams uint64) MoqIsEmptyTree_genType_paramsKey {
	m.Scene.T.Helper()
	var nUsed parse.Node
	var nUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.N == moq.ParamIndexByValue {
			nUsed = params.N
		} else {
			nUsedHash = hash.DeepHash(params.N)
		}
	}
	return MoqIsEmptyTree_genType_paramsKey{
		Params: struct{ N parse.Node }{
			N: nUsed,
		},
		Hashes: struct{ N hash.Hash }{
			N: nUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqIsEmptyTree_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqIsEmptyTree_genType) AssertExpectationsMet() {
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
