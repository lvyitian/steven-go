// Generated by protocol_builder
// Do not edit

package steven

import (
	"io"

	"github.com/thinkofdeath/steven/protocol"
)

func (p *pmMinecraftBrand) write(ww io.Writer) (err error) {
	if err = protocol.WriteString(ww, p.Brand); err != nil {
		return
	}
	return
}
func (p *pmMinecraftBrand) read(rr io.Reader) (err error) {
	if p.Brand, err = protocol.ReadString(rr); err != nil {
		return
	}
	return
}
