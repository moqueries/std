// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package sql

import (
	"database/sql"
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that sql.Result is mocked completely
var _ sql.Result = (*MoqResult_mock)(nil)

// MoqResult holds the state of a moq of the Result type
type MoqResult struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqResult_mock

	ResultsByParams_LastInsertId []MoqResult_LastInsertId_resultsByParams
	ResultsByParams_RowsAffected []MoqResult_RowsAffected_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			LastInsertId struct{}
			RowsAffected struct{}
		}
	}
}

// MoqResult_mock isolates the mock interface of the Result type
type MoqResult_mock struct {
	Moq *MoqResult
}

// MoqResult_recorder isolates the recorder interface of the Result type
type MoqResult_recorder struct {
	Moq *MoqResult
}

// MoqResult_LastInsertId_params holds the params of the Result type
type MoqResult_LastInsertId_params struct{}

// MoqResult_LastInsertId_paramsKey holds the map key params of the Result type
type MoqResult_LastInsertId_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqResult_LastInsertId_resultsByParams contains the results for a given set
// of parameters for the Result type
type MoqResult_LastInsertId_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqResult_LastInsertId_paramsKey]*MoqResult_LastInsertId_results
}

// MoqResult_LastInsertId_doFn defines the type of function needed when calling
// AndDo for the Result type
type MoqResult_LastInsertId_doFn func()

// MoqResult_LastInsertId_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Result type
type MoqResult_LastInsertId_doReturnFn func() (int64, error)

// MoqResult_LastInsertId_results holds the results of the Result type
type MoqResult_LastInsertId_results struct {
	Params  MoqResult_LastInsertId_params
	Results []struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqResult_LastInsertId_doFn
		DoReturnFn MoqResult_LastInsertId_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqResult_LastInsertId_fnRecorder routes recorded function calls to the
// MoqResult moq
type MoqResult_LastInsertId_fnRecorder struct {
	Params    MoqResult_LastInsertId_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqResult_LastInsertId_results
	Moq       *MoqResult
}

// MoqResult_LastInsertId_anyParams isolates the any params functions of the
// Result type
type MoqResult_LastInsertId_anyParams struct {
	Recorder *MoqResult_LastInsertId_fnRecorder
}

// MoqResult_RowsAffected_params holds the params of the Result type
type MoqResult_RowsAffected_params struct{}

// MoqResult_RowsAffected_paramsKey holds the map key params of the Result type
type MoqResult_RowsAffected_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqResult_RowsAffected_resultsByParams contains the results for a given set
// of parameters for the Result type
type MoqResult_RowsAffected_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqResult_RowsAffected_paramsKey]*MoqResult_RowsAffected_results
}

// MoqResult_RowsAffected_doFn defines the type of function needed when calling
// AndDo for the Result type
type MoqResult_RowsAffected_doFn func()

// MoqResult_RowsAffected_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Result type
type MoqResult_RowsAffected_doReturnFn func() (int64, error)

// MoqResult_RowsAffected_results holds the results of the Result type
type MoqResult_RowsAffected_results struct {
	Params  MoqResult_RowsAffected_params
	Results []struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqResult_RowsAffected_doFn
		DoReturnFn MoqResult_RowsAffected_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqResult_RowsAffected_fnRecorder routes recorded function calls to the
// MoqResult moq
type MoqResult_RowsAffected_fnRecorder struct {
	Params    MoqResult_RowsAffected_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqResult_RowsAffected_results
	Moq       *MoqResult
}

// MoqResult_RowsAffected_anyParams isolates the any params functions of the
// Result type
type MoqResult_RowsAffected_anyParams struct {
	Recorder *MoqResult_RowsAffected_fnRecorder
}

// NewMoqResult creates a new moq of the Result type
func NewMoqResult(scene *moq.Scene, config *moq.Config) *MoqResult {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqResult{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqResult_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				LastInsertId struct{}
				RowsAffected struct{}
			}
		}{ParameterIndexing: struct {
			LastInsertId struct{}
			RowsAffected struct{}
		}{
			LastInsertId: struct{}{},
			RowsAffected: struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Result type
func (m *MoqResult) Mock() *MoqResult_mock { return m.Moq }

func (m *MoqResult_mock) LastInsertId() (result1 int64, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqResult_LastInsertId_params{}
	var results *MoqResult_LastInsertId_results
	for _, resultsByParams := range m.Moq.ResultsByParams_LastInsertId {
		paramsKey := m.Moq.ParamsKey_LastInsertId(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_LastInsertId(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_LastInsertId(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_LastInsertId(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

func (m *MoqResult_mock) RowsAffected() (result1 int64, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqResult_RowsAffected_params{}
	var results *MoqResult_RowsAffected_results
	for _, resultsByParams := range m.Moq.ResultsByParams_RowsAffected {
		paramsKey := m.Moq.ParamsKey_RowsAffected(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_RowsAffected(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_RowsAffected(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_RowsAffected(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn()
	}
	return
}

// OnCall returns the recorder implementation of the Result type
func (m *MoqResult) OnCall() *MoqResult_recorder {
	return &MoqResult_recorder{
		Moq: m,
	}
}

func (m *MoqResult_recorder) LastInsertId() *MoqResult_LastInsertId_fnRecorder {
	return &MoqResult_LastInsertId_fnRecorder{
		Params:   MoqResult_LastInsertId_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqResult_LastInsertId_fnRecorder) Any() *MoqResult_LastInsertId_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_LastInsertId(r.Params))
		return nil
	}
	return &MoqResult_LastInsertId_anyParams{Recorder: r}
}

func (r *MoqResult_LastInsertId_fnRecorder) Seq() *MoqResult_LastInsertId_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_LastInsertId(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqResult_LastInsertId_fnRecorder) NoSeq() *MoqResult_LastInsertId_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_LastInsertId(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqResult_LastInsertId_fnRecorder) ReturnResults(result1 int64, result2 error) *MoqResult_LastInsertId_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqResult_LastInsertId_doFn
		DoReturnFn MoqResult_LastInsertId_doReturnFn
	}{
		Values: &struct {
			Result1 int64
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqResult_LastInsertId_fnRecorder) AndDo(fn MoqResult_LastInsertId_doFn) *MoqResult_LastInsertId_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqResult_LastInsertId_fnRecorder) DoReturnResults(fn MoqResult_LastInsertId_doReturnFn) *MoqResult_LastInsertId_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqResult_LastInsertId_doFn
		DoReturnFn MoqResult_LastInsertId_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqResult_LastInsertId_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqResult_LastInsertId_resultsByParams
	for n, res := range r.Moq.ResultsByParams_LastInsertId {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqResult_LastInsertId_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqResult_LastInsertId_paramsKey]*MoqResult_LastInsertId_results{},
		}
		r.Moq.ResultsByParams_LastInsertId = append(r.Moq.ResultsByParams_LastInsertId, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_LastInsertId) {
			copy(r.Moq.ResultsByParams_LastInsertId[insertAt+1:], r.Moq.ResultsByParams_LastInsertId[insertAt:0])
			r.Moq.ResultsByParams_LastInsertId[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_LastInsertId(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqResult_LastInsertId_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqResult_LastInsertId_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqResult_LastInsertId_fnRecorder {
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
					Result1 int64
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqResult_LastInsertId_doFn
				DoReturnFn MoqResult_LastInsertId_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqResult) PrettyParams_LastInsertId(params MoqResult_LastInsertId_params) string {
	return fmt.Sprintf("LastInsertId()")
}

func (m *MoqResult) ParamsKey_LastInsertId(params MoqResult_LastInsertId_params, anyParams uint64) MoqResult_LastInsertId_paramsKey {
	m.Scene.T.Helper()
	return MoqResult_LastInsertId_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqResult_recorder) RowsAffected() *MoqResult_RowsAffected_fnRecorder {
	return &MoqResult_RowsAffected_fnRecorder{
		Params:   MoqResult_RowsAffected_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqResult_RowsAffected_fnRecorder) Any() *MoqResult_RowsAffected_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RowsAffected(r.Params))
		return nil
	}
	return &MoqResult_RowsAffected_anyParams{Recorder: r}
}

func (r *MoqResult_RowsAffected_fnRecorder) Seq() *MoqResult_RowsAffected_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RowsAffected(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqResult_RowsAffected_fnRecorder) NoSeq() *MoqResult_RowsAffected_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RowsAffected(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqResult_RowsAffected_fnRecorder) ReturnResults(result1 int64, result2 error) *MoqResult_RowsAffected_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqResult_RowsAffected_doFn
		DoReturnFn MoqResult_RowsAffected_doReturnFn
	}{
		Values: &struct {
			Result1 int64
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqResult_RowsAffected_fnRecorder) AndDo(fn MoqResult_RowsAffected_doFn) *MoqResult_RowsAffected_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqResult_RowsAffected_fnRecorder) DoReturnResults(fn MoqResult_RowsAffected_doReturnFn) *MoqResult_RowsAffected_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqResult_RowsAffected_doFn
		DoReturnFn MoqResult_RowsAffected_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqResult_RowsAffected_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqResult_RowsAffected_resultsByParams
	for n, res := range r.Moq.ResultsByParams_RowsAffected {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqResult_RowsAffected_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqResult_RowsAffected_paramsKey]*MoqResult_RowsAffected_results{},
		}
		r.Moq.ResultsByParams_RowsAffected = append(r.Moq.ResultsByParams_RowsAffected, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_RowsAffected) {
			copy(r.Moq.ResultsByParams_RowsAffected[insertAt+1:], r.Moq.ResultsByParams_RowsAffected[insertAt:0])
			r.Moq.ResultsByParams_RowsAffected[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_RowsAffected(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqResult_RowsAffected_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqResult_RowsAffected_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqResult_RowsAffected_fnRecorder {
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
					Result1 int64
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqResult_RowsAffected_doFn
				DoReturnFn MoqResult_RowsAffected_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqResult) PrettyParams_RowsAffected(params MoqResult_RowsAffected_params) string {
	return fmt.Sprintf("RowsAffected()")
}

func (m *MoqResult) ParamsKey_RowsAffected(params MoqResult_RowsAffected_params, anyParams uint64) MoqResult_RowsAffected_paramsKey {
	m.Scene.T.Helper()
	return MoqResult_RowsAffected_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqResult) Reset() {
	m.ResultsByParams_LastInsertId = nil
	m.ResultsByParams_RowsAffected = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqResult) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_LastInsertId {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_LastInsertId(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_RowsAffected {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_RowsAffected(results.Params))
			}
		}
	}
}