<template>
  <div class="login-container">
    <h2>Login</h2>
    <form @submit.prevent="login">
      <div>
        <label for="email">Email:</label>
        <input type="email" v-model="email" required />
      </div>
      <div>
        <label for="password">Password:</label>
        <input type="password" v-model="password" required />
      </div>
      <button type="submit">Login</button>
      <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
    </form>
  </div>
</template>

<script>
import { loginUser } from '../services/api';

export default {
  data() {
    return {
      email: '',
      password: '',
      errorMessage: ''
    };
  },
  methods: {
    async login() {
      this.errorMessage = '';
      
      try {
        const response = await loginUser(this.email, this.password);
        console.log("Login response:", response);

        if (response.success) {
          // Store user data excluding sensitive information
          const user = response.data;
          localStorage.setItem('user', JSON.stringify(user));
          
          if (user.role === "admin") {
          this.$router.push({ name: 'AdminView' });
        } else {
          this.$router.push({ name: 'UserView' });
        }
        } 
        
      } catch (error) {
        this.errorMessage = 'Login failed. Please try again later.';
        console.error("Login error:", error);
      }
    }
  }
};
</script>

<style scoped>
.login-container {
  width: 300px;
  margin: 2rem auto;
  padding: 2rem;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

form div {
  margin-bottom: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
}

input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 0.75rem;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}

.error {
  color: #dc3545;
  margin-top: 1rem;
  text-align: center;
}
</style>