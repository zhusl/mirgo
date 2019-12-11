package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
	"testing"
)

func TestDB(t *testing.T) {
	path := os.Getenv("GOPATH") + "/src/github.com/yenkeia/mirgo/dotnettools/mir.sqlite"
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var basic Basic
	db.Table("basic").Where("game_version = ?", 65).Find(&basic)
	t.Log(basic.GameVersion, basic.MapIndex, basic.RespawnIndex)

	var gameShopItem GameShopItem
	db.Table("game_shop_item").Where("id = ?", 2).Find(&gameShopItem)
	t.Log(gameShopItem.GoldPrice, gameShopItem.Id)

	var magicInfo MagicInfo
	db.Table("magic_info").Where("name = ?", "Fencing").Find(&magicInfo)
	t.Log(magicInfo.Name, magicInfo.Icon)

	var mapInfo MapInfo
	db.Table("map_info").Where("id = ?", 1).Find(&mapInfo)
	t.Log(mapInfo.Title)

	//var mineZone com.MineZone
	//db.Table("mine_zone").Where("map_index = ?", )

	var monsterInfo MonsterInfo
	db.Table("monster_info").Where("id = ?", 1).Find(&monsterInfo)
	t.Log(monsterInfo.Name)

	var movementInfo MovementInfo
	db.Table("movement_info").Where("map_id = ?", 2).First(&movementInfo)
	t.Log(movementInfo.MapId, movementInfo.ConquestIndex, movementInfo.DestinationX, movementInfo.DestinationY)

	var npcInfo NpcInfo
	db.Table("npc_info").Where("id = ?", 1).Find(&npcInfo)
	t.Log(npcInfo.Filename)

	var questInfo QuestInfo
	db.Table("quest_info").Where("id = ?", 1).Find(&questInfo)
	t.Log(questInfo.Name)

	var respawnInfo RespawnInfo
	db.Table("respawn_info").Where("location_x = ?", 350).Find(&respawnInfo)
	t.Log(respawnInfo.MapId, respawnInfo.RespawnIndex, respawnInfo.Count)

	var safeZoneInfo SafeZoneInfo
	db.Table("safe_zone_info").Where("map_id = ?", 1).Find(&safeZoneInfo)
	t.Log(safeZoneInfo.MapId, safeZoneInfo.LocationX, safeZoneInfo.LocationY)
}
