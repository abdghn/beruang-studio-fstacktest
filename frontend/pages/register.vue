<template>
  <div class="container mt-4">
    <h2>Register</h2>
    <b-form @submit.prevent="submitForm">
      <b-form-group label="Email" label-for="email-input">
        <b-form-input
          id="email-input"
          v-model="email"
          type="email"
          placeholder="Enter email"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group label="Nama" label-for="name-input">
        <b-form-input
          id="name-input"
          v-model="nama"
          type="text"
          placeholder="Enter name"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group label="Tanggal Lahir" label-for="dob-input">
        <b-form-input
          id="dob-input"
          v-model="tanggalLahir"
          type="date"
          required
        ></b-form-input>
      </b-form-group>
      <b-form-group label="Jenis Kelamin" label-for="gender-select">
        <b-form-select
          id="gender-select"
          v-model="jenisKelamin"
          :options="[{ value: 'male', text: 'Male' }, { value: 'female', text: 'Female' }]"
          required
        ></b-form-select>
      </b-form-group>
      <b-form-group label="Alamat Lengkap" label-for="address-textarea">
        <b-form-textarea
          id="address-textarea"
          v-model="alamatLengkap"
          placeholder="Enter full address"
          required
        ></b-form-textarea>
      </b-form-group>
      <b-button type="submit" variant="primary">Register</b-button>
    </b-form>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import axios from 'axios'

export default defineComponent({
  name: 'Register',
  setup() {
    const email = ref('')
    const nama = ref('')
    const tanggalLahir = ref('')
    const jenisKelamin = ref('')
    const alamatLengkap = ref('')

    const submitForm = async () => {
      try {
        const response = await axios.post('http://localhost:8080/user/register', {
          email: email.value,
          nama: nama.value,
          tanggalLahir: tanggalLahir.value,
          jenisKelamin: jenisKelamin.value,
          alamatLengkap: alamatLengkap.value,
        })
        console.log(response.data)
        alert(`Registration successful. Your registration code is: ${response.data.registrationCode}`)
      } catch (error) {
        console.error(error)
      }
    }

    return {
      email,
      nama,
      tanggalLahir,
      jenisKelamin,
      alamatLengkap,
      submitForm
    }
  }
})
</script>
