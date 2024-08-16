package main

import (
	"os"
	"strings"
)
func ProcessCommands(player *Character, input string) {
	Output("yellow", "======================================================================")
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		Output("red", "No command received.")
		return
	}
	command := strings.ToLower(tokens[0])
	itemName := ""
	if len(tokens) > 1 {
		itemName = tokens[1]
	}
	loc := LocationMap[player.CurrentLocation]
	switch(command) {
	case "go":
		fallthrough
	case "goto":
		if loc.CanGoTo(strings.ToLower(itemName)) {
			locName, err := FindLocationName(strings.ToLower(itemName))
			if err != nil {
				Output("red", "Can't go to " + itemName + " from here.")
			} else {
				player.CurrentLocation = locName
			}
		} else {
			Output("red", "Can't go to " + itemName + " from here.")
		}
	case "get":
		err, index, itm := FindItemByName(itemName)
		//Make sure we do not pick it up twice
		if err == nil && itm.ItemInRoom(loc)  && !itm.ItemOnPlayer(player) {
			player.Items = append(player.Items, index)  // Add Item to Player's bag
			itm.RemoveItemFromRoom(loc)
		} else {
			Output("Could not get " + itemName)
		}
	case "open":
		OpenItem(player, itemName)
	case "inv":
		Output("yellow", "Your Inventory: ")
		for _, itm := range player.Items {
			Output("yellow", "\t" + Items[itm].Name)
		}
	case "help":
		Output("blue", "Commands:")
		Output("blue", "\tgo <Location Name> - Move to the new location")
		Output("blue", "\tattack - Attacks opponent(s)")
		Output("blue", "\tblock - Block incoming attack")
		Output("blue", "\trun - Escape attack")
		Output("blue", "\tget <Item Name> - Pick up item")
		Output("blue", "\topen <Item Name> - Open an iten if it can be opened")
		Output("blue", "\tinv - Show what you are carrying\n\n")
	case "quit":
		Output("green", "Goodbye...")
		os.Exit(0)
	default:
	}
}
