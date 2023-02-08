const otterFacts = [
    "Otters use rocks to smash food",
    "Otters keep favourite rocks in armpit pockets",
    "Otters mastered space travel before humans",
    "Most Otters have lovely singing voices",
    "An Otter can outrun a cheetah on a waterslide",
    "Otters can own humans as pets in Madagascar",
    "Otters are exempt from sales tax in most states",
    "Otters are prohibited from forming unions because of how good they are at it",
    "Otters can hold their breath for nearly 50 years by taking small breaths occasionally",
    "Otters love hot tubs and time machines",
    "An Otter currently holds the record for fastest speed run of E.T. on the Atari",
    "Otters do not wear shirts with buttons because they prefer zippers",
    "You are missing the match to that one sock because an Otter took it",
    "Otters are masters of all forms of warfare and kung-fu",
    "Otters prefer and often demand remote work",
    "Otters make great getaway drivers, obviously",
    "The record for most touchdowns scored underwater belongs to an Otter named Cheese",
    "The chances of losing to an Otter in Blackjack are never zero",
    "Otters love DJ Khalid and can often be seen at concerts near the front",
    "Dairy Queen will be releasing a fish flavoured shake specifically for Otters in 2023",
    "Can you definitively confirm that none of your cousins are Otters...I thought not"
]
export function GetRandomOtterFact() {
    return otterFacts[Math.floor(Math.random() * otterFacts.length)]
}