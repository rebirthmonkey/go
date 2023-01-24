package main

import (
	"fmt"
	"testing"
)

func TestInterpreter(t *testing.T) {
	isAntarcticaExpression := generateCheckAntarcticaExpression()
	// 大洲描述1
	continentDescription1 := "此大洲生活着大量企鹅，全年低温，并且伴随着有暴风雪"
	fmt.Printf("%s，是否是南极洲？%t\n", continentDescription1, isAntarcticaExpression.Interpret(continentDescription1))
	// 大洲描述2
	continentDescription2 := "此大洲生活着狮子，全年高温多雨"
	fmt.Printf("%s，是否是南极洲？%t\n", continentDescription2, isAntarcticaExpression.Interpret(continentDescription2))

	isAmericanExpression := generateCheckAmericanExpression()
	peopleDescription1 := "此人生活在北美洲的黑人，说着英语，持有美国绿卡"
	fmt.Printf("%s，是否是美国人？%t\n", peopleDescription1, isAmericanExpression.Interpret(peopleDescription1))

	peopleDescription2 := "此人生活在欧洲，说着英语，是欧洲议会议员"
	fmt.Printf("%s，是否是南极洲？%t\n", peopleDescription2, isAmericanExpression.Interpret(peopleDescription2))

}

// generateCheckAntarcticaExpression 生成校验是否是南极洲表达式
func generateCheckAntarcticaExpression() Expression {
	// 判断南极洲的动物，或关系
	animalExpression := NewOrExpression(NewTerminalExpression("企鹅"),
		NewTerminalExpression("蓝鲸"))
	// 判断南极洲的天气，与关系
	weatherExpression := NewAndExpression(NewTerminalExpression("低温"),
		NewTerminalExpression("暴风雪"))
	// 最终返回动物与天气的与关系
	return NewAndExpression(animalExpression, weatherExpression)
}

// generateCheckAmericanExpression 生成检查美国人表达式
func generateCheckAmericanExpression() Expression {
	// 人种判断，或关系
	raceExpression := NewOrExpression(NewTerminalExpression("白人"),
		NewTerminalExpression("黑人"))
	// 生活方式，与关系
	lifeStyleExpression := NewAndExpression(NewTerminalExpression("英语"),
		NewTerminalExpression("北美洲"))
	// 身份，与关系
	identityExpression := NewAndExpression(lifeStyleExpression, NewTerminalExpression("美国绿卡"))
	return NewAndExpression(raceExpression, identityExpression)
}
