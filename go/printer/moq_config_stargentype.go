// Code generated by Moqueries - https://moqueries.org - DO NOT EDIT!

package printer

import (
	"fmt"
	"go/token"
	"io"
	"math/bits"
	"sync/atomic"

	"moqueries.org/runtime/hash"
	"moqueries.org/runtime/moq"
)

// The following type assertion assures that printer.Config_starGenType is
// mocked completely
var _ Config_starGenType = (*MoqConfig_starGenType_mock)(nil)

// Config_starGenType is the fabricated implementation type of this mock
// (emitted when mocking a collections of methods directly and not from an
// interface type)
type Config_starGenType interface {
	Fprint(output io.Writer, fset *token.FileSet, node any) error
}

// MoqConfig_starGenType holds the state of a moq of the Config_starGenType
// type
type MoqConfig_starGenType struct {
	Scene  *moq.Scene
	Config moq.Config
	Moq    *MoqConfig_starGenType_mock

	ResultsByParams_Fprint []MoqConfig_starGenType_Fprint_resultsByParams

	Runtime struct {
		ParameterIndexing struct {
			Fprint struct {
				Output moq.ParamIndexing
				Fset   moq.ParamIndexing
				Node   moq.ParamIndexing
			}
		}
	}
	// MoqConfig_starGenType_mock isolates the mock interface of the
}

// Config_starGenType type
type MoqConfig_starGenType_mock struct {
	Moq *MoqConfig_starGenType
}

// MoqConfig_starGenType_recorder isolates the recorder interface of the
// Config_starGenType type
type MoqConfig_starGenType_recorder struct {
	Moq *MoqConfig_starGenType
}

// MoqConfig_starGenType_Fprint_params holds the params of the
// Config_starGenType type
type MoqConfig_starGenType_Fprint_params struct {
	Output io.Writer
	Fset   *token.FileSet
	Node   any
}

// MoqConfig_starGenType_Fprint_paramsKey holds the map key params of the
// Config_starGenType type
type MoqConfig_starGenType_Fprint_paramsKey struct {
	Params struct {
		Output io.Writer
		Fset   *token.FileSet
		Node   any
	}
	Hashes struct {
		Output hash.Hash
		Fset   hash.Hash
		Node   hash.Hash
	}
}

// MoqConfig_starGenType_Fprint_resultsByParams contains the results for a
// given set of parameters for the Config_starGenType type
type MoqConfig_starGenType_Fprint_resultsByParams struct {
	AnyCount  int
	AnyParams uint64
	Results   map[MoqConfig_starGenType_Fprint_paramsKey]*MoqConfig_starGenType_Fprint_results
}

// MoqConfig_starGenType_Fprint_doFn defines the type of function needed when
// calling AndDo for the Config_starGenType type
type MoqConfig_starGenType_Fprint_doFn func(output io.Writer, fset *token.FileSet, node any)

// MoqConfig_starGenType_Fprint_doReturnFn defines the type of function needed
// when calling DoReturnResults for the Config_starGenType type
type MoqConfig_starGenType_Fprint_doReturnFn func(output io.Writer, fset *token.FileSet, node any) error

// MoqConfig_starGenType_Fprint_results holds the results of the
// Config_starGenType type
type MoqConfig_starGenType_Fprint_results struct {
	Params  MoqConfig_starGenType_Fprint_params
	Results []struct {
		Values *struct {
			Result1 error
		}
		Sequence   uint32
		DoFn       MoqConfig_starGenType_Fprint_doFn
		DoReturnFn MoqConfig_starGenType_Fprint_doReturnFn
	}
	Index  uint32
	Repeat *moq.RepeatVal
}

// MoqConfig_starGenType_Fprint_fnRecorder routes recorded function calls to
// the MoqConfig_starGenType moq
type MoqConfig_starGenType_Fprint_fnRecorder struct {
	Params    MoqConfig_starGenType_Fprint_params
	AnyParams uint64
	Sequence  bool
	Results   *MoqConfig_starGenType_Fprint_results
	Moq       *MoqConfig_starGenType
}

// MoqConfig_starGenType_Fprint_anyParams isolates the any params functions of
// the Config_starGenType type
type MoqConfig_starGenType_Fprint_anyParams struct {
	Recorder *MoqConfig_starGenType_Fprint_fnRecorder
}

// NewMoqConfig_starGenType creates a new moq of the Config_starGenType type
func NewMoqConfig_starGenType(scene *moq.Scene, config *moq.Config) *MoqConfig_starGenType {
	if config == nil {
		config = &moq.Config{}
	}
	m := &MoqConfig_starGenType{
		Scene:  scene,
		Config: *config,
		Moq:    &MoqConfig_starGenType_mock{},

		Runtime: struct {
			ParameterIndexing struct {
				Fprint struct {
					Output moq.ParamIndexing
					Fset   moq.ParamIndexing
					Node   moq.ParamIndexing
				}
			}
		}{ParameterIndexing: struct {
			Fprint struct {
				Output moq.ParamIndexing
				Fset   moq.ParamIndexing
				Node   moq.ParamIndexing
			}
		}{
			Fprint: struct {
				Output moq.ParamIndexing
				Fset   moq.ParamIndexing
				Node   moq.ParamIndexing
			}{
				Output: moq.ParamIndexByHash,
				Fset:   moq.ParamIndexByHash,
				Node:   moq.ParamIndexByValue,
			},
		}},
	}
	m.Moq.Moq = m

	scene.AddMoq(m)
	return m
}

// Mock returns the mock implementation of the Config_starGenType type
func (m *MoqConfig_starGenType) Mock() *MoqConfig_starGenType_mock { return m.Moq }

func (m *MoqConfig_starGenType_mock) Fprint(output io.Writer, fset *token.FileSet, node any) (result1 error) {
	m.Moq.Scene.T.Helper()
	params := MoqConfig_starGenType_Fprint_params{
		Output: output,
		Fset:   fset,
		Node:   node,
	}
	var results *MoqConfig_starGenType_Fprint_results
	for _, resultsByParams := range m.Moq.ResultsByParams_Fprint {
		paramsKey := m.Moq.ParamsKey_Fprint(params, resultsByParams.AnyParams)
		var ok bool
		results, ok = resultsByParams.Results[paramsKey]
		if ok {
			break
		}
	}
	if results == nil {
		if m.Moq.Config.Expectation == moq.Strict {
			m.Moq.Scene.T.Fatalf("Unexpected call to %s", m.Moq.PrettyParams_Fprint(params))
		}
		return
	}

	i := int(atomic.AddUint32(&results.Index, 1)) - 1
	if i >= results.Repeat.ResultCount {
		if !results.Repeat.AnyTimes {
			if m.Moq.Config.Expectation == moq.Strict {
				m.Moq.Scene.T.Fatalf("Too many calls to %s", m.Moq.PrettyParams_Fprint(params))
			}
			return
		}
		i = results.Repeat.ResultCount - 1
	}

	result := results.Results[i]
	if result.Sequence != 0 {
		sequence := m.Moq.Scene.NextMockSequence()
		if (!results.Repeat.AnyTimes && result.Sequence != sequence) || result.Sequence > sequence {
			m.Moq.Scene.T.Fatalf("Call sequence does not match call to %s", m.Moq.PrettyParams_Fprint(params))
		}
	}

	if result.DoFn != nil {
		result.DoFn(output, fset, node)
	}

	if result.Values != nil {
		result1 = result.Values.Result1
	}
	if result.DoReturnFn != nil {
		result1 = result.DoReturnFn(output, fset, node)
	}
	return
}

// OnCall returns the recorder implementation of the Config_starGenType type
func (m *MoqConfig_starGenType) OnCall() *MoqConfig_starGenType_recorder {
	return &MoqConfig_starGenType_recorder{
		Moq: m,
	}
}

func (m *MoqConfig_starGenType_recorder) Fprint(output io.Writer, fset *token.FileSet, node any) *MoqConfig_starGenType_Fprint_fnRecorder {
	return &MoqConfig_starGenType_Fprint_fnRecorder{
		Params: MoqConfig_starGenType_Fprint_params{
			Output: output,
			Fset:   fset,
			Node:   node,
		},
		Sequence: m.Moq.Config.Sequence == moq.SeqDefaultOn,
		Moq:      m.Moq,
	}
}

func (r *MoqConfig_starGenType_Fprint_fnRecorder) Any() *MoqConfig_starGenType_Fprint_anyParams {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Any functions must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Fprint(r.Params))
		return nil
	}
	return &MoqConfig_starGenType_Fprint_anyParams{Recorder: r}
}

func (a *MoqConfig_starGenType_Fprint_anyParams) Output() *MoqConfig_starGenType_Fprint_fnRecorder {
	a.Recorder.AnyParams |= 1 << 0
	return a.Recorder
}

func (a *MoqConfig_starGenType_Fprint_anyParams) Fset() *MoqConfig_starGenType_Fprint_fnRecorder {
	a.Recorder.AnyParams |= 1 << 1
	return a.Recorder
}

func (a *MoqConfig_starGenType_Fprint_anyParams) Node() *MoqConfig_starGenType_Fprint_fnRecorder {
	a.Recorder.AnyParams |= 1 << 2
	return a.Recorder
}

func (r *MoqConfig_starGenType_Fprint_fnRecorder) Seq() *MoqConfig_starGenType_Fprint_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("Seq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Fprint(r.Params))
		return nil
	}
	r.Sequence = true
	return r
}

func (r *MoqConfig_starGenType_Fprint_fnRecorder) NoSeq() *MoqConfig_starGenType_Fprint_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Moq.Scene.T.Fatalf("NoSeq must be called before ReturnResults or DoReturnResults calls, recording %s", r.Moq.PrettyParams_Fprint(r.Params))
		return nil
	}
	r.Sequence = false
	return r
}

func (r *MoqConfig_starGenType_Fprint_fnRecorder) ReturnResults(result1 error) *MoqConfig_starGenType_Fprint_fnRecorder {
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
		DoFn       MoqConfig_starGenType_Fprint_doFn
		DoReturnFn MoqConfig_starGenType_Fprint_doReturnFn
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

func (r *MoqConfig_starGenType_Fprint_fnRecorder) AndDo(fn MoqConfig_starGenType_Fprint_doFn) *MoqConfig_starGenType_Fprint_fnRecorder {
	r.Moq.Scene.T.Helper()
	if r.Results == nil {
		r.Moq.Scene.T.Fatalf("ReturnResults must be called before calling AndDo")
		return nil
	}
	last := &r.Results.Results[len(r.Results.Results)-1]
	last.DoFn = fn
	return r
}

func (r *MoqConfig_starGenType_Fprint_fnRecorder) DoReturnResults(fn MoqConfig_starGenType_Fprint_doReturnFn) *MoqConfig_starGenType_Fprint_fnRecorder {
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
		DoFn       MoqConfig_starGenType_Fprint_doFn
		DoReturnFn MoqConfig_starGenType_Fprint_doReturnFn
	}{Sequence: sequence, DoReturnFn: fn})
	return r
}

func (r *MoqConfig_starGenType_Fprint_fnRecorder) FindResults() {
	r.Moq.Scene.T.Helper()
	if r.Results != nil {
		r.Results.Repeat.Increment(r.Moq.Scene.T)
		return
	}

	anyCount := bits.OnesCount64(r.AnyParams)
	insertAt := -1
	var results *MoqConfig_starGenType_Fprint_resultsByParams
	for n, res := range r.Moq.ResultsByParams_Fprint {
		if res.AnyParams == r.AnyParams {
			results = &res
			break
		}
		if res.AnyCount > anyCount {
			insertAt = n
		}
	}
	if results == nil {
		results = &MoqConfig_starGenType_Fprint_resultsByParams{
			AnyCount:  anyCount,
			AnyParams: r.AnyParams,
			Results:   map[MoqConfig_starGenType_Fprint_paramsKey]*MoqConfig_starGenType_Fprint_results{},
		}
		r.Moq.ResultsByParams_Fprint = append(r.Moq.ResultsByParams_Fprint, *results)
		if insertAt != -1 && insertAt+1 < len(r.Moq.ResultsByParams_Fprint) {
			copy(r.Moq.ResultsByParams_Fprint[insertAt+1:], r.Moq.ResultsByParams_Fprint[insertAt:0])
			r.Moq.ResultsByParams_Fprint[insertAt] = *results
		}
	}

	paramsKey := r.Moq.ParamsKey_Fprint(r.Params, r.AnyParams)

	var ok bool
	r.Results, ok = results.Results[paramsKey]
	if !ok {
		r.Results = &MoqConfig_starGenType_Fprint_results{
			Params:  r.Params,
			Results: nil,
			Index:   0,
			Repeat:  &moq.RepeatVal{},
		}
		results.Results[paramsKey] = r.Results
	}

	r.Results.Repeat.Increment(r.Moq.Scene.T)
}

func (r *MoqConfig_starGenType_Fprint_fnRecorder) Repeat(repeaters ...moq.Repeater) *MoqConfig_starGenType_Fprint_fnRecorder {
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
				DoFn       MoqConfig_starGenType_Fprint_doFn
				DoReturnFn MoqConfig_starGenType_Fprint_doReturnFn
			}{
				Values:   last.Values,
				Sequence: r.Moq.Scene.NextRecorderSequence(),
			}
		}
		r.Results.Results = append(r.Results.Results, last)
	}
	return r
}

func (m *MoqConfig_starGenType) PrettyParams_Fprint(params MoqConfig_starGenType_Fprint_params) string {
	return fmt.Sprintf("Fprint(%#v, %#v, %#v)", params.Output, params.Fset, params.Node)
}

func (m *MoqConfig_starGenType) ParamsKey_Fprint(params MoqConfig_starGenType_Fprint_params, anyParams uint64) MoqConfig_starGenType_Fprint_paramsKey {
	m.Scene.T.Helper()
	var outputUsed io.Writer
	var outputUsedHash hash.Hash
	if anyParams&(1<<0) == 0 {
		if m.Runtime.ParameterIndexing.Fprint.Output == moq.ParamIndexByValue {
			outputUsed = params.Output
		} else {
			outputUsedHash = hash.DeepHash(params.Output)
		}
	}
	var fsetUsed *token.FileSet
	var fsetUsedHash hash.Hash
	if anyParams&(1<<1) == 0 {
		if m.Runtime.ParameterIndexing.Fprint.Fset == moq.ParamIndexByValue {
			fsetUsed = params.Fset
		} else {
			fsetUsedHash = hash.DeepHash(params.Fset)
		}
	}
	var nodeUsed any
	var nodeUsedHash hash.Hash
	if anyParams&(1<<2) == 0 {
		if m.Runtime.ParameterIndexing.Fprint.Node == moq.ParamIndexByValue {
			nodeUsed = params.Node
		} else {
			nodeUsedHash = hash.DeepHash(params.Node)
		}
	}
	return MoqConfig_starGenType_Fprint_paramsKey{
		Params: struct {
			Output io.Writer
			Fset   *token.FileSet
			Node   any
		}{
			Output: outputUsed,
			Fset:   fsetUsed,
			Node:   nodeUsed,
		},
		Hashes: struct {
			Output hash.Hash
			Fset   hash.Hash
			Node   hash.Hash
		}{
			Output: outputUsedHash,
			Fset:   fsetUsedHash,
			Node:   nodeUsedHash,
		},
	}
}

// Reset resets the state of the moq
func (m *MoqConfig_starGenType) Reset() { m.ResultsByParams_Fprint = nil }

// AssertExpectationsMet asserts that all expectations have been met
func (m *MoqConfig_starGenType) AssertExpectationsMet() {
	m.Scene.T.Helper()
	for _, res := range m.ResultsByParams_Fprint {
		for _, results := range res.Results {
			missing := results.Repeat.MinTimes - int(atomic.LoadUint32(&results.Index))
			if missing > 0 {
				m.Scene.T.Errorf("Expected %d additional call(s) to %s", missing, m.PrettyParams_Fprint(results.Params))
			}
		}
	}
}
