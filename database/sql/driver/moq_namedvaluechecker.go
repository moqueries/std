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

// The following type assertion assures that driver.NamedValueChecker is mocked
// completely
var _ driver.NamedValueChecker = (*MoqNamedValueChecker_mock)(nil)

// MoqNamedValueChecker holds the state of a moq of the NamedValueChecker type
type MoqNamedValueChecker struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqNamedValueChecker_mock

	ResultsByParams_CheckNamedValue []MoqNamedValueChecker_CheckNamedValue_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			CheckNamedValue struct {
				Param1 moq.ParamIndexing
			}
		}
	}
	// MoqNamedValueChecker_mock isolates the mock interface of the
}

// NamedValueChecker type
type MoqNamedValueChecker_mock struct {
	Moq *MoqNamedValueChecker
}

// MoqNamedValueChecker_recorder isolates the recorder interface of the
// NamedValueChecker type
type MoqNamedValueChecker_recorder struct {
	Moq *MoqNamedValueChecker
}

// MoqNamedValueChecker_CheckNamedValue_params holds the params of the
// NamedValueChecker type
type MoqNamedValueChecker_CheckNamedValue_params struct{ Param1 *driver.NamedValue }

// MoqNamedValueChecker_CheckNamedValue_paramsKey holds the map key params of
// the NamedValueChecker type
type MoqNamedValueChecker_CheckNamedValue_paramsKey struct {
	Params struct{ Param1 *driver.NamedValue }
	Hashes struct{ Param1 hash.Hash }
}

// MoqNamedValueChecker_CheckNamedValue_resultsByParams contains the results
// for a given set of parameters for the NamedValueChecker type
type MoqNamedValueChecker_CheckNamedValue_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqNamedValueChecker_CheckNamedValue_paramsKey]*MoqNamedValueChecker_CheckNamedValue_results
}

// MoqNamedValueChecker_CheckNamedValue_doFn defines the type of function
// needed when calling AndDo for the NamedValueChecker type
type MoqNamedValueChecker_CheckNamedValue_doFn func(*driver.NamedValue)

// MoqNamedValueChecker_CheckNamedValue_doReturnFn defines the type of function
// needed when calling DoReturnResults for the NamedValueChecker type
type MoqNamedValueChecker_CheckNamedValue_doReturnFn func(*driver.NamedValue) error

// MoqNamedValueChecker_CheckNamedValue_results holds the results of the
// NamedValueChecker type
type MoqNamedValueChecker_CheckNamedValue_results struct {
	Params  MoqNamedValueChecker_CheckNamedValue_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqNamedValueChecker_CheckNamedValue_doFn
		DoReturnFn MoqNamedValueChecker_CheckNamedValue_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqNamedValueChecker_CheckNamedValue_fnRecorder routes recorded function
// calls to the MoqNamedValueChecker moq
type MoqNamedValueChecker_CheckNamedValue_fnRecorder struct {
	Params    MoqNamedValueChecker_CheckNamedValue_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqNamedValueChecker_CheckNamedValue_results
	Moq       *MoqNamedValueChecker
}

// MoqNamedValueChecker_CheckNamedValue_anyParams isolates the any params
// functions of the NamedValueChecker type
type MoqNamedValueChecker_CheckNamedValue_anyParams struct {
	Recorder *MoqNamedValueChecker_CheckNamedValue_fnRecorder
}

// NewMoqNamedValueChecker creates a new moq of the NamedValueChecker type
func NewMoqNamedValueChecker(scene *moq.Scene, config *moq.Config) *MoqNamedValueChecker {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqNamedValueChecker{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqNamedValueChecker_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				CheckNamedValue struct {
					Param1 moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			CheckNamedValue struct {
				Param1 moq.ParamIndexing
			}
		}{
			CheckNamedValue: struct {
				Param1 moq.ParamIndexing
			}{
				Param1: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the NamedValueChecker type
func (m *MoqNamedValueChecker) Mock() *MoqNamedValueChecker_mock { return m.Moq }

func (m *MoqNamedValueChecker_mock) CheckNamedValue(param1 *driver.NamedValue) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqNamedValueChecker_CheckNamedValue_params{
		Param1: param1,
	}
	var results *MoqNamedValueChecker_CheckNamedValue_results
	for _, resultsByParams := range m.Moq.ResultsByParams_CheckNamedValue {
		paramsKey := m.Moq.ParamsKey_CheckNamedValue(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_CheckNamedValue(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_CheckNamedValue(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_CheckNamedValue(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(param1)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(param1)
	}
	return
}

// OnCall returns the recorder implementation of the NamedValueChecker type
func (m *MoqNamedValueChecker) OnCall() *MoqNamedValueChecker_recorder {
	return &MoqNamedValueChecker_recorder{
		Moq: m,
	}
}

func (m *MoqNamedValueChecker_recorder) CheckNamedValue(param1 *driver.NamedValue) *MoqNamedValueChecker_CheckNamedValue_fnRecorder {
	return &MoqNamedValueChecker_CheckNamedValue_fnRecorder{
		Params: MoqNamedValueChecker_CheckNamedValue_params{
			Param1: param1,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqNamedValueChecker_CheckNamedValue_fnRecorder) Any() *MoqNamedValueChecker_CheckNamedValue_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_CheckNamedValue(r.Params))
		return nil
	}
	return &MoqNamedValueChecker_CheckNamedValue_anyParams{Recorder: r}
}

func (a *MoqNamedValueChecker_CheckNamedValue_anyParams) Param1() *MoqNamedValueChecker_CheckNamedValue_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqNamedValueChecker_CheckNamedValue_fnRecorder) Seq() *MoqNamedValueChecker_CheckNamedValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_CheckNamedValue(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqNamedValueChecker_CheckNamedValue_fnRecorder) NoSeq() *MoqNamedValueChecker_CheckNamedValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_CheckNamedValue(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqNamedValueChecker_CheckNamedValue_fnRecorder) ReturnResults(result1 error) *MoqNamedValueChecker_CheckNamedValue_fnRecorder {
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
		DoFn       MoqNamedValueChecker_CheckNamedValue_doFn
		DoReturnFn MoqNamedValueChecker_CheckNamedValue_doReturnFn
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

func (r *MoqNamedValueChecker_CheckNamedValue_fnRecorder) AndDo(fn MoqNamedValueChecker_CheckNamedValue_doFn) *MoqNamedValueChecker_CheckNamedValue_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqNamedValueChecker_CheckNamedValue_fnRecorder) DoReturnResults(fn MoqNamedValueChecker_CheckNamedValue_doReturnFn) *MoqNamedValueChecker_CheckNamedValue_fnRecorder {
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
		DoFn       MoqNamedValueChecker_CheckNamedValue_doFn
		DoReturnFn MoqNamedValueChecker_CheckNamedValue_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqNamedValueChecker_CheckNamedValue_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqNamedValueChecker_CheckNamedValue_resultsByParams
	for n, res := range r.Moq.ResultsByParams_CheckNamedValue {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqNamedValueChecker_CheckNamedValue_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqNamedValueChecker_CheckNamedValue_paramsKey]*MoqNamedValueChecker_CheckNamedValue_results{},
		}
		r.Moq.ResultsByParams_CheckNamedValue = append(r.Moq.ResultsByParams_CheckNamedValue, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_CheckNamedValue) {
			copy(r.Moq.ResultsByParams_CheckNamedValue[insertAt+1:], r.Moq.ResultsByParams_CheckNamedValue[insertAt:0])
			r.Moq.ResultsByParams_CheckNamedValue[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_CheckNamedValue(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqNamedValueChecker_CheckNamedValue_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqNamedValueChecker_CheckNamedValue_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqNamedValueChecker_CheckNamedValue_fnRecorder {
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
				DoFn       MoqNamedValueChecker_CheckNamedValue_doFn
				DoReturnFn MoqNamedValueChecker_CheckNamedValue_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqNamedValueChecker) PrettyParams_CheckNamedValue(params MoqNamedValueChecker_CheckNamedValue_params) string {
	return fmt.Sprintf("CheckNamedValue(%#v)", params.Param1)
}

func (m *MoqNamedValueChecker) ParamsKey_CheckNamedValue(params MoqNamedValueChecker_CheckNamedValue_params, anyParams uint64) MoqNamedValueChecker_CheckNamedValue_paramsKey {
	m.Scene.T.Helper()
	var param1Used *driver.NamedValue
	var param1UsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.CheckNamedValue.Param1 == moq.ParamIndexByValue {
			param1Used = params.Param1
		} else {
			param1UsedHash = hash.DeepHash(params.Param1)
		}
	}
	return MoqNamedValueChecker_CheckNamedValue_paramsKey{
		Params: struct{ Param1 *driver.NamedValue }{
			Param1: param1Used,
		},
		Hashes: struct{ Param1 hash.Hash }{
			Param1: param1UsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqNamedValueChecker) Reset() { m.ResultsByParams_CheckNamedValue = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqNamedValueChecker) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_CheckNamedValue {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_CheckNamedValue(results.Params))
			}
		}
	}
}
