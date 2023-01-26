// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package types

import (
	"fmt"
	"go/types"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/moq"
)

// The following type assertion assures that types.PkgName_starGenType is
// mocked completely
var _ PkgName_starGenType = (*MoqPkgName_starGenType_mock)(nil)

// PkgName_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type PkgName_starGenType interface {
	Imported() *types.Package
	String() string
}

// MoqPkgName_starGenType holds the state of a moq of the PkgName_starGenType
// type
type MoqPkgName_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqPkgName_starGenType_mock

	ResultsByParams_Imported []MoqPkgName_starGenType_Imported_resultsByParams
	ResultsByParams_String   []MoqPkgName_starGenType_String_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Imported struct{}
			String   struct{}
		}
	}
}

// MoqPkgName_starGenType_mock isolates the mock interface of the
// PkgName_starGenType type
type MoqPkgName_starGenType_mock struct {
	Moq *MoqPkgName_starGenType
}

// MoqPkgName_starGenType_recorder isolates the recorder interface of the
// PkgName_starGenType type
type MoqPkgName_starGenType_recorder struct {
	Moq *MoqPkgName_starGenType
}

// MoqPkgName_starGenType_Imported_params holds the params of the
// PkgName_starGenType type
type MoqPkgName_starGenType_Imported_params struct{}

// MoqPkgName_starGenType_Imported_paramsKey holds the map key params of the
// PkgName_starGenType type
type MoqPkgName_starGenType_Imported_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqPkgName_starGenType_Imported_resultsByParams contains the results for a
// given set of parameters for the PkgName_starGenType type
type MoqPkgName_starGenType_Imported_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPkgName_starGenType_Imported_paramsKey]*MoqPkgName_starGenType_Imported_results
}

// MoqPkgName_starGenType_Imported_doFn defines the type of function needed
// when calling AndDo for the PkgName_starGenType type
type MoqPkgName_starGenType_Imported_doFn func()

// MoqPkgName_starGenType_Imported_doReturnFn defines the type of function
// needed when calling DoReturnResults for the PkgName_starGenType type
type MoqPkgName_starGenType_Imported_doReturnFn func() *types.Package

// MoqPkgName_starGenType_Imported_results holds the results of the
// PkgName_starGenType type
type MoqPkgName_starGenType_Imported_results struct {
	Params  MoqPkgName_starGenType_Imported_params
	Results []struct {
		Values *struct {
			Result1 *types.Package
		}
		Sequence   uint32
		DoFn       MoqPkgName_starGenType_Imported_doFn
		DoReturnFn MoqPkgName_starGenType_Imported_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPkgName_starGenType_Imported_fnRecorder routes recorded function calls to
// the MoqPkgName_starGenType moq
type MoqPkgName_starGenType_Imported_fnRecorder struct {
	Params    MoqPkgName_starGenType_Imported_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPkgName_starGenType_Imported_results
	Moq       *MoqPkgName_starGenType
}

// MoqPkgName_starGenType_Imported_anyParams isolates the any params functions
// of the PkgName_starGenType type
type MoqPkgName_starGenType_Imported_anyParams struct {
	Recorder *MoqPkgName_starGenType_Imported_fnRecorder
}

// MoqPkgName_starGenType_String_params holds the params of the
// PkgName_starGenType type
type MoqPkgName_starGenType_String_params struct{}

// MoqPkgName_starGenType_String_paramsKey holds the map key params of the
// PkgName_starGenType type
type MoqPkgName_starGenType_String_paramsKey struct {
	Params struct{}
	Hashes struct{}
}

// MoqPkgName_starGenType_String_resultsByParams contains the results for a
// given set of parameters for the PkgName_starGenType type
type MoqPkgName_starGenType_String_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqPkgName_starGenType_String_paramsKey]*MoqPkgName_starGenType_String_results
}

// MoqPkgName_starGenType_String_doFn defines the type of function needed when
// calling AndDo for the PkgName_starGenType type
type MoqPkgName_starGenType_String_doFn func()

// MoqPkgName_starGenType_String_doReturnFn defines the type of function needed
// when calling DoReturnResults for the PkgName_starGenType type
type MoqPkgName_starGenType_String_doReturnFn func() string

// MoqPkgName_starGenType_String_results holds the results of the
// PkgName_starGenType type
type MoqPkgName_starGenType_String_results struct {
	Params  MoqPkgName_starGenType_String_params
	Results []struct {
		Values *struct {
			Result1 string
		}
		Sequence   uint32
		DoFn       MoqPkgName_starGenType_String_doFn
		DoReturnFn MoqPkgName_starGenType_String_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqPkgName_starGenType_String_fnRecorder routes recorded function calls to
// the MoqPkgName_starGenType moq
type MoqPkgName_starGenType_String_fnRecorder struct {
	Params    MoqPkgName_starGenType_String_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqPkgName_starGenType_String_results
	Moq       *MoqPkgName_starGenType
}

// MoqPkgName_starGenType_String_anyParams isolates the any params functions of
// the PkgName_starGenType type
type MoqPkgName_starGenType_String_anyParams struct {
	Recorder *MoqPkgName_starGenType_String_fnRecorder
}

// NewMoqPkgName_starGenType creates a new moq of the PkgName_starGenType type
func NewMoqPkgName_starGenType(scene *moq.Scene, config *moq.Config) *MoqPkgName_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqPkgName_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqPkgName_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Imported struct{}
				String   struct{}
			}
		}{ParameterIndexing: struct {
			Imported struct{}
			String   struct{}
		}{
			Imported: struct{}{},
			String:   struct{}{},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the PkgName_starGenType type
func (m *MoqPkgName_starGenType) Mock() *MoqPkgName_starGenType_mock { return m.Moq }

func (m *MoqPkgName_starGenType_mock) Imported() (result1 *types.Package) {
	m.Moq.Scene.T.Helper()
	params := MoqPkgName_starGenType_Imported_params{}
	var results *MoqPkgName_starGenType_Imported_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Imported {
		paramsKey := m.Moq.ParamsKey_Imported(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Imported(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Imported(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Imported(params))
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

func (m *MoqPkgName_starGenType_mock) String() (result1 string) {
	m.Moq.Scene.T.Helper()
	params := MoqPkgName_starGenType_String_params{}
	var results *MoqPkgName_starGenType_String_results
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

// OnCall returns the recorder implementation of the PkgName_starGenType type
func (m *MoqPkgName_starGenType) OnCall() *MoqPkgName_starGenType_recorder {
	return &MoqPkgName_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqPkgName_starGenType_recorder) Imported() *MoqPkgName_starGenType_Imported_fnRecorder {
	return &MoqPkgName_starGenType_Imported_fnRecorder{
		Params:   MoqPkgName_starGenType_Imported_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqPkgName_starGenType_Imported_fnRecorder) Any() *MoqPkgName_starGenType_Imported_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Imported(r.Params))
		return nil
	}
	return &MoqPkgName_starGenType_Imported_anyParams{Recorder: r}
}

func (r *MoqPkgName_starGenType_Imported_fnRecorder) Seq() *MoqPkgName_starGenType_Imported_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Imported(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPkgName_starGenType_Imported_fnRecorder) NoSeq() *MoqPkgName_starGenType_Imported_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Imported(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPkgName_starGenType_Imported_fnRecorder) ReturnResults(result1 *types.Package) *MoqPkgName_starGenType_Imported_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Package
		}
		Sequence   uint32
		DoFn       MoqPkgName_starGenType_Imported_doFn
		DoReturnFn MoqPkgName_starGenType_Imported_doReturnFn
	}{
		Values: &struct {
			Result1 *types.Package
		}{
			Result1: result1,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqPkgName_starGenType_Imported_fnRecorder) AndDo(fn MoqPkgName_starGenType_Imported_doFn) *MoqPkgName_starGenType_Imported_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPkgName_starGenType_Imported_fnRecorder) DoReturnResults(fn MoqPkgName_starGenType_Imported_doReturnFn) *MoqPkgName_starGenType_Imported_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Package
		}
		Sequence   uint32
		DoFn       MoqPkgName_starGenType_Imported_doFn
		DoReturnFn MoqPkgName_starGenType_Imported_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPkgName_starGenType_Imported_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPkgName_starGenType_Imported_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Imported {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqPkgName_starGenType_Imported_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPkgName_starGenType_Imported_paramsKey]*MoqPkgName_starGenType_Imported_results{},
		}
		r.Moq.ResultsByParams_Imported = append(r.Moq.ResultsByParams_Imported, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Imported) {
			copy(r.Moq.ResultsByParams_Imported[insertAt+1:], r.Moq.ResultsByParams_Imported[insertAt:0])
			r.Moq.ResultsByParams_Imported[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Imported(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqPkgName_starGenType_Imported_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPkgName_starGenType_Imported_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPkgName_starGenType_Imported_fnRecorder {
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
					Result1 *types.Package
				}
				Sequence   uint32
				DoFn       MoqPkgName_starGenType_Imported_doFn
				DoReturnFn MoqPkgName_starGenType_Imported_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPkgName_starGenType) PrettyParams_Imported(params MoqPkgName_starGenType_Imported_params) string {
	return fmt.Sprintf("Imported()")
}

func (m *MoqPkgName_starGenType) ParamsKey_Imported(params MoqPkgName_starGenType_Imported_params, anyParams uint64) MoqPkgName_starGenType_Imported_paramsKey {
	m.Scene.T.Helper()
	return MoqPkgName_starGenType_Imported_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

func (m *MoqPkgName_starGenType_recorder) String() *MoqPkgName_starGenType_String_fnRecorder {
	return &MoqPkgName_starGenType_String_fnRecorder{
		Params:   MoqPkgName_starGenType_String_params{},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqPkgName_starGenType_String_fnRecorder) Any() *MoqPkgName_starGenType_String_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	return &MoqPkgName_starGenType_String_anyParams{Recorder: r}
}

func (r *MoqPkgName_starGenType_String_fnRecorder) Seq() *MoqPkgName_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqPkgName_starGenType_String_fnRecorder) NoSeq() *MoqPkgName_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_String(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqPkgName_starGenType_String_fnRecorder) ReturnResults(result1 string) *MoqPkgName_starGenType_String_fnRecorder {
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
		DoFn       MoqPkgName_starGenType_String_doFn
		DoReturnFn MoqPkgName_starGenType_String_doReturnFn
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

func (r *MoqPkgName_starGenType_String_fnRecorder) AndDo(fn MoqPkgName_starGenType_String_doFn) *MoqPkgName_starGenType_String_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqPkgName_starGenType_String_fnRecorder) DoReturnResults(fn MoqPkgName_starGenType_String_doReturnFn) *MoqPkgName_starGenType_String_fnRecorder {
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
		DoFn       MoqPkgName_starGenType_String_doFn
		DoReturnFn MoqPkgName_starGenType_String_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqPkgName_starGenType_String_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqPkgName_starGenType_String_resultsByParams
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
		results = &MoqPkgName_starGenType_String_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqPkgName_starGenType_String_paramsKey]*MoqPkgName_starGenType_String_results{},
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
		r.Results = &MoqPkgName_starGenType_String_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqPkgName_starGenType_String_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqPkgName_starGenType_String_fnRecorder {
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
				DoFn       MoqPkgName_starGenType_String_doFn
				DoReturnFn MoqPkgName_starGenType_String_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqPkgName_starGenType) PrettyParams_String(params MoqPkgName_starGenType_String_params) string {
	return fmt.Sprintf("String()")
}

func (m *MoqPkgName_starGenType) ParamsKey_String(params MoqPkgName_starGenType_String_params, anyParams uint64) MoqPkgName_starGenType_String_paramsKey {
	m.Scene.T.Helper()
	return MoqPkgName_starGenType_String_paramsKey{
		Params: struct{}{},
		Hashes: struct{}{},
	}
}

// Reset resets the state of the moq
func (m *MoqPkgName_starGenType) Reset() {
	m.ResultsByParams_Imported = nil
	m.ResultsByParams_String = nil
}

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqPkgName_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Imported {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Imported(results.Params))
			}
		}
	}
	for _, res := range m.ResultsByParams_String {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_String(results.Params))
			}
		}
	}
}