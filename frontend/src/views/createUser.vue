<template>
    <div class="registration-page">
        <h1>Register</h1>
        <form @submit.prevent="createUsers">
            <div class="form-group">
                <label for="fname">First Name:</label>
                <input
                    type="text"
                    id="fname"
                    v-model="form.fname"
                    required
                />
            </div>
            <div class="form-group">
                <label for="lname">Last Name:</label>
                <input
                    type="text"
                    id="lname"
                    v-model="form.lname"
                    required
                />
            </div>
            <div class="form-group">
                <label for="email">Email:</label>
                <input
                    type="email"
                    id="email"
                    v-model="form.email"
                    required
                />
            </div>
            <div class="form-group">
                <label for="password">Password:</label>
                <input
                    type="password"
                    id="password"
                    v-model="form.password"
                    required
                />
            </div>
            <div class="form-group">
                <label for="role">Role:</label>
                <select id="role" v-model="form.role" required>
                    <option value="user">User</option>
                    <option value="admin">Admin</option>
                </select>
            </div>
            <button type="submit">Register</button>
        </form>
    </div>
</template>

<script>
import { createUser} from "../services/api"; // Adjust the import path as necessary
export default {
    data() {
        return {
            form: {
                fname: "",
                lname: "",
                email: "",
                password: "",
                role: "",
            },
        };
    },
    methods: {
        async createUsers() {
            try {
                // Check if form data is correct before sending
                if (!this.form.fname || !this.form.lname || !this.form.email || !this.form.password || !this.form.role) {
                    alert('All fields are required!');
                    return;
                }

                const response = await createUser(this.form);  // Pass the form directly

                if (!response.ok) {
                    throw new Error("Failed to register user");
                }

                const data = await response.json();
                alert("User registered successfully!");
                console.log(data);
            } catch (error) {
                console.error(error);
                alert("An error occurred while registering the user.");
            }
        },
    },
};

</script>

<style scoped>
.registration-page {
    max-width: 400px;
    margin: 0 auto;
    padding: 20px;
    border: 1px solid #ccc;
    border-radius: 5px;
    background-color: #f9f9f9;
}

.form-group {
    margin-bottom: 15px;
}

label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
}

input,
select {
    width: 100%;
    padding: 8px;
    box-sizing: border-box;
}

button {
    width: 100%;
    padding: 10px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
}

button:hover {
    background-color: #0056b3;
}
</style>
