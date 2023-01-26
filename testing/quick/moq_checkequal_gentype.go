// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package quick

import (
	"fmt"
	"math/bits"
	"sync/atomic"
	"testing/quick"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// CheckEqual_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type CheckEqual_genType func(f, g interface{}, config *quick.Config) error

// MoqCheckEqual_genType holds the state of a moq of the CheckEqual_genType
// type
type MoqCheckEqual_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCheckEqual_genType_mock

	ResultsByParams []MoqCheckEqual_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			F      moq.ParamIndexing
			G      moq.ParamIndexing
			Config moq.ParamIndexing
		}
	}
}

// MoqCheckEqual_genType_mock isolates the mock interface of the
// CheckEqual_genType type
type MoqCheckEqual_genType_mock struct {
	Moq *MoqCheckEqual_genType
}

// MoqCheckEqual_genType_params holds the params of the CheckEqual_genType type
type MoqCheckEqual_genType_params struct {
	F, G   interface{}
	Config *quick.Config
}

// MoqCheckEqual_genType_paramsKey holds the map key params of the
// CheckEqual_genType type
type MoqCheckEqual_genType_paramsKey struct {
	Params struct {
		F, G   interface{}
		Config *quick.Config
	}
	Hashes struct {
		F, G   hash.Hash
		Config hash.Hash
	}
}

// MoqCheckEqual_genType_resultsByParams contains the results for a given set
// of parameters for the CheckEqual_genType type
type MoqCheckEqual_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCheckEqual_genType_paramsKey]*MoqCheckEqual_genType_results
}

// MoqCheckEqual_genType_doFn defines the type of function needed when calling
// AndDo for the CheckEqual_genType type
type MoqCheckEqual_genType_doFn func(f, g interface{}, config *quick.Config)

// MoqCheckEqual_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the CheckEqual_genType type
type MoqCheckEqual_genType_doReturnFn func(f, g interface{}, config *quick.Config) error

// MoqCheckEqual_genType_results holds the results of the CheckEqual_genType
// type
type MoqCheckEqual_genType_results struct {
	Params  MoqCheckEqual_genType_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqCheckEqual_genType_doFn
		DoReturnFn MoqCheckEqual_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCheckEqual_genType_fnRecorder routes recorded function calls to the
// MoqCheckEqual_genType moq
type MoqCheckEqual_genType_fnRecorder struct {
	Params    MoqCheckEqual_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCheckEqual_genType_results
	Moq       *MoqCheckEqual_genType
}

// MoqCheckEqual_genType_anyParams isolates the any params functions of the
// CheckEqual_genType type
type MoqCheckEqual_genType_anyParams struct {
	Recorder *MoqCheckEqual_genType_fnRecorder
}

// NewMoqCheckEqual_genType creates a new moq of the CheckEqual_genType type
func NewMoqCheckEqual_genType(scene *moq.Scene, config *moq.Config) *MoqCheckEqual_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCheckEqual_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCheckEqual_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				F      moq.ParamIndexing
				G      moq.ParamIndexing
				Config moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			F      moq.ParamIndexing
			G      moq.ParamIndexing
			Config moq.ParamIndexing
		}{
			F:      moq.ParamIndexByHash,
			G:      moq.ParamIndexByHash,
			Config: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the CheckEqual_genType type
func (m *MoqCheckEqual_genType) Mock() CheckEqual_genType {
	return func(f, g interface{}, config *quick.Config) error {
		m.Scene.T.Helper()
		moq := &MoqCheckEqual_genType_mock{Moq: m}
		return moq.Fn(f, g, config)
	}
}

func (m *MoqCheckEqual_genType_mock) Fn(f, g interface{}, config *quick.Config) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqCheckEqual_genType_params{
		F:      f,
		G:      g,
		Config: config,
	}
	var results *MoqCheckEqual_genType_results
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
		result.DoFn(f, g, config)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(f, g, config)
	}
	return
}

func (m *MoqCheckEqual_genType) OnCall(f, g interface{}, config *quick.Config) *MoqCheckEqual_genType_fnRecorder {
	return &MoqCheckEqual_genType_fnRecorder{
		Params: MoqCheckEqual_genType_params{
			F:      f,
			G:      g,
			Config: config,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqCheckEqual_genType_fnRecorder) Any() *MoqCheckEqual_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqCheckEqual_genType_anyParams{Recorder: r}
}

func (a *MoqCheckEqual_genType_anyParams) F() *MoqCheckEqual_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqCheckEqual_genType_anyParams) G() *MoqCheckEqual_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqCheckEqual_genType_anyParams) Config() *MoqCheckEqual_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqCheckEqual_genType_fnRecorder) Seq() *MoqCheckEqual_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCheckEqual_genType_fnRecorder) NoSeq() *MoqCheckEqual_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCheckEqual_genType_fnRecorder) ReturnResults(result1 error) *MoqCheckEqual_genType_fnRecorder {
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
		DoFn       MoqCheckEqual_genType_doFn
		DoReturnFn MoqCheckEqual_genType_doReturnFn
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

func (r *MoqCheckEqual_genType_fnRecorder) AndDo(fn MoqCheckEqual_genType_doFn) *MoqCheckEqual_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCheckEqual_genType_fnRecorder) DoReturnResults(fn MoqCheckEqual_genType_doReturnFn) *MoqCheckEqual_genType_fnRecorder {
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
		DoFn       MoqCheckEqual_genType_doFn
		DoReturnFn MoqCheckEqual_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCheckEqual_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCheckEqual_genType_resultsByParams
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
		results = &MoqCheckEqual_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCheckEqual_genType_paramsKey]*MoqCheckEqual_genType_results{},
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
		r.Results = &MoqCheckEqual_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCheckEqual_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCheckEqual_genType_fnRecorder {
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
				DoFn       MoqCheckEqual_genType_doFn
				DoReturnFn MoqCheckEqual_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCheckEqual_genType) PrettyParams(params MoqCheckEqual_genType_params) string {
	return fmt.Sprintf("CheckEqual_genType(%#v, %#v, %#v)", params.F, params.G, params.Config)
}

func (m *MoqCheckEqual_genType) ParamsKey(params MoqCheckEqual_genType_params, anyParams uint64) MoqCheckEqual_genType_paramsKey {
	m.Scene.T.Helper()
	var fUsed interface{}
	var fUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.F == moq.ParamIndexByValue {
			fUsed = params.F
		} else {
			fUsedHash = hash.DeepHash(params.F)
		}
	}
	var gUsed interface{}
	var gUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.G == moq.ParamIndexByValue {
			gUsed = params.G
		} else {
			gUsedHash = hash.DeepHash(params.G)
		}
	}
	var configUsed *quick.Config
	var configUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Config == moq.ParamIndexByValue {
			configUsed = params.Config
		} else {
			configUsedHash = hash.DeepHash(params.Config)
		}
	}
	return MoqCheckEqual_genType_paramsKey{
		Params: struct {
			F, G   interface{}
			Config *quick.Config
		}{
			F:      fUsed,
			G:      gUsed,
			Config: configUsed,
		},
		Hashes: struct {
			F, G   hash.Hash
			Config hash.Hash
		}{
			F:      fUsedHash,
			G:      gUsedHash,
			Config: configUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCheckEqual_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCheckEqual_genType) AssertExpectationsMet() {
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