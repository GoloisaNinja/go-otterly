import {INVENTORY_MENU, POINTS, STATUS} from "./gameScreenElements.js";

const initialState = {
    nodeID: null,
    nodesCompleted: 1,
    validOptions: [],
    status: "unknown",
    inventory: ["empty"],
    points: 0,
}

export function gameState() {

    let state = {...initialState}

    function getNodeID() {
        return state.nodeID
    }

    function setNodeID(id) {
        state.nodeID = id
    }

    function getNodesCompleted() {
        return state.nodesCompleted
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

    function resetState() {
        state = initialState
        setStatus(state.status)
        resetInventory()
    }

    return {
        getNodeID,
        setNodeID,
        getValidOptions,
        setValidOptions,
        getStatus,
        setStatus,
        getInventory,
        addToInventory,
        getPoints,
        setPoints,
        resetState
    }
}