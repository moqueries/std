// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package driver

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that driver.RowsAffected_genType is
// mocked completely
var _ RowsAffected_genType = (*MoqRowsAffected_genType_mock)(nil)

// RowsAffected_genType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type RowsAffected_genType interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// MoqRowsAffected_genType holds the state of a moq of the RowsAffected_genType
// type
type MoqRowsAffected_genType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqRowsAffected_genType_mock

	ResultsByParams_LastInsertId []MoqRowsAffected_genType_LastInsertId_resultsByParams
	ResultsByParams_RowsAffected []MoqRowsAffected_genType_RowsAffected_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			LastInsertId struct{}
			RowsAffected struct{}
		}
	}
}

// MoqRowsAffected_genType_mock isolates the mock interface of the
// RowsAffected_genType type
type MoqRowsAffected_genType_mock struct {
	Moq *MoqRowsAffected_genType
}

// MoqRowsAffected_genType_recorder isolates the recorder interface of the
// RowsAffected_genType type
type MoqRowsAffected_genType_recorder struct {
	Moq *MoqRowsAffected_genType
}

// MoqRowsAffected_genType_LastInsertId_params holds the params of the
// RowsAffected_genType type
type MoqRowsAffected_genType_LastInsertId_params struct{}

// MoqRowsAffected_genType_LastInsertId_paramsKey holds the map key params of
// the RowsAffected_genType type
type MoqRowsAffected_genType_LastInsertId_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqRowsAffected_genType_LastInsertId_resultsByParams contains the results
// for a given set of parameters for the RowsAffected_genType type
type MoqRowsAffected_genType_LastInsertId_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRowsAffected_genType_LastInsertId_paramsKey]*MoqRowsAffected_genType_LastInsertId_results
}

// MoqRowsAffected_genType_LastInsertId_doFn defines the type of function
// needed when calling AndDo for the RowsAffected_genType type
type MoqRowsAffected_genType_LastInsertId_doFn func()

// MoqRowsAffected_genType_LastInsertId_doReturnFn defines the type of function
// needed when calling DoReturnResults for the RowsAffected_genType type
type MoqRowsAffected_genType_LastInsertId_doReturnFn func() (int64, error)

// MoqRowsAffected_genType_LastInsertId_results holds the results of the
// RowsAffected_genType type
type MoqRowsAffected_genType_LastInsertId_results struct {
	Params  MoqRowsAffected_genType_LastInsertId_params
	Results []struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRowsAffected_genType_LastInsertId_doFn
		DoReturnFn MoqRowsAffected_genType_LastInsertId_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRowsAffected_genType_LastInsertId_fnRecorder routes recorded function
// calls to the MoqRowsAffected_genType moq
type MoqRowsAffected_genType_LastInsertId_fnRecorder struct {
	Params    MoqRowsAffected_genType_LastInsertId_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRowsAffected_genType_LastInsertId_results
	Moq       *MoqRowsAffected_genType
}

// MoqRowsAffected_genType_LastInsertId_anyParams isolates the any params
// functions of the RowsAffected_genType type
type MoqRowsAffected_genType_LastInsertId_anyParams struct {
	Recorder *MoqRowsAffected_genType_LastInsertId_fnRecorder
}

// MoqRowsAffected_genType_RowsAffected_params holds the params of the
// RowsAffected_genType type
type MoqRowsAffected_genType_RowsAffected_params struct{}

// MoqRowsAffected_genType_RowsAffected_paramsKey holds the map key params of
// the RowsAffected_genType type
type MoqRowsAffected_genType_RowsAffected_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqRowsAffected_genType_RowsAffected_resultsByParams contains the results
// for a given set of parameters for the RowsAffected_genType type
type MoqRowsAffected_genType_RowsAffected_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqRowsAffected_genType_RowsAffected_paramsKey]*MoqRowsAffected_genType_RowsAffected_results
}

// MoqRowsAffected_genType_RowsAffected_doFn defines the type of function
// needed when calling AndDo for the RowsAffected_genType type
type MoqRowsAffected_genType_RowsAffected_doFn func()

// MoqRowsAffected_genType_RowsAffected_doReturnFn defines the type of function
// needed when calling DoReturnResults for the RowsAffected_genType type
type MoqRowsAffected_genType_RowsAffected_doReturnFn func() (int64, error)

// MoqRowsAffected_genType_RowsAffected_results holds the results of the
// RowsAffected_genType type
type MoqRowsAffected_genType_RowsAffected_results struct {
	Params  MoqRowsAffected_genType_RowsAffected_params
	Results []struct {
		Values *struct {
			Result1 int64
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqRowsAffected_genType_RowsAffected_doFn
		DoReturnFn MoqRowsAffected_genType_RowsAffected_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqRowsAffected_genType_RowsAffected_fnRecorder routes recorded function
// calls to the MoqRowsAffected_genType moq
type MoqRowsAffected_genType_RowsAffected_fnRecorder struct {
	Params    MoqRowsAffected_genType_RowsAffected_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqRowsAffected_genType_RowsAffected_results
	Moq       *MoqRowsAffected_genType
}

// MoqRowsAffected_genType_RowsAffected_anyParams isolates the any params
// functions of the RowsAffected_genType type
type MoqRowsAffected_genType_RowsAffected_anyParams struct {
	Recorder *MoqRowsAffected_genType_RowsAffected_fnRecorder
}

// NewMoqRowsAffected_genType creates a new moq of the RowsAffected_genType
// type
func NewMoqRowsAffected_genType(scene *moq.Scene, config *moq.Config) *MoqRowsAffected_genType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqRowsAffected_genType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqRowsAffected_genType_mock{},

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

// Mock returns the mock implementation of the RowsAffected_genType type
func (m *MoqRowsAffected_genType) Mock() *MoqRowsAffected_genType_mock { return m.Moq }

func (m *MoqRowsAffected_genType_mock) LastInsertId() (result1 int64, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqRowsAffected_genType_LastInsertId_params{}
	var results *MoqRowsAffected_genType_LastInsertId_results
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

func (m *MoqRowsAffected_genType_mock) RowsAffected() (result1 int64, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqRowsAffected_genType_RowsAffected_params{}
	var results *MoqRowsAffected_genType_RowsAffected_results
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

// OnCall returns the recorder implementation of the RowsAffected_genType type
func (m *MoqRowsAffected_genType) OnCall() *MoqRowsAffected_genType_recorder {
	return &MoqRowsAffected_genType_recorder{
		Moq: m,
	}
}

func (m *MoqRowsAffected_genType_recorder) LastInsertId() *MoqRowsAffected_genType_LastInsertId_fnRecorder {
	return &MoqRowsAffected_genType_LastInsertId_fnRecorder{
		Params:   MoqRowsAffected_genType_LastInsertId_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqRowsAffected_genType_LastInsertId_fnRecorder) Any() *MoqRowsAffected_genType_LastInsertId_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_LastInsertId(r.Params))
		return nil
	}
	return &MoqRowsAffected_genType_LastInsertId_anyParams{Recorder: r}
}

func (r *MoqRowsAffected_genType_LastInsertId_fnRecorder) Seq() *MoqRowsAffected_genType_LastInsertId_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_LastInsertId(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRowsAffected_genType_LastInsertId_fnRecorder) NoSeq() *MoqRowsAffected_genType_LastInsertId_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_LastInsertId(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRowsAffected_genType_LastInsertId_fnRecorder) ReturnResults(result1 int64, result2 error) *MoqRowsAffected_genType_LastInsertId_fnRecorder {
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
		DoFn       MoqRowsAffected_genType_LastInsertId_doFn
		DoReturnFn MoqRowsAffected_genType_LastInsertId_doReturnFn
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

func (r *MoqRowsAffected_genType_LastInsertId_fnRecorder) AndDo(fn MoqRowsAffected_genType_LastInsertId_doFn) *MoqRowsAffected_genType_LastInsertId_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRowsAffected_genType_LastInsertId_fnRecorder) DoReturnResults(fn MoqRowsAffected_genType_LastInsertId_doReturnFn) *MoqRowsAffected_genType_LastInsertId_fnRecorder {
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
		DoFn       MoqRowsAffected_genType_LastInsertId_doFn
		DoReturnFn MoqRowsAffected_genType_LastInsertId_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRowsAffected_genType_LastInsertId_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRowsAffected_genType_LastInsertId_resultsByParams
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
		results = &MoqRowsAffected_genType_LastInsertId_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRowsAffected_genType_LastInsertId_paramsKey]*MoqRowsAffected_genType_LastInsertId_results{},
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
		r.Results = &MoqRowsAffected_genType_LastInsertId_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRowsAffected_genType_LastInsertId_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRowsAffected_genType_LastInsertId_fnRecorder {
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
				DoFn       MoqRowsAffected_genType_LastInsertId_doFn
				DoReturnFn MoqRowsAffected_genType_LastInsertId_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRowsAffected_genType) PrettyParams_LastInsertId(params MoqRowsAffected_genType_LastInsertId_params) string {
	return fmt.Sprintf("LastInsertId()")
}

func (m *MoqRowsAffected_genType) ParamsKey_LastInsertId(params MoqRowsAffected_genType_LastInsertId_params, anyParams uint64) MoqRowsAffected_genType_LastInsertId_paramsKey {
	m.Scene.T.Helper()
	return MoqRowsAffected_genType_LastInsertId_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqRowsAffected_genType_recorder) RowsAffected() *MoqRowsAffected_genType_RowsAffected_fnRecorder {
	return &MoqRowsAffected_genType_RowsAffected_fnRecorder{
		Params:   MoqRowsAffected_genType_RowsAffected_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqRowsAffected_genType_RowsAffected_fnRecorder) Any() *MoqRowsAffected_genType_RowsAffected_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RowsAffected(r.Params))
		return nil
	}
	return &MoqRowsAffected_genType_RowsAffected_anyParams{Recorder: r}
}

func (r *MoqRowsAffected_genType_RowsAffected_fnRecorder) Seq() *MoqRowsAffected_genType_RowsAffected_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RowsAffected(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqRowsAffected_genType_RowsAffected_fnRecorder) NoSeq() *MoqRowsAffected_genType_RowsAffected_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_RowsAffected(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqRowsAffected_genType_RowsAffected_fnRecorder) ReturnResults(result1 int64, result2 error) *MoqRowsAffected_genType_RowsAffected_fnRecorder {
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
		DoFn       MoqRowsAffected_genType_RowsAffected_doFn
		DoReturnFn MoqRowsAffected_genType_RowsAffected_doReturnFn
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

func (r *MoqRowsAffected_genType_RowsAffected_fnRecorder) AndDo(fn MoqRowsAffected_genType_RowsAffected_doFn) *MoqRowsAffected_genType_RowsAffected_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqRowsAffected_genType_RowsAffected_fnRecorder) DoReturnResults(fn MoqRowsAffected_genType_RowsAffected_doReturnFn) *MoqRowsAffected_genType_RowsAffected_fnRecorder {
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
		DoFn       MoqRowsAffected_genType_RowsAffected_doFn
		DoReturnFn MoqRowsAffected_genType_RowsAffected_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqRowsAffected_genType_RowsAffected_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqRowsAffected_genType_RowsAffected_resultsByParams
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
		results = &MoqRowsAffected_genType_RowsAffected_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqRowsAffected_genType_RowsAffected_paramsKey]*MoqRowsAffected_genType_RowsAffected_results{},
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
		r.Results = &MoqRowsAffected_genType_RowsAffected_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqRowsAffected_genType_RowsAffected_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqRowsAffected_genType_RowsAffected_fnRecorder {
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
				DoFn       MoqRowsAffected_genType_RowsAffected_doFn
				DoReturnFn MoqRowsAffected_genType_RowsAffected_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqRowsAffected_genType) PrettyParams_RowsAffected(params MoqRowsAffected_genType_RowsAffected_params) string {
	return fmt.Sprintf("RowsAffected()")
}

func (m *MoqRowsAffected_genType) ParamsKey_RowsAffected(params MoqRowsAffected_genType_RowsAffected_params, anyParams uint64) MoqRowsAffected_genType_RowsAffected_paramsKey {
	m.Scene.T.Helper()
	return MoqRowsAffected_genType_RowsAffected_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqRowsAffected_genType) Reset() {
	m.ResultsByParams_LastInsertId = nil
	m.ResultsByParams_RowsAffected = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqRowsAffected_genType) AssertExpectationsMet() {
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