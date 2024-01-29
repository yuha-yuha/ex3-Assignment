import { ref, computed, type Ref } from 'vue'
import { defineStore } from 'pinia'
import { id } from 'vuetify/locale'
import axios from 'axios'
import UsersQuizesVue from '@/views/UsersQuizes.vue'


export type User = {
  id: number
  name: string
}

export type Word = {
  japanese: string
  english: string
  user_id: number
}


export const useSessionStore = defineStore('session', () => {

  const token = ref('')
  async function login(email:string,password:string) {
    const data = {
      email: email,
      password: password
    }
    try {
      const requestPath = `http://localhost:8080/api/login`
      const response = await axios.post(requestPath, data, {
        withCredentials: true 
      })
      console.log(response)
      token.value = response.data.token
      console.log(token)
    }
    catch(error) {
      console.log(error)
      throw error
    }
  }

  async function SignUp(email:string, password:string, name:string) {
    const data = {
      email: email,
      name: name,
      password: password,
    }

    try {
      const requestPath = `http://localhost:8080/api/signup`
      const response = await axios.post(requestPath, data)
    }catch(error) {
      console.log(error)
      throw error
    }
  }

  function istokenValid(): boolean {
    if (token.value.length !== 0) {
      return true
    }else{
      return false
    }
  }

  return {login, token, istokenValid, SignUp}
}) 
export const useUserStore = defineStore('user', () => {
  const user = ref<User>({id: 0, name: ''}) as Ref<User>
  async function userFetch(id:number) {
    try {
      const requestPath = `http://localhost:8080/api/users/${id}`
      const response = await axios.get(requestPath)
      console.log(response.data)

      user.value.id = response.data.ID
      user.value.name = response.data.user_name
    }
    catch (error) {
      throw error
    }
  }

  return { userFetch, user }
})

export const useWordStore = defineStore('word', () => {
  const wordList = ref<Array<Word>>([{japanese:'', english:'', user_id:0}]) as Ref<Array<Word>>

  async function AllUserDefineWordListFetch() {
    try {
      const reqestPath = `http://localhost:8080/api/words`
      const response = await axios.get<Array<Word>>(reqestPath)
      wordList.value = response.data
    }
    catch (error) {
      throw error
    }
  }

  async function createWord(japanese:string, english:string, token: string) {

    const data = {
      japanese: japanese,
      english: english,
    }
    try {
      const reqestPath = `http://localhost:8080/api/auth/words`
      const response = await axios.post(reqestPath,data, {
        params: {
          token: token,
        }
      })
    }
    catch (error) {
      throw error
    }
  }

  return {AllUserDefineWordListFetch, wordList, createWord}
})