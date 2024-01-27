<script setup lang="ts">
    import type { Ref } from 'vue';
    import { ref } from 'vue';
    import { useWordStore, type Word } from '@/stores/Request';
    import { storeToRefs } from 'pinia';
    import type { popScopeId } from 'vue';
    import { onMounted } from 'vue';

    const wordStore = useWordStore()
    const {wordList} = storeToRefs(wordStore)

    const UserAnswer = ref('')
    const word = ref<Word>({japanese:'', english:'', user_id:0})
    //interface function isWordListEmpty() {

    //}

    function PopWord() {
        if (wordList.value.length !== 0) {
        const NewWord = wordList.value.shift() as Word
        word.value = NewWord
        }
    } 

    function JudgeAnswer() {
        if (UserAnswer.value === word.value?.english) {
            if (wordList.value.length === 0) {
                //終了画面に移動
            } else {
                PopWord()
                
            }
        }
        UserAnswer.value = ''
    }

    onMounted(() => {
        PopWord()
    })
</script>

<template>
    <v-row
        justify="center" align-content="center">
        <v-col cols="6" lg="4" class="text-center">
                <v-card color="grey-lighten-3" width="600px" height="350px" style="margin-top: 80px;" text-color="grey lighten-1">
                    <v-card-text class="center-text text-h4 d-flex align-center justify-center">{{ word.japanese }}</v-card-text>
                </v-card>
        </v-col>

    </v-row>

    <v-row
        class="blue lighten-4" style="height: 25vh"
        justify="center" align-content="center">
  <v-col cols="12" sm="6" md="7" class="text-center">
    <v-text-field
        name="ans"
        label="回答"
        id="ans"
        clearable
        variant="outlined"
        autofocus
        v-model="UserAnswer"
        @keydown.enter="JudgeAnswer"
    ></v-text-field>
  </v-col>
</v-row>
</template>

<style>
.center-text {
    text-align: center;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 100%;
  }
</style>