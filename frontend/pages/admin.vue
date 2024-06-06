<template>
  <div class="container mt-4">
    <h2>Admin Dashboard</h2>
    <b-form @submit.prevent="sendInvitation">
      <b-form-group label="Email" label-for="email-input">
        <b-form-input
          id="email-input"
          v-model="email"
          type="email"
          placeholder="Enter email"
          required
        ></b-form-input>
      </b-form-group>
      <b-button type="submit" variant="primary">Send Invitation</b-button>
    </b-form>
    <b-list-group class="mt-4">
      <b-list-group-item v-for="invitation in invitations" :key="invitation.id">
        {{ invitation.email }} - {{ invitation.status }}
      </b-list-group-item>
    </b-list-group>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import axios from 'axios'

interface Invitation {
  id: number;
  email: string;
  status: string;
}

export default defineComponent({
  name: 'Admin',
  setup() {
    const email = ref('')
    const invitations = ref<Invitation[]>([])

    const sendInvitation = async () => {
      try {
        await axios.post('http://localhost:8080/admin/send-invitation', { email: email.value })
        fetchInvitations()
      } catch (error) {
        console.error(error)
      }
    }

    const fetchInvitations = async () => {
      try {
        const response = await axios.get('http://localhost:8080/admin/invitations')
        invitations.value = response.data
      } catch (error) {
        console.error(error)
      }
    }

    onMounted(() => {
      fetchInvitations()
    })

    return {
      email,
      invitations,
      sendInvitation
    }
  }
})
</script>
