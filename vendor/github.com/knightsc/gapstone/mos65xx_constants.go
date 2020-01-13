/*
Gapstone is a Go binding for the Capstone disassembly library. For examples,
try reading the *_test.go files.

	Library Author: Nguyen Anh Quynh
	Binding Author: Scott Knight
	License: BSD style - see LICENSE file for details
    (c) 2019 Scott Knight

    THIS FILE WAS AUTO-GENERATED -- DO NOT EDIT!
	Command: ./genconst ../capstone/bindings/python/capstone
	2019-12-30T10:12:58-05:00

*/

package gapstone

// #cgo LDFLAGS: -lcapstone
// #cgo freebsd CFLAGS: -I/usr/local/include
// #cgo freebsd LDFLAGS: -L/usr/local/lib
// #include <stdlib.h>
// #include <capstone/capstone.h>
import "C"

// For Capstone Engine. AUTO-GENERATED FILE, DO NOT EDIT [mos65xx_const.py]
const (
	MOS65XX_REG_INVALID = C.MOS65XX_REG_INVALID
	MOS65XX_REG_ACC     = C.MOS65XX_REG_ACC
	MOS65XX_REG_X       = C.MOS65XX_REG_X
	MOS65XX_REG_Y       = C.MOS65XX_REG_Y
	MOS65XX_REG_P       = C.MOS65XX_REG_P
	MOS65XX_REG_SP      = C.MOS65XX_REG_SP
	MOS65XX_REG_DP      = C.MOS65XX_REG_DP
	MOS65XX_REG_B       = C.MOS65XX_REG_B
	MOS65XX_REG_K       = C.MOS65XX_REG_K
	MOS65XX_REG_ENDING  = C.MOS65XX_REG_ENDING
)

const (
	MOS65XX_AM_NONE          = C.MOS65XX_AM_NONE
	MOS65XX_AM_IMP           = C.MOS65XX_AM_IMP
	MOS65XX_AM_ACC           = C.MOS65XX_AM_ACC
	MOS65XX_AM_IMM           = C.MOS65XX_AM_IMM
	MOS65XX_AM_REL           = C.MOS65XX_AM_REL
	MOS65XX_AM_INT           = C.MOS65XX_AM_INT
	MOS65XX_AM_BLOCK         = C.MOS65XX_AM_BLOCK
	MOS65XX_AM_ZP            = C.MOS65XX_AM_ZP
	MOS65XX_AM_ZP_X          = C.MOS65XX_AM_ZP_X
	MOS65XX_AM_ZP_Y          = C.MOS65XX_AM_ZP_Y
	MOS65XX_AM_ZP_REL        = C.MOS65XX_AM_ZP_REL
	MOS65XX_AM_ZP_IND        = C.MOS65XX_AM_ZP_IND
	MOS65XX_AM_ZP_X_IND      = C.MOS65XX_AM_ZP_X_IND
	MOS65XX_AM_ZP_IND_Y      = C.MOS65XX_AM_ZP_IND_Y
	MOS65XX_AM_ZP_IND_LONG   = C.MOS65XX_AM_ZP_IND_LONG
	MOS65XX_AM_ZP_IND_LONG_Y = C.MOS65XX_AM_ZP_IND_LONG_Y
	MOS65XX_AM_ABS           = C.MOS65XX_AM_ABS
	MOS65XX_AM_ABS_X         = C.MOS65XX_AM_ABS_X
	MOS65XX_AM_ABS_Y         = C.MOS65XX_AM_ABS_Y
	MOS65XX_AM_ABS_IND       = C.MOS65XX_AM_ABS_IND
	MOS65XX_AM_ABS_X_IND     = C.MOS65XX_AM_ABS_X_IND
	MOS65XX_AM_ABS_IND_LONG  = C.MOS65XX_AM_ABS_IND_LONG
	MOS65XX_AM_ABS_LONG      = C.MOS65XX_AM_ABS_LONG
	MOS65XX_AM_ABS_LONG_X    = C.MOS65XX_AM_ABS_LONG_X
	MOS65XX_AM_SR            = C.MOS65XX_AM_SR
	MOS65XX_AM_SR_IND_Y      = C.MOS65XX_AM_SR_IND_Y
)

const (
	MOS65XX_INS_INVALID = C.MOS65XX_INS_INVALID
	MOS65XX_INS_ADC     = C.MOS65XX_INS_ADC
	MOS65XX_INS_AND     = C.MOS65XX_INS_AND
	MOS65XX_INS_ASL     = C.MOS65XX_INS_ASL
	MOS65XX_INS_BBR     = C.MOS65XX_INS_BBR
	MOS65XX_INS_BBS     = C.MOS65XX_INS_BBS
	MOS65XX_INS_BCC     = C.MOS65XX_INS_BCC
	MOS65XX_INS_BCS     = C.MOS65XX_INS_BCS
	MOS65XX_INS_BEQ     = C.MOS65XX_INS_BEQ
	MOS65XX_INS_BIT     = C.MOS65XX_INS_BIT
	MOS65XX_INS_BMI     = C.MOS65XX_INS_BMI
	MOS65XX_INS_BNE     = C.MOS65XX_INS_BNE
	MOS65XX_INS_BPL     = C.MOS65XX_INS_BPL
	MOS65XX_INS_BRA     = C.MOS65XX_INS_BRA
	MOS65XX_INS_BRK     = C.MOS65XX_INS_BRK
	MOS65XX_INS_BRL     = C.MOS65XX_INS_BRL
	MOS65XX_INS_BVC     = C.MOS65XX_INS_BVC
	MOS65XX_INS_BVS     = C.MOS65XX_INS_BVS
	MOS65XX_INS_CLC     = C.MOS65XX_INS_CLC
	MOS65XX_INS_CLD     = C.MOS65XX_INS_CLD
	MOS65XX_INS_CLI     = C.MOS65XX_INS_CLI
	MOS65XX_INS_CLV     = C.MOS65XX_INS_CLV
	MOS65XX_INS_CMP     = C.MOS65XX_INS_CMP
	MOS65XX_INS_COP     = C.MOS65XX_INS_COP
	MOS65XX_INS_CPX     = C.MOS65XX_INS_CPX
	MOS65XX_INS_CPY     = C.MOS65XX_INS_CPY
	MOS65XX_INS_DEC     = C.MOS65XX_INS_DEC
	MOS65XX_INS_DEX     = C.MOS65XX_INS_DEX
	MOS65XX_INS_DEY     = C.MOS65XX_INS_DEY
	MOS65XX_INS_EOR     = C.MOS65XX_INS_EOR
	MOS65XX_INS_INC     = C.MOS65XX_INS_INC
	MOS65XX_INS_INX     = C.MOS65XX_INS_INX
	MOS65XX_INS_INY     = C.MOS65XX_INS_INY
	MOS65XX_INS_JML     = C.MOS65XX_INS_JML
	MOS65XX_INS_JMP     = C.MOS65XX_INS_JMP
	MOS65XX_INS_JSL     = C.MOS65XX_INS_JSL
	MOS65XX_INS_JSR     = C.MOS65XX_INS_JSR
	MOS65XX_INS_LDA     = C.MOS65XX_INS_LDA
	MOS65XX_INS_LDX     = C.MOS65XX_INS_LDX
	MOS65XX_INS_LDY     = C.MOS65XX_INS_LDY
	MOS65XX_INS_LSR     = C.MOS65XX_INS_LSR
	MOS65XX_INS_MVN     = C.MOS65XX_INS_MVN
	MOS65XX_INS_MVP     = C.MOS65XX_INS_MVP
	MOS65XX_INS_NOP     = C.MOS65XX_INS_NOP
	MOS65XX_INS_ORA     = C.MOS65XX_INS_ORA
	MOS65XX_INS_PEA     = C.MOS65XX_INS_PEA
	MOS65XX_INS_PEI     = C.MOS65XX_INS_PEI
	MOS65XX_INS_PER     = C.MOS65XX_INS_PER
	MOS65XX_INS_PHA     = C.MOS65XX_INS_PHA
	MOS65XX_INS_PHB     = C.MOS65XX_INS_PHB
	MOS65XX_INS_PHD     = C.MOS65XX_INS_PHD
	MOS65XX_INS_PHK     = C.MOS65XX_INS_PHK
	MOS65XX_INS_PHP     = C.MOS65XX_INS_PHP
	MOS65XX_INS_PHX     = C.MOS65XX_INS_PHX
	MOS65XX_INS_PHY     = C.MOS65XX_INS_PHY
	MOS65XX_INS_PLA     = C.MOS65XX_INS_PLA
	MOS65XX_INS_PLB     = C.MOS65XX_INS_PLB
	MOS65XX_INS_PLD     = C.MOS65XX_INS_PLD
	MOS65XX_INS_PLP     = C.MOS65XX_INS_PLP
	MOS65XX_INS_PLX     = C.MOS65XX_INS_PLX
	MOS65XX_INS_PLY     = C.MOS65XX_INS_PLY
	MOS65XX_INS_REP     = C.MOS65XX_INS_REP
	MOS65XX_INS_RMB     = C.MOS65XX_INS_RMB
	MOS65XX_INS_ROL     = C.MOS65XX_INS_ROL
	MOS65XX_INS_ROR     = C.MOS65XX_INS_ROR
	MOS65XX_INS_RTI     = C.MOS65XX_INS_RTI
	MOS65XX_INS_RTL     = C.MOS65XX_INS_RTL
	MOS65XX_INS_RTS     = C.MOS65XX_INS_RTS
	MOS65XX_INS_SBC     = C.MOS65XX_INS_SBC
	MOS65XX_INS_SEC     = C.MOS65XX_INS_SEC
	MOS65XX_INS_SED     = C.MOS65XX_INS_SED
	MOS65XX_INS_SEI     = C.MOS65XX_INS_SEI
	MOS65XX_INS_SEP     = C.MOS65XX_INS_SEP
	MOS65XX_INS_SMB     = C.MOS65XX_INS_SMB
	MOS65XX_INS_STA     = C.MOS65XX_INS_STA
	MOS65XX_INS_STP     = C.MOS65XX_INS_STP
	MOS65XX_INS_STX     = C.MOS65XX_INS_STX
	MOS65XX_INS_STY     = C.MOS65XX_INS_STY
	MOS65XX_INS_STZ     = C.MOS65XX_INS_STZ
	MOS65XX_INS_TAX     = C.MOS65XX_INS_TAX
	MOS65XX_INS_TAY     = C.MOS65XX_INS_TAY
	MOS65XX_INS_TCD     = C.MOS65XX_INS_TCD
	MOS65XX_INS_TCS     = C.MOS65XX_INS_TCS
	MOS65XX_INS_TDC     = C.MOS65XX_INS_TDC
	MOS65XX_INS_TRB     = C.MOS65XX_INS_TRB
	MOS65XX_INS_TSB     = C.MOS65XX_INS_TSB
	MOS65XX_INS_TSC     = C.MOS65XX_INS_TSC
	MOS65XX_INS_TSX     = C.MOS65XX_INS_TSX
	MOS65XX_INS_TXA     = C.MOS65XX_INS_TXA
	MOS65XX_INS_TXS     = C.MOS65XX_INS_TXS
	MOS65XX_INS_TXY     = C.MOS65XX_INS_TXY
	MOS65XX_INS_TYA     = C.MOS65XX_INS_TYA
	MOS65XX_INS_TYX     = C.MOS65XX_INS_TYX
	MOS65XX_INS_WAI     = C.MOS65XX_INS_WAI
	MOS65XX_INS_WDM     = C.MOS65XX_INS_WDM
	MOS65XX_INS_XBA     = C.MOS65XX_INS_XBA
	MOS65XX_INS_XCE     = C.MOS65XX_INS_XCE
	MOS65XX_INS_ENDING  = C.MOS65XX_INS_ENDING
)

const (
	MOS65XX_GRP_INVALID         = C.MOS65XX_GRP_INVALID
	MOS65XX_GRP_JUMP            = C.MOS65XX_GRP_JUMP
	MOS65XX_GRP_CALL            = C.MOS65XX_GRP_CALL
	MOS65XX_GRP_RET             = C.MOS65XX_GRP_RET
	MOS65XX_GRP_INT             = C.MOS65XX_GRP_INT
	MOS65XX_GRP_IRET            = C.MOS65XX_GRP_IRET
	MOS65XX_GRP_BRANCH_RELATIVE = C.MOS65XX_GRP_BRANCH_RELATIVE
	MOS65XX_GRP_ENDING          = C.MOS65XX_GRP_ENDING
)

const (
	MOS65XX_OP_INVALID = C.MOS65XX_OP_INVALID
	MOS65XX_OP_REG     = C.MOS65XX_OP_REG
	MOS65XX_OP_IMM     = C.MOS65XX_OP_IMM
	MOS65XX_OP_MEM     = C.MOS65XX_OP_MEM
)