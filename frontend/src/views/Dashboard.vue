<template>
  <div class="dashboard">
    <header>
      <h1>Welcome, {{ userName }}</h1>
      <button @click="logout" class="logout-button">Logout</button>
    </header>
    
    <main>
      <section class="stats">
        <div class="stat">
          <h2>Users</h2>
          <p>{{ stats.users }}</p>
        </div>
        <div class="stat">
          <h2>Sales</h2>
          <p>{{ formatCurrency(stats.sales) }}</p>
        </div>
        <div class="stat">
          <h2>Orders</h2>
          <p>{{ stats.orders }}</p>
        </div>
      </section>
    </main>
  </div>
</template>

<script>
export default {
  name: 'Dashboard',
  data() {
    return {
      stats: {
        users: 0,
        sales: 0,
        orders: 0
      },
      user: null
    };
  },
  computed: {
    userName() {
      return this.user ? `${this.user.fname} ${this.user.lname}` : 'Guest';
    }
  },
  created() {
    this.checkAuthentication();
    this.fetchDashboardData();
  },
  methods: {
    checkAuthentication() {
      const userData = localStorage.getItem('user');
      if (!userData) {
        this.$router.push({ name: 'Login' });
      } else {
        this.user = JSON.parse(userData);
      }
    },
    async fetchDashboardData() {
      try {
        // Replace with actual API call
        const mockData = {
          users: 152,
          sales: 7250,
          orders: 289
        };
        this.stats = mockData;
      } catch (error) {
        console.error('Error fetching dashboard data:', error);
      }
    },
    formatCurrency(value) {
      return new Intl.NumberFormat('en-US', {
        style: 'currency',
        currency: 'USD'
      }).format(value);
    },
    logout() {
      localStorage.removeItem('user');
      this.$router.push({ name: 'Login' });
    }
  }
};
</script>

<style scoped>
.dashboard {
  font-family: 'Segoe UI', Arial, sans-serif;
  padding: 20px;
  max-width: 1200px;
  margin: 0 auto;
}

header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: #2c3e50;
  color: white;
  padding: 1rem 2rem;
  border-radius: 8px;
  margin-bottom: 2rem;
}

.stats {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
}

.stat {
  background: white;
  padding: 1.5rem;
  border-radius: 10px;
  box-shadow: 0 2px 15px rgba(0, 0, 0, 0.1);
  transition: transform 0.3s ease;
}

.stat:hover {
  transform: translateY(-5px);
}

.stat h2 {
  color: #7f8c8d;
  margin: 0 0 1rem;
  font-size: 1.2rem;
}

.stat p {
  font-size: 2rem;
  font-weight: bold;
  color: #2c3e50;
  margin: 0;
}

.logout-button {
  background: #e74c3c;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 5px;
  cursor: pointer;
  transition: background 0.3s ease;
}

.logout-button:hover {
  background: #c0392b;
}

@media (max-width: 768px) {
  header {
    flex-direction: column;
    gap: 1rem;
    text-align: center;
  }
  
  .stats {
    grid-template-columns: 1fr;
  }
}
</style>