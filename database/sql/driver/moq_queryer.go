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

// The following type assertion assures that driver.Queryer is mocked
// completely
var _ driver.Queryer = (*MoqQueryer_mock)(nil)

// MoqQueryer holds the state of a moq of the Queryer type
type MoqQueryer struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqQueryer_mock

	ResultsByParams_Query []MoqQueryer_Query_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Query struct {
				Query moq.ParamIndexing
				Args  moq.ParamIndexing
			}
		}
	}
	// MoqQueryer_mock isolates the mock interface of the Queryer type
}

type MoqQueryer_mock struct {
	Moq *MoqQueryer
}

// MoqQueryer_recorder isolates the recorder interface of the Queryer type
type MoqQueryer_recorder struct {
	Moq *MoqQueryer
}

// MoqQueryer_Query_params holds the params of the Queryer type
type MoqQueryer_Query_params struct {
	Query string
	Args  []driver.Value
}

// MoqQueryer_Query_paramsKey holds the map key params of the Queryer type
type MoqQueryer_Query_paramsKey struct {
	Params struct{ Query string }
	Hashes struct {
		Query hash.Hash
		Args  hash.Hash
	}
}

// MoqQueryer_Query_resultsByParams contains the results for a given set of
// parameters for the Queryer type
type MoqQueryer_Query_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqQueryer_Query_paramsKey]*MoqQueryer_Query_results
}

// MoqQueryer_Query_doFn defines the type of function needed when calling AndDo
// for the Queryer type
type MoqQueryer_Query_doFn func(query string, args []driver.Value)

// MoqQueryer_Query_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Queryer type
type MoqQueryer_Query_doReturnFn func(query string, args []driver.Value) (driver.Rows, error)

// MoqQueryer_Query_results holds the results of the Queryer type
type MoqQueryer_Query_results struct {
	Params  MoqQueryer_Query_params
	Results []struct {
		Values *struct {
			Result1 driver.Rows
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqQueryer_Query_doFn
		DoReturnFn MoqQueryer_Query_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqQueryer_Query_fnRecorder routes recorded function calls to the MoqQueryer
// moq
type MoqQueryer_Query_fnRecorder struct {
	Params    MoqQueryer_Query_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqQueryer_Query_results
	Moq       *MoqQueryer
}

// MoqQueryer_Query_anyParams isolates the any params functions of the Queryer
// type
type MoqQueryer_Query_anyParams struct {
	Recorder *MoqQueryer_Query_fnRecorder
}

// NewMoqQueryer creates a new moq of the Queryer type
func NewMoqQueryer(scene *moq.Scene, config *moq.Config) *MoqQueryer {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqQueryer{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqQueryer_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Query struct {
					Query moq.ParamIndexing
					Args  moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Query struct {
				Query moq.ParamIndexing
				Args  moq.ParamIndexing
			}
		}{
			Query: struct {
				Query moq.ParamIndexing
				Args  moq.ParamIndexing
			}{
				Query: moq.ParamIndexByValue,
				Args:  moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Queryer type
func (m *MoqQueryer) Mock() *MoqQueryer_mock { return m.Moq }

func (m *MoqQueryer_mock) Query(query string, args []driver.Value) (result1 driver.Rows, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqQueryer_Query_params{
		Query: query,
		Args:  args,
	}
	var results *MoqQueryer_Query_results
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
		result.DoFn(query, args)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(query, args)
	}
	return
}

// OnCall returns the recorder implementation of the Queryer type
func (m *MoqQueryer) OnCall() *MoqQueryer_recorder {
	return &MoqQueryer_recorder{
		Moq: m,
	}
}

func (m *MoqQueryer_recorder) Query(query string, args []driver.Value) *MoqQueryer_Query_fnRecorder {
	return &MoqQueryer_Query_fnRecorder{
		Params: MoqQueryer_Query_params{
			Query: query,
			Args:  args,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqQueryer_Query_fnRecorder) Any() *MoqQueryer_Query_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Query(r.Params))
		return nil
	}
	return &MoqQueryer_Query_anyParams{Recorder: r}
}

func (a *MoqQueryer_Query_anyParams) Query() *MoqQueryer_Query_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqQueryer_Query_anyParams) Args() *MoqQueryer_Query_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (r *MoqQueryer_Query_fnRecorder) Seq() *MoqQueryer_Query_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Query(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqQueryer_Query_fnRecorder) NoSeq() *MoqQueryer_Query_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Query(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqQueryer_Query_fnRecorder) ReturnResults(result1 driver.Rows, result2 error) *MoqQueryer_Query_fnRecorder {
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
		DoFn       MoqQueryer_Query_doFn
		DoReturnFn MoqQueryer_Query_doReturnFn
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

func (r *MoqQueryer_Query_fnRecorder) AndDo(fn MoqQueryer_Query_doFn) *MoqQueryer_Query_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqQueryer_Query_fnRecorder) DoReturnResults(fn MoqQueryer_Query_doReturnFn) *MoqQueryer_Query_fnRecorder {
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
		DoFn       MoqQueryer_Query_doFn
		DoReturnFn MoqQueryer_Query_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqQueryer_Query_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqQueryer_Query_resultsByParams
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
		results = &MoqQueryer_Query_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqQueryer_Query_paramsKey]*MoqQueryer_Query_results{},
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
		r.Results = &MoqQueryer_Query_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqQueryer_Query_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqQueryer_Query_fnRecorder {
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
				DoFn       MoqQueryer_Query_doFn
				DoReturnFn MoqQueryer_Query_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqQueryer) PrettyParams_Query(params MoqQueryer_Query_params) string {
	return fmt.Sprintf("Query(%#v, %#v)", params.Query, params.Args)
}

func (m *MoqQueryer) ParamsKey_Query(params MoqQueryer_Query_params, anyParams uint64) MoqQueryer_Query_paramsKey {
	m.Scene.T.Helper()
	var queryUsed string
	var queryUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Query.Query == moq.ParamIndexByValue {
			queryUsed = params.Query
		} else {
			queryUsedHash = hash.DeepHash(params.Query)
		}
	}
	var argsUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Query.Args == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The args parameter of the Query function can't be indexed by value")
		}
		argsUsedHash = hash.DeepHash(params.Args)
	}
	return MoqQueryer_Query_paramsKey{
		Params: struct{ Query string }{
			Query: queryUsed,
		},
		Hashes: struct {
			Query hash.Hash
			Args  hash.Hash
		}{
			Query: queryUsedHash,
			Args:  argsUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqQueryer) Reset() { m.ResultsByParams_Query = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqQueryer) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Query {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Query(results.Params))
			}
		}
	}
}
