import {OPTIONS_WRAPPER, USER_INPUT, NODE_TEXT, ERROR_MESSAGE} from "./gameScreenElements.js";
import {Typewriter, ScrollNodeTextToTop} from "./utils.js";

function setNewNodeText(str) {
    NODE_TEXT.innerHTML = str
    ScrollNodeTextToTop()
}

function setNewNodeOptions(gs, arr) {
    gs.setValidOptions(arr)
    while (OPTIONS_WRAPPER.hasChildNodes()) {
        OPTIONS_WRAPPER.removeChild(OPTIONS_WRAPPER.lastChild)
    }
    arr.forEach((opt) => {
        let el = document.createElement("p")
        el.classList.add("option")
        el.innerHTML = opt.Text
        OPTIONS_WRAPPER.append(el)
        Typewriter(opt.Text, el)
    })
    USER_INPUT.value = ""
}

export async function GetNodeAndAlignState(gs, nodeId) {
    if (nodeId !== null) {
        const body = {
            nextNode: nodeId,
            mood: gs.getStatus(),
            inventory: gs.getInventory()
        }
        const node = await axios.post("/api/gameNode", body)
        console.log(node.data)
        gs.setNodeID(node.data.ID)
        gs.setPoints(node.data.EarnedPoints)
        setNewNodeText(node.data.Text)
        setNewNodeOptions(gs, node.data.Options)
    }
}

export function ReturnNextNode(gs, userCommand, defaultOptions) {
    let nextNode = null
    let options = gs.getValidOptions()
    options = options.length > 0 ? options : defaultOptions
    for (let opt of options) {
        if (userCommand === opt.Command) {
            if (ERROR_MESSAGE.innerHTML !== "null") {
                ERROR_MESSAGE.innerHTML = "null"
            }
            nextNode = opt.NextNode
            if (nextNode === "1") {
                gs.resetState()
                return nextNode
            } else {
                if (opt.Mood.length > 0) {
                    gs.setStatus(opt.Mood)
                }
                if (opt.Inventory.length > 0) {
                    gs.addToInventory(opt.Inventory)
                }
            }
            return nextNode
        } else {
            ERROR_MESSAGE.innerHTML = "command not recognized..."
        }
    }
    return nextNode
}