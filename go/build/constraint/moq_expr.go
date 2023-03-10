// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package constraint

import (
	"fmt"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that Expr_reduced is mocked completely
var _ Expr_reduced = (*MoqExpr_mock)(nil)

// Expr_reduced is the fabricated implementation type of this mock (emitted
// when the original interface contains non-exported methods)
type Expr_reduced interface {
	// String returns the string form of the expression,
	// using the boolean syntax used in //go:build lines.
	String() string

	// Eval reports whether the expression evaluates to true.
	// It calls ok(tag) as needed to find out whether a given build tag
	// is satisfied by the current build configuration.
	Eval(ok func(tag string) bool) bool
}

// MoqExpr holds the state of a moq of the Expr_reduced type
type MoqExpr struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqExpr_mock

	ResultsByParams_String []MoqExpr_String_resultsByParams
	ResultsByParams_Eval   []MoqExpr_Eval_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			String struct{}
			Eval   struct {
				Ok moq.ParamIndexing
			}
		}
	}
	// MoqExpr_mock isolates the mock interface of the Expr type
}

type MoqExpr_mock struct {
	Moq *MoqExpr
}

// MoqExpr_recorder isolates the recorder interface of the Expr type
type MoqExpr_recorder struct {
	Moq *MoqExpr
}

// MoqExpr_String_params holds the params of the Expr type
type MoqExpr_String_params struct{}

// MoqExpr_String_paramsKey holds the map key params of the Expr type
type MoqExpr_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqExpr_String_resultsByParams contains the results for a given set of
// parameters for the Expr type
type MoqExpr_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqExpr_String_paramsKey]*MoqExpr_String_results
}

// MoqExpr_String_doFn defines the type of function needed when calling AndDo
// for the Expr type
type MoqExpr_String_doFn func()

// MoqExpr_String_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Expr type
type MoqExpr_String_doReturnFn func() string

// MoqExpr_String_results holds the results of the Expr type
type MoqExpr_String_results struct {
	Params  MoqExpr_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqExpr_String_doFn
		DoReturnFn MoqExpr_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqExpr_String_fnRecorder routes recorded function calls to the MoqExpr moq
type MoqExpr_String_fnRecorder struct {
	Params    MoqExpr_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqExpr_String_results
	Moq       *MoqExpr
}

// MoqExpr_String_anyParams isolates the any params functions of the Expr type
type MoqExpr_String_anyParams struct {
	Recorder *MoqExpr_String_fnRecorder
}

// MoqExpr_Eval_params holds the params of the Expr type
type MoqExpr_Eval_params struct{ Ok func(tag string) bool }

// MoqExpr_Eval_paramsKey holds the map key params of the Expr type
type MoqExpr_Eval_paramsKey struct {
	Params struct{}
	Hashes struct{ Ok hash.Hash }
}

// MoqExpr_Eval_resultsByParams contains the results for a given set of
// parameters for the Expr type
type MoqExpr_Eval_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqExpr_Eval_paramsKey]*MoqExpr_Eval_results
}

// MoqExpr_Eval_doFn defines the type of function needed when calling AndDo for
// the Expr type
type MoqExpr_Eval_doFn func(ok func(tag string) bool)

// MoqExpr_Eval_doReturnFn defines the type of function needed when calling
// DoReturnResults for the Expr type
type MoqExpr_Eval_doReturnFn func(ok func(tag string) bool) bool

// MoqExpr_Eval_results holds the results of the Expr type
type MoqExpr_Eval_results struct {
	Params  MoqExpr_Eval_params
	Results []struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqExpr_Eval_doFn
		DoReturnFn MoqExpr_Eval_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqExpr_Eval_fnRecorder routes recorded function calls to the MoqExpr moq
type MoqExpr_Eval_fnRecorder struct {
	Params    MoqExpr_Eval_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqExpr_Eval_results
	Moq       *MoqExpr
}

// MoqExpr_Eval_anyParams isolates the any params functions of the Expr type
type MoqExpr_Eval_anyParams struct {
	Recorder *MoqExpr_Eval_fnRecorder
}

// NewMoqExpr creates a new moq of the Expr type
func NewMoqExpr(scene *moq.Scene, config *moq.Config) *MoqExpr {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqExpr{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqExpr_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				String struct{}
				Eval   struct {
					Ok moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			String struct{}
			Eval   struct {
				Ok moq.ParamIndexing
			}
		}{
			String: struct{}{},
			Eval: struct {
				Ok moq.ParamIndexing
			}{
				Ok: moq.ParamIndexByHash,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Expr type
func (m *MoqExpr) Mock() *MoqExpr_mock { return m.Moq }

func (m *MoqExpr_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqExpr_String_params{}
	var results *MoqExpr_String_results
	for _, resultsByParams := range m.Moq.ResultsByParams_String {
		paramsKey := m.Moq.ParamsKey_String(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_String(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_String(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_String(params))
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

func (m *MoqExpr_mock) Eval(ok func(tag string) bool) (result1 bool) {
	m.Moq.Scene.T.Helper()
	params := MoqExpr_Eval_params{
		Ok: ok,
	}
	var results *MoqExpr_Eval_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Eval {
		paramsKey := m.Moq.ParamsKey_Eval(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Eval(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Eval(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Eval(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(ok)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(ok)
	}
	return
}

// OnCall returns the recorder implementation of the Expr type
func (m *MoqExpr) OnCall() *MoqExpr_recorder {
	return &MoqExpr_recorder{
		Moq: m,
	}
}

func (m *MoqExpr_recorder) String() *MoqExpr_String_fnRecorder {
	return &MoqExpr_String_fnRecorder{
		Params:   MoqExpr_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqExpr_String_fnRecorder) Any() *MoqExpr_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqExpr_String_anyParams{Recorder: r}
}

func (r *MoqExpr_String_fnRecorder) Seq() *MoqExpr_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqExpr_String_fnRecorder) NoSeq() *MoqExpr_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqExpr_String_fnRecorder) ReturnResults(result1 string) *MoqExpr_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqExpr_String_doFn
		DoReturnFn MoqExpr_String_doReturnFn
	}{
		Values: &struct {
			Result1 string
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqExpr_String_fnRecorder) AndDo(fn MoqExpr_String_doFn) *MoqExpr_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqExpr_String_fnRecorder) DoReturnResults(fn MoqExpr_String_doReturnFn) *MoqExpr_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqExpr_String_doFn
		DoReturnFn MoqExpr_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqExpr_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqExpr_String_resultsByParams
	for n, res := range r.Moq.ResultsByParams_String {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqExpr_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqExpr_String_paramsKey]*MoqExpr_String_results{},
		}
		r.Moq.ResultsByParams_String = append(r.Moq.ResultsByParams_String, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_String) {
			copy(r.Moq.ResultsByParams_String[insertAt+1:], r.Moq.ResultsByParams_String[insertAt:0])
			r.Moq.ResultsByParams_String[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_String(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqExpr_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqExpr_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqExpr_String_fnRecorder {
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
					Result1 string
				}
				Sequence   uint32
				DoFn       MoqExpr_String_doFn
				DoReturnFn MoqExpr_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqExpr) PrettyParams_String(params MoqExpr_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqExpr) ParamsKey_String(params MoqExpr_String_params, anyParams uint64) MoqExpr_String_paramsKey {
	m.Scene.T.Helper()
	return MoqExpr_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqExpr_recorder) Eval(ok func(tag string) bool) *MoqExpr_Eval_fnRecorder {
	return &MoqExpr_Eval_fnRecorder{
		Params: MoqExpr_Eval_params{
			Ok: ok,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqExpr_Eval_fnRecorder) Any() *MoqExpr_Eval_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Eval(r.Params))
		return nil
	}
	return &MoqExpr_Eval_anyParams{Recorder: r}
}

func (a *MoqExpr_Eval_anyParams) Ok() *MoqExpr_Eval_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqExpr_Eval_fnRecorder) Seq() *MoqExpr_Eval_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Eval(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqExpr_Eval_fnRecorder) NoSeq() *MoqExpr_Eval_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Eval(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqExpr_Eval_fnRecorder) ReturnResults(result1 bool) *MoqExpr_Eval_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqExpr_Eval_doFn
		DoReturnFn MoqExpr_Eval_doReturnFn
	}{
		Values: &struct {
			Result1 bool
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqExpr_Eval_fnRecorder) AndDo(fn MoqExpr_Eval_doFn) *MoqExpr_Eval_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqExpr_Eval_fnRecorder) DoReturnResults(fn MoqExpr_Eval_doReturnFn) *MoqExpr_Eval_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 bool
		}
		Sequence   uint32
		DoFn       MoqExpr_Eval_doFn
		DoReturnFn MoqExpr_Eval_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqExpr_Eval_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqExpr_Eval_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Eval {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqExpr_Eval_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqExpr_Eval_paramsKey]*MoqExpr_Eval_results{},
		}
		r.Moq.ResultsByParams_Eval = append(r.Moq.ResultsByParams_Eval, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Eval) {
			copy(r.Moq.ResultsByParams_Eval[insertAt+1:], r.Moq.ResultsByParams_Eval[insertAt:0])
			r.Moq.ResultsByParams_Eval[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Eval(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqExpr_Eval_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqExpr_Eval_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqExpr_Eval_fnRecorder {
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
					Result1 bool
				}
				Sequence   uint32
				DoFn       MoqExpr_Eval_doFn
				DoReturnFn MoqExpr_Eval_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqExpr) PrettyParams_Eval(params MoqExpr_Eval_params) string {
	return fmt.Sprintf("Eval(%#v)", moq.FnString(params.Ok))
}

func (m *MoqExpr) ParamsKey_Eval(params MoqExpr_Eval_params, anyParams uint64) MoqExpr_Eval_paramsKey {
	m.Scene.T.Helper()
	var okUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Eval.Ok == moq.ParamIndexByValue {
			m.Scene.T.Fatalf("The ok parameter of the Eval function can't be indexed by value")
		}
		okUsedHash = hash.DeepHash(params.Ok)
	}
	return MoqExpr_Eval_paramsKey{
		Params: struct{}{},
		Hashes: struct{ Ok hash.Hash }{
			Ok: okUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqExpr) Reset() { m.ResultsByParams_String = nil; m.ResultsByParams_Eval = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqExpr) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_String {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_String(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_Eval {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Eval(results.Params))
			}
		}
	}
}
