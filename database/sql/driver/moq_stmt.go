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

// The following type assertion assures that driver.Stmt is mocked completely
var _ driver.Stmt = (*MoqStmt_mock)(nil)

// MoqStmt holds the state of a moq of the Stmt type
type MoqStmt struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqStmt_mock

	ResultsByParams_Close    []MoqStmt_Close_resultsByParams
	ResultsByParams_NumInput []MoqStmt_NumInput_resultsByParams
	ResultsByParams_Exec     []MoqStmt_Exec_resultsByParams
	ResultsByParams_Query    []MoqStmt_Query_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Close    struct{}
			NumInput struct{}
			Exec     struct {
				Args moq.ParamIndexing
			}
			Query struct {
				Args moq.ParamIndexing
			}
		}
	}
	// MoqStmt_mock isolates the mock interface of the Stmt type
}

type MoqStmt_mock struct {
	Moq *MoqStmt
}

// MoqStmt_recorder isolates the recorder interface of the Stmt type
type MoqStmt_recorder struct {
	Moq *MoqStmt
}

// MoqStmt_Close_params holds the params of the Stmt type
type MoqStmt_Close_params struct{}

// MoqStmt_Close_paramsKey holds the map key params of the Stmt type
type MoqStmt_Close_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqStmt_Close_resultsByParams contains the results for a given set of
// parameters for the Stmt type
type MoqStmt_Close_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStmt_Close_paramsKey]*MoqStmt_Close_results
}

// MoqStmt_Close_doFn defines the type of function needed when calling AndDo
// for the Stmt type
type MoqStmt_Close_doFn func()

// MoqStmt_Close_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Stmt type
type MoqStmt_Close_doReturnFn func() error

// MoqStmt_Close_results holds the results of the Stmt type
type MoqStmt_Close_results struct {
	Params  MoqStmt_Close_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqStmt_Close_doFn
		DoReturnFn MoqStmt_Close_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStmt_Close_fnRecorder routes recorded function calls to the MoqStmt moq
type MoqStmt_Close_fnRecorder struct {
	Params    MoqStmt_Close_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStmt_Close_results
	Moq       *MoqStmt
}

// MoqStmt_Close_anyParams isolates the any params functions of the Stmt type
type MoqStmt_Close_anyParams struct {
	Recorder *MoqStmt_Close_fnRecorder
}

// MoqStmt_NumInput_params holds the params of the Stmt type
type MoqStmt_NumInput_params struct{}

// MoqStmt_NumInput_paramsKey holds the map key params of the Stmt type
type MoqStmt_NumInput_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqStmt_NumInput_resultsByParams contains the results for a given set of
// parameters for the Stmt type
type MoqStmt_NumInput_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStmt_NumInput_paramsKey]*MoqStmt_NumInput_results
}

// MoqStmt_NumInput_doFn defines the type of function needed when calling AndDo
// for the Stmt type
type MoqStmt_NumInput_doFn func()

// MoqStmt_NumInput_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Stmt type
type MoqStmt_NumInput_doReturnFn func() int

// MoqStmt_NumInput_results holds the results of the Stmt type
type MoqStmt_NumInput_results struct {
	Params  MoqStmt_NumInput_params
	Results []struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqStmt_NumInput_doFn
		DoReturnFn MoqStmt_NumInput_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStmt_NumInput_fnRecorder routes recorded function calls to the MoqStmt
// moq
type MoqStmt_NumInput_fnRecorder struct {
	Params    MoqStmt_NumInput_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStmt_NumInput_results
	Moq       *MoqStmt
}

// MoqStmt_NumInput_anyParams isolates the any params functions of the Stmt
// type
type MoqStmt_NumInput_anyParams struct {
	Recorder *MoqStmt_NumInput_fnRecorder
}

// MoqStmt_Exec_params holds the params of the Stmt type
type MoqStmt_Exec_params struct{ Args []driver.Value }

// MoqStmt_Exec_paramsKey holds the map key params of the Stmt type
type MoqStmt_Exec_paramsKey struct {
	Params struct{}
	Hashes struct{ Args hash.Hash }
}

// MoqStmt_Exec_resultsByParams contains the results for a given set of
// parameters for the Stmt type
type MoqStmt_Exec_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStmt_Exec_paramsKey]*MoqStmt_Exec_results
}

// MoqStmt_Exec_doFn defines the type of function needed when calling AndDo for
// the Stmt type
type MoqStmt_Exec_doFn func(args []driver.Value)

// MoqStmt_Exec_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Stmt type
type MoqStmt_Exec_doReturnFn func(args []driver.Value) (driver.Result, error)

// MoqStmt_Exec_results holds the results of the Stmt type
type MoqStmt_Exec_results struct {
	Params  MoqStmt_Exec_params
	Results []struct {
		Values *struct {
			Result1 driver.Result
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqStmt_Exec_doFn
		DoReturnFn MoqStmt_Exec_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStmt_Exec_fnRecorder routes recorded function calls to the MoqStmt moq
type MoqStmt_Exec_fnRecorder struct {
	Params    MoqStmt_Exec_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStmt_Exec_results
	Moq       *MoqStmt
}

// MoqStmt_Exec_anyParams isolates the any params functions of the Stmt type
type MoqStmt_Exec_anyParams struct {
	Recorder *MoqStmt_Exec_fnRecorder
}

// MoqStmt_Query_params holds the params of the Stmt type
type MoqStmt_Query_params struct{ Args []driver.Value }

// MoqStmt_Query_paramsKey holds the map key params of the Stmt type
type MoqStmt_Query_paramsKey struct {
	Params struct{}
	Hashes struct{ Args hash.Hash }
}

// MoqStmt_Query_resultsByParams contains the results for a given set of
// parameters for the Stmt type
type MoqStmt_Query_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqStmt_Query_paramsKey]*MoqStmt_Query_results
}

// MoqStmt_Query_doFn defines the type of function needed when calling AndDo
// for the Stmt type
type MoqStmt_Query_doFn func(args []driver.Value)

// MoqStmt_Query_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Stmt type
type MoqStmt_Query_doReturnFn func(args []driver.Value) (driver.Rows, error)

// MoqStmt_Query_results holds the results of the Stmt type
type MoqStmt_Query_results struct {
	Params  MoqStmt_Query_params
	Results []struct {
		Values *struct {
			Result1 driver.Rows
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqStmt_Query_doFn
		DoReturnFn MoqStmt_Query_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqStmt_Query_fnRecorder routes recorded function calls to the MoqStmt moq
type MoqStmt_Query_fnRecorder struct {
	Params    MoqStmt_Query_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqStmt_Query_results
	Moq       *MoqStmt
}

// MoqStmt_Query_anyParams isolates the any params functions of the Stmt type
type MoqStmt_Query_anyParams struct {
	Recorder *MoqStmt_Query_fnRecorder
}

// NewMoqStmt creates a new moq of the Stmt type
func NewMoqStmt(scene *moq.Scene, config *moq.Config) *MoqStmt {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqStmt{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqStmt_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Close    struct{}
				NumInput struct{}
				Exec     struct {
					Args moq.ParamIndexing
				}
				Query struct {
					Args moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Close    struct{}
			NumInput struct{}
			Exec     struct {
				Args moq.ParamIndexing
			}
			Query struct {
				Args moq.ParamIndexing
			}
		}{
			Close:    struct{}{},
			NumInput: struct{}{},
			Exec: struct {
				Args moq.ParamIndexing
			}{
				Args: moq.ParamIndexByHash,
			},
			Query: struct {
				Args moq.ParamIndexing
			}{
				Args: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Stmt type
func (m *MoqStmt) Mock() *MoqStmt_mock { return m.Moq }

func (m *MoqStmt_mock) Close() (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqStmt_Close_params{}
	var results *MoqStmt_Close_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Close {
		paramsKey := m.Moq.ParamsKey_Close(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Close(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Close(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Close(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqStmt_mock) NumInput() (result1 int) {
	m.Moq.Scene.T.Helper()
	params := MoqStmt_NumInput_params{}
	var results *MoqStmt_NumInput_results
	for _, resultsByParams := range m.Moq.ResultsByParams_NumInput {
		paramsKey := m.Moq.ParamsKey_NumInput(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_NumInput(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_NumInput(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_NumInput(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn()
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn()
	}
	return
}

func (m *MoqStmt_mock) Exec(args []driver.Value) (result1 driver.Result, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqStmt_Exec_params{
		Args: args,
	}
	var results *MoqStmt_Exec_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Exec {
		paramsKey := m.Moq.ParamsKey_Exec(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Exec(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Exec(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Exec(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(args)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(args)
	}
	return
}

func (m *MoqStmt_mock) Query(args []driver.Value) (result1 driver.Rows, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqStmt_Query_params{
		Args: args,
	}
	var results *MoqStmt_Query_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Query {
		paramsKey := m.Moq.ParamsKey_Query(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Query(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Query(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Query(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(args)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(args)
	}
	return
}

// OnCall returns the recorder implementation of the Stmt type
func (m *MoqStmt) OnCall() *MoqStmt_recorder {
	return &MoqStmt_recorder{
		Moq: m,
	}
}

func (m *MoqStmt_recorder) Close() *MoqStmt_Close_fnRecorder {
	return &MoqStmt_Close_fnRecorder{
		Params:   MoqStmt_Close_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqStmt_Close_fnRecorder) Any() *MoqStmt_Close_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	return &MoqStmt_Close_anyParams{Recorder: r}
}

func (r *MoqStmt_Close_fnRecorder) Seq() *MoqStmt_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStmt_Close_fnRecorder) NoSeq() *MoqStmt_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Close(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStmt_Close_fnRecorder) ReturnResults(result1 error) *MoqStmt_Close_fnRecorder {
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
		DoFn       MoqStmt_Close_doFn
		DoReturnFn MoqStmt_Close_doReturnFn
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

func (r *MoqStmt_Close_fnRecorder) AndDo(fn MoqStmt_Close_doFn) *MoqStmt_Close_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStmt_Close_fnRecorder) DoReturnResults(fn MoqStmt_Close_doReturnFn) *MoqStmt_Close_fnRecorder {
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
		DoFn       MoqStmt_Close_doFn
		DoReturnFn MoqStmt_Close_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStmt_Close_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStmt_Close_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Close {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqStmt_Close_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStmt_Close_paramsKey]*MoqStmt_Close_results{},
		}
		r.Moq.ResultsByParams_Close = append(r.Moq.ResultsByParams_Close, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Close) {
			copy(r.Moq.ResultsByParams_Close[insertAt+1:], r.Moq.ResultsByParams_Close[insertAt:0])
			r.Moq.ResultsByParams_Close[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Close(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqStmt_Close_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStmt_Close_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStmt_Close_fnRecorder {
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
				DoFn       MoqStmt_Close_doFn
				DoReturnFn MoqStmt_Close_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStmt) PrettyParams_Close(params MoqStmt_Close_params) string {
	return fmt.Sprintf("Close()")
}

func (m *MoqStmt) ParamsKey_Close(params MoqStmt_Close_params, anyParams uint64) MoqStmt_Close_paramsKey {
	m.Scene.T.Helper()
	return MoqStmt_Close_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqStmt_recorder) NumInput() *MoqStmt_NumInput_fnRecorder {
	return &MoqStmt_NumInput_fnRecorder{
		Params:   MoqStmt_NumInput_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqStmt_NumInput_fnRecorder) Any() *MoqStmt_NumInput_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_NumInput(r.Params))
		return nil
	}
	return &MoqStmt_NumInput_anyParams{Recorder: r}
}

func (r *MoqStmt_NumInput_fnRecorder) Seq() *MoqStmt_NumInput_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_NumInput(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStmt_NumInput_fnRecorder) NoSeq() *MoqStmt_NumInput_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_NumInput(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStmt_NumInput_fnRecorder) ReturnResults(result1 int) *MoqStmt_NumInput_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqStmt_NumInput_doFn
		DoReturnFn MoqStmt_NumInput_doReturnFn
	}{
		Values: &struct {
			Result1 int
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStmt_NumInput_fnRecorder) AndDo(fn MoqStmt_NumInput_doFn) *MoqStmt_NumInput_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStmt_NumInput_fnRecorder) DoReturnResults(fn MoqStmt_NumInput_doReturnFn) *MoqStmt_NumInput_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 int
		}
		Sequence   uint32
		DoFn       MoqStmt_NumInput_doFn
		DoReturnFn MoqStmt_NumInput_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStmt_NumInput_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStmt_NumInput_resultsByParams
	for n, res := range r.Moq.ResultsByParams_NumInput {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqStmt_NumInput_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStmt_NumInput_paramsKey]*MoqStmt_NumInput_results{},
		}
		r.Moq.ResultsByParams_NumInput = append(r.Moq.ResultsByParams_NumInput, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_NumInput) {
			copy(r.Moq.ResultsByParams_NumInput[insertAt+1:], r.Moq.ResultsByParams_NumInput[insertAt:0])
			r.Moq.ResultsByParams_NumInput[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_NumInput(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqStmt_NumInput_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStmt_NumInput_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStmt_NumInput_fnRecorder {
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
					Result1 int
				}
				Sequence   uint32
				DoFn       MoqStmt_NumInput_doFn
				DoReturnFn MoqStmt_NumInput_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStmt) PrettyParams_NumInput(params MoqStmt_NumInput_params) string {
	return fmt.Sprintf("NumInput()")
}

func (m *MoqStmt) ParamsKey_NumInput(params MoqStmt_NumInput_params, anyParams uint64) MoqStmt_NumInput_paramsKey {
	m.Scene.T.Helper()
	return MoqStmt_NumInput_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqStmt_recorder) Exec(args []driver.Value) *MoqStmt_Exec_fnRecorder {
	return &MoqStmt_Exec_fnRecorder{
		Params: MoqStmt_Exec_params{
			Args: args,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqStmt_Exec_fnRecorder) Any() *MoqStmt_Exec_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Exec(r.Params))
		return nil
	}
	return &MoqStmt_Exec_anyParams{Recorder: r}
}

func (a *MoqStmt_Exec_anyParams) Args() *MoqStmt_Exec_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqStmt_Exec_fnRecorder) Seq() *MoqStmt_Exec_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Exec(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStmt_Exec_fnRecorder) NoSeq() *MoqStmt_Exec_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Exec(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStmt_Exec_fnRecorder) ReturnResults(result1 driver.Result, result2 error) *MoqStmt_Exec_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Result
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqStmt_Exec_doFn
		DoReturnFn MoqStmt_Exec_doReturnFn
	}{
		Values: &struct {
			Result1 driver.Result
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStmt_Exec_fnRecorder) AndDo(fn MoqStmt_Exec_doFn) *MoqStmt_Exec_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStmt_Exec_fnRecorder) DoReturnResults(fn MoqStmt_Exec_doReturnFn) *MoqStmt_Exec_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Result
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqStmt_Exec_doFn
		DoReturnFn MoqStmt_Exec_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStmt_Exec_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStmt_Exec_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Exec {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqStmt_Exec_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStmt_Exec_paramsKey]*MoqStmt_Exec_results{},
		}
		r.Moq.ResultsByParams_Exec = append(r.Moq.ResultsByParams_Exec, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Exec) {
			copy(r.Moq.ResultsByParams_Exec[insertAt+1:], r.Moq.ResultsByParams_Exec[insertAt:0])
			r.Moq.ResultsByParams_Exec[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Exec(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqStmt_Exec_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStmt_Exec_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStmt_Exec_fnRecorder {
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
					Result1 driver.Result
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqStmt_Exec_doFn
				DoReturnFn MoqStmt_Exec_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStmt) PrettyParams_Exec(params MoqStmt_Exec_params) string {
	return fmt.Sprintf("Exec(%#v)", params.Args)
}

func (m *MoqStmt) ParamsKey_Exec(params MoqStmt_Exec_params, anyParams uint64) MoqStmt_Exec_paramsKey {
	m.Scene.T.Helper()
	var argsUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Exec.Args == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The args parameter of the Exec function can't be indexed by value")
		}
		argsUsedHash = hash.DeepHash(params.Args)
	}
	return MoqStmt_Exec_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Args hash.Hash }{
			Args: argsUsedHash,
		},
	}
}

func (m *MoqStmt_recorder) Query(args []driver.Value) *MoqStmt_Query_fnRecorder {
	return &MoqStmt_Query_fnRecorder{
		Params: MoqStmt_Query_params{
			Args: args,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqStmt_Query_fnRecorder) Any() *MoqStmt_Query_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Query(r.Params))
		return nil
	}
	return &MoqStmt_Query_anyParams{Recorder: r}
}

func (a *MoqStmt_Query_anyParams) Args() *MoqStmt_Query_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqStmt_Query_fnRecorder) Seq() *MoqStmt_Query_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Query(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqStmt_Query_fnRecorder) NoSeq() *MoqStmt_Query_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Query(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqStmt_Query_fnRecorder) ReturnResults(result1 driver.Rows, result2 error) *MoqStmt_Query_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Rows
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqStmt_Query_doFn
		DoReturnFn MoqStmt_Query_doReturnFn
	}{
		Values: &struct {
			Result1 driver.Rows
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqStmt_Query_fnRecorder) AndDo(fn MoqStmt_Query_doFn) *MoqStmt_Query_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqStmt_Query_fnRecorder) DoReturnResults(fn MoqStmt_Query_doReturnFn) *MoqStmt_Query_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 driver.Rows
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqStmt_Query_doFn
		DoReturnFn MoqStmt_Query_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqStmt_Query_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqStmt_Query_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Query {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqStmt_Query_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqStmt_Query_paramsKey]*MoqStmt_Query_results{},
		}
		r.Moq.ResultsByParams_Query = append(r.Moq.ResultsByParams_Query, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Query) {
			copy(r.Moq.ResultsByParams_Query[insertAt+1:], r.Moq.ResultsByParams_Query[insertAt:0])
			r.Moq.ResultsByParams_Query[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Query(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqStmt_Query_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqStmt_Query_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqStmt_Query_fnRecorder {
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
					Result1 driver.Rows
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqStmt_Query_doFn
				DoReturnFn MoqStmt_Query_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqStmt) PrettyParams_Query(params MoqStmt_Query_params) string {
	return fmt.Sprintf("Query(%#v)", params.Args)
}

func (m *MoqStmt) ParamsKey_Query(params MoqStmt_Query_params, anyParams uint64) MoqStmt_Query_paramsKey {
	m.Scene.T.Helper()
	var argsUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Query.Args == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The args parameter of the Query function can't be indexed by value")
		}
		argsUsedHash = hash.DeepHash(params.Args)
	}
	return MoqStmt_Query_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Args hash.Hash }{
			Args: argsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqStmt) Reset() {
	m.ResultsByParams_Close = nil
	m.ResultsByParams_NumInput = nil
	m.ResultsByParams_Exec = nil
	m.ResultsByParams_Query = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqStmt) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Close {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Close(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_NumInput {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_NumInput(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Exec {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Exec(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Query {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Query(results.Params))
			}
		}
	}
}
