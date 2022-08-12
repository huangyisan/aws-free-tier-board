<template>
  <div class="about">
    <!-- <h1>This is filter traffic page</h1> -->
    <div>
      <v-row justify="center">
        <v-col cols="12" sm="4" class="ma-2">
          <v-date-picker v-model="dates" range></v-date-picker>
        </v-col>
        <v-col cols="12" sm="4" class="ma-2">
          <v-card elevation="2">
            <v-card-title>Tag Filter</v-card-title>
            <v-container fluid>
              <v-row align="center">
                <v-col class="d-flex" cols="12" sm="6">
                  <v-select
                    :items="groups"
                    label="Group"
                    @change="getTagList"
                  ></v-select>
                </v-col>

                <v-col class="d-flex" cols="12" sm="6">
                  <v-select
                    :items="tags"
                    label="Tag"
                    @change="setTag"
                  ></v-select>
                </v-col>
              </v-row>
              <v-btn block color="primary" @click="submitTagFilter">
                Tag Filter
              </v-btn>
            </v-container>
          </v-card>

          <v-card elevation="2" class="mt-2">
            <v-card-title>Group Filter</v-card-title>
            <v-container fluid>
              <v-row align="center">
                <v-col class="d-flex" cols="12" sm="12">
                  <v-select
                    :items="groups"
                    label="Group"
                    @change="setGroup"
                  ></v-select>
                </v-col>
              </v-row>
              <v-btn block color="primary" @click="submitGroupFilter">
                Filter
              </v-btn>
            </v-container>
          </v-card>
        </v-col>
      </v-row>

      <v-col cols="12" sm="12">
        <g-chart :type="type" :data="data" :options="options" />
      </v-col>
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
  genTrafficDataByTag,
  genTrafficDataByGroup,
  setDatesRange,
} from "./Traffic.js";

export default {
  components: {
    GChart,
  },
  name: "trojanTraffic",
  methods: {
    setTag(value) {
      this.currentTag = value;
    },
    setGroup(value) {
      this.currentGroup = value;
    },
    getTagList(value) {
      axios({
        method: "GET",
        url: `http://127.0.0.1:8888/v1/servers/group/${value}`,
      })
        .then((res) => {
          this.tags = res.data;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    getGroupList() {
      axios({
        method: "GET",
        url: "http://127.0.0.1:8888/v1/servers/group",
      })
        .then((res) => {
          this.groups = res.data;
        })
        .catch((err) => {
          console.error(err);
        });
    },
    getTagTraffic(dates) {
      dates.sort();
      axios({
        method: "GET",
        url: `http://127.0.0.1:8888/v1/traffic/tag/${this.currentTag}`,
        params: {
          start: dates[0],
          end: dates[1],
        },
      })
        .then((res) => {
          let tmpdata = [["Date", "download", "upload", "total"]];
          if (res.data !== null) {
            res.data.forEach((element) => {
              genTrafficDataByTag(element, tmpdata);
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
    getGroupTraffic(dates) {
      dates.sort();
      axios({
        method: "GET",
        url: `http://127.0.0.1:8888/v1/traffic/group/${this.currentGroup}`,
        params: {
          start: dates[0],
          end: dates[1],
        },
      })
        .then((res) => {
          let tmpdata = [["Date", "download", "upload", "total"]];
          if (res.data !== null) {
            res.data.forEach((element) => {
              genTrafficDataByGroup(element, tmpdata);
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
    submitTagFilter() {
      this.getTagTraffic(this.dates);
    },
    submitGroupFilter() {
      this.getGroupTraffic(this.dates);
    },
  },
  data() {
    return {
      type: chartType,
      data: chartData,
      options: chartOptions,
      dates: [],
      items: ["Foo", "Bar", "Fizz", "Buzz"],
      tags: [],
      groups: [],
      currentTag: "",
      currentGroup: "",
    };
  },
  computed: {
    dateRangeText() {
      return this.dates.join(" ~ ");
    },
  },
  mounted() {
    this.getGroupList();
    setDatesRange(this.dates);
    // this.getAllTraffic(this.dates);
  },
};
</script>

<style scoped>
</style>