package games

import "github.com/GoloisaNinja/go-otterly/pkg/models"

var OneWayOrTheOtter = models.Game{
	ID:          "1A",
	Title:       "One Way or the Otter",
	TitleColor:  "#d84dff",
	Image:       "/static/images/owoto.webp",
	Description: "You wake up from cryo-sleep to find alarms blaring and the ship you're on in shambles! Are you the captain or a prisoner? Why does the ship sound like it's coming apart? You'll need to work quickly to solve the mysteries and live to cuddle more space rocks, in this otterly futuristic sci-fi adventure!",
	IsAvailable: true,
	Nodes: []models.GameNode{
		{
			ID:                 "1",
			Text:               `The world is dark and quiet. Or, at least, it was. You are aware of something rhythmic. It’s just at the edge of your perception. But it’s there. Thudding, or maybe, rapping. It grows in intensity until finally you recognize it as an alarm sounding off. Light creeps into the corners of your field of vision and you slowly, arduously, open the heavy lids of your eyes. Blurry at first, your vision finally focuses somewhat and you can see that you are in a cryo-pod. You are suspended in gooey cryo-gel, but through the pod’s forward window, you can see the cryo-bay. There are lockers, a few terminal stations, and there’s a light on the far wall. The light is flashing red. The thudding of the alarm is muted through the gel, but it’s definitely a ship’s alarm sequence. Something is wrong! There doesn’t appear to be anyone in the pod bay. No one to let you out? What is going on? Looks like you’ll have to get out of this pod yourself. But your memory is so foggy from the after-effects of cryo-sleep. Wait. How do you even know that? Whatever, you’ll figure that out later. The point is, there’s got to be a way to exit the pod from the inside. But how? You glance around the inside of the pod. There are all the various tubes keeping you alive and such, and then you spot it. It’s some kind of user panel. The panel is protected from the cryo-gel but appears to be accessible by some physical switches. There’s two switches that grab your attention. One labeled with the characters “ECPP”, and another labeled “EPEP”. Your brain tells you these are significant. You just can’t remember why! Blasted cryo-mind-fog! You’ve got to get out of this pod. Time to make a choice - which switch do you want to activate?`,
			CodeNode:           false,
			CodeLength:         0,
			CodeFailedNextNode: 0,
			Options: []models.GameOptions{
				{
					ID:        "1a",
					Text:      `Type “ecpp” to Engage ECPP switch`,
					Command:   "ecpp",
					Mood:      "space popsicle",
					HasReqs:   false,
					NextNode:  "2",
					PlayType:  "reckless",
					DeathNode: true,
				},
				{
					ID:        "1b",
					Text:      `Type "epep" to Engage EPEP switch`,
					Command:   "epep",
					Mood:      "relief",
					HasReqs:   false,
					Inventory: "test slot",
					NextNode:  "3",
					PlayType:  "reckless",
				},
			},
			EarnedPoints: 0,
		},
		{
			ID:                 "2",
			Text:               `You furiously try to remember what these switches stand for, but it just isn’t coming to you. You know they are both significant, but that’s about it. There are no colors or warnings to indicate that one switch is better than the other. Argh, this brain fog is so frustrating! You slam the switch labeled ECPP and quickly look to the user panel for some feedback. There are a series of muted clicks and you can hear some kind of motorized servo engaging somewhere. This is good! It’s probably the main pod chamber release! Even with cryo-brain-fog you are a genuine, rockstar genius of unfathomable proportion. Suddenly you see a flash across the user panel. The text is bold but it’s so small and hard to read through the cryo-gel. Who designed this crap? You try to rub your eyes, but that actually makes things worse. The text on the user panel keeps changing. Ugh. What does that say? Then you hear several, very serious sounding clicks. No, not clicks. Legitimate thuds. Like, very loud, concerning, thuds. You finally realize the user panel text keeps changing because, well, it’s a countdown. As your cryo-gel covered body is ejected into the cold vastness of space, you suddenly remember that ECPP stands for Evacuate Cryo Pod and Purge. Would have been nice to have remembered that 2 minutes ago. Thanks brain. Before the last breath in your lungs freezes and everything goes black, you curse the makers of the user panel for their poor font family and font size choices. Game Over.`,
			CodeNode:           false,
			CodeLength:         0,
			CodeFailedNextNode: 0,
			Options: []models.GameOptions{
				{
					ID:        "2a",
					Text:      `Type "s" to Start Over`,
					Command:   "s",
					Mood:      "",
					HasReqs:   false,
					Inventory: "",
					NextNode:  "1",
				},
			},
			EarnedPoints: -500,
		},
		{
			ID:                 "3",
			Text:               `You furiously try to remember what these switches stand for, but it just isn’t coming to you. You know they are both significant, but that’s about it. There are no colors or warnings to indicate that one switch is better than the other. Argh, this brain fog is so frustrating! You slam the switch labeled EPEP and quickly look to the user panel for some feedback. You can hear several clicks and the whirring of a motor somewhere. The user panel is flashing some information very quickly across the screen, but the text is so small, you can’t make out what’s happening. Suddenly there is an extremely loud hissing sound that nearly shatters your ear drums, followed by a deep, resonating thud. You are violently ejected from the cryo-pod. All the monitoring nodes and tubes are unceremoniously ripped off or out of you. A speaker near the pod you just got eject from says “Emergency Pod Eject Procedure…is now complete…have a nice day”. Oh right. That’s what EPEP means. Now you remember. Hmm. Still can’t remember what that other switch stands for though. You cough and your throat is scratchy from just having a breathing tube ripped from it, but you glance around the pod bay. All the other pods are empty. You can hear the alarm more clearly now, it’s very annoying, and your vision is clearing up. You move to the lockers and instinctively pick the second one from the right. There’s a blue uniform inside. The patch on the arm reads “One Way Or the Otter”, and the patch on the chest reads “Titan Station - LHF - Liquid HydroCarbon Frigate”. The uniform doesn’t quite fit. You must have put on some weight in cryo-sleep. You look in the mirror next to the lockers and instantly recognize your face. You are an adorable otter. Almost too adorable. And this uniform. Wow. It’s smashing on you. I mean you look fantastic. Just as you’re about to give yourself some finger guns and a few winks. The ship violently jerks to port, throwing you into the row of lockers. A whole new set of alarms start blaring. This is not good. You are trying to desperately to remember who you are. But it’s a total fog. Regardless, you need to do something. And fast! You see what looks like a communication panel on the wall by the only door in the cryo-bay. You could try to use the ship’s general comms and see if anyone responds? You can’t be on this ship alone, can you? Another part of you is thinking that the insane, and violent movement the ship just experienced can only be explained by a problem with the engines. You aren’t sure how you know that, but you’re pretty sure you know that. And yet, still another part of you thinks, that with how good looking you are, that it’s entirely possible, and indeed, likely that you are the captain of this ship called the “One Way or the Otter”. And captains belong on the bridge! What are you going to do?`,
			CodeNode:           false,
			CodeLength:         0,
			CodeFailedNextNode: 0,
			Options: []models.GameOptions{
				{
					ID:      "3a",
					Text:    `Type "comms" to use the Comm Panel`,
					Command: "comms",
					Mood:    "Uhura",
					HasReqs: true,
					Requires: models.OptionReqs{
						Mood: "anger",
					},
					Inventory: "",
					NextNode:  "4",
				},
				{
					ID:        "3b",
					Text:      `Type "engine" to visit Engineering`,
					Command:   "engine",
					Mood:      "Scotty",
					HasReqs:   false,
					Inventory: "",
					NextNode:  "5",
				},
				{
					ID:      "3c",
					Text:    `Type "bridge" to visit the Bridge`,
					Command: "bridge",
					Mood:    "Kirk",
					HasReqs: true,
					Requires: models.OptionReqs{
						Mood:      "",
						Inventory: []string{"test slot", "microphone"},
					},
					Inventory: "",
					NextNode:  "6",
				},
			},
			EarnedPoints: 5,
		},
	},
}
