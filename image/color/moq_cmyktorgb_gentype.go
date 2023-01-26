// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package color

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// CMYKToRGB_genType is the fabricated implementation type of this mock
// (emitted when mocking functions directly and not from a function type)
type CMYKToRGB_genType func(c, m, y, k uint8) (uint8, uint8, uint8)

// MoqCMYKToRGB_genType holds the state of a moq of the CMYKToRGB_genType type
type MoqCMYKToRGB_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqCMYKToRGB_genType_mock

	ResultsByParams []MoqCMYKToRGB_genType_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			C      moq.ParamIndexing
			Param2 moq.ParamIndexing
			Y      moq.ParamIndexing
			K      moq.ParamIndexing
		}
	}
}

// MoqCMYKToRGB_genType_mock isolates the mock interface of the
// CMYKToRGB_genType type
type MoqCMYKToRGB_genType_mock struct {
	Moq *MoqCMYKToRGB_genType
}

// MoqCMYKToRGB_genType_params holds the params of the CMYKToRGB_genType type
type MoqCMYKToRGB_genType_params struct{ C, Param2, Y, K uint8 }

// MoqCMYKToRGB_genType_paramsKey holds the map key params of the
// CMYKToRGB_genType type
type MoqCMYKToRGB_genType_paramsKey struct {
	Params struct{ C, Param2, Y, K uint8 }
	Hashes struct{ C, Param2, Y, K hash.Hash }
}

// MoqCMYKToRGB_genType_resultsByParams contains the results for a given set of
// parameters for the CMYKToRGB_genType type
type MoqCMYKToRGB_genType_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqCMYKToRGB_genType_paramsKey]*MoqCMYKToRGB_genType_results
}

// MoqCMYKToRGB_genType_doFn defines the type of function needed when calling
// AndDo for the CMYKToRGB_genType type
type MoqCMYKToRGB_genType_doFn func(c, m, y, k uint8)

// MoqCMYKToRGB_genType_doReturnFn defines the type of function needed when
// calling DoReturnResults for the CMYKToRGB_genType type
type MoqCMYKToRGB_genType_doReturnFn func(c, m, y, k uint8) (uint8, uint8, uint8)

// MoqCMYKToRGB_genType_results holds the results of the CMYKToRGB_genType type
type MoqCMYKToRGB_genType_results struct {
	Params  MoqCMYKToRGB_genType_params
	Results []struct {
		Values *struct {
			Result1 uint8
			Result2 uint8
			Result3 uint8
		}
		Sequence   uint32
		DoFn       MoqCMYKToRGB_genType_doFn
		DoReturnFn MoqCMYKToRGB_genType_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqCMYKToRGB_genType_fnRecorder routes recorded function calls to the
// MoqCMYKToRGB_genType moq
type MoqCMYKToRGB_genType_fnRecorder struct {
	Params    MoqCMYKToRGB_genType_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqCMYKToRGB_genType_results
	Moq       *MoqCMYKToRGB_genType
}

// MoqCMYKToRGB_genType_anyParams isolates the any params functions of the
// CMYKToRGB_genType type
type MoqCMYKToRGB_genType_anyParams struct {
	Recorder *MoqCMYKToRGB_genType_fnRecorder
}

// NewMoqCMYKToRGB_genType creates a new moq of the CMYKToRGB_genType type
func NewMoqCMYKToRGB_genType(scene *moq.Scene, config *moq.Config) *MoqCMYKToRGB_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqCMYKToRGB_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqCMYKToRGB_genType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				C      moq.ParamIndexing
				Param2 moq.ParamIndexing
				Y      moq.ParamIndexing
				K      moq.ParamIndexing
			}
		}{ParameterIndexing: struct {
			C      moq.ParamIndexing
			Param2 moq.ParamIndexing
			Y      moq.ParamIndexing
			K      moq.ParamIndexing
		}{
			C:      moq.ParamIndexByValue,
			Param2: moq.ParamIndexByValue,
			Y:      moq.ParamIndexByValue,
			K:      moq.ParamIndexByValue,
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the moq implementation of the CMYKToRGB_genType type
func (m *MoqCMYKToRGB_genType) Mock() CMYKToRGB_genType {
	return func(c, param2, y, k uint8) (uint8, uint8, uint8) {
		m.Scene.T.Helper()
		moq := &MoqCMYKToRGB_genType_mock{Moq: m}
		return moq.Fn(c, param2, y, k)
	}
}

func (m *MoqCMYKToRGB_genType_mock) Fn(c, param2, y, k uint8) (result1 uint8, result2 uint8, result3 uint8) {
	m.Moq.Scene.T.Helper()
	params := MoqCMYKToRGB_genType_params{
		C:      c,
		Param2: param2,
		Y:      y,
		K:      k,
	}
	var results *MoqCMYKToRGB_genType_results
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
		result.DoFn(c, param2, y, k)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
		result3 = result.Values.Result3
	}
	if result.DoReturnFn != nil {
		result1, result2, result3 = result.DoReturnFn(c, param2, y, k)
	}
	return
}

func (m *MoqCMYKToRGB_genType) OnCall(c, param2, y, k uint8) *MoqCMYKToRGB_genType_fnRecorder {
	return &MoqCMYKToRGB_genType_fnRecorder{
		Params: MoqCMYKToRGB_genType_params{
			C:      c,
			Param2: param2,
			Y:      y,
			K:      k,
		},
		Sequence: m.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m,
	}
}

func (r *MoqCMYKToRGB_genType_fnRecorder) Any() *MoqCMYKToRGB_genType_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	return &MoqCMYKToRGB_genType_anyParams{Recorder: r}
}

func (a *MoqCMYKToRGB_genType_anyParams) C() *MoqCMYKToRGB_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqCMYKToRGB_genType_anyParams) Param2() *MoqCMYKToRGB_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqCMYKToRGB_genType_anyParams) Y() *MoqCMYKToRGB_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (a *MoqCMYKToRGB_genType_anyParams) K() *MoqCMYKToRGB_genType_fnRecorder {
	a.Recorder.AnyParams |= 1 << 3
	return a.Recorder
}

func (r *MoqCMYKToRGB_genType_fnRecorder) Seq() *MoqCMYKToRGB_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqCMYKToRGB_genType_fnRecorder) NoSeq() *MoqCMYKToRGB_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqCMYKToRGB_genType_fnRecorder) ReturnResults(result1 uint8, result2 uint8, result3 uint8) *MoqCMYKToRGB_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 uint8
			Result2 uint8
			Result3 uint8
		}
		Sequence   uint32
		DoFn       MoqCMYKToRGB_genType_doFn
		DoReturnFn MoqCMYKToRGB_genType_doReturnFn
	}{
		Values: &struct {
			Result1 uint8
			Result2 uint8
			Result3 uint8
		}{
			Result1: result1,
			Result2: result2,
			Result3: result3,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqCMYKToRGB_genType_fnRecorder) AndDo(fn MoqCMYKToRGB_genType_doFn) *MoqCMYKToRGB_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqCMYKToRGB_genType_fnRecorder) DoReturnResults(fn MoqCMYKToRGB_genType_doReturnFn) *MoqCMYKToRGB_genType_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 uint8
			Result2 uint8
			Result3 uint8
		}
		Sequence   uint32
		DoFn       MoqCMYKToRGB_genType_doFn
		DoReturnFn MoqCMYKToRGB_genType_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqCMYKToRGB_genType_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqCMYKToRGB_genType_resultsByParams
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
		results = &MoqCMYKToRGB_genType_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqCMYKToRGB_genType_paramsKey]*MoqCMYKToRGB_genType_results{},
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
		r.Results = &MoqCMYKToRGB_genType_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqCMYKToRGB_genType_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqCMYKToRGB_genType_fnRecorder {
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
					Result1 uint8
					Result2 uint8
					Result3 uint8
				}
				Sequence   uint32
				DoFn       MoqCMYKToRGB_genType_doFn
				DoReturnFn MoqCMYKToRGB_genType_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqCMYKToRGB_genType) PrettyParams(params MoqCMYKToRGB_genType_params) string {
	return fmt.Sprintf("CMYKToRGB_genType(%#v, %#v, %#v, %#v)", params.C, params.Param2, params.Y, params.K)
}

func (m *MoqCMYKToRGB_genType) ParamsKey(params MoqCMYKToRGB_genType_params, anyParams uint64) MoqCMYKToRGB_genType_paramsKey {
	m.Scene.T.Helper()
	var cUsed uint8
	var cUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.C == moq.ParamIndexByValue {
			cUsed = params.C
		} else {
			cUsedHash = hash.DeepHash(params.C)
		}
	}
	var param2Used uint8
	var param2UsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Param2 == moq.ParamIndexByValue {
			param2Used = params.Param2
		} else {
			param2UsedHash = hash.DeepHash(params.Param2)
		}
	}
	var yUsed uint8
	var yUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Y == moq.ParamIndexByValue {
			yUsed = params.Y
		} else {
			yUsedHash = hash.DeepHash(params.Y)
		}
	}
	var kUsed uint8
	var kUsedHash hash.Hash
	if anyParams&(1<<3) == 0 {
		if m.Runtime.ParameterIndexing.K == moq.ParamIndexByValue {
			kUsed = params.K
		} else {
			kUsedHash = hash.DeepHash(params.K)
		}
	}
	return MoqCMYKToRGB_genType_paramsKey{
		Params: struct{ C, Param2, Y, K uint8 }{
			C:      cUsed,
			Param2: param2Used,
			Y:      yUsed,
			K:      kUsed,
		},
		Hashes: struct{ C, Param2, Y, K hash.Hash }{
			C:      cUsedHash,
			Param2: param2UsedHash,
			Y:      yUsedHash,
			K:      kUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqCMYKToRGB_genType) Reset() { m.ResultsByParams = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqCMYKToRGB_genType) AssertExpectationsMet() {
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
