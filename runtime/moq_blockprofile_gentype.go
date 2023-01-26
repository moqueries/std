// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package runtime

import (
	"fmt"
	"math/bits"
	"runtime"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// BlockProfile_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type BlockProfile_genType func(p []runtime.BlockProfileRecord) (n int, ok bool)

// MoqBlockProfile_genType holds the state of a moq of the BlockProfile_genType
// type
type MoqBlockProfile_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqBlockProfile_genType_mock

	ResultsByParams []MoqBlockProfile_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			P moq.ParamIndexing
		}
	}
}

// MoqBlockProfile_genType_mock isolates the mock interface of the
// BlockProfile_genType type
type MoqBlockProfile_genType_mock struct {
	Moq *MoqBlockProfile_genType
}

// MoqBlockProfile_genType_params holds the params of the BlockProfile_genType
// type
type MoqBlockProfile_genType_params struct{ P []runtime.BlockProfileRecord }

// MoqBlockProfile_genType_paramsKey holds the map key params of the
// BlockProfile_genType type
type MoqBlockProfile_genType_paramsKey struct {
	Params struct{}
	Hashes struct{ P hash.Hash }
}

// MoqBlockProfile_genType_resultsByParams contains the results for a given set
// of parameters for the BlockProfile_genType type
type MoqBlockProfile_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqBlockProfile_genType_paramsKey]*MoqBlockProfile_genType_results
}

// MoqBlockProfile_genType_doFn defines the type of function needed when
// calling AndDo for the BlockProfile_genType type
type MoqBlockProfile_genType_doFn func(p []runtime.BlockProfileRecord)

// MoqBlockProfile_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the BlockProfile_genType type
type MoqBlockProfile_genType_doReturnFn func(p []runtime.BlockProfileRecord) (n int, ok bool)

// MoqBlockProfile_genType_results holds the results of the
// BlockProfile_genType type
type MoqBlockProfile_genType_results struct {
	Params  MoqBlockProfile_genType_params
	Results []struct {
		Values *struct {
			N  int
			Ok bool
		}
		Sequence   uint32
		DoFn       MoqBlockProfile_genType_doFn
		DoReturnFn MoqBlockProfile_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqBlockProfile_genType_fnRecorder routes recorded function calls to the
// MoqBlockProfile_genType moq
type MoqBlockProfile_genType_fnRecorder struct {
	Params    MoqBlockProfile_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqBlockProfile_genType_results
	Moq       *MoqBlockProfile_genType
}

// MoqBlockProfile_genType_anyParams isolates the any params functions of the
// BlockProfile_genType type
type MoqBlockProfile_genType_anyParams struct {
	Recorder *MoqBlockProfile_genType_fnRecorder
}

// NewMoqBlockProfile_genType creates a new moq of the BlockProfile_genType
// type
func NewMoqBlockProfile_genType(scene *moq.Scene, config *moq.Config) *MoqBlockProfile_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqBlockProfile_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqBlockProfile_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				P moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			P moq.ParamIndexing
		}{
			P: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the BlockProfile_genType type
func (m *MoqBlockProfile_genType) Mock() BlockProfile_genType {
	return func(p []runtime.BlockProfileRecord) (_ int, _ bool) {
		m.Scene.T.Helper()
		moq := &MoqBlockProfile_genType_mock{Moq: m}
		return moq.Fn(p)
	}
}

func (m *MoqBlockProfile_genType_mock) Fn(p []runtime.BlockProfileRecord) (n int, ok bool) {
	m.Moq.Scene.T.Helper()
	params := MoqBlockProfile_genType_params{
		P: p,
	}
	var results *MoqBlockProfile_genType_results
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
		result.DoFn(p)
	}

	if result.Values != nil {
		n = result.Values.N
		ok = result.Values.Ok
	}
	if result.DoReturnFn != nil {
		n, ok = result.DoReturnFn(p)
	}
	return
}

func (m *MoqBlockProfile_genType) OnCall(p []runtime.BlockProfileRecord) *MoqBlockProfile_genType_fnRecorder {
	return &MoqBlockProfile_genType_fnRecorder{
		Params: MoqBlockProfile_genType_params{
			P: p,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqBlockProfile_genType_fnRecorder) Any() *MoqBlockProfile_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqBlockProfile_genType_anyParams{Recorder: r}
}

func (a *MoqBlockProfile_genType_anyParams) P() *MoqBlockProfile_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqBlockProfile_genType_fnRecorder) Seq() *MoqBlockProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqBlockProfile_genType_fnRecorder) NoSeq() *MoqBlockProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqBlockProfile_genType_fnRecorder) ReturnResults(n int, ok bool) *MoqBlockProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N  int
			Ok bool
		}
		Sequence   uint32
		DoFn       MoqBlockProfile_genType_doFn
		DoReturnFn MoqBlockProfile_genType_doReturnFn
	}{
		Values: &struct {
			N  int
			Ok bool
		}{
			N:  n,
			Ok: ok,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqBlockProfile_genType_fnRecorder) AndDo(fn MoqBlockProfile_genType_doFn) *MoqBlockProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqBlockProfile_genType_fnRecorder) DoReturnResults(fn MoqBlockProfile_genType_doReturnFn) *MoqBlockProfile_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			N  int
			Ok bool
		}
		Sequence   uint32
		DoFn       MoqBlockProfile_genType_doFn
		DoReturnFn MoqBlockProfile_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqBlockProfile_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqBlockProfile_genType_resultsByParams
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
		results = &MoqBlockProfile_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqBlockProfile_genType_paramsKey]*MoqBlockProfile_genType_results{},
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
		r.Results = &MoqBlockProfile_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqBlockProfile_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqBlockProfile_genType_fnRecorder {
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
					N  int
					Ok bool
				}
				Sequence   uint32
				DoFn       MoqBlockProfile_genType_doFn
				DoReturnFn MoqBlockProfile_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqBlockProfile_genType) PrettyParams(params MoqBlockProfile_genType_params) string {
	return fmt.Sprintf("BlockProfile_genType(%#v)", params.P)
}

func (m *MoqBlockProfile_genType) ParamsKey(params MoqBlockProfile_genType_params, anyParams uint64) MoqBlockProfile_genType_paramsKey {
	m.Scene.T.Helper()
	var pUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.P == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The p parameter can't be indexed by value")
		}
		pUsedHash = hash.DeepHash(params.P)
	}
	return MoqBlockProfile_genType_paramsKey{
		Params: struct{}{},
		Hashes: struct{ P hash.Hash }{
			P: pUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqBlockProfile_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqBlockProfile_genType) AssertExpectationsMet() {
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
