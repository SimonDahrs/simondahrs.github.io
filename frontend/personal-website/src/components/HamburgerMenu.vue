<script setup>
import { ref } from 'vue'
import NavigationBar from './NavigationBar.vue';
const menuOpened = ref(false)

</script>
<template>
    <div id="hamburger-menu">
        <div id="hamburger-toggle" @click="menuOpened = !menuOpened" :class="{ opened: menuOpened }">
            <span class="button-icon" :class="{ opened: menuOpened }" id="openMenu">☰</span>
            <span class="button-icon" :class="{ opened: menuOpened }" id="closeMenu">✕</span>
        </div>
        <div id="hamburger-menu-item-container" :class="{ opened: menuOpened }">
            <NavigationBar />
        </div>
    </div>
</template>
<style>
#openMenu {
  opacity: 1;
  transition: transform .2s ease-in-out;
  transform: scale(1) rotate(0);
}

#openMenu.opened {
  transform: scale(0) rotate(90deg);
}

#closeMenu {
  transition: transform .2s ease-in-out;
  transform: scale(0) rotate(-90deg);
}

.button-icon {
  position: absolute;
  width: 50px;
  text-align: center;
}

#closeMenu.opened {
  opacity: 1;
  transform: scale(1) rotate(0);
}
#hamburger-menu {
    display: none;
    position: fixed;
    top: .5rem;
    left: .5rem;
    z-index: 5;
    display: none;
    row-gap: 1em;
}

#hamburger-toggle {
    font-weight: bold;
    font-size: 24px;
    box-sizing: border-box;
    background: rgba(var(--vt-c-accent-light), 0.5);
    border-radius: 50%;
    display: flex;
    justify-content: center;
    height: 50px;
    width: 50px;
    text-align: center;
    align-items: center;
    cursor: pointer;
    transition: transform .3s ease-out;
    box-shadow:
        0.1px 0.1px 0.3px rgba(0, 0, 0, 0.02),
        0.1px 0.1px 0.8px rgba(0, 0, 0, 0.028),
        0.3px 0.3px 1.5px rgba(0, 0, 0, 0.035),
        0.4px 0.4px 2.7px rgba(0, 0, 0, 0.042),
        0.8px 0.8px 5px rgba(0, 0, 0, 0.05),
        2px 2px 12px rgba(0, 0, 0, 0.07);
}

#hamburger-toggle:hover {
    background: rgb(var(--vt-c-accent));
    transform: scale(1.1);
    box-shadow:
        0.4px 0.4px 0.3px rgba(0, 0, 0, 0.02),
        1px 1px 0.8px rgba(0, 0, 0, 0.028),
        1.9px 1.9px 1.5px rgba(0, 0, 0, 0.035),
        3.4px 3.4px 2.7px rgba(0, 0, 0, 0.042),
        6.3px 6.3px 5px rgba(0, 0, 0, 0.05),
        15px 15px 12px rgba(0, 0, 0, 0.07);
}

#hamburger-toggle.opened {
    background: rgba(199, 217, 254, 1);
}

#hamburger-menu-item-container {
    transform: scaleY(0);
    background-color: var(--vt-c-background-right);
    z-index: 5;
    box-shadow: 0 3px 10px rgb(0 1 1 / 0.3);
    border: 1px solid rgba(200, 200, 200, 0.3);
    border-radius: 7px;
    transition: transform .15s;
    transform-origin: top;
}

#hamburger-menu-item-container.opened {
    transform: scaleY(1);
}

#hamburger-menu-item-container a {
    color: var(--vt-c-text-right);
}

@media only screen and (max-width: 768px) {

    #hamburger-menu {
        display: grid;
    }

}
</style>