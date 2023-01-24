package main

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	// 创建RPG游戏存档管理器
	rpgManager := NewRPGArchiveManager()
	// 创建RPG游戏
	rpg := NewRolesPlayGame("暗黑破坏神2", "野蛮人战士")
	fmt.Println(rpg)                  // 输出游戏当前状态
	rpgManager.Put(rpg.Save("第一关存档")) // 游戏存档

	// 第一关闯关失败
	rpg.SetRolesState([]string{"野蛮人战士", "死亡"})
	rpg.SetScenarioState("第一关闯关失败")
	fmt.Println(rpg)

	// 恢复存档，重新闯关
	rpgManager.Reload("第一关存档")
	fmt.Println(rpg)
}
