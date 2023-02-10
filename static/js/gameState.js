import {INVENTORY_MENU, POINTS, STATUS} from "./screenElements.js";

const initialState = {
    nodeID: null,
    gameTitle: "",
    totalNodes: 0,
    nodesCompleted: 1,
    validOptions: [],
    status: "unknown",
    inventory: ["empty"],
    points: 0,
    playType: new Map([
        ["passive", 0],
        ["intelligent", 0],
        ["aggressive", 0],
        ["reckless", 0],
    ]),
    storyArc: "undetermined...",
    survived: true,
    codedNode: {
        codeNode: false,
        codeLength: 0,
        codeFailedNextNode: 0,
    }
}

export function gameState() {

    let state = {...initialState}

    function getNodeID() {
        return state.nodeID
    }

    function setNodeID(id) {
        state.nodeID = id
    }

    function getGameTitle() {
        return state.gameTitle
    }

    function setGameTitle(gt) {
        state.gameTitle = gt
    }

    function getTotalNodes() {
        return state.totalNodes
    }

    function setTotalNodes(tn) {
        state.totalNodes = tn
    }

    function getNodesCompleted() {
        return state.nodesCompleted
    }

    function increaseNodesCompleted() {
        state.nodesCompleted++
    }

    function increaseNodesCompleted() {
        state.nodesCompleted++
    }

    function getValidOptions() {
        return state.validOptions
    }

    function setValidOptions(arr) {
        state.validOptions = arr
    }

    function getStatus() {
        return state.status
    }

    function setStatus(s) {
        state.status = s
        STATUS.innerHTML = s
    }

    function getInventory() {
        return state.inventory
    }

    function addToInventory(item) {
        let ni = document.createElement("p")
        ni.innerHTML = item
        if (state.inventory.includes("empty")) {
            state.inventory = [item]
            INVENTORY_MENU.lastChild.remove()
        } else {
            state.inventory = [...state.inventory, item]
        }
        INVENTORY_MENU.appendChild(ni)

    }
    function resetInventory() {
        while(INVENTORY_MENU.lastChild) {
            INVENTORY_MENU.lastChild.remove()
        }
        let defaultInventory = document.createElement("p")
        defaultInventory.innerHTML = "empty"
        INVENTORY_MENU.appendChild(defaultInventory)
    }

    function getPoints() {
        return state.points
    }

    function setPoints(p) {
        state.points += p
        POINTS.innerHTML = state.points
    }

    function recordPlayType(playType) {
        let ptCurrentValue = state.playType.get(playType)
        ptCurrentValue += 1
        state.playType.set(playType, ptCurrentValue)
    }

    function getDominantPlayType() {
        let dominantPlayType = "passive"
        let dominantPTValue = 0;
        for (const [key, value] of state.playType) {
            if (value > dominantPTValue) {
                dominantPlayType = key
                dominantPTValue = value
            }
        }
        return dominantPlayType
    }

    function getStoryArc() {
        return state.storyArc
    }

    function setStoryArc(storyArc) {
        state.storyArc = storyArc
    }

    function getSurvived() {
        return state.survived
    }

    function setSurvived() {
        state.survived = false
    }

    function getCodedNode() {
        return state.codedNode
    }

    function setCodedNode(cn, cl, cnn) {
        state.codedNode = {
            codeNode: cn,
            codeLength: cl,
            codeFailedNextNode: cnn,
        }
    }


    function resetState(gt, tn) {
        state = {...initialState, gameTitle: gt, totalNodes: tn}
        setStatus(state.status)
        resetInventory()
    }

    return {
        getNodeID,
        setNodeID,
        getGameTitle,
        setGameTitle,
        getTotalNodes,
        setTotalNodes,
        getNodesCompleted,
        increaseNodesCompleted,
        getValidOptions,
        setValidOptions,
        getStatus,
        setStatus,
        getInventory,
        addToInventory,
        getPoints,
        setPoints,
        recordPlayType,
        getDominantPlayType,
        getStoryArc,
        setStoryArc,
        getSurvived,
        setSurvived,
        getCodedNode,
        setCodedNode,
        resetState
    }
}