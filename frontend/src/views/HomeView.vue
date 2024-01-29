<script setup lang="ts">
  import { RouterLink, RouterView, useRoute } from 'vue-router'
  import { useRouter } from 'vue-router';
  import { useWordStore } from '@/stores/Request';
import { ref } from 'vue';
  import { useSessionStore } from '@/stores/Request';


  const wordStore = useWordStore()
  const router = useRouter()
  const sessionstore = useSessionStore()

  async function navWordQuiz() 
  {
    try {
    await wordStore.AllUserDefineWordListFetch()
    router.push('/quiz')
    }catch(error) {
      console.log(error)
    }
  }
</script>

<template>

  <v-row justify="center" align-content="center" style="height: 40%;">
    <v-col cols="5"> 
      <v-btn block style="min-height: 200px;" @click="navWordQuiz">クイズを解く</v-btn> 
    </v-col>
    <v-col cols="5"> <v-btn v-bind:disabled="!sessionstore.istokenValid()" block style="min-height: 200px;" @click="() => {$router.push('/quiz/new')}">クイズを作る</v-btn>  </v-col>
  </v-row>
</template>
