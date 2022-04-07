<script setup lang="ts">

</script>

<template>
  <aside class="section">
    <div class="login-form">
      <form>
        <div>
          <h3 class="title has-text-centered">Login</h3>
        </div>

        <div class="field">
          <label for="email" class="label">Email</label>
          <div class="control has-icons-left">
            <input
              id="email"
              class="input"
              type="text"
              placeholder="Enter your email"
            />
            <span class="icon is-small is-left">
              <i class="fas fa-user"></i>
            </span>
          </div>
        </div>

        <div class="field">
          <label for="password" class="label">Password</label>
          <div class="control has-icons-left has-icons-right">
            <input
              id="password"
              class="input"
              type="text"
              placeholder="Enter your password"
            />
            <span class="icon is-small is-left">
              <i class="fas fa-lock"></i>
            </span>
            <span class="icon is-small is-right">
              <i class="fas fa-eye"></i>
            </span>
          </div>
        </div>

        <div class="field">
          <div class="control">
            <label class="checkbox">
              <input
                type="checkbox"
              />
              Show password
            </label>
          </div>
        </div>

        <div class="control">
          <button
            class="button is-primary"
          >
            Login
          </button>
        </div>

      </form>
    </div>
  </aside>
</template>

<style lang="scss" scoped>
@media (min-width: 768px) {
  .login-form {
    width: 600px;
    margin: auto;
  }
}
</style>