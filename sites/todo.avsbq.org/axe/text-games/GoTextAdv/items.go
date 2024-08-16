package main

import (
	"errors"
)

type Item struct {
	Name string
	Action string
	ItemForUse int
	Contains []int
}


func FindItemByName(itemName string) (error, int, *Item) {
	for index , itm := range Items {
		if itm.Name == itemName {
			return nil, index, itm
		}
	}
	return errors.New("Item not found"), -1, nil
}

func OpenItem(pla *Character, itemName string) {
	loc := LocationMap[player.CurrentLocation]
	for _, itm := range loc.Items {
		if Items[itm].Name == itemName {
			if Items[itm].ItemForUse != 0 && PlayerHasItem(pla, Items[itm].ItemForUse) {
				loc.Items = append(loc.Items, Items[itm].Contains...)
				Items[itm].Contains = *new([]int)
			} else {
				Output("red", "Could not open the " + itemName)
				return
			}
		} else {
			Output("red", "Could not open the " + itemName)
		}
	}
}

func PlayerHasItem(pla *Character, itm int) bool {
	for _, pitm := range pla.Items {
		if pitm == itm {
			return true
		}
	}
	return false
}

func (it *Item)RemoveItemFromRoom(loc *Location) {
	for index, itm := range loc.Items {
		if Items[itm].Name == it.Name {
			loc.Items = append(loc.Items[:index], loc.Items[index+1:]...)
		}
	}
}

//To be refactored on a location struct
func (it *Item)ItemInRoom(loc *Location) bool {
	for _, itm := range loc.Items {
		if Items[itm].Name == it.Name {
			return true
		}
	}
	return false
}

//To be refactored on a character struct
func (it *Item)ItemOnPlayer(pla *Character) bool {
	for _, itm := range pla.Items {
		if Items[itm].Name == it.Name {
			return true
		}
	}
	return false
}

//To be refactored on a location struct
func describeItems(player Character) {
	l := LocationMap[player.CurrentLocation]

	Output("You see:")
	for _, itm := range l.Items {
		Outputf("\t%s\n", Items[itm].Name)
	}
}

