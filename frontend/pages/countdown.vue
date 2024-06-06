<template>
  <div class="container mt-4">
    <h2>Countdown to Election</h2>
    <CountdownTimer />
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import CountdownTimer from '~/components/CountdownTimer.vue'

export default defineComponent({
  name: 'CountdownPage',
  components: {
    CountdownTimer
  }
})
</script>
