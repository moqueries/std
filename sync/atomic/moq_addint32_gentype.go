// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package atomic

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// AddInt32_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type AddInt32_genType func(addr *int32, delta int32) (new int32)

// MoqAddInt32_genType holds the state of a moq of the AddInt32_genType type
type MoqAddInt32_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqAddInt32_genType_mock

	ResultsByParams []MoqAddInt32_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Addr  moq.ParamIndexing
			Delta moq.ParamIndexing
		}
	}
}

// MoqAddInt32_genType_mock isolates the mock interface of the AddInt32_genType
// type
type MoqAddInt32_genType_mock struct {
	Moq *MoqAddInt32_genType
}

// MoqAddInt32_genType_params holds the params of the AddInt32_genType type
type MoqAddInt32_genType_params struct {
	Addr  *int32
	Delta int32
}

// MoqAddInt32_genType_paramsKey holds the map key params of the
// AddInt32_genType type
type MoqAddInt32_genType_paramsKey struct {
	Params struct {
		Addr  *int32
		Delta int32
	}
	Hashes struct {
		Addr  hash.Hash
		Delta hash.Hash
	}
}

// MoqAddInt32_genType_resultsByParams contains the results for a given set of
// parameters for the AddInt32_genType type
type MoqAddInt32_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqAddInt32_genType_paramsKey]*MoqAddInt32_genType_results
}

// MoqAddInt32_genType_doFn defines the type of function needed when calling
// AndDo for the AddInt32_genType type
type MoqAddInt32_genType_doFn func(addr *int32, delta int32)

// MoqAddInt32_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the AddInt32_genType type
type MoqAddInt32_genType_doReturnFn func(addr *int32, delta int32) (new int32)

// MoqAddInt32_genType_results holds the results of the AddInt32_genType type
type MoqAddInt32_genType_results struct {
	Params  MoqAddInt32_genType_params
	Results []struct {
		Values     *struct{ New int32 }
		Sequence   uint32
		DoFn       MoqAddInt32_genType_doFn
		DoReturnFn MoqAddInt32_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqAddInt32_genType_fnRecorder routes recorded function calls to the
// MoqAddInt32_genType moq
type MoqAddInt32_genType_fnRecorder struct {
	Params    MoqAddInt32_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqAddInt32_genType_results
	Moq       *MoqAddInt32_genType
}

// MoqAddInt32_genType_anyParams isolates the any params functions of the
// AddInt32_genType type
type MoqAddInt32_genType_anyParams struct {
	Recorder *MoqAddInt32_genType_fnRecorder
}

// NewMoqAddInt32_genType creates a new moq of the AddInt32_genType type
func NewMoqAddInt32_genType(scene *moq.Scene, config *moq.Config) *MoqAddInt32_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqAddInt32_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqAddInt32_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Addr  moq.ParamIndexing
				Delta moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Addr  moq.ParamIndexing
			Delta moq.ParamIndexing
		}{
			Addr:  moq.ParamIndexByHash,
			Delta: moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the AddInt32_genType type
func (m *MoqAddInt32_genType) Mock() AddInt32_genType {
	return func(addr *int32, delta int32) (_ int32) {
		m.Scene.T.Helper()
		moq := &MoqAddInt32_genType_mock{Moq: m}
		return moq.Fn(addr, delta)
	}
}

func (m *MoqAddInt32_genType_mock) Fn(addr *int32, delta int32) (new int32) {
	m.Moq.Scene.T.Helper()
	params := MoqAddInt32_genType_params{
		Addr:  addr,
		Delta: delta,
	}
	var results *MoqAddInt32_genType_results
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
		result.DoFn(addr, delta)
	}

	if result.Values != nil {
		new = result.Values.New
	}
	if result.DoReturnFn != nil {
		new = result.DoReturnFn(addr, delta)
	}
	return
}

func (m *MoqAddInt32_genType) OnCall(addr *int32, delta int32) *MoqAddInt32_genType_fnRecorder {
	return &MoqAddInt32_genType_fnRecorder{
		Params: MoqAddInt32_genType_params{
			Addr:  addr,
			Delta: delta,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqAddInt32_genType_fnRecorder) Any() *MoqAddInt32_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqAddInt32_genType_anyParams{Recorder: r}
}

func (a *MoqAddInt32_genType_anyParams) Addr() *MoqAddInt32_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqAddInt32_genType_anyParams) Delta() *MoqAddInt32_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqAddInt32_genType_fnRecorder) Seq() *MoqAddInt32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqAddInt32_genType_fnRecorder) NoSeq() *MoqAddInt32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqAddInt32_genType_fnRecorder) ReturnResults(new int32) *MoqAddInt32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ New int32 }
		Sequence   uint32
		DoFn       MoqAddInt32_genType_doFn
		DoReturnFn MoqAddInt32_genType_doReturnFn
	}{
		Values: &struct{ New int32 }{
			New: new,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqAddInt32_genType_fnRecorder) AndDo(fn MoqAddInt32_genType_doFn) *MoqAddInt32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqAddInt32_genType_fnRecorder) DoReturnResults(fn MoqAddInt32_genType_doReturnFn) *MoqAddInt32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ New int32 }
		Sequence   uint32
		DoFn       MoqAddInt32_genType_doFn
		DoReturnFn MoqAddInt32_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqAddInt32_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqAddInt32_genType_resultsByParams
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
		results = &MoqAddInt32_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqAddInt32_genType_paramsKey]*MoqAddInt32_genType_results{},
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
		r.Results = &MoqAddInt32_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqAddInt32_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqAddInt32_genType_fnRecorder {
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
				Values     *struct{ New int32 }
				Sequence   uint32
				DoFn       MoqAddInt32_genType_doFn
				DoReturnFn MoqAddInt32_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqAddInt32_genType) PrettyParams(params MoqAddInt32_genType_params) string {
	return fmt.Sprintf("AddInt32_genType(%#v, %#v)", params.Addr, params.Delta)
}

func (m *MoqAddInt32_genType) ParamsKey(params MoqAddInt32_genType_params, anyParams uint64) MoqAddInt32_genType_paramsKey {
	m.Scene.T.Helper()
	var addrUsed *int32
	var addrUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Addr == moq.ParamIndexByValue {
			addrUsed = params.Addr
		} else {
			addrUsedHash = hash.DeepHash(params.Addr)
		}
	}
	var deltaUsed int32
	var deltaUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Delta == moq.ParamIndexByValue {
			deltaUsed = params.Delta
		} else {
			deltaUsedHash = hash.DeepHash(params.Delta)
		}
	}
	return MoqAddInt32_genType_paramsKey{
		Params: struct {
			Addr  *int32
			Delta int32
		}{
			Addr:  addrUsed,
			Delta: deltaUsed,
		},
		Hashes: struct {
			Addr  hash.Hash
			Delta hash.Hash
		}{
			Addr:  addrUsedHash,
			Delta: deltaUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqAddInt32_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqAddInt32_genType) AssertExpectationsMet() {
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
