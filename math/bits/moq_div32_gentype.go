// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package bits

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// Div32_genType is the fabricated implementation type of this mock (emitted
// when mocking functions directly and not from a function type)
type Div32_genType func(hi, lo, y uint32) (quo, rem uint32)

// MoqDiv32_genType holds the state of a moq of the Div32_genType type
type MoqDiv32_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqDiv32_genType_mock

	ResultsByParams []MoqDiv32_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Hi moq.ParamIndexing
			Lo moq.ParamIndexing
			Y  moq.ParamIndexing
		}
	}
}

// MoqDiv32_genType_mock isolates the mock interface of the Div32_genType type
type MoqDiv32_genType_mock struct {
	Moq *MoqDiv32_genType
}

// MoqDiv32_genType_params holds the params of the Div32_genType type
type MoqDiv32_genType_params struct{ Hi, Lo, Y uint32 }

// MoqDiv32_genType_paramsKey holds the map key params of the Div32_genType
// type
type MoqDiv32_genType_paramsKey struct {
	Params struct{ Hi, Lo, Y uint32 }
	Hashes struct{ Hi, Lo, Y hash.Hash }
}

// MoqDiv32_genType_resultsByParams contains the results for a given set of
// parameters for the Div32_genType type
type MoqDiv32_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqDiv32_genType_paramsKey]*MoqDiv32_genType_results
}

// MoqDiv32_genType_doFn defines the type of function needed when calling AndDo
// for the Div32_genType type
type MoqDiv32_genType_doFn func(hi, lo, y uint32)

// MoqDiv32_genType_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Div32_genType type
type MoqDiv32_genType_doReturnFn func(hi, lo, y uint32) (quo, rem uint32)

// MoqDiv32_genType_results holds the results of the Div32_genType type
type MoqDiv32_genType_results struct {
	Params  MoqDiv32_genType_params
	Results []struct {
		Values     *struct{ Quo, Rem uint32 }
		Sequence   uint32
		DoFn       MoqDiv32_genType_doFn
		DoReturnFn MoqDiv32_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqDiv32_genType_fnRecorder routes recorded function calls to the
// MoqDiv32_genType moq
type MoqDiv32_genType_fnRecorder struct {
	Params    MoqDiv32_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqDiv32_genType_results
	Moq       *MoqDiv32_genType
}

// MoqDiv32_genType_anyParams isolates the any params functions of the
// Div32_genType type
type MoqDiv32_genType_anyParams struct {
	Recorder *MoqDiv32_genType_fnRecorder
}

// NewMoqDiv32_genType creates a new moq of the Div32_genType type
func NewMoqDiv32_genType(scene *moq.Scene, config *moq.Config) *MoqDiv32_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqDiv32_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqDiv32_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Hi moq.ParamIndexing
				Lo moq.ParamIndexing
				Y  moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			Hi moq.ParamIndexing
			Lo moq.ParamIndexing
			Y  moq.ParamIndexing
		}{
			Hi: moq.ParamIndexByValue,
			Lo: moq.ParamIndexByValue,
			Y:  moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the Div32_genType type
func (m *MoqDiv32_genType) Mock() Div32_genType {
	return func(hi, lo, y uint32) (_, _ uint32) {
		m.Scene.T.Helper()
		moq := &MoqDiv32_genType_mock{Moq: m}
		return moq.Fn(hi, lo, y)
	}
}

func (m *MoqDiv32_genType_mock) Fn(hi, lo, y uint32) (quo, rem uint32) {
	m.Moq.Scene.T.Helper()
	params := MoqDiv32_genType_params{
		Hi: hi,
		Lo: lo,
		Y:  y,
	}
	var results *MoqDiv32_genType_results
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
		result.DoFn(hi, lo, y)
	}

	if result.Values != nil {
		quo = result.Values.Quo
		rem = result.Values.Rem
	}
	if result.DoReturnFn != nil {
		quo, rem = result.DoReturnFn(hi, lo, y)
	}
	return
}

func (m *MoqDiv32_genType) OnCall(hi, lo, y uint32) *MoqDiv32_genType_fnRecorder {
	return &MoqDiv32_genType_fnRecorder{
		Params: MoqDiv32_genType_params{
			Hi: hi,
			Lo: lo,
			Y:  y,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqDiv32_genType_fnRecorder) Any() *MoqDiv32_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqDiv32_genType_anyParams{Recorder: r}
}

func (a *MoqDiv32_genType_anyParams) Hi() *MoqDiv32_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqDiv32_genType_anyParams) Lo() *MoqDiv32_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqDiv32_genType_anyParams) Y() *MoqDiv32_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqDiv32_genType_fnRecorder) Seq() *MoqDiv32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqDiv32_genType_fnRecorder) NoSeq() *MoqDiv32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqDiv32_genType_fnRecorder) ReturnResults(quo, rem uint32) *MoqDiv32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Quo, Rem uint32 }
		Sequence   uint32
		DoFn       MoqDiv32_genType_doFn
		DoReturnFn MoqDiv32_genType_doReturnFn
	}{
		Values: &struct{ Quo, Rem uint32 }{
			Quo: quo,
			Rem: rem,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqDiv32_genType_fnRecorder) AndDo(fn MoqDiv32_genType_doFn) *MoqDiv32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqDiv32_genType_fnRecorder) DoReturnResults(fn MoqDiv32_genType_doReturnFn) *MoqDiv32_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values     *struct{ Quo, Rem uint32 }
		Sequence   uint32
		DoFn       MoqDiv32_genType_doFn
		DoReturnFn MoqDiv32_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqDiv32_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqDiv32_genType_resultsByParams
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
		results = &MoqDiv32_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqDiv32_genType_paramsKey]*MoqDiv32_genType_results{},
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
		r.Results = &MoqDiv32_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqDiv32_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqDiv32_genType_fnRecorder {
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
				Values     *struct{ Quo, Rem uint32 }
				Sequence   uint32
				DoFn       MoqDiv32_genType_doFn
				DoReturnFn MoqDiv32_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqDiv32_genType) PrettyParams(params MoqDiv32_genType_params) string {
	return fmt.Sprintf("Div32_genType(%#v, %#v, %#v)", params.Hi, params.Lo, params.Y)
}

func (m *MoqDiv32_genType) ParamsKey(params MoqDiv32_genType_params, anyParams uint64) MoqDiv32_genType_paramsKey {
	m.Scene.T.Helper()
	var hiUsed uint32
	var hiUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Hi == moq.ParamIndexByValue {
			hiUsed = params.Hi
		} else {
			hiUsedHash = hash.DeepHash(params.Hi)
		}
	}
	var loUsed uint32
	var loUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Lo == moq.ParamIndexByValue {
			loUsed = params.Lo
		} else {
			loUsedHash = hash.DeepHash(params.Lo)
		}
	}
	var yUsed uint32
	var yUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Y == moq.ParamIndexByValue {
			yUsed = params.Y
		} else {
			yUsedHash = hash.DeepHash(params.Y)
		}
	}
	return MoqDiv32_genType_paramsKey{
		Params: struct{ Hi, Lo, Y uint32 }{
			Hi: hiUsed,
			Lo: loUsed,
			Y:  yUsed,
		},
		Hashes: struct{ Hi, Lo, Y hash.Hash }{
			Hi: hiUsedHash,
			Lo: loUsedHash,
			Y:  yUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqDiv32_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqDiv32_genType) AssertExpectationsMet() {
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
