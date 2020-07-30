Package eew is EEW reader/decoder
===================
[![GoDev][godev-image]][godev-url]
![Go](https://github.com/skubota/eew/workflows/Go/badge.svg)

[godev-image]: https://pkg.go.dev/badge/github.com/skubota/eew
[godev-url]: https://pkg.go.dev/github.com/skubota/eew


DESCRIPTION

Earthquake Early Warning(Japan)
The Earthquake Early Warning (EEW) (緊急地震速報 ,Kinkyu- Jishin Sokuho-) is a warning which is issued just after an earthquake in Japan is detected.
The warnings are issued mainly by Japan Meteorological Agency (JMA).
JMA has two EEW schemes. One is for advanced users. The other is for the general public, which is mainly mentioned in detail in this article.
this package support for advanced users scheme.

usage

	$ go get github.com/skubota/eew

package eew 

<https://pkg.go.dev/github.com/skubota/eew>

sample

<https://play.golang.org/p/twDmdbRdCHV>

```go
package main

import (
	"fmt"
	"github.com/skubota/eew"
)

func main() {
	data := `
37 03 00 110311144702 C11
110311144616
ND20110311144640 NCN009 JD////////////// JN///
288 N381 E1429 010 76 5- RK66444 RT11/// RC0////
EBI 222 S5-04 ////// 11 220 S5-04 ////// 11 211 S5-04 ////// 11
210 S5-04 144703 10 221 S5-04 144703 10 213 S0404 ////// 11
251 S0404 144704 10 250 S0404 144711 10 241 S0404 144715 10
212 S0404 144715 10 242 S0404 144715 10 233 S0404 144715 10
300 S0404 144721 00 252 S0404 144722 00 240 S0404 144727 00
243 S0403 144719 00 231 S0403 144730 00 202 S0403 144732 00
372 S0403 144732 00 301 S0403 144733 00 230 S0403 144736 00
340 S0403 144739 00 331 S0403 144748 00
9999=
`
	fmt.Printf("%#v", eew.Decoder(data))
}
```



