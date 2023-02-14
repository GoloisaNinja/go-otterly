import {
    ERROR_MESSAGE,
    MODAL_INVENTORY_DYNAMIC,
    MODAL_OVERLAY_WRAPPER,
    MODAL_WRAPPER,
    NODE_TEXT,
    OPTIONS_WRAPPER,
    USER_INPUT
} from "./screenElements.js";
import {ScrollNodeTextToTop, Typewriter} from "./utils.js";
import {GetRandomOtterFact} from "./otterFacts.js";

function showInventoryModal(inv) {
    MODAL_INVENTORY_DYNAMIC.innerHTML = inv
    MODAL_WRAPPER.classList.add("show")
    MODAL_OVERLAY_WRAPPER.classList.add("show")
}

function setNewNodeText(str) {
    NODE_TEXT.innerHTML = str
    setTimeout(() => {
        ScrollNodeTextToTop()
    }, 300)

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
        Typewriter(opt.Text, el, 50)
    })
    USER_INPUT.removeAttribute("disabled")
    USER_INPUT.value = ""
    USER_INPUT.focus({preventScroll: true})
}

function HandleAfterAction(gs) {
    const option = [{
        ID: "afterAction-999",
        Text: `Type "s" to Start Over`,
        Command: "s",
        Mood: "",
        HasReqs: false,
        Inventory: "",
        NextNode: "1",
    }]
    const afterActionReport =
        `****AFTER ACTION REPORT****
            
        
            GAME PLAYED: ${gs.getGameTitle()}
        
            NODES COMPLETED: ${gs.getNodesCompleted()} / ${gs.getTotalNodes()}
        
            SELECTED STORY ARC: ${gs.getStoryArc()}
        
            DOMINANT PLAY TYPE: ${gs.getDominantPlayType()}
        
            SURVIVED ENCOUNTER: ${gs.getSurvived()}
        
            TOTAL POINTS EARNED: ${gs.getPoints()}
            
            DID YOU KNOW?:
            ${GetRandomOtterFact()}
        
        
            ****THANK YOU FOR PLAYING****`

    setNewNodeText(afterActionReport)
    setNewNodeOptions(gs, option)
}

export async function GetNodeAndAlignState(gs, nodeId) {
    const body = {
        nextNode: nodeId,
        status: gs.getStatus(),
        inventory: gs.getInventory()
    }
    // reset codednode state if it's in a true position
    if (gs.getCodedNode().codeNode) {
        gs.setCodedNode(false, 0, 0)
    }
    // reset all state if nodeId === 1
    if (nodeId === "1") {
        gs.resetState(gs.getGameTitle(), gs.getTotalNodes())
    }
    USER_INPUT.value = "...computing input"
    USER_INPUT.setAttribute("disabled", "true")
    const node = await axios.post("/api/gameNode", body)
    gs.setNodeID(node.data.ID)
    gs.setPoints(node.data.EarnedPoints)
    gs.increaseNodesCompleted()
    if (node.data.CodeNode) {
        gs.setCodedNode(true, node.data.CodeLength, node.data.CodeFailedNextNode)
    }
    setNewNodeText(node.data.Text)
    setNewNodeOptions(gs, node.data.Options)
}

function goodUserInput(options, userCommand, cns) {
    const {codeNode, codeLength, codeFailedNextNode} = cns
    const inputResult = {
        goodInput: true,
        optionIndex: -1,
        nextNode: null,
    }
    if (codeNode) {
        if (codeLength !== userCommand.length) {
            inputResult.goodInput = false
            return inputResult
        }
        if (userCommand === options[0].Command) {
            inputResult.optionIndex = 0
            return inputResult
        } else {
            inputResult.nextNode = codeFailedNextNode
            return inputResult
        }
    }
    for (let i = 0; i < options.length; i++) {
        if (userCommand === options[i].Command) {
            inputResult.optionIndex = i
            return inputResult
        }
    }
    inputResult.goodInput = false
    return inputResult
}

export function ReturnNextNode(gs, userCommand, defaultOptions) {
    let options = gs.getValidOptions()
    options = options.length > 0 ? options : defaultOptions
    const {goodInput, optionIndex, nextNode} = goodUserInput(options, userCommand, gs.getCodedNode())
    // Step 1: is the user input any good? Handle ERROR_MESSAGE accordingly
    if (!goodInput) {
        ERROR_MESSAGE.innerHTML = "command not recognized..."
        return null
    } else {
        if (ERROR_MESSAGE.innerHTML !== "null") {
            ERROR_MESSAGE.innerHTML = "null"
        }
    }
    // Step 2: determine if a codednode/nextnode is occurring
    // if it is reset codednode state and return the failedNextNode to axios call
    if (nextNode !== null) {
        return nextNode.toString()
    }
    const uo = options[optionIndex]
    // Step 3: is the option select an AfterAction Option - if so
    // we need to display the after action and not make an axios call to the server
    if (uo.AfterAction) {
        HandleAfterAction(gs)
        return null
    }
    // Step 4: user input is good, option chosen is not an after action
    // at this point we can process Mood/Status, Playtype, DeathNode, Inventory,
    // Points, and possible StoryArc and return the nextNode for axios call

    // DeathNode Check
    if (uo.DeathNode) {
        gs.setSurvived()
    }
    // Mood Check
    if (uo.Mood.length > 0) {
        gs.setStatus(uo.Mood)
    }
    // Inventory Check
    if (uo.Inventory.length > 0) {
        gs.addToInventory(uo.Inventory)
        showInventoryModal(uo.Inventory)
    }
    // PlayType Check
    if (uo.PlayType !== undefined && uo.PlayType.length > 0) {
        gs.recordPlayType(uo.PlayType)
    }
    // StoryArc Check
    if (uo.StoryArc !== undefined && uo.StoryArc.length > 0) {
        gs.setStoryArc(uo.StoryArc)
    }

    return uo.NextNode
}