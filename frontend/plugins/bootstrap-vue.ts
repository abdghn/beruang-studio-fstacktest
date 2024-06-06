import { defineNuxtPlugin } from '#app'
import { BootstrapVue3 } from 'bootstrap-vue-3'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

export default defineNuxtPlugin(nuxtApp => {
  nuxtApp.vueApp.use(BootstrapVue3)
})
