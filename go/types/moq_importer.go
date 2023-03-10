// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package types

import (
	"fmt"
	"go/types"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that types.Importer is mocked
// completely
var _ types.Importer = (*MoqImporter_mock)(nil)

// MoqImporter holds the state of a moq of the Importer type
type MoqImporter struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqImporter_mock

	ResultsByParams_Import []MoqImporter_Import_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Import struct {
				Path moq.ParamIndexing
			}
		}
	}
	// MoqImporter_mock isolates the mock interface of the Importer type
}

type MoqImporter_mock struct {
	Moq *MoqImporter
}

// MoqImporter_recorder isolates the recorder interface of the Importer type
type MoqImporter_recorder struct {
	Moq *MoqImporter
}

// MoqImporter_Import_params holds the params of the Importer type
type MoqImporter_Import_params struct{ Path string }

// MoqImporter_Import_paramsKey holds the map key params of the Importer type
type MoqImporter_Import_paramsKey struct {
	Params struct{ Path string }
	Hashes struct{ Path hash.Hash }
}

// MoqImporter_Import_resultsByParams contains the results for a given set of
// parameters for the Importer type
type MoqImporter_Import_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqImporter_Import_paramsKey]*MoqImporter_Import_results
}

// MoqImporter_Import_doFn defines the type of function needed when calling
// AndDo for the Importer type
type MoqImporter_Import_doFn func(path string)

// MoqImporter_Import_doReturnFn defines the type of function needed when
// calling DoReturnResults for the Importer type
type MoqImporter_Import_doReturnFn func(path string) (*types.Package, error)

// MoqImporter_Import_results holds the results of the Importer type
type MoqImporter_Import_results struct {
	Params  MoqImporter_Import_params
	Results []struct {
		Values *struct {
			Result1 *types.Package
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqImporter_Import_doFn
		DoReturnFn MoqImporter_Import_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqImporter_Import_fnRecorder routes recorded function calls to the
// MoqImporter moq
type MoqImporter_Import_fnRecorder struct {
	Params    MoqImporter_Import_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqImporter_Import_results
	Moq       *MoqImporter
}

// MoqImporter_Import_anyParams isolates the any params functions of the
// Importer type
type MoqImporter_Import_anyParams struct {
	Recorder *MoqImporter_Import_fnRecorder
}

// NewMoqImporter creates a new moq of the Importer type
func NewMoqImporter(scene *moq.Scene, config *moq.Config) *MoqImporter {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqImporter{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqImporter_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Import struct {
					Path moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Import struct {
				Path moq.ParamIndexing
			}
		}{
			Import: struct {
				Path moq.ParamIndexing
			}{
				Path: moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Importer type
func (m *MoqImporter) Mock() *MoqImporter_mock { return m.Moq }

func (m *MoqImporter_mock) Import(path string) (result1 *types.Package, result2 error) {
	m.Moq.Scene.T.Helper()
	params := MoqImporter_Import_params{
		Path: path,
	}
	var results *MoqImporter_Import_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Import {
		paramsKey := m.Moq.ParamsKey_Import(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Import(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Import(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Import(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(path)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
		result2 = result.Values.Result2
	}
	if result.DoReturnFn != nil {
		result1, result2 = result.DoReturnFn(path)
	}
	return
}

// OnCall returns the recorder implementation of the Importer type
func (m *MoqImporter) OnCall() *MoqImporter_recorder {
	return &MoqImporter_recorder{
		Moq: m,
	}
}

func (m *MoqImporter_recorder) Import(path string) *MoqImporter_Import_fnRecorder {
	return &MoqImporter_Import_fnRecorder{
		Params: MoqImporter_Import_params{
			Path: path,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqImporter_Import_fnRecorder) Any() *MoqImporter_Import_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Import(r.Params))
		return nil
	}
	return &MoqImporter_Import_anyParams{Recorder: r}
}

func (a *MoqImporter_Import_anyParams) Path() *MoqImporter_Import_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (r *MoqImporter_Import_fnRecorder) Seq() *MoqImporter_Import_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Import(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqImporter_Import_fnRecorder) NoSeq() *MoqImporter_Import_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Import(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqImporter_Import_fnRecorder) ReturnResults(result1 *types.Package, result2 error) *MoqImporter_Import_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Package
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqImporter_Import_doFn
		DoReturnFn MoqImporter_Import_doReturnFn
	}{
		Values: &struct {
			Result1 *types.Package
			Result2 error
		}{
			Result1: result1,
			Result2: result2,
		},
		Sequence: sequence,
	})
	return r
}

func (r *MoqImporter_Import_fnRecorder) AndDo(fn MoqImporter_Import_doFn) *MoqImporter_Import_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqImporter_Import_fnRecorder) DoReturnResults(fn MoqImporter_Import_doReturnFn) *MoqImporter_Import_fnRecorder {
	r.Moq.Scene.T.Helper()
	r.FindResults()

	var sequence uint32
	if r.Sequence {
		sequence = r.Moq.Scene.NextRecorderSequence()
	}

	r.Results.Results = append(r.Results.Results, struct {
		Values *struct {
			Result1 *types.Package
			Result2 error
		}
		Sequence   uint32
		DoFn       MoqImporter_Import_doFn
		DoReturnFn MoqImporter_Import_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqImporter_Import_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqImporter_Import_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Import {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqImporter_Import_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqImporter_Import_paramsKey]*MoqImporter_Import_results{},
		}
		r.Moq.ResultsByParams_Import = append(r.Moq.ResultsByParams_Import, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Import) {
			copy(r.Moq.ResultsByParams_Import[insertAt+1:], r.Moq.ResultsByParams_Import[insertAt:0])
			r.Moq.ResultsByParams_Import[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Import(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqImporter_Import_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqImporter_Import_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqImporter_Import_fnRecorder {
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
					Result2 error
				}
				Sequence   uint32
				DoFn       MoqImporter_Import_doFn
				DoReturnFn MoqImporter_Import_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqImporter) PrettyParams_Import(params MoqImporter_Import_params) string {
	return fmt.Sprintf("Import(%#v)", params.Path)
}

func (m *MoqImporter) ParamsKey_Import(params MoqImporter_Import_params, anyParams uint64) MoqImporter_Import_paramsKey {
	m.Scene.T.Helper()
	var pathUsed string
	var pathUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Import.Path == moq.ParamIndexByValue {
			pathUsed = params.Path
		} else {
			pathUsedHash = hash.DeepHash(params.Path)
		}
	}
	return MoqImporter_Import_paramsKey{
		Params: struct{ Path string }{
			Path: pathUsed,
		},
		Hashes: struct{ Path hash.Hash }{
			Path: pathUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqImporter) Reset() { m.ResultsByParams_Import = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqImporter) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Import {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Import(results.Params))
			}
		}
	}
}
