package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	tm "github.com/buger/goterm"
)

func main() {
	tm.Clear()
	tm.MoveCursor(1, 1)
	tm.Flush()

	mainPath0.commands["left"] = func(s *state) {
		fmt.Println("You Turn Left")
		s.room = mRoom
		s.RoomNo = 4
		s.HiddenCommands["Mausoleum/go through tunnel"] = struct{}{}
	}
	mainPath0.commands["forward"] = func(s *state) {
		fmt.Println("You Move forward")
		s.room = mainPath2
		s.RoomNo = 7
	}

	mRoom.commands["backward"] = func(s *state) {
		fmt.Println("You Move Back to where you came from")
		s.room = mainPath0
		s.RoomNo = 3
	}
	mRoom.commands["pick up hammer"] = func(s *state) {
		fmt.Println("You Pick up the hammer")
		s.Hammergot = true
		s.HiddenCommands["Mausoleum/pick up hammer"] = struct{}{}
	}

	mainPath2.commands["right"] = func(s *state) {
		fmt.Println("You Turn Right")
		s.HiddenCommands["Guard Tower/pick up safe code"] = struct{}{}
		s.room = gTRoom
		s.RoomNo = 5
	}

	mainPath2.commands["backward"] = func(s *state) {
		fmt.Println("You go back")
		s.room = mainPath0
		s.RoomNo = 3
	}
	mainPath2.commands["forward"] = func(s *state) {
		fmt.Println("You continue walking")
		s.room = mainPath3
		s.RoomNo = 9
	}

	gTRoom.commands["backward"] = func(s *state) {
		fmt.Println("You go back")
		s.room = mainPath2
		s.RoomNo = 7
	}
	gTRoom.commands["pick up safe code"] = func(s *state) {
		fmt.Println("You pick up the safe code")
		s.Safecodegot = true
		s.HiddenCommands["Guard Tower/pick up safe code"] = struct{}{}
	}
	mainPath3.commands["backward"] = func(s *state) {
		fmt.Println("You go back")
		s.room = mainPath2
		s.RoomNo = 7
	}
	mainPath3.commands["left"] = func(s *state) {
		s.HiddenCommands["Mansion/grab key"] = struct{}{}
		fmt.Println("You go left")
		s.room = maroom
		s.RoomNo = 6
		s.HiddenCommands["Mansion/grab key"] = struct{}{}
	}
	mainPath3.commands["right"] = func(s *state) {
		fmt.Println("You go right")
		s.room = gRoom
		s.RoomNo = 8
	}
	gRoom.commands["turn on power"] = func(s *state) {
		fmt.Println("You turn on the power")
		s.Electricity = true
		s.HiddenCommands["Generator/turn on power"] = struct{}{}
	}
	gRoom.commands["backward"] = func(s *state) {
		fmt.Println("You turn back")
		s.room = mainPath3
		s.RoomNo = 9
	}
	maroom.commands["backward"] = func(s *state) {
		fmt.Println("You turn back")
		s.room = mainPath3
		s.RoomNo = 9
	}
	mRoom.commands["go through tunnel"] = func(s *state) {
		fmt.Println("You go through the tunnel")
		s.room = mABRoom
		s.RoomNo = 14
	}
	mABRoom.commands["pull lever"] = func(s *state) {
		fmt.Println("You pull the lever and rocks fall into a pit in the center of the room")
		s.room = mABRoom
		s.RoomNo = 14
		s.RocksFallen = true
		s.HiddenCommands["Basement/pull lever"] = struct{}{}
	}
	mABRoom.commands["backward"] = func(s *state) {
		fmt.Println("You go back to the mausoleum")
		s.room = mRoom
		s.RoomNo = 4
		s.HiddenCommands["Mausoleum/go through tunnel"] = struct{}{}
	}
	mainPath3.commands["forward"] = func(s *state) {
		fmt.Println("You forward")
		s.room = mainpath5
		s.RoomNo = 10
		s.HiddenCommands["End of path/exit"] = struct{}{}
	}
	maroom.commands["grab key"] = func(s *state) {
		fmt.Println("You grab the key")
		s.KeyGot = true
		s.HiddenCommands["Mansion/grab key"] = struct{}{}
	}
	mainpath5.commands["exit"] = func(s *state) {
		fmt.Println("You leave.")
		s.room = gamefinish
		s.RoomNo = 12

	}
	mainpath5.commands["backward"] = func(s *state) {
		fmt.Println("You go back")
		s.room = mainPath3
		s.RoomNo = 9
	}
	mainpath5.commands["left"] = func(s *state) {
		fmt.Println("You go left")
		s.room = bRoom
		s.RoomNo = 11
	}
	bRoom.commands["switch off"] = func(s *state) {
		fmt.Println("You switch off the electromagnet")
		s.BreakerRoomUsed = true
		s.HiddenCommands["Breaker Room/switch off"] = struct{}{}
	}
	bRoom.commands["backward"] = func(s *state) {
		fmt.Println("You go back")
		s.room = mainpath5
		s.RoomNo = 10
	}
	s := &state{
		room:           titleRoom,
		HiddenCommands: map[string]struct{}{},
		Movestaken:     0,
	}
	reader := bufio.NewReader(os.Stdin)

	for {
		renderRoom(s.room, s)
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		action, ok := s.room.commands[strings.TrimSpace(text)]
		if strings.TrimSpace(text) == "quit" {
			fmt.Println("\n You Quit the game.\n ")
			os.Exit(0)
		} else if strings.TrimSpace(text) == "save" && s.room != titleRoom && s.room != leftStartPath {
			fmt.Println("\n You save the game.\n ")
			save(s)
		} else if ok && !commandIsHidden(strings.TrimSpace(text), s) {
			s.Movestaken++
			action(s)
		} else {
			fmt.Println()
			fmt.Println("The command You have entered is not valid")
		}
		//if  condition-1 {
		//	// code to be executed if condition-1 is true
		//} else if condition-2 {
		//	// code to be executed if condition-2 is true
		//} else {
		//	// code to be executed if both condition1 and condition2 are false
		//}
	}
}

type state struct {
	//game_data
	room            *room
	Hammergot       bool
	Safecodegot     bool
	Electricity     bool
	BreakerRoomUsed bool
	RocksFallen     bool
	KeyGot          bool
	RoomNo          int
	HiddenCommands  map[string]struct{}
	//info_data
	Movestaken int
}

type action func(s *state)

type room struct {
	Title     string
	Desc      string
	commands  map[string]action
	stateDesc func(s *state) string
}

func renderRoom(r *room, s *state) {
	fmt.Println(r.Title)
	fmt.Println()
	if r.stateDesc != nil {
		fmt.Println(r.stateDesc(s))
	} else {
		fmt.Println(r.Desc)
	}
	fmt.Println()
	fmt.Println("Available Commands:")
	fmt.Println(strings.Join(getCommands(r.commands, s), "|"))
	fmt.Println()
}

func commandIsHidden(cmd string, s *state) bool {
	_, ok := s.HiddenCommands[s.room.Title+"/"+cmd]
	return ok
}

func getCommands(m map[string]action, s *state) (commands []string) {
	commands = append(commands, "quit")
	if s.room != titleRoom && s.room != leftStartPath {
		commands = append(commands, "save")
	}
	for k := range m {
		if _, ok := s.HiddenCommands[s.room.Title+"/"+k]; !ok {
			commands = append(commands, k)

		}
	}
	sort.Strings(commands)
	return
}
func getRoomFromR(r int) *room {
	if r == 1 {
		var n = startRoom
		return n
	} else if r == 2 {
		var n = leftStartPath
		return n
	} else if r == 3 {
		var n = mainPath0
		return n
	} else if r == 4 {
		var n = mRoom
		return n
	} else if r == 5 {
		var n = gTRoom
		return n
	} else if r == 6 {
		var n = maroom
		return n
	} else if r == 7 {
		var n = mainPath2
		return n
	} else if r == 8 {
		var n = gRoom
		return n
	} else if r == 9 {
		var n = mainPath3
		return n
	} else if r == 10 {
		var n = mainpath5
		return n
	} else if r == 11 {
		var n = bRoom
		return n
	} else if r == 12 {
		var n = gamefinish
		return n
	} else if r == 14 {
		var n = mABRoom
		return n
	}
	return startRoom
}

var startRoom = &room{
	Title: "Road",
	Desc:  "You are sprinting over a road. You don't know why you are here, and doubt you ever will know. 3 men in a vehicle are chasing you and you come accross a split in the path. You see a long winding path to your left, and an entrance to what seems like a park on your right. The entrance to the park has a metal gate which you can lock, But once you are inside, there seems to be no way out.",
	commands: map[string]action{
		"left": func(s *state) {
			fmt.Println("\n You turn left.\n ")
			s.room = leftStartPath
			s.RoomNo = 2
		},
		"right": func(s *state) {
			fmt.Println("\n You turn right\n ")
			s.room = mainPath0
			s.RoomNo = 3
		},
	},
}

var leftStartPath = &room{
	Title:    "Left Path",
	Desc:     " You run as fast as you can along the left path. You notice the group in the car continue approching. You are fast, but not fast enough, the car stops and the men get out. They seem to want to kill you. You cannot escape as they drag you into their vehicle. This is the end for you.\n You died.",
	commands: map[string]action{},
}

var mainPath0 = &room{
	Title:    "Right Path",
	Desc:     " You run as fast as you can along the right path. You dive into the enclosed space and lock the gate. You're safe for the moment, but you know that they will wait for you to come out from that gate, no matter what. It is getting dark now and you pull out your lantern. You look around and see a path to your left, There is also a path forewards.\n Which Direction do you go in?",
	commands: map[string]action{},
}
var mRoom = &room{
	Title:    "Mausoleum",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You turn left and walk along the path towards a building that looks like a mausoleum. You look around yourself and notice graves, all over the floor some with bones still half sticking out. You begin to wonder what you got yourself into as you approach the Mausoleum.\n"
		if s.Hammergot == true {
			delete(s.HiddenCommands, "Mausoleum/go through tunnel")
			return general + "There used to be a hammer on the floor, but you picked it up. There is also now a tunnel that you can go through."
		}

		return general + "There is a hammer on the floor."
	},
}
var mainPath2 = &room{
	Title:    "Path",
	Desc:     "You walk to the next intersection on the path and notice a prison guard tower to your right. It streaches far into the sky but has no staircase. You can continue walking forward, or take the path to your right.",
	commands: map[string]action{},
}
var gTRoom = &room{
	Title:    "Guard Tower",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You walk towards the guard tower, the night falling around you, and see that"
		if s.Electricity == true && s.BreakerRoomUsed == true {
			delete(s.HiddenCommands, "Guard Tower/pick up safe code")
			return general + " the door is unlocked, You enter and find a metal block on the floor with a safe code sticking out of it"
		}
		if s.Electricity == true {

			return general + " the door is unlocked, You enter and find a metal block on the floor with a safe code sticking out of it. You cannot lift it as an electromagnet has it stuck to the floor"
		}

		return general + " the door is locked, it requires electricity to function"
	},
}
var mainPath3 = &room{
	Title:    "Path",
	Desc:     "You walk to the next intersection on the path and look around yourself. You look to your left and see an abandoned mansion, covered in winding vines. To your left you see a small house, Presumably where the generator is kept.",
	commands: map[string]action{},
}
var maroom = &room{
	Title:    "Mansion",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You walk along the short path, observing your surroundings. You notice the distant screams coming from the car that was chasing you. Shivers make their way down your spine as you peek inside the old building. Inside you notice"
		if s.RocksFallen == true && s.Safecodegot == true && s.Hammergot == true {
			delete(s.HiddenCommands, "Mansion/grab key")
			return general + " a key guarded by a safe. You smash the glass, unlock the safe with the code and see a key to the exit."
		}
		if s.RocksFallen == true {
			if s.Hammergot {
				return general + " a safe guarded by a sheet of glass. You smash the glass using a hammer to reveal a safe with a code."
			}
			return general + " a safe guarded by a sheet of glass. A hammer could be useful."
		}

		return general + " a pile of rocks blocking your path."
	},
}
var gRoom = &room{
	Title:    "Generator",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You enter the room with the power generator, the switch on the wall controlling the power output is set to "
		if !s.Electricity {
			delete(s.HiddenCommands, "Generator/turn on power")
			return general + "off."
		}

		return general + "on."
	},
}
var mABRoom = &room{
	Title:    "Basement",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You crawl through the tunnel and reach the basement of the Mansion. There is a lever controlling a trapdoor in the floor above. The lever is "
		if !s.RocksFallen {
			delete(s.HiddenCommands, "Basement/pull lever")
			return general + "not pulled."
		}

		return general + "pulled."
	},
}

var mainpath5 = &room{
	Title:    "End of path",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You walk along the path and reach a gate. This appears to be your way out. You can turn right."
		if !s.KeyGot {
			delete(s.HiddenCommands, "mainpath5/exit")
		}
		return general
	},
}

var bRoom = &room{
	Title:    "Breaker Room",
	commands: map[string]action{},
	stateDesc: func(s *state) string {
		general := "You Enter the breaker room and notice a panel of switches. You see one labeled \"Electromagnet\" ."
		if !s.BreakerRoomUsed {
			delete(s.HiddenCommands, "Breaker Room/switch off electromagnet")
		}
		return general
	},
}
var titleRoom = &room{
	Title: "Eddie's Text Adventure",
	Desc:  "MENU",
	commands: map[string]action{
		"start": func(s *state) {
			s.room = startRoom
			s.RoomNo = 1
		},
		"load": func(s *state) {
			ok, err := load(s)
			if err != nil {
				fmt.Println(err)
			}
			if !ok {
				fmt.Println("\n No file to load.")
			}
			if err == nil && ok == true {
				s.room = getRoomFromR(s.RoomNo)
			}

		},
	},
}
var gamefinish = &room{
	Title:    "Escape",
	Desc:     "You escape from the place you were in and run from the people in the car. You escape and get home. The end",
	commands: map[string]action{},
}
