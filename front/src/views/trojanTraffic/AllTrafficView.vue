<template>
  <div class="about">
    <h1>This is traffic page</h1>
    <div>
      <v-row justify="center">
        <v-col cols="12" sm="5" class="mt-2">
          <v-date-picker v-model="dates" range></v-date-picker>
        </v-col>
        <v-col cols="12" sm="5" class="mt-2">
          <v-text-field
            v-model="dateRangeText"
            label="Date range"
            prepend-icon="mdi-calendar"
            readonly
          ></v-text-field>
        </v-col>
        <v-col cols="12" sm="10" class="mt-2">
          <v-btn block color="primary" @click="submitDate">
            Block Button
          </v-btn>
        </v-col>
      </v-row>
      <v-row justify="center">
        <g-chart :type="type" :data="data" :options="options" />
      </v-row>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import { GChart } from "vue-google-charts/legacy";

import {
  chartType,
  chartData,
  chartOptions,
  genAllTrafficData,
  setDatesRange,
} from "./Traffic.js";

export default {
  components: {
    GChart,
  },
  name: "trojanTraffic",
  methods: {
    getAllTraffic(dates) {
      dates.sort();
      axios({
        method: "GET",
        url: "http://127.0.0.1:8888/v1/traffic",
        params: {
          start: dates[0],
          end: dates[1],
        },
      })
        .then((res) => {
          let tmpdata = [["Endpoint", "download", "upload", "total"]];
          if (res.data !== null) {
            res.data.forEach((element) => {
              genAllTrafficData(element, tmpdata);
            });
            this.data = tmpdata;
          } else {
            tmpdata.push(["null", 0, 0, 0]);
            this.data = tmpdata;
          }
        })
        .catch((err) => {
          console.error(err);
        });
    },
    submitDate() {
      this.getAllTraffic(this.dates);
    },
  },
  data() {
    return {
      type: chartType,
      data: chartData,
      options: chartOptions,
      dates: [],
    };
  },
  computed: {
    dateRangeText() {
      return this.dates.join(" ~ ");
    },
  },
  mounted() {
    setDatesRange(this.dates);
    this.getAllTraffic(this.dates);
  },
};
</script>

<style scoped>
</style>