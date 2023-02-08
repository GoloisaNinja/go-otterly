import { ReturnLoadingTexts } from './utils.js'
import {LOADING_SCREEN, LOADING_TEXT_WRAPPER, GAME_WRAPPER} from "./gameScreenElements.js";

export function ExecuteLoadingScreen(gameTitle) {
    const lt = ReturnLoadingTexts(gameTitle)
    let delay = 0
    for (let i = 0; i < lt.length; i++) {
        let childEl = document.createElement("p")
        childEl.innerHTML = lt[i]
        setTimeout(() => {
            LOADING_TEXT_WRAPPER.appendChild(childEl)
        }, delay)
        delay += 1000
    }
    setTimeout(() => {
        LOADING_SCREEN.style.display = "none"
        GAME_WRAPPER.style.display = "flex"
    }, 15000)
}