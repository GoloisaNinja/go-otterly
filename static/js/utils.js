import {MENU_DROPDOWNS, BODY, GAME_WRAPPER, MODAL_OVERLAY_WRAPPER, MODAL_WRAPPER, THEME_TOGGLE_BTN} from "./screenElements.js";

export function Typewriter(str, el, speed = 125) {
    let textPosition = 0;
    function typeEffect() {
        el.innerHTML = str.substring(0, textPosition);
        if (textPosition !== str.length + 1) {
            textPosition++;
            setTimeout(() => {
                typeEffect();
            }, speed);
        }
    }
    typeEffect();
}

export const Quotes = [
    `"Otterly has improved my day immensly. No matter what, at least I know I didn't create this trash!" - Barnaby, UK`,
    `"The great thing about Otterly is that there is so little that's great about it!" - Lowie, CA`,
    `"If there's one thing, I wish I had less of, it's definitely Otterly, but in a good way." - Ash, LV426`,
    `"Sometimes you just need a good text adventure. After visiting Otterly, I'm still searching." - Martin, MA`,
    `"Playing Otterly games is a bit like meeting that new friend that kinda scares you with their crazy." - Lou, QLD`,
    `"I wasn't sure about Otterly. Then I played a game. Now I'm sure. It's just as I thought. Trash." - Marc, NY`,
    `"It's called Otterly Ridiculous Games. There's otters. There's ridiculous. Games is a stretch." - Geralt, Rivia`,
    '"Otterly is like an ice cream sundae for the soul, with candy toppings, so your soul is like, extra tasty." - Willy, KY',
    `"I like to relax with a good text adventure game. Otterly is not relaxing. Or good. There is text." - Samus, Brinstar`,
]

export function ReturnLoadingTexts(gameTitle) {
    const texts = [
        "BOOT SEQUENCE INITIATED...",
        "CHECKING RAM...",
        "64K RAM SYSTEM 38911 BASIC BYTES FREE",
        "CHECKING HARD DISK...",
        "HARD DISK FOUND...",
        "LOADING SYSTEM BIOS...",
        "BIOS LOADED SUCCESSFULLY",
        "OTTER OS VERSION 1.4 FOUND",
        "JOYSTICK PORT NOT DETECTED...",
        "ENGAGING RIDICULOUSNESS ENGINE...",
        "MAX RIDICULOUSNESS ACHIEVED",
        `${gameTitle.toUpperCase()} REQUESTED...`,
        "RESOLVING OTTERS...",
        "READY...",
    ]
    return texts;
}


const menuItemMap = new Map([
    ["Game-drop-content", ["File-drop-content", "Inventory-drop-content"]],
    ["File-drop-content", ["Game-drop-content", "Inventory-drop-content"]],
    ["Inventory-drop-content", ["Game-drop-content", "File-drop-content"]]
])

export function ToggleMenuShow(id) {
    const clickedElement = document.getElementById(id)
    const nonClickedElements = menuItemMap.get(id)
    for (let elId of nonClickedElements) {
        let el = document.getElementById(elId)
        if (el.classList.contains("show")) {
            el.classList.toggle("show")
        }
    }
    clickedElement.classList.toggle("show")
}

export function CheckMenuDropDowns() {
    for (let dd of MENU_DROPDOWNS) {
        if (dd.classList.contains("show")) {
            dd.classList.toggle("show")
        }
    }
}
export function ThemeToggle() {
    let gt = localStorage.getItem("game-theme")
    if (gt === "theme-dark") {
        localStorage.setItem("game-theme", "theme-light")
        BODY.classList.remove("theme-dark")
        BODY.classList.add("theme-light")
        THEME_TOGGLE_BTN.innerHTML = "light theme"
    } else {
        localStorage.setItem("game-theme", "theme-dark")
        BODY.classList.remove("theme-light")
        BODY.classList.add("theme-dark")
        THEME_TOGGLE_BTN.innerHTML = "dark theme"
    }
}
export function ScrollNodeTextToTop() {
    const yPos = GAME_WRAPPER.getBoundingClientRect().top
    const yOffset = yPos + window.scrollY
    window.scroll({left: 0, top: 0, behavior: "smooth"})
}
export function InventoryModalDismiss() {
    MODAL_WRAPPER.classList.remove("show")
    MODAL_OVERLAY_WRAPPER.classList.remove("show")
}
export function HandleGameCardClick(element, id) {
    let spinnerEls = document.querySelectorAll(`#spinner${id}`)
    spinnerEls.forEach((el) => el.classList.add('isActive'))
    element.click()
}