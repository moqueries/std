// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package driver

import (
	"database/sql/driver"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that driver.ValueConverter is mocked
// completely
var _ driver.ValueConverter = (*MoqValueConverter_mock)(nil)

// MoqValueConverter holds the state of a moq of the ValueConverter type
type MoqValueConverter struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqValueConverter_mock

	ResultsByParams_ConvertValue []MoqValueConverter_ConvertValue_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			ConvertValue struct {
				V moq.ParamIndexing
			}
		}
	}
	// MoqValueConverter_mock isolates the mock interface of the ValueConverter
}

// type
type MoqValueConverter_mock struct {
	Moq *MoqValueConverter
}

// MoqValueConverter_recorder isolates the recorder interface of the
// ValueConverter type
type MoqValueConverter_recorder struct {
	Moq *MoqValueConverter
}

// MoqValueConverter_ConvertValue_params holds the params of the ValueConverter
// type
type MoqValueConverter_ConvertValue_params struct{ V any }

// MoqValueConverter_ConvertValue_paramsKey holds the map key params of the
// ValueConverter type
type MoqValueConverter_ConvertValue_paramsKey struct {
	Params struct{ V any }
	Hashes struct{ V hash.Hash }
}

// MoqValueConverter_ConvertValue_resultsByParams contains the results for a
// given set of parameters for the ValueConverter type
type MoqValueConverter_ConvertValue_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqValueConverter_ConvertValue_paramsKey]*MoqValueConverter_ConvertValue_results
}

// MoqValueConverter_ConvertValue_doFn defines the type of function needed when
// calling AndDo for the ValueConverter type
type MoqValueConverter_ConvertValue_doFn func(v any)

// MoqValueConverter_ConvertValue_doReturnFn defines the type of function
// needed when calling DoReturnResults for the ValueConverter type
type MoqValueConverter_ConvertValue_doReturnFn func(v any) (driver.Value, error)

// MoqValueConverter_ConvertValue_results holds the results of the
// ValueConverter type
type MoqValueConverter_ConvertValue_results struct {
	Params  MoqValueConverter_ConvertValue_params
	Results []struct {
		Values *struct {
			Result1 driver.Value
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqValueConverter_ConvertValue_doFn
		DoReturnFn MoqValueConverter_ConvertValue_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqValueConverter_ConvertValue_fnRecorder routes recorded function calls to
// the MoqValueConverter moq
type MoqValueConverter_ConvertValue_fnRecorder struct {
	Params    MoqValueConverter_ConvertValue_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqValueConverter_ConvertValue_results
	Moq       *MoqValueConverter
}

// MoqValueConverter_ConvertValue_anyParams isolates the any params functions
// of the ValueConverter type
type MoqValueConverter_ConvertValue_anyParams struct {
	Recorder *MoqValueConverter_ConvertValue_fnRecorder
}

// NewMoqValueConverter creates a new moq of the ValueConverter type
func NewMoqValueConverter(scene *moq.Scene, config *moq.Config) *MoqValueConverter {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqValueConverter{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqValueConverter_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				ConvertValue struct {
					V moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			ConvertValue struct {
				V moq.ParamIndexing
			}
		}{
			ConvertValue: struct {
				V moq.ParamIndexing
			}{
				V: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the ValueConverter type
func (m *MoqValueConverter) Mock() *MoqValueConverter_mock { return m.Moq }

func (m *MoqValueConverter_mock) ConvertValue(v any) (result1 driver.Value, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqValueConverter_ConvertValue_params{
		V: v,
	}
	var results *MoqValueConverter_ConvertValue_results
	for _, resultsByParams := range m.Moq.ResultsByParams_ConvertValue {
		paramsKey := m.Moq.ParamsKey_ConvertValue(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_ConvertValue(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_ConvertValue(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_ConvertValue(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(v)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(v)
	}
	return
}

// OnCall returns the recorder implementation of the ValueConverter type
func (m *MoqValueConverter) OnCall() *MoqValueConverter_recorder {
	return &MoqValueConverter_recorder{
		Moq: m,
	}
}

func (m *MoqValueConverter_recorder) ConvertValue(v any) *MoqValueConverter_ConvertValue_fnRecorder {
	return &MoqValueConverter_ConvertValue_fnRecorder{
		Params: MoqValueConverter_ConvertValue_params{
			V: v,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqValueConverter_ConvertValue_fnRecorder) Any() *MoqValueConverter_ConvertValue_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ConvertValue(r.Params))
		return nil
	}
	return &MoqValueConverter_ConvertValue_anyParams{Recorder: r}
}

func (a *MoqValueConverter_ConvertValue_anyParams) V() *MoqValueConverter_ConvertValue_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqValueConverter_ConvertValue_fnRecorder) Seq() *MoqValueConverter_ConvertValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ConvertValue(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqValueConverter_ConvertValue_fnRecorder) NoSeq() *MoqValueConverter_ConvertValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_ConvertValue(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqValueConverter_ConvertValue_fnRecorder) ReturnResults(result1 driver.Value, result2 error) *MoqValueConverter_ConvertValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Value
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqValueConverter_ConvertValue_doFn
		DoReturnFn MoqValueConverter_ConvertValue_doReturnFn
	}{
		Values: &struct {
			Result1 driver.Value
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqValueConverter_ConvertValue_fnRecorder) AndDo(fn MoqValueConverter_ConvertValue_doFn) *MoqValueConverter_ConvertValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqValueConverter_ConvertValue_fnRecorder) DoReturnResults(fn MoqValueConverter_ConvertValue_doReturnFn) *MoqValueConverter_ConvertValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Value
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqValueConverter_ConvertValue_doFn
		DoReturnFn MoqValueConverter_ConvertValue_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqValueConverter_ConvertValue_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqValueConverter_ConvertValue_resultsByParams
	for n, res := range r.Moq.ResultsByParams_ConvertValue {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqValueConverter_ConvertValue_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqValueConverter_ConvertValue_paramsKey]*MoqValueConverter_ConvertValue_results{},
		}
		r.Moq.ResultsByParams_ConvertValue = append(r.Moq.ResultsByParams_ConvertValue, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_ConvertValue) {
			copy(r.Moq.ResultsByParams_ConvertValue[insertAt+1:], r.Moq.ResultsByParams_ConvertValue[insertAt:0])
			r.Moq.ResultsByParams_ConvertValue[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_ConvertValue(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqValueConverter_ConvertValue_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqValueConverter_ConvertValue_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqValueConverter_ConvertValue_fnRecorder {
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
					Result1 driver.Value
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqValueConverter_ConvertValue_doFn
				DoReturnFn MoqValueConverter_ConvertValue_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqValueConverter) PrettyParams_ConvertValue(params MoqValueConverter_ConvertValue_params) string {
	return fmt.Sprintf("ConvertValue(%#v)", params.V)
}

func (m *MoqValueConverter) ParamsKey_ConvertValue(params MoqValueConverter_ConvertValue_params, anyParams uint64) MoqValueConverter_ConvertValue_paramsKey {
	m.Scene.T.Helper()
	var vUsed any
	var vUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.ConvertValue.V == moq.ParamIndexByValue {
			vUsed = params.V
		} else {
			vUsedHash = hash.DeepHash(params.V)
		}
	}
	return MoqValueConverter_ConvertValue_paramsKey{
		Params: struct{ V any }{
			V: vUsed,
		},
		Hashes: struct{ V hash.Hash }{
			V: vUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqValueConverter) Reset() { m.ResultsByParams_ConvertValue = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqValueConverter) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_ConvertValue {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_ConvertValue(results.Params))
			}
		}
	}
}
