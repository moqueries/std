// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package doc

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Synopsis_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Synopsis_genType func(s string) string

// MoqSynopsis_genType holds the state of a moq of the Synopsis_genType type
type MoqSynopsis_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqSynopsis_genType_mock

	ResultsByParams []MoqSynopsis_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			S moq.ParamIndexing
		}
	}
}

// MoqSynopsis_genType_mock isolates the mock interface of the Synopsis_genType
// type
type MoqSynopsis_genType_mock struct {
	Moq *MoqSynopsis_genType
}

// MoqSynopsis_genType_params holds the params of the Synopsis_genType type
type MoqSynopsis_genType_params struct{ S string }

// MoqSynopsis_genType_paramsKey holds the map key params of the
// Synopsis_genType type
type MoqSynopsis_genType_paramsKey struct {
	Params struct{ S string }
	Hashes struct{ S hash.Hash }
}

// MoqSynopsis_genType_resultsByParams contains the results for a given set of
// parameters for the Synopsis_genType type
type MoqSynopsis_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqSynopsis_genType_paramsKey]*MoqSynopsis_genType_results
}

// MoqSynopsis_genType_doFn defines the type of function needed when calling
// AndDo for the Synopsis_genType type
type MoqSynopsis_genType_doFn func(s string)

// MoqSynopsis_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Synopsis_genType type
type MoqSynopsis_genType_doReturnFn func(s string) string

// MoqSynopsis_genType_results holds the results of the Synopsis_genType type
type MoqSynopsis_genType_results struct {
	Params  MoqSynopsis_genType_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqSynopsis_genType_doFn
		DoReturnFn MoqSynopsis_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqSynopsis_genType_fnRecorder routes recorded function calls to the
// MoqSynopsis_genType moq
type MoqSynopsis_genType_fnRecorder struct {
	Params    MoqSynopsis_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqSynopsis_genType_results
	Moq       *MoqSynopsis_genType
}

// MoqSynopsis_genType_anyParams isolates the any params functions of the
// Synopsis_genType type
type MoqSynopsis_genType_anyParams struct {
	Recorder *MoqSynopsis_genType_fnRecorder
}

// NewMoqSynopsis_genType creates a new moq of the Synopsis_genType type
func NewMoqSynopsis_genType(scene *moq.Scene, config *moq.Config) *MoqSynopsis_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqSynopsis_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqSynopsis_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				S moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			S moq.ParamIndexing
		}{
			S: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Synopsis_genType type
func (m *MoqSynopsis_genType) Mock() Synopsis_genType {
	return func(s string) string { m.Scene.T.Helper(); moq := &MoqSynopsis_genType_mock{Moq: m}; return moq.Fn(s) }
}

func (m *MoqSynopsis_genType_mock) Fn(s string) (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqSynopsis_genType_params{
		S: s,
	}
	var results *MoqSynopsis_genType_results
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
		result.DoFn(s)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(s)
	}
	return
}

func (m *MoqSynopsis_genType) OnCall(s string) *MoqSynopsis_genType_fnRecorder {
	return &MoqSynopsis_genType_fnRecorder{
		Params: MoqSynopsis_genType_params{
			S: s,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqSynopsis_genType_fnRecorder) Any() *MoqSynopsis_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqSynopsis_genType_anyParams{Recorder: r}
}

func (a *MoqSynopsis_genType_anyParams) S() *MoqSynopsis_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqSynopsis_genType_fnRecorder) Seq() *MoqSynopsis_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqSynopsis_genType_fnRecorder) NoSeq() *MoqSynopsis_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqSynopsis_genType_fnRecorder) ReturnResults(result1 string) *MoqSynopsis_genType_fnRecorder {
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
		DoFn       MoqSynopsis_genType_doFn
		DoReturnFn MoqSynopsis_genType_doReturnFn
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

func (r *MoqSynopsis_genType_fnRecorder) AndDo(fn MoqSynopsis_genType_doFn) *MoqSynopsis_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqSynopsis_genType_fnRecorder) DoReturnResults(fn MoqSynopsis_genType_doReturnFn) *MoqSynopsis_genType_fnRecorder {
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
		DoFn       MoqSynopsis_genType_doFn
		DoReturnFn MoqSynopsis_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqSynopsis_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqSynopsis_genType_resultsByParams
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
		results = &MoqSynopsis_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqSynopsis_genType_paramsKey]*MoqSynopsis_genType_results{},
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
		r.Results = &MoqSynopsis_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqSynopsis_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqSynopsis_genType_fnRecorder {
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
				DoFn       MoqSynopsis_genType_doFn
				DoReturnFn MoqSynopsis_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqSynopsis_genType) PrettyParams(params MoqSynopsis_genType_params) string {
	return fmt.Sprintf("Synopsis_genType(%#v)", params.S)
}

func (m *MoqSynopsis_genType) ParamsKey(params MoqSynopsis_genType_params, anyParams uint64) MoqSynopsis_genType_paramsKey {
	m.Scene.T.Helper()
	var sUsed string
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			sUsed = params.S
		} else {
			sUsedHash = hash.DeepHash(params.S)
		}
	}
	return MoqSynopsis_genType_paramsKey{
		Params: struct{ S string }{
			S: sUsed,
		},
		Hashes: struct{ S hash.Hash }{
			S: sUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqSynopsis_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqSynopsis_genType) AssertExpectationsMet() {
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
