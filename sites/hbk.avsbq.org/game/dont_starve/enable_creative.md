# Enable creative mode

I found out how from a 11 year old reddit [post](https://www.reddit.com/r/dontstarve/comments/1fbaqt/how_to_enable_creative_mode_including_how_to/)

How to enable console commands:

1. On Windows & Mac go to <Documents>\Klei\DoNotStarve\settings.ini On Linux go to ~/.klei/DoNotStarve/settings.ini

2. In settings.ini change

`ENABLECONSOLE = false`

to

`ENABLECONSOLE = true`

## Console Commands:

Use the '~' key in game to open the console

Use the Enter/Return key to execute the command

Warning: All commands will wear off upon leaving the game

Creative Mode:

This will give you the ability to craft every recipe wether you have the resources or not.

`GetPlayer().components.builder:GiveAllRecipes()`

Invincibility:

`GetPlayer().components.health:SetInvincible(true)`

Pause Hunger:

`GetPlayer().components.hunger:Pause(true)`

Set Max Sanity:

`GetPlayer().components.sanity:SetMax(insert number)`

Set Max Health:

`GetPlayer().components.health:SetMaxHealth(insert number)`

Skip days:

`for x = 1, 15 do GetClock():MakeNextDay() end`

Replace the 15 with the number corresponding to the number of days you want to skip. Be aware this is very slow because the game saves once everyday

Reveal Full Map: First execute

`minimap = TheSim:FindFirstEntityWithTag("minimap")`

Next execute

`minimap.MiniMap:ShowArea(0,0,0, 10000)`

Useful functions

Enabling useful functions

`RunScript("consolecommands")`

You must execute this command before entering any command starting with

	c_

Also all commands needing this command I tagged with UF

Set Hunger (UF)

	c_sethunger(percent)

For Example: Entering

	c_sethunger(0.50)

will set your hunger to 50%

Set Health (UF)

	c_sethealth(0.percent)

Set Sanity (UF)

	c_setsanity(0.percent)

Invincibility (UF)

	c_godmode()

Spawning Items

Pre-fabricated (aka prefab) names are used to identify objects in the game. A list of all prefab names can be found on this Don't Starve Wiki page.

Spawning 1 item

`DebugSpawn("prefabname")`

Spawning more than 1 item

`for x = 1, 40 do DebugSpawn("prefabname") end`

Spawning items (UF)

	c_spawn("prefabname", amount)

Spawning items directly into inventory (UF)

	c_give("prefabname", amount)

Only spawn items that can be stored in your inventory with this command
