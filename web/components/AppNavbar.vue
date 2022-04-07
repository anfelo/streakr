<script lang="ts" setup>
  let isNavbarOpen = false;
  const handleBurgerClick = () => {
    isNavbarOpen = !isNavbarOpen;
  };
</script>

<template>
  <nav
    class="navbar is-light"
    role="navigation"
    aria-label="main navigation"
  >
    <div class="navbar-brand">
      <nuxt-link class="navbar-item" to="/">
        <span>Streakr</span>
      </nuxt-link>

      <span
        role="button"
        :class="`navbar-burger burger ${isNavbarOpen ? 'is-active' : ''}`"
        aria-label="menu"
        aria-expanded="false"
        data-target="mainNavbar"
        @click="handleBurgerClick"
      >
        <span aria-hidden="true" />
        <span aria-hidden="true" />
        <span aria-hidden="true" />
      </span>
    </div>

    <div
      id="mainNavbar"
      :class="`navbar-menu ${isNavbarOpen ? 'is-active' : ''}`"
    >
      <div class="navbar-end">
        <div class="navbar-item">
          <div class="buttons">
              <a class="button is-light">
                Login
              </a>
          </div>
        </div>
      </div>
    </div>
  </nav>
</template>
