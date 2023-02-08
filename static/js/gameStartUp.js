import {CheckMenuDropDowns, ToggleMenuShow, Typewriter, ThemeToggle} from "./utils.js";
import {BODY,GAME_AREA,MENU_BUTTONS,THEME_TOGGLE_BTN, RESET_BUTTON,GAME_OPTIONS,INVENTORY_MENU,STATUS,POINTS} from "./gameScreenElements.js";
import {GetNodeAndAlignState} from "./gameLogic.js";
import {ExecuteLoadingScreen} from "./loadingScreen.js";

export function GameStartUp(gs, gameTitle) {
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
    })()
    // Get inventory, status and points from Game State
    const inv = gs.getInventory()
    const status = gs.getStatus()
    const points = gs.getPoints()
    // Run the Loading Screen Func Animation
    ExecuteLoadingScreen(gameTitle)
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
        gs.resetState()
        CheckMenuDropDowns()
        await GetNodeAndAlignState(gs, "1")
    }, false)
    // Run the Typewriter Animation on all initial options elements
    GAME_OPTIONS.forEach((opt) => {
        let str = opt.childNodes[0].nodeValue
        Typewriter(str, opt)
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