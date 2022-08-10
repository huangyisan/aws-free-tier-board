<template>
  <div class="about">
    <h1>This is cost page</h1>
    <div>
      <v-row>
        <v-col cols="12" sm="6">
          <v-date-picker v-model="dates" range></v-date-picker>
        </v-col>
        <v-col cols="12" sm="6">
          <v-text-field
            v-model="dateRangeText"
            label="Date range"
            prepend-icon="mdi-calendar"
            readonly
          ></v-text-field>
          model: {{ dates }}
        </v-col>
      </v-row>
    </div>
    <div>
      <v-btn block color="primary" @click="submitDate"> Block Button </v-btn>
    </div>
    <div>
      <v-chart class="chart" :option="option" />
    </div>
  </div>
</template>

<script>
// import dateFormat from "./tools.js";
import axios from "axios";
import dateformat from "dateformat";
// import VChart from "vue-echarts";

import { use } from "echarts/core";
import { CanvasRenderer } from "echarts/renderers";
import { PieChart } from "echarts/charts";
import { BarChart } from "echarts/charts";
import "echarts/lib/component/grid";
import "echarts/lib/chart/line";
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
} from "echarts/components";
import VChart, { THEME_KEY } from "vue-echarts";

use([
  CanvasRenderer,
  PieChart,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  BarChart,
]);

export default {
  components: {
    VChart,
  },
  name: "Cost",
  methods: {
    submitDate() {
      console.log(this.dates);
      axios({
        method: "GET",
        url: "http://127.0.0.1:8888/v1/traffic",
        params: {
          start: this.dates[0],
          end: this.dates[1],
        },
      })
        .then((res) => {
          console.log(res.data);
          this.data = res.data;
        })
        .catch((err) => {
          console.error(err);
        });
    },
  },
  data() {
    return {
      dates: [],
      data: [],
    };
  },
  computed: {
    dateRangeText() {
      return this.dates.join(" ~ ");
    },
    option() {
      return {
        xAxis: {
          type: "category",
          data: this.data.map((d) => d.tag),
        },
        yAxis: {
          type: "value",
        },
        series: [
          {
            data: this.data.map((d) => d.value),
            type: "bar",
          },
        ],
      };
    },
  },
  mounted() {
    setDatesRange(this.dates);
    axios({
      method: "GET",
      url: "http://127.0.0.1:8888/v1/traffic",
      params: {
        start: this.dates[0],
        end: this.dates[1],
      },
    })
      .then((res) => {
        console.log(res.data);
        this.data = res.data;
      })
      .catch((err) => {
        console.error(err);
      });
    // axios
    //   .get("http://127.0.0.1:8888/v1/traffic")
    //   .then((res) => {
    //     console.log(res.data);
    //     this.data = res.data;
    //   })
    //   .catch((err) => {
    //     console.error(err);
    //   });

    //
    // setDatesRange(this.dates);
  },
};

function setDatesRange(dates) {
  let today = new Date();
  today = dateformat(today, "yyyy-mm-dd");
  let tomorrow = new Date().setDate(new Date().getDate() + 1);
  tomorrow = dateformat(tomorrow, "yyyy-mm-dd");
  dates.push(today, tomorrow);
}
</script>

<style scoped>
.chart {
  height: 400px;
}
</style>