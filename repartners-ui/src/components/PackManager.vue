<template>
  <v-container class="pa-0">
    <v-row class="header">
      <v-col cols="12" class="text-center">
        <h1>RE PARTNERS TEST</h1>
      </v-col>
    </v-row>
    <v-row class="content">
      <v-col cols="12" md="6">
        <v-card class="card">
          <v-card-title>Set Config</v-card-title>
          <v-card-text>
            <v-text-field
              v-model.trim="configArray"
              label="Enter array of integers separated by commas with no spaces"
              outlined
            ></v-text-field>
            <v-btn @click="updateConfig">Update Config</v-btn>
          </v-card-text>
          <div v-if="latestConfig" class="latest-config">
              <strong>Latest Config:</strong> {{ latestConfig }}
            </div>
        </v-card>
      </v-col>
      <v-col cols="12" md="6">
        <v-card class="card">
          <v-card-title>Get Packs</v-card-title>
          <v-card-text>
            <v-text-field
              v-model.trim="itemId"
              label="Enter number of items"
              outlined
            ></v-text-field>
            <v-btn @click="getItems">Calculate Packs</v-btn>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
    <v-row v-if="items" class="content">
      <v-col cols="12">
        <v-card class="card">
          <v-card-title>Packs</v-card-title>
          <v-card-text>
            <div v-for="(pack, index) in items.packs" :key="index" class="item">
              <p><strong>Pack:</strong> {{ pack.value }} <strong>    Quantity:</strong> {{ pack.quantity }}</p>
              <v-divider v-if="index < items.length - 1"></v-divider>
            </div>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      configArray: '',
      itemId: '',
      items: null,
      latestConfig: null,
    };
  },
  methods: {
    async updateConfig() {
      try {
        const apiUrl = import.meta.env.VITE_API_BASE_URL;
        const configArray = this.configArray.split(',').map(item => parseInt(item.trim(), 10));
        if (configArray.length === 0 || configArray.some(isNaN)) {
          alert('Config array cannot be empty');
          return;
        }
        const response = await axios.put(`${apiUrl}/packs/config`, {
          packs_config: configArray,
        });
        if (response.status === 200) {
          alert('Config updated successfully');
          this.latestConfig = this.configArray;

        } else {
          alert('Failed to update config');
        }
      } catch (error) {
        console.error('Error updating config:', error);
        alert('Failed to update config');
      }
    },
    async getItems() {
      try {
        const apiUrl = import.meta.env.VITE_API_BASE_URL;
        const response = await axios.get(`${apiUrl}/packs/${this.itemId}`);
        this.items = response.data;
      } catch (error) {
        console.error('Error fetching items:', error);
        alert('Failed to fetch items');
      }
    },
  },
};
</script>

<style scoped>
.vcontainer {
  padding: 0;
}
.header {
  background-color: #000;
  color: #ffcc00;
  padding: 20px 0;
}

.content {
  background-color: #1a1a1a;
  padding: 20px 0;
}

.card {
  background-color: #333;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  color: #ffcc00;
}

.text-center {
  text-align: center;
}

h1 {
  font-family: 'Arial', sans-serif;
  font-size: 2.5rem;
  color: #ffcc00;
}

.v-card-title {
  font-family: 'Arial', sans-serif;
  font-size: 1.5rem;
  color: #ffcc00;
}

.v-card-text {
  font-family: 'Arial', sans-serif;
  font-size: 1rem;
  color: #ffcc00;
}

.v-btn {
  background-color: #ffcc00;
  color: #000;
  font-family: 'Arial', sans-serif;
  font-size: 1rem;
  font-weight: bold;
}
.latest-config {
  margin-top: 10px;
  margin-left: 15px;
  font-family: 'Arial', sans-serif;
  font-size: 1rem;
  color: #ffcc00;
}

</style>