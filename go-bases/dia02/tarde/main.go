package main

import (
	"banana/controle"
	"banana/lists"
	"banana/loops"
	"banana/operadores"
)

func main() {
	operadores.Operators()
	controle.ControleComIfElse()
	controle.ControleComSwitch()
	controle.ControleComSwitchMultiplosCases()
	controle.ControleComSwitchComFallthrough()
	loops.LoopInfinito()
	loops.LoopWhile()
	loops.LoopComContinue()
	loops.LoopComBreak()
	lists.ManipularArray()
}
