// by catfish.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package report

import (
	"regexp"

	"github.com/google/syzkaller/pkg/symbolizer"
)

type ucore struct {
	kernelSrc string
	kernelObj string
	symbols   map[string][]symbolizer.Symbol
	ignores   []*regexp.Regexp
}

func ctorUcore(kernelSrc, kernelObj string, symbols map[string][]symbolizer.Symbol,
	ignores []*regexp.Regexp) (Reporter, error) {
	ctx := &ucore{
		kernelSrc: kernelSrc,
		kernelObj: kernelObj,
		symbols:   symbols,
		ignores:   ignores,
	}
	return ctx, nil
}

func (ctx *ucore) ContainsCrash(output []byte) bool {
	panic("not implemented")
}

func (ctx *ucore) Parse(output []byte) *Report {
	panic("not implemented")
}

func (ctx *ucore) Symbolize(rep *Report) error {
	panic("not implemented")
}
