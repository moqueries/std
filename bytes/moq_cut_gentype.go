// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package bytes

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Cut_genType is the fabricated implementation type of this mock (emitted when
// mocking functions directly and not from a function type)
type Cut_genType func(s, sep []byte) (before, after []byte, found bool)

// MoqCut_genType holds the state of a moq of the Cut_genType type
type MoqCut_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCut_genType_mock

	ResultsByParams []MoqCut_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			S   moq.ParamIndexing
			Sep moq.ParamIndexing
		}
	}
}

// MoqCut_genType_mock isolates the mock interface of the Cut_genType type
type MoqCut_genType_mock struct {
	Moq *MoqCut_genType
}

// MoqCut_genType_params holds the params of the Cut_genType type
type MoqCut_genType_params struct{ S, Sep []byte }

// MoqCut_genType_paramsKey holds the map key params of the Cut_genType type
type MoqCut_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ S, Sep hash.Hash }
}

// MoqCut_genType_resultsByParams contains the results for a given set of
// parameters for the Cut_genType type
type MoqCut_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCut_genType_paramsKey]*MoqCut_genType_results
}

// MoqCut_genType_doFn defines the type of function needed when calling AndDo
// for the Cut_genType type
type MoqCut_genType_doFn func(s, sep []byte)

// MoqCut_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Cut_genType type
type MoqCut_genType_doReturnFn func(s, sep []byte) (before, after []byte, found bool)

// MoqCut_genType_results holds the results of the Cut_genType type
type MoqCut_genType_results struct {
	Params  MoqCut_genType_params
	Results []struct {
		Values *struct {
			Before, After []byte
			Found         bool
		}
		Sequence   uint32
		DoFn       MoqCut_genType_doFn
		DoReturnFn MoqCut_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCut_genType_fnRecorder routes recorded function calls to the
// MoqCut_genType moq
type MoqCut_genType_fnRecorder struct {
	Params    MoqCut_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCut_genType_results
	Moq       *MoqCut_genType
}

// MoqCut_genType_anyParams isolates the any params functions of the
// Cut_genType type
type MoqCut_genType_anyParams struct {
	Recorder *MoqCut_genType_fnRecorder
}

// NewMoqCut_genType creates a new moq of the Cut_genType type
func NewMoqCut_genType(scene *moq.Scene, config *moq.Config) *MoqCut_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCut_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCut_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				S   moq.ParamIndexing
				Sep moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			S   moq.ParamIndexing
			Sep moq.ParamIndexing
		}{
			S:   moq.ParamIndexByHash,
			Sep: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Cut_genType type
func (m *MoqCut_genType) Mock() Cut_genType {
	return func(s, sep []byte) (_, _ []byte, _ bool) {
		m.Scene.T.Helper()
		moq := &MoqCut_genType_mock{Moq: m}
		return moq.Fn(s, sep)
	}
}

func (m *MoqCut_genType_mock) Fn(s, sep []byte) (before, after []byte, found bool) {
	m.Moq.Scene.T.Helper()
	params := MoqCut_genType_params{
		S:   s,
		Sep: sep,
	}
	var results *MoqCut_genType_results
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
		result.DoFn(s, sep)
	}

	if result.Values != nil {
		before = result.Values.Before
		after = result.Values.After
		found = result.Values.Found
	}
	if result.DoReturnFn != nil {
		before, after, found = result.DoReturnFn(s, sep)
	}
	return
}

func (m *MoqCut_genType) OnCall(s, sep []byte) *MoqCut_genType_fnRecorder {
	return &MoqCut_genType_fnRecorder{
		Params: MoqCut_genType_params{
			S:   s,
			Sep: sep,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqCut_genType_fnRecorder) Any() *MoqCut_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqCut_genType_anyParams{Recorder: r}
}

func (a *MoqCut_genType_anyParams) S() *MoqCut_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqCut_genType_anyParams) Sep() *MoqCut_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqCut_genType_fnRecorder) Seq() *MoqCut_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCut_genType_fnRecorder) NoSeq() *MoqCut_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCut_genType_fnRecorder) ReturnResults(before, after []byte, found bool) *MoqCut_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Before, After []byte
			Found         bool
		}
		Sequence   uint32
		DoFn       MoqCut_genType_doFn
		DoReturnFn MoqCut_genType_doReturnFn
	}{
		Values: &struct {
			Before, After []byte
			Found         bool
		}{
			Before: before,
			After:  after,
			Found:  found,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCut_genType_fnRecorder) AndDo(fn MoqCut_genType_doFn) *MoqCut_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCut_genType_fnRecorder) DoReturnResults(fn MoqCut_genType_doReturnFn) *MoqCut_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Before, After []byte
			Found         bool
		}
		Sequence   uint32
		DoFn       MoqCut_genType_doFn
		DoReturnFn MoqCut_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCut_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCut_genType_resultsByParams
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
		results = &MoqCut_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCut_genType_paramsKey]*MoqCut_genType_results{},
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
		r.Results = &MoqCut_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCut_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCut_genType_fnRecorder {
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
					Before, After []byte
					Found         bool
				}
				Sequence   uint32
				DoFn       MoqCut_genType_doFn
				DoReturnFn MoqCut_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCut_genType) PrettyParams(params MoqCut_genType_params) string {
	return fmt.Sprintf("Cut_genType(%#v, %#v)", params.S, params.Sep)
}

func (m *MoqCut_genType) ParamsKey(params MoqCut_genType_params, anyParams uint64) MoqCut_genType_paramsKey {
	m.Scene.T.Helper()
	var sUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.S == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The s parameter can't be indexed by value")
		}
		sUsedHash = hash.DeepHash(params.S)
	}
	var sepUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Sep == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The sep parameter can't be indexed by value")
		}
		sepUsedHash = hash.DeepHash(params.Sep)
	}
	return MoqCut_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ S, Sep hash.Hash }{
			S:   sUsedHash,
			Sep: sepUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCut_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCut_genType) AssertExpectationsMet() {
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
