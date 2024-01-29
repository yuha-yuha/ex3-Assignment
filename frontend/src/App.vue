<script setup lang="ts">
import { RouterLink, RouterView, useRouter } from 'vue-router'
import { useSessionStore } from './stores/Request';
const router = useRouter()
const sessionstore = useSessionStore()

</script>

<template>
  <v-app>
    <v-app-bar color="blue-lighten-1">
      <v-icon>mdi-home</v-icon>
      <v-app-var-title> 
        <v-toolbar-title>実験3ej-app</v-toolbar-title>
      </v-app-var-title>

      <v-spacer></v-spacer>

        <v-btn variant="flat" class="ma-1" v-if="!sessionstore.istokenValid()" @click="() => {router.push('/signup')}">Sign Up</v-btn>
        <v-btn variant="flat" class="ma-1" @click="() => {router.push('/login')}" v-if="!sessionstore.istokenValid()">Login</v-btn>
        <v-btn variant="flat" class="ma-1" v-if="sessionstore.istokenValid()">Profile</v-btn>

    </v-app-bar>
    <RouterView />
    <v-bottom-navigation >
      <v-btn @click="() => {router.push('/')}">home</v-btn>
    </v-bottom-navigation>
  </v-app>
</template>

<style scoped>
</style>
