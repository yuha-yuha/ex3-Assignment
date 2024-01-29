<script setup lang="ts">
    import { ref } from 'vue';
    import { useSessionStore } from '@/stores/Request';
import router from '@/router';

    const sessionStore = useSessionStore()

    const inputEmail = ref('') 
    const inputPassword = ref('')
    const inputName = ref('')
    const faildLogin = ref('')

    async function Submit() {
        try {
            await sessionStore.SignUp(inputEmail.value, inputPassword.value, inputName.value)
            router.push('/')
        }catch(error) {
            faildLogin.value = '失敗w' + error
        };

    }
</script>

<template>
    <v-app>
        <v-card width="80vw" class="login-block">
            <v-card-title >
                SignUp
            </v-card-title>

            <v-text-field 
                v-model="inputEmail"
                label="email"
                variant="outlined">
            </v-text-field>
            <v-text-field
                v-model="inputName"
                label="name"
                variant="outlined">
            </v-text-field>

            <v-text-field 
                v-model="inputPassword"
                type="password"
                label="password"
                variant="outlined">
                
            </v-text-field>

            

            <v-btn @click="Submit">送信</v-btn>
            {{ faildLogin }}
        </v-card>
      </v-app>
</template>

<style>
.login-block {
    margin-top: 120px;
}
</style>