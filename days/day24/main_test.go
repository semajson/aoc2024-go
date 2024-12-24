package day24

import "testing"

func TestDayExampleP1(t *testing.T) {
	example := `x00: 1
x01: 1
x02: 1
y00: 0
y01: 1
y02: 0

x00 AND y00 -> z00
x01 XOR y01 -> z01
x02 OR y02 -> z02`
	result := Solve1(example)
	expect := 4

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExample2P1(t *testing.T) {
	example := `x00: 1
x01: 0
x02: 1
x03: 1
x04: 0
y00: 1
y01: 1
y02: 1
y03: 1
y04: 1

ntg XOR fgs -> mjb
y02 OR x01 -> tnw
kwq OR kpj -> z05
x00 OR x03 -> fst
tgd XOR rvg -> z01
vdt OR tnw -> bfw
bfw AND frj -> z10
ffh OR nrd -> bqk
y00 AND y03 -> djm
y03 OR y00 -> psh
bqk OR frj -> z08
tnw OR fst -> frj
gnj AND tgd -> z11
bfw XOR mjb -> z00
x03 OR x00 -> vdt
gnj AND wpb -> z02
x04 AND y00 -> kjc
djm OR pbm -> qhw
nrd AND vdt -> hwm
kjc AND fst -> rvg
y04 OR y02 -> fgs
y01 AND x02 -> pbm
ntg OR kjc -> kwq
psh XOR fgs -> tgd
qhw XOR tgd -> z09
pbm OR djm -> kpj
x03 XOR y03 -> ffh
x00 XOR y04 -> ntg
bfw OR bqk -> z06
nrd XOR fgs -> wpb
frj XOR qhw -> z04
bqk OR frj -> z07
y03 OR x01 -> nrd
hwm AND bqk -> z03
tgd XOR rvg -> z12
tnw OR pbm -> gnj`
	result := Solve1(example)
	expect := 2024

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %d; \nWanted: %d", example, result, expect)
	}
}

func TestDayExampleP2(t *testing.T) {
	example := `x00: 0
x01: 1
x02: 0
x03: 1
x04: 0
x05: 1
y00: 0
y01: 0
y02: 1
y03: 1
y04: 0
y05: 1

x00 AND y00 -> z05
x01 AND y01 -> z02
x02 AND y02 -> z01
x03 AND y03 -> z03
x04 AND y04 -> z04
x05 AND y05 -> z00`
	result := Solve2(example)
	expect := "z00,z01,z02,z05"

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %s; \nWanted: %s", example, result, expect)
	}
}

func TestSet_x_y(t *testing.T) {
	wire_vals := map[string]int{}

	Set_x_y(wire_vals, 11, 13)

	// x00: 1
	val, exists := wire_vals["x00"]
	if !exists || val != 1 {
		t.Errorf("Failed for x00")
	}

	// x01: 1
	val, exists = wire_vals["x01"]
	if !exists || val != 1 {
		t.Errorf("Failed for x01")
	}

	// x02: 0
	val, exists = wire_vals["x02"]
	if !exists || val != 0 {
		t.Errorf("Failed for x02")
	}

	// x03: 1
	val, exists = wire_vals["x03"]
	if !exists || val != 1 {
		t.Errorf("Failed for x03")
	}

	// y00: 1
	val, exists = wire_vals["y00"]
	if !exists || val != 1 {
		t.Errorf("Failed for y00")
	}

	// y01: 0
	val, exists = wire_vals["y01"]
	if !exists || val != 0 {
		t.Errorf("Failed for y01")
	}

	// y02: 1
	val, exists = wire_vals["y02"]
	if !exists || val != 1 {
		t.Errorf("Failed for y02")
	}

	// y03: 1
	val, exists = wire_vals["y03"]
	if !exists || val != 1 {
		t.Errorf("Failed for y03")
	}
}

func TestSet_x_y_2(t *testing.T) {
	wire_vals := map[string]int{}
	Set_x_y(wire_vals, 8, 0)

	// x03: 1
	val, exists := wire_vals["x03"]
	if !exists || val != 1 {
		t.Errorf("Failed for x03")
	}
	// x04: 0
	val, exists = wire_vals["x04"]
	if !exists || val != 0 {
		t.Errorf("Failed for x04")
	}
	// x02: 0
	val, exists = wire_vals["x02"]
	if !exists || val != 0 {
		t.Errorf("Failed for x02")
	}

	wire_vals = map[string]int{}
	Set_x_y(wire_vals, 7, 0)

	// x02: 1
	val, exists = wire_vals["x02"]
	if !exists || val != 1 {
		t.Errorf("Failed for x02")
	}
	// x03: 0
	val, exists = wire_vals["x03"]
	if !exists || val != 0 {
		t.Errorf("Failed for x03")
	}

	wire_vals = map[string]int{}
	Set_x_y(wire_vals, 9, 0)

	// x03: 1
	val, exists = wire_vals["x03"]
	if !exists || val != 1 {
		t.Errorf("Failed for x03")
	}
	// x04: 0
	val, exists = wire_vals["x04"]
	if !exists || val != 0 {
		t.Errorf("Failed for x04")
	}
	// x00: 1
	val, exists = wire_vals["x00"]
	if !exists || val != 1 {
		t.Errorf("Failed for x00")
	}
	// x02: 0
	val, exists = wire_vals["x02"]
	if !exists || val != 0 {
		t.Errorf("Failed for x02")
	}
}

func TestDayExample3P2(t *testing.T) {
	example := `x00: 1
x01: 0
x02: 0
x03: 1
x04: 1
x05: 1
x06: 0
x07: 0
x08: 0
x09: 1
x10: 0
x11: 0
x12: 0
x13: 1
x14: 1
x15: 1
x16: 0
x17: 0
x18: 0
x19: 1
x20: 0
x21: 1
x22: 0
x23: 1
x24: 1
x25: 0
x26: 1
x27: 1
x28: 1
x29: 0
x30: 0
x31: 0
x32: 0
x33: 0
x34: 0
x35: 1
x36: 0
x37: 1
x38: 1
x39: 0
x40: 1
x41: 0
x42: 0
x43: 0
x44: 1
y00: 1
y01: 1
y02: 1
y03: 1
y04: 0
y05: 1
y06: 0
y07: 1
y08: 0
y09: 1
y10: 1
y11: 1
y12: 1
y13: 0
y14: 1
y15: 0
y16: 0
y17: 0
y18: 1
y19: 1
y20: 0
y21: 0
y22: 1
y23: 0
y24: 0
y25: 1
y26: 0
y27: 1
y28: 0
y29: 0
y30: 0
y31: 1
y32: 1
y33: 1
y34: 1
y35: 0
y36: 1
y37: 1
y38: 1
y39: 1
y40: 0
y41: 1
y42: 0
y43: 1
y44: 1

nsg AND sdh -> nbq
nhs XOR qvh -> z21
htr OR kdm -> vkt
dvb OR wtv -> cdm
x37 AND y37 -> jdk
y27 AND x27 -> snj
cnp OR ddh -> pjt
fvp AND kgr -> vbr
fvd OR rpm -> srq
y08 AND x08 -> pvm
y43 AND x43 -> nhq
vrk OR hmt -> tpv
y18 AND x18 -> kdm
x27 XOR y27 -> nhh
jmc XOR qkk -> z29
x01 AND y01 -> hqh
y36 AND x36 -> hsc
y43 XOR x43 -> wjf
x26 AND y26 -> dvb
y31 AND x31 -> mds
qkk AND jmc -> hnj
wjt AND ftt -> htr
x40 XOR y40 -> bwc
x42 AND y42 -> tbp
x34 AND y34 -> tgb
nqg AND hbn -> wqc
hkt OR sms -> fqk
x04 AND y04 -> cjg
cfd XOR cqd -> z28
jpp AND tpv -> vqk
x17 AND y17 -> ppw
tvs XOR mkh -> z31
cqw OR wtc -> dnp
y21 XOR x21 -> qvh
y24 AND x24 -> bvn
y32 XOR x32 -> cwr
hsp AND qnp -> dds
x03 XOR y03 -> vng
hvg XOR gtg -> z10
hpm OR nhq -> kkt
cwn XOR pmt -> z07
y12 XOR x12 -> nsg
cnv AND jsj -> ddh
srq AND sdr -> dpm
ktj XOR dmh -> z38
tcv AND gqh -> cdk
y13 AND x13 -> bfw
ptt OR bvn -> nnw
x01 XOR y01 -> nmk
vnn XOR qjh -> z14
y39 AND x39 -> sms
y07 AND x07 -> tcb
x23 XOR y23 -> cnv
nqg XOR hbn -> z08
ppp AND pjt -> ptt
ngc AND sdt -> rbk
hsp XOR qnp -> z15
y07 XOR x07 -> cwn
x17 XOR y17 -> mhm
y12 AND x12 -> qjf
ssq AND skd -> nqw
y24 XOR x24 -> ppp
x34 XOR y34 -> bmh
x15 AND y15 -> btb
cmd AND pfk -> cqw
scb XOR dnp -> z36
mmr AND cgr -> vrk
cdk OR rmn -> z20
x44 XOR y44 -> cbq
nnf OR vbr -> hvg
nnw XOR qnd -> z25
y05 XOR x05 -> cgr
y37 XOR x37 -> vnh
jmp OR nqw -> vnw
y20 AND x20 -> rmn
y44 AND x44 -> rqv
bjm AND vkt -> wtr
x29 AND y29 -> ddt
kbj XOR nsd -> z13
nss OR vjd -> vss
hjp XOR mhm -> z17
y35 AND x35 -> wtc
x02 XOR y02 -> ssq
fnk XOR cwr -> z32
jsj XOR cnv -> z23
rtw XOR wpq -> z11
x22 AND y22 -> gwt
vng XOR vnw -> z03
vnn AND qjh -> fgv
kwf OR nbh -> sdh
y11 XOR x11 -> rtw
kbv XOR bfc -> z42
x29 XOR y29 -> qkk
hnj OR ddt -> fmn
wjf AND dcd -> hpm
ddr XOR wkn -> z16
krm AND vdc -> svv
wjf XOR dcd -> z43
y26 XOR x26 -> cpg
y40 AND x40 -> rpm
qnd AND nnw -> gnv
x33 XOR y33 -> ngc
y06 AND x06 -> kgh
nhh AND cdm -> vmf
y33 AND x33 -> mdv
tcb OR fgg -> nqg
sdt XOR ngc -> z33
y28 XOR x28 -> cqd
vmf OR snj -> cfd
nmk AND dsr -> qrt
y18 XOR x18 -> ftt
y19 XOR x19 -> bjm
x19 AND y19 -> rwr
x00 XOR y00 -> z00
bwc XOR fqk -> z40
y31 XOR x31 -> tvs
ktj AND dmh -> msm
nhh XOR cdm -> z27
vng AND vnw -> fcg
wrc OR tgb -> pfk
mds OR cmn -> fnk
mmr XOR cgr -> z05
pjj OR wvn -> wpq
mtq AND vnh -> jtf
hsc OR dpq -> mtq
x42 XOR y42 -> kbv
y30 XOR x30 -> ddn
scb AND dnp -> dpq
y09 XOR x09 -> fvp
x20 XOR y20 -> gqh
bss XOR fch -> z39
x28 AND y28 -> dgj
y41 AND x41 -> pqq
vss AND wnq -> tjq
cfd AND cqd -> pht
tbp OR tts -> dcd
bfc AND kbv -> tts
wnq XOR vss -> z22
bmh AND spk -> z34
dpm OR pqq -> bfc
hjp AND mhm -> jqs
srq XOR sdr -> z41
jhb OR fgv -> hsp
gqt OR gnv -> rvb
y39 XOR x39 -> bss
y05 AND x05 -> hmt
gtg AND hvg -> wvn
y02 AND x02 -> jmp
svv OR cjg -> mmr
pfk XOR cmd -> z35
krm XOR vdc -> z04
x15 XOR y15 -> qnp
kkt XOR cbq -> z44
y10 XOR x10 -> gtg
vqk OR kgh -> pmt
kbj AND nsd -> hdc
mkh AND tvs -> cmn
fmn XOR kqh -> z30
ppw OR jqs -> wjt
wtr OR rwr -> tcv
jpp XOR tpv -> z06
y36 XOR x36 -> scb
nhs AND qvh -> vjd
x03 AND y03 -> qds
y21 AND x21 -> nss
x04 XOR y04 -> krm
bjm XOR vkt -> z19
x25 AND y25 -> gqt
rtw AND wpq -> nbh
cwn AND pmt -> fgg
x23 AND y23 -> cnp
y25 XOR x25 -> qnd
vnh XOR mtq -> z37
nmk XOR dsr -> z01
dds OR btb -> ddr
dnw OR dtt -> sdt
y35 XOR x35 -> cmd
y38 XOR x38 -> dmh
y08 XOR x08 -> hbn
x16 XOR y16 -> wkn
dtk OR ddn -> mkh
pht OR dgj -> jmc
kgr XOR fvp -> nnf
y06 XOR x06 -> jpp
pjt XOR ppp -> z24
skd XOR ssq -> z02
x14 XOR y14 -> qjh
x30 AND y30 -> kqh
x00 AND y00 -> dsr
rvb AND cpg -> wtv
y16 AND x16 -> qsd
x13 XOR y13 -> nsd
nbq OR qjf -> kbj
kqh AND fmn -> dtk
spk XOR bmh -> wrc
ddr AND wkn -> jsv
y11 AND x11 -> kwf
y38 AND x38 -> gcn
gqh XOR tcv -> nhs
y41 XOR x41 -> sdr
y14 AND x14 -> jhb
hqh OR qrt -> skd
rbk OR mdv -> spk
bss AND fch -> hkt
jtf OR jdk -> ktj
bwc AND fqk -> fvd
cbq AND kkt -> cdh
hdc OR bfw -> vnn
y09 AND x09 -> z09
rvb XOR cpg -> z26
qds OR fcg -> vdc
wqc OR pvm -> kgr
jsv OR qsd -> hjp
cdh OR rqv -> z45
tjq OR gwt -> jsj
msm OR gcn -> fch
sdh XOR nsg -> z12
y10 AND x10 -> pjj
cwr AND fnk -> dtt
y22 XOR x22 -> wnq
wjt XOR ftt -> z18
y32 AND x32 -> dnw`
	result := Solve2(example)
	expect := "z00,z01,z02,z05"

	if result != expect {
		t.Errorf("Failed! \nSolve(%q)\nEquals: %s; \nWanted: %s", example, result, expect)
	}
}
