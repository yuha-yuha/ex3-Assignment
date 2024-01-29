<script setup lang="ts">
    import { ref } from 'vue';
    import { useWordStore, useSessionStore } from '@/stores/Request';
    import router from '@/router';
import { storeToRefs } from 'pinia';

    const wordStore = useWordStore()
    const sessionStore = useSessionStore() 

    const {token} = storeToRefs(sessionStore)

    const Japanese = ref('') 
    const English = ref('')
    const faildLogin = ref('')

    async function Submit() {
        try {
            await wordStore.createWord(Japanese.value, English.value, token.value)
            faildLogin.value='送信かんりょう！！！'
            Japanese.value = ''
            English.value = ''
            
        }catch(error) {
            faildLogin.value = '失敗!ｗｗｗｗ err:' + error
        };

    }
</script>

<template>
    <v-app>
        <v-card width="80vw" class="login-block">
            <v-card-title >
                クイズ作成
            </v-card-title>

            <v-text-field 
                v-model="Japanese"
                label="日本語"
                variant="outlined">
            </v-text-field>

            <v-text-field 
                v-model="English"
                label="英語"
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