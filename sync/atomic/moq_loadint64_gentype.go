// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package atomic

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// LoadInt64_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type LoadInt64_genType func(addr *int64) (val int64)

// MoqLoadInt64_genType holds the state of a moq of the LoadInt64_genType type
type MoqLoadInt64_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqLoadInt64_genType_mock

	ResultsByParams []MoqLoadInt64_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Addr moq.ParamIndexing
		}
	}
}

// MoqLoadInt64_genType_mock isolates the mock interface of the
// LoadInt64_genType type
type MoqLoadInt64_genType_mock struct {
	Moq *MoqLoadInt64_genType
}

// MoqLoadInt64_genType_params holds the params of the LoadInt64_genType type
type MoqLoadInt64_genType_params struct{ Addr *int64 }

// MoqLoadInt64_genType_paramsKey holds the map key params of the
// LoadInt64_genType type
type MoqLoadInt64_genType_paramsKey struct {
	Params struct{ Addr *int64 }
	Hashes struct{ Addr hash.Hash }
}

// MoqLoadInt64_genType_resultsByParams contains the results for a given set of
// parameters for the LoadInt64_genType type
type MoqLoadInt64_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqLoadInt64_genType_paramsKey]*MoqLoadInt64_genType_results
}

// MoqLoadInt64_genType_doFn defines the type of function needed when calling
// AndDo for the LoadInt64_genType type
type MoqLoadInt64_genType_doFn func(addr *int64)

// MoqLoadInt64_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the LoadInt64_genType type
type MoqLoadInt64_genType_doReturnFn func(addr *int64) (val int64)

// MoqLoadInt64_genType_results holds the results of the LoadInt64_genType type
type MoqLoadInt64_genType_results struct {
	Params  MoqLoadInt64_genType_params
	Results []struct {
		Values     *struct{ Val int64 }
		Sequence   uint32
		DoFn       MoqLoadInt64_genType_doFn
		DoReturnFn MoqLoadInt64_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqLoadInt64_genType_fnRecorder routes recorded function calls to the
// MoqLoadInt64_genType moq
type MoqLoadInt64_genType_fnRecorder struct {
	Params    MoqLoadInt64_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqLoadInt64_genType_results
	Moq       *MoqLoadInt64_genType
}

// MoqLoadInt64_genType_anyParams isolates the any params functions of the
// LoadInt64_genType type
type MoqLoadInt64_genType_anyParams struct {
	Recorder *MoqLoadInt64_genType_fnRecorder
}

// NewMoqLoadInt64_genType creates a new moq of the LoadInt64_genType type
func NewMoqLoadInt64_genType(scene *moq.Scene, config *moq.Config) *MoqLoadInt64_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqLoadInt64_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqLoadInt64_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Addr moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Addr moq.ParamIndexing
		}{
			Addr: moq.ParamIndexByHash,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the LoadInt64_genType type
func (m *MoqLoadInt64_genType) Mock() LoadInt64_genType {
	return func(addr *int64) (_ int64) {
		m.Scene.T.Helper()
		moq := &MoqLoadInt64_genType_mock{Moq: m}
		return moq.Fn(addr)
	}
}

func (m *MoqLoadInt64_genType_mock) Fn(addr *int64) (val int64) {
	m.Moq.Scene.T.Helper()
	params := MoqLoadInt64_genType_params{
		Addr: addr,
	}
	var results *MoqLoadInt64_genType_results
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
		result.DoFn(addr)
	}

	if result.Values != nil {
		val = result.Values.Val
	}
	if result.DoReturnFn != nil {
		val = result.DoReturnFn(addr)
	}
	return
}

func (m *MoqLoadInt64_genType) OnCall(addr *int64) *MoqLoadInt64_genType_fnRecorder {
	return &MoqLoadInt64_genType_fnRecorder{
		Params: MoqLoadInt64_genType_params{
			Addr: addr,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqLoadInt64_genType_fnRecorder) Any() *MoqLoadInt64_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqLoadInt64_genType_anyParams{Recorder: r}
}

func (a *MoqLoadInt64_genType_anyParams) Addr() *MoqLoadInt64_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqLoadInt64_genType_fnRecorder) Seq() *MoqLoadInt64_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqLoadInt64_genType_fnRecorder) NoSeq() *MoqLoadInt64_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqLoadInt64_genType_fnRecorder) ReturnResults(val int64) *MoqLoadInt64_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Val int64 }
		Sequence   uint32
		DoFn       MoqLoadInt64_genType_doFn
		DoReturnFn MoqLoadInt64_genType_doReturnFn
	}{
		Values: &struct{ Val int64 }{
			Val: val,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqLoadInt64_genType_fnRecorder) AndDo(fn MoqLoadInt64_genType_doFn) *MoqLoadInt64_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqLoadInt64_genType_fnRecorder) DoReturnResults(fn MoqLoadInt64_genType_doReturnFn) *MoqLoadInt64_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Val int64 }
		Sequence   uint32
		DoFn       MoqLoadInt64_genType_doFn
		DoReturnFn MoqLoadInt64_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqLoadInt64_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqLoadInt64_genType_resultsByParams
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
		results = &MoqLoadInt64_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqLoadInt64_genType_paramsKey]*MoqLoadInt64_genType_results{},
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
		r.Results = &MoqLoadInt64_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqLoadInt64_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqLoadInt64_genType_fnRecorder {
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
				Values     *struct{ Val int64 }
				Sequence   uint32
				DoFn       MoqLoadInt64_genType_doFn
				DoReturnFn MoqLoadInt64_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqLoadInt64_genType) PrettyParams(params MoqLoadInt64_genType_params) string {
	return fmt.Sprintf("LoadInt64_genType(%#v)", params.Addr)
}

func (m *MoqLoadInt64_genType) ParamsKey(params MoqLoadInt64_genType_params, anyParams uint64) MoqLoadInt64_genType_paramsKey {
	m.Scene.T.Helper()
	var addrUsed *int64
	var addrUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Addr == moq.ParamIndexByValue {
			addrUsed = params.Addr
		} else {
			addrUsedHash = hash.DeepHash(params.Addr)
		}
	}
	return MoqLoadInt64_genType_paramsKey{
		Params: struct{ Addr *int64 }{
			Addr: addrUsed,
		},
		Hashes: struct{ Addr hash.Hash }{
			Addr: addrUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqLoadInt64_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqLoadInt64_genType) AssertExpectationsMet() {
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
