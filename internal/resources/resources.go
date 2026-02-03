package resources

import (
	"embed"
	"encoding/json"

	"github.com/WaitWut862/cli-dungeon-crawler/internal/common"
	"github.com/WaitWut862/cli-dungeon-crawler/internal/mob"
)

//go:embed items.json
var itemsFile embed.FS

//go:embed mobs.json
var mobsFile embed.FS

type itemsData struct {
	Items []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Durability  int    `json:"durability"`
	} `json:"items"`
}

type mobsData struct {
	Mobs []struct {
		Name          string `json:"name"`
		Description   string `json:"description"`
		Health        int    `json:"health"`
		InventorySize int    `json:"inventorySize"`
	} `json:"mobs"`
}

type MobTemplate struct {
	Name          string
	Description   string
	Health        int
	InventorySize int
}

var LoadedItems []*common.Item
var LoadedMobs []*MobTemplate

func LoadItems() error {
	data, err := itemsFile.ReadFile("items.json")
	if err != nil {
		return err
	}

	var items itemsData
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}

	LoadedItems = make([]*common.Item, len(items.Items))
	for i, item := range items.Items {
		LoadedItems[i] = &common.Item{
			Name:        item.Name,
			Description: item.Description,
			Durability:  item.Durability,
		}
	}

	return nil
}

func GetItemByName(name string) *common.Item {
	for _, item := range LoadedItems {
		if item.Name == name {
			return &common.Item{
				Name:        item.Name,
				Description: item.Description,
				Durability:  item.Durability,
			}
		}
	}
	return nil
}

func LoadMobs() error {
	data, err := mobsFile.ReadFile("mobs.json")
	if err != nil {
		return err
	}

	var mobs mobsData
	if err := json.Unmarshal(data, &mobs); err != nil {
		return err
	}

	LoadedMobs = make([]*MobTemplate, len(mobs.Mobs))
	for i, m := range mobs.Mobs {
		LoadedMobs[i] = &MobTemplate{
			Name:          m.Name,
			Description:   m.Description,
			Health:        m.Health,
			InventorySize: m.InventorySize,
		}
	}

	return nil
}

func GetMobByName(name string) *MobTemplate {
	for _, m := range LoadedMobs {
		if m.Name == name {
			return &MobTemplate{
				Name:          m.Name,
				Description:   m.Description,
				Health:        m.Health,
				InventorySize: m.InventorySize,
			}
		}
	}
	return nil
}

func CreateMob(name string) *mob.Mob {
	template := GetMobByName(name)
	if template == nil {
		return nil
	}
	return mob.NewMob(&mob.MobParams{
		Name:          template.Name,
		Description:   template.Description,
		Health:        template.Health,
		InventorySize: template.InventorySize,
	})
}
