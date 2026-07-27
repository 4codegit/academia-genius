<script setup lang='ts'>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuth } from '~/composables/useAuth'
const isLogin = ref(true)
const isForgot = ref(false)
const email = ref('')
const password = ref('')
const confirmPassword = ref('')
const username = ref('')
const fullName = ref('')
const errorMsg = ref('')
const successMsg = ref('')
const loading = ref(false)
const router = useRouter()
const { login, register, forgotPassword } = useAuth()
const submit = async () => {
  errorMsg.value = ''
  successMsg.value = ''
  loading.value = true
  try {
    if (isForgot.value) {
      if (password.value !== confirmPassword.value) {
        throw new Error('Пароли не совпадают')
      }
      await forgotPassword(email.value, password.value)
      successMsg.value = 'Пароль успешно обновлён. Войдите в аккаунт.'
      isForgot.value = false
    } else if (isLogin.value) {
      await login(email.value, password.value)
      router.push('/cabinet')
    } else {
      await register(username.value, email.value, password.value, fullName.value)
      router.push('/cabinet')
    }
  } catch (e: any) {
    errorMsg.value = e?.data?.error || e?.message || 'Ошибка'
  } finally {
    loading.value = false
  }
}
</script>
<template>
  <div class='min-h-[80vh] flex items-center justify-center px-4'>
    <div class='w-full max-w-md bg-[#16213e] rounded-2xl border border-gray-700/50 p-8'>
      <div class='text-center mb-8'>
        <h1 class='text-3xl font-bold text-[#e2b04a]'>Academy Genius</h1>
        <p class='text-gray-400 mt-2'>
          {{ isForgot ? 'Восстановите пароль' : isLogin ? 'Войдите в аккаунт' : 'Создайте аккаунт' }}
        </p>
      </div>
      <div v-if='errorMsg' class='mb-6 p-3 bg-red-500/10 border border-red-500/30 rounded-lg text-red-400 text-sm'>{{errorMsg}}</div>
      <div v-if='successMsg' class='mb-6 p-3 bg-green-500/10 border border-green-500/30 rounded-lg text-green-400 text-sm'>{{successMsg}}</div>
      <form @submit.prevent='submit' class='space-y-4'>
        <template v-if='!isLogin && !isForgot'>
          <div><label class='block text-gray-300 text-sm mb-1'>Имя пользователя</label><input v-model='username' required minlength='3' class='w-full bg-[#0f0f23] border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:border-[#e2b04a] transition' placeholder='username' /></div>
          <div><label class='block text-gray-300 text-sm mb-1'>Полное имя</label><input v-model='fullName' required class='w-full bg-[#0f0f23] border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:border-[#e2b04a] transition' placeholder='Иван Иванов' /></div>
        </template>
        <div><label class='block text-gray-300 text-sm mb-1'>Email</label><input v-model='email' type='email' required class='w-full bg-[#0f0f23] border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:border-[#e2b04a] transition' placeholder='example@mail.com' /></div>
        <template v-if='!isForgot'>
          <div><label class='block text-gray-300 text-sm mb-1'>Пароль</label><input v-model='password' type='password' required minlength='6' class='w-full bg-[#0f0f23] border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:border-[#e2b04a] transition' placeholder='Минимум 6 символов' /></div>
        </template>
        <template v-else>
          <div><label class='block text-gray-300 text-sm mb-1'>Новый пароль</label><input v-model='password' type='password' required minlength='6' class='w-full bg-[#0f0f23] border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:border-[#e2b04a] transition' placeholder='Минимум 6 символов' /></div>
          <div><label class='block text-gray-300 text-sm mb-1'>Повторите пароль</label><input v-model='confirmPassword' type='password' required minlength='6' class='w-full bg-[#0f0f23] border border-gray-600 rounded-lg px-4 py-2.5 text-white focus:outline-none focus:border-[#e2b04a] transition' placeholder='Повторите пароль' /></div>
        </template>
        <button type='submit' :disabled='loading' class='w-full bg-[#e2b04a] hover:bg-[#f0c75e] text-[#1a1a2e] font-bold py-2.5 rounded-lg transition disabled:opacity-50'>
          {{ loading ? '...' : isForgot ? 'Восстановить пароль' : isLogin ? 'Войти' : 'Зарегистрироваться' }}
        </button>
      </form>
      <div class='text-center mt-6 space-y-2'>
        <button v-if='!isForgot' @click='isLogin = !isLogin; errorMsg = ""; successMsg = ""' class='text-[#e2b04a] hover:underline text-sm'>{{ isLogin ? 'Нет аккаунта? Зарегистрируйтесь' : 'Уже есть аккаунт? Войдите' }}</button>
        <button v-if='!isLogin && !isForgot' @click='isForgot = true; isLogin = false; errorMsg = ""; successMsg = ""' class='text-[#e2b04a] hover:underline text-sm'>Забыли пароль?</button>
        <button v-if='isForgot' @click='isForgot = false; isLogin = true; errorMsg = ""; successMsg = ""' class='text-[#e2b04a] hover:underline text-sm'>Войти</button>
      </div>
    </div>
  </div>
</template>