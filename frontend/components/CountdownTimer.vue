<template>
    <b-card class="text-center">
      <b-card-text>
        <h3>{{ days }} Days {{ hours }} Hours {{ minutes }} Minutes {{ seconds }} Seconds</h3>
      </b-card-text>
    </b-card>
  </template>
  
  <script lang="ts">
  import { defineComponent, ref, onMounted } from 'vue'
  
  export default defineComponent({
    name: 'CountdownTimer',
    setup() {
      const days = ref(0)
      const hours = ref(0)
      const minutes = ref(0)
      const seconds = ref(0)
  
      const updateCountdown = () => {
        const targetDate = new Date('2024-06-01T00:00:00').getTime()
        const now = new Date().getTime()
        const distance = targetDate - now
  
        days.value = Math.floor(distance / (1000 * 60 * 60 * 24))
        hours.value = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60))
        minutes.value = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60))
        seconds.value = Math.floor((distance % (1000 * 60)) / 1000)
      }
  
      onMounted(() => {
        updateCountdown()
        setInterval(updateCountdown, 1000)
      })
  
      return {
        days,
        hours,
        minutes,
        seconds
      }
    }
  })
  </script>
  
  <style scoped>
  /* Additional styles if needed */
  </style>
  