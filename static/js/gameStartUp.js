import {CheckMenuDropDowns, ToggleMenuShow, Typewriter, ThemeToggle, InventoryModalDismiss} from "./utils.js";
import {BODY,GAME_AREA,MENU_BUTTONS,THEME_TOGGLE_BTN, RESET_BUTTON,GAME_OPTIONS,INVENTORY_MENU,STATUS,POINTS, MODAL_DISMISS_BUTTON} from "./screenElements.js";
import {GetNodeAndAlignState} from "./gameLogic.js";
import {ExecuteLoadingScreen} from "./loadingScreen.js";

export function GameStartUp(gs, gameTitle, totalNodes) {
    // IIFE to check/set theme from local storage
    (function() {
        if (localStorage.getItem("game-theme") !== null) {
            let gt = localStorage.getItem("game-theme")
            if (gt === "theme-light") {
                BODY.classList.add("theme-light")
            } else {
                BODY.classList.add("theme-dark")
            }
        } else {
            BODY.classList.add("theme-dark")
            localStorage.setItem("game-theme", "theme-dark")
        }
        const t = localStorage.getItem("game-theme").split("-")
        THEME_TOGGLE_BTN.innerHTML = `${t[1]} theme`
    })()
    // Get inventory, status and points from Game State
    const inv = gs.getInventory()
    const status = gs.getStatus()
    const points = gs.getPoints()
    // Set the Game Title and Total Nodes
    gs.setGameTitle(gameTitle)
    gs.setTotalNodes(totalNodes)
    // Run the Loading Screen Func Animation
    ExecuteLoadingScreen(gameTitle) // Comment this out to bypass loading during dev
    // Set event listeners for the game area and the menu buttons
    GAME_AREA.addEventListener("click", CheckMenuDropDowns)
    MENU_BUTTONS.forEach((node) => {
        node.addEventListener("click", (e) => {
            let clickedId = e.target.innerHTML + "-drop-content"
            ToggleMenuShow(clickedId)
        })
    })
    THEME_TOGGLE_BTN.addEventListener("click", ThemeToggle)
    RESET_BUTTON.addEventListener("click", async () => {
        CheckMenuDropDowns()
        await GetNodeAndAlignState(gs, "1")
    }, false)
    // Set the listener for the modal dismiss button
    MODAL_DISMISS_BUTTON.addEventListener("click", InventoryModalDismiss)
    // Run the Typewriter Animation on all initial options elements
    GAME_OPTIONS.forEach((opt) => {
        let str = opt.childNodes[0].nodeValue
        Typewriter(str, opt, 50)
    })
    // Sets initial Inventory
    const initialInventory = document.createElement("p")
    initialInventory.innerHTML = inv
    INVENTORY_MENU.appendChild(initialInventory)
    // Sets initial Status
    STATUS.innerHTML = status
    // Set initial points
    POINTS.innerHTML = points
}