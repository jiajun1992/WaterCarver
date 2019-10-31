// Copyright 2017-2018 DERO Project. All rights reserved.
// Use of this source code in any form is governed by RESEARCH license.
// license can be found in the LICENSE file.
// GPG: 0F39 E425 8C65 3947 702A  8234 08B2 0360 A03A 9DE8
//
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
// STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF
// THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// TODO merge this file with the file in ringct

package crypto

// Zero, Identity and L
var Zero = Key{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

var Identity = Key{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

// basepoint of curve basically G
//var GBASE = ScalarmultBase(*(d2h(1)))
var GBASE = HexToKey("5866666666666666666666666666666666666666666666666666666666666666")
var EIGHT = Key{0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
var INV_EIGHT = HexToKey("792fdce229e50661d0da1c7db39dd30700000000000000000000000000000006")
var MINUS_INV_EIGHT = HexToKey("74a4197af07d0bf705c2da252b5c0b0d0000000000000000000000000000000a")
var MINUS_ONE = HexToKey("ecd3f55c1a631258d69cf7a2def9de1400000000000000000000000000000010")

var GBASE_Cached [8]CachedGroupElement // filled in below
var H_Cached [8]CachedGroupElement
var GBASE_H_TABLE PRECOMPUTE_TABLE
func init() { // calculate G cached for faster computations later on
	var B ExtendedGroupElement
	B.FromBytes(&GBASE)
	GePrecompute(&GBASE_Cached, &B)
	B.FromBytes(&H)
	GePrecompute(&H_Cached, &B)
	GenDoublePrecompute(&GBASE_H_TABLE, GBASE, H)
	H_p3.FromBytes(&H)
}

// this is the curve order of ed25519
var L = Key{0xed, 0xd3, 0xf5, 0x5c, 0x1a, 0x63, 0x12, 0x58, 0xd6, 0x9c, 0xf7, 0xa2, 0xde, 0xf9, 0xde, 0x14, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10}

// The other basepoint for use in Pedersen Commitments, which is used for
// Confidential Transactions
// H = G.HashToEC(), where G is the basepoint basically  H = toPoint(cn_fast_hash(G))
var H = Key{0x8b, 0x65, 0x59, 0x70, 0x15, 0x37, 0x99, 0xaf, 0x2a, 0xea, 0xdc, 0x9f, 0xf1, 0xad, 0xd0, 0xea, 0x6c, 0x72, 0x51, 0xd5, 0x41, 0x54, 0xcf, 0xa9, 0x2c, 0x17, 0x3a, 0x0d, 0xd3, 0x9c, 0x1f, 0x94}
var H_p3 ExtendedGroupElement