import { defineNuxtConfig } from 'nuxt3'

// https://v3.nuxtjs.org/docs/directory-structure/nuxt.config
export default defineNuxtConfig({
  vite: {
    css: {
      preprocessorOptions: {
        sass: {
          additionalData: '@import "@/assets/styles/global.scss"',
        }
      }
    }
  },

  app: {
    head: {
      title: 'streakr',
      htmlAttrs: {
        lang: 'en',
      },
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { hid: 'description', name: 'description', content: '' },
        { name: 'format-detection', content: 'telephone=no' },
      ],
      link: [
        { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
        {
          rel: "stylesheet",
          href:
            "https://fonts.googleapis.com/css2?family=Fira+Mono&family=Montserrat&family=Poppins:wght@700&display=swap"
        }
      ],
      script: [
        {
          hid: "fontawesome",
          src: "https://use.fontawesome.com/releases/v5.3.1/js/all.js",
          defer: true
        }
      ],
    },
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: ["@/assets/styles/global.scss"],
})
