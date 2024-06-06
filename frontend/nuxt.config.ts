export default defineNuxtConfig({
  ssr: false,
  target: 'static',
  head: {
    title: 'Quick Count Pemilu',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
    ],
    link: [{ rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }],
  },
  css: [
    'bootstrap/dist/css/bootstrap.css',
    'bootstrap-vue/dist/bootstrap-vue.css'
  ],
  plugins: [
    '~/plugins/bootstrap-vue.ts'
  ],
  // components: true,
  buildModules: [
    '@nuxt/typescript-build',
    '@pinia/nuxt',
  ],
  modules: [],
  build: {
    transpile: ['bootstrap-vue']
  },
})
