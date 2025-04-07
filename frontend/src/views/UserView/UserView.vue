<template>
    <div class="user-view">
        <header class="user-header">
                <h1>Welcome, {{ userName }}</h1>
                <button @click="logout" class="logout-button">Logout</button>
        </header>
        <main class="user-content">
            <section class="user-section">
                <h2>Manage Users</h2>
                <p>View, edit, or delete user accounts.</p>
                <button @click="navigateTo('users')">Go to Users</button>
            </section>
            <section class="user-section">
                <h2>Manage Posts</h2>
                <p>Review and moderate posts.</p>
                <button @click="navigateTo('posts')">Go to Posts</button>
            </section>
            <section class="user-section">
                <h2>Settings</h2>
                <p>Update application settings.</p>
                <button @click="navigateTo('settings')">Go to Settings</button>
            </section>
        </main>
    </div>
</template>

<script>
export default {
  name: "UserView",
  data() {
    return {
      user: JSON.parse(localStorage.getItem('user')) || null,
    };
  },
  computed: {
    userName() {
      return this.user ? `${this.user.fname} ${this.user.lname}` : 'Guest';
    }
  },
  methods: {
    navigateTo(page) {
      this.$router.push({ name: page });
    },
    logout() {
      localStorage.removeItem('user');
      this.$router.push({ name: 'Login' });
    }
  }
};
</script>

<style scoped>
.user-view {
    font-family: Arial, sans-serif;
    padding: 20px;
}

.user-header {
    background-color: #4c82e6;
    color: white;
    padding: 10px;
    text-align: center;
}

.user-content {
    display: flex;
    flex-direction: column;
    gap: 20px;
    margin-top: 20px;
}

.user-section {
    border: 1px solid #ddd;
    padding: 15px;
    border-radius: 5px;
    background-color: #f9f9f9;
}

.user-section h2 {
    margin: 0 0 10px;
}

button {
    background-color: #4CAF50;
    color: white;
    border: none;
    padding: 10px 15px;
    border-radius: 5px;
    cursor: pointer;
}

button:hover {
    background-color: #b5ca2b;
}
</style>